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

package aiplatform

import (
	"context"
	"fmt"
	"log/slog"
	"math"
	"net/url"

	aiplatformpb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	iampb "cloud.google.com/go/iam/apiv1/iampb"
	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

var newDeploymentResourcePoolClientHook clientHook

// DeploymentResourcePoolCallOptions contains the retry settings for each method of DeploymentResourcePoolClient.
type DeploymentResourcePoolCallOptions struct {
	CreateDeploymentResourcePool []gax.CallOption
	GetDeploymentResourcePool    []gax.CallOption
	ListDeploymentResourcePools  []gax.CallOption
	UpdateDeploymentResourcePool []gax.CallOption
	DeleteDeploymentResourcePool []gax.CallOption
	QueryDeployedModels          []gax.CallOption
	GetLocation                  []gax.CallOption
	ListLocations                []gax.CallOption
	GetIamPolicy                 []gax.CallOption
	SetIamPolicy                 []gax.CallOption
	TestIamPermissions           []gax.CallOption
	CancelOperation              []gax.CallOption
	DeleteOperation              []gax.CallOption
	GetOperation                 []gax.CallOption
	ListOperations               []gax.CallOption
	WaitOperation                []gax.CallOption
}

func defaultDeploymentResourcePoolGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("aiplatform.googleapis.com:443"),
		internaloption.WithDefaultEndpointTemplate("aiplatform.UNIVERSE_DOMAIN:443"),
		internaloption.WithDefaultMTLSEndpoint("aiplatform.mtls.googleapis.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://aiplatform.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		internaloption.EnableNewAuthLibrary(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultDeploymentResourcePoolCallOptions() *DeploymentResourcePoolCallOptions {
	return &DeploymentResourcePoolCallOptions{
		CreateDeploymentResourcePool: []gax.CallOption{},
		GetDeploymentResourcePool:    []gax.CallOption{},
		ListDeploymentResourcePools:  []gax.CallOption{},
		UpdateDeploymentResourcePool: []gax.CallOption{},
		DeleteDeploymentResourcePool: []gax.CallOption{},
		QueryDeployedModels:          []gax.CallOption{},
		GetLocation:                  []gax.CallOption{},
		ListLocations:                []gax.CallOption{},
		GetIamPolicy:                 []gax.CallOption{},
		SetIamPolicy:                 []gax.CallOption{},
		TestIamPermissions:           []gax.CallOption{},
		CancelOperation:              []gax.CallOption{},
		DeleteOperation:              []gax.CallOption{},
		GetOperation:                 []gax.CallOption{},
		ListOperations:               []gax.CallOption{},
		WaitOperation:                []gax.CallOption{},
	}
}

// internalDeploymentResourcePoolClient is an interface that defines the methods available from Vertex AI API.
type internalDeploymentResourcePoolClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateDeploymentResourcePool(context.Context, *aiplatformpb.CreateDeploymentResourcePoolRequest, ...gax.CallOption) (*CreateDeploymentResourcePoolOperation, error)
	CreateDeploymentResourcePoolOperation(name string) *CreateDeploymentResourcePoolOperation
	GetDeploymentResourcePool(context.Context, *aiplatformpb.GetDeploymentResourcePoolRequest, ...gax.CallOption) (*aiplatformpb.DeploymentResourcePool, error)
	ListDeploymentResourcePools(context.Context, *aiplatformpb.ListDeploymentResourcePoolsRequest, ...gax.CallOption) *DeploymentResourcePoolIterator
	UpdateDeploymentResourcePool(context.Context, *aiplatformpb.UpdateDeploymentResourcePoolRequest, ...gax.CallOption) (*UpdateDeploymentResourcePoolOperation, error)
	UpdateDeploymentResourcePoolOperation(name string) *UpdateDeploymentResourcePoolOperation
	DeleteDeploymentResourcePool(context.Context, *aiplatformpb.DeleteDeploymentResourcePoolRequest, ...gax.CallOption) (*DeleteDeploymentResourcePoolOperation, error)
	DeleteDeploymentResourcePoolOperation(name string) *DeleteDeploymentResourcePoolOperation
	QueryDeployedModels(context.Context, *aiplatformpb.QueryDeployedModelsRequest, ...gax.CallOption) *DeployedModelIterator
	GetLocation(context.Context, *locationpb.GetLocationRequest, ...gax.CallOption) (*locationpb.Location, error)
	ListLocations(context.Context, *locationpb.ListLocationsRequest, ...gax.CallOption) *LocationIterator
	GetIamPolicy(context.Context, *iampb.GetIamPolicyRequest, ...gax.CallOption) (*iampb.Policy, error)
	SetIamPolicy(context.Context, *iampb.SetIamPolicyRequest, ...gax.CallOption) (*iampb.Policy, error)
	TestIamPermissions(context.Context, *iampb.TestIamPermissionsRequest, ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error)
	CancelOperation(context.Context, *longrunningpb.CancelOperationRequest, ...gax.CallOption) error
	DeleteOperation(context.Context, *longrunningpb.DeleteOperationRequest, ...gax.CallOption) error
	GetOperation(context.Context, *longrunningpb.GetOperationRequest, ...gax.CallOption) (*longrunningpb.Operation, error)
	ListOperations(context.Context, *longrunningpb.ListOperationsRequest, ...gax.CallOption) *OperationIterator
	WaitOperation(context.Context, *longrunningpb.WaitOperationRequest, ...gax.CallOption) (*longrunningpb.Operation, error)
}

