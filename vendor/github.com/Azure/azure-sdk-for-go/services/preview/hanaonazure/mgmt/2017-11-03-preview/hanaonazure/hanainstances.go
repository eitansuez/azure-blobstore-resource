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
func (client HanaInstancesClient) listNextResults(lastResults HanaInstancesListResult) (result HanaInstancesListResult, err error) {
	req, err := lastResults.hanaInstancesListResultPreparer()
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
	result.page, err = client.List(ctx)
	return
}

// ListByResourceGroup gets a list of SAP HANA instances in the specified subscription and the resource group. The
// operations returns various properties of each SAP HANA on Azure instance.
// Parameters:
// resourceGroupName - name of the resource group.
func (client HanaInstancesClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result HanaInstancesListResultPage, err error) {
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
func (client HanaInstancesClient) listByResourceGroupNextResults(lastResults HanaInstancesListResult) (result HanaInstancesListResult, err error) {
	req, err := lastResults.hanaInstancesListResultPreparer()
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
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName)
	return
}

// Restart the operation to restart a SAP HANA instance.
// Parameters:
// resourceGroupName - name of the resource group.
// hanaInstanceName - name of the SAP HANA on Azure instance.
func (client HanaInstancesClient) Restart(ctx context.Context, resourceGroupName string, hanaInstanceName string) (result autorest.Response, err error) {
	req, err := client.RestartPreparer(ctx, resourceGroupName, hanaInstanceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hanaonazure.HanaInstancesClient", "Restart", nil, "Failure preparing request")
		return
	}

	resp, err := client.RestartSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "hanaonazure.HanaInstancesClient", "Restart", resp, "Failure sending request")
		return
	}

	result, err = client.RestartResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hanaonazure.HanaInstancesClient", "Restart", resp, "Failure responding to request")
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
func (client HanaInstancesClient) RestartSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
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
