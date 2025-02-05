/*
 * Diagnostics API reference
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

// DiagnosticsAutoMsgModuleWatchHelper AutoMsgModuleWatchHelper is a wrapper object for watch events for Module objects.
type DiagnosticsAutoMsgModuleWatchHelper struct {
	Events *[]DiagnosticsAutoMsgModuleWatchHelperWatchEvent `json:"events,omitempty"`
}

// NewDiagnosticsAutoMsgModuleWatchHelper instantiates a new DiagnosticsAutoMsgModuleWatchHelper object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiagnosticsAutoMsgModuleWatchHelper() *DiagnosticsAutoMsgModuleWatchHelper {
	this := DiagnosticsAutoMsgModuleWatchHelper{}
	return &this
}

// NewDiagnosticsAutoMsgModuleWatchHelperWithDefaults instantiates a new DiagnosticsAutoMsgModuleWatchHelper object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiagnosticsAutoMsgModuleWatchHelperWithDefaults() *DiagnosticsAutoMsgModuleWatchHelper {
	this := DiagnosticsAutoMsgModuleWatchHelper{}
	return &this
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *DiagnosticsAutoMsgModuleWatchHelper) GetEvents() []DiagnosticsAutoMsgModuleWatchHelperWatchEvent {
	if o == nil || o.Events == nil {
		var ret []DiagnosticsAutoMsgModuleWatchHelperWatchEvent
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticsAutoMsgModuleWatchHelper) GetEventsOk() (*[]DiagnosticsAutoMsgModuleWatchHelperWatchEvent, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *DiagnosticsAutoMsgModuleWatchHelper) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []DiagnosticsAutoMsgModuleWatchHelperWatchEvent and assigns it to the Events field.
func (o *DiagnosticsAutoMsgModuleWatchHelper) SetEvents(v []DiagnosticsAutoMsgModuleWatchHelperWatchEvent) {
	o.Events = &v
}

func (o DiagnosticsAutoMsgModuleWatchHelper) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableDiagnosticsAutoMsgModuleWatchHelper struct {
	value *DiagnosticsAutoMsgModuleWatchHelper
	isSet bool
}

func (v NullableDiagnosticsAutoMsgModuleWatchHelper) Get() *DiagnosticsAutoMsgModuleWatchHelper {
	return v.value
}

func (v *NullableDiagnosticsAutoMsgModuleWatchHelper) Set(val *DiagnosticsAutoMsgModuleWatchHelper) {
	v.value = val
	v.isSet = true
}

func (v NullableDiagnosticsAutoMsgModuleWatchHelper) IsSet() bool {
	return v.isSet
}

func (v *NullableDiagnosticsAutoMsgModuleWatchHelper) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiagnosticsAutoMsgModuleWatchHelper(val *DiagnosticsAutoMsgModuleWatchHelper) *NullableDiagnosticsAutoMsgModuleWatchHelper {
	return &NullableDiagnosticsAutoMsgModuleWatchHelper{value: val, isSet: true}
}

func (v NullableDiagnosticsAutoMsgModuleWatchHelper) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiagnosticsAutoMsgModuleWatchHelper) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


