package graphrbac

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// SignedInUserClient is the the Graph RBAC Management Client
type SignedInUserClient struct {
	BaseClient
}

// NewSignedInUserClient creates an instance of the SignedInUserClient client.
func NewSignedInUserClient(tenantID string) SignedInUserClient {
	return NewSignedInUserClientWithBaseURI(DefaultBaseURI, tenantID)
}

// NewSignedInUserClientWithBaseURI creates an instance of the SignedInUserClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewSignedInUserClientWithBaseURI(baseURI string, tenantID string) SignedInUserClient {
	return SignedInUserClient{NewWithBaseURI(baseURI, tenantID)}
}

// Get gets the details for the currently logged-in user.
func (client SignedInUserClient) Get(ctx context.Context) (result User, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SignedInUserClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.SignedInUserClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "graphrbac.SignedInUserClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.SignedInUserClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client SignedInUserClient) GetPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"tenantID": autorest.Encode("path", client.TenantID),
	}

	const APIVersion = "1.6"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{tenantID}/me", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SignedInUserClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SignedInUserClient) GetResponder(resp *http.Response) (result User, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListOwnedObjects get the list of directory objects that are owned by the user.
func (client SignedInUserClient) ListOwnedObjects(ctx context.Context) (result DirectoryObjectListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SignedInUserClient.ListOwnedObjects")
		defer func() {
			sc := -1
			if result.dolr.Response.Response != nil {
				sc = result.dolr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = func(ctx context.Context, lastResult DirectoryObjectListResult) (DirectoryObjectListResult, error) {
		if lastResult.OdataNextLink == nil || len(to.String(lastResult.OdataNextLink)) < 1 {
			return DirectoryObjectListResult{}, nil
		}
		return client.ListOwnedObjectsNext(ctx, *lastResult.OdataNextLink)
	}
	req, err := client.ListOwnedObjectsPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.SignedInUserClient", "ListOwnedObjects", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListOwnedObjectsSender(req)
	if err != nil {
		result.dolr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "graphrbac.SignedInUserClient", "ListOwnedObjects", resp, "Failure sending request")
		return
	}

	result.dolr, err = client.ListOwnedObjectsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.SignedInUserClient", "ListOwnedObjects", resp, "Failure responding to request")
		return
	}
	if result.dolr.hasNextLink() && result.dolr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListOwnedObjectsPreparer prepares the ListOwnedObjects request.
func (client SignedInUserClient) ListOwnedObjectsPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"tenantID": autorest.Encode("path", client.TenantID),
	}

	const APIVersion = "1.6"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{tenantID}/me/ownedObjects", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListOwnedObjectsSender sends the ListOwnedObjects request. The method will close the
// http.Response Body if it receives an error.
func (client SignedInUserClient) ListOwnedObjectsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListOwnedObjectsResponder handles the response to the ListOwnedObjects request. The method always
// closes the http.Response Body.
func (client SignedInUserClient) ListOwnedObjectsResponder(resp *http.Response) (result DirectoryObjectListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListOwnedObjectsComplete enumerates all values, automatically crossing page boundaries as required.
func (client SignedInUserClient) ListOwnedObjectsComplete(ctx context.Context) (result DirectoryObjectListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SignedInUserClient.ListOwnedObjects")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListOwnedObjects(ctx)
	return
}

// ListOwnedObjectsNext get the list of directory objects that are owned by the user.
// Parameters:
// nextLink - next link for the list operation.
func (client SignedInUserClient) ListOwnedObjectsNext(ctx context.Context, nextLink string) (result DirectoryObjectListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SignedInUserClient.ListOwnedObjectsNext")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListOwnedObjectsNextPreparer(ctx, nextLink)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.SignedInUserClient", "ListOwnedObjectsNext", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListOwnedObjectsNextSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "graphrbac.SignedInUserClient", "ListOwnedObjectsNext", resp, "Failure sending request")
		return
	}

	result, err = client.ListOwnedObjectsNextResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.SignedInUserClient", "ListOwnedObjectsNext", resp, "Failure responding to request")
		return
	}

	return
}

// ListOwnedObjectsNextPreparer prepares the ListOwnedObjectsNext request.
func (client SignedInUserClient) ListOwnedObjectsNextPreparer(ctx context.Context, nextLink string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"nextLink": nextLink,
		"tenantID": autorest.Encode("path", client.TenantID),
	}

	const APIVersion = "1.6"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{tenantID}/{nextLink}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListOwnedObjectsNextSender sends the ListOwnedObjectsNext request. The method will close the
// http.Response Body if it receives an error.
func (client SignedInUserClient) ListOwnedObjectsNextSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListOwnedObjectsNextResponder handles the response to the ListOwnedObjectsNext request. The method always
// closes the http.Response Body.
func (client SignedInUserClient) ListOwnedObjectsNextResponder(resp *http.Response) (result DirectoryObjectListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
