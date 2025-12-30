# Implementation Plan: OpenTelemetry Attribute Enrichment

This plan outlines the steps required to implement the design documented in `design.md`.

**Global Requirement:** All changes must be gated by calling `gax.IsFeatureEnabled("TRACING")` to check for the presence of `GOOGLE_SDK_GO_EXPERIMENTAL_TRACING=true` (value case-insensitive).

## Phase 1: Transport Layer (`google-cloud-go/auth`)
**Goal:** Enable `grpctransport` and `httptransport` to inject static attributes, capture dynamic attributes from context/metadata, and strictly enforce error attribute semantics.

1.  **Modify `grpctransport/grpctransport.go`:**
    *   Add `TelemetryAttributes map[string]string` to `InternalOptions`.
    *   Implement `otelHandler` struct wrapping `stats.Handler`.
    *   **TagRPC (Start):**
        *   Call underlying `otelgrpc` handler.
        *   Extract "gcp.resource.name" and "gcp.grpc.resend_count" from `metadata.FromOutgoingContext`.
        *   Enrich the created span with these attributes.
    *   **HandleRPC (End):**
        *   Implement `HandleRPC` method.
        *   Inspect `stats.End` for errors.
        *   If error exists, set the following attributes ensuring **UPPER_CASE** values for codes:
            *   `error.type`: e.g., "UNAVAILABLE".
            *   `grpc.status`: e.g., "UNAVAILABLE".
            *   `status.message`: The raw error message string.
            *   `exception.type`: The Go type of the error.
            *   `url.domain`: Extract from `:authority` header.
    *   Update `addOpenTelemetryStatsHandler` to use this wrapper.
    *   **Check Feature Flag:** Only install `otelHandler` wrapper if `gax.IsFeatureEnabled("TRACING")` is true. Otherwise, use standard `otelgrpc` handler.

2.  **Modify `httptransport/httptransport.go`:**
    *   Add `TelemetryAttributes map[string]string` to `InternalOptions`.
    *   Implement a custom `http.RoundTripper` wrapper (`otelAttributeTransport`).
    *   **RoundTrip:**
        *   Retrieve the span via `trace.SpanFromContext`.
        *   Extracts "gcp.resource.name" and "gcp.grpc.resend_count" (or `http.request.resend_count`) from `metadata.FromOutgoingContext(req.Context())`.
        *   Enrich the span with static and dynamic attributes.
        *   Call `base.RoundTrip`.
        *   **On Return (Error Handling):** Inspect `err` and `resp`.
        *   If `err` is not nil or `resp.StatusCode` >= 400:
            *   Set `error.type` (e.g., "503" or error class).
            *   Set `status.message`.
            *   Set `rpc.system` to "http".
            *   Set `url.domain` from the request URL host.
            *   Set `exception.type` from the error.
    *   Update `NewClient` (or `addOpenTelemetryTransport`) to use this wrapper.
    *   **Check Feature Flag:** Only install `otelAttributeTransport` wrapper if `gax.IsFeatureEnabled("TRACING")` is true. Otherwise, use standard `otelhttp` transport.

3.  **Verify with Tests (`*_otel_test.go`)**:
    *   Add test cases ensuring `TelemetryAttributes` appear on spans (gRPC and HTTP) **when enabled**.
    *   Add test cases ensuring metadata/context injection works **when enabled**.
    *   Add test cases ensuring error attributes are correctly populated on failure, verifying UPPER_CASE formatting for gRPC codes.
    *   **Gate Coverage:** Verify that legacy behavior is preserved when the feature flag is disabled.

## Phase 1b: API Option Bridge (`google-api-go-client`)
**Goal:** Allow the generated clients to pass telemetry metadata through the standard `option` package to the `auth` library.

1.  **Modify `internal/settings.go`**:
    *   Add `TelemetryAttributes map[string]string` field to `DialSettings` struct.
2.  **Modify `option/internaloption/internaloption.go`**:
    *   Implement `WithTelemetryAttributes(map[string]string) option.ClientOption`.
    *   Implement the `Apply` method to populate `DialSettings.TelemetryAttributes`.
3.  **Modify `transport/grpc/dial.go`**:
    *   Update `dialPoolNewAuth` to read `ds.TelemetryAttributes` and pass it to `grpctransport.InternalOptions`.
4.  **Modify `transport/http/dial.go`**:
    *   Update `newClientNewAuth` to read `ds.TelemetryAttributes` and pass it to `httptransport.InternalOptions`.

## Phase 2: Retry Logic (`gax-go`)
**Goal:** Expose the retry attempt count to the transport layer.

1.  **Update `Invoke` Logic**:
    *   Modify `gax.Invoke` to track the attempt count.
    *   **Check Feature Flag:** If `gax.IsFeatureEnabled("TRACING")`, inject this count into the `context` using `metadata.AppendToOutgoingContext(ctx, "gcp.grpc.resend_count", count)`.
    *   Ensure this is done in a backward-compatible way (no signature changes).

## Phase 3: Code Generation (`gapic-generator-go`)
**Goal:** Automate the extraction of resource names and static identity.

1.  **Update Templates**:
    *   **Service Identity:** Modify `grpcClientOptions` (gRPC) and `restClientOptions` (REST).
        *   **Check Feature Flag:** Generate logic to call `internaloption.WithTelemetryAttributes` *only if* `gax.IsFeatureEnabled("TRACING")` in client code.
    *   **Resource Name:** Modify `insertRequestHeaders` in `gengapic.go`.
        *   Reuse the existing logic that parses `google.api.http` and `google.api.routing` annotations to identify the primary resource field (e.g., `req.Name`).
        *   **Check Feature Flag:** Generate code to inject this value into the context metadata using `metadata.AppendToOutgoingContext` *only if* `gax.IsFeatureEnabled("TRACING")` in client code.

## Phase 4: Integration & Rollout
1.  **Release `google-cloud-go/auth`**: Publish the changes from Phase 1.
2.  **Release `google-api-go-client`**: Publish the changes from Phase 1b.
3.  **Release `gax-go`**: Publish the changes from Phase 2.
4.  **Release `gapic-generator-go`**: Publish the updated generator.
5.  **Regenerate Clients**: Run the new generator against Google Cloud Go client libraries to pick up the changes.