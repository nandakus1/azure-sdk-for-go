package windowsesu

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
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/preview/windowsesu/2019-09-16-preview/windowsesu"

// OsType enumerates the values for os type.
type OsType string

const (
	// Windows7 ...
	Windows7 OsType = "Windows7"
	// WindowsServer2008 ...
	WindowsServer2008 OsType = "WindowsServer2008"
	// WindowsServer2008R2 ...
	WindowsServer2008R2 OsType = "WindowsServer2008R2"
)

// PossibleOsTypeValues returns an array of possible values for the OsType const type.
func PossibleOsTypeValues() []OsType {
	return []OsType{Windows7, WindowsServer2008, WindowsServer2008R2}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Accepted ...
	Accepted ProvisioningState = "Accepted"
	// Canceled ...
	Canceled ProvisioningState = "Canceled"
	// Failed ...
	Failed ProvisioningState = "Failed"
	// Provisioning ...
	Provisioning ProvisioningState = "Provisioning"
	// Succeeded ...
	Succeeded ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{Accepted, Canceled, Failed, Provisioning, Succeeded}
}

// SupportType enumerates the values for support type.
type SupportType string

const (
	// PremiumAssurance ...
	PremiumAssurance SupportType = "PremiumAssurance"
	// SupplementalServicing ...
	SupplementalServicing SupportType = "SupplementalServicing"
)

// PossibleSupportTypeValues returns an array of possible values for the SupportType const type.
func PossibleSupportTypeValues() []SupportType {
	return []SupportType{PremiumAssurance, SupplementalServicing}
}

// AzureEntityResource the resource model definition for a Azure Resource Manager resource with an etag.
type AzureEntityResource struct {
	// Etag - READ-ONLY; Resource Etag.
	Etag *string `json:"etag,omitempty"`
	// ID - READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty"`
}

// ErrorDefinition error definition.
type ErrorDefinition struct {
	// Code - READ-ONLY; Service specific error code which serves as the substatus for the HTTP error code.
	Code *string `json:"code,omitempty"`
	// Message - READ-ONLY; Description of the error.
	Message *string `json:"message,omitempty"`
	// Details - READ-ONLY; Internal error details.
	Details *[]ErrorDefinition `json:"details,omitempty"`
}

// ErrorResponse error response.
type ErrorResponse struct {
	// Error - The error details.
	Error *ErrorDefinition `json:"error,omitempty"`
}

// MultipleActivationKey MAK key details.
type MultipleActivationKey struct {
	autorest.Response `json:"-"`
	// MultipleActivationKeyProperties - MAK key specific properties.
	*MultipleActivationKeyProperties `json:"properties,omitempty"`
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// Location - The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	// ID - READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for MultipleActivationKey.
func (mak MultipleActivationKey) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if mak.MultipleActivationKeyProperties != nil {
		objectMap["properties"] = mak.MultipleActivationKeyProperties
	}
	if mak.Tags != nil {
		objectMap["tags"] = mak.Tags
	}
	if mak.Location != nil {
		objectMap["location"] = mak.Location
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for MultipleActivationKey struct.
func (mak *MultipleActivationKey) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var multipleActivationKeyProperties MultipleActivationKeyProperties
				err = json.Unmarshal(*v, &multipleActivationKeyProperties)
				if err != nil {
					return err
				}
				mak.MultipleActivationKeyProperties = &multipleActivationKeyProperties
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				mak.Tags = tags
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				mak.Location = &location
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				mak.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				mak.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				mak.Type = &typeVar
			}
		}
	}

	return nil
}

// MultipleActivationKeyList list of MAK keys.
type MultipleActivationKeyList struct {
	autorest.Response `json:"-"`
	// Value - List of MAK keys.
	Value *[]MultipleActivationKey `json:"value,omitempty"`
	// NextLink - READ-ONLY; Link to the next page of resources.
	NextLink *string `json:"nextLink,omitempty"`
}

