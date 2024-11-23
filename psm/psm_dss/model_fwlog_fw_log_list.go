/*
 * Fwlog API reference
 *
 * Service name  
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package psm_dss

import (
	"encoding/json"
)

// FwlogFwLogList list of fw logs.
type FwlogFwLogList struct {
	ApiVersion *string `json:"api-version,omitempty"`
	// CountRelation to qualify total count of search results.
	CountRelation *string `json:"count-relation,omitempty"`
	Items *[]FwlogFwLog `json:"items,omitempty"`
	Kind *string `json:"kind,omitempty"`
	ListMeta *ApiListMeta `json:"list-meta,omitempty"`
	// ScrollID to scroll through results.
	ScrollId *string `json:"scroll-id,omitempty"`
}

// NewFwlogFwLogList instantiates a new FwlogFwLogList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFwlogFwLogList() *FwlogFwLogList {
	this := FwlogFwLogList{}
	var countRelation string = "notsupported"
	this.CountRelation = &countRelation
	return &this
}

// NewFwlogFwLogListWithDefaults instantiates a new FwlogFwLogList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFwlogFwLogListWithDefaults() *FwlogFwLogList {
	this := FwlogFwLogList{}
	var countRelation string = "notsupported"
	this.CountRelation = &countRelation
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *FwlogFwLogList) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FwlogFwLogList) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *FwlogFwLogList) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *FwlogFwLogList) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetCountRelation returns the CountRelation field value if set, zero value otherwise.
func (o *FwlogFwLogList) GetCountRelation() string {
	if o == nil || o.CountRelation == nil {
		var ret string
		return ret
	}
	return *o.CountRelation
}

// GetCountRelationOk returns a tuple with the CountRelation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FwlogFwLogList) GetCountRelationOk() (*string, bool) {
	if o == nil || o.CountRelation == nil {
		return nil, false
	}
	return o.CountRelation, true
}

// HasCountRelation returns a boolean if a field has been set.
func (o *FwlogFwLogList) HasCountRelation() bool {
	if o != nil && o.CountRelation != nil {
		return true
	}

	return false
}

// SetCountRelation gets a reference to the given string and assigns it to the CountRelation field.
func (o *FwlogFwLogList) SetCountRelation(v string) {
	o.CountRelation = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *FwlogFwLogList) GetItems() []FwlogFwLog {
	if o == nil || o.Items == nil {
		var ret []FwlogFwLog
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FwlogFwLogList) GetItemsOk() (*[]FwlogFwLog, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *FwlogFwLogList) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []FwlogFwLog and assigns it to the Items field.
func (o *FwlogFwLogList) SetItems(v []FwlogFwLog) {
	o.Items = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *FwlogFwLogList) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FwlogFwLogList) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *FwlogFwLogList) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *FwlogFwLogList) SetKind(v string) {
	o.Kind = &v
}

// GetListMeta returns the ListMeta field value if set, zero value otherwise.
func (o *FwlogFwLogList) GetListMeta() ApiListMeta {
	if o == nil || o.ListMeta == nil {
		var ret ApiListMeta
		return ret
	}
	return *o.ListMeta
}

// GetListMetaOk returns a tuple with the ListMeta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FwlogFwLogList) GetListMetaOk() (*ApiListMeta, bool) {
	if o == nil || o.ListMeta == nil {
		return nil, false
	}
	return o.ListMeta, true
}

// HasListMeta returns a boolean if a field has been set.
func (o *FwlogFwLogList) HasListMeta() bool {
	if o != nil && o.ListMeta != nil {
		return true
	}

	return false
}

// SetListMeta gets a reference to the given ApiListMeta and assigns it to the ListMeta field.
func (o *FwlogFwLogList) SetListMeta(v ApiListMeta) {
	o.ListMeta = &v
}

// GetScrollId returns the ScrollId field value if set, zero value otherwise.
func (o *FwlogFwLogList) GetScrollId() string {
	if o == nil || o.ScrollId == nil {
		var ret string
		return ret
	}
	return *o.ScrollId
}

// GetScrollIdOk returns a tuple with the ScrollId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FwlogFwLogList) GetScrollIdOk() (*string, bool) {
	if o == nil || o.ScrollId == nil {
		return nil, false
	}
	return o.ScrollId, true
}

// HasScrollId returns a boolean if a field has been set.
func (o *FwlogFwLogList) HasScrollId() bool {
	if o != nil && o.ScrollId != nil {
		return true
	}

	return false
}

// SetScrollId gets a reference to the given string and assigns it to the ScrollId field.
func (o *FwlogFwLogList) SetScrollId(v string) {
	o.ScrollId = &v
}

func (o FwlogFwLogList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiVersion != nil {
		toSerialize["api-version"] = o.ApiVersion
	}
	if o.CountRelation != nil {
		toSerialize["count-relation"] = o.CountRelation
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.ListMeta != nil {
		toSerialize["list-meta"] = o.ListMeta
	}
	if o.ScrollId != nil {
		toSerialize["scroll-id"] = o.ScrollId
	}
	return json.Marshal(toSerialize)
}

type NullableFwlogFwLogList struct {
	value *FwlogFwLogList
	isSet bool
}

func (v NullableFwlogFwLogList) Get() *FwlogFwLogList {
	return v.value
}

func (v *NullableFwlogFwLogList) Set(val *FwlogFwLogList) {
	v.value = val
	v.isSet = true
}

func (v NullableFwlogFwLogList) IsSet() bool {
	return v.isSet
}

func (v *NullableFwlogFwLogList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFwlogFwLogList(val *FwlogFwLogList) *NullableFwlogFwLogList {
	return &NullableFwlogFwLogList{value: val, isSet: true}
}

func (v NullableFwlogFwLogList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFwlogFwLogList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


