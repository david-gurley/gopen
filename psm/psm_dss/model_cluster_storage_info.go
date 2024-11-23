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

// ClusterStorageInfo Storage information.
type ClusterStorageInfo struct {
	// List of storage devices.
	Devices *[]ClusterStorageDeviceInfo `json:"devices,omitempty"`
}

// NewClusterStorageInfo instantiates a new ClusterStorageInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterStorageInfo() *ClusterStorageInfo {
	this := ClusterStorageInfo{}
	return &this
}

// NewClusterStorageInfoWithDefaults instantiates a new ClusterStorageInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterStorageInfoWithDefaults() *ClusterStorageInfo {
	this := ClusterStorageInfo{}
	return &this
}

// GetDevices returns the Devices field value if set, zero value otherwise.
func (o *ClusterStorageInfo) GetDevices() []ClusterStorageDeviceInfo {
	if o == nil || o.Devices == nil {
		var ret []ClusterStorageDeviceInfo
		return ret
	}
	return *o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterStorageInfo) GetDevicesOk() (*[]ClusterStorageDeviceInfo, bool) {
	if o == nil || o.Devices == nil {
		return nil, false
	}
	return o.Devices, true
}

// HasDevices returns a boolean if a field has been set.
func (o *ClusterStorageInfo) HasDevices() bool {
	if o != nil && o.Devices != nil {
		return true
	}

	return false
}

// SetDevices gets a reference to the given []ClusterStorageDeviceInfo and assigns it to the Devices field.
func (o *ClusterStorageInfo) SetDevices(v []ClusterStorageDeviceInfo) {
	o.Devices = &v
}

func (o ClusterStorageInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Devices != nil {
		toSerialize["devices"] = o.Devices
	}
	return json.Marshal(toSerialize)
}

type NullableClusterStorageInfo struct {
	value *ClusterStorageInfo
	isSet bool
}

func (v NullableClusterStorageInfo) Get() *ClusterStorageInfo {
	return v.value
}

func (v *NullableClusterStorageInfo) Set(val *ClusterStorageInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterStorageInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterStorageInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterStorageInfo(val *ClusterStorageInfo) *NullableClusterStorageInfo {
	return &NullableClusterStorageInfo{value: val, isSet: true}
}

func (v NullableClusterStorageInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterStorageInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


