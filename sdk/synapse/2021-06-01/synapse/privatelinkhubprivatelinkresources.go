package synapse

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
)

// PrivateLinkHubPrivateLinkResourcesClient is the azure Synapse Analytics Management Client
type PrivateLinkHubPrivateLinkResourcesClient struct {
	BaseClient
}

// NewPrivateLinkHubPrivateLinkResourcesClient creates an instance of the PrivateLinkHubPrivateLinkResourcesClient
// client.
func NewPrivateLinkHubPrivateLinkResourcesClient(subscriptionID string) PrivateLinkHubPrivateLinkResourcesClient {
	return NewPrivateLinkHubPrivateLinkResourcesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPrivateLinkHubPrivateLinkResourcesClientWithBaseURI creates an instance of the
// PrivateLinkHubPrivateLinkResourcesClient client using a custom endpoint.  Use this when interacting with an Azure
// cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewPrivateLinkHubPrivateLinkResourcesClientWithBaseURI(baseURI string, subscriptionID string) PrivateLinkHubPrivateLinkResourcesClient {
	return PrivateLinkHubPrivateLinkResourcesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get get private link resource in private link hub
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// privateLinkHubName - the name of the private link hub
// privateLinkResourceName - the name of the private link resource
func (client PrivateLinkHubPrivateLinkResourcesClient) Get(ctx context.Context, resourceGroupName string, privateLinkHubName string, privateLinkResourceName string) (result PrivateLinkResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateLinkHubPrivateLinkResourcesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("synapse.PrivateLinkHubPrivateLinkResourcesClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, privateLinkHubName, privateLinkResourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.PrivateLinkHubPrivateLinkResourcesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "synapse.PrivateLinkHubPrivateLinkResourcesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.PrivateLinkHubPrivateLinkResourcesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client PrivateLinkHubPrivateLinkResourcesClient) GetPreparer(ctx context.Context, resourceGroupName string, privateLinkHubName string, privateLinkResourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"privateLinkHubName":      autorest.Encode("path", privateLinkHubName),
		"privateLinkResourceName": autorest.Encode("path", privateLinkResourceName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/privateLinkHubs/{privateLinkHubName}/privateLinkResources/{privateLinkResourceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateLinkHubPrivateLinkResourcesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client PrivateLinkHubPrivateLinkResourcesClient) GetResponder(resp *http.Response) (result PrivateLinkResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List get all private link resources for a private link hub
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// privateLinkHubName - the name of the private link hub
func (client PrivateLinkHubPrivateLinkResourcesClient) List(ctx context.Context, resourceGroupName string, privateLinkHubName string) (result PrivateLinkResourceListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateLinkHubPrivateLinkResourcesClient.List")
		defer func() {
			sc := -1
			if result.plrlr.Response.Response != nil {
				sc = result.plrlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("synapse.PrivateLinkHubPrivateLinkResourcesClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, privateLinkHubName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.PrivateLinkHubPrivateLinkResourcesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.plrlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "synapse.PrivateLinkHubPrivateLinkResourcesClient", "List", resp, "Failure sending request")
		return
	}

	result.plrlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.PrivateLinkHubPrivateLinkResourcesClient", "List", resp, "Failure responding to request")
		return
	}
	if result.plrlr.hasNextLink() && result.plrlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client PrivateLinkHubPrivateLinkResourcesClient) ListPreparer(ctx context.Context, resourceGroupName string, privateLinkHubName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"privateLinkHubName": autorest.Encode("path", privateLinkHubName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/privateLinkHubs/{privateLinkHubName}/privateLinkResources", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateLinkHubPrivateLinkResourcesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client PrivateLinkHubPrivateLinkResourcesClient) ListResponder(resp *http.Response) (result PrivateLinkResourceListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client PrivateLinkHubPrivateLinkResourcesClient) listNextResults(ctx context.Context, lastResults PrivateLinkResourceListResult) (result PrivateLinkResourceListResult, err error) {
	req, err := lastResults.privateLinkResourceListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "synapse.PrivateLinkHubPrivateLinkResourcesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "synapse.PrivateLinkHubPrivateLinkResourcesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.PrivateLinkHubPrivateLinkResourcesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client PrivateLinkHubPrivateLinkResourcesClient) ListComplete(ctx context.Context, resourceGroupName string, privateLinkHubName string) (result PrivateLinkResourceListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateLinkHubPrivateLinkResourcesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, privateLinkHubName)
	return
}
