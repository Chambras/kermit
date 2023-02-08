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

// LoadBalancerBackendAddressPoolsClient is the network Client
type LoadBalancerBackendAddressPoolsClient struct {
	BaseClient
}

// NewLoadBalancerBackendAddressPoolsClient creates an instance of the LoadBalancerBackendAddressPoolsClient client.
func NewLoadBalancerBackendAddressPoolsClient(subscriptionID string) LoadBalancerBackendAddressPoolsClient {
	return NewLoadBalancerBackendAddressPoolsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewLoadBalancerBackendAddressPoolsClientWithBaseURI creates an instance of the LoadBalancerBackendAddressPoolsClient
// client using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI
// (sovereign clouds, Azure stack).
func NewLoadBalancerBackendAddressPoolsClientWithBaseURI(baseURI string, subscriptionID string) LoadBalancerBackendAddressPoolsClient {
	return LoadBalancerBackendAddressPoolsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a load balancer backend address pool.
// Parameters:
// resourceGroupName - the name of the resource group.
// loadBalancerName - the name of the load balancer.
// backendAddressPoolName - the name of the backend address pool.
// parameters - parameters supplied to the create or update load balancer backend address pool operation.
func (client LoadBalancerBackendAddressPoolsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, loadBalancerName string, backendAddressPoolName string, parameters BackendAddressPool) (result LoadBalancerBackendAddressPoolsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LoadBalancerBackendAddressPoolsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, loadBalancerName, backendAddressPoolName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LoadBalancerBackendAddressPoolsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LoadBalancerBackendAddressPoolsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client LoadBalancerBackendAddressPoolsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, loadBalancerName string, backendAddressPoolName string, parameters BackendAddressPool) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"backendAddressPoolName": autorest.Encode("path", backendAddressPoolName),
		"loadBalancerName":       autorest.Encode("path", loadBalancerName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	parameters.Etag = nil
	parameters.Type = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/backendAddressPools/{backendAddressPoolName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client LoadBalancerBackendAddressPoolsClient) CreateOrUpdateSender(req *http.Request) (future LoadBalancerBackendAddressPoolsCreateOrUpdateFuture, err error) {
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
func (client LoadBalancerBackendAddressPoolsClient) CreateOrUpdateResponder(resp *http.Response) (result BackendAddressPool, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified load balancer backend address pool.
// Parameters:
// resourceGroupName - the name of the resource group.
// loadBalancerName - the name of the load balancer.
// backendAddressPoolName - the name of the backend address pool.
func (client LoadBalancerBackendAddressPoolsClient) Delete(ctx context.Context, resourceGroupName string, loadBalancerName string, backendAddressPoolName string) (result LoadBalancerBackendAddressPoolsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LoadBalancerBackendAddressPoolsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, loadBalancerName, backendAddressPoolName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LoadBalancerBackendAddressPoolsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LoadBalancerBackendAddressPoolsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client LoadBalancerBackendAddressPoolsClient) DeletePreparer(ctx context.Context, resourceGroupName string, loadBalancerName string, backendAddressPoolName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"backendAddressPoolName": autorest.Encode("path", backendAddressPoolName),
		"loadBalancerName":       autorest.Encode("path", loadBalancerName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/backendAddressPools/{backendAddressPoolName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client LoadBalancerBackendAddressPoolsClient) DeleteSender(req *http.Request) (future LoadBalancerBackendAddressPoolsDeleteFuture, err error) {
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
func (client LoadBalancerBackendAddressPoolsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets load balancer backend address pool.
// Parameters:
// resourceGroupName - the name of the resource group.
// loadBalancerName - the name of the load balancer.
// backendAddressPoolName - the name of the backend address pool.
func (client LoadBalancerBackendAddressPoolsClient) Get(ctx context.Context, resourceGroupName string, loadBalancerName string, backendAddressPoolName string) (result BackendAddressPool, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LoadBalancerBackendAddressPoolsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, loadBalancerName, backendAddressPoolName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LoadBalancerBackendAddressPoolsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.LoadBalancerBackendAddressPoolsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LoadBalancerBackendAddressPoolsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client LoadBalancerBackendAddressPoolsClient) GetPreparer(ctx context.Context, resourceGroupName string, loadBalancerName string, backendAddressPoolName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"backendAddressPoolName": autorest.Encode("path", backendAddressPoolName),
		"loadBalancerName":       autorest.Encode("path", loadBalancerName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/backendAddressPools/{backendAddressPoolName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client LoadBalancerBackendAddressPoolsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client LoadBalancerBackendAddressPoolsClient) GetResponder(resp *http.Response) (result BackendAddressPool, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all the load balancer backed address pools.
// Parameters:
// resourceGroupName - the name of the resource group.
// loadBalancerName - the name of the load balancer.
func (client LoadBalancerBackendAddressPoolsClient) List(ctx context.Context, resourceGroupName string, loadBalancerName string) (result LoadBalancerBackendAddressPoolListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LoadBalancerBackendAddressPoolsClient.List")
		defer func() {
			sc := -1
			if result.lbbaplr.Response.Response != nil {
				sc = result.lbbaplr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, loadBalancerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LoadBalancerBackendAddressPoolsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.lbbaplr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.LoadBalancerBackendAddressPoolsClient", "List", resp, "Failure sending request")
		return
	}

	result.lbbaplr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LoadBalancerBackendAddressPoolsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.lbbaplr.hasNextLink() && result.lbbaplr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client LoadBalancerBackendAddressPoolsClient) ListPreparer(ctx context.Context, resourceGroupName string, loadBalancerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"loadBalancerName":  autorest.Encode("path", loadBalancerName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/backendAddressPools", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client LoadBalancerBackendAddressPoolsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client LoadBalancerBackendAddressPoolsClient) ListResponder(resp *http.Response) (result LoadBalancerBackendAddressPoolListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client LoadBalancerBackendAddressPoolsClient) listNextResults(ctx context.Context, lastResults LoadBalancerBackendAddressPoolListResult) (result LoadBalancerBackendAddressPoolListResult, err error) {
	req, err := lastResults.loadBalancerBackendAddressPoolListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.LoadBalancerBackendAddressPoolsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.LoadBalancerBackendAddressPoolsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LoadBalancerBackendAddressPoolsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client LoadBalancerBackendAddressPoolsClient) ListComplete(ctx context.Context, resourceGroupName string, loadBalancerName string) (result LoadBalancerBackendAddressPoolListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LoadBalancerBackendAddressPoolsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, loadBalancerName)
	return
}
