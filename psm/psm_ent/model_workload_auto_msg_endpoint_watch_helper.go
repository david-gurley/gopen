/*
 * Workload API reference
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

// WorkloadAutoMsgEndpointWatchHelper AutoMsgEndpointWatchHelper is a wrapper object for watch events for Endpoint objects.
type WorkloadAutoMsgEndpointWatchHelper struct {
	Events *[]WorkloadAutoMsgEndpointWatchHelperWatchEvent `json:"events,omitempty"`
}

// NewWorkloadAutoMsgEndpointWatchHelper instantiates a new WorkloadAutoMsgEndpointWatchHelper object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkloadAutoMsgEndpointWatchHelper() *WorkloadAutoMsgEndpointWatchHelper {
	this := WorkloadAutoMsgEndpointWatchHelper{}
	return &this
}

// NewWorkloadAutoMsgEndpointWatchHelperWithDefaults instantiates a new WorkloadAutoMsgEndpointWatchHelper object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkloadAutoMsgEndpointWatchHelperWithDefaults() *WorkloadAutoMsgEndpointWatchHelper {
	this := WorkloadAutoMsgEndpointWatchHelper{}
	return &this
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *WorkloadAutoMsgEndpointWatchHelper) GetEvents() []WorkloadAutoMsgEndpointWatchHelperWatchEvent {
	if o == nil || o.Events == nil {
		var ret []WorkloadAutoMsgEndpointWatchHelperWatchEvent
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadAutoMsgEndpointWatchHelper) GetEventsOk() (*[]WorkloadAutoMsgEndpointWatchHelperWatchEvent, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *WorkloadAutoMsgEndpointWatchHelper) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []WorkloadAutoMsgEndpointWatchHelperWatchEvent and assigns it to the Events field.
func (o *WorkloadAutoMsgEndpointWatchHelper) SetEvents(v []WorkloadAutoMsgEndpointWatchHelperWatchEvent) {
	o.Events = &v
}

func (o WorkloadAutoMsgEndpointWatchHelper) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableWorkloadAutoMsgEndpointWatchHelper struct {
	value *WorkloadAutoMsgEndpointWatchHelper
	isSet bool
}

func (v NullableWorkloadAutoMsgEndpointWatchHelper) Get() *WorkloadAutoMsgEndpointWatchHelper {
	return v.value
}

func (v *NullableWorkloadAutoMsgEndpointWatchHelper) Set(val *WorkloadAutoMsgEndpointWatchHelper) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkloadAutoMsgEndpointWatchHelper) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkloadAutoMsgEndpointWatchHelper) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkloadAutoMsgEndpointWatchHelper(val *WorkloadAutoMsgEndpointWatchHelper) *NullableWorkloadAutoMsgEndpointWatchHelper {
	return &NullableWorkloadAutoMsgEndpointWatchHelper{value: val, isSet: true}
}

func (v NullableWorkloadAutoMsgEndpointWatchHelper) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkloadAutoMsgEndpointWatchHelper) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

