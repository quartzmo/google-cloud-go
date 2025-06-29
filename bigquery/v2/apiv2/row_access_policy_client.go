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

package bigquery

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"math"
	"net/http"
	"net/url"
	"time"

	bigquerypb "cloud.google.com/go/bigquery/v2/apiv2/bigquerypb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var newRowAccessPolicyClientHook clientHook

// RowAccessPolicyCallOptions contains the retry settings for each method of RowAccessPolicyClient.
type RowAccessPolicyCallOptions struct {
	ListRowAccessPolicies        []gax.CallOption
	GetRowAccessPolicy           []gax.CallOption
	CreateRowAccessPolicy        []gax.CallOption
	UpdateRowAccessPolicy        []gax.CallOption
	DeleteRowAccessPolicy        []gax.CallOption
	BatchDeleteRowAccessPolicies []gax.CallOption
}

func defaultRowAccessPolicyGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("bigquery.googleapis.com:443"),
		internaloption.WithDefaultEndpointTemplate("bigquery.UNIVERSE_DOMAIN:443"),
		internaloption.WithDefaultMTLSEndpoint("bigquery.mtls.googleapis.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://bigquery.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		internaloption.EnableNewAuthLibrary(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultRowAccessPolicyCallOptions() *RowAccessPolicyCallOptions {
	return &RowAccessPolicyCallOptions{
		ListRowAccessPolicies: []gax.CallOption{
			gax.WithTimeout(64000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
					codes.ResourceExhausted,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetRowAccessPolicy: []gax.CallOption{
			gax.WithTimeout(64000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
					codes.ResourceExhausted,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		CreateRowAccessPolicy: []gax.CallOption{
			gax.WithTimeout(64000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
					codes.ResourceExhausted,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		UpdateRowAccessPolicy: []gax.CallOption{
			gax.WithTimeout(64000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
					codes.ResourceExhausted,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		DeleteRowAccessPolicy: []gax.CallOption{
			gax.WithTimeout(64000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
					codes.ResourceExhausted,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		BatchDeleteRowAccessPolicies: []gax.CallOption{
			gax.WithTimeout(64000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
					codes.ResourceExhausted,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

func defaultRowAccessPolicyRESTCallOptions() *RowAccessPolicyCallOptions {
	return &RowAccessPolicyCallOptions{
		ListRowAccessPolicies: []gax.CallOption{
			gax.WithTimeout(64000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusGatewayTimeout,
					http.StatusServiceUnavailable,
					http.StatusTooManyRequests)
			}),
		},
		GetRowAccessPolicy: []gax.CallOption{
			gax.WithTimeout(64000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusGatewayTimeout,
					http.StatusServiceUnavailable,
					http.StatusTooManyRequests)
			}),
		},
		CreateRowAccessPolicy: []gax.CallOption{
			gax.WithTimeout(64000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusGatewayTimeout,
					http.StatusServiceUnavailable,
					http.StatusTooManyRequests)
			}),
		},
		UpdateRowAccessPolicy: []gax.CallOption{
			gax.WithTimeout(64000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusGatewayTimeout,
					http.StatusServiceUnavailable,
					http.StatusTooManyRequests)
			}),
		},
		DeleteRowAccessPolicy: []gax.CallOption{
			gax.WithTimeout(64000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusGatewayTimeout,
					http.StatusServiceUnavailable,
					http.StatusTooManyRequests)
			}),
		},
		BatchDeleteRowAccessPolicies: []gax.CallOption{
			gax.WithTimeout(64000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusGatewayTimeout,
					http.StatusServiceUnavailable,
					http.StatusTooManyRequests)
			}),
		},
	}
}

// internalRowAccessPolicyClient is an interface that defines the methods available from BigQuery API.
type internalRowAccessPolicyClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListRowAccessPolicies(context.Context, *bigquerypb.ListRowAccessPoliciesRequest, ...gax.CallOption) *RowAccessPolicyIterator
	GetRowAccessPolicy(context.Context, *bigquerypb.GetRowAccessPolicyRequest, ...gax.CallOption) (*bigquerypb.RowAccessPolicy, error)
	CreateRowAccessPolicy(context.Context, *bigquerypb.CreateRowAccessPolicyRequest, ...gax.CallOption) (*bigquerypb.RowAccessPolicy, error)
	UpdateRowAccessPolicy(context.Context, *bigquerypb.UpdateRowAccessPolicyRequest, ...gax.CallOption) (*bigquerypb.RowAccessPolicy, error)
	DeleteRowAccessPolicy(context.Context, *bigquerypb.DeleteRowAccessPolicyRequest, ...gax.CallOption) error
	BatchDeleteRowAccessPolicies(context.Context, *bigquerypb.BatchDeleteRowAccessPoliciesRequest, ...gax.CallOption) error
}

// RowAccessPolicyClient is a client for interacting with BigQuery API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service for interacting with row access policies.
type RowAccessPolicyClient struct {
	// The internal transport-dependent client.
	internalClient internalRowAccessPolicyClient

	// The call options for this service.
	CallOptions *RowAccessPolicyCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *RowAccessPolicyClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *RowAccessPolicyClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *RowAccessPolicyClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ListRowAccessPolicies lists all row access policies on the specified table.
func (c *RowAccessPolicyClient) ListRowAccessPolicies(ctx context.Context, req *bigquerypb.ListRowAccessPoliciesRequest, opts ...gax.CallOption) *RowAccessPolicyIterator {
	return c.internalClient.ListRowAccessPolicies(ctx, req, opts...)
}

// GetRowAccessPolicy gets the specified row access policy by policy ID.
func (c *RowAccessPolicyClient) GetRowAccessPolicy(ctx context.Context, req *bigquerypb.GetRowAccessPolicyRequest, opts ...gax.CallOption) (*bigquerypb.RowAccessPolicy, error) {
	return c.internalClient.GetRowAccessPolicy(ctx, req, opts...)
}

// CreateRowAccessPolicy creates a row access policy.
func (c *RowAccessPolicyClient) CreateRowAccessPolicy(ctx context.Context, req *bigquerypb.CreateRowAccessPolicyRequest, opts ...gax.CallOption) (*bigquerypb.RowAccessPolicy, error) {
	return c.internalClient.CreateRowAccessPolicy(ctx, req, opts...)
}

// UpdateRowAccessPolicy updates a row access policy.
func (c *RowAccessPolicyClient) UpdateRowAccessPolicy(ctx context.Context, req *bigquerypb.UpdateRowAccessPolicyRequest, opts ...gax.CallOption) (*bigquerypb.RowAccessPolicy, error) {
	return c.internalClient.UpdateRowAccessPolicy(ctx, req, opts...)
}

// DeleteRowAccessPolicy deletes a row access policy.
func (c *RowAccessPolicyClient) DeleteRowAccessPolicy(ctx context.Context, req *bigquerypb.DeleteRowAccessPolicyRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteRowAccessPolicy(ctx, req, opts...)
}

// BatchDeleteRowAccessPolicies deletes provided row access policies.
func (c *RowAccessPolicyClient) BatchDeleteRowAccessPolicies(ctx context.Context, req *bigquerypb.BatchDeleteRowAccessPoliciesRequest, opts ...gax.CallOption) error {
	return c.internalClient.BatchDeleteRowAccessPolicies(ctx, req, opts...)
}

// rowAccessPolicyGRPCClient is a client for interacting with BigQuery API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type rowAccessPolicyGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing RowAccessPolicyClient
	CallOptions **RowAccessPolicyCallOptions

	// The gRPC API client.
	rowAccessPolicyClient bigquerypb.RowAccessPolicyServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string

	logger *slog.Logger
}

// NewRowAccessPolicyClient creates a new row access policy service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service for interacting with row access policies.
func NewRowAccessPolicyClient(ctx context.Context, opts ...option.ClientOption) (*RowAccessPolicyClient, error) {
	clientOpts := defaultRowAccessPolicyGRPCClientOptions()
	if newRowAccessPolicyClientHook != nil {
		hookOpts, err := newRowAccessPolicyClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := RowAccessPolicyClient{CallOptions: defaultRowAccessPolicyCallOptions()}

	c := &rowAccessPolicyGRPCClient{
		connPool:              connPool,
		rowAccessPolicyClient: bigquerypb.NewRowAccessPolicyServiceClient(connPool),
		CallOptions:           &client.CallOptions,
		logger:                internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *rowAccessPolicyGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *rowAccessPolicyGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version, "pb", protoVersion)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *rowAccessPolicyGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type rowAccessPolicyRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing RowAccessPolicyClient
	CallOptions **RowAccessPolicyCallOptions

	logger *slog.Logger
}

// NewRowAccessPolicyRESTClient creates a new row access policy service rest client.
//
// Service for interacting with row access policies.
func NewRowAccessPolicyRESTClient(ctx context.Context, opts ...option.ClientOption) (*RowAccessPolicyClient, error) {
	clientOpts := append(defaultRowAccessPolicyRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultRowAccessPolicyRESTCallOptions()
	c := &rowAccessPolicyRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
		logger:      internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	return &RowAccessPolicyClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultRowAccessPolicyRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://bigquery.googleapis.com"),
		internaloption.WithDefaultEndpointTemplate("https://bigquery.UNIVERSE_DOMAIN"),
		internaloption.WithDefaultMTLSEndpoint("https://bigquery.mtls.googleapis.com"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://bigquery.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableNewAuthLibrary(),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *rowAccessPolicyRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN", "pb", protoVersion)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *rowAccessPolicyRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *rowAccessPolicyRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *rowAccessPolicyGRPCClient) ListRowAccessPolicies(ctx context.Context, req *bigquerypb.ListRowAccessPoliciesRequest, opts ...gax.CallOption) *RowAccessPolicyIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "dataset_id", url.QueryEscape(req.GetDatasetId()), "table_id", url.QueryEscape(req.GetTableId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListRowAccessPolicies[0:len((*c.CallOptions).ListRowAccessPolicies):len((*c.CallOptions).ListRowAccessPolicies)], opts...)
	it := &RowAccessPolicyIterator{}
	req = proto.Clone(req).(*bigquerypb.ListRowAccessPoliciesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*bigquerypb.RowAccessPolicy, string, error) {
		resp := &bigquerypb.ListRowAccessPoliciesResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = executeRPC(ctx, c.rowAccessPolicyClient.ListRowAccessPolicies, req, settings.GRPC, c.logger, "ListRowAccessPolicies")
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetRowAccessPolicies(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *rowAccessPolicyGRPCClient) GetRowAccessPolicy(ctx context.Context, req *bigquerypb.GetRowAccessPolicyRequest, opts ...gax.CallOption) (*bigquerypb.RowAccessPolicy, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "dataset_id", url.QueryEscape(req.GetDatasetId()), "table_id", url.QueryEscape(req.GetTableId()), "policy_id", url.QueryEscape(req.GetPolicyId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetRowAccessPolicy[0:len((*c.CallOptions).GetRowAccessPolicy):len((*c.CallOptions).GetRowAccessPolicy)], opts...)
	var resp *bigquerypb.RowAccessPolicy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.rowAccessPolicyClient.GetRowAccessPolicy, req, settings.GRPC, c.logger, "GetRowAccessPolicy")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *rowAccessPolicyGRPCClient) CreateRowAccessPolicy(ctx context.Context, req *bigquerypb.CreateRowAccessPolicyRequest, opts ...gax.CallOption) (*bigquerypb.RowAccessPolicy, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "dataset_id", url.QueryEscape(req.GetDatasetId()), "table_id", url.QueryEscape(req.GetTableId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).CreateRowAccessPolicy[0:len((*c.CallOptions).CreateRowAccessPolicy):len((*c.CallOptions).CreateRowAccessPolicy)], opts...)
	var resp *bigquerypb.RowAccessPolicy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.rowAccessPolicyClient.CreateRowAccessPolicy, req, settings.GRPC, c.logger, "CreateRowAccessPolicy")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *rowAccessPolicyGRPCClient) UpdateRowAccessPolicy(ctx context.Context, req *bigquerypb.UpdateRowAccessPolicyRequest, opts ...gax.CallOption) (*bigquerypb.RowAccessPolicy, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "dataset_id", url.QueryEscape(req.GetDatasetId()), "table_id", url.QueryEscape(req.GetTableId()), "policy_id", url.QueryEscape(req.GetPolicyId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).UpdateRowAccessPolicy[0:len((*c.CallOptions).UpdateRowAccessPolicy):len((*c.CallOptions).UpdateRowAccessPolicy)], opts...)
	var resp *bigquerypb.RowAccessPolicy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.rowAccessPolicyClient.UpdateRowAccessPolicy, req, settings.GRPC, c.logger, "UpdateRowAccessPolicy")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *rowAccessPolicyGRPCClient) DeleteRowAccessPolicy(ctx context.Context, req *bigquerypb.DeleteRowAccessPolicyRequest, opts ...gax.CallOption) error {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "dataset_id", url.QueryEscape(req.GetDatasetId()), "table_id", url.QueryEscape(req.GetTableId()), "policy_id", url.QueryEscape(req.GetPolicyId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).DeleteRowAccessPolicy[0:len((*c.CallOptions).DeleteRowAccessPolicy):len((*c.CallOptions).DeleteRowAccessPolicy)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = executeRPC(ctx, c.rowAccessPolicyClient.DeleteRowAccessPolicy, req, settings.GRPC, c.logger, "DeleteRowAccessPolicy")
		return err
	}, opts...)
	return err
}

