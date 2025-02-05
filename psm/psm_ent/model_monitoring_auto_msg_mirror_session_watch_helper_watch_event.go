/*
 * Monitoring API reference
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

// MonitoringAutoMsgMirrorSessionWatchHelperWatchEvent struct for MonitoringAutoMsgMirrorSessionWatchHelperWatchEvent
type MonitoringAutoMsgMirrorSessionWatchHelperWatchEvent struct {
	Object *MonitoringMirrorSession `json:"object,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewMonitoringAutoMsgMirrorSessionWatchHelperWatchEvent instantiates a new MonitoringAutoMsgMirrorSessionWatchHelperWatchEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitoringAutoMsgMirrorSessionWatchHelperWatchEvent() *MonitoringAutoMsgMirrorSessionWatchHelperWatchEvent {
	this := MonitoringAutoMsgMirrorSessionWatchHelperWatchEvent{}
	return &this
}

// NewMonitoringAutoMsgMirrorSessionWatchHelperWatchEventWithDefaults instantiates a new MonitoringAutoMsgMirrorSessionWatchHelperWatchEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitoringAutoMsgMirrorSessionWatchHelperWatchEventWithDefaults() *MonitoringAutoMsgMirrorSessionWatchHelperWatchEvent {
	this := MonitoringAutoMsgMirrorSessionWatchHelperWatchEvent{}
	return &this
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *MonitoringAutoMsgMirrorSessionWatchHelperWatchEvent) GetObject() MonitoringMirrorSession {
	if o == nil || o.Object == nil {
		var ret MonitoringMirrorSession
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringAutoMsgMirrorSessionWatchHelperWatchEvent) GetObjectOk() (*MonitoringMirrorSession, bool) {
	if o == nil || o.Object == nil {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *MonitoringAutoMsgMirrorSessionWatchHelperWatchEvent) HasObject() bool {
	if o != nil && o.Object != nil {
		return true
	}

	return false
}

// SetObject gets a reference to the given MonitoringMirrorSession and assigns it to the Object field.
func (o *MonitoringAutoMsgMirrorSessionWatchHelperWatchEvent) SetObject(v MonitoringMirrorSession) {
	o.Object = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MonitoringAutoMsgMirrorSessionWatchHelperWatchEvent) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringAutoMsgMirrorSessionWatchHelperWatchEvent) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MonitoringAutoMsgMirrorSessionWatchHelperWatchEvent) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MonitoringAutoMsgMirrorSessionWatchHelperWatchEvent) SetType(v string) {
	o.Type = &v
}

func (o MonitoringAutoMsgMirrorSessionWatchHelperWatchEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Object != nil {
		toSerialize["object"] = o.Object
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableMonitoringAutoMsgMirrorSessionWatchHelperWatchEvent struct {
	value *MonitoringAutoMsgMirrorSessionWatchHelperWatchEvent
	isSet bool
}

func (v NullableMonitoringAutoMsgMirrorSessionWatchHelperWatchEvent) Get() *MonitoringAutoMsgMirrorSessionWatchHelperWatchEvent {
	return v.value
}

func (v *NullableMonitoringAutoMsgMirrorSessionWatchHelperWatchEvent) Set(val *MonitoringAutoMsgMirrorSessionWatchHelperWatchEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringAutoMsgMirrorSessionWatchHelperWatchEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringAutoMsgMirrorSessionWatchHelperWatchEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringAutoMsgMirrorSessionWatchHelperWatchEvent(val *MonitoringAutoMsgMirrorSessionWatchHelperWatchEvent) *NullableMonitoringAutoMsgMirrorSessionWatchHelperWatchEvent {
	return &NullableMonitoringAutoMsgMirrorSessionWatchHelperWatchEvent{value: val, isSet: true}
}

func (v NullableMonitoringAutoMsgMirrorSessionWatchHelperWatchEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringAutoMsgMirrorSessionWatchHelperWatchEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


