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

// MonitoringAlertSpec User can change the state of the alert by changing the spec.
type MonitoringAlertSpec struct {
	State *string `json:"state,omitempty"`
}

// NewMonitoringAlertSpec instantiates a new MonitoringAlertSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitoringAlertSpec() *MonitoringAlertSpec {
	this := MonitoringAlertSpec{}
	var state string = "open"
	this.State = &state
	return &this
}

// NewMonitoringAlertSpecWithDefaults instantiates a new MonitoringAlertSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitoringAlertSpecWithDefaults() *MonitoringAlertSpec {
	this := MonitoringAlertSpec{}
	var state string = "open"
	this.State = &state
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *MonitoringAlertSpec) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringAlertSpec) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *MonitoringAlertSpec) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *MonitoringAlertSpec) SetState(v string) {
	o.State = &v
}

func (o MonitoringAlertSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableMonitoringAlertSpec struct {
	value *MonitoringAlertSpec
	isSet bool
}

func (v NullableMonitoringAlertSpec) Get() *MonitoringAlertSpec {
	return v.value
}

func (v *NullableMonitoringAlertSpec) Set(val *MonitoringAlertSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringAlertSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringAlertSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringAlertSpec(val *MonitoringAlertSpec) *NullableMonitoringAlertSpec {
	return &NullableMonitoringAlertSpec{value: val, isSet: true}
}

func (v NullableMonitoringAlertSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringAlertSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


