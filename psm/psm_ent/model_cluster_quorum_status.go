/*
 * Cluster API reference
 *
 * Service name  
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package psm_ent

import (
	"encoding/json"
)

// ClusterQuorumStatus QuorumStatus contains the current state of the quorum, including registered members and health.
type ClusterQuorumStatus struct {
	Members *[]ClusterQuorumMemberStatus `json:"members,omitempty"`
}

// NewClusterQuorumStatus instantiates a new ClusterQuorumStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterQuorumStatus() *ClusterQuorumStatus {
	this := ClusterQuorumStatus{}
	return &this
}

// NewClusterQuorumStatusWithDefaults instantiates a new ClusterQuorumStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterQuorumStatusWithDefaults() *ClusterQuorumStatus {
	this := ClusterQuorumStatus{}
	return &this
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *ClusterQuorumStatus) GetMembers() []ClusterQuorumMemberStatus {
	if o == nil || o.Members == nil {
		var ret []ClusterQuorumMemberStatus
		return ret
	}
	return *o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterQuorumStatus) GetMembersOk() (*[]ClusterQuorumMemberStatus, bool) {
	if o == nil || o.Members == nil {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *ClusterQuorumStatus) HasMembers() bool {
	if o != nil && o.Members != nil {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []ClusterQuorumMemberStatus and assigns it to the Members field.
func (o *ClusterQuorumStatus) SetMembers(v []ClusterQuorumMemberStatus) {
	o.Members = &v
}

func (o ClusterQuorumStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Members != nil {
		toSerialize["members"] = o.Members
	}
	return json.Marshal(toSerialize)
}

type NullableClusterQuorumStatus struct {
	value *ClusterQuorumStatus
	isSet bool
}

func (v NullableClusterQuorumStatus) Get() *ClusterQuorumStatus {
	return v.value
}

func (v *NullableClusterQuorumStatus) Set(val *ClusterQuorumStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterQuorumStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterQuorumStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterQuorumStatus(val *ClusterQuorumStatus) *NullableClusterQuorumStatus {
	return &NullableClusterQuorumStatus{value: val, isSet: true}
}

func (v NullableClusterQuorumStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterQuorumStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


