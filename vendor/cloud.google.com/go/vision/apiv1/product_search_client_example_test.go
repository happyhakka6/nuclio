// Copyright 2018 Google LLC
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

// AUTO-GENERATED CODE. DO NOT EDIT.

package vision_test

import (
	"context"

	"cloud.google.com/go/vision/apiv1"
	"google.golang.org/api/iterator"
	visionpb "google.golang.org/genproto/googleapis/cloud/vision/v1"
)

func ExampleNewProductSearchClient() {
	ctx := context.Background()
	c, err := vision.NewProductSearchClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use client.
	_ = c
}

func ExampleProductSearchClient_CreateProduct() {
	ctx := context.Background()
	c, err := vision.NewProductSearchClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &visionpb.CreateProductRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateProduct(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleProductSearchClient_ListProducts() {
	ctx := context.Background()
	c, err := vision.NewProductSearchClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &visionpb.ListProductsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListProducts(ctx, req)
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
	}
}

func ExampleProductSearchClient_GetProduct() {
	ctx := context.Background()
	c, err := vision.NewProductSearchClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &visionpb.GetProductRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetProduct(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleProductSearchClient_UpdateProduct() {
	ctx := context.Background()
	c, err := vision.NewProductSearchClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &visionpb.UpdateProductRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateProduct(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleProductSearchClient_DeleteProduct() {
	ctx := context.Background()
	c, err := vision.NewProductSearchClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &visionpb.DeleteProductRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteProduct(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleProductSearchClient_ListReferenceImages() {
	ctx := context.Background()
	c, err := vision.NewProductSearchClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &visionpb.ListReferenceImagesRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListReferenceImages(ctx, req)
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
	}
}

func ExampleProductSearchClient_GetReferenceImage() {
	ctx := context.Background()
	c, err := vision.NewProductSearchClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &visionpb.GetReferenceImageRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetReferenceImage(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleProductSearchClient_DeleteReferenceImage() {
	ctx := context.Background()
	c, err := vision.NewProductSearchClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &visionpb.DeleteReferenceImageRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteReferenceImage(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleProductSearchClient_CreateReferenceImage() {
	ctx := context.Background()
	c, err := vision.NewProductSearchClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &visionpb.CreateReferenceImageRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateReferenceImage(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleProductSearchClient_CreateProductSet() {
	ctx := context.Background()
	c, err := vision.NewProductSearchClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &visionpb.CreateProductSetRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateProductSet(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleProductSearchClient_ListProductSets() {
	ctx := context.Background()
	c, err := vision.NewProductSearchClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &visionpb.ListProductSetsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListProductSets(ctx, req)
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
	}
}

func ExampleProductSearchClient_GetProductSet() {
	ctx := context.Background()
	c, err := vision.NewProductSearchClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &visionpb.GetProductSetRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetProductSet(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleProductSearchClient_UpdateProductSet() {
	ctx := context.Background()
	c, err := vision.NewProductSearchClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &visionpb.UpdateProductSetRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateProductSet(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleProductSearchClient_DeleteProductSet() {
	ctx := context.Background()
	c, err := vision.NewProductSearchClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &visionpb.DeleteProductSetRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteProductSet(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleProductSearchClient_AddProductToProductSet() {
	ctx := context.Background()
	c, err := vision.NewProductSearchClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &visionpb.AddProductToProductSetRequest{
		// TODO: Fill request struct fields.
	}
	err = c.AddProductToProductSet(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleProductSearchClient_RemoveProductFromProductSet() {
	ctx := context.Background()
	c, err := vision.NewProductSearchClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &visionpb.RemoveProductFromProductSetRequest{
		// TODO: Fill request struct fields.
	}
	err = c.RemoveProductFromProductSet(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleProductSearchClient_ListProductsInProductSet() {
	ctx := context.Background()
	c, err := vision.NewProductSearchClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &visionpb.ListProductsInProductSetRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListProductsInProductSet(ctx, req)
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
	}
}

func ExampleProductSearchClient_ImportProductSets() {
	ctx := context.Background()
	c, err := vision.NewProductSearchClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &visionpb.ImportProductSetsRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.ImportProductSets(ctx, req)
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