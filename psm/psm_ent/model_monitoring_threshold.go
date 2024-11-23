/*
 * Monitoring API reference
 *
 * Service name  
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package psm_ent

import (
	"encoding/json"
)

// MonitoringThreshold Threshold represents the threshold value of a metric against different severities.
type MonitoringThreshold struct {
	// Raise/Create an alert when the threshold reaches this value.
	RaiseValue *string `json:"raise-value,omitempty"`
	// Severity of the alert to be created.
	Severity *string `json:"severity,omitempty"`
}

// NewMonitoringThreshold instantiates a new MonitoringThreshold object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitoringThreshold() *MonitoringThreshold {
	this := MonitoringThreshold{}
	var severity string = "info"
	this.Severity = &severity
	return &this
}

// NewMonitoringThresholdWithDefaults instantiates a new MonitoringThreshold object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitoringThresholdWithDefaults() *MonitoringThreshold {
	this := MonitoringThreshold{}
	var severity string = "info"
	this.Severity = &severity
	return &this
}

// GetRaiseValue returns the RaiseValue field value if set, zero value otherwise.
func (o *MonitoringThreshold) GetRaiseValue() string {
	if o == nil || o.RaiseValue == nil {
		var ret string
		return ret
	}
	return *o.RaiseValue
}

// GetRaiseValueOk returns a tuple with the RaiseValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringThreshold) GetRaiseValueOk() (*string, bool) {
	if o == nil || o.RaiseValue == nil {
		return nil, false
	}
	return o.RaiseValue, true
}

// HasRaiseValue returns a boolean if a field has been set.
func (o *MonitoringThreshold) HasRaiseValue() bool {
	if o != nil && o.RaiseValue != nil {
		return true
	}

	return false
}

// SetRaiseValue gets a reference to the given string and assigns it to the RaiseValue field.
func (o *MonitoringThreshold) SetRaiseValue(v string) {
	o.RaiseValue = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *MonitoringThreshold) GetSeverity() string {
	if o == nil || o.Severity == nil {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringThreshold) GetSeverityOk() (*string, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *MonitoringThreshold) HasSeverity() bool {
	if o != nil && o.Severity != nil {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *MonitoringThreshold) SetSeverity(v string) {
	o.Severity = &v
}

func (o MonitoringThreshold) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RaiseValue != nil {
		toSerialize["raise-value"] = o.RaiseValue
	}
	if o.Severity != nil {
		toSerialize["severity"] = o.Severity
	}
	return json.Marshal(toSerialize)
}

type NullableMonitoringThreshold struct {
	value *MonitoringThreshold
	isSet bool
}

func (v NullableMonitoringThreshold) Get() *MonitoringThreshold {
	return v.value
}

func (v *NullableMonitoringThreshold) Set(val *MonitoringThreshold) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringThreshold) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringThreshold) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringThreshold(val *MonitoringThreshold) *NullableMonitoringThreshold {
	return &NullableMonitoringThreshold{value: val, isSet: true}
}

func (v NullableMonitoringThreshold) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringThreshold) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


