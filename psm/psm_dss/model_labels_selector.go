/*
 * Workload API reference
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

// LabelsSelector Selector is used to select objects by labels. Requirements in the selector are ANDed. A selector with no Requirements does not select anything.
type LabelsSelector struct {
	// Requirements are ANDed.
	Requirements *[]LabelsRequirement `json:"requirements,omitempty"`
}

// NewLabelsSelector instantiates a new LabelsSelector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLabelsSelector() *LabelsSelector {
	this := LabelsSelector{}
	return &this
}

// NewLabelsSelectorWithDefaults instantiates a new LabelsSelector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLabelsSelectorWithDefaults() *LabelsSelector {
	this := LabelsSelector{}
	return &this
}

// GetRequirements returns the Requirements field value if set, zero value otherwise.
func (o *LabelsSelector) GetRequirements() []LabelsRequirement {
	if o == nil || o.Requirements == nil {
		var ret []LabelsRequirement
		return ret
	}
	return *o.Requirements
}

// GetRequirementsOk returns a tuple with the Requirements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelsSelector) GetRequirementsOk() (*[]LabelsRequirement, bool) {
	if o == nil || o.Requirements == nil {
		return nil, false
	}
	return o.Requirements, true
}

// HasRequirements returns a boolean if a field has been set.
func (o *LabelsSelector) HasRequirements() bool {
	if o != nil && o.Requirements != nil {
		return true
	}

	return false
}

// SetRequirements gets a reference to the given []LabelsRequirement and assigns it to the Requirements field.
func (o *LabelsSelector) SetRequirements(v []LabelsRequirement) {
	o.Requirements = &v
}

func (o LabelsSelector) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Requirements != nil {
		toSerialize["requirements"] = o.Requirements
	}
	return json.Marshal(toSerialize)
}

type NullableLabelsSelector struct {
	value *LabelsSelector
	isSet bool
}

func (v NullableLabelsSelector) Get() *LabelsSelector {
	return v.value
}

func (v *NullableLabelsSelector) Set(val *LabelsSelector) {
	v.value = val
	v.isSet = true
}

func (v NullableLabelsSelector) IsSet() bool {
	return v.isSet
}

func (v *NullableLabelsSelector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabelsSelector(val *LabelsSelector) *NullableLabelsSelector {
	return &NullableLabelsSelector{value: val, isSet: true}
}

func (v NullableLabelsSelector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLabelsSelector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


