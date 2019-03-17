// +build go1.9

// Copyright 2018 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package storageimportexport

import original "github.com/Azure/azure-sdk-for-go/services/storageimportexport/mgmt/2016-11-01/storageimportexport"

type BitLockerKeysClient = original.BitLockerKeysClient

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type JobsClient = original.JobsClient
type LocationsClient = original.LocationsClient
type DriveState = original.DriveState

const (
	Completed         DriveState = original.Completed
	CompletedMoreInfo DriveState = original.CompletedMoreInfo
	NeverReceived     DriveState = original.NeverReceived
	Received          DriveState = original.Received
	ShippedBack       DriveState = original.ShippedBack
	Specified         DriveState = original.Specified
	Transferring      DriveState = original.Transferring
)

type DriveBitLockerKey = original.DriveBitLockerKey
type DriveStatus = original.DriveStatus
type ErrorResponse = original.ErrorResponse
type ErrorResponseError = original.ErrorResponseError
type ErrorResponseErrorDetailsItem = original.ErrorResponseErrorDetailsItem
type Export = original.Export
type ExportBlobList = original.ExportBlobList
type GetBitLockerKeysResponse = original.GetBitLockerKeysResponse
type JobDetails = original.JobDetails
type JobResponse = original.JobResponse
type ListJobsResponse = original.ListJobsResponse
type ListJobsResponseIterator = original.ListJobsResponseIterator
type ListJobsResponsePage = original.ListJobsResponsePage
type ListOperationsResponse = original.ListOperationsResponse
type Location = original.Location
type LocationProperties = original.LocationProperties
type LocationsResponse = original.LocationsResponse
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type PackageInfomation = original.PackageInfomation
type PutJobParameters = original.PutJobParameters
type ReturnAddress = original.ReturnAddress
type ReturnShipping = original.ReturnShipping
type ShippingInformation = original.ShippingInformation
type UpdateJobParameters = original.UpdateJobParameters
type UpdateJobParametersProperties = original.UpdateJobParametersProperties
type OperationsClient = original.OperationsClient

func NewBitLockerKeysClient(subscriptionID string, acceptLanguage string) BitLockerKeysClient {
	return original.NewBitLockerKeysClient(subscriptionID, acceptLanguage)
}
func NewBitLockerKeysClientWithBaseURI(baseURI string, subscriptionID string, acceptLanguage string) BitLockerKeysClient {
	return original.NewBitLockerKeysClientWithBaseURI(baseURI, subscriptionID, acceptLanguage)
}
func New(subscriptionID string, acceptLanguage string) BaseClient {
	return original.New(subscriptionID, acceptLanguage)
}
func NewWithBaseURI(baseURI string, subscriptionID string, acceptLanguage string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID, acceptLanguage)
}
func NewJobsClient(subscriptionID string, acceptLanguage string) JobsClient {
	return original.NewJobsClient(subscriptionID, acceptLanguage)
}
func NewJobsClientWithBaseURI(baseURI string, subscriptionID string, acceptLanguage string) JobsClient {
	return original.NewJobsClientWithBaseURI(baseURI, subscriptionID, acceptLanguage)
}
func NewLocationsClient(subscriptionID string, acceptLanguage string) LocationsClient {
	return original.NewLocationsClient(subscriptionID, acceptLanguage)
}
func NewLocationsClientWithBaseURI(baseURI string, subscriptionID string, acceptLanguage string) LocationsClient {
	return original.NewLocationsClientWithBaseURI(baseURI, subscriptionID, acceptLanguage)
}
func PossibleDriveStateValues() []DriveState {
	return original.PossibleDriveStateValues()
}
func NewOperationsClient(subscriptionID string, acceptLanguage string) OperationsClient {
	return original.NewOperationsClient(subscriptionID, acceptLanguage)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string, acceptLanguage string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID, acceptLanguage)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
