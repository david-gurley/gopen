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
	"time"
)

// MonitoringMirrorSessionStatus MirrorSessionStatus.
type MonitoringMirrorSessionStatus struct {
	PropagationStatus *MonitoringPropagationStatus `json:"propagation-status,omitempty"`
	ScheduleState *string `json:"schedule-state,omitempty"`
	StartedAt *time.Time `json:"started-at,omitempty"`
}

// NewMonitoringMirrorSessionStatus instantiates a new MonitoringMirrorSessionStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitoringMirrorSessionStatus() *MonitoringMirrorSessionStatus {
	this := MonitoringMirrorSessionStatus{}
	var scheduleState string = "none"
	this.ScheduleState = &scheduleState
	return &this
}

// NewMonitoringMirrorSessionStatusWithDefaults instantiates a new MonitoringMirrorSessionStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitoringMirrorSessionStatusWithDefaults() *MonitoringMirrorSessionStatus {
	this := MonitoringMirrorSessionStatus{}
	var scheduleState string = "none"
	this.ScheduleState = &scheduleState
	return &this
}

// GetPropagationStatus returns the PropagationStatus field value if set, zero value otherwise.
func (o *MonitoringMirrorSessionStatus) GetPropagationStatus() MonitoringPropagationStatus {
	if o == nil || o.PropagationStatus == nil {
		var ret MonitoringPropagationStatus
		return ret
	}
	return *o.PropagationStatus
}

// GetPropagationStatusOk returns a tuple with the PropagationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringMirrorSessionStatus) GetPropagationStatusOk() (*MonitoringPropagationStatus, bool) {
	if o == nil || o.PropagationStatus == nil {
		return nil, false
	}
	return o.PropagationStatus, true
}

// HasPropagationStatus returns a boolean if a field has been set.
func (o *MonitoringMirrorSessionStatus) HasPropagationStatus() bool {
	if o != nil && o.PropagationStatus != nil {
		return true
	}

	return false
}

// SetPropagationStatus gets a reference to the given MonitoringPropagationStatus and assigns it to the PropagationStatus field.
func (o *MonitoringMirrorSessionStatus) SetPropagationStatus(v MonitoringPropagationStatus) {
	o.PropagationStatus = &v
}

// GetScheduleState returns the ScheduleState field value if set, zero value otherwise.
func (o *MonitoringMirrorSessionStatus) GetScheduleState() string {
	if o == nil || o.ScheduleState == nil {
		var ret string
		return ret
	}
	return *o.ScheduleState
}

// GetScheduleStateOk returns a tuple with the ScheduleState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringMirrorSessionStatus) GetScheduleStateOk() (*string, bool) {
	if o == nil || o.ScheduleState == nil {
		return nil, false
	}
	return o.ScheduleState, true
}

// HasScheduleState returns a boolean if a field has been set.
func (o *MonitoringMirrorSessionStatus) HasScheduleState() bool {
	if o != nil && o.ScheduleState != nil {
		return true
	}

	return false
}

// SetScheduleState gets a reference to the given string and assigns it to the ScheduleState field.
func (o *MonitoringMirrorSessionStatus) SetScheduleState(v string) {
	o.ScheduleState = &v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *MonitoringMirrorSessionStatus) GetStartedAt() time.Time {
	if o == nil || o.StartedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringMirrorSessionStatus) GetStartedAtOk() (*time.Time, bool) {
	if o == nil || o.StartedAt == nil {
		return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *MonitoringMirrorSessionStatus) HasStartedAt() bool {
	if o != nil && o.StartedAt != nil {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given time.Time and assigns it to the StartedAt field.
func (o *MonitoringMirrorSessionStatus) SetStartedAt(v time.Time) {
	o.StartedAt = &v
}

func (o MonitoringMirrorSessionStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PropagationStatus != nil {
		toSerialize["propagation-status"] = o.PropagationStatus
	}
	if o.ScheduleState != nil {
		toSerialize["schedule-state"] = o.ScheduleState
	}
	if o.StartedAt != nil {
		toSerialize["started-at"] = o.StartedAt
	}
	return json.Marshal(toSerialize)
}

type NullableMonitoringMirrorSessionStatus struct {
	value *MonitoringMirrorSessionStatus
	isSet bool
}

func (v NullableMonitoringMirrorSessionStatus) Get() *MonitoringMirrorSessionStatus {
	return v.value
}

func (v *NullableMonitoringMirrorSessionStatus) Set(val *MonitoringMirrorSessionStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringMirrorSessionStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringMirrorSessionStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringMirrorSessionStatus(val *MonitoringMirrorSessionStatus) *NullableMonitoringMirrorSessionStatus {
	return &NullableMonitoringMirrorSessionStatus{value: val, isSet: true}
}

func (v NullableMonitoringMirrorSessionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringMirrorSessionStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


