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

// TelemetryQueryResultSeries struct for TelemetryQueryResultSeries
type TelemetryQueryResultSeries struct {
	// columns list all available fields in tsdb.
	Columns *[]string `json:"columns,omitempty"`
	// Name of the series.
	Name *string `json:"name,omitempty"`
	// Tags are the TSDB tags in the query response.
	Tags *map[string]string `json:"tags,omitempty"`
	// values contain field values received frpm tsdb, it is in the form of [][]interface{}.
	Values *[]ApiInterfaceSlice `json:"values,omitempty"`
}

// NewTelemetryQueryResultSeries instantiates a new TelemetryQueryResultSeries object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryQueryResultSeries() *TelemetryQueryResultSeries {
	this := TelemetryQueryResultSeries{}
	return &this
}

// NewTelemetryQueryResultSeriesWithDefaults instantiates a new TelemetryQueryResultSeries object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryQueryResultSeriesWithDefaults() *TelemetryQueryResultSeries {
	this := TelemetryQueryResultSeries{}
	return &this
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *TelemetryQueryResultSeries) GetColumns() []string {
	if o == nil || o.Columns == nil {
		var ret []string
		return ret
	}
	return *o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryQueryResultSeries) GetColumnsOk() (*[]string, bool) {
	if o == nil || o.Columns == nil {
		return nil, false
	}
	return o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *TelemetryQueryResultSeries) HasColumns() bool {
	if o != nil && o.Columns != nil {
		return true
	}

	return false
}

// SetColumns gets a reference to the given []string and assigns it to the Columns field.
func (o *TelemetryQueryResultSeries) SetColumns(v []string) {
	o.Columns = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TelemetryQueryResultSeries) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryQueryResultSeries) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TelemetryQueryResultSeries) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TelemetryQueryResultSeries) SetName(v string) {
	o.Name = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *TelemetryQueryResultSeries) GetTags() map[string]string {
	if o == nil || o.Tags == nil {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryQueryResultSeries) GetTagsOk() (*map[string]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *TelemetryQueryResultSeries) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *TelemetryQueryResultSeries) SetTags(v map[string]string) {
	o.Tags = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *TelemetryQueryResultSeries) GetValues() []ApiInterfaceSlice {
	if o == nil || o.Values == nil {
		var ret []ApiInterfaceSlice
		return ret
	}
	return *o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryQueryResultSeries) GetValuesOk() (*[]ApiInterfaceSlice, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *TelemetryQueryResultSeries) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given []ApiInterfaceSlice and assigns it to the Values field.
func (o *TelemetryQueryResultSeries) SetValues(v []ApiInterfaceSlice) {
	o.Values = &v
}

func (o TelemetryQueryResultSeries) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Columns != nil {
		toSerialize["columns"] = o.Columns
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}
	return json.Marshal(toSerialize)
}

type NullableTelemetryQueryResultSeries struct {
	value *TelemetryQueryResultSeries
	isSet bool
}

func (v NullableTelemetryQueryResultSeries) Get() *TelemetryQueryResultSeries {
	return v.value
}

func (v *NullableTelemetryQueryResultSeries) Set(val *TelemetryQueryResultSeries) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryQueryResultSeries) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryQueryResultSeries) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryQueryResultSeries(val *TelemetryQueryResultSeries) *NullableTelemetryQueryResultSeries {
	return &NullableTelemetryQueryResultSeries{value: val, isSet: true}
}

func (v NullableTelemetryQueryResultSeries) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryQueryResultSeries) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