func (c *rowAccessPolicyGRPCClient) BatchDeleteRowAccessPolicies(ctx context.Context, req *bigquerypb.BatchDeleteRowAccessPoliciesRequest, opts ...gax.CallOption) error {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "dataset_id", url.QueryEscape(req.GetDatasetId()), "table_id", url.QueryEscape(req.GetTableId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).BatchDeleteRowAccessPolicies[0:len((*c.CallOptions).BatchDeleteRowAccessPolicies):len((*c.CallOptions).BatchDeleteRowAccessPolicies)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = executeRPC(ctx, c.rowAccessPolicyClient.BatchDeleteRowAccessPolicies, req, settings.GRPC, c.logger, "BatchDeleteRowAccessPolicies")
		return err
	}, opts...)
	return err
}

// ListRowAccessPolicies lists all row access policies on the specified table.
func (c *rowAccessPolicyRESTClient) ListRowAccessPolicies(ctx context.Context, req *bigquerypb.ListRowAccessPoliciesRequest, opts ...gax.CallOption) *RowAccessPolicyIterator {
	it := &RowAccessPolicyIterator{}
	req = proto.Clone(req).(*bigquerypb.ListRowAccessPoliciesRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*bigquerypb.RowAccessPolicy, string, error) {
		resp := &bigquerypb.ListRowAccessPoliciesResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		baseUrl, err := url.Parse(c.endpoint)
		if err != nil {
			return nil, "", err
		}
		baseUrl.Path += fmt.Sprintf("/bigquery/v2/projects/%v/datasets/%v/tables/%v/rowAccessPolicies", req.GetProjectId(), req.GetDatasetId(), req.GetTableId())

		params := url.Values{}
		if req.GetPageSize() != 0 {
			params.Add("pageSize", fmt.Sprintf("%v", req.GetPageSize()))
		}
		if req.GetPageToken() != "" {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}

		baseUrl.RawQuery = params.Encode()

		// Build HTTP headers from client and context metadata.
		hds := append(c.xGoogHeaders, "Content-Type", "application/json")
		headers := gax.BuildHeaders(ctx, hds...)
		e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			if settings.Path != "" {
				baseUrl.Path = settings.Path
			}
			httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
			if err != nil {
				return err
			}
			httpReq.Header = headers

			buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "ListRowAccessPolicies")
			if err != nil {
				return err
			}
			if err := unm.Unmarshal(buf, resp); err != nil {
				return err
			}

			return nil
		}, opts...)
		if e != nil {
			return nil, "", e
		}
		it.Response = resp
		return resp.GetRowAccessPolicies(), resp.GetNextPageToken(), nil
	}

	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// GetRowAccessPolicy gets the specified row access policy by policy ID.
