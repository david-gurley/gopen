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

// MonitoringMeasurementCriteria Measurement window and function to be used for monitoring the metric. nil `measurement` indicates that the policy will act on the instantaneous value of the metric that gets reported every 30s.
type MonitoringMeasurementCriteria struct {
	// Aggregate function to be applied on the metric values that were monitored over a window/interval.
	Function *string `json:"function,omitempty"`
	// The length of time the metric will be monitored/observed before running the values against thresholds for alert creation/resolution. ui-hint: Allowed values - 5m, 10m, 30m, 1h.
	Window *string `json:"window,omitempty"`
}

// NewMonitoringMeasurementCriteria instantiates a new MonitoringMeasurementCriteria object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitoringMeasurementCriteria() *MonitoringMeasurementCriteria {
	this := MonitoringMeasurementCriteria{}
	var function string = "none_function"
	this.Function = &function
	return &this
}

// NewMonitoringMeasurementCriteriaWithDefaults instantiates a new MonitoringMeasurementCriteria object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitoringMeasurementCriteriaWithDefaults() *MonitoringMeasurementCriteria {
	this := MonitoringMeasurementCriteria{}
	var function string = "none_function"
	this.Function = &function
	return &this
}

// GetFunction returns the Function field value if set, zero value otherwise.
func (o *MonitoringMeasurementCriteria) GetFunction() string {
	if o == nil || o.Function == nil {
		var ret string
		return ret
	}
	return *o.Function
}

// GetFunctionOk returns a tuple with the Function field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringMeasurementCriteria) GetFunctionOk() (*string, bool) {
	if o == nil || o.Function == nil {
		return nil, false
	}
	return o.Function, true
}

// HasFunction returns a boolean if a field has been set.
func (o *MonitoringMeasurementCriteria) HasFunction() bool {
	if o != nil && o.Function != nil {
		return true
	}

	return false
}

// SetFunction gets a reference to the given string and assigns it to the Function field.
func (o *MonitoringMeasurementCriteria) SetFunction(v string) {
	o.Function = &v
}

// GetWindow returns the Window field value if set, zero value otherwise.
func (o *MonitoringMeasurementCriteria) GetWindow() string {
	if o == nil || o.Window == nil {
		var ret string
		return ret
	}
	return *o.Window
}

// GetWindowOk returns a tuple with the Window field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringMeasurementCriteria) GetWindowOk() (*string, bool) {
	if o == nil || o.Window == nil {
		return nil, false
	}
	return o.Window, true
}

// HasWindow returns a boolean if a field has been set.
func (o *MonitoringMeasurementCriteria) HasWindow() bool {
	if o != nil && o.Window != nil {
		return true
	}

	return false
}

// SetWindow gets a reference to the given string and assigns it to the Window field.
func (o *MonitoringMeasurementCriteria) SetWindow(v string) {
	o.Window = &v
}

func (o MonitoringMeasurementCriteria) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Function != nil {
		toSerialize["function"] = o.Function
	}
	if o.Window != nil {
		toSerialize["window"] = o.Window
	}
	return json.Marshal(toSerialize)
}

type NullableMonitoringMeasurementCriteria struct {
	value *MonitoringMeasurementCriteria
	isSet bool
}

func (v NullableMonitoringMeasurementCriteria) Get() *MonitoringMeasurementCriteria {
	return v.value
}

func (v *NullableMonitoringMeasurementCriteria) Set(val *MonitoringMeasurementCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringMeasurementCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringMeasurementCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringMeasurementCriteria(val *MonitoringMeasurementCriteria) *NullableMonitoringMeasurementCriteria {
	return &NullableMonitoringMeasurementCriteria{value: val, isSet: true}
}

func (v NullableMonitoringMeasurementCriteria) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringMeasurementCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


