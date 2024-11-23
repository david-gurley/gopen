/*
 * Telemetry_query API reference
 *
 * Service name  
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package psm_dss

import (
	"encoding/json"
)

// TelemetryQueryTopSpec TopSpec is optional parameters for top aggregation function.
type TelemetryQueryTopSpec struct {
	// TopN defines how many entries returned for top aggregation function and by default it is 10.
	TopN *int64 `json:"top-n,omitempty"`
}

// NewTelemetryQueryTopSpec instantiates a new TelemetryQueryTopSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryQueryTopSpec() *TelemetryQueryTopSpec {
	this := TelemetryQueryTopSpec{}
	return &this
}

// NewTelemetryQueryTopSpecWithDefaults instantiates a new TelemetryQueryTopSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryQueryTopSpecWithDefaults() *TelemetryQueryTopSpec {
	this := TelemetryQueryTopSpec{}
	return &this
}

// GetTopN returns the TopN field value if set, zero value otherwise.
func (o *TelemetryQueryTopSpec) GetTopN() int64 {
	if o == nil || o.TopN == nil {
		var ret int64
		return ret
	}
	return *o.TopN
}

// GetTopNOk returns a tuple with the TopN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryQueryTopSpec) GetTopNOk() (*int64, bool) {
	if o == nil || o.TopN == nil {
		return nil, false
	}
	return o.TopN, true
}

// HasTopN returns a boolean if a field has been set.
func (o *TelemetryQueryTopSpec) HasTopN() bool {
	if o != nil && o.TopN != nil {
		return true
	}

	return false
}

// SetTopN gets a reference to the given int64 and assigns it to the TopN field.
func (o *TelemetryQueryTopSpec) SetTopN(v int64) {
	o.TopN = &v
}

func (o TelemetryQueryTopSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TopN != nil {
		toSerialize["top-n"] = o.TopN
	}
	return json.Marshal(toSerialize)
}

type NullableTelemetryQueryTopSpec struct {
	value *TelemetryQueryTopSpec
	isSet bool
}

func (v NullableTelemetryQueryTopSpec) Get() *TelemetryQueryTopSpec {
	return v.value
}

func (v *NullableTelemetryQueryTopSpec) Set(val *TelemetryQueryTopSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryQueryTopSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryQueryTopSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryQueryTopSpec(val *TelemetryQueryTopSpec) *NullableTelemetryQueryTopSpec {
	return &NullableTelemetryQueryTopSpec{value: val, isSet: true}
}

func (v NullableTelemetryQueryTopSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryQueryTopSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