// MultipleActivationKeyListIterator provides access to a complete listing of MultipleActivationKey values.
type MultipleActivationKeyListIterator struct {
	i    int
	page MultipleActivationKeyListPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *MultipleActivationKeyListIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MultipleActivationKeyListIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *MultipleActivationKeyListIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter MultipleActivationKeyListIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter MultipleActivationKeyListIterator) Response() MultipleActivationKeyList {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter MultipleActivationKeyListIterator) Value() MultipleActivationKey {
	if !iter.page.NotDone() {
		return MultipleActivationKey{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the MultipleActivationKeyListIterator type.
func NewMultipleActivationKeyListIterator(page MultipleActivationKeyListPage) MultipleActivationKeyListIterator {
	return MultipleActivationKeyListIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (makl MultipleActivationKeyList) IsEmpty() bool {
	return makl.Value == nil || len(*makl.Value) == 0
}

// multipleActivationKeyListPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (makl MultipleActivationKeyList) multipleActivationKeyListPreparer(ctx context.Context) (*http.Request, error) {
	if makl.NextLink == nil || len(to.String(makl.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(makl.NextLink)))
}

// MultipleActivationKeyListPage contains a page of MultipleActivationKey values.
type MultipleActivationKeyListPage struct {
	fn   func(context.Context, MultipleActivationKeyList) (MultipleActivationKeyList, error)
	makl MultipleActivationKeyList
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *MultipleActivationKeyListPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MultipleActivationKeyListPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.makl)
	if err != nil {
		return err
	}
	page.makl = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *MultipleActivationKeyListPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page MultipleActivationKeyListPage) NotDone() bool {
	return !page.makl.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page MultipleActivationKeyListPage) Response() MultipleActivationKeyList {
	return page.makl
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page MultipleActivationKeyListPage) Values() []MultipleActivationKey {
	if page.makl.IsEmpty() {
		return nil
	}
	return *page.makl.Value
}

// Creates a new instance of the MultipleActivationKeyListPage type.
func NewMultipleActivationKeyListPage(getNextPage func(context.Context, MultipleActivationKeyList) (MultipleActivationKeyList, error)) MultipleActivationKeyListPage {
	return MultipleActivationKeyListPage{fn: getNextPage}
}

// MultipleActivationKeyProperties MAK key specific properties.
type MultipleActivationKeyProperties struct {
	// MultipleActivationKey - READ-ONLY; MAK 5x5 key.
	MultipleActivationKey *string `json:"multipleActivationKey,omitempty"`
	// ExpirationDate - READ-ONLY; End of support of security updates activated by the MAK key.
	ExpirationDate *date.Time `json:"expirationDate,omitempty"`
	// OsType - Type of OS for which the key is requested. Possible values include: 'Windows7', 'WindowsServer2008', 'WindowsServer2008R2'
	OsType OsType `json:"osType,omitempty"`
	// SupportType - Type of support. Possible values include: 'SupplementalServicing', 'PremiumAssurance'
	SupportType SupportType `json:"supportType,omitempty"`
	// InstalledServerNumber - Number of activations/servers using the MAK key.
	InstalledServerNumber *int32 `json:"installedServerNumber,omitempty"`
	// AgreementNumber - Agreement number under which the key is requested.
	AgreementNumber *string `json:"agreementNumber,omitempty"`
	// IsEligible - <code> true </code> if user has eligible on-premises Windows physical or virtual machines, and that the requested key will only be used in their organization; <code> false </code> otherwise.
	IsEligible *bool `json:"isEligible,omitempty"`
	// ProvisioningState - READ-ONLY; Possible values include: 'Succeeded', 'Failed', 'Canceled', 'Accepted', 'Provisioning'
	ProvisioningState ProvisioningState `json:"provisioningState,omitempty"`
}

// MultipleActivationKeysCreateFuture an abstraction for monitoring and retrieving the results of a
// long-running operation.
type MultipleActivationKeysCreateFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *MultipleActivationKeysCreateFuture) Result(client MultipleActivationKeysClient) (mak MultipleActivationKey, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "windowsesu.MultipleActivationKeysCreateFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("windowsesu.MultipleActivationKeysCreateFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if mak.Response.Response, err = future.GetResult(sender); err == nil && mak.Response.Response.StatusCode != http.StatusNoContent {
		mak, err = client.CreateResponder(mak.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "windowsesu.MultipleActivationKeysCreateFuture", "Result", mak.Response.Response, "Failure responding to request")
		}
	}
	return
}

