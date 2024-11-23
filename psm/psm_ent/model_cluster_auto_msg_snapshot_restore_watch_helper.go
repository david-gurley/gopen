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

// ClusterAutoMsgSnapshotRestoreWatchHelper AutoMsgSnapshotRestoreWatchHelper is a wrapper object for watch events for SnapshotRestore objects.
type ClusterAutoMsgSnapshotRestoreWatchHelper struct {
	Events *[]ClusterAutoMsgSnapshotRestoreWatchHelperWatchEvent `json:"events,omitempty"`
}

// NewClusterAutoMsgSnapshotRestoreWatchHelper instantiates a new ClusterAutoMsgSnapshotRestoreWatchHelper object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterAutoMsgSnapshotRestoreWatchHelper() *ClusterAutoMsgSnapshotRestoreWatchHelper {
	this := ClusterAutoMsgSnapshotRestoreWatchHelper{}
	return &this
}

// NewClusterAutoMsgSnapshotRestoreWatchHelperWithDefaults instantiates a new ClusterAutoMsgSnapshotRestoreWatchHelper object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterAutoMsgSnapshotRestoreWatchHelperWithDefaults() *ClusterAutoMsgSnapshotRestoreWatchHelper {
	this := ClusterAutoMsgSnapshotRestoreWatchHelper{}
	return &this
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *ClusterAutoMsgSnapshotRestoreWatchHelper) GetEvents() []ClusterAutoMsgSnapshotRestoreWatchHelperWatchEvent {
	if o == nil || o.Events == nil {
		var ret []ClusterAutoMsgSnapshotRestoreWatchHelperWatchEvent
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAutoMsgSnapshotRestoreWatchHelper) GetEventsOk() (*[]ClusterAutoMsgSnapshotRestoreWatchHelperWatchEvent, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *ClusterAutoMsgSnapshotRestoreWatchHelper) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []ClusterAutoMsgSnapshotRestoreWatchHelperWatchEvent and assigns it to the Events field.
func (o *ClusterAutoMsgSnapshotRestoreWatchHelper) SetEvents(v []ClusterAutoMsgSnapshotRestoreWatchHelperWatchEvent) {
	o.Events = &v
}

func (o ClusterAutoMsgSnapshotRestoreWatchHelper) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableClusterAutoMsgSnapshotRestoreWatchHelper struct {
	value *ClusterAutoMsgSnapshotRestoreWatchHelper
	isSet bool
}

func (v NullableClusterAutoMsgSnapshotRestoreWatchHelper) Get() *ClusterAutoMsgSnapshotRestoreWatchHelper {
	return v.value
}

func (v *NullableClusterAutoMsgSnapshotRestoreWatchHelper) Set(val *ClusterAutoMsgSnapshotRestoreWatchHelper) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterAutoMsgSnapshotRestoreWatchHelper) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterAutoMsgSnapshotRestoreWatchHelper) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterAutoMsgSnapshotRestoreWatchHelper(val *ClusterAutoMsgSnapshotRestoreWatchHelper) *NullableClusterAutoMsgSnapshotRestoreWatchHelper {
	return &NullableClusterAutoMsgSnapshotRestoreWatchHelper{value: val, isSet: true}
}

func (v NullableClusterAutoMsgSnapshotRestoreWatchHelper) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterAutoMsgSnapshotRestoreWatchHelper) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


