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

package resourcemanager

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"math"
	"net/http"
	"net/url"
	"time"

	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	resourcemanagerpb "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
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

var newTagBindingsClientHook clientHook

// TagBindingsCallOptions contains the retry settings for each method of TagBindingsClient.
type TagBindingsCallOptions struct {
	ListTagBindings   []gax.CallOption
	CreateTagBinding  []gax.CallOption
	DeleteTagBinding  []gax.CallOption
	ListEffectiveTags []gax.CallOption
	GetOperation      []gax.CallOption
}

func defaultTagBindingsGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("cloudresourcemanager.googleapis.com:443"),
		internaloption.WithDefaultEndpointTemplate("cloudresourcemanager.UNIVERSE_DOMAIN:443"),
		internaloption.WithDefaultMTLSEndpoint("cloudresourcemanager.mtls.googleapis.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://cloudresourcemanager.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		internaloption.EnableNewAuthLibrary(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultTagBindingsCallOptions() *TagBindingsCallOptions {
	return &TagBindingsCallOptions{
		ListTagBindings: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		CreateTagBinding: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		DeleteTagBinding: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		ListEffectiveTags: []gax.CallOption{},
		GetOperation:      []gax.CallOption{},
	}
}

func defaultTagBindingsRESTCallOptions() *TagBindingsCallOptions {
	return &TagBindingsCallOptions{
		ListTagBindings: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
		CreateTagBinding: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		DeleteTagBinding: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		ListEffectiveTags: []gax.CallOption{},
		GetOperation:      []gax.CallOption{},
	}
}

// internalTagBindingsClient is an interface that defines the methods available from Cloud Resource Manager API.
type internalTagBindingsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListTagBindings(context.Context, *resourcemanagerpb.ListTagBindingsRequest, ...gax.CallOption) *TagBindingIterator
	CreateTagBinding(context.Context, *resourcemanagerpb.CreateTagBindingRequest, ...gax.CallOption) (*CreateTagBindingOperation, error)
	CreateTagBindingOperation(name string) *CreateTagBindingOperation
	DeleteTagBinding(context.Context, *resourcemanagerpb.DeleteTagBindingRequest, ...gax.CallOption) (*DeleteTagBindingOperation, error)
	DeleteTagBindingOperation(name string) *DeleteTagBindingOperation
	ListEffectiveTags(context.Context, *resourcemanagerpb.ListEffectiveTagsRequest, ...gax.CallOption) *EffectiveTagIterator
	GetOperation(context.Context, *longrunningpb.GetOperationRequest, ...gax.CallOption) (*longrunningpb.Operation, error)
}

// TagBindingsClient is a client for interacting with Cloud Resource Manager API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Allow users to create and manage TagBindings between TagValues and
// different Google Cloud resources throughout the GCP resource hierarchy.
type TagBindingsClient struct {
	// The internal transport-dependent client.
	internalClient internalTagBindingsClient

	// The call options for this service.
	CallOptions *TagBindingsCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *TagBindingsClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *TagBindingsClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *TagBindingsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ListTagBindings lists the TagBindings for the given Google Cloud resource, as specified
// with parent.
//
// NOTE: The parent field is expected to be a full resource name:
// https://cloud.google.com/apis/design/resource_names#full_resource_name (at https://cloud.google.com/apis/design/resource_names#full_resource_name)
func (c *TagBindingsClient) ListTagBindings(ctx context.Context, req *resourcemanagerpb.ListTagBindingsRequest, opts ...gax.CallOption) *TagBindingIterator {
	return c.internalClient.ListTagBindings(ctx, req, opts...)
}

// CreateTagBinding creates a TagBinding between a TagValue and a Google Cloud resource.
func (c *TagBindingsClient) CreateTagBinding(ctx context.Context, req *resourcemanagerpb.CreateTagBindingRequest, opts ...gax.CallOption) (*CreateTagBindingOperation, error) {
	return c.internalClient.CreateTagBinding(ctx, req, opts...)
}

// CreateTagBindingOperation returns a new CreateTagBindingOperation from a given name.
// The name must be that of a previously created CreateTagBindingOperation, possibly from a different process.
func (c *TagBindingsClient) CreateTagBindingOperation(name string) *CreateTagBindingOperation {
	return c.internalClient.CreateTagBindingOperation(name)
}

// DeleteTagBinding deletes a TagBinding.
func (c *TagBindingsClient) DeleteTagBinding(ctx context.Context, req *resourcemanagerpb.DeleteTagBindingRequest, opts ...gax.CallOption) (*DeleteTagBindingOperation, error) {
	return c.internalClient.DeleteTagBinding(ctx, req, opts...)
}

// DeleteTagBindingOperation returns a new DeleteTagBindingOperation from a given name.
// The name must be that of a previously created DeleteTagBindingOperation, possibly from a different process.
func (c *TagBindingsClient) DeleteTagBindingOperation(name string) *DeleteTagBindingOperation {
	return c.internalClient.DeleteTagBindingOperation(name)
}

// ListEffectiveTags return a list of effective tags for the given Google Cloud resource, as
// specified in parent.
func (c *TagBindingsClient) ListEffectiveTags(ctx context.Context, req *resourcemanagerpb.ListEffectiveTagsRequest, opts ...gax.CallOption) *EffectiveTagIterator {
	return c.internalClient.ListEffectiveTags(ctx, req, opts...)
}

// GetOperation is a utility method from google.longrunning.Operations.
func (c *TagBindingsClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	return c.internalClient.GetOperation(ctx, req, opts...)
}

// tagBindingsGRPCClient is a client for interacting with Cloud Resource Manager API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type tagBindingsGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing TagBindingsClient
	CallOptions **TagBindingsCallOptions

	// The gRPC API client.
	tagBindingsClient resourcemanagerpb.TagBindingsClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	operationsClient longrunningpb.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string

	logger *slog.Logger
}

