/*
 * Rollout API reference
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

// RolloutRolloutActionStatus RolloutAction Status gives the status of the rollout at the top level as well as details of individual components.
type RolloutRolloutActionStatus struct {
	// Heuristic value of percentage completion of the rollout.
	CompletionPercent *int64 `json:"completion-percent,omitempty"`
	// End time of Rollout.
	EndTime *time.Time `json:"end-time,omitempty"`
	// Version of the cluster before the start of rollout.
	PrevVersion *string `json:"prev-version,omitempty"`
	// Start time of Rollout.
	StartTime *time.Time `json:"start-time,omitempty"`
	State *string `json:"state,omitempty"`
}

// NewRolloutRolloutActionStatus instantiates a new RolloutRolloutActionStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRolloutRolloutActionStatus() *RolloutRolloutActionStatus {
	this := RolloutRolloutActionStatus{}
	var state string = "progressing"
	this.State = &state
	return &this
}

// NewRolloutRolloutActionStatusWithDefaults instantiates a new RolloutRolloutActionStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRolloutRolloutActionStatusWithDefaults() *RolloutRolloutActionStatus {
	this := RolloutRolloutActionStatus{}
	var state string = "progressing"
	this.State = &state
	return &this
}

// GetCompletionPercent returns the CompletionPercent field value if set, zero value otherwise.
func (o *RolloutRolloutActionStatus) GetCompletionPercent() int64 {
	if o == nil || o.CompletionPercent == nil {
		var ret int64
		return ret
	}
	return *o.CompletionPercent
}

// GetCompletionPercentOk returns a tuple with the CompletionPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RolloutRolloutActionStatus) GetCompletionPercentOk() (*int64, bool) {
	if o == nil || o.CompletionPercent == nil {
		return nil, false
	}
	return o.CompletionPercent, true
}

// HasCompletionPercent returns a boolean if a field has been set.
func (o *RolloutRolloutActionStatus) HasCompletionPercent() bool {
	if o != nil && o.CompletionPercent != nil {
		return true
	}

	return false
}

// SetCompletionPercent gets a reference to the given int64 and assigns it to the CompletionPercent field.
func (o *RolloutRolloutActionStatus) SetCompletionPercent(v int64) {
	o.CompletionPercent = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *RolloutRolloutActionStatus) GetEndTime() time.Time {
	if o == nil || o.EndTime == nil {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RolloutRolloutActionStatus) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || o.EndTime == nil {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *RolloutRolloutActionStatus) HasEndTime() bool {
	if o != nil && o.EndTime != nil {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *RolloutRolloutActionStatus) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetPrevVersion returns the PrevVersion field value if set, zero value otherwise.
func (o *RolloutRolloutActionStatus) GetPrevVersion() string {
	if o == nil || o.PrevVersion == nil {
		var ret string
		return ret
	}
	return *o.PrevVersion
}

// GetPrevVersionOk returns a tuple with the PrevVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RolloutRolloutActionStatus) GetPrevVersionOk() (*string, bool) {
	if o == nil || o.PrevVersion == nil {
		return nil, false
	}
	return o.PrevVersion, true
}

// HasPrevVersion returns a boolean if a field has been set.
func (o *RolloutRolloutActionStatus) HasPrevVersion() bool {
	if o != nil && o.PrevVersion != nil {
		return true
	}

	return false
}

// SetPrevVersion gets a reference to the given string and assigns it to the PrevVersion field.
func (o *RolloutRolloutActionStatus) SetPrevVersion(v string) {
	o.PrevVersion = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *RolloutRolloutActionStatus) GetStartTime() time.Time {
	if o == nil || o.StartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RolloutRolloutActionStatus) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *RolloutRolloutActionStatus) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *RolloutRolloutActionStatus) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *RolloutRolloutActionStatus) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RolloutRolloutActionStatus) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *RolloutRolloutActionStatus) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *RolloutRolloutActionStatus) SetState(v string) {
	o.State = &v
}

func (o RolloutRolloutActionStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CompletionPercent != nil {
		toSerialize["completion-percent"] = o.CompletionPercent
	}
	if o.EndTime != nil {
		toSerialize["end-time"] = o.EndTime
	}
	if o.PrevVersion != nil {
		toSerialize["prev-version"] = o.PrevVersion
	}
	if o.StartTime != nil {
		toSerialize["start-time"] = o.StartTime
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableRolloutRolloutActionStatus struct {
	value *RolloutRolloutActionStatus
	isSet bool
}

func (v NullableRolloutRolloutActionStatus) Get() *RolloutRolloutActionStatus {
	return v.value
}

func (v *NullableRolloutRolloutActionStatus) Set(val *RolloutRolloutActionStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableRolloutRolloutActionStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableRolloutRolloutActionStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRolloutRolloutActionStatus(val *RolloutRolloutActionStatus) *NullableRolloutRolloutActionStatus {
	return &NullableRolloutRolloutActionStatus{value: val, isSet: true}
}

func (v NullableRolloutRolloutActionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRolloutRolloutActionStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

