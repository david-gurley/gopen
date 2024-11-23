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

// ClusterVersionSpec VersionSpec contains the configuration of the Version.
type ClusterVersionSpec struct {
	// AutoRolloutDSCVersion indicates the version DSCs will be automatically upgraded upon admission. Empty value indicates no auto-rollout.
	AutoRolloutDscVersion *string `json:"auto-rollout-dsc-version,omitempty"`
}

// NewClusterVersionSpec instantiates a new ClusterVersionSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterVersionSpec() *ClusterVersionSpec {
	this := ClusterVersionSpec{}
	return &this
}

// NewClusterVersionSpecWithDefaults instantiates a new ClusterVersionSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterVersionSpecWithDefaults() *ClusterVersionSpec {
	this := ClusterVersionSpec{}
	return &this
}

// GetAutoRolloutDscVersion returns the AutoRolloutDscVersion field value if set, zero value otherwise.
func (o *ClusterVersionSpec) GetAutoRolloutDscVersion() string {
	if o == nil || o.AutoRolloutDscVersion == nil {
		var ret string
		return ret
	}
	return *o.AutoRolloutDscVersion
}

// GetAutoRolloutDscVersionOk returns a tuple with the AutoRolloutDscVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterVersionSpec) GetAutoRolloutDscVersionOk() (*string, bool) {
	if o == nil || o.AutoRolloutDscVersion == nil {
		return nil, false
	}
	return o.AutoRolloutDscVersion, true
}

// HasAutoRolloutDscVersion returns a boolean if a field has been set.
func (o *ClusterVersionSpec) HasAutoRolloutDscVersion() bool {
	if o != nil && o.AutoRolloutDscVersion != nil {
		return true
	}

	return false
}

// SetAutoRolloutDscVersion gets a reference to the given string and assigns it to the AutoRolloutDscVersion field.
func (o *ClusterVersionSpec) SetAutoRolloutDscVersion(v string) {
	o.AutoRolloutDscVersion = &v
}

func (o ClusterVersionSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AutoRolloutDscVersion != nil {
		toSerialize["auto-rollout-dsc-version"] = o.AutoRolloutDscVersion
	}
	return json.Marshal(toSerialize)
}

type NullableClusterVersionSpec struct {
	value *ClusterVersionSpec
	isSet bool
}

func (v NullableClusterVersionSpec) Get() *ClusterVersionSpec {
	return v.value
}

func (v *NullableClusterVersionSpec) Set(val *ClusterVersionSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterVersionSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterVersionSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterVersionSpec(val *ClusterVersionSpec) *NullableClusterVersionSpec {
	return &NullableClusterVersionSpec{value: val, isSet: true}
}

func (v NullableClusterVersionSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterVersionSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


