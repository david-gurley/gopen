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

// MonitoringAutoMsgAlertDestinationWatchHelperWatchEvent struct for MonitoringAutoMsgAlertDestinationWatchHelperWatchEvent
type MonitoringAutoMsgAlertDestinationWatchHelperWatchEvent struct {
	Object *MonitoringAlertDestination `json:"object,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewMonitoringAutoMsgAlertDestinationWatchHelperWatchEvent instantiates a new MonitoringAutoMsgAlertDestinationWatchHelperWatchEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitoringAutoMsgAlertDestinationWatchHelperWatchEvent() *MonitoringAutoMsgAlertDestinationWatchHelperWatchEvent {
	this := MonitoringAutoMsgAlertDestinationWatchHelperWatchEvent{}
	return &this
}

// NewMonitoringAutoMsgAlertDestinationWatchHelperWatchEventWithDefaults instantiates a new MonitoringAutoMsgAlertDestinationWatchHelperWatchEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitoringAutoMsgAlertDestinationWatchHelperWatchEventWithDefaults() *MonitoringAutoMsgAlertDestinationWatchHelperWatchEvent {
	this := MonitoringAutoMsgAlertDestinationWatchHelperWatchEvent{}
	return &this
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *MonitoringAutoMsgAlertDestinationWatchHelperWatchEvent) GetObject() MonitoringAlertDestination {
	if o == nil || o.Object == nil {
		var ret MonitoringAlertDestination
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringAutoMsgAlertDestinationWatchHelperWatchEvent) GetObjectOk() (*MonitoringAlertDestination, bool) {
	if o == nil || o.Object == nil {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *MonitoringAutoMsgAlertDestinationWatchHelperWatchEvent) HasObject() bool {
	if o != nil && o.Object != nil {
		return true
	}

	return false
}

// SetObject gets a reference to the given MonitoringAlertDestination and assigns it to the Object field.
func (o *MonitoringAutoMsgAlertDestinationWatchHelperWatchEvent) SetObject(v MonitoringAlertDestination) {
	o.Object = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MonitoringAutoMsgAlertDestinationWatchHelperWatchEvent) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringAutoMsgAlertDestinationWatchHelperWatchEvent) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MonitoringAutoMsgAlertDestinationWatchHelperWatchEvent) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MonitoringAutoMsgAlertDestinationWatchHelperWatchEvent) SetType(v string) {
	o.Type = &v
}

func (o MonitoringAutoMsgAlertDestinationWatchHelperWatchEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Object != nil {
		toSerialize["object"] = o.Object
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableMonitoringAutoMsgAlertDestinationWatchHelperWatchEvent struct {
	value *MonitoringAutoMsgAlertDestinationWatchHelperWatchEvent
	isSet bool
}

func (v NullableMonitoringAutoMsgAlertDestinationWatchHelperWatchEvent) Get() *MonitoringAutoMsgAlertDestinationWatchHelperWatchEvent {
	return v.value
}

func (v *NullableMonitoringAutoMsgAlertDestinationWatchHelperWatchEvent) Set(val *MonitoringAutoMsgAlertDestinationWatchHelperWatchEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringAutoMsgAlertDestinationWatchHelperWatchEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringAutoMsgAlertDestinationWatchHelperWatchEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringAutoMsgAlertDestinationWatchHelperWatchEvent(val *MonitoringAutoMsgAlertDestinationWatchHelperWatchEvent) *NullableMonitoringAutoMsgAlertDestinationWatchHelperWatchEvent {
	return &NullableMonitoringAutoMsgAlertDestinationWatchHelperWatchEvent{value: val, isSet: true}
}

func (v NullableMonitoringAutoMsgAlertDestinationWatchHelperWatchEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringAutoMsgAlertDestinationWatchHelperWatchEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


