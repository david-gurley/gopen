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

// ClusterSnapshotDestination Destination for saving the configuration snapshot.
type ClusterSnapshotDestination struct {
	Type *string `json:"Type,omitempty"`
}

// NewClusterSnapshotDestination instantiates a new ClusterSnapshotDestination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterSnapshotDestination() *ClusterSnapshotDestination {
	this := ClusterSnapshotDestination{}
	var type_ string = "objectstore"
	this.Type = &type_
	return &this
}

// NewClusterSnapshotDestinationWithDefaults instantiates a new ClusterSnapshotDestination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterSnapshotDestinationWithDefaults() *ClusterSnapshotDestination {
	this := ClusterSnapshotDestination{}
	var type_ string = "objectstore"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ClusterSnapshotDestination) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSnapshotDestination) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ClusterSnapshotDestination) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ClusterSnapshotDestination) SetType(v string) {
	o.Type = &v
}

func (o ClusterSnapshotDestination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableClusterSnapshotDestination struct {
	value *ClusterSnapshotDestination
	isSet bool
}

func (v NullableClusterSnapshotDestination) Get() *ClusterSnapshotDestination {
	return v.value
}

func (v *NullableClusterSnapshotDestination) Set(val *ClusterSnapshotDestination) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterSnapshotDestination) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterSnapshotDestination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterSnapshotDestination(val *ClusterSnapshotDestination) *NullableClusterSnapshotDestination {
	return &NullableClusterSnapshotDestination{value: val, isSet: true}
}

func (v NullableClusterSnapshotDestination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterSnapshotDestination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


