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
)

// ManagedPrivateEndpointsClient is the client for the ManagedPrivateEndpoints methods of the Synapse service.
type ManagedPrivateEndpointsClient struct {
    BaseClient
}
// NewManagedPrivateEndpointsClient creates an instance of the ManagedPrivateEndpointsClient client.
func NewManagedPrivateEndpointsClient(endpoint string) ManagedPrivateEndpointsClient {
    return ManagedPrivateEndpointsClient{ New(endpoint)}
}

// Create create Managed Private Endpoints
    // Parameters:
        // managedVirtualNetworkName - managed virtual network name
        // managedPrivateEndpointName - managed private endpoint name
        // managedPrivateEndpoint - managed private endpoint properties.
func (client ManagedPrivateEndpointsClient) Create(ctx context.Context, managedVirtualNetworkName string, managedPrivateEndpointName string, managedPrivateEndpoint ManagedPrivateEndpoint) (result ManagedPrivateEndpoint, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ManagedPrivateEndpointsClient.Create")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.CreatePreparer(ctx, managedVirtualNetworkName, managedPrivateEndpointName, managedPrivateEndpoint)
    if err != nil {
    err = autorest.NewErrorWithError(err, "synapse.ManagedPrivateEndpointsClient", "Create", nil , "Failure preparing request")
    return
    }

        resp, err := client.CreateSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "synapse.ManagedPrivateEndpointsClient", "Create", resp, "Failure sending request")
        return
        }

        result, err = client.CreateResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "synapse.ManagedPrivateEndpointsClient", "Create", resp, "Failure responding to request")
        return
        }

    return
}

    // CreatePreparer prepares the Create request.
    func (client ManagedPrivateEndpointsClient) CreatePreparer(ctx context.Context, managedVirtualNetworkName string, managedPrivateEndpointName string, managedPrivateEndpoint ManagedPrivateEndpoint) (*http.Request, error) {
        urlParameters := map[string]interface{} {
        "endpoint": client.Endpoint,
        }

        pathParameters := map[string]interface{} {
        "managedPrivateEndpointName": autorest.Encode("path",managedPrivateEndpointName),
        "managedVirtualNetworkName": autorest.Encode("path",managedVirtualNetworkName),
        }

            const APIVersion = "2019-06-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

            managedPrivateEndpoint.ID = nil
            managedPrivateEndpoint.Name = nil
            managedPrivateEndpoint.Type = nil
    preparer := autorest.CreatePreparer(
autorest.AsContentType("application/json; charset=utf-8"),
autorest.AsPut(),
autorest.WithCustomBaseURL("{endpoint}", urlParameters),
autorest.WithPathParameters("/managedVirtualNetworks/{managedVirtualNetworkName}/managedPrivateEndpoints/{managedPrivateEndpointName}",pathParameters),
autorest.WithJSON(managedPrivateEndpoint),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateSender sends the Create request. The method will close the
    // http.Response Body if it receives an error.
    func (client ManagedPrivateEndpointsClient) CreateSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
                }

    // CreateResponder handles the response to the Create request. The method always
    // closes the http.Response Body.
    func (client ManagedPrivateEndpointsClient) CreateResponder(resp *http.Response) (result ManagedPrivateEndpoint, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// Delete delete Managed Private Endpoints
    // Parameters:
        // managedVirtualNetworkName - managed virtual network name
        // managedPrivateEndpointName - managed private endpoint name
func (client ManagedPrivateEndpointsClient) Delete(ctx context.Context, managedVirtualNetworkName string, managedPrivateEndpointName string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ManagedPrivateEndpointsClient.Delete")
        defer func() {
            sc := -1
        if result.Response != nil {
        sc = result.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.DeletePreparer(ctx, managedVirtualNetworkName, managedPrivateEndpointName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "synapse.ManagedPrivateEndpointsClient", "Delete", nil , "Failure preparing request")
    return
    }

        resp, err := client.DeleteSender(req)
        if err != nil {
        result.Response = resp
        err = autorest.NewErrorWithError(err, "synapse.ManagedPrivateEndpointsClient", "Delete", resp, "Failure sending request")
        return
        }

        result, err = client.DeleteResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "synapse.ManagedPrivateEndpointsClient", "Delete", resp, "Failure responding to request")
        return
        }

    return
}

    // DeletePreparer prepares the Delete request.
    func (client ManagedPrivateEndpointsClient) DeletePreparer(ctx context.Context, managedVirtualNetworkName string, managedPrivateEndpointName string) (*http.Request, error) {
        urlParameters := map[string]interface{} {
        "endpoint": client.Endpoint,
        }

        pathParameters := map[string]interface{} {
        "managedPrivateEndpointName": autorest.Encode("path",managedPrivateEndpointName),
        "managedVirtualNetworkName": autorest.Encode("path",managedVirtualNetworkName),
        }

            const APIVersion = "2019-06-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsDelete(),
autorest.WithCustomBaseURL("{endpoint}", urlParameters),
autorest.WithPathParameters("/managedVirtualNetworks/{managedVirtualNetworkName}/managedPrivateEndpoints/{managedPrivateEndpointName}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client ManagedPrivateEndpointsClient) DeleteSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
                }

    // DeleteResponder handles the response to the Delete request. The method always
    // closes the http.Response Body.
    func (client ManagedPrivateEndpointsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted,http.StatusNoContent),
            autorest.ByClosing())
            result.Response = resp
            return
    }

