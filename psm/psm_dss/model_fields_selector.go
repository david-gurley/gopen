/*
 * Telemetry_query API reference
 *
 * Service name  
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package psm_dss

import (
	"encoding/json"
)

// FieldsSelector Selector is used to select objects by fields. Requirements in the selector are ANDed. A selector with no Requirements does not select anything.
type FieldsSelector struct {
	// Requirements are ANDed.
	Requirements *[]FieldsRequirement `json:"requirements,omitempty"`
}

// NewFieldsSelector instantiates a new FieldsSelector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFieldsSelector() *FieldsSelector {
	this := FieldsSelector{}
	return &this
}

// NewFieldsSelectorWithDefaults instantiates a new FieldsSelector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFieldsSelectorWithDefaults() *FieldsSelector {
	this := FieldsSelector{}
	return &this
}

// GetRequirements returns the Requirements field value if set, zero value otherwise.
func (o *FieldsSelector) GetRequirements() []FieldsRequirement {
	if o == nil || o.Requirements == nil {
		var ret []FieldsRequirement
		return ret
	}
	return *o.Requirements
}

// GetRequirementsOk returns a tuple with the Requirements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldsSelector) GetRequirementsOk() (*[]FieldsRequirement, bool) {
	if o == nil || o.Requirements == nil {
		return nil, false
	}
	return o.Requirements, true
}

// HasRequirements returns a boolean if a field has been set.
func (o *FieldsSelector) HasRequirements() bool {
	if o != nil && o.Requirements != nil {
		return true
	}

	return false
}

// SetRequirements gets a reference to the given []FieldsRequirement and assigns it to the Requirements field.
func (o *FieldsSelector) SetRequirements(v []FieldsRequirement) {
	o.Requirements = &v
}

func (o FieldsSelector) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Requirements != nil {
		toSerialize["requirements"] = o.Requirements
	}
	return json.Marshal(toSerialize)
}

type NullableFieldsSelector struct {
	value *FieldsSelector
	isSet bool
}

func (v NullableFieldsSelector) Get() *FieldsSelector {
	return v.value
}

func (v *NullableFieldsSelector) Set(val *FieldsSelector) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldsSelector) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldsSelector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldsSelector(val *FieldsSelector) *NullableFieldsSelector {
	return &NullableFieldsSelector{value: val, isSet: true}
}

func (v NullableFieldsSelector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldsSelector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


