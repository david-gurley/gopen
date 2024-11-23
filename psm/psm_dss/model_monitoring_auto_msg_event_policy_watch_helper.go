/*
 * Monitoring API reference
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

// MonitoringAutoMsgEventPolicyWatchHelper AutoMsgEventPolicyWatchHelper is a wrapper object for watch events for EventPolicy objects.
type MonitoringAutoMsgEventPolicyWatchHelper struct {
	Events *[]MonitoringAutoMsgEventPolicyWatchHelperWatchEvent `json:"events,omitempty"`
}

// NewMonitoringAutoMsgEventPolicyWatchHelper instantiates a new MonitoringAutoMsgEventPolicyWatchHelper object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitoringAutoMsgEventPolicyWatchHelper() *MonitoringAutoMsgEventPolicyWatchHelper {
	this := MonitoringAutoMsgEventPolicyWatchHelper{}
	return &this
}

// NewMonitoringAutoMsgEventPolicyWatchHelperWithDefaults instantiates a new MonitoringAutoMsgEventPolicyWatchHelper object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitoringAutoMsgEventPolicyWatchHelperWithDefaults() *MonitoringAutoMsgEventPolicyWatchHelper {
	this := MonitoringAutoMsgEventPolicyWatchHelper{}
	return &this
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *MonitoringAutoMsgEventPolicyWatchHelper) GetEvents() []MonitoringAutoMsgEventPolicyWatchHelperWatchEvent {
	if o == nil || o.Events == nil {
		var ret []MonitoringAutoMsgEventPolicyWatchHelperWatchEvent
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringAutoMsgEventPolicyWatchHelper) GetEventsOk() (*[]MonitoringAutoMsgEventPolicyWatchHelperWatchEvent, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *MonitoringAutoMsgEventPolicyWatchHelper) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []MonitoringAutoMsgEventPolicyWatchHelperWatchEvent and assigns it to the Events field.
func (o *MonitoringAutoMsgEventPolicyWatchHelper) SetEvents(v []MonitoringAutoMsgEventPolicyWatchHelperWatchEvent) {
	o.Events = &v
}

func (o MonitoringAutoMsgEventPolicyWatchHelper) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableMonitoringAutoMsgEventPolicyWatchHelper struct {
	value *MonitoringAutoMsgEventPolicyWatchHelper
	isSet bool
}

func (v NullableMonitoringAutoMsgEventPolicyWatchHelper) Get() *MonitoringAutoMsgEventPolicyWatchHelper {
	return v.value
}

func (v *NullableMonitoringAutoMsgEventPolicyWatchHelper) Set(val *MonitoringAutoMsgEventPolicyWatchHelper) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringAutoMsgEventPolicyWatchHelper) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringAutoMsgEventPolicyWatchHelper) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringAutoMsgEventPolicyWatchHelper(val *MonitoringAutoMsgEventPolicyWatchHelper) *NullableMonitoringAutoMsgEventPolicyWatchHelper {
	return &NullableMonitoringAutoMsgEventPolicyWatchHelper{value: val, isSet: true}
}

func (v NullableMonitoringAutoMsgEventPolicyWatchHelper) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringAutoMsgEventPolicyWatchHelper) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

