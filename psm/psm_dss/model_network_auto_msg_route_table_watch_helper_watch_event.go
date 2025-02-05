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

// NetworkAutoMsgRouteTableWatchHelperWatchEvent struct for NetworkAutoMsgRouteTableWatchHelperWatchEvent
type NetworkAutoMsgRouteTableWatchHelperWatchEvent struct {
	Object *NetworkRouteTable `json:"object,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewNetworkAutoMsgRouteTableWatchHelperWatchEvent instantiates a new NetworkAutoMsgRouteTableWatchHelperWatchEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkAutoMsgRouteTableWatchHelperWatchEvent() *NetworkAutoMsgRouteTableWatchHelperWatchEvent {
	this := NetworkAutoMsgRouteTableWatchHelperWatchEvent{}
	return &this
}

// NewNetworkAutoMsgRouteTableWatchHelperWatchEventWithDefaults instantiates a new NetworkAutoMsgRouteTableWatchHelperWatchEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkAutoMsgRouteTableWatchHelperWatchEventWithDefaults() *NetworkAutoMsgRouteTableWatchHelperWatchEvent {
	this := NetworkAutoMsgRouteTableWatchHelperWatchEvent{}
	return &this
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *NetworkAutoMsgRouteTableWatchHelperWatchEvent) GetObject() NetworkRouteTable {
	if o == nil || o.Object == nil {
		var ret NetworkRouteTable
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkAutoMsgRouteTableWatchHelperWatchEvent) GetObjectOk() (*NetworkRouteTable, bool) {
	if o == nil || o.Object == nil {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *NetworkAutoMsgRouteTableWatchHelperWatchEvent) HasObject() bool {
	if o != nil && o.Object != nil {
		return true
	}

	return false
}

// SetObject gets a reference to the given NetworkRouteTable and assigns it to the Object field.
func (o *NetworkAutoMsgRouteTableWatchHelperWatchEvent) SetObject(v NetworkRouteTable) {
	o.Object = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NetworkAutoMsgRouteTableWatchHelperWatchEvent) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkAutoMsgRouteTableWatchHelperWatchEvent) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NetworkAutoMsgRouteTableWatchHelperWatchEvent) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NetworkAutoMsgRouteTableWatchHelperWatchEvent) SetType(v string) {
	o.Type = &v
}

func (o NetworkAutoMsgRouteTableWatchHelperWatchEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Object != nil {
		toSerialize["object"] = o.Object
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkAutoMsgRouteTableWatchHelperWatchEvent struct {
	value *NetworkAutoMsgRouteTableWatchHelperWatchEvent
	isSet bool
}

func (v NullableNetworkAutoMsgRouteTableWatchHelperWatchEvent) Get() *NetworkAutoMsgRouteTableWatchHelperWatchEvent {
	return v.value
}

func (v *NullableNetworkAutoMsgRouteTableWatchHelperWatchEvent) Set(val *NetworkAutoMsgRouteTableWatchHelperWatchEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkAutoMsgRouteTableWatchHelperWatchEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkAutoMsgRouteTableWatchHelperWatchEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkAutoMsgRouteTableWatchHelperWatchEvent(val *NetworkAutoMsgRouteTableWatchHelperWatchEvent) *NullableNetworkAutoMsgRouteTableWatchHelperWatchEvent {
	return &NullableNetworkAutoMsgRouteTableWatchHelperWatchEvent{value: val, isSet: true}
}

func (v NullableNetworkAutoMsgRouteTableWatchHelperWatchEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkAutoMsgRouteTableWatchHelperWatchEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


