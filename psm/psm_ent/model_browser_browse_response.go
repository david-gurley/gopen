/*
 * Browser API reference
 *
 * Service name  
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package psm_ent

import (
	"encoding/json"
)

// BrowserBrowseResponse struct for BrowserBrowseResponse
type BrowserBrowseResponse struct {
	ApiVersion *string `json:"api-version,omitempty"`
	Kind *string `json:"kind,omitempty"`
	MaxDepth *int64 `json:"max-depth,omitempty"`
	Meta *ApiObjectMeta `json:"meta,omitempty"`
	Objects *map[string]BrowserObject `json:"objects,omitempty"`
	QueryType *string `json:"query-type,omitempty"`
	RootUri *string `json:"root-uri,omitempty"`
	TotalCount *int64 `json:"total-count,omitempty"`
}

// NewBrowserBrowseResponse instantiates a new BrowserBrowseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrowserBrowseResponse() *BrowserBrowseResponse {
	this := BrowserBrowseResponse{}
	return &this
}

// NewBrowserBrowseResponseWithDefaults instantiates a new BrowserBrowseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrowserBrowseResponseWithDefaults() *BrowserBrowseResponse {
	this := BrowserBrowseResponse{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *BrowserBrowseResponse) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrowserBrowseResponse) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *BrowserBrowseResponse) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *BrowserBrowseResponse) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *BrowserBrowseResponse) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrowserBrowseResponse) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *BrowserBrowseResponse) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *BrowserBrowseResponse) SetKind(v string) {
	o.Kind = &v
}

// GetMaxDepth returns the MaxDepth field value if set, zero value otherwise.
func (o *BrowserBrowseResponse) GetMaxDepth() int64 {
	if o == nil || o.MaxDepth == nil {
		var ret int64
		return ret
	}
	return *o.MaxDepth
}

// GetMaxDepthOk returns a tuple with the MaxDepth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrowserBrowseResponse) GetMaxDepthOk() (*int64, bool) {
	if o == nil || o.MaxDepth == nil {
		return nil, false
	}
	return o.MaxDepth, true
}

// HasMaxDepth returns a boolean if a field has been set.
func (o *BrowserBrowseResponse) HasMaxDepth() bool {
	if o != nil && o.MaxDepth != nil {
		return true
	}

	return false
}

// SetMaxDepth gets a reference to the given int64 and assigns it to the MaxDepth field.
func (o *BrowserBrowseResponse) SetMaxDepth(v int64) {
	o.MaxDepth = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *BrowserBrowseResponse) GetMeta() ApiObjectMeta {
	if o == nil || o.Meta == nil {
		var ret ApiObjectMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrowserBrowseResponse) GetMetaOk() (*ApiObjectMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *BrowserBrowseResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ApiObjectMeta and assigns it to the Meta field.
func (o *BrowserBrowseResponse) SetMeta(v ApiObjectMeta) {
	o.Meta = &v
}

// GetObjects returns the Objects field value if set, zero value otherwise.
func (o *BrowserBrowseResponse) GetObjects() map[string]BrowserObject {
	if o == nil || o.Objects == nil {
		var ret map[string]BrowserObject
		return ret
	}
	return *o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrowserBrowseResponse) GetObjectsOk() (*map[string]BrowserObject, bool) {
	if o == nil || o.Objects == nil {
		return nil, false
	}
	return o.Objects, true
}

// HasObjects returns a boolean if a field has been set.
func (o *BrowserBrowseResponse) HasObjects() bool {
	if o != nil && o.Objects != nil {
		return true
	}

	return false
}

// SetObjects gets a reference to the given map[string]BrowserObject and assigns it to the Objects field.
func (o *BrowserBrowseResponse) SetObjects(v map[string]BrowserObject) {
	o.Objects = &v
}

// GetQueryType returns the QueryType field value if set, zero value otherwise.
func (o *BrowserBrowseResponse) GetQueryType() string {
	if o == nil || o.QueryType == nil {
		var ret string
		return ret
	}
	return *o.QueryType
}

// GetQueryTypeOk returns a tuple with the QueryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrowserBrowseResponse) GetQueryTypeOk() (*string, bool) {
	if o == nil || o.QueryType == nil {
		return nil, false
	}
	return o.QueryType, true
}

// HasQueryType returns a boolean if a field has been set.
func (o *BrowserBrowseResponse) HasQueryType() bool {
	if o != nil && o.QueryType != nil {
		return true
	}

	return false
}

// SetQueryType gets a reference to the given string and assigns it to the QueryType field.
func (o *BrowserBrowseResponse) SetQueryType(v string) {
	o.QueryType = &v
}

// GetRootUri returns the RootUri field value if set, zero value otherwise.
func (o *BrowserBrowseResponse) GetRootUri() string {
	if o == nil || o.RootUri == nil {
		var ret string
		return ret
	}
	return *o.RootUri
}

// GetRootUriOk returns a tuple with the RootUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrowserBrowseResponse) GetRootUriOk() (*string, bool) {
	if o == nil || o.RootUri == nil {
		return nil, false
	}
	return o.RootUri, true
}

// HasRootUri returns a boolean if a field has been set.
func (o *BrowserBrowseResponse) HasRootUri() bool {
	if o != nil && o.RootUri != nil {
		return true
	}

	return false
}

// SetRootUri gets a reference to the given string and assigns it to the RootUri field.
func (o *BrowserBrowseResponse) SetRootUri(v string) {
	o.RootUri = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *BrowserBrowseResponse) GetTotalCount() int64 {
	if o == nil || o.TotalCount == nil {
		var ret int64
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrowserBrowseResponse) GetTotalCountOk() (*int64, bool) {
	if o == nil || o.TotalCount == nil {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *BrowserBrowseResponse) HasTotalCount() bool {
	if o != nil && o.TotalCount != nil {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int64 and assigns it to the TotalCount field.
func (o *BrowserBrowseResponse) SetTotalCount(v int64) {
	o.TotalCount = &v
}

func (o BrowserBrowseResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiVersion != nil {
		toSerialize["api-version"] = o.ApiVersion
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.MaxDepth != nil {
		toSerialize["max-depth"] = o.MaxDepth
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.Objects != nil {
		toSerialize["objects"] = o.Objects
	}
	if o.QueryType != nil {
		toSerialize["query-type"] = o.QueryType
	}
	if o.RootUri != nil {
		toSerialize["root-uri"] = o.RootUri
	}
	if o.TotalCount != nil {
		toSerialize["total-count"] = o.TotalCount
	}
	return json.Marshal(toSerialize)
}

type NullableBrowserBrowseResponse struct {
	value *BrowserBrowseResponse
	isSet bool
}

func (v NullableBrowserBrowseResponse) Get() *BrowserBrowseResponse {
	return v.value
}

func (v *NullableBrowserBrowseResponse) Set(val *BrowserBrowseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBrowserBrowseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBrowserBrowseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrowserBrowseResponse(val *BrowserBrowseResponse) *NullableBrowserBrowseResponse {
	return &NullableBrowserBrowseResponse{value: val, isSet: true}
}

func (v NullableBrowserBrowseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrowserBrowseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


