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
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var newTableClientHook clientHook

// TableCallOptions contains the retry settings for each method of TableClient.
type TableCallOptions struct {
	GetTable    []gax.CallOption
	InsertTable []gax.CallOption
	PatchTable  []gax.CallOption
	UpdateTable []gax.CallOption
	DeleteTable []gax.CallOption
	ListTables  []gax.CallOption
}

func defaultTableGRPCClientOptions() []option.ClientOption {
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

func defaultTableCallOptions() *TableCallOptions {
	return &TableCallOptions{
		GetTable: []gax.CallOption{
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
		InsertTable: []gax.CallOption{
			gax.WithTimeout(240000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
					codes.ResourceExhausted,
				}, gax.Backoff{
					Initial:    400 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 2.00,
				})
			}),
		},
		PatchTable: []gax.CallOption{
			gax.WithTimeout(240000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
					codes.ResourceExhausted,
				}, gax.Backoff{
					Initial:    400 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 2.00,
				})
			}),
		},
		UpdateTable: []gax.CallOption{
			gax.WithTimeout(240000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
					codes.ResourceExhausted,
				}, gax.Backoff{
					Initial:    400 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 2.00,
				})
			}),
		},
		DeleteTable: []gax.CallOption{
			gax.WithTimeout(240000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
					codes.ResourceExhausted,
				}, gax.Backoff{
					Initial:    400 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 2.00,
				})
			}),
		},
		ListTables: []gax.CallOption{
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

func defaultTableRESTCallOptions() *TableCallOptions {
	return &TableCallOptions{
		GetTable: []gax.CallOption{
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
		InsertTable: []gax.CallOption{
			gax.WithTimeout(240000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    400 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 2.00,
				},
					http.StatusGatewayTimeout,
					http.StatusServiceUnavailable,
					http.StatusTooManyRequests)
			}),
		},
		PatchTable: []gax.CallOption{
			gax.WithTimeout(240000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    400 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 2.00,
				},
					http.StatusGatewayTimeout,
					http.StatusServiceUnavailable,
					http.StatusTooManyRequests)
			}),
		},
		UpdateTable: []gax.CallOption{
			gax.WithTimeout(240000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    400 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 2.00,
				},
					http.StatusGatewayTimeout,
					http.StatusServiceUnavailable,
					http.StatusTooManyRequests)
			}),
		},
		DeleteTable: []gax.CallOption{
			gax.WithTimeout(240000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    400 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 2.00,
				},
					http.StatusGatewayTimeout,
					http.StatusServiceUnavailable,
					http.StatusTooManyRequests)
			}),
		},
		ListTables: []gax.CallOption{
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

// internalTableClient is an interface that defines the methods available from BigQuery API.
type internalTableClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetTable(context.Context, *bigquerypb.GetTableRequest, ...gax.CallOption) (*bigquerypb.Table, error)
	InsertTable(context.Context, *bigquerypb.InsertTableRequest, ...gax.CallOption) (*bigquerypb.Table, error)
	PatchTable(context.Context, *bigquerypb.UpdateOrPatchTableRequest, ...gax.CallOption) (*bigquerypb.Table, error)
	UpdateTable(context.Context, *bigquerypb.UpdateOrPatchTableRequest, ...gax.CallOption) (*bigquerypb.Table, error)
	DeleteTable(context.Context, *bigquerypb.DeleteTableRequest, ...gax.CallOption) error
	ListTables(context.Context, *bigquerypb.ListTablesRequest, ...gax.CallOption) *ListFormatTableIterator
}

