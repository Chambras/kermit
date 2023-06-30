package appplatform

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

// ApplicationAcceleratorsClient is the REST API for Azure Spring Apps
type ApplicationAcceleratorsClient struct {
    BaseClient
}
// NewApplicationAcceleratorsClient creates an instance of the ApplicationAcceleratorsClient client.
func NewApplicationAcceleratorsClient(subscriptionID string) ApplicationAcceleratorsClient {
    return NewApplicationAcceleratorsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewApplicationAcceleratorsClientWithBaseURI creates an instance of the ApplicationAcceleratorsClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
    func NewApplicationAcceleratorsClientWithBaseURI(baseURI string, subscriptionID string) ApplicationAcceleratorsClient {
        return ApplicationAcceleratorsClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// CreateOrUpdate create or update the application accelerator.
    // Parameters:
        // resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
        // from the Azure Resource Manager API or the portal.
        // serviceName - the name of the Service resource.
        // applicationAcceleratorName - the name of the application accelerator.
        // applicationAcceleratorResource - the application accelerator for the create or update operation
func (client ApplicationAcceleratorsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string, applicationAcceleratorResource ApplicationAcceleratorResource) (result ApplicationAcceleratorsCreateOrUpdateFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ApplicationAcceleratorsClient.CreateOrUpdate")
        defer func() {
            sc := -1
        if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
        sc = result.FutureAPI.Response().StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        if err := validation.Validate([]validation.Validation{
        { TargetValue: serviceName,
         Constraints: []validation.Constraint{	{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-z][a-z0-9-]*[a-z0-9]$`, Chain: nil }}}}); err != nil {
        return result, validation.NewError("appplatform.ApplicationAcceleratorsClient", "CreateOrUpdate", err.Error())
        }

        req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serviceName, applicationAcceleratorName, applicationAcceleratorResource)
    if err != nil {
    err = autorest.NewErrorWithError(err, "appplatform.ApplicationAcceleratorsClient", "CreateOrUpdate", nil , "Failure preparing request")
    return
    }

        result, err = client.CreateOrUpdateSender(req)
        if err != nil {
        err = autorest.NewErrorWithError(err, "appplatform.ApplicationAcceleratorsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
        return
        }

    return
}

    // CreateOrUpdatePreparer prepares the CreateOrUpdate request.
    func (client ApplicationAcceleratorsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string, applicationAcceleratorResource ApplicationAcceleratorResource) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "applicationAcceleratorName": autorest.Encode("path",applicationAcceleratorName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "serviceName": autorest.Encode("path",serviceName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2023-05-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsContentType("application/json; charset=utf-8"),
autorest.AsPut(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/applicationAccelerators/{applicationAcceleratorName}",pathParameters),
autorest.WithJSON(applicationAcceleratorResource),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
    // http.Response Body if it receives an error.
    func (client ApplicationAcceleratorsClient) CreateOrUpdateSender(req *http.Request) (future ApplicationAcceleratorsCreateOrUpdateFuture, err error) {
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
    func (client ApplicationAcceleratorsClient) CreateOrUpdateResponder(resp *http.Response) (result ApplicationAcceleratorResource, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// Delete delete the application accelerator.
    // Parameters:
        // resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
        // from the Azure Resource Manager API or the portal.
        // serviceName - the name of the Service resource.
        // applicationAcceleratorName - the name of the application accelerator.
func (client ApplicationAcceleratorsClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string) (result ApplicationAcceleratorsDeleteFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ApplicationAcceleratorsClient.Delete")
        defer func() {
            sc := -1
        if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
        sc = result.FutureAPI.Response().StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        if err := validation.Validate([]validation.Validation{
        { TargetValue: serviceName,
         Constraints: []validation.Constraint{	{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-z][a-z0-9-]*[a-z0-9]$`, Chain: nil }}}}); err != nil {
        return result, validation.NewError("appplatform.ApplicationAcceleratorsClient", "Delete", err.Error())
        }

        req, err := client.DeletePreparer(ctx, resourceGroupName, serviceName, applicationAcceleratorName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "appplatform.ApplicationAcceleratorsClient", "Delete", nil , "Failure preparing request")
    return
    }

        result, err = client.DeleteSender(req)
        if err != nil {
        err = autorest.NewErrorWithError(err, "appplatform.ApplicationAcceleratorsClient", "Delete", result.Response(), "Failure sending request")
        return
        }

    return
}

    // DeletePreparer prepares the Delete request.
    func (client ApplicationAcceleratorsClient) DeletePreparer(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "applicationAcceleratorName": autorest.Encode("path",applicationAcceleratorName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "serviceName": autorest.Encode("path",serviceName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2023-05-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsDelete(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/applicationAccelerators/{applicationAcceleratorName}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client ApplicationAcceleratorsClient) DeleteSender(req *http.Request) (future ApplicationAcceleratorsDeleteFuture, err error) {
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
    func (client ApplicationAcceleratorsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted,http.StatusNoContent),
            autorest.ByClosing())
            result.Response = resp
            return
    }

// Get get the application accelerator.
    // Parameters:
        // resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
        // from the Azure Resource Manager API or the portal.
        // serviceName - the name of the Service resource.
        // applicationAcceleratorName - the name of the application accelerator.
func (client ApplicationAcceleratorsClient) Get(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string) (result ApplicationAcceleratorResource, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ApplicationAcceleratorsClient.Get")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        if err := validation.Validate([]validation.Validation{
        { TargetValue: serviceName,
         Constraints: []validation.Constraint{	{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-z][a-z0-9-]*[a-z0-9]$`, Chain: nil }}}}); err != nil {
        return result, validation.NewError("appplatform.ApplicationAcceleratorsClient", "Get", err.Error())
        }

        req, err := client.GetPreparer(ctx, resourceGroupName, serviceName, applicationAcceleratorName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "appplatform.ApplicationAcceleratorsClient", "Get", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "appplatform.ApplicationAcceleratorsClient", "Get", resp, "Failure sending request")
        return
        }

        result, err = client.GetResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "appplatform.ApplicationAcceleratorsClient", "Get", resp, "Failure responding to request")
        return
        }

    return
}

    // GetPreparer prepares the Get request.
    func (client ApplicationAcceleratorsClient) GetPreparer(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "applicationAcceleratorName": autorest.Encode("path",applicationAcceleratorName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "serviceName": autorest.Encode("path",serviceName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2023-05-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/applicationAccelerators/{applicationAcceleratorName}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client ApplicationAcceleratorsClient) GetSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // GetResponder handles the response to the Get request. The method always
    // closes the http.Response Body.
    func (client ApplicationAcceleratorsClient) GetResponder(resp *http.Response) (result ApplicationAcceleratorResource, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// List handle requests to list all application accelerator.
    // Parameters:
        // resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
        // from the Azure Resource Manager API or the portal.
        // serviceName - the name of the Service resource.
func (client ApplicationAcceleratorsClient) List(ctx context.Context, resourceGroupName string, serviceName string) (result ApplicationAcceleratorResourceCollectionPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ApplicationAcceleratorsClient.List")
        defer func() {
            sc := -1
        if result.aarc.Response.Response != nil {
        sc = result.aarc.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        if err := validation.Validate([]validation.Validation{
        { TargetValue: serviceName,
         Constraints: []validation.Constraint{	{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-z][a-z0-9-]*[a-z0-9]$`, Chain: nil }}}}); err != nil {
        return result, validation.NewError("appplatform.ApplicationAcceleratorsClient", "List", err.Error())
        }

            result.fn = client.listNextResults
    req, err := client.ListPreparer(ctx, resourceGroupName, serviceName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "appplatform.ApplicationAcceleratorsClient", "List", nil , "Failure preparing request")
    return
    }

        resp, err := client.ListSender(req)
        if err != nil {
        result.aarc.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "appplatform.ApplicationAcceleratorsClient", "List", resp, "Failure sending request")
        return
        }

        result.aarc, err = client.ListResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "appplatform.ApplicationAcceleratorsClient", "List", resp, "Failure responding to request")
        return
        }
            if result.aarc.hasNextLink() && result.aarc.IsEmpty() {
            err = result.NextWithContext(ctx)
            return
            }

    return
}

    // ListPreparer prepares the List request.
    func (client ApplicationAcceleratorsClient) ListPreparer(ctx context.Context, resourceGroupName string, serviceName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "serviceName": autorest.Encode("path",serviceName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2023-05-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/applicationAccelerators",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client ApplicationAcceleratorsClient) ListSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // ListResponder handles the response to the List request. The method always
    // closes the http.Response Body.
    func (client ApplicationAcceleratorsClient) ListResponder(resp *http.Response) (result ApplicationAcceleratorResourceCollection, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

            // listNextResults retrieves the next set of results, if any.
            func (client ApplicationAcceleratorsClient) listNextResults(ctx context.Context, lastResults ApplicationAcceleratorResourceCollection) (result ApplicationAcceleratorResourceCollection, err error) {
            req, err := lastResults.applicationAcceleratorResourceCollectionPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "appplatform.ApplicationAcceleratorsClient", "listNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "appplatform.ApplicationAcceleratorsClient", "listNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "appplatform.ApplicationAcceleratorsClient", "listNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

            // ListComplete enumerates all values, automatically crossing page boundaries as required.
            func (client ApplicationAcceleratorsClient) ListComplete(ctx context.Context, resourceGroupName string, serviceName string) (result ApplicationAcceleratorResourceCollectionIterator, err error) {
            if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/ApplicationAcceleratorsClient.List")
            defer func() {
            sc := -1
            if result.Response().Response.Response != nil {
            sc = result.page.Response().Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
            }()
            }
                    result.page, err = client.List(ctx, resourceGroupName, serviceName)
                            return
            }

