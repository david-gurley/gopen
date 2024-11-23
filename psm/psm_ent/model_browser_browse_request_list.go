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

// BrowserBrowseRequestList BrowseRequest is the query request for the dependency tree.
type BrowserBrowseRequestList struct {
	ApiVersion *string `json:"api-version,omitempty"`
	Kind *string `json:"kind,omitempty"`
	Meta *ApiObjectMeta `json:"meta,omitempty"`
	Requestlist *[]BrowserBrowseRequestObject `json:"requestlist,omitempty"`
}

// NewBrowserBrowseRequestList instantiates a new BrowserBrowseRequestList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrowserBrowseRequestList() *BrowserBrowseRequestList {
	this := BrowserBrowseRequestList{}
	return &this
}

// NewBrowserBrowseRequestListWithDefaults instantiates a new BrowserBrowseRequestList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrowserBrowseRequestListWithDefaults() *BrowserBrowseRequestList {
	this := BrowserBrowseRequestList{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *BrowserBrowseRequestList) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrowserBrowseRequestList) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *BrowserBrowseRequestList) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *BrowserBrowseRequestList) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *BrowserBrowseRequestList) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrowserBrowseRequestList) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *BrowserBrowseRequestList) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *BrowserBrowseRequestList) SetKind(v string) {
	o.Kind = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *BrowserBrowseRequestList) GetMeta() ApiObjectMeta {
	if o == nil || o.Meta == nil {
		var ret ApiObjectMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrowserBrowseRequestList) GetMetaOk() (*ApiObjectMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *BrowserBrowseRequestList) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ApiObjectMeta and assigns it to the Meta field.
func (o *BrowserBrowseRequestList) SetMeta(v ApiObjectMeta) {
	o.Meta = &v
}

// GetRequestlist returns the Requestlist field value if set, zero value otherwise.
func (o *BrowserBrowseRequestList) GetRequestlist() []BrowserBrowseRequestObject {
	if o == nil || o.Requestlist == nil {
		var ret []BrowserBrowseRequestObject
		return ret
	}
	return *o.Requestlist
}

// GetRequestlistOk returns a tuple with the Requestlist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrowserBrowseRequestList) GetRequestlistOk() (*[]BrowserBrowseRequestObject, bool) {
	if o == nil || o.Requestlist == nil {
		return nil, false
	}
	return o.Requestlist, true
}

// HasRequestlist returns a boolean if a field has been set.
func (o *BrowserBrowseRequestList) HasRequestlist() bool {
	if o != nil && o.Requestlist != nil {
		return true
	}

	return false
}

// SetRequestlist gets a reference to the given []BrowserBrowseRequestObject and assigns it to the Requestlist field.
func (o *BrowserBrowseRequestList) SetRequestlist(v []BrowserBrowseRequestObject) {
	o.Requestlist = &v
}

func (o BrowserBrowseRequestList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiVersion != nil {
		toSerialize["api-version"] = o.ApiVersion
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.Requestlist != nil {
		toSerialize["requestlist"] = o.Requestlist
	}
	return json.Marshal(toSerialize)
}

type NullableBrowserBrowseRequestList struct {
	value *BrowserBrowseRequestList
	isSet bool
}

func (v NullableBrowserBrowseRequestList) Get() *BrowserBrowseRequestList {
	return v.value
}

func (v *NullableBrowserBrowseRequestList) Set(val *BrowserBrowseRequestList) {
	v.value = val
	v.isSet = true
}

func (v NullableBrowserBrowseRequestList) IsSet() bool {
	return v.isSet
}

func (v *NullableBrowserBrowseRequestList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrowserBrowseRequestList(val *BrowserBrowseRequestList) *NullableBrowserBrowseRequestList {
	return &NullableBrowserBrowseRequestList{value: val, isSet: true}
}

func (v NullableBrowserBrowseRequestList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrowserBrowseRequestList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


