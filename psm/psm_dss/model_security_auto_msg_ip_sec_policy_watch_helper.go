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

// SecurityAutoMsgIPSecPolicyWatchHelper AutoMsgIPSecPolicyWatchHelper is a wrapper object for watch events for IPSecPolicy objects.
type SecurityAutoMsgIPSecPolicyWatchHelper struct {
	Events *[]SecurityAutoMsgIPSecPolicyWatchHelperWatchEvent `json:"events,omitempty"`
}

// NewSecurityAutoMsgIPSecPolicyWatchHelper instantiates a new SecurityAutoMsgIPSecPolicyWatchHelper object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityAutoMsgIPSecPolicyWatchHelper() *SecurityAutoMsgIPSecPolicyWatchHelper {
	this := SecurityAutoMsgIPSecPolicyWatchHelper{}
	return &this
}

// NewSecurityAutoMsgIPSecPolicyWatchHelperWithDefaults instantiates a new SecurityAutoMsgIPSecPolicyWatchHelper object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityAutoMsgIPSecPolicyWatchHelperWithDefaults() *SecurityAutoMsgIPSecPolicyWatchHelper {
	this := SecurityAutoMsgIPSecPolicyWatchHelper{}
	return &this
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *SecurityAutoMsgIPSecPolicyWatchHelper) GetEvents() []SecurityAutoMsgIPSecPolicyWatchHelperWatchEvent {
	if o == nil || o.Events == nil {
		var ret []SecurityAutoMsgIPSecPolicyWatchHelperWatchEvent
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAutoMsgIPSecPolicyWatchHelper) GetEventsOk() (*[]SecurityAutoMsgIPSecPolicyWatchHelperWatchEvent, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *SecurityAutoMsgIPSecPolicyWatchHelper) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []SecurityAutoMsgIPSecPolicyWatchHelperWatchEvent and assigns it to the Events field.
func (o *SecurityAutoMsgIPSecPolicyWatchHelper) SetEvents(v []SecurityAutoMsgIPSecPolicyWatchHelperWatchEvent) {
	o.Events = &v
}

func (o SecurityAutoMsgIPSecPolicyWatchHelper) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableSecurityAutoMsgIPSecPolicyWatchHelper struct {
	value *SecurityAutoMsgIPSecPolicyWatchHelper
	isSet bool
}

func (v NullableSecurityAutoMsgIPSecPolicyWatchHelper) Get() *SecurityAutoMsgIPSecPolicyWatchHelper {
	return v.value
}

func (v *NullableSecurityAutoMsgIPSecPolicyWatchHelper) Set(val *SecurityAutoMsgIPSecPolicyWatchHelper) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityAutoMsgIPSecPolicyWatchHelper) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityAutoMsgIPSecPolicyWatchHelper) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityAutoMsgIPSecPolicyWatchHelper(val *SecurityAutoMsgIPSecPolicyWatchHelper) *NullableSecurityAutoMsgIPSecPolicyWatchHelper {
	return &NullableSecurityAutoMsgIPSecPolicyWatchHelper{value: val, isSet: true}
}

func (v NullableSecurityAutoMsgIPSecPolicyWatchHelper) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityAutoMsgIPSecPolicyWatchHelper) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

