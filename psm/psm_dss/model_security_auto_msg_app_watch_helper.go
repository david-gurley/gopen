/*
 * Security API reference
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

// SecurityAutoMsgAppWatchHelper AutoMsgAppWatchHelper is a wrapper object for watch events for App objects.
type SecurityAutoMsgAppWatchHelper struct {
	Events *[]SecurityAutoMsgAppWatchHelperWatchEvent `json:"events,omitempty"`
}

// NewSecurityAutoMsgAppWatchHelper instantiates a new SecurityAutoMsgAppWatchHelper object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityAutoMsgAppWatchHelper() *SecurityAutoMsgAppWatchHelper {
	this := SecurityAutoMsgAppWatchHelper{}
	return &this
}

// NewSecurityAutoMsgAppWatchHelperWithDefaults instantiates a new SecurityAutoMsgAppWatchHelper object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityAutoMsgAppWatchHelperWithDefaults() *SecurityAutoMsgAppWatchHelper {
	this := SecurityAutoMsgAppWatchHelper{}
	return &this
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *SecurityAutoMsgAppWatchHelper) GetEvents() []SecurityAutoMsgAppWatchHelperWatchEvent {
	if o == nil || o.Events == nil {
		var ret []SecurityAutoMsgAppWatchHelperWatchEvent
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAutoMsgAppWatchHelper) GetEventsOk() (*[]SecurityAutoMsgAppWatchHelperWatchEvent, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *SecurityAutoMsgAppWatchHelper) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []SecurityAutoMsgAppWatchHelperWatchEvent and assigns it to the Events field.
func (o *SecurityAutoMsgAppWatchHelper) SetEvents(v []SecurityAutoMsgAppWatchHelperWatchEvent) {
	o.Events = &v
}

func (o SecurityAutoMsgAppWatchHelper) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableSecurityAutoMsgAppWatchHelper struct {
	value *SecurityAutoMsgAppWatchHelper
	isSet bool
}

func (v NullableSecurityAutoMsgAppWatchHelper) Get() *SecurityAutoMsgAppWatchHelper {
	return v.value
}

func (v *NullableSecurityAutoMsgAppWatchHelper) Set(val *SecurityAutoMsgAppWatchHelper) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityAutoMsgAppWatchHelper) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityAutoMsgAppWatchHelper) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityAutoMsgAppWatchHelper(val *SecurityAutoMsgAppWatchHelper) *NullableSecurityAutoMsgAppWatchHelper {
	return &NullableSecurityAutoMsgAppWatchHelper{value: val, isSet: true}
}

func (v NullableSecurityAutoMsgAppWatchHelper) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityAutoMsgAppWatchHelper) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

