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

// ClusterHostStatus HostStatus contains the current state of the Host.
type ClusterHostStatus struct {
	// AdmittedDSCs contains a list of admitted DistributedServiceCards that are on this host.
	AdmittedDscs *[]string `json:"admitted-dscs,omitempty"`
	// MirrorSessions list of mirror sessions enabled on this host.
	MirrorSessions *[]string `json:"mirror-sessions,omitempty"`
}

// NewClusterHostStatus instantiates a new ClusterHostStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterHostStatus() *ClusterHostStatus {
	this := ClusterHostStatus{}
	return &this
}

// NewClusterHostStatusWithDefaults instantiates a new ClusterHostStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterHostStatusWithDefaults() *ClusterHostStatus {
	this := ClusterHostStatus{}
	return &this
}

// GetAdmittedDscs returns the AdmittedDscs field value if set, zero value otherwise.
func (o *ClusterHostStatus) GetAdmittedDscs() []string {
	if o == nil || o.AdmittedDscs == nil {
		var ret []string
		return ret
	}
	return *o.AdmittedDscs
}

// GetAdmittedDscsOk returns a tuple with the AdmittedDscs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterHostStatus) GetAdmittedDscsOk() (*[]string, bool) {
	if o == nil || o.AdmittedDscs == nil {
		return nil, false
	}
	return o.AdmittedDscs, true
}

// HasAdmittedDscs returns a boolean if a field has been set.
func (o *ClusterHostStatus) HasAdmittedDscs() bool {
	if o != nil && o.AdmittedDscs != nil {
		return true
	}

	return false
}

// SetAdmittedDscs gets a reference to the given []string and assigns it to the AdmittedDscs field.
func (o *ClusterHostStatus) SetAdmittedDscs(v []string) {
	o.AdmittedDscs = &v
}

// GetMirrorSessions returns the MirrorSessions field value if set, zero value otherwise.
func (o *ClusterHostStatus) GetMirrorSessions() []string {
	if o == nil || o.MirrorSessions == nil {
		var ret []string
		return ret
	}
	return *o.MirrorSessions
}

// GetMirrorSessionsOk returns a tuple with the MirrorSessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterHostStatus) GetMirrorSessionsOk() (*[]string, bool) {
	if o == nil || o.MirrorSessions == nil {
		return nil, false
	}
	return o.MirrorSessions, true
}

// HasMirrorSessions returns a boolean if a field has been set.
func (o *ClusterHostStatus) HasMirrorSessions() bool {
	if o != nil && o.MirrorSessions != nil {
		return true
	}

	return false
}

// SetMirrorSessions gets a reference to the given []string and assigns it to the MirrorSessions field.
func (o *ClusterHostStatus) SetMirrorSessions(v []string) {
	o.MirrorSessions = &v
}

func (o ClusterHostStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdmittedDscs != nil {
		toSerialize["admitted-dscs"] = o.AdmittedDscs
	}
	if o.MirrorSessions != nil {
		toSerialize["mirror-sessions"] = o.MirrorSessions
	}
	return json.Marshal(toSerialize)
}

type NullableClusterHostStatus struct {
	value *ClusterHostStatus
	isSet bool
}

func (v NullableClusterHostStatus) Get() *ClusterHostStatus {
	return v.value
}

func (v *NullableClusterHostStatus) Set(val *ClusterHostStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterHostStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterHostStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterHostStatus(val *ClusterHostStatus) *NullableClusterHostStatus {
	return &NullableClusterHostStatus{value: val, isSet: true}
}

func (v NullableClusterHostStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterHostStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