func (c *rowAccessPolicyRESTClient) GetRowAccessPolicy(ctx context.Context, req *bigquerypb.GetRowAccessPolicyRequest, opts ...gax.CallOption) (*bigquerypb.RowAccessPolicy, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/bigquery/v2/projects/%v/datasets/%v/tables/%v/rowAccessPolicies/%v", req.GetProjectId(), req.GetDatasetId(), req.GetTableId(), req.GetPolicyId())

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "dataset_id", url.QueryEscape(req.GetDatasetId()), "table_id", url.QueryEscape(req.GetTableId()), "policy_id", url.QueryEscape(req.GetPolicyId()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).GetRowAccessPolicy[0:len((*c.CallOptions).GetRowAccessPolicy):len((*c.CallOptions).GetRowAccessPolicy)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &bigquerypb.RowAccessPolicy{}
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

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "GetRowAccessPolicy")
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

// CreateRowAccessPolicy creates a row access policy.
func (c *rowAccessPolicyRESTClient) CreateRowAccessPolicy(ctx context.Context, req *bigquerypb.CreateRowAccessPolicyRequest, opts ...gax.CallOption) (*bigquerypb.RowAccessPolicy, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetRowAccessPolicy()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/bigquery/v2/projects/%v/datasets/%v/tables/%v/rowAccessPolicies", req.GetProjectId(), req.GetDatasetId(), req.GetTableId())

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "dataset_id", url.QueryEscape(req.GetDatasetId()), "table_id", url.QueryEscape(req.GetTableId()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).CreateRowAccessPolicy[0:len((*c.CallOptions).CreateRowAccessPolicy):len((*c.CallOptions).CreateRowAccessPolicy)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &bigquerypb.RowAccessPolicy{}
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

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, jsonReq, "CreateRowAccessPolicy")
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

// UpdateRowAccessPolicy updates a row access policy.
func (c *rowAccessPolicyRESTClient) UpdateRowAccessPolicy(ctx context.Context, req *bigquerypb.UpdateRowAccessPolicyRequest, opts ...gax.CallOption) (*bigquerypb.RowAccessPolicy, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetRowAccessPolicy()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/bigquery/v2/projects/%v/datasets/%v/tables/%v/rowAccessPolicies/%v", req.GetProjectId(), req.GetDatasetId(), req.GetTableId(), req.GetPolicyId())

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "dataset_id", url.QueryEscape(req.GetDatasetId()), "table_id", url.QueryEscape(req.GetTableId()), "policy_id", url.QueryEscape(req.GetPolicyId()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).UpdateRowAccessPolicy[0:len((*c.CallOptions).UpdateRowAccessPolicy):len((*c.CallOptions).UpdateRowAccessPolicy)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &bigquerypb.RowAccessPolicy{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("PUT", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, jsonReq, "UpdateRowAccessPolicy")
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

// DeleteRowAccessPolicy deletes a row access policy.
func (c *rowAccessPolicyRESTClient) DeleteRowAccessPolicy(ctx context.Context, req *bigquerypb.DeleteRowAccessPolicyRequest, opts ...gax.CallOption) error {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return err
	}
	baseUrl.Path += fmt.Sprintf("/bigquery/v2/projects/%v/datasets/%v/tables/%v/rowAccessPolicies/%v", req.GetProjectId(), req.GetDatasetId(), req.GetTableId(), req.GetPolicyId())

	params := url.Values{}
	if req != nil && req.Force != nil {
		params.Add("force", fmt.Sprintf("%v", req.GetForce()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "dataset_id", url.QueryEscape(req.GetDatasetId()), "table_id", url.QueryEscape(req.GetTableId()), "policy_id", url.QueryEscape(req.GetPolicyId()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	return gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("DELETE", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		_, err = executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "DeleteRowAccessPolicy")
		return err
	}, opts...)
}

// BatchDeleteRowAccessPolicies deletes provided row access policies.
func (c *rowAccessPolicyRESTClient) BatchDeleteRowAccessPolicies(ctx context.Context, req *bigquerypb.BatchDeleteRowAccessPoliciesRequest, opts ...gax.CallOption) error {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return err
	}
	baseUrl.Path += fmt.Sprintf("/bigquery/v2/projects/%v/datasets/%v/tables/%v/rowAccessPolicies:batchDelete", req.GetProjectId(), req.GetDatasetId(), req.GetTableId())

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "dataset_id", url.QueryEscape(req.GetDatasetId()), "table_id", url.QueryEscape(req.GetTableId()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	return gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		_, err = executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, jsonReq, "BatchDeleteRowAccessPolicies")
		return err
	}, opts...)
}