// DeploymentResourcePoolClient is a client for interacting with Vertex AI API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// A service that manages the DeploymentResourcePool resource.
type DeploymentResourcePoolClient struct {
	// The internal transport-dependent client.
	internalClient internalDeploymentResourcePoolClient

	// The call options for this service.
	CallOptions *DeploymentResourcePoolCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *DeploymentResourcePoolClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *DeploymentResourcePoolClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *DeploymentResourcePoolClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreateDeploymentResourcePool create a DeploymentResourcePool.
func (c *DeploymentResourcePoolClient) CreateDeploymentResourcePool(ctx context.Context, req *aiplatformpb.CreateDeploymentResourcePoolRequest, opts ...gax.CallOption) (*CreateDeploymentResourcePoolOperation, error) {
	return c.internalClient.CreateDeploymentResourcePool(ctx, req, opts...)
}

// CreateDeploymentResourcePoolOperation returns a new CreateDeploymentResourcePoolOperation from a given name.
// The name must be that of a previously created CreateDeploymentResourcePoolOperation, possibly from a different process.
func (c *DeploymentResourcePoolClient) CreateDeploymentResourcePoolOperation(name string) *CreateDeploymentResourcePoolOperation {
	return c.internalClient.CreateDeploymentResourcePoolOperation(name)
}

// GetDeploymentResourcePool get a DeploymentResourcePool.
func (c *DeploymentResourcePoolClient) GetDeploymentResourcePool(ctx context.Context, req *aiplatformpb.GetDeploymentResourcePoolRequest, opts ...gax.CallOption) (*aiplatformpb.DeploymentResourcePool, error) {
	return c.internalClient.GetDeploymentResourcePool(ctx, req, opts...)
}

// ListDeploymentResourcePools list DeploymentResourcePools in a location.
func (c *DeploymentResourcePoolClient) ListDeploymentResourcePools(ctx context.Context, req *aiplatformpb.ListDeploymentResourcePoolsRequest, opts ...gax.CallOption) *DeploymentResourcePoolIterator {
	return c.internalClient.ListDeploymentResourcePools(ctx, req, opts...)
}

// UpdateDeploymentResourcePool update a DeploymentResourcePool.
func (c *DeploymentResourcePoolClient) UpdateDeploymentResourcePool(ctx context.Context, req *aiplatformpb.UpdateDeploymentResourcePoolRequest, opts ...gax.CallOption) (*UpdateDeploymentResourcePoolOperation, error) {
	return c.internalClient.UpdateDeploymentResourcePool(ctx, req, opts...)
}

// UpdateDeploymentResourcePoolOperation returns a new UpdateDeploymentResourcePoolOperation from a given name.
// The name must be that of a previously created UpdateDeploymentResourcePoolOperation, possibly from a different process.
func (c *DeploymentResourcePoolClient) UpdateDeploymentResourcePoolOperation(name string) *UpdateDeploymentResourcePoolOperation {
	return c.internalClient.UpdateDeploymentResourcePoolOperation(name)
}

// DeleteDeploymentResourcePool delete a DeploymentResourcePool.
func (c *DeploymentResourcePoolClient) DeleteDeploymentResourcePool(ctx context.Context, req *aiplatformpb.DeleteDeploymentResourcePoolRequest, opts ...gax.CallOption) (*DeleteDeploymentResourcePoolOperation, error) {
	return c.internalClient.DeleteDeploymentResourcePool(ctx, req, opts...)
}

// DeleteDeploymentResourcePoolOperation returns a new DeleteDeploymentResourcePoolOperation from a given name.
// The name must be that of a previously created DeleteDeploymentResourcePoolOperation, possibly from a different process.
func (c *DeploymentResourcePoolClient) DeleteDeploymentResourcePoolOperation(name string) *DeleteDeploymentResourcePoolOperation {
	return c.internalClient.DeleteDeploymentResourcePoolOperation(name)
}

// QueryDeployedModels list DeployedModels that have been deployed on this DeploymentResourcePool.
func (c *DeploymentResourcePoolClient) QueryDeployedModels(ctx context.Context, req *aiplatformpb.QueryDeployedModelsRequest, opts ...gax.CallOption) *DeployedModelIterator {
	return c.internalClient.QueryDeployedModels(ctx, req, opts...)
}

// GetLocation gets information about a location.
func (c *DeploymentResourcePoolClient) GetLocation(ctx context.Context, req *locationpb.GetLocationRequest, opts ...gax.CallOption) (*locationpb.Location, error) {
	return c.internalClient.GetLocation(ctx, req, opts...)
}

// ListLocations lists information about the supported locations for this service.
func (c *DeploymentResourcePoolClient) ListLocations(ctx context.Context, req *locationpb.ListLocationsRequest, opts ...gax.CallOption) *LocationIterator {
	return c.internalClient.ListLocations(ctx, req, opts...)
}

// GetIamPolicy gets the access control policy for a resource. Returns an empty policy
// if the resource exists and does not have a policy set.
func (c *DeploymentResourcePoolClient) GetIamPolicy(ctx context.Context, req *iampb.GetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	return c.internalClient.GetIamPolicy(ctx, req, opts...)
}

// SetIamPolicy sets the access control policy on the specified resource. Replaces
// any existing policy.
//
// Can return NOT_FOUND, INVALID_ARGUMENT, and PERMISSION_DENIED
// errors.
func (c *DeploymentResourcePoolClient) SetIamPolicy(ctx context.Context, req *iampb.SetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	return c.internalClient.SetIamPolicy(ctx, req, opts...)
}

// TestIamPermissions returns permissions that a caller has on the specified resource. If the
// resource does not exist, this will return an empty set of
// permissions, not a NOT_FOUND error.
//
// Note: This operation is designed to be used for building
// permission-aware UIs and command-line tools, not for authorization
// checking. This operation may “fail open” without warning.
func (c *DeploymentResourcePoolClient) TestIamPermissions(ctx context.Context, req *iampb.TestIamPermissionsRequest, opts ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error) {
	return c.internalClient.TestIamPermissions(ctx, req, opts...)
}

// CancelOperation is a utility method from google.longrunning.Operations.
func (c *DeploymentResourcePoolClient) CancelOperation(ctx context.Context, req *longrunningpb.CancelOperationRequest, opts ...gax.CallOption) error {
	return c.internalClient.CancelOperation(ctx, req, opts...)
}

// DeleteOperation is a utility method from google.longrunning.Operations.
func (c *DeploymentResourcePoolClient) DeleteOperation(ctx context.Context, req *longrunningpb.DeleteOperationRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteOperation(ctx, req, opts...)
}

// GetOperation is a utility method from google.longrunning.Operations.
func (c *DeploymentResourcePoolClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	return c.internalClient.GetOperation(ctx, req, opts...)
}

// ListOperations is a utility method from google.longrunning.Operations.
func (c *DeploymentResourcePoolClient) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	return c.internalClient.ListOperations(ctx, req, opts...)
}

