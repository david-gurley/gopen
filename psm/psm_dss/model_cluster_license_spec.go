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

// ClusterLicenseSpec struct for ClusterLicenseSpec
type ClusterLicenseSpec struct {
	// List of Feature licences applied.
	Features *[]ClusterFeature `json:"features,omitempty"`
}

// NewClusterLicenseSpec instantiates a new ClusterLicenseSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterLicenseSpec() *ClusterLicenseSpec {
	this := ClusterLicenseSpec{}
	return &this
}

// NewClusterLicenseSpecWithDefaults instantiates a new ClusterLicenseSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterLicenseSpecWithDefaults() *ClusterLicenseSpec {
	this := ClusterLicenseSpec{}
	return &this
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *ClusterLicenseSpec) GetFeatures() []ClusterFeature {
	if o == nil || o.Features == nil {
		var ret []ClusterFeature
		return ret
	}
	return *o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterLicenseSpec) GetFeaturesOk() (*[]ClusterFeature, bool) {
	if o == nil || o.Features == nil {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *ClusterLicenseSpec) HasFeatures() bool {
	if o != nil && o.Features != nil {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []ClusterFeature and assigns it to the Features field.
func (o *ClusterLicenseSpec) SetFeatures(v []ClusterFeature) {
	o.Features = &v
}

func (o ClusterLicenseSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Features != nil {
		toSerialize["features"] = o.Features
	}
	return json.Marshal(toSerialize)
}

type NullableClusterLicenseSpec struct {
	value *ClusterLicenseSpec
	isSet bool
}

func (v NullableClusterLicenseSpec) Get() *ClusterLicenseSpec {
	return v.value
}

func (v *NullableClusterLicenseSpec) Set(val *ClusterLicenseSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterLicenseSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterLicenseSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterLicenseSpec(val *ClusterLicenseSpec) *NullableClusterLicenseSpec {
	return &NullableClusterLicenseSpec{value: val, isSet: true}
}

func (v NullableClusterLicenseSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterLicenseSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


