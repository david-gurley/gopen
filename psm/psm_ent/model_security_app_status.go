/*
 * Security API reference
 *
 *  
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package psm_ent

import (
	"encoding/json"
)

// SecurityAppStatus AppStatus - status part of App object.
type SecurityAppStatus struct {
	// List of security group policies attached to the app.
	AttachedPolicies *[]string `json:"attached-policies,omitempty"`
}

// NewSecurityAppStatus instantiates a new SecurityAppStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityAppStatus() *SecurityAppStatus {
	this := SecurityAppStatus{}
	return &this
}

// NewSecurityAppStatusWithDefaults instantiates a new SecurityAppStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityAppStatusWithDefaults() *SecurityAppStatus {
	this := SecurityAppStatus{}
	return &this
}

// GetAttachedPolicies returns the AttachedPolicies field value if set, zero value otherwise.
func (o *SecurityAppStatus) GetAttachedPolicies() []string {
	if o == nil || o.AttachedPolicies == nil {
		var ret []string
		return ret
	}
	return *o.AttachedPolicies
}

// GetAttachedPoliciesOk returns a tuple with the AttachedPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAppStatus) GetAttachedPoliciesOk() (*[]string, bool) {
	if o == nil || o.AttachedPolicies == nil {
		return nil, false
	}
	return o.AttachedPolicies, true
}

// HasAttachedPolicies returns a boolean if a field has been set.
func (o *SecurityAppStatus) HasAttachedPolicies() bool {
	if o != nil && o.AttachedPolicies != nil {
		return true
	}

	return false
}

// SetAttachedPolicies gets a reference to the given []string and assigns it to the AttachedPolicies field.
func (o *SecurityAppStatus) SetAttachedPolicies(v []string) {
	o.AttachedPolicies = &v
}

func (o SecurityAppStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AttachedPolicies != nil {
		toSerialize["attached-policies"] = o.AttachedPolicies
	}
	return json.Marshal(toSerialize)
}

type NullableSecurityAppStatus struct {
	value *SecurityAppStatus
	isSet bool
}

func (v NullableSecurityAppStatus) Get() *SecurityAppStatus {
	return v.value
}

func (v *NullableSecurityAppStatus) Set(val *SecurityAppStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityAppStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityAppStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityAppStatus(val *SecurityAppStatus) *NullableSecurityAppStatus {
	return &NullableSecurityAppStatus{value: val, isSet: true}
}

func (v NullableSecurityAppStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityAppStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


