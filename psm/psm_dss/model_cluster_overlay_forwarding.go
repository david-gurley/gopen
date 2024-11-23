/*
 * Cluster API reference
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

// ClusterOverlayForwarding struct for ClusterOverlayForwarding
type ClusterOverlayForwarding struct {
	SymmetricIrb *string `json:"symmetric-irb,omitempty"`
}

// NewClusterOverlayForwarding instantiates a new ClusterOverlayForwarding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterOverlayForwarding() *ClusterOverlayForwarding {
	this := ClusterOverlayForwarding{}
	var symmetricIrb string = "notset"
	this.SymmetricIrb = &symmetricIrb
	return &this
}

// NewClusterOverlayForwardingWithDefaults instantiates a new ClusterOverlayForwarding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterOverlayForwardingWithDefaults() *ClusterOverlayForwarding {
	this := ClusterOverlayForwarding{}
	var symmetricIrb string = "notset"
	this.SymmetricIrb = &symmetricIrb
	return &this
}

// GetSymmetricIrb returns the SymmetricIrb field value if set, zero value otherwise.
func (o *ClusterOverlayForwarding) GetSymmetricIrb() string {
	if o == nil || o.SymmetricIrb == nil {
		var ret string
		return ret
	}
	return *o.SymmetricIrb
}

// GetSymmetricIrbOk returns a tuple with the SymmetricIrb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterOverlayForwarding) GetSymmetricIrbOk() (*string, bool) {
	if o == nil || o.SymmetricIrb == nil {
		return nil, false
	}
	return o.SymmetricIrb, true
}

// HasSymmetricIrb returns a boolean if a field has been set.
func (o *ClusterOverlayForwarding) HasSymmetricIrb() bool {
	if o != nil && o.SymmetricIrb != nil {
		return true
	}

	return false
}

// SetSymmetricIrb gets a reference to the given string and assigns it to the SymmetricIrb field.
func (o *ClusterOverlayForwarding) SetSymmetricIrb(v string) {
	o.SymmetricIrb = &v
}

func (o ClusterOverlayForwarding) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SymmetricIrb != nil {
		toSerialize["symmetric-irb"] = o.SymmetricIrb
	}
	return json.Marshal(toSerialize)
}

type NullableClusterOverlayForwarding struct {
	value *ClusterOverlayForwarding
	isSet bool
}

func (v NullableClusterOverlayForwarding) Get() *ClusterOverlayForwarding {
	return v.value
}

func (v *NullableClusterOverlayForwarding) Set(val *ClusterOverlayForwarding) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterOverlayForwarding) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterOverlayForwarding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterOverlayForwarding(val *ClusterOverlayForwarding) *NullableClusterOverlayForwarding {
	return &NullableClusterOverlayForwarding{value: val, isSet: true}
}

func (v NullableClusterOverlayForwarding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterOverlayForwarding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


