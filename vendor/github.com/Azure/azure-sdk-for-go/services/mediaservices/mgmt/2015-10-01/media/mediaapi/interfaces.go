package mediaapi

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
	"github.com/Azure/azure-sdk-for-go/services/mediaservices/mgmt/2015-10-01/media"
	"github.com/Azure/go-autorest/autorest"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result media.OperationListResult, err error)
}

var _ OperationsClientAPI = (*media.OperationsClient)(nil)

// ServiceClientAPI contains the set of methods on the ServiceClient type.
type ServiceClientAPI interface {
	CheckNameAvailability(ctx context.Context, parameters media.CheckNameAvailabilityInput) (result media.CheckNameAvailabilityOutput, err error)
	Create(ctx context.Context, resourceGroupName string, mediaServiceName string, parameters media.Service) (result media.Service, err error)
	Delete(ctx context.Context, resourceGroupName string, mediaServiceName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, mediaServiceName string) (result media.Service, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result media.ServiceCollection, err error)
	ListKeys(ctx context.Context, resourceGroupName string, mediaServiceName string) (result media.ServiceKeys, err error)
	RegenerateKey(ctx context.Context, resourceGroupName string, mediaServiceName string, parameters media.RegenerateKeyInput) (result media.RegenerateKeyOutput, err error)
	SyncStorageKeys(ctx context.Context, resourceGroupName string, mediaServiceName string, parameters media.SyncStorageKeysInput) (result autorest.Response, err error)
	Update(ctx context.Context, resourceGroupName string, mediaServiceName string, parameters media.Service) (result media.Service, err error)
}

var _ ServiceClientAPI = (*media.ServiceClient)(nil)
