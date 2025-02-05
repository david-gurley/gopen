/*
 * Monitoring API reference
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

// MonitoringInterfaceMirror struct for MonitoringInterfaceMirror
type MonitoringInterfaceMirror struct {
	Direction *string `json:"direction,omitempty"`
	Selectors *[]LabelsSelector `json:"selectors,omitempty"`
}

// NewMonitoringInterfaceMirror instantiates a new MonitoringInterfaceMirror object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitoringInterfaceMirror() *MonitoringInterfaceMirror {
	this := MonitoringInterfaceMirror{}
	var direction string = "both"
	this.Direction = &direction
	return &this
}

// NewMonitoringInterfaceMirrorWithDefaults instantiates a new MonitoringInterfaceMirror object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitoringInterfaceMirrorWithDefaults() *MonitoringInterfaceMirror {
	this := MonitoringInterfaceMirror{}
	var direction string = "both"
	this.Direction = &direction
	return &this
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *MonitoringInterfaceMirror) GetDirection() string {
	if o == nil || o.Direction == nil {
		var ret string
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringInterfaceMirror) GetDirectionOk() (*string, bool) {
	if o == nil || o.Direction == nil {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *MonitoringInterfaceMirror) HasDirection() bool {
	if o != nil && o.Direction != nil {
		return true
	}

	return false
}

// SetDirection gets a reference to the given string and assigns it to the Direction field.
func (o *MonitoringInterfaceMirror) SetDirection(v string) {
	o.Direction = &v
}

// GetSelectors returns the Selectors field value if set, zero value otherwise.
func (o *MonitoringInterfaceMirror) GetSelectors() []LabelsSelector {
	if o == nil || o.Selectors == nil {
		var ret []LabelsSelector
		return ret
	}
	return *o.Selectors
}

// GetSelectorsOk returns a tuple with the Selectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringInterfaceMirror) GetSelectorsOk() (*[]LabelsSelector, bool) {
	if o == nil || o.Selectors == nil {
		return nil, false
	}
	return o.Selectors, true
}

// HasSelectors returns a boolean if a field has been set.
func (o *MonitoringInterfaceMirror) HasSelectors() bool {
	if o != nil && o.Selectors != nil {
		return true
	}

	return false
}

// SetSelectors gets a reference to the given []LabelsSelector and assigns it to the Selectors field.
func (o *MonitoringInterfaceMirror) SetSelectors(v []LabelsSelector) {
	o.Selectors = &v
}

func (o MonitoringInterfaceMirror) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Direction != nil {
		toSerialize["direction"] = o.Direction
	}
	if o.Selectors != nil {
		toSerialize["selectors"] = o.Selectors
	}
	return json.Marshal(toSerialize)
}

type NullableMonitoringInterfaceMirror struct {
	value *MonitoringInterfaceMirror
	isSet bool
}

func (v NullableMonitoringInterfaceMirror) Get() *MonitoringInterfaceMirror {
	return v.value
}

func (v *NullableMonitoringInterfaceMirror) Set(val *MonitoringInterfaceMirror) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringInterfaceMirror) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringInterfaceMirror) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringInterfaceMirror(val *MonitoringInterfaceMirror) *NullableMonitoringInterfaceMirror {
	return &NullableMonitoringInterfaceMirror{value: val, isSet: true}
}

func (v NullableMonitoringInterfaceMirror) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringInterfaceMirror) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


