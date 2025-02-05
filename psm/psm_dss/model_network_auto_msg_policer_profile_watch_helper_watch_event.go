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

// NetworkAutoMsgPolicerProfileWatchHelperWatchEvent struct for NetworkAutoMsgPolicerProfileWatchHelperWatchEvent
type NetworkAutoMsgPolicerProfileWatchHelperWatchEvent struct {
	Object *NetworkPolicerProfile `json:"object,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewNetworkAutoMsgPolicerProfileWatchHelperWatchEvent instantiates a new NetworkAutoMsgPolicerProfileWatchHelperWatchEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkAutoMsgPolicerProfileWatchHelperWatchEvent() *NetworkAutoMsgPolicerProfileWatchHelperWatchEvent {
	this := NetworkAutoMsgPolicerProfileWatchHelperWatchEvent{}
	return &this
}

// NewNetworkAutoMsgPolicerProfileWatchHelperWatchEventWithDefaults instantiates a new NetworkAutoMsgPolicerProfileWatchHelperWatchEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkAutoMsgPolicerProfileWatchHelperWatchEventWithDefaults() *NetworkAutoMsgPolicerProfileWatchHelperWatchEvent {
	this := NetworkAutoMsgPolicerProfileWatchHelperWatchEvent{}
	return &this
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *NetworkAutoMsgPolicerProfileWatchHelperWatchEvent) GetObject() NetworkPolicerProfile {
	if o == nil || o.Object == nil {
		var ret NetworkPolicerProfile
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkAutoMsgPolicerProfileWatchHelperWatchEvent) GetObjectOk() (*NetworkPolicerProfile, bool) {
	if o == nil || o.Object == nil {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *NetworkAutoMsgPolicerProfileWatchHelperWatchEvent) HasObject() bool {
	if o != nil && o.Object != nil {
		return true
	}

	return false
}

// SetObject gets a reference to the given NetworkPolicerProfile and assigns it to the Object field.
func (o *NetworkAutoMsgPolicerProfileWatchHelperWatchEvent) SetObject(v NetworkPolicerProfile) {
	o.Object = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NetworkAutoMsgPolicerProfileWatchHelperWatchEvent) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkAutoMsgPolicerProfileWatchHelperWatchEvent) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NetworkAutoMsgPolicerProfileWatchHelperWatchEvent) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NetworkAutoMsgPolicerProfileWatchHelperWatchEvent) SetType(v string) {
	o.Type = &v
}

func (o NetworkAutoMsgPolicerProfileWatchHelperWatchEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Object != nil {
		toSerialize["object"] = o.Object
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkAutoMsgPolicerProfileWatchHelperWatchEvent struct {
	value *NetworkAutoMsgPolicerProfileWatchHelperWatchEvent
	isSet bool
}

func (v NullableNetworkAutoMsgPolicerProfileWatchHelperWatchEvent) Get() *NetworkAutoMsgPolicerProfileWatchHelperWatchEvent {
	return v.value
}

func (v *NullableNetworkAutoMsgPolicerProfileWatchHelperWatchEvent) Set(val *NetworkAutoMsgPolicerProfileWatchHelperWatchEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkAutoMsgPolicerProfileWatchHelperWatchEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkAutoMsgPolicerProfileWatchHelperWatchEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkAutoMsgPolicerProfileWatchHelperWatchEvent(val *NetworkAutoMsgPolicerProfileWatchHelperWatchEvent) *NullableNetworkAutoMsgPolicerProfileWatchHelperWatchEvent {
	return &NullableNetworkAutoMsgPolicerProfileWatchHelperWatchEvent{value: val, isSet: true}
}

func (v NullableNetworkAutoMsgPolicerProfileWatchHelperWatchEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkAutoMsgPolicerProfileWatchHelperWatchEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