// TableClient is a client for interacting with BigQuery API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// TableService provides methods for managing BigQuery tables and table-like
// entities such as views and snapshots.
type TableClient struct {
	// The internal transport-dependent client.
	internalClient internalTableClient

	// The call options for this service.
	CallOptions *TableCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *TableClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *TableClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *TableClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetTable gets the specified table resource by table ID.
// This method does not return the data in the table, it only returns the
// table resource, which describes the structure of this table.
func (c *TableClient) GetTable(ctx context.Context, req *bigquerypb.GetTableRequest, opts ...gax.CallOption) (*bigquerypb.Table, error) {
	return c.internalClient.GetTable(ctx, req, opts...)
}

// InsertTable creates a new, empty table in the dataset.
func (c *TableClient) InsertTable(ctx context.Context, req *bigquerypb.InsertTableRequest, opts ...gax.CallOption) (*bigquerypb.Table, error) {
	return c.internalClient.InsertTable(ctx, req, opts...)
}

// PatchTable updates information in an existing table. The update method replaces the
// entire table resource, whereas the patch method only replaces fields that
// are provided in the submitted table resource.
// This method supports RFC5789 patch semantics.
func (c *TableClient) PatchTable(ctx context.Context, req *bigquerypb.UpdateOrPatchTableRequest, opts ...gax.CallOption) (*bigquerypb.Table, error) {
	return c.internalClient.PatchTable(ctx, req, opts...)
}

// UpdateTable updates information in an existing table. The update method replaces the
// entire Table resource, whereas the patch method only replaces fields that
// are provided in the submitted Table resource.
func (c *TableClient) UpdateTable(ctx context.Context, req *bigquerypb.UpdateOrPatchTableRequest, opts ...gax.CallOption) (*bigquerypb.Table, error) {
	return c.internalClient.UpdateTable(ctx, req, opts...)
}

// DeleteTable deletes the table specified by tableId from the dataset.
// If the table contains data, all the data will be deleted.
func (c *TableClient) DeleteTable(ctx context.Context, req *bigquerypb.DeleteTableRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteTable(ctx, req, opts...)
}

// ListTables lists all tables in the specified dataset. Requires the READER dataset
// role.
func (c *TableClient) ListTables(ctx context.Context, req *bigquerypb.ListTablesRequest, opts ...gax.CallOption) *ListFormatTableIterator {
	return c.internalClient.ListTables(ctx, req, opts...)
}

// tableGRPCClient is a client for interacting with BigQuery API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type tableGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing TableClient
	CallOptions **TableCallOptions

	// The gRPC API client.
	tableClient bigquerypb.TableServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string

	logger *slog.Logger
}

// NewTableClient creates a new table service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// TableService provides methods for managing BigQuery tables and table-like
// entities such as views and snapshots.
func NewTableClient(ctx context.Context, opts ...option.ClientOption) (*TableClient, error) {
	clientOpts := defaultTableGRPCClientOptions()
	if newTableClientHook != nil {
		hookOpts, err := newTableClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := TableClient{CallOptions: defaultTableCallOptions()}

	c := &tableGRPCClient{
		connPool:    connPool,
		tableClient: bigquerypb.NewTableServiceClient(connPool),
		CallOptions: &client.CallOptions,
		logger:      internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *tableGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *tableGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version, "pb", protoVersion)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *tableGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type tableRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing TableClient
	CallOptions **TableCallOptions

	logger *slog.Logger
}

// NewTableRESTClient creates a new table service rest client.
//
// TableService provides methods for managing BigQuery tables and table-like
// entities such as views and snapshots.
func NewTableRESTClient(ctx context.Context, opts ...option.ClientOption) (*TableClient, error) {
	clientOpts := append(defaultTableRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultTableRESTCallOptions()
	c := &tableRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
		logger:      internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	return &TableClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultTableRESTClientOptions() []option.ClientOption {
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
func (c *tableRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN", "pb", protoVersion)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *tableRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *tableRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *tableGRPCClient) GetTable(ctx context.Context, req *bigquerypb.GetTableRequest, opts ...gax.CallOption) (*bigquerypb.Table, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "dataset_id", url.QueryEscape(req.GetDatasetId()), "table_id", url.QueryEscape(req.GetTableId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetTable[0:len((*c.CallOptions).GetTable):len((*c.CallOptions).GetTable)], opts...)
	var resp *bigquerypb.Table
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.tableClient.GetTable, req, settings.GRPC, c.logger, "GetTable")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *tableGRPCClient) InsertTable(ctx context.Context, req *bigquerypb.InsertTableRequest, opts ...gax.CallOption) (*bigquerypb.Table, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "dataset_id", url.QueryEscape(req.GetDatasetId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).InsertTable[0:len((*c.CallOptions).InsertTable):len((*c.CallOptions).InsertTable)], opts...)
	var resp *bigquerypb.Table
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.tableClient.InsertTable, req, settings.GRPC, c.logger, "InsertTable")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *tableGRPCClient) PatchTable(ctx context.Context, req *bigquerypb.UpdateOrPatchTableRequest, opts ...gax.CallOption) (*bigquerypb.Table, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "dataset_id", url.QueryEscape(req.GetDatasetId()), "table_id", url.QueryEscape(req.GetTableId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).PatchTable[0:len((*c.CallOptions).PatchTable):len((*c.CallOptions).PatchTable)], opts...)
	var resp *bigquerypb.Table
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.tableClient.PatchTable, req, settings.GRPC, c.logger, "PatchTable")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *tableGRPCClient) UpdateTable(ctx context.Context, req *bigquerypb.UpdateOrPatchTableRequest, opts ...gax.CallOption) (*bigquerypb.Table, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "dataset_id", url.QueryEscape(req.GetDatasetId()), "table_id", url.QueryEscape(req.GetTableId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).UpdateTable[0:len((*c.CallOptions).UpdateTable):len((*c.CallOptions).UpdateTable)], opts...)
	var resp *bigquerypb.Table
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.tableClient.UpdateTable, req, settings.GRPC, c.logger, "UpdateTable")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *tableGRPCClient) DeleteTable(ctx context.Context, req *bigquerypb.DeleteTableRequest, opts ...gax.CallOption) error {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "dataset_id", url.QueryEscape(req.GetDatasetId()), "table_id", url.QueryEscape(req.GetTableId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).DeleteTable[0:len((*c.CallOptions).DeleteTable):len((*c.CallOptions).DeleteTable)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = executeRPC(ctx, c.tableClient.DeleteTable, req, settings.GRPC, c.logger, "DeleteTable")
		return err
	}, opts...)
	return err
}

