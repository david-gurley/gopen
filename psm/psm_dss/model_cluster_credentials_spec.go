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

// ClusterCredentialsSpec struct for ClusterCredentialsSpec
type ClusterCredentialsSpec struct {
	// KeyValuePairs contains all key,value pairs that constitute credentials to access a service.
	KeyValuePairs *[]ClusterKeyValue `json:"key-value-pairs,omitempty"`
}

// NewClusterCredentialsSpec instantiates a new ClusterCredentialsSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterCredentialsSpec() *ClusterCredentialsSpec {
	this := ClusterCredentialsSpec{}
	return &this
}

// NewClusterCredentialsSpecWithDefaults instantiates a new ClusterCredentialsSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterCredentialsSpecWithDefaults() *ClusterCredentialsSpec {
	this := ClusterCredentialsSpec{}
	return &this
}

// GetKeyValuePairs returns the KeyValuePairs field value if set, zero value otherwise.
func (o *ClusterCredentialsSpec) GetKeyValuePairs() []ClusterKeyValue {
	if o == nil || o.KeyValuePairs == nil {
		var ret []ClusterKeyValue
		return ret
	}
	return *o.KeyValuePairs
}

// GetKeyValuePairsOk returns a tuple with the KeyValuePairs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterCredentialsSpec) GetKeyValuePairsOk() (*[]ClusterKeyValue, bool) {
	if o == nil || o.KeyValuePairs == nil {
		return nil, false
	}
	return o.KeyValuePairs, true
}

// HasKeyValuePairs returns a boolean if a field has been set.
func (o *ClusterCredentialsSpec) HasKeyValuePairs() bool {
	if o != nil && o.KeyValuePairs != nil {
		return true
	}

	return false
}

// SetKeyValuePairs gets a reference to the given []ClusterKeyValue and assigns it to the KeyValuePairs field.
func (o *ClusterCredentialsSpec) SetKeyValuePairs(v []ClusterKeyValue) {
	o.KeyValuePairs = &v
}

func (o ClusterCredentialsSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.KeyValuePairs != nil {
		toSerialize["key-value-pairs"] = o.KeyValuePairs
	}
	return json.Marshal(toSerialize)
}

type NullableClusterCredentialsSpec struct {
	value *ClusterCredentialsSpec
	isSet bool
}

func (v NullableClusterCredentialsSpec) Get() *ClusterCredentialsSpec {
	return v.value
}

func (v *NullableClusterCredentialsSpec) Set(val *ClusterCredentialsSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterCredentialsSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterCredentialsSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterCredentialsSpec(val *ClusterCredentialsSpec) *NullableClusterCredentialsSpec {
	return &NullableClusterCredentialsSpec{value: val, isSet: true}
}

func (v NullableClusterCredentialsSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterCredentialsSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