// WaitOperation is a utility method from google.longrunning.Operations.
func (c *DeploymentResourcePoolClient) WaitOperation(ctx context.Context, req *longrunningpb.WaitOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	return c.internalClient.WaitOperation(ctx, req, opts...)
}

// deploymentResourcePoolGRPCClient is a client for interacting with Vertex AI API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type deploymentResourcePoolGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing DeploymentResourcePoolClient
	CallOptions **DeploymentResourcePoolCallOptions

	// The gRPC API client.
	deploymentResourcePoolClient aiplatformpb.DeploymentResourcePoolServiceClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	operationsClient longrunningpb.OperationsClient

	iamPolicyClient iampb.IAMPolicyClient

	locationsClient locationpb.LocationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string

	logger *slog.Logger
}

// NewDeploymentResourcePoolClient creates a new deployment resource pool service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// A service that manages the DeploymentResourcePool resource.
func NewDeploymentResourcePoolClient(ctx context.Context, opts ...option.ClientOption) (*DeploymentResourcePoolClient, error) {
	clientOpts := defaultDeploymentResourcePoolGRPCClientOptions()
	if newDeploymentResourcePoolClientHook != nil {
		hookOpts, err := newDeploymentResourcePoolClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := DeploymentResourcePoolClient{CallOptions: defaultDeploymentResourcePoolCallOptions()}

	c := &deploymentResourcePoolGRPCClient{
		connPool:                     connPool,
		deploymentResourcePoolClient: aiplatformpb.NewDeploymentResourcePoolServiceClient(connPool),
		CallOptions:                  &client.CallOptions,
		logger:                       internaloption.GetLogger(opts),
		operationsClient:             longrunningpb.NewOperationsClient(connPool),
		iamPolicyClient:              iampb.NewIAMPolicyClient(connPool),
		locationsClient:              locationpb.NewLocationsClient(connPool),
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
func (c *deploymentResourcePoolGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *deploymentResourcePoolGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version, "pb", protoVersion)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *deploymentResourcePoolGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *deploymentResourcePoolGRPCClient) CreateDeploymentResourcePool(ctx context.Context, req *aiplatformpb.CreateDeploymentResourcePoolRequest, opts ...gax.CallOption) (*CreateDeploymentResourcePoolOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).CreateDeploymentResourcePool[0:len((*c.CallOptions).CreateDeploymentResourcePool):len((*c.CallOptions).CreateDeploymentResourcePool)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.deploymentResourcePoolClient.CreateDeploymentResourcePool, req, settings.GRPC, c.logger, "CreateDeploymentResourcePool")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &CreateDeploymentResourcePoolOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *deploymentResourcePoolGRPCClient) GetDeploymentResourcePool(ctx context.Context, req *aiplatformpb.GetDeploymentResourcePoolRequest, opts ...gax.CallOption) (*aiplatformpb.DeploymentResourcePool, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetDeploymentResourcePool[0:len((*c.CallOptions).GetDeploymentResourcePool):len((*c.CallOptions).GetDeploymentResourcePool)], opts...)
	var resp *aiplatformpb.DeploymentResourcePool
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.deploymentResourcePoolClient.GetDeploymentResourcePool, req, settings.GRPC, c.logger, "GetDeploymentResourcePool")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *deploymentResourcePoolGRPCClient) ListDeploymentResourcePools(ctx context.Context, req *aiplatformpb.ListDeploymentResourcePoolsRequest, opts ...gax.CallOption) *DeploymentResourcePoolIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListDeploymentResourcePools[0:len((*c.CallOptions).ListDeploymentResourcePools):len((*c.CallOptions).ListDeploymentResourcePools)], opts...)
	it := &DeploymentResourcePoolIterator{}
	req = proto.Clone(req).(*aiplatformpb.ListDeploymentResourcePoolsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*aiplatformpb.DeploymentResourcePool, string, error) {
		resp := &aiplatformpb.ListDeploymentResourcePoolsResponse{}
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
			resp, err = executeRPC(ctx, c.deploymentResourcePoolClient.ListDeploymentResourcePools, req, settings.GRPC, c.logger, "ListDeploymentResourcePools")
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetDeploymentResourcePools(), resp.GetNextPageToken(), nil
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

func (c *deploymentResourcePoolGRPCClient) UpdateDeploymentResourcePool(ctx context.Context, req *aiplatformpb.UpdateDeploymentResourcePoolRequest, opts ...gax.CallOption) (*UpdateDeploymentResourcePoolOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "deployment_resource_pool.name", url.QueryEscape(req.GetDeploymentResourcePool().GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).UpdateDeploymentResourcePool[0:len((*c.CallOptions).UpdateDeploymentResourcePool):len((*c.CallOptions).UpdateDeploymentResourcePool)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.deploymentResourcePoolClient.UpdateDeploymentResourcePool, req, settings.GRPC, c.logger, "UpdateDeploymentResourcePool")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &UpdateDeploymentResourcePoolOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *deploymentResourcePoolGRPCClient) DeleteDeploymentResourcePool(ctx context.Context, req *aiplatformpb.DeleteDeploymentResourcePoolRequest, opts ...gax.CallOption) (*DeleteDeploymentResourcePoolOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).DeleteDeploymentResourcePool[0:len((*c.CallOptions).DeleteDeploymentResourcePool):len((*c.CallOptions).DeleteDeploymentResourcePool)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.deploymentResourcePoolClient.DeleteDeploymentResourcePool, req, settings.GRPC, c.logger, "DeleteDeploymentResourcePool")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &DeleteDeploymentResourcePoolOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *deploymentResourcePoolGRPCClient) QueryDeployedModels(ctx context.Context, req *aiplatformpb.QueryDeployedModelsRequest, opts ...gax.CallOption) *DeployedModelIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "deployment_resource_pool", url.QueryEscape(req.GetDeploymentResourcePool()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).QueryDeployedModels[0:len((*c.CallOptions).QueryDeployedModels):len((*c.CallOptions).QueryDeployedModels)], opts...)
	it := &DeployedModelIterator{}
	req = proto.Clone(req).(*aiplatformpb.QueryDeployedModelsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*aiplatformpb.DeployedModel, string, error) {
		resp := &aiplatformpb.QueryDeployedModelsResponse{}
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
			resp, err = executeRPC(ctx, c.deploymentResourcePoolClient.QueryDeployedModels, req, settings.GRPC, c.logger, "QueryDeployedModels")
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetDeployedModels(), resp.GetNextPageToken(), nil
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

func (c *deploymentResourcePoolGRPCClient) GetLocation(ctx context.Context, req *locationpb.GetLocationRequest, opts ...gax.CallOption) (*locationpb.Location, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetLocation[0:len((*c.CallOptions).GetLocation):len((*c.CallOptions).GetLocation)], opts...)
	var resp *locationpb.Location
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.locationsClient.GetLocation, req, settings.GRPC, c.logger, "GetLocation")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *deploymentResourcePoolGRPCClient) ListLocations(ctx context.Context, req *locationpb.ListLocationsRequest, opts ...gax.CallOption) *LocationIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListLocations[0:len((*c.CallOptions).ListLocations):len((*c.CallOptions).ListLocations)], opts...)
	it := &LocationIterator{}
	req = proto.Clone(req).(*locationpb.ListLocationsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*locationpb.Location, string, error) {
		resp := &locationpb.ListLocationsResponse{}
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
			resp, err = executeRPC(ctx, c.locationsClient.ListLocations, req, settings.GRPC, c.logger, "ListLocations")
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetLocations(), resp.GetNextPageToken(), nil
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

func (c *deploymentResourcePoolGRPCClient) GetIamPolicy(ctx context.Context, req *iampb.GetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetIamPolicy[0:len((*c.CallOptions).GetIamPolicy):len((*c.CallOptions).GetIamPolicy)], opts...)
	var resp *iampb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.iamPolicyClient.GetIamPolicy, req, settings.GRPC, c.logger, "GetIamPolicy")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *deploymentResourcePoolGRPCClient) SetIamPolicy(ctx context.Context, req *iampb.SetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).SetIamPolicy[0:len((*c.CallOptions).SetIamPolicy):len((*c.CallOptions).SetIamPolicy)], opts...)
	var resp *iampb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.iamPolicyClient.SetIamPolicy, req, settings.GRPC, c.logger, "SetIamPolicy")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *deploymentResourcePoolGRPCClient) TestIamPermissions(ctx context.Context, req *iampb.TestIamPermissionsRequest, opts ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).TestIamPermissions[0:len((*c.CallOptions).TestIamPermissions):len((*c.CallOptions).TestIamPermissions)], opts...)
	var resp *iampb.TestIamPermissionsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.iamPolicyClient.TestIamPermissions, req, settings.GRPC, c.logger, "TestIamPermissions")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *deploymentResourcePoolGRPCClient) CancelOperation(ctx context.Context, req *longrunningpb.CancelOperationRequest, opts ...gax.CallOption) error {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).CancelOperation[0:len((*c.CallOptions).CancelOperation):len((*c.CallOptions).CancelOperation)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = executeRPC(ctx, c.operationsClient.CancelOperation, req, settings.GRPC, c.logger, "CancelOperation")
		return err
	}, opts...)
	return err
}

