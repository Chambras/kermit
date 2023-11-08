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

// PredefinedAcceleratorsClient is the REST API for Azure Spring Apps
type PredefinedAcceleratorsClient struct {
    BaseClient
}
// NewPredefinedAcceleratorsClient creates an instance of the PredefinedAcceleratorsClient client.
func NewPredefinedAcceleratorsClient(subscriptionID string) PredefinedAcceleratorsClient {
    return NewPredefinedAcceleratorsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPredefinedAcceleratorsClientWithBaseURI creates an instance of the PredefinedAcceleratorsClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
    func NewPredefinedAcceleratorsClientWithBaseURI(baseURI string, subscriptionID string) PredefinedAcceleratorsClient {
        return PredefinedAcceleratorsClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// Disable disable predefined accelerator.
    // Parameters:
        // resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
        // from the Azure Resource Manager API or the portal.
        // serviceName - the name of the Service resource.
        // applicationAcceleratorName - the name of the application accelerator.
        // predefinedAcceleratorName - the name of the predefined accelerator.
func (client PredefinedAcceleratorsClient) Disable(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string, predefinedAcceleratorName string) (result PredefinedAcceleratorsDisableFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/PredefinedAcceleratorsClient.Disable")
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
        return result, validation.NewError("appplatform.PredefinedAcceleratorsClient", "Disable", err.Error())
        }

        req, err := client.DisablePreparer(ctx, resourceGroupName, serviceName, applicationAcceleratorName, predefinedAcceleratorName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "appplatform.PredefinedAcceleratorsClient", "Disable", nil , "Failure preparing request")
    return
    }

        result, err = client.DisableSender(req)
        if err != nil {
        err = autorest.NewErrorWithError(err, "appplatform.PredefinedAcceleratorsClient", "Disable", result.Response(), "Failure sending request")
        return
        }

    return
}

    // DisablePreparer prepares the Disable request.
    func (client PredefinedAcceleratorsClient) DisablePreparer(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string, predefinedAcceleratorName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "applicationAcceleratorName": autorest.Encode("path",applicationAcceleratorName),
        "predefinedAcceleratorName": autorest.Encode("path",predefinedAcceleratorName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "serviceName": autorest.Encode("path",serviceName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2023-03-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsPost(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/applicationAccelerators/{applicationAcceleratorName}/predefinedAccelerators/{predefinedAcceleratorName}/disable",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DisableSender sends the Disable request. The method will close the
    // http.Response Body if it receives an error.
    func (client PredefinedAcceleratorsClient) DisableSender(req *http.Request) (future PredefinedAcceleratorsDisableFuture, err error) {
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

    // DisableResponder handles the response to the Disable request. The method always
    // closes the http.Response Body.
    func (client PredefinedAcceleratorsClient) DisableResponder(resp *http.Response) (result autorest.Response, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted),
            autorest.ByClosing())
            result.Response = resp
            return
    }

// Enable enable predefined accelerator.
    // Parameters:
        // resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
        // from the Azure Resource Manager API or the portal.
        // serviceName - the name of the Service resource.
        // applicationAcceleratorName - the name of the application accelerator.
        // predefinedAcceleratorName - the name of the predefined accelerator.
func (client PredefinedAcceleratorsClient) Enable(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string, predefinedAcceleratorName string) (result PredefinedAcceleratorsEnableFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/PredefinedAcceleratorsClient.Enable")
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
        return result, validation.NewError("appplatform.PredefinedAcceleratorsClient", "Enable", err.Error())
        }

        req, err := client.EnablePreparer(ctx, resourceGroupName, serviceName, applicationAcceleratorName, predefinedAcceleratorName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "appplatform.PredefinedAcceleratorsClient", "Enable", nil , "Failure preparing request")
    return
    }

        result, err = client.EnableSender(req)
        if err != nil {
        err = autorest.NewErrorWithError(err, "appplatform.PredefinedAcceleratorsClient", "Enable", result.Response(), "Failure sending request")
        return
        }

    return
}

    // EnablePreparer prepares the Enable request.
    func (client PredefinedAcceleratorsClient) EnablePreparer(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string, predefinedAcceleratorName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "applicationAcceleratorName": autorest.Encode("path",applicationAcceleratorName),
        "predefinedAcceleratorName": autorest.Encode("path",predefinedAcceleratorName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "serviceName": autorest.Encode("path",serviceName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2023-03-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsPost(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/applicationAccelerators/{applicationAcceleratorName}/predefinedAccelerators/{predefinedAcceleratorName}/enable",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // EnableSender sends the Enable request. The method will close the
    // http.Response Body if it receives an error.
    func (client PredefinedAcceleratorsClient) EnableSender(req *http.Request) (future PredefinedAcceleratorsEnableFuture, err error) {
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

    // EnableResponder handles the response to the Enable request. The method always
    // closes the http.Response Body.
    func (client PredefinedAcceleratorsClient) EnableResponder(resp *http.Response) (result autorest.Response, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted),
            autorest.ByClosing())
            result.Response = resp
            return
    }

// Get get the predefined accelerator.
    // Parameters:
        // resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
        // from the Azure Resource Manager API or the portal.
        // serviceName - the name of the Service resource.
        // applicationAcceleratorName - the name of the application accelerator.
        // predefinedAcceleratorName - the name of the predefined accelerator.
func (client PredefinedAcceleratorsClient) Get(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string, predefinedAcceleratorName string) (result PredefinedAcceleratorResource, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/PredefinedAcceleratorsClient.Get")
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
        return result, validation.NewError("appplatform.PredefinedAcceleratorsClient", "Get", err.Error())
        }

        req, err := client.GetPreparer(ctx, resourceGroupName, serviceName, applicationAcceleratorName, predefinedAcceleratorName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "appplatform.PredefinedAcceleratorsClient", "Get", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "appplatform.PredefinedAcceleratorsClient", "Get", resp, "Failure sending request")
        return
        }

        result, err = client.GetResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "appplatform.PredefinedAcceleratorsClient", "Get", resp, "Failure responding to request")
        return
        }

    return
}

    // GetPreparer prepares the Get request.
    func (client PredefinedAcceleratorsClient) GetPreparer(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string, predefinedAcceleratorName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "applicationAcceleratorName": autorest.Encode("path",applicationAcceleratorName),
        "predefinedAcceleratorName": autorest.Encode("path",predefinedAcceleratorName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "serviceName": autorest.Encode("path",serviceName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2023-03-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/applicationAccelerators/{applicationAcceleratorName}/predefinedAccelerators/{predefinedAcceleratorName}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client PredefinedAcceleratorsClient) GetSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // GetResponder handles the response to the Get request. The method always
    // closes the http.Response Body.
    func (client PredefinedAcceleratorsClient) GetResponder(resp *http.Response) (result PredefinedAcceleratorResource, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// List handle requests to list all predefined accelerators.
    // Parameters:
        // resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
        // from the Azure Resource Manager API or the portal.
        // serviceName - the name of the Service resource.
        // applicationAcceleratorName - the name of the application accelerator.
func (client PredefinedAcceleratorsClient) List(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string) (result PredefinedAcceleratorResourceCollectionPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/PredefinedAcceleratorsClient.List")
        defer func() {
            sc := -1
        if result.parc.Response.Response != nil {
        sc = result.parc.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        if err := validation.Validate([]validation.Validation{
        { TargetValue: serviceName,
         Constraints: []validation.Constraint{	{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-z][a-z0-9-]*[a-z0-9]$`, Chain: nil }}}}); err != nil {
        return result, validation.NewError("appplatform.PredefinedAcceleratorsClient", "List", err.Error())
        }

            result.fn = client.listNextResults
    req, err := client.ListPreparer(ctx, resourceGroupName, serviceName, applicationAcceleratorName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "appplatform.PredefinedAcceleratorsClient", "List", nil , "Failure preparing request")
    return
    }

        resp, err := client.ListSender(req)
        if err != nil {
        result.parc.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "appplatform.PredefinedAcceleratorsClient", "List", resp, "Failure sending request")
        return
        }

        result.parc, err = client.ListResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "appplatform.PredefinedAcceleratorsClient", "List", resp, "Failure responding to request")
        return
        }
            if result.parc.hasNextLink() && result.parc.IsEmpty() {
            err = result.NextWithContext(ctx)
            return
            }

    return
}

    // ListPreparer prepares the List request.
    func (client PredefinedAcceleratorsClient) ListPreparer(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "applicationAcceleratorName": autorest.Encode("path",applicationAcceleratorName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "serviceName": autorest.Encode("path",serviceName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2023-03-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/applicationAccelerators/{applicationAcceleratorName}/predefinedAccelerators",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client PredefinedAcceleratorsClient) ListSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // ListResponder handles the response to the List request. The method always
    // closes the http.Response Body.
    func (client PredefinedAcceleratorsClient) ListResponder(resp *http.Response) (result PredefinedAcceleratorResourceCollection, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

            // listNextResults retrieves the next set of results, if any.
            func (client PredefinedAcceleratorsClient) listNextResults(ctx context.Context, lastResults PredefinedAcceleratorResourceCollection) (result PredefinedAcceleratorResourceCollection, err error) {
            req, err := lastResults.predefinedAcceleratorResourceCollectionPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "appplatform.PredefinedAcceleratorsClient", "listNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "appplatform.PredefinedAcceleratorsClient", "listNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "appplatform.PredefinedAcceleratorsClient", "listNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

            // ListComplete enumerates all values, automatically crossing page boundaries as required.
            func (client PredefinedAcceleratorsClient) ListComplete(ctx context.Context, resourceGroupName string, serviceName string, applicationAcceleratorName string) (result PredefinedAcceleratorResourceCollectionIterator, err error) {
            if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/PredefinedAcceleratorsClient.List")
            defer func() {
            sc := -1
            if result.Response().Response.Response != nil {
            sc = result.page.Response().Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
            }()
            }
                    result.page, err = client.List(ctx, resourceGroupName, serviceName, applicationAcceleratorName)
                            return
            }

