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

// PipelineClient is the client for the Pipeline methods of the Synapse service.
type PipelineClient struct {
    BaseClient
}
// NewPipelineClient creates an instance of the PipelineClient client.
func NewPipelineClient(endpoint string) PipelineClient {
    return PipelineClient{ New(endpoint)}
}

// CreateOrUpdatePipeline creates or updates a pipeline.
    // Parameters:
        // pipelineName - the pipeline name.
        // pipeline - pipeline resource definition.
        // ifMatch - eTag of the pipeline entity.  Should only be specified for update, for which it should match
        // existing entity or can be * for unconditional update.
func (client PipelineClient) CreateOrUpdatePipeline(ctx context.Context, pipelineName string, pipeline PipelineResource, ifMatch string) (result PipelineCreateOrUpdatePipelineFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/PipelineClient.CreateOrUpdatePipeline")
        defer func() {
            sc := -1
        if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
        sc = result.FutureAPI.Response().StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        if err := validation.Validate([]validation.Validation{
        { TargetValue: pipelineName,
         Constraints: []validation.Constraint{	{Target: "pipelineName", Name: validation.MaxLength, Rule: 260, Chain: nil },
        	{Target: "pipelineName", Name: validation.MinLength, Rule: 1, Chain: nil },
        	{Target: "pipelineName", Name: validation.Pattern, Rule: `^[A-Za-z0-9_][^<>*#.%&:\\+?/]*$`, Chain: nil }}},
        { TargetValue: pipeline,
         Constraints: []validation.Constraint{	{Target: "pipeline.Pipeline", Name: validation.Null, Rule: true ,
        Chain: []validation.Constraint{	{Target: "pipeline.Pipeline.Concurrency", Name: validation.Null, Rule: false ,
        Chain: []validation.Constraint{	{Target: "pipeline.Pipeline.Concurrency", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil },
        }},
        }}}}}); err != nil {
        return result, validation.NewError("synapse.PipelineClient", "CreateOrUpdatePipeline", err.Error())
        }

        req, err := client.CreateOrUpdatePipelinePreparer(ctx, pipelineName, pipeline, ifMatch)
    if err != nil {
    err = autorest.NewErrorWithError(err, "synapse.PipelineClient", "CreateOrUpdatePipeline", nil , "Failure preparing request")
    return
    }

        result, err = client.CreateOrUpdatePipelineSender(req)
        if err != nil {
        err = autorest.NewErrorWithError(err, "synapse.PipelineClient", "CreateOrUpdatePipeline", result.Response(), "Failure sending request")
        return
        }

    return
}

    // CreateOrUpdatePipelinePreparer prepares the CreateOrUpdatePipeline request.
    func (client PipelineClient) CreateOrUpdatePipelinePreparer(ctx context.Context, pipelineName string, pipeline PipelineResource, ifMatch string) (*http.Request, error) {
        urlParameters := map[string]interface{} {
        "endpoint": client.Endpoint,
        }

        pathParameters := map[string]interface{} {
        "pipelineName": autorest.Encode("path",pipelineName),
        }

            const APIVersion = "2021-06-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsContentType("application/json; charset=utf-8"),
autorest.AsPut(),
autorest.WithCustomBaseURL("{endpoint}", urlParameters),
autorest.WithPathParameters("/pipelines/{pipelineName}",pathParameters),
autorest.WithJSON(pipeline),
autorest.WithQueryParameters(queryParameters))
        if len(ifMatch) > 0 {
        preparer = autorest.DecoratePreparer(preparer,
        autorest.WithHeader("If-Match",autorest.String(ifMatch)))
        }
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateOrUpdatePipelineSender sends the CreateOrUpdatePipeline request. The method will close the
    // http.Response Body if it receives an error.
    func (client PipelineClient) CreateOrUpdatePipelineSender(req *http.Request) (future PipelineCreateOrUpdatePipelineFuture, err error) {
            var resp *http.Response
            future.FutureAPI = &azure.Future{}
            resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            if err != nil {
            return
            }
            var azf azure.Future
            azf, err = azure.NewFutureFromResponse(resp)
            future.FutureAPI = &azf
            future.Result = future.result
            return
                }

    // CreateOrUpdatePipelineResponder handles the response to the CreateOrUpdatePipeline request. The method always
    // closes the http.Response Body.
    func (client PipelineClient) CreateOrUpdatePipelineResponder(resp *http.Response) (result PipelineResource, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// CreatePipelineRun creates a run of a pipeline.
    // Parameters:
        // pipelineName - the pipeline name.
        // referencePipelineRunID - the pipeline run identifier. If run ID is specified the parameters of the specified
        // run will be used to create a new run.
        // isRecovery - recovery mode flag. If recovery mode is set to true, the specified referenced pipeline run and
        // the new run will be grouped under the same groupId.
        // startActivityName - in recovery mode, the rerun will start from this activity. If not specified, all
        // activities will run.
        // parameters - parameters of the pipeline run. These parameters will be used only if the runId is not
        // specified.
func (client PipelineClient) CreatePipelineRun(ctx context.Context, pipelineName string, referencePipelineRunID string, isRecovery *bool, startActivityName string, parameters map[string]interface{}) (result CreateRunResponse, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/PipelineClient.CreatePipelineRun")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        if err := validation.Validate([]validation.Validation{
        { TargetValue: pipelineName,
         Constraints: []validation.Constraint{	{Target: "pipelineName", Name: validation.MaxLength, Rule: 260, Chain: nil },
        	{Target: "pipelineName", Name: validation.MinLength, Rule: 1, Chain: nil },
        	{Target: "pipelineName", Name: validation.Pattern, Rule: `^[A-Za-z0-9_][^<>*#.%&:\\+?/]*$`, Chain: nil }}}}); err != nil {
        return result, validation.NewError("synapse.PipelineClient", "CreatePipelineRun", err.Error())
        }

        req, err := client.CreatePipelineRunPreparer(ctx, pipelineName, referencePipelineRunID, isRecovery, startActivityName, parameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "synapse.PipelineClient", "CreatePipelineRun", nil , "Failure preparing request")
    return
    }

        resp, err := client.CreatePipelineRunSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "synapse.PipelineClient", "CreatePipelineRun", resp, "Failure sending request")
        return
        }

        result, err = client.CreatePipelineRunResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "synapse.PipelineClient", "CreatePipelineRun", resp, "Failure responding to request")
        return
        }

    return
}

    // CreatePipelineRunPreparer prepares the CreatePipelineRun request.
    func (client PipelineClient) CreatePipelineRunPreparer(ctx context.Context, pipelineName string, referencePipelineRunID string, isRecovery *bool, startActivityName string, parameters map[string]interface{}) (*http.Request, error) {
        urlParameters := map[string]interface{} {
        "endpoint": client.Endpoint,
        }

        pathParameters := map[string]interface{} {
        "pipelineName": autorest.Encode("path",pipelineName),
        }

            const APIVersion = "2021-06-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }
        if len(referencePipelineRunID) > 0 {
        queryParameters["referencePipelineRunId"] = autorest.Encode("query",referencePipelineRunID)
        }
        if isRecovery != nil {
        queryParameters["isRecovery"] = autorest.Encode("query",*isRecovery)
        }
        if len(startActivityName) > 0 {
        queryParameters["startActivityName"] = autorest.Encode("query",startActivityName)
        }

    preparer := autorest.CreatePreparer(
autorest.AsContentType("application/json; charset=utf-8"),
autorest.AsPost(),
autorest.WithCustomBaseURL("{endpoint}", urlParameters),
autorest.WithPathParameters("/pipelines/{pipelineName}/createRun",pathParameters),
autorest.WithQueryParameters(queryParameters))
        if parameters != nil && len(parameters) > 0 {
        preparer = autorest.DecoratePreparer(preparer,
        autorest.WithJSON(parameters))
        }
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreatePipelineRunSender sends the CreatePipelineRun request. The method will close the
    // http.Response Body if it receives an error.
    func (client PipelineClient) CreatePipelineRunSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
                }

    // CreatePipelineRunResponder handles the response to the CreatePipelineRun request. The method always
    // closes the http.Response Body.
    func (client PipelineClient) CreatePipelineRunResponder(resp *http.Response) (result CreateRunResponse, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// DeletePipeline deletes a pipeline.
    // Parameters:
        // pipelineName - the pipeline name.
func (client PipelineClient) DeletePipeline(ctx context.Context, pipelineName string) (result PipelineDeletePipelineFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/PipelineClient.DeletePipeline")
        defer func() {
            sc := -1
        if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
        sc = result.FutureAPI.Response().StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        if err := validation.Validate([]validation.Validation{
        { TargetValue: pipelineName,
         Constraints: []validation.Constraint{	{Target: "pipelineName", Name: validation.MaxLength, Rule: 260, Chain: nil },
        	{Target: "pipelineName", Name: validation.MinLength, Rule: 1, Chain: nil },
        	{Target: "pipelineName", Name: validation.Pattern, Rule: `^[A-Za-z0-9_][^<>*#.%&:\\+?/]*$`, Chain: nil }}}}); err != nil {
        return result, validation.NewError("synapse.PipelineClient", "DeletePipeline", err.Error())
        }

        req, err := client.DeletePipelinePreparer(ctx, pipelineName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "synapse.PipelineClient", "DeletePipeline", nil , "Failure preparing request")
    return
    }

        result, err = client.DeletePipelineSender(req)
        if err != nil {
        err = autorest.NewErrorWithError(err, "synapse.PipelineClient", "DeletePipeline", result.Response(), "Failure sending request")
        return
        }

    return
}

    // DeletePipelinePreparer prepares the DeletePipeline request.
    func (client PipelineClient) DeletePipelinePreparer(ctx context.Context, pipelineName string) (*http.Request, error) {
        urlParameters := map[string]interface{} {
        "endpoint": client.Endpoint,
        }

        pathParameters := map[string]interface{} {
        "pipelineName": autorest.Encode("path",pipelineName),
        }

            const APIVersion = "2021-06-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsDelete(),
autorest.WithCustomBaseURL("{endpoint}", urlParameters),
autorest.WithPathParameters("/pipelines/{pipelineName}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeletePipelineSender sends the DeletePipeline request. The method will close the
    // http.Response Body if it receives an error.
    func (client PipelineClient) DeletePipelineSender(req *http.Request) (future PipelineDeletePipelineFuture, err error) {
            var resp *http.Response
            future.FutureAPI = &azure.Future{}
            resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            if err != nil {
            return
            }
            var azf azure.Future
            azf, err = azure.NewFutureFromResponse(resp)
            future.FutureAPI = &azf
            future.Result = future.result
            return
                }

    // DeletePipelineResponder handles the response to the DeletePipeline request. The method always
    // closes the http.Response Body.
    func (client PipelineClient) DeletePipelineResponder(resp *http.Response) (result autorest.Response, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted,http.StatusNoContent),
            autorest.ByClosing())
            result.Response = resp
            return
    }

// GetPipeline gets a pipeline.
    // Parameters:
        // pipelineName - the pipeline name.
        // ifNoneMatch - eTag of the pipeline entity. Should only be specified for get. If the ETag matches the
        // existing entity tag, or if * was provided, then no content will be returned.
func (client PipelineClient) GetPipeline(ctx context.Context, pipelineName string, ifNoneMatch string) (result PipelineResource, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/PipelineClient.GetPipeline")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        if err := validation.Validate([]validation.Validation{
        { TargetValue: pipelineName,
         Constraints: []validation.Constraint{	{Target: "pipelineName", Name: validation.MaxLength, Rule: 260, Chain: nil },
        	{Target: "pipelineName", Name: validation.MinLength, Rule: 1, Chain: nil },
        	{Target: "pipelineName", Name: validation.Pattern, Rule: `^[A-Za-z0-9_][^<>*#.%&:\\+?/]*$`, Chain: nil }}}}); err != nil {
        return result, validation.NewError("synapse.PipelineClient", "GetPipeline", err.Error())
        }

        req, err := client.GetPipelinePreparer(ctx, pipelineName, ifNoneMatch)
    if err != nil {
    err = autorest.NewErrorWithError(err, "synapse.PipelineClient", "GetPipeline", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetPipelineSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "synapse.PipelineClient", "GetPipeline", resp, "Failure sending request")
        return
        }

        result, err = client.GetPipelineResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "synapse.PipelineClient", "GetPipeline", resp, "Failure responding to request")
        return
        }

    return
}

    // GetPipelinePreparer prepares the GetPipeline request.
    func (client PipelineClient) GetPipelinePreparer(ctx context.Context, pipelineName string, ifNoneMatch string) (*http.Request, error) {
        urlParameters := map[string]interface{} {
        "endpoint": client.Endpoint,
        }

        pathParameters := map[string]interface{} {
        "pipelineName": autorest.Encode("path",pipelineName),
        }

            const APIVersion = "2021-06-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithCustomBaseURL("{endpoint}", urlParameters),
autorest.WithPathParameters("/pipelines/{pipelineName}",pathParameters),
autorest.WithQueryParameters(queryParameters))
        if len(ifNoneMatch) > 0 {
        preparer = autorest.DecoratePreparer(preparer,
        autorest.WithHeader("If-None-Match",autorest.String(ifNoneMatch)))
        }
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetPipelineSender sends the GetPipeline request. The method will close the
    // http.Response Body if it receives an error.
    func (client PipelineClient) GetPipelineSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
                }

    // GetPipelineResponder handles the response to the GetPipeline request. The method always
    // closes the http.Response Body.
    func (client PipelineClient) GetPipelineResponder(resp *http.Response) (result PipelineResource, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusNotModified),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// GetPipelinesByWorkspace lists pipelines.
