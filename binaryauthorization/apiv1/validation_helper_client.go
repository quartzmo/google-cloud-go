// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package binaryauthorization

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"math"
	"net/http"
	"net/url"

	binaryauthorizationpb "cloud.google.com/go/binaryauthorization/apiv1/binaryauthorizationpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

var newValidationHelperClientHook clientHook

// ValidationHelperCallOptions contains the retry settings for each method of ValidationHelperClient.
type ValidationHelperCallOptions struct {
	ValidateAttestationOccurrence []gax.CallOption
}

func defaultValidationHelperGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("binaryauthorization.googleapis.com:443"),
		internaloption.WithDefaultEndpointTemplate("binaryauthorization.UNIVERSE_DOMAIN:443"),
		internaloption.WithDefaultMTLSEndpoint("binaryauthorization.mtls.googleapis.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://binaryauthorization.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		internaloption.EnableNewAuthLibrary(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultValidationHelperCallOptions() *ValidationHelperCallOptions {
	return &ValidationHelperCallOptions{
		ValidateAttestationOccurrence: []gax.CallOption{},
	}
}

func defaultValidationHelperRESTCallOptions() *ValidationHelperCallOptions {
	return &ValidationHelperCallOptions{
		ValidateAttestationOccurrence: []gax.CallOption{},
	}
}

// internalValidationHelperClient is an interface that defines the methods available from Binary Authorization API.
type internalValidationHelperClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ValidateAttestationOccurrence(context.Context, *binaryauthorizationpb.ValidateAttestationOccurrenceRequest, ...gax.CallOption) (*binaryauthorizationpb.ValidateAttestationOccurrenceResponse, error)
}

// ValidationHelperClient is a client for interacting with Binary Authorization API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// BinAuthz Attestor verification
type ValidationHelperClient struct {
	// The internal transport-dependent client.
	internalClient internalValidationHelperClient

	// The call options for this service.
	CallOptions *ValidationHelperCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ValidationHelperClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ValidationHelperClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *ValidationHelperClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ValidateAttestationOccurrence returns whether the given Attestation for the given image URI
// was signed by the given Attestor
func (c *ValidationHelperClient) ValidateAttestationOccurrence(ctx context.Context, req *binaryauthorizationpb.ValidateAttestationOccurrenceRequest, opts ...gax.CallOption) (*binaryauthorizationpb.ValidateAttestationOccurrenceResponse, error) {
	return c.internalClient.ValidateAttestationOccurrence(ctx, req, opts...)
}

// validationHelperGRPCClient is a client for interacting with Binary Authorization API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type validationHelperGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing ValidationHelperClient
	CallOptions **ValidationHelperCallOptions

	// The gRPC API client.
	validationHelperClient binaryauthorizationpb.ValidationHelperV1Client

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string

	logger *slog.Logger
}

// NewValidationHelperClient creates a new validation helper v1 client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// BinAuthz Attestor verification
func NewValidationHelperClient(ctx context.Context, opts ...option.ClientOption) (*ValidationHelperClient, error) {
	clientOpts := defaultValidationHelperGRPCClientOptions()
	if newValidationHelperClientHook != nil {
		hookOpts, err := newValidationHelperClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := ValidationHelperClient{CallOptions: defaultValidationHelperCallOptions()}

	c := &validationHelperGRPCClient{
		connPool:               connPool,
		validationHelperClient: binaryauthorizationpb.NewValidationHelperV1Client(connPool),
		CallOptions:            &client.CallOptions,
		logger:                 internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *validationHelperGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *validationHelperGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version, "pb", protoVersion)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *validationHelperGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type validationHelperRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing ValidationHelperClient
	CallOptions **ValidationHelperCallOptions

	logger *slog.Logger
}

// NewValidationHelperRESTClient creates a new validation helper v1 rest client.
//
// BinAuthz Attestor verification
func NewValidationHelperRESTClient(ctx context.Context, opts ...option.ClientOption) (*ValidationHelperClient, error) {
	clientOpts := append(defaultValidationHelperRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultValidationHelperRESTCallOptions()
	c := &validationHelperRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
		logger:      internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	return &ValidationHelperClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultValidationHelperRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://binaryauthorization.googleapis.com"),
		internaloption.WithDefaultEndpointTemplate("https://binaryauthorization.UNIVERSE_DOMAIN"),
		internaloption.WithDefaultMTLSEndpoint("https://binaryauthorization.mtls.googleapis.com"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://binaryauthorization.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableNewAuthLibrary(),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *validationHelperRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN", "pb", protoVersion)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *validationHelperRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *validationHelperRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *validationHelperGRPCClient) ValidateAttestationOccurrence(ctx context.Context, req *binaryauthorizationpb.ValidateAttestationOccurrenceRequest, opts ...gax.CallOption) (*binaryauthorizationpb.ValidateAttestationOccurrenceResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "attestor", url.QueryEscape(req.GetAttestor()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ValidateAttestationOccurrence[0:len((*c.CallOptions).ValidateAttestationOccurrence):len((*c.CallOptions).ValidateAttestationOccurrence)], opts...)
	var resp *binaryauthorizationpb.ValidateAttestationOccurrenceResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.validationHelperClient.ValidateAttestationOccurrence, req, settings.GRPC, c.logger, "ValidateAttestationOccurrence")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ValidateAttestationOccurrence returns whether the given Attestation for the given image URI
// was signed by the given Attestor
func (c *validationHelperRESTClient) ValidateAttestationOccurrence(ctx context.Context, req *binaryauthorizationpb.ValidateAttestationOccurrenceRequest, opts ...gax.CallOption) (*binaryauthorizationpb.ValidateAttestationOccurrenceResponse, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v:validateAttestationOccurrence", req.GetAttestor())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "attestor", url.QueryEscape(req.GetAttestor()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).ValidateAttestationOccurrence[0:len((*c.CallOptions).ValidateAttestationOccurrence):len((*c.CallOptions).ValidateAttestationOccurrence)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &binaryauthorizationpb.ValidateAttestationOccurrenceResponse{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, jsonReq, "ValidateAttestationOccurrence")
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}
