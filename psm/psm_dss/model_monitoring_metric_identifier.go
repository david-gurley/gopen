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

// MonitoringMetricIdentifier MetricIdentifier - uniquely identifies an metric that needs to be monitored as part of the policy.
type MonitoringMetricIdentifier struct {
	// Field belonging to the kind e.g. ConnectionsPerSecond. This is the attribute that will be monitored and alerts will be created/resolved based on the thresholds.
	FieldName *string `json:"field-name,omitempty"`
	// Metric group - e.g. ftestats, flowstats, etc.
	Group *string `json:"group,omitempty"`
	// Sub-category within the group e.g. MaxSessionThresholdDrops, FlowMissPackets.
	Kind *string `json:"kind,omitempty"`
}

// NewMonitoringMetricIdentifier instantiates a new MonitoringMetricIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitoringMetricIdentifier() *MonitoringMetricIdentifier {
	this := MonitoringMetricIdentifier{}
	return &this
}

// NewMonitoringMetricIdentifierWithDefaults instantiates a new MonitoringMetricIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitoringMetricIdentifierWithDefaults() *MonitoringMetricIdentifier {
	this := MonitoringMetricIdentifier{}
	return &this
}

// GetFieldName returns the FieldName field value if set, zero value otherwise.
func (o *MonitoringMetricIdentifier) GetFieldName() string {
	if o == nil || o.FieldName == nil {
		var ret string
		return ret
	}
	return *o.FieldName
}

// GetFieldNameOk returns a tuple with the FieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringMetricIdentifier) GetFieldNameOk() (*string, bool) {
	if o == nil || o.FieldName == nil {
		return nil, false
	}
	return o.FieldName, true
}

// HasFieldName returns a boolean if a field has been set.
func (o *MonitoringMetricIdentifier) HasFieldName() bool {
	if o != nil && o.FieldName != nil {
		return true
	}

	return false
}

// SetFieldName gets a reference to the given string and assigns it to the FieldName field.
func (o *MonitoringMetricIdentifier) SetFieldName(v string) {
	o.FieldName = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *MonitoringMetricIdentifier) GetGroup() string {
	if o == nil || o.Group == nil {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringMetricIdentifier) GetGroupOk() (*string, bool) {
	if o == nil || o.Group == nil {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *MonitoringMetricIdentifier) HasGroup() bool {
	if o != nil && o.Group != nil {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *MonitoringMetricIdentifier) SetGroup(v string) {
	o.Group = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *MonitoringMetricIdentifier) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringMetricIdentifier) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *MonitoringMetricIdentifier) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *MonitoringMetricIdentifier) SetKind(v string) {
	o.Kind = &v
}

func (o MonitoringMetricIdentifier) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FieldName != nil {
		toSerialize["field-name"] = o.FieldName
	}
	if o.Group != nil {
		toSerialize["group"] = o.Group
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	return json.Marshal(toSerialize)
}

type NullableMonitoringMetricIdentifier struct {
	value *MonitoringMetricIdentifier
	isSet bool
}

func (v NullableMonitoringMetricIdentifier) Get() *MonitoringMetricIdentifier {
	return v.value
}

func (v *NullableMonitoringMetricIdentifier) Set(val *MonitoringMetricIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringMetricIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringMetricIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringMetricIdentifier(val *MonitoringMetricIdentifier) *NullableMonitoringMetricIdentifier {
	return &NullableMonitoringMetricIdentifier{value: val, isSet: true}
}

func (v NullableMonitoringMetricIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringMetricIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

