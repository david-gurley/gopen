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

// AuthAutoMsgUserPreferenceWatchHelper AutoMsgUserPreferenceWatchHelper is a wrapper object for watch events for UserPreference objects.
type AuthAutoMsgUserPreferenceWatchHelper struct {
	Events *[]AuthAutoMsgUserPreferenceWatchHelperWatchEvent `json:"events,omitempty"`
}

// NewAuthAutoMsgUserPreferenceWatchHelper instantiates a new AuthAutoMsgUserPreferenceWatchHelper object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthAutoMsgUserPreferenceWatchHelper() *AuthAutoMsgUserPreferenceWatchHelper {
	this := AuthAutoMsgUserPreferenceWatchHelper{}
	return &this
}

// NewAuthAutoMsgUserPreferenceWatchHelperWithDefaults instantiates a new AuthAutoMsgUserPreferenceWatchHelper object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthAutoMsgUserPreferenceWatchHelperWithDefaults() *AuthAutoMsgUserPreferenceWatchHelper {
	this := AuthAutoMsgUserPreferenceWatchHelper{}
	return &this
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *AuthAutoMsgUserPreferenceWatchHelper) GetEvents() []AuthAutoMsgUserPreferenceWatchHelperWatchEvent {
	if o == nil || o.Events == nil {
		var ret []AuthAutoMsgUserPreferenceWatchHelperWatchEvent
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthAutoMsgUserPreferenceWatchHelper) GetEventsOk() (*[]AuthAutoMsgUserPreferenceWatchHelperWatchEvent, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *AuthAutoMsgUserPreferenceWatchHelper) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []AuthAutoMsgUserPreferenceWatchHelperWatchEvent and assigns it to the Events field.
func (o *AuthAutoMsgUserPreferenceWatchHelper) SetEvents(v []AuthAutoMsgUserPreferenceWatchHelperWatchEvent) {
	o.Events = &v
}

func (o AuthAutoMsgUserPreferenceWatchHelper) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableAuthAutoMsgUserPreferenceWatchHelper struct {
	value *AuthAutoMsgUserPreferenceWatchHelper
	isSet bool
}

func (v NullableAuthAutoMsgUserPreferenceWatchHelper) Get() *AuthAutoMsgUserPreferenceWatchHelper {
	return v.value
}

func (v *NullableAuthAutoMsgUserPreferenceWatchHelper) Set(val *AuthAutoMsgUserPreferenceWatchHelper) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthAutoMsgUserPreferenceWatchHelper) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthAutoMsgUserPreferenceWatchHelper) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthAutoMsgUserPreferenceWatchHelper(val *AuthAutoMsgUserPreferenceWatchHelper) *NullableAuthAutoMsgUserPreferenceWatchHelper {
	return &NullableAuthAutoMsgUserPreferenceWatchHelper{value: val, isSet: true}
}

func (v NullableAuthAutoMsgUserPreferenceWatchHelper) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthAutoMsgUserPreferenceWatchHelper) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


