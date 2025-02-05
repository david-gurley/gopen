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

// DSCProfileSpecInterfaces struct for DSCProfileSpecInterfaces
type DSCProfileSpecInterfaces struct {
	NumPfs *int64 `json:"num-pfs,omitempty"`
	NumVfs *int64 `json:"num-vfs,omitempty"`
}

// NewDSCProfileSpecInterfaces instantiates a new DSCProfileSpecInterfaces object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDSCProfileSpecInterfaces() *DSCProfileSpecInterfaces {
	this := DSCProfileSpecInterfaces{}
	return &this
}

// NewDSCProfileSpecInterfacesWithDefaults instantiates a new DSCProfileSpecInterfaces object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDSCProfileSpecInterfacesWithDefaults() *DSCProfileSpecInterfaces {
	this := DSCProfileSpecInterfaces{}
	return &this
}

// GetNumPfs returns the NumPfs field value if set, zero value otherwise.
func (o *DSCProfileSpecInterfaces) GetNumPfs() int64 {
	if o == nil || o.NumPfs == nil {
		var ret int64
		return ret
	}
	return *o.NumPfs
}

// GetNumPfsOk returns a tuple with the NumPfs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DSCProfileSpecInterfaces) GetNumPfsOk() (*int64, bool) {
	if o == nil || o.NumPfs == nil {
		return nil, false
	}
	return o.NumPfs, true
}

// HasNumPfs returns a boolean if a field has been set.
func (o *DSCProfileSpecInterfaces) HasNumPfs() bool {
	if o != nil && o.NumPfs != nil {
		return true
	}

	return false
}

// SetNumPfs gets a reference to the given int64 and assigns it to the NumPfs field.
func (o *DSCProfileSpecInterfaces) SetNumPfs(v int64) {
	o.NumPfs = &v
}

// GetNumVfs returns the NumVfs field value if set, zero value otherwise.
func (o *DSCProfileSpecInterfaces) GetNumVfs() int64 {
	if o == nil || o.NumVfs == nil {
		var ret int64
		return ret
	}
	return *o.NumVfs
}

// GetNumVfsOk returns a tuple with the NumVfs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DSCProfileSpecInterfaces) GetNumVfsOk() (*int64, bool) {
	if o == nil || o.NumVfs == nil {
		return nil, false
	}
	return o.NumVfs, true
}

// HasNumVfs returns a boolean if a field has been set.
func (o *DSCProfileSpecInterfaces) HasNumVfs() bool {
	if o != nil && o.NumVfs != nil {
		return true
	}

	return false
}

// SetNumVfs gets a reference to the given int64 and assigns it to the NumVfs field.
func (o *DSCProfileSpecInterfaces) SetNumVfs(v int64) {
	o.NumVfs = &v
}

func (o DSCProfileSpecInterfaces) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NumPfs != nil {
		toSerialize["num-pfs"] = o.NumPfs
	}
	if o.NumVfs != nil {
		toSerialize["num-vfs"] = o.NumVfs
	}
	return json.Marshal(toSerialize)
}

type NullableDSCProfileSpecInterfaces struct {
	value *DSCProfileSpecInterfaces
	isSet bool
}

func (v NullableDSCProfileSpecInterfaces) Get() *DSCProfileSpecInterfaces {
	return v.value
}

func (v *NullableDSCProfileSpecInterfaces) Set(val *DSCProfileSpecInterfaces) {
	v.value = val
	v.isSet = true
}

func (v NullableDSCProfileSpecInterfaces) IsSet() bool {
	return v.isSet
}

func (v *NullableDSCProfileSpecInterfaces) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDSCProfileSpecInterfaces(val *DSCProfileSpecInterfaces) *NullableDSCProfileSpecInterfaces {
	return &NullableDSCProfileSpecInterfaces{value: val, isSet: true}
}

func (v NullableDSCProfileSpecInterfaces) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDSCProfileSpecInterfaces) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


