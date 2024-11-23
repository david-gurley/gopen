/*
 * Objstore API reference
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

// ObjstoreBucketSpec struct for ObjstoreBucketSpec
type ObjstoreBucketSpec struct {
	Description *string `json:"description,omitempty"`
}

// NewObjstoreBucketSpec instantiates a new ObjstoreBucketSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjstoreBucketSpec() *ObjstoreBucketSpec {
	this := ObjstoreBucketSpec{}
	return &this
}

// NewObjstoreBucketSpecWithDefaults instantiates a new ObjstoreBucketSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjstoreBucketSpecWithDefaults() *ObjstoreBucketSpec {
	this := ObjstoreBucketSpec{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ObjstoreBucketSpec) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjstoreBucketSpec) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ObjstoreBucketSpec) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ObjstoreBucketSpec) SetDescription(v string) {
	o.Description = &v
}

func (o ObjstoreBucketSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableObjstoreBucketSpec struct {
	value *ObjstoreBucketSpec
	isSet bool
}

func (v NullableObjstoreBucketSpec) Get() *ObjstoreBucketSpec {
	return v.value
}

func (v *NullableObjstoreBucketSpec) Set(val *ObjstoreBucketSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableObjstoreBucketSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableObjstoreBucketSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjstoreBucketSpec(val *ObjstoreBucketSpec) *NullableObjstoreBucketSpec {
	return &NullableObjstoreBucketSpec{value: val, isSet: true}
}

func (v NullableObjstoreBucketSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjstoreBucketSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


