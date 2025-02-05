/*
 * Network API reference
 *
 *  
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package psm_ent

import (
	"encoding/json"
)

// NetworkAutoMsgVirtualRouterPeeringGroupWatchHelperWatchEvent struct for NetworkAutoMsgVirtualRouterPeeringGroupWatchHelperWatchEvent
type NetworkAutoMsgVirtualRouterPeeringGroupWatchHelperWatchEvent struct {
	Object *NetworkVirtualRouterPeeringGroup `json:"object,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewNetworkAutoMsgVirtualRouterPeeringGroupWatchHelperWatchEvent instantiates a new NetworkAutoMsgVirtualRouterPeeringGroupWatchHelperWatchEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkAutoMsgVirtualRouterPeeringGroupWatchHelperWatchEvent() *NetworkAutoMsgVirtualRouterPeeringGroupWatchHelperWatchEvent {
	this := NetworkAutoMsgVirtualRouterPeeringGroupWatchHelperWatchEvent{}
	return &this
}

// NewNetworkAutoMsgVirtualRouterPeeringGroupWatchHelperWatchEventWithDefaults instantiates a new NetworkAutoMsgVirtualRouterPeeringGroupWatchHelperWatchEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkAutoMsgVirtualRouterPeeringGroupWatchHelperWatchEventWithDefaults() *NetworkAutoMsgVirtualRouterPeeringGroupWatchHelperWatchEvent {
	this := NetworkAutoMsgVirtualRouterPeeringGroupWatchHelperWatchEvent{}
	return &this
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *NetworkAutoMsgVirtualRouterPeeringGroupWatchHelperWatchEvent) GetObject() NetworkVirtualRouterPeeringGroup {
	if o == nil || o.Object == nil {
		var ret NetworkVirtualRouterPeeringGroup
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkAutoMsgVirtualRouterPeeringGroupWatchHelperWatchEvent) GetObjectOk() (*NetworkVirtualRouterPeeringGroup, bool) {
	if o == nil || o.Object == nil {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *NetworkAutoMsgVirtualRouterPeeringGroupWatchHelperWatchEvent) HasObject() bool {
	if o != nil && o.Object != nil {
		return true
	}

	return false
}

// SetObject gets a reference to the given NetworkVirtualRouterPeeringGroup and assigns it to the Object field.
func (o *NetworkAutoMsgVirtualRouterPeeringGroupWatchHelperWatchEvent) SetObject(v NetworkVirtualRouterPeeringGroup) {
	o.Object = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NetworkAutoMsgVirtualRouterPeeringGroupWatchHelperWatchEvent) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkAutoMsgVirtualRouterPeeringGroupWatchHelperWatchEvent) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NetworkAutoMsgVirtualRouterPeeringGroupWatchHelperWatchEvent) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NetworkAutoMsgVirtualRouterPeeringGroupWatchHelperWatchEvent) SetType(v string) {
	o.Type = &v
}

func (o NetworkAutoMsgVirtualRouterPeeringGroupWatchHelperWatchEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Object != nil {
		toSerialize["object"] = o.Object
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkAutoMsgVirtualRouterPeeringGroupWatchHelperWatchEvent struct {
	value *NetworkAutoMsgVirtualRouterPeeringGroupWatchHelperWatchEvent
	isSet bool
}

func (v NullableNetworkAutoMsgVirtualRouterPeeringGroupWatchHelperWatchEvent) Get() *NetworkAutoMsgVirtualRouterPeeringGroupWatchHelperWatchEvent {
	return v.value
}

func (v *NullableNetworkAutoMsgVirtualRouterPeeringGroupWatchHelperWatchEvent) Set(val *NetworkAutoMsgVirtualRouterPeeringGroupWatchHelperWatchEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkAutoMsgVirtualRouterPeeringGroupWatchHelperWatchEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkAutoMsgVirtualRouterPeeringGroupWatchHelperWatchEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkAutoMsgVirtualRouterPeeringGroupWatchHelperWatchEvent(val *NetworkAutoMsgVirtualRouterPeeringGroupWatchHelperWatchEvent) *NullableNetworkAutoMsgVirtualRouterPeeringGroupWatchHelperWatchEvent {
	return &NullableNetworkAutoMsgVirtualRouterPeeringGroupWatchHelperWatchEvent{value: val, isSet: true}
}

func (v NullableNetworkAutoMsgVirtualRouterPeeringGroupWatchHelperWatchEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkAutoMsgVirtualRouterPeeringGroupWatchHelperWatchEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