// MultipleActivationKeyUpdate MAK key details.
type MultipleActivationKeyUpdate struct {
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for MultipleActivationKeyUpdate.
func (maku MultipleActivationKeyUpdate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if maku.Tags != nil {
		objectMap["tags"] = maku.Tags
	}
	return json.Marshal(objectMap)
}

// Operation REST API operation details.
type Operation struct {
	// Name - READ-ONLY; Name of the operation.
	Name    *string           `json:"name,omitempty"`
	Display *OperationDisplay `json:"display,omitempty"`
}

// OperationDisplay meta data about operation used for display in portal.
type OperationDisplay struct {
	Provider    *string `json:"provider,omitempty"`
	Resource    *string `json:"resource,omitempty"`
	Operation   *string `json:"operation,omitempty"`
	Description *string `json:"description,omitempty"`
}

// OperationList list of available REST API operations.
type OperationList struct {
	autorest.Response `json:"-"`
	// Value - List of operations.
	Value *[]Operation `json:"value,omitempty"`
	// NextLink - READ-ONLY; Link to the next page of resources.
	NextLink *string `json:"nextLink,omitempty"`
}

// OperationListIterator provides access to a complete listing of Operation values.
type OperationListIterator struct {
	i    int
	page OperationListPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *OperationListIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationListIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *OperationListIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter OperationListIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter OperationListIterator) Response() OperationList {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter OperationListIterator) Value() Operation {
	if !iter.page.NotDone() {
		return Operation{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the OperationListIterator type.
func NewOperationListIterator(page OperationListPage) OperationListIterator {
	return OperationListIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (ol OperationList) IsEmpty() bool {
	return ol.Value == nil || len(*ol.Value) == 0
}

// operationListPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (ol OperationList) operationListPreparer(ctx context.Context) (*http.Request, error) {
	if ol.NextLink == nil || len(to.String(ol.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(ol.NextLink)))
}

// OperationListPage contains a page of Operation values.
type OperationListPage struct {
	fn func(context.Context, OperationList) (OperationList, error)
	ol OperationList
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *OperationListPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationListPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.ol)
	if err != nil {
		return err
	}
	page.ol = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *OperationListPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page OperationListPage) NotDone() bool {
	return !page.ol.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page OperationListPage) Response() OperationList {
	return page.ol
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page OperationListPage) Values() []Operation {
	if page.ol.IsEmpty() {
		return nil
	}
	return *page.ol.Value
}

// Creates a new instance of the OperationListPage type.
func NewOperationListPage(getNextPage func(context.Context, OperationList) (OperationList, error)) OperationListPage {
	return OperationListPage{fn: getNextPage}
}

// ProxyResource the resource model definition for a ARM proxy resource. It will have everything other than
// required location and tags
type ProxyResource struct {
	// ID - READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty"`
}

// Resource ...
type Resource struct {
	// ID - READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty"`
}

// TrackedResource the resource model definition for a ARM tracked top level resource
type TrackedResource struct {
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// Location - The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	// ID - READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for TrackedResource.
func (tr TrackedResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if tr.Tags != nil {
		objectMap["tags"] = tr.Tags
	}
	if tr.Location != nil {
		objectMap["location"] = tr.Location
	}
	return json.Marshal(objectMap)
}
