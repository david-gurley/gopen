/*
 * Network API reference
 *
 *  
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package psm_dss

import (
	"encoding/json"
)

// ApiRDAdminValue struct for ApiRDAdminValue
type ApiRDAdminValue struct {
	Format *string `json:"Format,omitempty"`
	Value *int64 `json:"Value,omitempty"`
}

// NewApiRDAdminValue instantiates a new ApiRDAdminValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiRDAdminValue() *ApiRDAdminValue {
	this := ApiRDAdminValue{}
	return &this
}

// NewApiRDAdminValueWithDefaults instantiates a new ApiRDAdminValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiRDAdminValueWithDefaults() *ApiRDAdminValue {
	this := ApiRDAdminValue{}
	return &this
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *ApiRDAdminValue) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRDAdminValue) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *ApiRDAdminValue) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *ApiRDAdminValue) SetFormat(v string) {
	o.Format = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ApiRDAdminValue) GetValue() int64 {
	if o == nil || o.Value == nil {
		var ret int64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRDAdminValue) GetValueOk() (*int64, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ApiRDAdminValue) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given int64 and assigns it to the Value field.
func (o *ApiRDAdminValue) SetValue(v int64) {
	o.Value = &v
}

func (o ApiRDAdminValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Format != nil {
		toSerialize["Format"] = o.Format
	}
	if o.Value != nil {
		toSerialize["Value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableApiRDAdminValue struct {
	value *ApiRDAdminValue
	isSet bool
}

func (v NullableApiRDAdminValue) Get() *ApiRDAdminValue {
	return v.value
}

func (v *NullableApiRDAdminValue) Set(val *ApiRDAdminValue) {
	v.value = val
	v.isSet = true
}

func (v NullableApiRDAdminValue) IsSet() bool {
	return v.isSet
}

func (v *NullableApiRDAdminValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiRDAdminValue(val *ApiRDAdminValue) *NullableApiRDAdminValue {
	return &NullableApiRDAdminValue{value: val, isSet: true}
}

func (v NullableApiRDAdminValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiRDAdminValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