func (client PipelineClient) GetPipelinesByWorkspace(ctx context.Context) (result PipelineListResponsePage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/PipelineClient.GetPipelinesByWorkspace")
        defer func() {
            sc := -1
        if result.plr.Response.Response != nil {
        sc = result.plr.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        result.fn = client.getPipelinesByWorkspaceNextResults
    req, err := client.GetPipelinesByWorkspacePreparer(ctx)
    if err != nil {
    err = autorest.NewErrorWithError(err, "synapse.PipelineClient", "GetPipelinesByWorkspace", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetPipelinesByWorkspaceSender(req)
        if err != nil {
        result.plr.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "synapse.PipelineClient", "GetPipelinesByWorkspace", resp, "Failure sending request")
        return
        }

        result.plr, err = client.GetPipelinesByWorkspaceResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "synapse.PipelineClient", "GetPipelinesByWorkspace", resp, "Failure responding to request")
        return
        }
            if result.plr.hasNextLink() && result.plr.IsEmpty() {
            err = result.NextWithContext(ctx)
            return
            }

    return
}

    // GetPipelinesByWorkspacePreparer prepares the GetPipelinesByWorkspace request.
    func (client PipelineClient) GetPipelinesByWorkspacePreparer(ctx context.Context) (*http.Request, error) {
        urlParameters := map[string]interface{} {
        "endpoint": client.Endpoint,
        }

            const APIVersion = "2021-06-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithCustomBaseURL("{endpoint}", urlParameters),
autorest.WithPath("/pipelines"),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetPipelinesByWorkspaceSender sends the GetPipelinesByWorkspace request. The method will close the
    // http.Response Body if it receives an error.
    func (client PipelineClient) GetPipelinesByWorkspaceSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
                }

    // GetPipelinesByWorkspaceResponder handles the response to the GetPipelinesByWorkspace request. The method always
    // closes the http.Response Body.
    func (client PipelineClient) GetPipelinesByWorkspaceResponder(resp *http.Response) (result PipelineListResponse, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

            // getPipelinesByWorkspaceNextResults retrieves the next set of results, if any.
            func (client PipelineClient) getPipelinesByWorkspaceNextResults(ctx context.Context, lastResults PipelineListResponse) (result PipelineListResponse, err error) {
            req, err := lastResults.pipelineListResponsePreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "synapse.PipelineClient", "getPipelinesByWorkspaceNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.GetPipelinesByWorkspaceSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "synapse.PipelineClient", "getPipelinesByWorkspaceNextResults", resp, "Failure sending next results request")
            }
            result, err = client.GetPipelinesByWorkspaceResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "synapse.PipelineClient", "getPipelinesByWorkspaceNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

            // GetPipelinesByWorkspaceComplete enumerates all values, automatically crossing page boundaries as required.
            func (client PipelineClient) GetPipelinesByWorkspaceComplete(ctx context.Context) (result PipelineListResponseIterator, err error) {
            if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/PipelineClient.GetPipelinesByWorkspace")
            defer func() {
            sc := -1
            if result.Response().Response.Response != nil {
            sc = result.page.Response().Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
            }()
            }
                    result.page, err = client.GetPipelinesByWorkspace(ctx)
                            return
            }

// RenamePipeline renames a pipeline.
    // Parameters:
        // pipelineName - the pipeline name.
        // request - proposed new name.
func (client PipelineClient) RenamePipeline(ctx context.Context, pipelineName string, request ArtifactRenameRequest) (result PipelineRenamePipelineFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/PipelineClient.RenamePipeline")
        defer func() {
            sc := -1
        if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
        sc = result.FutureAPI.Response().StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        if err := validation.Validate([]validation.Validation{
        { TargetValue: pipelineName,
         Constraints: []validation.Constraint{	{Target: "pipelineName", Name: validation.MaxLength, Rule: 260, Chain: nil },
        	{Target: "pipelineName", Name: validation.MinLength, Rule: 1, Chain: nil },
        	{Target: "pipelineName", Name: validation.Pattern, Rule: `^[A-Za-z0-9_][^<>*#.%&:\\+?/]*$`, Chain: nil }}},
        { TargetValue: request,
         Constraints: []validation.Constraint{	{Target: "request.NewName", Name: validation.Null, Rule: false ,
        Chain: []validation.Constraint{	{Target: "request.NewName", Name: validation.MaxLength, Rule: 260, Chain: nil },
        	{Target: "request.NewName", Name: validation.MinLength, Rule: 1, Chain: nil },
        	{Target: "request.NewName", Name: validation.Pattern, Rule: `^[A-Za-z0-9_][^<>*#.%&:\\+?/]*$`, Chain: nil },
        }}}}}); err != nil {
        return result, validation.NewError("synapse.PipelineClient", "RenamePipeline", err.Error())
        }

        req, err := client.RenamePipelinePreparer(ctx, pipelineName, request)
    if err != nil {
    err = autorest.NewErrorWithError(err, "synapse.PipelineClient", "RenamePipeline", nil , "Failure preparing request")
    return
    }

        result, err = client.RenamePipelineSender(req)
        if err != nil {
        err = autorest.NewErrorWithError(err, "synapse.PipelineClient", "RenamePipeline", result.Response(), "Failure sending request")
        return
        }

    return
}

    // RenamePipelinePreparer prepares the RenamePipeline request.
    func (client PipelineClient) RenamePipelinePreparer(ctx context.Context, pipelineName string, request ArtifactRenameRequest) (*http.Request, error) {
        urlParameters := map[string]interface{} {
        "endpoint": client.Endpoint,
        }

        pathParameters := map[string]interface{} {
        "pipelineName": autorest.Encode("path",pipelineName),
        }

            const APIVersion = "2021-06-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsContentType("application/json; charset=utf-8"),
autorest.AsPost(),
autorest.WithCustomBaseURL("{endpoint}", urlParameters),
autorest.WithPathParameters("/pipelines/{pipelineName}/rename",pathParameters),
autorest.WithJSON(request),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // RenamePipelineSender sends the RenamePipeline request. The method will close the
    // http.Response Body if it receives an error.
    func (client PipelineClient) RenamePipelineSender(req *http.Request) (future PipelineRenamePipelineFuture, err error) {
            var resp *http.Response
            future.FutureAPI = &azure.Future{}
            resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            if err != nil {
            return
            }
            var azf azure.Future
            azf, err = azure.NewFutureFromResponse(resp)
            future.FutureAPI = &azf
            future.Result = future.result
            return
                }

    // RenamePipelineResponder handles the response to the RenamePipeline request. The method always
    // closes the http.Response Body.
    func (client PipelineClient) RenamePipelineResponder(resp *http.Response) (result autorest.Response, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted),
            autorest.ByClosing())
            result.Response = resp
            return
    }

