/*
 * Rollout API reference
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

// RolloutAutoMsgRolloutActionWatchHelper AutoMsgRolloutActionWatchHelper is a wrapper object for watch events for RolloutAction objects.
type RolloutAutoMsgRolloutActionWatchHelper struct {
	Events *[]RolloutAutoMsgRolloutActionWatchHelperWatchEvent `json:"events,omitempty"`
}

// NewRolloutAutoMsgRolloutActionWatchHelper instantiates a new RolloutAutoMsgRolloutActionWatchHelper object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRolloutAutoMsgRolloutActionWatchHelper() *RolloutAutoMsgRolloutActionWatchHelper {
	this := RolloutAutoMsgRolloutActionWatchHelper{}
	return &this
}

// NewRolloutAutoMsgRolloutActionWatchHelperWithDefaults instantiates a new RolloutAutoMsgRolloutActionWatchHelper object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRolloutAutoMsgRolloutActionWatchHelperWithDefaults() *RolloutAutoMsgRolloutActionWatchHelper {
	this := RolloutAutoMsgRolloutActionWatchHelper{}
	return &this
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *RolloutAutoMsgRolloutActionWatchHelper) GetEvents() []RolloutAutoMsgRolloutActionWatchHelperWatchEvent {
	if o == nil || o.Events == nil {
		var ret []RolloutAutoMsgRolloutActionWatchHelperWatchEvent
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RolloutAutoMsgRolloutActionWatchHelper) GetEventsOk() (*[]RolloutAutoMsgRolloutActionWatchHelperWatchEvent, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *RolloutAutoMsgRolloutActionWatchHelper) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []RolloutAutoMsgRolloutActionWatchHelperWatchEvent and assigns it to the Events field.
func (o *RolloutAutoMsgRolloutActionWatchHelper) SetEvents(v []RolloutAutoMsgRolloutActionWatchHelperWatchEvent) {
	o.Events = &v
}

func (o RolloutAutoMsgRolloutActionWatchHelper) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableRolloutAutoMsgRolloutActionWatchHelper struct {
	value *RolloutAutoMsgRolloutActionWatchHelper
	isSet bool
}

func (v NullableRolloutAutoMsgRolloutActionWatchHelper) Get() *RolloutAutoMsgRolloutActionWatchHelper {
	return v.value
}

func (v *NullableRolloutAutoMsgRolloutActionWatchHelper) Set(val *RolloutAutoMsgRolloutActionWatchHelper) {
	v.value = val
	v.isSet = true
}

func (v NullableRolloutAutoMsgRolloutActionWatchHelper) IsSet() bool {
	return v.isSet
}

func (v *NullableRolloutAutoMsgRolloutActionWatchHelper) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRolloutAutoMsgRolloutActionWatchHelper(val *RolloutAutoMsgRolloutActionWatchHelper) *NullableRolloutAutoMsgRolloutActionWatchHelper {
	return &NullableRolloutAutoMsgRolloutActionWatchHelper{value: val, isSet: true}
}

func (v NullableRolloutAutoMsgRolloutActionWatchHelper) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRolloutAutoMsgRolloutActionWatchHelper) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