// Get get Managed Private Endpoints
    // Parameters:
        // managedVirtualNetworkName - managed virtual network name
        // managedPrivateEndpointName - managed private endpoint name
func (client ManagedPrivateEndpointsClient) Get(ctx context.Context, managedVirtualNetworkName string, managedPrivateEndpointName string) (result ManagedPrivateEndpoint, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ManagedPrivateEndpointsClient.Get")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.GetPreparer(ctx, managedVirtualNetworkName, managedPrivateEndpointName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "synapse.ManagedPrivateEndpointsClient", "Get", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "synapse.ManagedPrivateEndpointsClient", "Get", resp, "Failure sending request")
        return
        }

        result, err = client.GetResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "synapse.ManagedPrivateEndpointsClient", "Get", resp, "Failure responding to request")
        return
        }

    return
}

    // GetPreparer prepares the Get request.
    func (client ManagedPrivateEndpointsClient) GetPreparer(ctx context.Context, managedVirtualNetworkName string, managedPrivateEndpointName string) (*http.Request, error) {
        urlParameters := map[string]interface{} {
        "endpoint": client.Endpoint,
        }

        pathParameters := map[string]interface{} {
        "managedPrivateEndpointName": autorest.Encode("path",managedPrivateEndpointName),
        "managedVirtualNetworkName": autorest.Encode("path",managedVirtualNetworkName),
        }

            const APIVersion = "2019-06-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithCustomBaseURL("{endpoint}", urlParameters),
autorest.WithPathParameters("/managedVirtualNetworks/{managedVirtualNetworkName}/managedPrivateEndpoints/{managedPrivateEndpointName}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client ManagedPrivateEndpointsClient) GetSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
                }

    // GetResponder handles the response to the Get request. The method always
    // closes the http.Response Body.
    func (client ManagedPrivateEndpointsClient) GetResponder(resp *http.Response) (result ManagedPrivateEndpoint, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// List list Managed Private Endpoints
    // Parameters:
        // managedVirtualNetworkName - managed virtual network name
func (client ManagedPrivateEndpointsClient) List(ctx context.Context, managedVirtualNetworkName string) (result ManagedPrivateEndpointListResponsePage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ManagedPrivateEndpointsClient.List")
        defer func() {
            sc := -1
        if result.mpelr.Response.Response != nil {
        sc = result.mpelr.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        result.fn = client.listNextResults
    req, err := client.ListPreparer(ctx, managedVirtualNetworkName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "synapse.ManagedPrivateEndpointsClient", "List", nil , "Failure preparing request")
    return
    }

        resp, err := client.ListSender(req)
        if err != nil {
        result.mpelr.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "synapse.ManagedPrivateEndpointsClient", "List", resp, "Failure sending request")
        return
        }

        result.mpelr, err = client.ListResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "synapse.ManagedPrivateEndpointsClient", "List", resp, "Failure responding to request")
        return
        }
            if result.mpelr.hasNextLink() && result.mpelr.IsEmpty() {
            err = result.NextWithContext(ctx)
            return
            }

    return
}

    // ListPreparer prepares the List request.
    func (client ManagedPrivateEndpointsClient) ListPreparer(ctx context.Context, managedVirtualNetworkName string) (*http.Request, error) {
        urlParameters := map[string]interface{} {
        "endpoint": client.Endpoint,
        }

        pathParameters := map[string]interface{} {
        "managedVirtualNetworkName": autorest.Encode("path",managedVirtualNetworkName),
        }

            const APIVersion = "2019-06-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithCustomBaseURL("{endpoint}", urlParameters),
autorest.WithPathParameters("/managedVirtualNetworks/{managedVirtualNetworkName}/managedPrivateEndpoints",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client ManagedPrivateEndpointsClient) ListSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
                }

    // ListResponder handles the response to the List request. The method always
    // closes the http.Response Body.
    func (client ManagedPrivateEndpointsClient) ListResponder(resp *http.Response) (result ManagedPrivateEndpointListResponse, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

            // listNextResults retrieves the next set of results, if any.
            func (client ManagedPrivateEndpointsClient) listNextResults(ctx context.Context, lastResults ManagedPrivateEndpointListResponse) (result ManagedPrivateEndpointListResponse, err error) {
            req, err := lastResults.managedPrivateEndpointListResponsePreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "synapse.ManagedPrivateEndpointsClient", "listNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "synapse.ManagedPrivateEndpointsClient", "listNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "synapse.ManagedPrivateEndpointsClient", "listNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

            // ListComplete enumerates all values, automatically crossing page boundaries as required.
            func (client ManagedPrivateEndpointsClient) ListComplete(ctx context.Context, managedVirtualNetworkName string) (result ManagedPrivateEndpointListResponseIterator, err error) {
            if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/ManagedPrivateEndpointsClient.List")
            defer func() {
            sc := -1
            if result.Response().Response.Response != nil {
            sc = result.page.Response().Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
            }()
            }
                    result.page, err = client.List(ctx, managedVirtualNetworkName)
                            return
            }

