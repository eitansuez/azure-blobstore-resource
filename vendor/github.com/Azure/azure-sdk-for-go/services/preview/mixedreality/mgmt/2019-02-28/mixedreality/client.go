// Package mixedreality implements the Azure ARM Mixedreality service API version 2019-02-28-preview.
//
// Mixed Reality Client
package mixedreality

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
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

const (
	// DefaultBaseURI is the default URI used for the service Mixedreality
	DefaultBaseURI = "https://management.azure.com"
)

// BaseClient is the base client for Mixedreality.
type BaseClient struct {
	autorest.Client
	BaseURI        string
	SubscriptionID string
}

// New creates an instance of the BaseClient client.
func New(subscriptionID string) BaseClient {
	return NewWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWithBaseURI creates an instance of the BaseClient client.
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return BaseClient{
		Client:         autorest.NewClientWithUserAgent(UserAgent()),
		BaseURI:        baseURI,
		SubscriptionID: subscriptionID,
	}
}

// CheckNameAvailabilityLocal check Name Availability for global uniqueness
// Parameters:
// location - the location in which uniqueness will be verified.
// checkNameAvailability - check Name Availability Request.
func (client BaseClient) CheckNameAvailabilityLocal(ctx context.Context, location string, checkNameAvailability CheckNameAvailabilityRequest) (result CheckNameAvailabilityResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.CheckNameAvailabilityLocal")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: location,
			Constraints: []validation.Constraint{{Target: "location", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "location", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "location", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: checkNameAvailability,
			Constraints: []validation.Constraint{{Target: "checkNameAvailability.Name", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "checkNameAvailability.Type", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("mixedreality.BaseClient", "CheckNameAvailabilityLocal", err.Error())
	}

	req, err := client.CheckNameAvailabilityLocalPreparer(ctx, location, checkNameAvailability)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mixedreality.BaseClient", "CheckNameAvailabilityLocal", nil, "Failure preparing request")
		return
	}

	resp, err := client.CheckNameAvailabilityLocalSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "mixedreality.BaseClient", "CheckNameAvailabilityLocal", resp, "Failure sending request")
		return
	}

	result, err = client.CheckNameAvailabilityLocalResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mixedreality.BaseClient", "CheckNameAvailabilityLocal", resp, "Failure responding to request")
	}

	return
}

// CheckNameAvailabilityLocalPreparer prepares the CheckNameAvailabilityLocal request.
func (client BaseClient) CheckNameAvailabilityLocalPreparer(ctx context.Context, location string, checkNameAvailability CheckNameAvailabilityRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-02-28-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.MixedReality/locations/{location}/checkNameAvailability", pathParameters),
		autorest.WithJSON(checkNameAvailability),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CheckNameAvailabilityLocalSender sends the CheckNameAvailabilityLocal request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) CheckNameAvailabilityLocalSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// CheckNameAvailabilityLocalResponder handles the response to the CheckNameAvailabilityLocal request. The method always
// closes the http.Response Body.
func (client BaseClient) CheckNameAvailabilityLocalResponder(resp *http.Response) (result CheckNameAvailabilityResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