// NewTagBindingsClient creates a new tag bindings client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Allow users to create and manage TagBindings between TagValues and
// different Google Cloud resources throughout the GCP resource hierarchy.
func NewTagBindingsClient(ctx context.Context, opts ...option.ClientOption) (*TagBindingsClient, error) {
	clientOpts := defaultTagBindingsGRPCClientOptions()
	if newTagBindingsClientHook != nil {
		hookOpts, err := newTagBindingsClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := TagBindingsClient{CallOptions: defaultTagBindingsCallOptions()}

	c := &tagBindingsGRPCClient{
		connPool:          connPool,
		tagBindingsClient: resourcemanagerpb.NewTagBindingsClient(connPool),
		CallOptions:       &client.CallOptions,
		logger:            internaloption.GetLogger(opts),
		operationsClient:  longrunningpb.NewOperationsClient(connPool),
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	client.LROClient, err = lroauto.NewOperationsClient(ctx, gtransport.WithConnPool(connPool))
	if err != nil {
		// This error "should not happen", since we are just reusing old connection pool
		// and never actually need to dial.
		// If this does happen, we could leak connp. However, we cannot close conn:
		// If the user invoked the constructor with option.WithGRPCConn,
		// we would close a connection that's still in use.
		// TODO: investigate error conditions.
		return nil, err
	}
	c.LROClient = &client.LROClient
	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *tagBindingsGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *tagBindingsGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version, "pb", protoVersion)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *tagBindingsGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type tagBindingsRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing TagBindingsClient
	CallOptions **TagBindingsCallOptions

	logger *slog.Logger
}

// NewTagBindingsRESTClient creates a new tag bindings rest client.
//
// Allow users to create and manage TagBindings between TagValues and
// different Google Cloud resources throughout the GCP resource hierarchy.
func NewTagBindingsRESTClient(ctx context.Context, opts ...option.ClientOption) (*TagBindingsClient, error) {
	clientOpts := append(defaultTagBindingsRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultTagBindingsRESTCallOptions()
	c := &tagBindingsRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
		logger:      internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	lroOpts := []option.ClientOption{
		option.WithHTTPClient(httpClient),
		option.WithEndpoint(endpoint),
	}
	opClient, err := lroauto.NewOperationsRESTClient(ctx, lroOpts...)
	if err != nil {
		return nil, err
	}
	c.LROClient = &opClient

	return &TagBindingsClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultTagBindingsRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://cloudresourcemanager.googleapis.com"),
		internaloption.WithDefaultEndpointTemplate("https://cloudresourcemanager.UNIVERSE_DOMAIN"),
		internaloption.WithDefaultMTLSEndpoint("https://cloudresourcemanager.mtls.googleapis.com"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://cloudresourcemanager.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableNewAuthLibrary(),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *tagBindingsRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN", "pb", protoVersion)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *tagBindingsRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *tagBindingsRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *tagBindingsGRPCClient) ListTagBindings(ctx context.Context, req *resourcemanagerpb.ListTagBindingsRequest, opts ...gax.CallOption) *TagBindingIterator {
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
	opts = append((*c.CallOptions).ListTagBindings[0:len((*c.CallOptions).ListTagBindings):len((*c.CallOptions).ListTagBindings)], opts...)
	it := &TagBindingIterator{}
	req = proto.Clone(req).(*resourcemanagerpb.ListTagBindingsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*resourcemanagerpb.TagBinding, string, error) {
		resp := &resourcemanagerpb.ListTagBindingsResponse{}
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
			resp, err = executeRPC(ctx, c.tagBindingsClient.ListTagBindings, req, settings.GRPC, c.logger, "ListTagBindings")
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetTagBindings(), resp.GetNextPageToken(), nil
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

func (c *tagBindingsGRPCClient) CreateTagBinding(ctx context.Context, req *resourcemanagerpb.CreateTagBindingRequest, opts ...gax.CallOption) (*CreateTagBindingOperation, error) {
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
	opts = append((*c.CallOptions).CreateTagBinding[0:len((*c.CallOptions).CreateTagBinding):len((*c.CallOptions).CreateTagBinding)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.tagBindingsClient.CreateTagBinding, req, settings.GRPC, c.logger, "CreateTagBinding")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &CreateTagBindingOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *tagBindingsGRPCClient) DeleteTagBinding(ctx context.Context, req *resourcemanagerpb.DeleteTagBindingRequest, opts ...gax.CallOption) (*DeleteTagBindingOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).DeleteTagBinding[0:len((*c.CallOptions).DeleteTagBinding):len((*c.CallOptions).DeleteTagBinding)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.tagBindingsClient.DeleteTagBinding, req, settings.GRPC, c.logger, "DeleteTagBinding")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &DeleteTagBindingOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *tagBindingsGRPCClient) ListEffectiveTags(ctx context.Context, req *resourcemanagerpb.ListEffectiveTagsRequest, opts ...gax.CallOption) *EffectiveTagIterator {
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
	opts = append((*c.CallOptions).ListEffectiveTags[0:len((*c.CallOptions).ListEffectiveTags):len((*c.CallOptions).ListEffectiveTags)], opts...)
	it := &EffectiveTagIterator{}
	req = proto.Clone(req).(*resourcemanagerpb.ListEffectiveTagsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*resourcemanagerpb.EffectiveTag, string, error) {
		resp := &resourcemanagerpb.ListEffectiveTagsResponse{}
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
			resp, err = executeRPC(ctx, c.tagBindingsClient.ListEffectiveTags, req, settings.GRPC, c.logger, "ListEffectiveTags")
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetEffectiveTags(), resp.GetNextPageToken(), nil
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

func (c *tagBindingsGRPCClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetOperation[0:len((*c.CallOptions).GetOperation):len((*c.CallOptions).GetOperation)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.operationsClient.GetOperation, req, settings.GRPC, c.logger, "GetOperation")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ListTagBindings lists the TagBindings for the given Google Cloud resource, as specified
// with parent.
//
// NOTE: The parent field is expected to be a full resource name:
// https://cloud.google.com/apis/design/resource_names#full_resource_name (at https://cloud.google.com/apis/design/resource_names#full_resource_name)
func (c *tagBindingsRESTClient) ListTagBindings(ctx context.Context, req *resourcemanagerpb.ListTagBindingsRequest, opts ...gax.CallOption) *TagBindingIterator {
	it := &TagBindingIterator{}
	req = proto.Clone(req).(*resourcemanagerpb.ListTagBindingsRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*resourcemanagerpb.TagBinding, string, error) {
		resp := &resourcemanagerpb.ListTagBindingsResponse{}
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
		baseUrl.Path += fmt.Sprintf("/v3/tagBindings")

		params := url.Values{}
		params.Add("$alt", "json;enum-encoding=int")
		if req.GetPageSize() != 0 {
			params.Add("pageSize", fmt.Sprintf("%v", req.GetPageSize()))
		}
		if req.GetPageToken() != "" {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}
		params.Add("parent", fmt.Sprintf("%v", req.GetParent()))

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

			buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "ListTagBindings")
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
		return resp.GetTagBindings(), resp.GetNextPageToken(), nil
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

// CreateTagBinding creates a TagBinding between a TagValue and a Google Cloud resource.
func (c *tagBindingsRESTClient) CreateTagBinding(ctx context.Context, req *resourcemanagerpb.CreateTagBindingRequest, opts ...gax.CallOption) (*CreateTagBindingOperation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetTagBinding()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v3/tagBindings")

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")
	if req.GetValidateOnly() {
		params.Add("validateOnly", fmt.Sprintf("%v", req.GetValidateOnly()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := append(c.xGoogHeaders, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &longrunningpb.Operation{}
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

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, jsonReq, "CreateTagBinding")
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

	override := fmt.Sprintf("/v3/%s", resp.GetName())
	return &CreateTagBindingOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, resp),
		pollPath: override,
	}, nil
}

// DeleteTagBinding deletes a TagBinding.
func (c *tagBindingsRESTClient) DeleteTagBinding(ctx context.Context, req *resourcemanagerpb.DeleteTagBindingRequest, opts ...gax.CallOption) (*DeleteTagBindingOperation, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v3/%v", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &longrunningpb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("DELETE", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "DeleteTagBinding")
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

	override := fmt.Sprintf("/v3/%s", resp.GetName())
	return &DeleteTagBindingOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, resp),
		pollPath: override,
	}, nil
}

// ListEffectiveTags return a list of effective tags for the given Google Cloud resource, as
// specified in parent.
func (c *tagBindingsRESTClient) ListEffectiveTags(ctx context.Context, req *resourcemanagerpb.ListEffectiveTagsRequest, opts ...gax.CallOption) *EffectiveTagIterator {
	it := &EffectiveTagIterator{}
	req = proto.Clone(req).(*resourcemanagerpb.ListEffectiveTagsRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*resourcemanagerpb.EffectiveTag, string, error) {
		resp := &resourcemanagerpb.ListEffectiveTagsResponse{}
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
		baseUrl.Path += fmt.Sprintf("/v3/effectiveTags")

		params := url.Values{}
		params.Add("$alt", "json;enum-encoding=int")
		if req.GetPageSize() != 0 {
			params.Add("pageSize", fmt.Sprintf("%v", req.GetPageSize()))
		}
		if req.GetPageToken() != "" {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}
		params.Add("parent", fmt.Sprintf("%v", req.GetParent()))

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

			buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "ListEffectiveTags")
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
		return resp.GetEffectiveTags(), resp.GetNextPageToken(), nil
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

// GetOperation is a utility method from google.longrunning.Operations.
func (c *tagBindingsRESTClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v3/%v", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).GetOperation[0:len((*c.CallOptions).GetOperation):len((*c.CallOptions).GetOperation)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &longrunningpb.Operation{}
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

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "GetOperation")
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

// CreateTagBindingOperation returns a new CreateTagBindingOperation from a given name.
// The name must be that of a previously created CreateTagBindingOperation, possibly from a different process.
func (c *tagBindingsGRPCClient) CreateTagBindingOperation(name string) *CreateTagBindingOperation {
	return &CreateTagBindingOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// CreateTagBindingOperation returns a new CreateTagBindingOperation from a given name.
// The name must be that of a previously created CreateTagBindingOperation, possibly from a different process.
func (c *tagBindingsRESTClient) CreateTagBindingOperation(name string) *CreateTagBindingOperation {
	override := fmt.Sprintf("/v3/%s", name)
	return &CreateTagBindingOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
		pollPath: override,
	}
}

// DeleteTagBindingOperation returns a new DeleteTagBindingOperation from a given name.
// The name must be that of a previously created DeleteTagBindingOperation, possibly from a different process.
func (c *tagBindingsGRPCClient) DeleteTagBindingOperation(name string) *DeleteTagBindingOperation {
	return &DeleteTagBindingOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// DeleteTagBindingOperation returns a new DeleteTagBindingOperation from a given name.
// The name must be that of a previously created DeleteTagBindingOperation, possibly from a different process.
func (c *tagBindingsRESTClient) DeleteTagBindingOperation(name string) *DeleteTagBindingOperation {
	override := fmt.Sprintf("/v3/%s", name)
	return &DeleteTagBindingOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
		pollPath: override,
	}
}
