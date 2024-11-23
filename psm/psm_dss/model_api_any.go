/*
 * Staging API reference
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

// ApiAny Any is wrapper around the proto Any object.
type ApiAny struct {
	TypeUrl *string `json:"type_url,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewApiAny instantiates a new ApiAny object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAny() *ApiAny {
	this := ApiAny{}
	return &this
}

// NewApiAnyWithDefaults instantiates a new ApiAny object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAnyWithDefaults() *ApiAny {
	this := ApiAny{}
	return &this
}

// GetTypeUrl returns the TypeUrl field value if set, zero value otherwise.
func (o *ApiAny) GetTypeUrl() string {
	if o == nil || o.TypeUrl == nil {
		var ret string
		return ret
	}
	return *o.TypeUrl
}

// GetTypeUrlOk returns a tuple with the TypeUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAny) GetTypeUrlOk() (*string, bool) {
	if o == nil || o.TypeUrl == nil {
		return nil, false
	}
	return o.TypeUrl, true
}

// HasTypeUrl returns a boolean if a field has been set.
func (o *ApiAny) HasTypeUrl() bool {
	if o != nil && o.TypeUrl != nil {
		return true
	}

	return false
}

// SetTypeUrl gets a reference to the given string and assigns it to the TypeUrl field.
func (o *ApiAny) SetTypeUrl(v string) {
	o.TypeUrl = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ApiAny) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAny) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ApiAny) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ApiAny) SetValue(v string) {
	o.Value = &v
}

func (o ApiAny) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TypeUrl != nil {
		toSerialize["type_url"] = o.TypeUrl
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableApiAny struct {
	value *ApiAny
	isSet bool
}

func (v NullableApiAny) Get() *ApiAny {
	return v.value
}

func (v *NullableApiAny) Set(val *ApiAny) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAny) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAny) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAny(val *ApiAny) *NullableApiAny {
	return &NullableApiAny{value: val, isSet: true}
}

func (v NullableApiAny) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAny) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


