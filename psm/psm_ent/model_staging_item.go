/*
 * Staging API reference
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

// StagingItem struct for StagingItem
type StagingItem struct {
	Method *string `json:"method,omitempty"`
	Object *ApiAny `json:"object,omitempty"`
	Uri *string `json:"uri,omitempty"`
}

// NewStagingItem instantiates a new StagingItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStagingItem() *StagingItem {
	this := StagingItem{}
	return &this
}

// NewStagingItemWithDefaults instantiates a new StagingItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStagingItemWithDefaults() *StagingItem {
	this := StagingItem{}
	return &this
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *StagingItem) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagingItem) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *StagingItem) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *StagingItem) SetMethod(v string) {
	o.Method = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *StagingItem) GetObject() ApiAny {
	if o == nil || o.Object == nil {
		var ret ApiAny
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagingItem) GetObjectOk() (*ApiAny, bool) {
	if o == nil || o.Object == nil {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *StagingItem) HasObject() bool {
	if o != nil && o.Object != nil {
		return true
	}

	return false
}

// SetObject gets a reference to the given ApiAny and assigns it to the Object field.
func (o *StagingItem) SetObject(v ApiAny) {
	o.Object = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *StagingItem) GetUri() string {
	if o == nil || o.Uri == nil {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagingItem) GetUriOk() (*string, bool) {
	if o == nil || o.Uri == nil {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *StagingItem) HasUri() bool {
	if o != nil && o.Uri != nil {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *StagingItem) SetUri(v string) {
	o.Uri = &v
}

func (o StagingItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}
	if o.Object != nil {
		toSerialize["object"] = o.Object
	}
	if o.Uri != nil {
		toSerialize["uri"] = o.Uri
	}
	return json.Marshal(toSerialize)
}

type NullableStagingItem struct {
	value *StagingItem
	isSet bool
}

func (v NullableStagingItem) Get() *StagingItem {
	return v.value
}

func (v *NullableStagingItem) Set(val *StagingItem) {
	v.value = val
	v.isSet = true
}

func (v NullableStagingItem) IsSet() bool {
	return v.isSet
}

func (v *NullableStagingItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStagingItem(val *StagingItem) *NullableStagingItem {
	return &NullableStagingItem{value: val, isSet: true}
}

func (v NullableStagingItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStagingItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

