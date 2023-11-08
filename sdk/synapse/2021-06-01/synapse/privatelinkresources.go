package synapse

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
    "github.com/Azure/go-autorest/autorest"
    "github.com/Azure/go-autorest/autorest/azure"
    "net/http"
    "context"
    "github.com/Azure/go-autorest/tracing"
    "github.com/Azure/go-autorest/autorest/validation"
)

// PrivateLinkResourcesClient is the azure Synapse Analytics Management Client
type PrivateLinkResourcesClient struct {
    BaseClient
}
// NewPrivateLinkResourcesClient creates an instance of the PrivateLinkResourcesClient client.
func NewPrivateLinkResourcesClient(subscriptionID string) PrivateLinkResourcesClient {
    return NewPrivateLinkResourcesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPrivateLinkResourcesClientWithBaseURI creates an instance of the PrivateLinkResourcesClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
    func NewPrivateLinkResourcesClientWithBaseURI(baseURI string, subscriptionID string) PrivateLinkResourcesClient {
        return PrivateLinkResourcesClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// Get get private link resource in workspace
    // Parameters:
        // resourceGroupName - the name of the resource group. The name is case insensitive.
        // workspaceName - the name of the workspace.
        // privateLinkResourceName - the name of the private link resource
func (client PrivateLinkResourcesClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, privateLinkResourceName string) (result PrivateLinkResource, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/PrivateLinkResourcesClient.Get")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        if err := validation.Validate([]validation.Validation{
        { TargetValue: client.SubscriptionID,
         Constraints: []validation.Constraint{	{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil }}},
        { TargetValue: resourceGroupName,
         Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil },
        	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil }}}}); err != nil {
        return result, validation.NewError("synapse.PrivateLinkResourcesClient", "Get", err.Error())
        }

        req, err := client.GetPreparer(ctx, resourceGroupName, workspaceName, privateLinkResourceName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "synapse.PrivateLinkResourcesClient", "Get", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "synapse.PrivateLinkResourcesClient", "Get", resp, "Failure sending request")
        return
        }

        result, err = client.GetResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "synapse.PrivateLinkResourcesClient", "Get", resp, "Failure responding to request")
        return
        }

    return
}

    // GetPreparer prepares the Get request.
    func (client PrivateLinkResourcesClient) GetPreparer(ctx context.Context, resourceGroupName string, workspaceName string, privateLinkResourceName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "privateLinkResourceName": autorest.Encode("path",privateLinkResourceName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        "workspaceName": autorest.Encode("path",workspaceName),
        }

            const APIVersion = "2021-06-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/privateLinkResources/{privateLinkResourceName}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client PrivateLinkResourcesClient) GetSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // GetResponder handles the response to the Get request. The method always
    // closes the http.Response Body.
    func (client PrivateLinkResourcesClient) GetResponder(resp *http.Response) (result PrivateLinkResource, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// List get all private link resources for a workspaces
    // Parameters:
        // resourceGroupName - the name of the resource group. The name is case insensitive.
        // workspaceName - the name of the workspace.
func (client PrivateLinkResourcesClient) List(ctx context.Context, resourceGroupName string, workspaceName string) (result PrivateLinkResourceListResultPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/PrivateLinkResourcesClient.List")
        defer func() {
            sc := -1
        if result.plrlr.Response.Response != nil {
        sc = result.plrlr.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        if err := validation.Validate([]validation.Validation{
        { TargetValue: client.SubscriptionID,
         Constraints: []validation.Constraint{	{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil }}},
        { TargetValue: resourceGroupName,
         Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil },
        	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil }}}}); err != nil {
        return result, validation.NewError("synapse.PrivateLinkResourcesClient", "List", err.Error())
        }

            result.fn = client.listNextResults
    req, err := client.ListPreparer(ctx, resourceGroupName, workspaceName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "synapse.PrivateLinkResourcesClient", "List", nil , "Failure preparing request")
    return
    }

        resp, err := client.ListSender(req)
        if err != nil {
        result.plrlr.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "synapse.PrivateLinkResourcesClient", "List", resp, "Failure sending request")
        return
        }

        result.plrlr, err = client.ListResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "synapse.PrivateLinkResourcesClient", "List", resp, "Failure responding to request")
        return
        }
            if result.plrlr.hasNextLink() && result.plrlr.IsEmpty() {
            err = result.NextWithContext(ctx)
            return
            }

    return
}

    // ListPreparer prepares the List request.
    func (client PrivateLinkResourcesClient) ListPreparer(ctx context.Context, resourceGroupName string, workspaceName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        "workspaceName": autorest.Encode("path",workspaceName),
        }

            const APIVersion = "2021-06-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/privateLinkResources",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client PrivateLinkResourcesClient) ListSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // ListResponder handles the response to the List request. The method always
    // closes the http.Response Body.
    func (client PrivateLinkResourcesClient) ListResponder(resp *http.Response) (result PrivateLinkResourceListResult, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

            // listNextResults retrieves the next set of results, if any.
            func (client PrivateLinkResourcesClient) listNextResults(ctx context.Context, lastResults PrivateLinkResourceListResult) (result PrivateLinkResourceListResult, err error) {
            req, err := lastResults.privateLinkResourceListResultPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "synapse.PrivateLinkResourcesClient", "listNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "synapse.PrivateLinkResourcesClient", "listNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "synapse.PrivateLinkResourcesClient", "listNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

            // ListComplete enumerates all values, automatically crossing page boundaries as required.
            func (client PrivateLinkResourcesClient) ListComplete(ctx context.Context, resourceGroupName string, workspaceName string) (result PrivateLinkResourceListResultIterator, err error) {
            if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/PrivateLinkResourcesClient.List")
            defer func() {
            sc := -1
            if result.Response().Response.Response != nil {
            sc = result.page.Response().Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
            }()
            }
                    result.page, err = client.List(ctx, resourceGroupName, workspaceName)
                            return
            }

