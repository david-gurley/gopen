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

// AuthAutoMsgAuthenticationPolicyWatchHelper AutoMsgAuthenticationPolicyWatchHelper is a wrapper object for watch events for AuthenticationPolicy objects.
type AuthAutoMsgAuthenticationPolicyWatchHelper struct {
	Events *[]AuthAutoMsgAuthenticationPolicyWatchHelperWatchEvent `json:"events,omitempty"`
}

// NewAuthAutoMsgAuthenticationPolicyWatchHelper instantiates a new AuthAutoMsgAuthenticationPolicyWatchHelper object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthAutoMsgAuthenticationPolicyWatchHelper() *AuthAutoMsgAuthenticationPolicyWatchHelper {
	this := AuthAutoMsgAuthenticationPolicyWatchHelper{}
	return &this
}

// NewAuthAutoMsgAuthenticationPolicyWatchHelperWithDefaults instantiates a new AuthAutoMsgAuthenticationPolicyWatchHelper object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthAutoMsgAuthenticationPolicyWatchHelperWithDefaults() *AuthAutoMsgAuthenticationPolicyWatchHelper {
	this := AuthAutoMsgAuthenticationPolicyWatchHelper{}
	return &this
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *AuthAutoMsgAuthenticationPolicyWatchHelper) GetEvents() []AuthAutoMsgAuthenticationPolicyWatchHelperWatchEvent {
	if o == nil || o.Events == nil {
		var ret []AuthAutoMsgAuthenticationPolicyWatchHelperWatchEvent
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthAutoMsgAuthenticationPolicyWatchHelper) GetEventsOk() (*[]AuthAutoMsgAuthenticationPolicyWatchHelperWatchEvent, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *AuthAutoMsgAuthenticationPolicyWatchHelper) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []AuthAutoMsgAuthenticationPolicyWatchHelperWatchEvent and assigns it to the Events field.
func (o *AuthAutoMsgAuthenticationPolicyWatchHelper) SetEvents(v []AuthAutoMsgAuthenticationPolicyWatchHelperWatchEvent) {
	o.Events = &v
}

func (o AuthAutoMsgAuthenticationPolicyWatchHelper) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableAuthAutoMsgAuthenticationPolicyWatchHelper struct {
	value *AuthAutoMsgAuthenticationPolicyWatchHelper
	isSet bool
}

func (v NullableAuthAutoMsgAuthenticationPolicyWatchHelper) Get() *AuthAutoMsgAuthenticationPolicyWatchHelper {
	return v.value
}

func (v *NullableAuthAutoMsgAuthenticationPolicyWatchHelper) Set(val *AuthAutoMsgAuthenticationPolicyWatchHelper) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthAutoMsgAuthenticationPolicyWatchHelper) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthAutoMsgAuthenticationPolicyWatchHelper) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthAutoMsgAuthenticationPolicyWatchHelper(val *AuthAutoMsgAuthenticationPolicyWatchHelper) *NullableAuthAutoMsgAuthenticationPolicyWatchHelper {
	return &NullableAuthAutoMsgAuthenticationPolicyWatchHelper{value: val, isSet: true}
}

func (v NullableAuthAutoMsgAuthenticationPolicyWatchHelper) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthAutoMsgAuthenticationPolicyWatchHelper) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


