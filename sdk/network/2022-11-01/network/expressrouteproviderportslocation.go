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

// ExpressRouteProviderPortsLocationClient is the network Client
type ExpressRouteProviderPortsLocationClient struct {
	BaseClient
}

// NewExpressRouteProviderPortsLocationClient creates an instance of the ExpressRouteProviderPortsLocationClient
// client.
func NewExpressRouteProviderPortsLocationClient(subscriptionID string) ExpressRouteProviderPortsLocationClient {
	return NewExpressRouteProviderPortsLocationClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewExpressRouteProviderPortsLocationClientWithBaseURI creates an instance of the
// ExpressRouteProviderPortsLocationClient client using a custom endpoint.  Use this when interacting with an Azure
// cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewExpressRouteProviderPortsLocationClientWithBaseURI(baseURI string, subscriptionID string) ExpressRouteProviderPortsLocationClient {
	return ExpressRouteProviderPortsLocationClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List retrieves all the ExpressRouteProviderPorts in a subscription.
// Parameters:
// filter - the filter to apply on the operation. For example, you can use $filter=location eq '{state}'.
func (client ExpressRouteProviderPortsLocationClient) List(ctx context.Context, filter string) (result ExpressRouteProviderPortListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExpressRouteProviderPortsLocationClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListPreparer(ctx, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteProviderPortsLocationClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.ExpressRouteProviderPortsLocationClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ExpressRouteProviderPortsLocationClient", "List", resp, "Failure responding to request")
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client ExpressRouteProviderPortsLocationClient) ListPreparer(ctx context.Context, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Network/expressRouteProviderPorts", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteProviderPortsLocationClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ExpressRouteProviderPortsLocationClient) ListResponder(resp *http.Response) (result ExpressRouteProviderPortListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
