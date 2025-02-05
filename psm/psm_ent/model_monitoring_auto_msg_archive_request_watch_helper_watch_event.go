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

// MonitoringAutoMsgArchiveRequestWatchHelperWatchEvent struct for MonitoringAutoMsgArchiveRequestWatchHelperWatchEvent
type MonitoringAutoMsgArchiveRequestWatchHelperWatchEvent struct {
	Object *MonitoringArchiveRequest `json:"object,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewMonitoringAutoMsgArchiveRequestWatchHelperWatchEvent instantiates a new MonitoringAutoMsgArchiveRequestWatchHelperWatchEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitoringAutoMsgArchiveRequestWatchHelperWatchEvent() *MonitoringAutoMsgArchiveRequestWatchHelperWatchEvent {
	this := MonitoringAutoMsgArchiveRequestWatchHelperWatchEvent{}
	return &this
}

// NewMonitoringAutoMsgArchiveRequestWatchHelperWatchEventWithDefaults instantiates a new MonitoringAutoMsgArchiveRequestWatchHelperWatchEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitoringAutoMsgArchiveRequestWatchHelperWatchEventWithDefaults() *MonitoringAutoMsgArchiveRequestWatchHelperWatchEvent {
	this := MonitoringAutoMsgArchiveRequestWatchHelperWatchEvent{}
	return &this
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *MonitoringAutoMsgArchiveRequestWatchHelperWatchEvent) GetObject() MonitoringArchiveRequest {
	if o == nil || o.Object == nil {
		var ret MonitoringArchiveRequest
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringAutoMsgArchiveRequestWatchHelperWatchEvent) GetObjectOk() (*MonitoringArchiveRequest, bool) {
	if o == nil || o.Object == nil {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *MonitoringAutoMsgArchiveRequestWatchHelperWatchEvent) HasObject() bool {
	if o != nil && o.Object != nil {
		return true
	}

	return false
}

// SetObject gets a reference to the given MonitoringArchiveRequest and assigns it to the Object field.
func (o *MonitoringAutoMsgArchiveRequestWatchHelperWatchEvent) SetObject(v MonitoringArchiveRequest) {
	o.Object = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MonitoringAutoMsgArchiveRequestWatchHelperWatchEvent) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringAutoMsgArchiveRequestWatchHelperWatchEvent) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MonitoringAutoMsgArchiveRequestWatchHelperWatchEvent) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MonitoringAutoMsgArchiveRequestWatchHelperWatchEvent) SetType(v string) {
	o.Type = &v
}

func (o MonitoringAutoMsgArchiveRequestWatchHelperWatchEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Object != nil {
		toSerialize["object"] = o.Object
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableMonitoringAutoMsgArchiveRequestWatchHelperWatchEvent struct {
	value *MonitoringAutoMsgArchiveRequestWatchHelperWatchEvent
	isSet bool
}

func (v NullableMonitoringAutoMsgArchiveRequestWatchHelperWatchEvent) Get() *MonitoringAutoMsgArchiveRequestWatchHelperWatchEvent {
	return v.value
}

func (v *NullableMonitoringAutoMsgArchiveRequestWatchHelperWatchEvent) Set(val *MonitoringAutoMsgArchiveRequestWatchHelperWatchEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringAutoMsgArchiveRequestWatchHelperWatchEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringAutoMsgArchiveRequestWatchHelperWatchEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringAutoMsgArchiveRequestWatchHelperWatchEvent(val *MonitoringAutoMsgArchiveRequestWatchHelperWatchEvent) *NullableMonitoringAutoMsgArchiveRequestWatchHelperWatchEvent {
	return &NullableMonitoringAutoMsgArchiveRequestWatchHelperWatchEvent{value: val, isSet: true}
}

func (v NullableMonitoringAutoMsgArchiveRequestWatchHelperWatchEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringAutoMsgArchiveRequestWatchHelperWatchEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


