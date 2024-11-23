/*
 * Sysruntime API reference
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

// SysruntimeFlowSpec FlowSpec contains all information needed to install a uni-directional flow.
type SysruntimeFlowSpec struct {
	FlowData *SysruntimeFlowData `json:"flow-data,omitempty"`
	FlowKey *SysruntimeFlowKey `json:"flow-key,omitempty"`
}

// NewSysruntimeFlowSpec instantiates a new SysruntimeFlowSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSysruntimeFlowSpec() *SysruntimeFlowSpec {
	this := SysruntimeFlowSpec{}
	return &this
}

// NewSysruntimeFlowSpecWithDefaults instantiates a new SysruntimeFlowSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSysruntimeFlowSpecWithDefaults() *SysruntimeFlowSpec {
	this := SysruntimeFlowSpec{}
	return &this
}

// GetFlowData returns the FlowData field value if set, zero value otherwise.
func (o *SysruntimeFlowSpec) GetFlowData() SysruntimeFlowData {
	if o == nil || o.FlowData == nil {
		var ret SysruntimeFlowData
		return ret
	}
	return *o.FlowData
}

// GetFlowDataOk returns a tuple with the FlowData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SysruntimeFlowSpec) GetFlowDataOk() (*SysruntimeFlowData, bool) {
	if o == nil || o.FlowData == nil {
		return nil, false
	}
	return o.FlowData, true
}

// HasFlowData returns a boolean if a field has been set.
func (o *SysruntimeFlowSpec) HasFlowData() bool {
	if o != nil && o.FlowData != nil {
		return true
	}

	return false
}

// SetFlowData gets a reference to the given SysruntimeFlowData and assigns it to the FlowData field.
func (o *SysruntimeFlowSpec) SetFlowData(v SysruntimeFlowData) {
	o.FlowData = &v
}

// GetFlowKey returns the FlowKey field value if set, zero value otherwise.
func (o *SysruntimeFlowSpec) GetFlowKey() SysruntimeFlowKey {
	if o == nil || o.FlowKey == nil {
		var ret SysruntimeFlowKey
		return ret
	}
	return *o.FlowKey
}

// GetFlowKeyOk returns a tuple with the FlowKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SysruntimeFlowSpec) GetFlowKeyOk() (*SysruntimeFlowKey, bool) {
	if o == nil || o.FlowKey == nil {
		return nil, false
	}
	return o.FlowKey, true
}

// HasFlowKey returns a boolean if a field has been set.
func (o *SysruntimeFlowSpec) HasFlowKey() bool {
	if o != nil && o.FlowKey != nil {
		return true
	}

	return false
}

// SetFlowKey gets a reference to the given SysruntimeFlowKey and assigns it to the FlowKey field.
func (o *SysruntimeFlowSpec) SetFlowKey(v SysruntimeFlowKey) {
	o.FlowKey = &v
}

func (o SysruntimeFlowSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FlowData != nil {
		toSerialize["flow-data"] = o.FlowData
	}
	if o.FlowKey != nil {
		toSerialize["flow-key"] = o.FlowKey
	}
	return json.Marshal(toSerialize)
}

type NullableSysruntimeFlowSpec struct {
	value *SysruntimeFlowSpec
	isSet bool
}

func (v NullableSysruntimeFlowSpec) Get() *SysruntimeFlowSpec {
	return v.value
}

func (v *NullableSysruntimeFlowSpec) Set(val *SysruntimeFlowSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableSysruntimeFlowSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableSysruntimeFlowSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSysruntimeFlowSpec(val *SysruntimeFlowSpec) *NullableSysruntimeFlowSpec {
	return &NullableSysruntimeFlowSpec{value: val, isSet: true}
}

func (v NullableSysruntimeFlowSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSysruntimeFlowSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