func (c *deploymentResourcePoolGRPCClient) DeleteOperation(ctx context.Context, req *longrunningpb.DeleteOperationRequest, opts ...gax.CallOption) error {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).DeleteOperation[0:len((*c.CallOptions).DeleteOperation):len((*c.CallOptions).DeleteOperation)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = executeRPC(ctx, c.operationsClient.DeleteOperation, req, settings.GRPC, c.logger, "DeleteOperation")
		return err
	}, opts...)
	return err
}

func (c *deploymentResourcePoolGRPCClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
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

func (c *deploymentResourcePoolGRPCClient) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListOperations[0:len((*c.CallOptions).ListOperations):len((*c.CallOptions).ListOperations)], opts...)
	it := &OperationIterator{}
	req = proto.Clone(req).(*longrunningpb.ListOperationsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*longrunningpb.Operation, string, error) {
		resp := &longrunningpb.ListOperationsResponse{}
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
			resp, err = executeRPC(ctx, c.operationsClient.ListOperations, req, settings.GRPC, c.logger, "ListOperations")
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetOperations(), resp.GetNextPageToken(), nil
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

func (c *deploymentResourcePoolGRPCClient) WaitOperation(ctx context.Context, req *longrunningpb.WaitOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).WaitOperation[0:len((*c.CallOptions).WaitOperation):len((*c.CallOptions).WaitOperation)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.operationsClient.WaitOperation, req, settings.GRPC, c.logger, "WaitOperation")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CreateDeploymentResourcePoolOperation returns a new CreateDeploymentResourcePoolOperation from a given name.
// The name must be that of a previously created CreateDeploymentResourcePoolOperation, possibly from a different process.
func (c *deploymentResourcePoolGRPCClient) CreateDeploymentResourcePoolOperation(name string) *CreateDeploymentResourcePoolOperation {
	return &CreateDeploymentResourcePoolOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// DeleteDeploymentResourcePoolOperation returns a new DeleteDeploymentResourcePoolOperation from a given name.
// The name must be that of a previously created DeleteDeploymentResourcePoolOperation, possibly from a different process.
func (c *deploymentResourcePoolGRPCClient) DeleteDeploymentResourcePoolOperation(name string) *DeleteDeploymentResourcePoolOperation {
	return &DeleteDeploymentResourcePoolOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// UpdateDeploymentResourcePoolOperation returns a new UpdateDeploymentResourcePoolOperation from a given name.
// The name must be that of a previously created UpdateDeploymentResourcePoolOperation, possibly from a different process.
func (c *deploymentResourcePoolGRPCClient) UpdateDeploymentResourcePoolOperation(name string) *UpdateDeploymentResourcePoolOperation {
	return &UpdateDeploymentResourcePoolOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}
