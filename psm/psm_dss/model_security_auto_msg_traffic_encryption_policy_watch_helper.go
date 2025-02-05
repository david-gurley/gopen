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

// SecurityAutoMsgTrafficEncryptionPolicyWatchHelper AutoMsgTrafficEncryptionPolicyWatchHelper is a wrapper object for watch events for TrafficEncryptionPolicy objects.
type SecurityAutoMsgTrafficEncryptionPolicyWatchHelper struct {
	Events *[]SecurityAutoMsgTrafficEncryptionPolicyWatchHelperWatchEvent `json:"events,omitempty"`
}

// NewSecurityAutoMsgTrafficEncryptionPolicyWatchHelper instantiates a new SecurityAutoMsgTrafficEncryptionPolicyWatchHelper object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityAutoMsgTrafficEncryptionPolicyWatchHelper() *SecurityAutoMsgTrafficEncryptionPolicyWatchHelper {
	this := SecurityAutoMsgTrafficEncryptionPolicyWatchHelper{}
	return &this
}

// NewSecurityAutoMsgTrafficEncryptionPolicyWatchHelperWithDefaults instantiates a new SecurityAutoMsgTrafficEncryptionPolicyWatchHelper object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityAutoMsgTrafficEncryptionPolicyWatchHelperWithDefaults() *SecurityAutoMsgTrafficEncryptionPolicyWatchHelper {
	this := SecurityAutoMsgTrafficEncryptionPolicyWatchHelper{}
	return &this
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *SecurityAutoMsgTrafficEncryptionPolicyWatchHelper) GetEvents() []SecurityAutoMsgTrafficEncryptionPolicyWatchHelperWatchEvent {
	if o == nil || o.Events == nil {
		var ret []SecurityAutoMsgTrafficEncryptionPolicyWatchHelperWatchEvent
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAutoMsgTrafficEncryptionPolicyWatchHelper) GetEventsOk() (*[]SecurityAutoMsgTrafficEncryptionPolicyWatchHelperWatchEvent, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *SecurityAutoMsgTrafficEncryptionPolicyWatchHelper) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []SecurityAutoMsgTrafficEncryptionPolicyWatchHelperWatchEvent and assigns it to the Events field.
func (o *SecurityAutoMsgTrafficEncryptionPolicyWatchHelper) SetEvents(v []SecurityAutoMsgTrafficEncryptionPolicyWatchHelperWatchEvent) {
	o.Events = &v
}

func (o SecurityAutoMsgTrafficEncryptionPolicyWatchHelper) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableSecurityAutoMsgTrafficEncryptionPolicyWatchHelper struct {
	value *SecurityAutoMsgTrafficEncryptionPolicyWatchHelper
	isSet bool
}

func (v NullableSecurityAutoMsgTrafficEncryptionPolicyWatchHelper) Get() *SecurityAutoMsgTrafficEncryptionPolicyWatchHelper {
	return v.value
}

func (v *NullableSecurityAutoMsgTrafficEncryptionPolicyWatchHelper) Set(val *SecurityAutoMsgTrafficEncryptionPolicyWatchHelper) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityAutoMsgTrafficEncryptionPolicyWatchHelper) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityAutoMsgTrafficEncryptionPolicyWatchHelper) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityAutoMsgTrafficEncryptionPolicyWatchHelper(val *SecurityAutoMsgTrafficEncryptionPolicyWatchHelper) *NullableSecurityAutoMsgTrafficEncryptionPolicyWatchHelper {
	return &NullableSecurityAutoMsgTrafficEncryptionPolicyWatchHelper{value: val, isSet: true}
}

func (v NullableSecurityAutoMsgTrafficEncryptionPolicyWatchHelper) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityAutoMsgTrafficEncryptionPolicyWatchHelper) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


