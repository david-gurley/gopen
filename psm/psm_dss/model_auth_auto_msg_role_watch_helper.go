/*
 * Auth API reference
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

// AuthAutoMsgRoleWatchHelper AutoMsgRoleWatchHelper is a wrapper object for watch events for Role objects.
type AuthAutoMsgRoleWatchHelper struct {
	Events *[]AuthAutoMsgRoleWatchHelperWatchEvent `json:"events,omitempty"`
}

// NewAuthAutoMsgRoleWatchHelper instantiates a new AuthAutoMsgRoleWatchHelper object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthAutoMsgRoleWatchHelper() *AuthAutoMsgRoleWatchHelper {
	this := AuthAutoMsgRoleWatchHelper{}
	return &this
}

// NewAuthAutoMsgRoleWatchHelperWithDefaults instantiates a new AuthAutoMsgRoleWatchHelper object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthAutoMsgRoleWatchHelperWithDefaults() *AuthAutoMsgRoleWatchHelper {
	this := AuthAutoMsgRoleWatchHelper{}
	return &this
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *AuthAutoMsgRoleWatchHelper) GetEvents() []AuthAutoMsgRoleWatchHelperWatchEvent {
	if o == nil || o.Events == nil {
		var ret []AuthAutoMsgRoleWatchHelperWatchEvent
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthAutoMsgRoleWatchHelper) GetEventsOk() (*[]AuthAutoMsgRoleWatchHelperWatchEvent, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *AuthAutoMsgRoleWatchHelper) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []AuthAutoMsgRoleWatchHelperWatchEvent and assigns it to the Events field.
func (o *AuthAutoMsgRoleWatchHelper) SetEvents(v []AuthAutoMsgRoleWatchHelperWatchEvent) {
	o.Events = &v
}

func (o AuthAutoMsgRoleWatchHelper) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableAuthAutoMsgRoleWatchHelper struct {
	value *AuthAutoMsgRoleWatchHelper
	isSet bool
}

func (v NullableAuthAutoMsgRoleWatchHelper) Get() *AuthAutoMsgRoleWatchHelper {
	return v.value
}

func (v *NullableAuthAutoMsgRoleWatchHelper) Set(val *AuthAutoMsgRoleWatchHelper) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthAutoMsgRoleWatchHelper) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthAutoMsgRoleWatchHelper) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthAutoMsgRoleWatchHelper(val *AuthAutoMsgRoleWatchHelper) *NullableAuthAutoMsgRoleWatchHelper {
	return &NullableAuthAutoMsgRoleWatchHelper{value: val, isSet: true}
}

func (v NullableAuthAutoMsgRoleWatchHelper) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthAutoMsgRoleWatchHelper) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


