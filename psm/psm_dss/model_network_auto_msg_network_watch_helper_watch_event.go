/*
 * Network API reference
 *
 *  
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package psm_dss

import (
	"encoding/json"
)

// NetworkAutoMsgNetworkWatchHelperWatchEvent struct for NetworkAutoMsgNetworkWatchHelperWatchEvent
type NetworkAutoMsgNetworkWatchHelperWatchEvent struct {
	Object *NetworkNetwork `json:"object,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewNetworkAutoMsgNetworkWatchHelperWatchEvent instantiates a new NetworkAutoMsgNetworkWatchHelperWatchEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkAutoMsgNetworkWatchHelperWatchEvent() *NetworkAutoMsgNetworkWatchHelperWatchEvent {
	this := NetworkAutoMsgNetworkWatchHelperWatchEvent{}
	return &this
}

// NewNetworkAutoMsgNetworkWatchHelperWatchEventWithDefaults instantiates a new NetworkAutoMsgNetworkWatchHelperWatchEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkAutoMsgNetworkWatchHelperWatchEventWithDefaults() *NetworkAutoMsgNetworkWatchHelperWatchEvent {
	this := NetworkAutoMsgNetworkWatchHelperWatchEvent{}
	return &this
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *NetworkAutoMsgNetworkWatchHelperWatchEvent) GetObject() NetworkNetwork {
	if o == nil || o.Object == nil {
		var ret NetworkNetwork
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkAutoMsgNetworkWatchHelperWatchEvent) GetObjectOk() (*NetworkNetwork, bool) {
	if o == nil || o.Object == nil {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *NetworkAutoMsgNetworkWatchHelperWatchEvent) HasObject() bool {
	if o != nil && o.Object != nil {
		return true
	}

	return false
}

// SetObject gets a reference to the given NetworkNetwork and assigns it to the Object field.
func (o *NetworkAutoMsgNetworkWatchHelperWatchEvent) SetObject(v NetworkNetwork) {
	o.Object = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NetworkAutoMsgNetworkWatchHelperWatchEvent) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkAutoMsgNetworkWatchHelperWatchEvent) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NetworkAutoMsgNetworkWatchHelperWatchEvent) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NetworkAutoMsgNetworkWatchHelperWatchEvent) SetType(v string) {
	o.Type = &v
}

func (o NetworkAutoMsgNetworkWatchHelperWatchEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Object != nil {
		toSerialize["object"] = o.Object
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkAutoMsgNetworkWatchHelperWatchEvent struct {
	value *NetworkAutoMsgNetworkWatchHelperWatchEvent
	isSet bool
}

func (v NullableNetworkAutoMsgNetworkWatchHelperWatchEvent) Get() *NetworkAutoMsgNetworkWatchHelperWatchEvent {
	return v.value
}

func (v *NullableNetworkAutoMsgNetworkWatchHelperWatchEvent) Set(val *NetworkAutoMsgNetworkWatchHelperWatchEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkAutoMsgNetworkWatchHelperWatchEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkAutoMsgNetworkWatchHelperWatchEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkAutoMsgNetworkWatchHelperWatchEvent(val *NetworkAutoMsgNetworkWatchHelperWatchEvent) *NullableNetworkAutoMsgNetworkWatchHelperWatchEvent {
	return &NullableNetworkAutoMsgNetworkWatchHelperWatchEvent{value: val, isSet: true}
}

func (v NullableNetworkAutoMsgNetworkWatchHelperWatchEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkAutoMsgNetworkWatchHelperWatchEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


