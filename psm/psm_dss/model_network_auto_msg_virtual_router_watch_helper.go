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

// NetworkAutoMsgVirtualRouterWatchHelper AutoMsgVirtualRouterWatchHelper is a wrapper object for watch events for VirtualRouter objects.
type NetworkAutoMsgVirtualRouterWatchHelper struct {
	Events *[]NetworkAutoMsgVirtualRouterWatchHelperWatchEvent `json:"events,omitempty"`
}

// NewNetworkAutoMsgVirtualRouterWatchHelper instantiates a new NetworkAutoMsgVirtualRouterWatchHelper object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkAutoMsgVirtualRouterWatchHelper() *NetworkAutoMsgVirtualRouterWatchHelper {
	this := NetworkAutoMsgVirtualRouterWatchHelper{}
	return &this
}

// NewNetworkAutoMsgVirtualRouterWatchHelperWithDefaults instantiates a new NetworkAutoMsgVirtualRouterWatchHelper object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkAutoMsgVirtualRouterWatchHelperWithDefaults() *NetworkAutoMsgVirtualRouterWatchHelper {
	this := NetworkAutoMsgVirtualRouterWatchHelper{}
	return &this
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *NetworkAutoMsgVirtualRouterWatchHelper) GetEvents() []NetworkAutoMsgVirtualRouterWatchHelperWatchEvent {
	if o == nil || o.Events == nil {
		var ret []NetworkAutoMsgVirtualRouterWatchHelperWatchEvent
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkAutoMsgVirtualRouterWatchHelper) GetEventsOk() (*[]NetworkAutoMsgVirtualRouterWatchHelperWatchEvent, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *NetworkAutoMsgVirtualRouterWatchHelper) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []NetworkAutoMsgVirtualRouterWatchHelperWatchEvent and assigns it to the Events field.
func (o *NetworkAutoMsgVirtualRouterWatchHelper) SetEvents(v []NetworkAutoMsgVirtualRouterWatchHelperWatchEvent) {
	o.Events = &v
}

func (o NetworkAutoMsgVirtualRouterWatchHelper) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkAutoMsgVirtualRouterWatchHelper struct {
	value *NetworkAutoMsgVirtualRouterWatchHelper
	isSet bool
}

func (v NullableNetworkAutoMsgVirtualRouterWatchHelper) Get() *NetworkAutoMsgVirtualRouterWatchHelper {
	return v.value
}

func (v *NullableNetworkAutoMsgVirtualRouterWatchHelper) Set(val *NetworkAutoMsgVirtualRouterWatchHelper) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkAutoMsgVirtualRouterWatchHelper) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkAutoMsgVirtualRouterWatchHelper) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkAutoMsgVirtualRouterWatchHelper(val *NetworkAutoMsgVirtualRouterWatchHelper) *NullableNetworkAutoMsgVirtualRouterWatchHelper {
	return &NullableNetworkAutoMsgVirtualRouterWatchHelper{value: val, isSet: true}
}

func (v NullableNetworkAutoMsgVirtualRouterWatchHelper) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkAutoMsgVirtualRouterWatchHelper) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

