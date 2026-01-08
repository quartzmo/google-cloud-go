# Scoped TracerProvider for Go Client Libraries

| |
| :---: |

Author: chrisdsmith@google.com | Last Updated: 2026-01-02 | Status: Draft
Project: Cloud SDK Go Tracing | Self link: N/A

# Objective

Enable Google Cloud Go client libraries to accept a user-provided `trace.TracerProvider` scoped to a specific client instance. This allows for advanced telemetry configurations (e.g., multi-tenant applications, separate logging projects, custom sampling per client) that override the global OpenTelemetry default.

# Background

Currently, the `grpctransport` and `httptransport` layers in `google-cloud-go/auth` implicitly rely on `otel.GetTracerProvider()` (the global provider) when initializing OpenTelemetry instrumentation. This limits flexibility for users who need isolated telemetry pipelines for different clients within the same process.

# Detailed Design

## 1. API Layer (`google-api-go-client`)
**Owner:** `google-api-go-client/option`

We need a standard mechanism for users to pass a `TracerProvider` via `ClientOptions`.

*   **New Option:** Add `func WithTracerProvider(tp trace.TracerProvider) ClientOption` to `google.golang.org/api/option`.
*   **Storage:** Add a `TracerProvider trace.TracerProvider` field to `internal.DialSettings` in `google.golang.org/api/internal`.
*   **Application:** The `WithTracerProvider` option's `Apply` method will populate this field in `DialSettings`.

## 2. Transport Bridge
**Owner:** `google-api-go-client/transport`

The transport dialers must propagate this configuration from the generated client options to the underlying authentication transport layer.

*   **gRPC:** Update `dialPoolNewAuth` in `transport/grpc/dial.go` to read `ds.TracerProvider` and pass it to `grpctransport.Options`.
*   **HTTP:** Update `newClientNewAuth` in `transport/http/dial.go` to read `ds.TracerProvider` and pass it to `httptransport.Options`.

## 3. Transport Layer (`google-cloud-go/auth`)
**Owner:** `cloud.google.com/go/auth`

The transport libraries must accept and apply the provider.

*   **Configuration:** Add `TracerProvider trace.TracerProvider` to the public `Options` struct in both `grpctransport` and `httptransport`.

### gRPC Implementation (`grpctransport`)
*   In `addOpenTelemetryStatsHandler`:
    *   Check if `opts.TracerProvider` is non-nil.
    *   If set, append `otelgrpc.WithTracerProvider(opts.TracerProvider)` to the `otelgrpc` options.
    *   If nil, `otelgrpc` defaults to the global provider (preserving backward compatibility).

### HTTP Implementation (`httptransport`)
*   In `addOpenTelemetryTransport`:
    *   Check if `opts.TracerProvider` is non-nil.
    *   If set, pass `otelhttp.WithTracerProvider(opts.TracerProvider)` to `otelhttp.NewTransport`.
    *   If nil, `otelhttp` defaults to the global provider.

# Validation Plan

1.  **Unit Tests:**
    *   Create tests in `option` package verifying `WithTracerProvider` sets the `DialSettings` field.
    *   Create tests in `auth` transport packages using a mock/test `TracerProvider`. Verify that spans are created using the *mock* provider and not the global one.
2.  **Integration:**
    *   Verify a generated client configured with a custom provider exports spans to the expected destination (and not to the destination configured globally, if different).
