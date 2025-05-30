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

package accounts

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"math"
	"net/http"
	"net/url"
	"time"

	accountspb "cloud.google.com/go/shopping/merchant/accounts/apiv1beta/accountspb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/encoding/protojson"
)

var newAutofeedSettingsClientHook clientHook

// AutofeedSettingsCallOptions contains the retry settings for each method of AutofeedSettingsClient.
type AutofeedSettingsCallOptions struct {
	GetAutofeedSettings    []gax.CallOption
	UpdateAutofeedSettings []gax.CallOption
}

func defaultAutofeedSettingsGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("merchantapi.googleapis.com:443"),
		internaloption.WithDefaultEndpointTemplate("merchantapi.UNIVERSE_DOMAIN:443"),
		internaloption.WithDefaultMTLSEndpoint("merchantapi.mtls.googleapis.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://merchantapi.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		internaloption.EnableNewAuthLibrary(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultAutofeedSettingsCallOptions() *AutofeedSettingsCallOptions {
	return &AutofeedSettingsCallOptions{
		GetAutofeedSettings: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		UpdateAutofeedSettings: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

func defaultAutofeedSettingsRESTCallOptions() *AutofeedSettingsCallOptions {
	return &AutofeedSettingsCallOptions{
		GetAutofeedSettings: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
		UpdateAutofeedSettings: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
	}
}

// internalAutofeedSettingsClient is an interface that defines the methods available from Merchant API.
type internalAutofeedSettingsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetAutofeedSettings(context.Context, *accountspb.GetAutofeedSettingsRequest, ...gax.CallOption) (*accountspb.AutofeedSettings, error)
	UpdateAutofeedSettings(context.Context, *accountspb.UpdateAutofeedSettingsRequest, ...gax.CallOption) (*accountspb.AutofeedSettings, error)
}

// AutofeedSettingsClient is a client for interacting with Merchant API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to support
// autofeed (at https://support.google.com/merchants/answer/7538732) setting.
type AutofeedSettingsClient struct {
	// The internal transport-dependent client.
	internalClient internalAutofeedSettingsClient

	// The call options for this service.
	CallOptions *AutofeedSettingsCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *AutofeedSettingsClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *AutofeedSettingsClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *AutofeedSettingsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetAutofeedSettings retrieves the autofeed settings of an account.
func (c *AutofeedSettingsClient) GetAutofeedSettings(ctx context.Context, req *accountspb.GetAutofeedSettingsRequest, opts ...gax.CallOption) (*accountspb.AutofeedSettings, error) {
	return c.internalClient.GetAutofeedSettings(ctx, req, opts...)
}

// UpdateAutofeedSettings updates the autofeed settings of an account.
func (c *AutofeedSettingsClient) UpdateAutofeedSettings(ctx context.Context, req *accountspb.UpdateAutofeedSettingsRequest, opts ...gax.CallOption) (*accountspb.AutofeedSettings, error) {
	return c.internalClient.UpdateAutofeedSettings(ctx, req, opts...)
}

// autofeedSettingsGRPCClient is a client for interacting with Merchant API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type autofeedSettingsGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing AutofeedSettingsClient
	CallOptions **AutofeedSettingsCallOptions

	// The gRPC API client.
	autofeedSettingsClient accountspb.AutofeedSettingsServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string

	logger *slog.Logger
}

// NewAutofeedSettingsClient creates a new autofeed settings service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to support
// autofeed (at https://support.google.com/merchants/answer/7538732) setting.
func NewAutofeedSettingsClient(ctx context.Context, opts ...option.ClientOption) (*AutofeedSettingsClient, error) {
	clientOpts := defaultAutofeedSettingsGRPCClientOptions()
	if newAutofeedSettingsClientHook != nil {
		hookOpts, err := newAutofeedSettingsClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := AutofeedSettingsClient{CallOptions: defaultAutofeedSettingsCallOptions()}

	c := &autofeedSettingsGRPCClient{
		connPool:               connPool,
		autofeedSettingsClient: accountspb.NewAutofeedSettingsServiceClient(connPool),
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
func (c *autofeedSettingsGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *autofeedSettingsGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version, "pb", protoVersion)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *autofeedSettingsGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type autofeedSettingsRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing AutofeedSettingsClient
	CallOptions **AutofeedSettingsCallOptions

	logger *slog.Logger
}

// NewAutofeedSettingsRESTClient creates a new autofeed settings service rest client.
//
// Service to support
// autofeed (at https://support.google.com/merchants/answer/7538732) setting.
func NewAutofeedSettingsRESTClient(ctx context.Context, opts ...option.ClientOption) (*AutofeedSettingsClient, error) {
	clientOpts := append(defaultAutofeedSettingsRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultAutofeedSettingsRESTCallOptions()
	c := &autofeedSettingsRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
		logger:      internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	return &AutofeedSettingsClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultAutofeedSettingsRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://merchantapi.googleapis.com"),
		internaloption.WithDefaultEndpointTemplate("https://merchantapi.UNIVERSE_DOMAIN"),
		internaloption.WithDefaultMTLSEndpoint("https://merchantapi.mtls.googleapis.com"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://merchantapi.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableNewAuthLibrary(),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *autofeedSettingsRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN", "pb", protoVersion)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *autofeedSettingsRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *autofeedSettingsRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *autofeedSettingsGRPCClient) GetAutofeedSettings(ctx context.Context, req *accountspb.GetAutofeedSettingsRequest, opts ...gax.CallOption) (*accountspb.AutofeedSettings, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetAutofeedSettings[0:len((*c.CallOptions).GetAutofeedSettings):len((*c.CallOptions).GetAutofeedSettings)], opts...)
	var resp *accountspb.AutofeedSettings
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.autofeedSettingsClient.GetAutofeedSettings, req, settings.GRPC, c.logger, "GetAutofeedSettings")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *autofeedSettingsGRPCClient) UpdateAutofeedSettings(ctx context.Context, req *accountspb.UpdateAutofeedSettingsRequest, opts ...gax.CallOption) (*accountspb.AutofeedSettings, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "autofeed_settings.name", url.QueryEscape(req.GetAutofeedSettings().GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).UpdateAutofeedSettings[0:len((*c.CallOptions).UpdateAutofeedSettings):len((*c.CallOptions).UpdateAutofeedSettings)], opts...)
	var resp *accountspb.AutofeedSettings
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.autofeedSettingsClient.UpdateAutofeedSettings, req, settings.GRPC, c.logger, "UpdateAutofeedSettings")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetAutofeedSettings retrieves the autofeed settings of an account.
func (c *autofeedSettingsRESTClient) GetAutofeedSettings(ctx context.Context, req *accountspb.GetAutofeedSettingsRequest, opts ...gax.CallOption) (*accountspb.AutofeedSettings, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/accounts/v1beta/%v", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).GetAutofeedSettings[0:len((*c.CallOptions).GetAutofeedSettings):len((*c.CallOptions).GetAutofeedSettings)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &accountspb.AutofeedSettings{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "GetAutofeedSettings")
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

// UpdateAutofeedSettings updates the autofeed settings of an account.
func (c *autofeedSettingsRESTClient) UpdateAutofeedSettings(ctx context.Context, req *accountspb.UpdateAutofeedSettingsRequest, opts ...gax.CallOption) (*accountspb.AutofeedSettings, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetAutofeedSettings()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/accounts/v1beta/%v", req.GetAutofeedSettings().GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")
	if req.GetUpdateMask() != nil {
		field, err := protojson.Marshal(req.GetUpdateMask())
		if err != nil {
			return nil, err
		}
		params.Add("updateMask", string(field[1:len(field)-1]))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "autofeed_settings.name", url.QueryEscape(req.GetAutofeedSettings().GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).UpdateAutofeedSettings[0:len((*c.CallOptions).UpdateAutofeedSettings):len((*c.CallOptions).UpdateAutofeedSettings)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &accountspb.AutofeedSettings{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("PATCH", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, jsonReq, "UpdateAutofeedSettings")
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
