package network

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
	"github.com/Azure/go-autorest/tracing"
)

// HubVirtualNetworkConnectionsClient is the network Client
type HubVirtualNetworkConnectionsClient struct {
	BaseClient
}

// NewHubVirtualNetworkConnectionsClient creates an instance of the HubVirtualNetworkConnectionsClient client.
func NewHubVirtualNetworkConnectionsClient(subscriptionID string) HubVirtualNetworkConnectionsClient {
	return NewHubVirtualNetworkConnectionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewHubVirtualNetworkConnectionsClientWithBaseURI creates an instance of the HubVirtualNetworkConnectionsClient
// client using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI
// (sovereign clouds, Azure stack).
func NewHubVirtualNetworkConnectionsClientWithBaseURI(baseURI string, subscriptionID string) HubVirtualNetworkConnectionsClient {
	return HubVirtualNetworkConnectionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates a hub virtual network connection if it doesn't exist else updates the existing one.
// Parameters:
// resourceGroupName - the resource group name of the HubVirtualNetworkConnection.
// virtualHubName - the name of the VirtualHub.
// connectionName - the name of the HubVirtualNetworkConnection.
// hubVirtualNetworkConnectionParameters - parameters supplied to create or update a hub virtual network
// connection.
func (client HubVirtualNetworkConnectionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, virtualHubName string, connectionName string, hubVirtualNetworkConnectionParameters HubVirtualNetworkConnection) (result HubVirtualNetworkConnectionsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/HubVirtualNetworkConnectionsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, virtualHubName, connectionName, hubVirtualNetworkConnectionParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.HubVirtualNetworkConnectionsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.HubVirtualNetworkConnectionsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client HubVirtualNetworkConnectionsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, virtualHubName string, connectionName string, hubVirtualNetworkConnectionParameters HubVirtualNetworkConnection) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"connectionName":    autorest.Encode("path", connectionName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"virtualHubName":    autorest.Encode("path", virtualHubName),
	}

	const APIVersion = "2022-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	hubVirtualNetworkConnectionParameters.Etag = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/hubVirtualNetworkConnections/{connectionName}", pathParameters),
		autorest.WithJSON(hubVirtualNetworkConnectionParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client HubVirtualNetworkConnectionsClient) CreateOrUpdateSender(req *http.Request) (future HubVirtualNetworkConnectionsCreateOrUpdateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client HubVirtualNetworkConnectionsClient) CreateOrUpdateResponder(resp *http.Response) (result HubVirtualNetworkConnection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a HubVirtualNetworkConnection.
// Parameters:
// resourceGroupName - the resource group name of the VirtualHub.
// virtualHubName - the name of the VirtualHub.
// connectionName - the name of the HubVirtualNetworkConnection.
func (client HubVirtualNetworkConnectionsClient) Delete(ctx context.Context, resourceGroupName string, virtualHubName string, connectionName string) (result HubVirtualNetworkConnectionsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/HubVirtualNetworkConnectionsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, virtualHubName, connectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.HubVirtualNetworkConnectionsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.HubVirtualNetworkConnectionsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client HubVirtualNetworkConnectionsClient) DeletePreparer(ctx context.Context, resourceGroupName string, virtualHubName string, connectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"connectionName":    autorest.Encode("path", connectionName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"virtualHubName":    autorest.Encode("path", virtualHubName),
	}

	const APIVersion = "2022-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/hubVirtualNetworkConnections/{connectionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client HubVirtualNetworkConnectionsClient) DeleteSender(req *http.Request) (future HubVirtualNetworkConnectionsDeleteFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client HubVirtualNetworkConnectionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get retrieves the details of a HubVirtualNetworkConnection.
// Parameters:
// resourceGroupName - the resource group name of the VirtualHub.
// virtualHubName - the name of the VirtualHub.
// connectionName - the name of the vpn connection.
func (client HubVirtualNetworkConnectionsClient) Get(ctx context.Context, resourceGroupName string, virtualHubName string, connectionName string) (result HubVirtualNetworkConnection, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/HubVirtualNetworkConnectionsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, virtualHubName, connectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.HubVirtualNetworkConnectionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.HubVirtualNetworkConnectionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.HubVirtualNetworkConnectionsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client HubVirtualNetworkConnectionsClient) GetPreparer(ctx context.Context, resourceGroupName string, virtualHubName string, connectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"connectionName":    autorest.Encode("path", connectionName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"virtualHubName":    autorest.Encode("path", virtualHubName),
	}

	const APIVersion = "2022-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/hubVirtualNetworkConnections/{connectionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client HubVirtualNetworkConnectionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client HubVirtualNetworkConnectionsClient) GetResponder(resp *http.Response) (result HubVirtualNetworkConnection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List retrieves the details of all HubVirtualNetworkConnections.
// Parameters:
// resourceGroupName - the resource group name of the VirtualHub.
// virtualHubName - the name of the VirtualHub.
func (client HubVirtualNetworkConnectionsClient) List(ctx context.Context, resourceGroupName string, virtualHubName string) (result ListHubVirtualNetworkConnectionsResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/HubVirtualNetworkConnectionsClient.List")
		defer func() {
			sc := -1
			if result.lhvncr.Response.Response != nil {
				sc = result.lhvncr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, virtualHubName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.HubVirtualNetworkConnectionsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.lhvncr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.HubVirtualNetworkConnectionsClient", "List", resp, "Failure sending request")
		return
	}

	result.lhvncr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.HubVirtualNetworkConnectionsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.lhvncr.hasNextLink() && result.lhvncr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client HubVirtualNetworkConnectionsClient) ListPreparer(ctx context.Context, resourceGroupName string, virtualHubName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"virtualHubName":    autorest.Encode("path", virtualHubName),
	}

	const APIVersion = "2022-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/hubVirtualNetworkConnections", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client HubVirtualNetworkConnectionsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client HubVirtualNetworkConnectionsClient) ListResponder(resp *http.Response) (result ListHubVirtualNetworkConnectionsResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client HubVirtualNetworkConnectionsClient) listNextResults(ctx context.Context, lastResults ListHubVirtualNetworkConnectionsResult) (result ListHubVirtualNetworkConnectionsResult, err error) {
	req, err := lastResults.listHubVirtualNetworkConnectionsResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.HubVirtualNetworkConnectionsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.HubVirtualNetworkConnectionsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.HubVirtualNetworkConnectionsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client HubVirtualNetworkConnectionsClient) ListComplete(ctx context.Context, resourceGroupName string, virtualHubName string) (result ListHubVirtualNetworkConnectionsResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/HubVirtualNetworkConnectionsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, virtualHubName)
	return
}
