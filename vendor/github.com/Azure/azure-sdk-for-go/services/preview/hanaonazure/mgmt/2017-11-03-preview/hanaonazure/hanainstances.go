package hanaonazure

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
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// HanaInstancesClient is the HANA on Azure Client
type HanaInstancesClient struct {
	BaseClient
}

// NewHanaInstancesClient creates an instance of the HanaInstancesClient client.
func NewHanaInstancesClient(subscriptionID string) HanaInstancesClient {
	return NewHanaInstancesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewHanaInstancesClientWithBaseURI creates an instance of the HanaInstancesClient client.
func NewHanaInstancesClientWithBaseURI(baseURI string, subscriptionID string) HanaInstancesClient {
	return HanaInstancesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets properties of a SAP HANA instance for the specified subscription, resource group, and instance name.
// Parameters:
// resourceGroupName - name of the resource group.
// hanaInstanceName - name of the SAP HANA on Azure instance.
func (client HanaInstancesClient) Get(ctx context.Context, resourceGroupName string, hanaInstanceName string) (result HanaInstance, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/HanaInstancesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, hanaInstanceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hanaonazure.HanaInstancesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "hanaonazure.HanaInstancesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hanaonazure.HanaInstancesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client HanaInstancesClient) GetPreparer(ctx context.Context, resourceGroupName string, hanaInstanceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hanaInstanceName":  autorest.Encode("path", hanaInstanceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-11-03-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HanaOnAzure/hanaInstances/{hanaInstanceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client HanaInstancesClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client HanaInstancesClient) GetResponder(resp *http.Response) (result HanaInstance, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets a list of SAP HANA instances in the specified subscription. The operations returns various properties of
// each SAP HANA on Azure instance.
func (client HanaInstancesClient) List(ctx context.Context) (result HanaInstancesListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/HanaInstancesClient.List")
		defer func() {
			sc := -1
			if result.hilr.Response.Response != nil {
				sc = result.hilr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hanaonazure.HanaInstancesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.hilr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "hanaonazure.HanaInstancesClient", "List", resp, "Failure sending request")
		return
	}

	result.hilr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hanaonazure.HanaInstancesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client HanaInstancesClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-11-03-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.HanaOnAzure/hanaInstances", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client HanaInstancesClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client HanaInstancesClient) ListResponder(resp *http.Response) (result HanaInstancesListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client HanaInstancesClient) listNextResults(ctx context.Context, lastResults HanaInstancesListResult) (result HanaInstancesListResult, err error) {
	req, err := lastResults.hanaInstancesListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "hanaonazure.HanaInstancesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "hanaonazure.HanaInstancesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hanaonazure.HanaInstancesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client HanaInstancesClient) ListComplete(ctx context.Context) (result HanaInstancesListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/HanaInstancesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx)
	return
}

// ListByResourceGroup gets a list of SAP HANA instances in the specified subscription and the resource group. The
// operations returns various properties of each SAP HANA on Azure instance.
// Parameters:
// resourceGroupName - name of the resource group.
func (client HanaInstancesClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result HanaInstancesListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/HanaInstancesClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.hilr.Response.Response != nil {
				sc = result.hilr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hanaonazure.HanaInstancesClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.hilr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "hanaonazure.HanaInstancesClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.hilr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hanaonazure.HanaInstancesClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client HanaInstancesClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-11-03-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HanaOnAzure/hanaInstances", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client HanaInstancesClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client HanaInstancesClient) ListByResourceGroupResponder(resp *http.Response) (result HanaInstancesListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client HanaInstancesClient) listByResourceGroupNextResults(ctx context.Context, lastResults HanaInstancesListResult) (result HanaInstancesListResult, err error) {
	req, err := lastResults.hanaInstancesListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "hanaonazure.HanaInstancesClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "hanaonazure.HanaInstancesClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hanaonazure.HanaInstancesClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client HanaInstancesClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result HanaInstancesListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/HanaInstancesClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName)
	return
}

// Restart the operation to restart a SAP HANA instance.
// Parameters:
// resourceGroupName - name of the resource group.
// hanaInstanceName - name of the SAP HANA on Azure instance.
func (client HanaInstancesClient) Restart(ctx context.Context, resourceGroupName string, hanaInstanceName string) (result HanaInstancesRestartFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/HanaInstancesClient.Restart")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.RestartPreparer(ctx, resourceGroupName, hanaInstanceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hanaonazure.HanaInstancesClient", "Restart", nil, "Failure preparing request")
		return
	}

	result, err = client.RestartSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hanaonazure.HanaInstancesClient", "Restart", result.Response(), "Failure sending request")
		return
	}

	return
}

// RestartPreparer prepares the Restart request.
func (client HanaInstancesClient) RestartPreparer(ctx context.Context, resourceGroupName string, hanaInstanceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hanaInstanceName":  autorest.Encode("path", hanaInstanceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-11-03-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HanaOnAzure/hanaInstances/{hanaInstanceName}/restart", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RestartSender sends the Restart request. The method will close the
// http.Response Body if it receives an error.
func (client HanaInstancesClient) RestartSender(req *http.Request) (future HanaInstancesRestartFuture, err error) {
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// RestartResponder handles the response to the Restart request. The method always
// closes the http.Response Body.
func (client HanaInstancesClient) RestartResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Update patches the Tags field of a SAP HANA instance for the specified subscription, resource group, and instance
// name.
// Parameters:
// resourceGroupName - name of the resource group.
// hanaInstanceName - name of the SAP HANA on Azure instance.
// tagsParameter - request body that only contains the new Tags field
func (client HanaInstancesClient) Update(ctx context.Context, resourceGroupName string, hanaInstanceName string, tagsParameter Tags) (result HanaInstance, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/HanaInstancesClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, hanaInstanceName, tagsParameter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hanaonazure.HanaInstancesClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "hanaonazure.HanaInstancesClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hanaonazure.HanaInstancesClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client HanaInstancesClient) UpdatePreparer(ctx context.Context, resourceGroupName string, hanaInstanceName string, tagsParameter Tags) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hanaInstanceName":  autorest.Encode("path", hanaInstanceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-11-03-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HanaOnAzure/hanaInstances/{hanaInstanceName}", pathParameters),
		autorest.WithJSON(tagsParameter),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client HanaInstancesClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client HanaInstancesClient) UpdateResponder(resp *http.Response) (result HanaInstance, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
