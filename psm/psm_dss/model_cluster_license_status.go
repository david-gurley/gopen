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

// ClusterLicenseStatus struct for ClusterLicenseStatus
type ClusterLicenseStatus struct {
	// Status of current Licenced features.
	Features *[]ClusterFeatureStatus `json:"features,omitempty"`
	// Licenses that are not understood by the current running version of software.
	Unknown *[]string `json:"unknown,omitempty"`
}

// NewClusterLicenseStatus instantiates a new ClusterLicenseStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterLicenseStatus() *ClusterLicenseStatus {
	this := ClusterLicenseStatus{}
	return &this
}

// NewClusterLicenseStatusWithDefaults instantiates a new ClusterLicenseStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterLicenseStatusWithDefaults() *ClusterLicenseStatus {
	this := ClusterLicenseStatus{}
	return &this
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *ClusterLicenseStatus) GetFeatures() []ClusterFeatureStatus {
	if o == nil || o.Features == nil {
		var ret []ClusterFeatureStatus
		return ret
	}
	return *o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterLicenseStatus) GetFeaturesOk() (*[]ClusterFeatureStatus, bool) {
	if o == nil || o.Features == nil {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *ClusterLicenseStatus) HasFeatures() bool {
	if o != nil && o.Features != nil {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []ClusterFeatureStatus and assigns it to the Features field.
func (o *ClusterLicenseStatus) SetFeatures(v []ClusterFeatureStatus) {
	o.Features = &v
}

// GetUnknown returns the Unknown field value if set, zero value otherwise.
func (o *ClusterLicenseStatus) GetUnknown() []string {
	if o == nil || o.Unknown == nil {
		var ret []string
		return ret
	}
	return *o.Unknown
}

// GetUnknownOk returns a tuple with the Unknown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterLicenseStatus) GetUnknownOk() (*[]string, bool) {
	if o == nil || o.Unknown == nil {
		return nil, false
	}
	return o.Unknown, true
}

// HasUnknown returns a boolean if a field has been set.
func (o *ClusterLicenseStatus) HasUnknown() bool {
	if o != nil && o.Unknown != nil {
		return true
	}

	return false
}

// SetUnknown gets a reference to the given []string and assigns it to the Unknown field.
func (o *ClusterLicenseStatus) SetUnknown(v []string) {
	o.Unknown = &v
}

func (o ClusterLicenseStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Features != nil {
		toSerialize["features"] = o.Features
	}
	if o.Unknown != nil {
		toSerialize["unknown"] = o.Unknown
	}
	return json.Marshal(toSerialize)
}

type NullableClusterLicenseStatus struct {
	value *ClusterLicenseStatus
	isSet bool
}

func (v NullableClusterLicenseStatus) Get() *ClusterLicenseStatus {
	return v.value
}

func (v *NullableClusterLicenseStatus) Set(val *ClusterLicenseStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterLicenseStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterLicenseStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterLicenseStatus(val *ClusterLicenseStatus) *NullableClusterLicenseStatus {
	return &NullableClusterLicenseStatus{value: val, isSet: true}
}

func (v NullableClusterLicenseStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterLicenseStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


