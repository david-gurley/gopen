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

// WorkloadIPBlock IPBlock.
type WorkloadIPBlock struct {
	// CIDRs prefix networks. Should be a valid v4 or v6 CIDR block.
	Cidr *string `json:"cidr,omitempty"`
}

// NewWorkloadIPBlock instantiates a new WorkloadIPBlock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkloadIPBlock() *WorkloadIPBlock {
	this := WorkloadIPBlock{}
	return &this
}

// NewWorkloadIPBlockWithDefaults instantiates a new WorkloadIPBlock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkloadIPBlockWithDefaults() *WorkloadIPBlock {
	this := WorkloadIPBlock{}
	return &this
}

// GetCidr returns the Cidr field value if set, zero value otherwise.
func (o *WorkloadIPBlock) GetCidr() string {
	if o == nil || o.Cidr == nil {
		var ret string
		return ret
	}
	return *o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadIPBlock) GetCidrOk() (*string, bool) {
	if o == nil || o.Cidr == nil {
		return nil, false
	}
	return o.Cidr, true
}

// HasCidr returns a boolean if a field has been set.
func (o *WorkloadIPBlock) HasCidr() bool {
	if o != nil && o.Cidr != nil {
		return true
	}

	return false
}

// SetCidr gets a reference to the given string and assigns it to the Cidr field.
func (o *WorkloadIPBlock) SetCidr(v string) {
	o.Cidr = &v
}

func (o WorkloadIPBlock) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Cidr != nil {
		toSerialize["cidr"] = o.Cidr
	}
	return json.Marshal(toSerialize)
}

type NullableWorkloadIPBlock struct {
	value *WorkloadIPBlock
	isSet bool
}

func (v NullableWorkloadIPBlock) Get() *WorkloadIPBlock {
	return v.value
}

func (v *NullableWorkloadIPBlock) Set(val *WorkloadIPBlock) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkloadIPBlock) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkloadIPBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkloadIPBlock(val *WorkloadIPBlock) *NullableWorkloadIPBlock {
	return &NullableWorkloadIPBlock{value: val, isSet: true}
}

func (v NullableWorkloadIPBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkloadIPBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

