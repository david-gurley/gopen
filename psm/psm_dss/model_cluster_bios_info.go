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

// ClusterBiosInfo BIOS information.
type ClusterBiosInfo struct {
	// Firmware major release info.
	FwMajorVer *string `json:"fw-major-ver,omitempty"`
	// Firmware minor release info.
	FwMinorVer *string `json:"fw-minor-ver,omitempty"`
	// Vendor name.
	Vendor *string `json:"vendor,omitempty"`
	// BIOS version.
	Version *string `json:"version,omitempty"`
}

// NewClusterBiosInfo instantiates a new ClusterBiosInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterBiosInfo() *ClusterBiosInfo {
	this := ClusterBiosInfo{}
	return &this
}

// NewClusterBiosInfoWithDefaults instantiates a new ClusterBiosInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterBiosInfoWithDefaults() *ClusterBiosInfo {
	this := ClusterBiosInfo{}
	return &this
}

// GetFwMajorVer returns the FwMajorVer field value if set, zero value otherwise.
func (o *ClusterBiosInfo) GetFwMajorVer() string {
	if o == nil || o.FwMajorVer == nil {
		var ret string
		return ret
	}
	return *o.FwMajorVer
}

// GetFwMajorVerOk returns a tuple with the FwMajorVer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterBiosInfo) GetFwMajorVerOk() (*string, bool) {
	if o == nil || o.FwMajorVer == nil {
		return nil, false
	}
	return o.FwMajorVer, true
}

// HasFwMajorVer returns a boolean if a field has been set.
func (o *ClusterBiosInfo) HasFwMajorVer() bool {
	if o != nil && o.FwMajorVer != nil {
		return true
	}

	return false
}

// SetFwMajorVer gets a reference to the given string and assigns it to the FwMajorVer field.
func (o *ClusterBiosInfo) SetFwMajorVer(v string) {
	o.FwMajorVer = &v
}

// GetFwMinorVer returns the FwMinorVer field value if set, zero value otherwise.
func (o *ClusterBiosInfo) GetFwMinorVer() string {
	if o == nil || o.FwMinorVer == nil {
		var ret string
		return ret
	}
	return *o.FwMinorVer
}

// GetFwMinorVerOk returns a tuple with the FwMinorVer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterBiosInfo) GetFwMinorVerOk() (*string, bool) {
	if o == nil || o.FwMinorVer == nil {
		return nil, false
	}
	return o.FwMinorVer, true
}

// HasFwMinorVer returns a boolean if a field has been set.
func (o *ClusterBiosInfo) HasFwMinorVer() bool {
	if o != nil && o.FwMinorVer != nil {
		return true
	}

	return false
}

// SetFwMinorVer gets a reference to the given string and assigns it to the FwMinorVer field.
func (o *ClusterBiosInfo) SetFwMinorVer(v string) {
	o.FwMinorVer = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *ClusterBiosInfo) GetVendor() string {
	if o == nil || o.Vendor == nil {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterBiosInfo) GetVendorOk() (*string, bool) {
	if o == nil || o.Vendor == nil {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *ClusterBiosInfo) HasVendor() bool {
	if o != nil && o.Vendor != nil {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *ClusterBiosInfo) SetVendor(v string) {
	o.Vendor = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ClusterBiosInfo) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterBiosInfo) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ClusterBiosInfo) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ClusterBiosInfo) SetVersion(v string) {
	o.Version = &v
}

func (o ClusterBiosInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FwMajorVer != nil {
		toSerialize["fw-major-ver"] = o.FwMajorVer
	}
	if o.FwMinorVer != nil {
		toSerialize["fw-minor-ver"] = o.FwMinorVer
	}
	if o.Vendor != nil {
		toSerialize["vendor"] = o.Vendor
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableClusterBiosInfo struct {
	value *ClusterBiosInfo
	isSet bool
}

func (v NullableClusterBiosInfo) Get() *ClusterBiosInfo {
	return v.value
}

func (v *NullableClusterBiosInfo) Set(val *ClusterBiosInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterBiosInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterBiosInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterBiosInfo(val *ClusterBiosInfo) *NullableClusterBiosInfo {
	return &NullableClusterBiosInfo{value: val, isSet: true}
}

func (v NullableClusterBiosInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterBiosInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


