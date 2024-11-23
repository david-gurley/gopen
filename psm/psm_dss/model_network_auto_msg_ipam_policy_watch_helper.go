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

// NetworkAutoMsgIPAMPolicyWatchHelper AutoMsgIPAMPolicyWatchHelper is a wrapper object for watch events for IPAMPolicy objects.
type NetworkAutoMsgIPAMPolicyWatchHelper struct {
	Events *[]NetworkAutoMsgIPAMPolicyWatchHelperWatchEvent `json:"events,omitempty"`
}

// NewNetworkAutoMsgIPAMPolicyWatchHelper instantiates a new NetworkAutoMsgIPAMPolicyWatchHelper object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkAutoMsgIPAMPolicyWatchHelper() *NetworkAutoMsgIPAMPolicyWatchHelper {
	this := NetworkAutoMsgIPAMPolicyWatchHelper{}
	return &this
}

// NewNetworkAutoMsgIPAMPolicyWatchHelperWithDefaults instantiates a new NetworkAutoMsgIPAMPolicyWatchHelper object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkAutoMsgIPAMPolicyWatchHelperWithDefaults() *NetworkAutoMsgIPAMPolicyWatchHelper {
	this := NetworkAutoMsgIPAMPolicyWatchHelper{}
	return &this
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *NetworkAutoMsgIPAMPolicyWatchHelper) GetEvents() []NetworkAutoMsgIPAMPolicyWatchHelperWatchEvent {
	if o == nil || o.Events == nil {
		var ret []NetworkAutoMsgIPAMPolicyWatchHelperWatchEvent
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkAutoMsgIPAMPolicyWatchHelper) GetEventsOk() (*[]NetworkAutoMsgIPAMPolicyWatchHelperWatchEvent, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *NetworkAutoMsgIPAMPolicyWatchHelper) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []NetworkAutoMsgIPAMPolicyWatchHelperWatchEvent and assigns it to the Events field.
func (o *NetworkAutoMsgIPAMPolicyWatchHelper) SetEvents(v []NetworkAutoMsgIPAMPolicyWatchHelperWatchEvent) {
	o.Events = &v
}

func (o NetworkAutoMsgIPAMPolicyWatchHelper) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkAutoMsgIPAMPolicyWatchHelper struct {
	value *NetworkAutoMsgIPAMPolicyWatchHelper
	isSet bool
}

func (v NullableNetworkAutoMsgIPAMPolicyWatchHelper) Get() *NetworkAutoMsgIPAMPolicyWatchHelper {
	return v.value
}

func (v *NullableNetworkAutoMsgIPAMPolicyWatchHelper) Set(val *NetworkAutoMsgIPAMPolicyWatchHelper) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkAutoMsgIPAMPolicyWatchHelper) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkAutoMsgIPAMPolicyWatchHelper) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkAutoMsgIPAMPolicyWatchHelper(val *NetworkAutoMsgIPAMPolicyWatchHelper) *NullableNetworkAutoMsgIPAMPolicyWatchHelper {
	return &NullableNetworkAutoMsgIPAMPolicyWatchHelper{value: val, isSet: true}
}

func (v NullableNetworkAutoMsgIPAMPolicyWatchHelper) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkAutoMsgIPAMPolicyWatchHelper) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