func (c *tableGRPCClient) ListTables(ctx context.Context, req *bigquerypb.ListTablesRequest, opts ...gax.CallOption) *ListFormatTableIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "dataset_id", url.QueryEscape(req.GetDatasetId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListTables[0:len((*c.CallOptions).ListTables):len((*c.CallOptions).ListTables)], opts...)
	it := &ListFormatTableIterator{}
	req = proto.Clone(req).(*bigquerypb.ListTablesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*bigquerypb.ListFormatTable, string, error) {
		resp := &bigquerypb.TableList{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.MaxResults = &wrapperspb.UInt32Value{Value: uint32(math.MaxInt32)}
		} else if pageSize != 0 {
			req.MaxResults = &wrapperspb.UInt32Value{Value: uint32(pageSize)}
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = executeRPC(ctx, c.tableClient.ListTables, req, settings.GRPC, c.logger, "ListTables")
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetTables(), resp.GetNextPageToken(), nil
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
	if psVal := req.GetMaxResults(); psVal != nil {
		it.pageInfo.MaxSize = int(psVal.GetValue())
	}
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// GetTable gets the specified table resource by table ID.
// This method does not return the data in the table, it only returns the
// table resource, which describes the structure of this table.
func (c *tableRESTClient) GetTable(ctx context.Context, req *bigquerypb.GetTableRequest, opts ...gax.CallOption) (*bigquerypb.Table, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/bigquery/v2/projects/%v/datasets/%v/tables/%v", req.GetProjectId(), req.GetDatasetId(), req.GetTableId())

	params := url.Values{}
	if req.GetSelectedFields() != "" {
		params.Add("selectedFields", fmt.Sprintf("%v", req.GetSelectedFields()))
	}
	if req.GetView() != 0 {
		params.Add("view", fmt.Sprintf("%v", req.GetView()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "dataset_id", url.QueryEscape(req.GetDatasetId()), "table_id", url.QueryEscape(req.GetTableId()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).GetTable[0:len((*c.CallOptions).GetTable):len((*c.CallOptions).GetTable)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &bigquerypb.Table{}
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

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "GetTable")
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

// InsertTable creates a new, empty table in the dataset.
func (c *tableRESTClient) InsertTable(ctx context.Context, req *bigquerypb.InsertTableRequest, opts ...gax.CallOption) (*bigquerypb.Table, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetTable()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/bigquery/v2/projects/%v/datasets/%v/tables", req.GetProjectId(), req.GetDatasetId())

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "dataset_id", url.QueryEscape(req.GetDatasetId()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).InsertTable[0:len((*c.CallOptions).InsertTable):len((*c.CallOptions).InsertTable)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &bigquerypb.Table{}
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

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, jsonReq, "InsertTable")
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

// PatchTable updates information in an existing table. The update method replaces the
// entire table resource, whereas the patch method only replaces fields that
// are provided in the submitted table resource.
// This method supports RFC5789 patch semantics.
func (c *tableRESTClient) PatchTable(ctx context.Context, req *bigquerypb.UpdateOrPatchTableRequest, opts ...gax.CallOption) (*bigquerypb.Table, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetTable()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/bigquery/v2/projects/%v/datasets/%v/tables/%v", req.GetProjectId(), req.GetDatasetId(), req.GetTableId())

	params := url.Values{}
	if req.GetAutodetectSchema() {
		params.Add("autodetectSchema", fmt.Sprintf("%v", req.GetAutodetectSchema()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "dataset_id", url.QueryEscape(req.GetDatasetId()), "table_id", url.QueryEscape(req.GetTableId()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).PatchTable[0:len((*c.CallOptions).PatchTable):len((*c.CallOptions).PatchTable)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &bigquerypb.Table{}
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

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, jsonReq, "PatchTable")
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

// UpdateTable updates information in an existing table. The update method replaces the
// entire Table resource, whereas the patch method only replaces fields that
// are provided in the submitted Table resource.
func (c *tableRESTClient) UpdateTable(ctx context.Context, req *bigquerypb.UpdateOrPatchTableRequest, opts ...gax.CallOption) (*bigquerypb.Table, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetTable()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/bigquery/v2/projects/%v/datasets/%v/tables/%v", req.GetProjectId(), req.GetDatasetId(), req.GetTableId())

	params := url.Values{}
	if req.GetAutodetectSchema() {
		params.Add("autodetectSchema", fmt.Sprintf("%v", req.GetAutodetectSchema()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "dataset_id", url.QueryEscape(req.GetDatasetId()), "table_id", url.QueryEscape(req.GetTableId()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).UpdateTable[0:len((*c.CallOptions).UpdateTable):len((*c.CallOptions).UpdateTable)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &bigquerypb.Table{}
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

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, jsonReq, "UpdateTable")
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

// DeleteTable deletes the table specified by tableId from the dataset.
// If the table contains data, all the data will be deleted.
func (c *tableRESTClient) DeleteTable(ctx context.Context, req *bigquerypb.DeleteTableRequest, opts ...gax.CallOption) error {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return err
	}
	baseUrl.Path += fmt.Sprintf("/bigquery/v2/projects/%v/datasets/%v/tables/%v", req.GetProjectId(), req.GetDatasetId(), req.GetTableId())

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "dataset_id", url.QueryEscape(req.GetDatasetId()), "table_id", url.QueryEscape(req.GetTableId()))}

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

		_, err = executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "DeleteTable")
		return err
	}, opts...)
}

// ListTables lists all tables in the specified dataset. Requires the READER dataset
// role.
func (c *tableRESTClient) ListTables(ctx context.Context, req *bigquerypb.ListTablesRequest, opts ...gax.CallOption) *ListFormatTableIterator {
	it := &ListFormatTableIterator{}
	req = proto.Clone(req).(*bigquerypb.ListTablesRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*bigquerypb.ListFormatTable, string, error) {
		resp := &bigquerypb.TableList{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.MaxResults = &wrapperspb.UInt32Value{Value: uint32(math.MaxInt32)}
		} else if pageSize != 0 {
			req.MaxResults = &wrapperspb.UInt32Value{Value: uint32(pageSize)}
		}
		baseUrl, err := url.Parse(c.endpoint)
		if err != nil {
			return nil, "", err
		}
		baseUrl.Path += fmt.Sprintf("/bigquery/v2/projects/%v/datasets/%v/tables", req.GetProjectId(), req.GetDatasetId())

		params := url.Values{}
		if req.GetMaxResults() != nil {
			field, err := protojson.Marshal(req.GetMaxResults())
			if err != nil {
				return nil, "", err
			}
			params.Add("maxResults", string(field))
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

			buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "ListTables")
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
		return resp.GetTables(), resp.GetNextPageToken(), nil
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
	if psVal := req.GetMaxResults(); psVal != nil {
		it.pageInfo.MaxSize = int(psVal.GetValue())
	}
	it.pageInfo.Token = req.GetPageToken()

	return it
}
