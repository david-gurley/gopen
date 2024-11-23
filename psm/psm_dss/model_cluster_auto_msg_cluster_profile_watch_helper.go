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

// ClusterAutoMsgClusterProfileWatchHelper AutoMsgClusterProfileWatchHelper is a wrapper object for watch events for ClusterProfile objects.
type ClusterAutoMsgClusterProfileWatchHelper struct {
	Events *[]ClusterAutoMsgClusterProfileWatchHelperWatchEvent `json:"events,omitempty"`
}

// NewClusterAutoMsgClusterProfileWatchHelper instantiates a new ClusterAutoMsgClusterProfileWatchHelper object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterAutoMsgClusterProfileWatchHelper() *ClusterAutoMsgClusterProfileWatchHelper {
	this := ClusterAutoMsgClusterProfileWatchHelper{}
	return &this
}

// NewClusterAutoMsgClusterProfileWatchHelperWithDefaults instantiates a new ClusterAutoMsgClusterProfileWatchHelper object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterAutoMsgClusterProfileWatchHelperWithDefaults() *ClusterAutoMsgClusterProfileWatchHelper {
	this := ClusterAutoMsgClusterProfileWatchHelper{}
	return &this
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *ClusterAutoMsgClusterProfileWatchHelper) GetEvents() []ClusterAutoMsgClusterProfileWatchHelperWatchEvent {
	if o == nil || o.Events == nil {
		var ret []ClusterAutoMsgClusterProfileWatchHelperWatchEvent
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAutoMsgClusterProfileWatchHelper) GetEventsOk() (*[]ClusterAutoMsgClusterProfileWatchHelperWatchEvent, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *ClusterAutoMsgClusterProfileWatchHelper) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []ClusterAutoMsgClusterProfileWatchHelperWatchEvent and assigns it to the Events field.
func (o *ClusterAutoMsgClusterProfileWatchHelper) SetEvents(v []ClusterAutoMsgClusterProfileWatchHelperWatchEvent) {
	o.Events = &v
}

func (o ClusterAutoMsgClusterProfileWatchHelper) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableClusterAutoMsgClusterProfileWatchHelper struct {
	value *ClusterAutoMsgClusterProfileWatchHelper
	isSet bool
}

func (v NullableClusterAutoMsgClusterProfileWatchHelper) Get() *ClusterAutoMsgClusterProfileWatchHelper {
	return v.value
}

func (v *NullableClusterAutoMsgClusterProfileWatchHelper) Set(val *ClusterAutoMsgClusterProfileWatchHelper) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterAutoMsgClusterProfileWatchHelper) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterAutoMsgClusterProfileWatchHelper) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterAutoMsgClusterProfileWatchHelper(val *ClusterAutoMsgClusterProfileWatchHelper) *NullableClusterAutoMsgClusterProfileWatchHelper {
	return &NullableClusterAutoMsgClusterProfileWatchHelper{value: val, isSet: true}
}

func (v NullableClusterAutoMsgClusterProfileWatchHelper) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterAutoMsgClusterProfileWatchHelper) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

