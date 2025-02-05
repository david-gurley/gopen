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

// MonitoringFlowExportPolicyStatus FlowExportPolicyStatus.
type MonitoringFlowExportPolicyStatus struct {
	PropagationStatus *MonitoringPropagationStatus `json:"propagation-status,omitempty"`
}

// NewMonitoringFlowExportPolicyStatus instantiates a new MonitoringFlowExportPolicyStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitoringFlowExportPolicyStatus() *MonitoringFlowExportPolicyStatus {
	this := MonitoringFlowExportPolicyStatus{}
	return &this
}

// NewMonitoringFlowExportPolicyStatusWithDefaults instantiates a new MonitoringFlowExportPolicyStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitoringFlowExportPolicyStatusWithDefaults() *MonitoringFlowExportPolicyStatus {
	this := MonitoringFlowExportPolicyStatus{}
	return &this
}

// GetPropagationStatus returns the PropagationStatus field value if set, zero value otherwise.
func (o *MonitoringFlowExportPolicyStatus) GetPropagationStatus() MonitoringPropagationStatus {
	if o == nil || o.PropagationStatus == nil {
		var ret MonitoringPropagationStatus
		return ret
	}
	return *o.PropagationStatus
}

// GetPropagationStatusOk returns a tuple with the PropagationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringFlowExportPolicyStatus) GetPropagationStatusOk() (*MonitoringPropagationStatus, bool) {
	if o == nil || o.PropagationStatus == nil {
		return nil, false
	}
	return o.PropagationStatus, true
}

// HasPropagationStatus returns a boolean if a field has been set.
func (o *MonitoringFlowExportPolicyStatus) HasPropagationStatus() bool {
	if o != nil && o.PropagationStatus != nil {
		return true
	}

	return false
}

// SetPropagationStatus gets a reference to the given MonitoringPropagationStatus and assigns it to the PropagationStatus field.
func (o *MonitoringFlowExportPolicyStatus) SetPropagationStatus(v MonitoringPropagationStatus) {
	o.PropagationStatus = &v
}

func (o MonitoringFlowExportPolicyStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PropagationStatus != nil {
		toSerialize["propagation-status"] = o.PropagationStatus
	}
	return json.Marshal(toSerialize)
}

type NullableMonitoringFlowExportPolicyStatus struct {
	value *MonitoringFlowExportPolicyStatus
	isSet bool
}

func (v NullableMonitoringFlowExportPolicyStatus) Get() *MonitoringFlowExportPolicyStatus {
	return v.value
}

func (v *NullableMonitoringFlowExportPolicyStatus) Set(val *MonitoringFlowExportPolicyStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringFlowExportPolicyStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringFlowExportPolicyStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringFlowExportPolicyStatus(val *MonitoringFlowExportPolicyStatus) *NullableMonitoringFlowExportPolicyStatus {
	return &NullableMonitoringFlowExportPolicyStatus{value: val, isSet: true}
}

func (v NullableMonitoringFlowExportPolicyStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringFlowExportPolicyStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


