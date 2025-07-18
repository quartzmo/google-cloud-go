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

package control_test

import (
	"context"

	control "cloud.google.com/go/storage/control/apiv2"
	controlpb "cloud.google.com/go/storage/control/apiv2/controlpb"
	"google.golang.org/api/iterator"
)

func ExampleNewStorageControlClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := control.NewStorageControlClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleNewStorageControlRESTClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := control.NewStorageControlRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleStorageControlClient_CreateAnywhereCache() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := control.NewStorageControlClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &controlpb.CreateAnywhereCacheRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/storage/control/apiv2/controlpb#CreateAnywhereCacheRequest.
	}
	op, err := c.CreateAnywhereCache(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleStorageControlClient_CreateFolder() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := control.NewStorageControlClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &controlpb.CreateFolderRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/storage/control/apiv2/controlpb#CreateFolderRequest.
	}
	resp, err := c.CreateFolder(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleStorageControlClient_CreateManagedFolder() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := control.NewStorageControlClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &controlpb.CreateManagedFolderRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/storage/control/apiv2/controlpb#CreateManagedFolderRequest.
	}
	resp, err := c.CreateManagedFolder(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleStorageControlClient_DeleteFolder() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := control.NewStorageControlClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &controlpb.DeleteFolderRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/storage/control/apiv2/controlpb#DeleteFolderRequest.
	}
	err = c.DeleteFolder(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleStorageControlClient_DeleteManagedFolder() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := control.NewStorageControlClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &controlpb.DeleteManagedFolderRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/storage/control/apiv2/controlpb#DeleteManagedFolderRequest.
	}
	err = c.DeleteManagedFolder(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleStorageControlClient_DisableAnywhereCache() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := control.NewStorageControlClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &controlpb.DisableAnywhereCacheRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/storage/control/apiv2/controlpb#DisableAnywhereCacheRequest.
	}
	resp, err := c.DisableAnywhereCache(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleStorageControlClient_GetAnywhereCache() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := control.NewStorageControlClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &controlpb.GetAnywhereCacheRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/storage/control/apiv2/controlpb#GetAnywhereCacheRequest.
	}
	resp, err := c.GetAnywhereCache(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleStorageControlClient_GetFolder() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := control.NewStorageControlClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &controlpb.GetFolderRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/storage/control/apiv2/controlpb#GetFolderRequest.
	}
	resp, err := c.GetFolder(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleStorageControlClient_GetFolderIntelligenceConfig() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := control.NewStorageControlClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &controlpb.GetFolderIntelligenceConfigRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/storage/control/apiv2/controlpb#GetFolderIntelligenceConfigRequest.
	}
	resp, err := c.GetFolderIntelligenceConfig(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleStorageControlClient_GetManagedFolder() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := control.NewStorageControlClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &controlpb.GetManagedFolderRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/storage/control/apiv2/controlpb#GetManagedFolderRequest.
	}
	resp, err := c.GetManagedFolder(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleStorageControlClient_GetOrganizationIntelligenceConfig() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := control.NewStorageControlClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &controlpb.GetOrganizationIntelligenceConfigRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/storage/control/apiv2/controlpb#GetOrganizationIntelligenceConfigRequest.
	}
	resp, err := c.GetOrganizationIntelligenceConfig(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleStorageControlClient_GetProjectIntelligenceConfig() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := control.NewStorageControlClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &controlpb.GetProjectIntelligenceConfigRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/storage/control/apiv2/controlpb#GetProjectIntelligenceConfigRequest.
	}
	resp, err := c.GetProjectIntelligenceConfig(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleStorageControlClient_GetStorageLayout() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := control.NewStorageControlClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &controlpb.GetStorageLayoutRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/storage/control/apiv2/controlpb#GetStorageLayoutRequest.
	}
	resp, err := c.GetStorageLayout(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleStorageControlClient_ListAnywhereCaches() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := control.NewStorageControlClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &controlpb.ListAnywhereCachesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/storage/control/apiv2/controlpb#ListAnywhereCachesRequest.
	}
	it := c.ListAnywhereCaches(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp

		// If you need to access the underlying RPC response,
		// you can do so by casting the `Response` as below.
		// Otherwise, remove this line. Only populated after
		// first call to Next(). Not safe for concurrent access.
		_ = it.Response.(*controlpb.ListAnywhereCachesResponse)
	}
}

func ExampleStorageControlClient_ListFolders() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := control.NewStorageControlClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &controlpb.ListFoldersRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/storage/control/apiv2/controlpb#ListFoldersRequest.
	}
	it := c.ListFolders(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp

		// If you need to access the underlying RPC response,
		// you can do so by casting the `Response` as below.
		// Otherwise, remove this line. Only populated after
		// first call to Next(). Not safe for concurrent access.
		_ = it.Response.(*controlpb.ListFoldersResponse)
	}
}

func ExampleStorageControlClient_ListManagedFolders() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := control.NewStorageControlClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &controlpb.ListManagedFoldersRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/storage/control/apiv2/controlpb#ListManagedFoldersRequest.
	}
	it := c.ListManagedFolders(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp

		// If you need to access the underlying RPC response,
		// you can do so by casting the `Response` as below.
		// Otherwise, remove this line. Only populated after
		// first call to Next(). Not safe for concurrent access.
		_ = it.Response.(*controlpb.ListManagedFoldersResponse)
	}
}

func ExampleStorageControlClient_PauseAnywhereCache() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := control.NewStorageControlClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &controlpb.PauseAnywhereCacheRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/storage/control/apiv2/controlpb#PauseAnywhereCacheRequest.
	}
	resp, err := c.PauseAnywhereCache(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleStorageControlClient_RenameFolder() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := control.NewStorageControlClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &controlpb.RenameFolderRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/storage/control/apiv2/controlpb#RenameFolderRequest.
	}
	op, err := c.RenameFolder(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleStorageControlClient_ResumeAnywhereCache() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := control.NewStorageControlClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &controlpb.ResumeAnywhereCacheRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/storage/control/apiv2/controlpb#ResumeAnywhereCacheRequest.
	}
	resp, err := c.ResumeAnywhereCache(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleStorageControlClient_UpdateAnywhereCache() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := control.NewStorageControlClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &controlpb.UpdateAnywhereCacheRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/storage/control/apiv2/controlpb#UpdateAnywhereCacheRequest.
	}
	op, err := c.UpdateAnywhereCache(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleStorageControlClient_UpdateFolderIntelligenceConfig() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := control.NewStorageControlClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &controlpb.UpdateFolderIntelligenceConfigRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/storage/control/apiv2/controlpb#UpdateFolderIntelligenceConfigRequest.
	}
	resp, err := c.UpdateFolderIntelligenceConfig(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleStorageControlClient_UpdateOrganizationIntelligenceConfig() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := control.NewStorageControlClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &controlpb.UpdateOrganizationIntelligenceConfigRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/storage/control/apiv2/controlpb#UpdateOrganizationIntelligenceConfigRequest.
	}
	resp, err := c.UpdateOrganizationIntelligenceConfig(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleStorageControlClient_UpdateProjectIntelligenceConfig() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := control.NewStorageControlClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &controlpb.UpdateProjectIntelligenceConfigRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/storage/control/apiv2/controlpb#UpdateProjectIntelligenceConfigRequest.
	}
	resp, err := c.UpdateProjectIntelligenceConfig(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
