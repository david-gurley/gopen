/*
 * Staging API reference
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

// StagingClearActionStatus struct for StagingClearActionStatus
type StagingClearActionStatus struct {
	Reason *string `json:"reason,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewStagingClearActionStatus instantiates a new StagingClearActionStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStagingClearActionStatus() *StagingClearActionStatus {
	this := StagingClearActionStatus{}
	var status string = "success"
	this.Status = &status
	return &this
}

// NewStagingClearActionStatusWithDefaults instantiates a new StagingClearActionStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStagingClearActionStatusWithDefaults() *StagingClearActionStatus {
	this := StagingClearActionStatus{}
	var status string = "success"
	this.Status = &status
	return &this
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *StagingClearActionStatus) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagingClearActionStatus) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *StagingClearActionStatus) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *StagingClearActionStatus) SetReason(v string) {
	o.Reason = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *StagingClearActionStatus) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagingClearActionStatus) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *StagingClearActionStatus) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *StagingClearActionStatus) SetStatus(v string) {
	o.Status = &v
}

func (o StagingClearActionStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableStagingClearActionStatus struct {
	value *StagingClearActionStatus
	isSet bool
}

func (v NullableStagingClearActionStatus) Get() *StagingClearActionStatus {
	return v.value
}

func (v *NullableStagingClearActionStatus) Set(val *StagingClearActionStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableStagingClearActionStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableStagingClearActionStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStagingClearActionStatus(val *StagingClearActionStatus) *NullableStagingClearActionStatus {
	return &NullableStagingClearActionStatus{value: val, isSet: true}
}

func (v NullableStagingClearActionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStagingClearActionStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


