/*
 * Rollout API reference
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

// RolloutAutoMsgRolloutWatchHelper AutoMsgRolloutWatchHelper is a wrapper object for watch events for Rollout objects.
type RolloutAutoMsgRolloutWatchHelper struct {
	Events *[]RolloutAutoMsgRolloutWatchHelperWatchEvent `json:"events,omitempty"`
}

// NewRolloutAutoMsgRolloutWatchHelper instantiates a new RolloutAutoMsgRolloutWatchHelper object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRolloutAutoMsgRolloutWatchHelper() *RolloutAutoMsgRolloutWatchHelper {
	this := RolloutAutoMsgRolloutWatchHelper{}
	return &this
}

// NewRolloutAutoMsgRolloutWatchHelperWithDefaults instantiates a new RolloutAutoMsgRolloutWatchHelper object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRolloutAutoMsgRolloutWatchHelperWithDefaults() *RolloutAutoMsgRolloutWatchHelper {
	this := RolloutAutoMsgRolloutWatchHelper{}
	return &this
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *RolloutAutoMsgRolloutWatchHelper) GetEvents() []RolloutAutoMsgRolloutWatchHelperWatchEvent {
	if o == nil || o.Events == nil {
		var ret []RolloutAutoMsgRolloutWatchHelperWatchEvent
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RolloutAutoMsgRolloutWatchHelper) GetEventsOk() (*[]RolloutAutoMsgRolloutWatchHelperWatchEvent, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *RolloutAutoMsgRolloutWatchHelper) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []RolloutAutoMsgRolloutWatchHelperWatchEvent and assigns it to the Events field.
func (o *RolloutAutoMsgRolloutWatchHelper) SetEvents(v []RolloutAutoMsgRolloutWatchHelperWatchEvent) {
	o.Events = &v
}

func (o RolloutAutoMsgRolloutWatchHelper) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableRolloutAutoMsgRolloutWatchHelper struct {
	value *RolloutAutoMsgRolloutWatchHelper
	isSet bool
}

func (v NullableRolloutAutoMsgRolloutWatchHelper) Get() *RolloutAutoMsgRolloutWatchHelper {
	return v.value
}

func (v *NullableRolloutAutoMsgRolloutWatchHelper) Set(val *RolloutAutoMsgRolloutWatchHelper) {
	v.value = val
	v.isSet = true
}

func (v NullableRolloutAutoMsgRolloutWatchHelper) IsSet() bool {
	return v.isSet
}

func (v *NullableRolloutAutoMsgRolloutWatchHelper) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRolloutAutoMsgRolloutWatchHelper(val *RolloutAutoMsgRolloutWatchHelper) *NullableRolloutAutoMsgRolloutWatchHelper {
	return &NullableRolloutAutoMsgRolloutWatchHelper{value: val, isSet: true}
}

func (v NullableRolloutAutoMsgRolloutWatchHelper) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRolloutAutoMsgRolloutWatchHelper) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


