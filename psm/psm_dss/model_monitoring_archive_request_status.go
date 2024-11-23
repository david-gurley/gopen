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

// MonitoringArchiveRequestStatus struct for MonitoringArchiveRequestStatus
type MonitoringArchiveRequestStatus struct {
	Reason *string `json:"reason,omitempty"`
	Status *string `json:"status,omitempty"`
	Uri *string `json:"uri,omitempty"`
}

// NewMonitoringArchiveRequestStatus instantiates a new MonitoringArchiveRequestStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitoringArchiveRequestStatus() *MonitoringArchiveRequestStatus {
	this := MonitoringArchiveRequestStatus{}
	var status string = "scheduled"
	this.Status = &status
	return &this
}

// NewMonitoringArchiveRequestStatusWithDefaults instantiates a new MonitoringArchiveRequestStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitoringArchiveRequestStatusWithDefaults() *MonitoringArchiveRequestStatus {
	this := MonitoringArchiveRequestStatus{}
	var status string = "scheduled"
	this.Status = &status
	return &this
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *MonitoringArchiveRequestStatus) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringArchiveRequestStatus) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *MonitoringArchiveRequestStatus) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *MonitoringArchiveRequestStatus) SetReason(v string) {
	o.Reason = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MonitoringArchiveRequestStatus) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringArchiveRequestStatus) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MonitoringArchiveRequestStatus) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *MonitoringArchiveRequestStatus) SetStatus(v string) {
	o.Status = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *MonitoringArchiveRequestStatus) GetUri() string {
	if o == nil || o.Uri == nil {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringArchiveRequestStatus) GetUriOk() (*string, bool) {
	if o == nil || o.Uri == nil {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *MonitoringArchiveRequestStatus) HasUri() bool {
	if o != nil && o.Uri != nil {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *MonitoringArchiveRequestStatus) SetUri(v string) {
	o.Uri = &v
}

func (o MonitoringArchiveRequestStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Uri != nil {
		toSerialize["uri"] = o.Uri
	}
	return json.Marshal(toSerialize)
}

type NullableMonitoringArchiveRequestStatus struct {
	value *MonitoringArchiveRequestStatus
	isSet bool
}

func (v NullableMonitoringArchiveRequestStatus) Get() *MonitoringArchiveRequestStatus {
	return v.value
}

func (v *NullableMonitoringArchiveRequestStatus) Set(val *MonitoringArchiveRequestStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringArchiveRequestStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringArchiveRequestStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringArchiveRequestStatus(val *MonitoringArchiveRequestStatus) *NullableMonitoringArchiveRequestStatus {
	return &NullableMonitoringArchiveRequestStatus{value: val, isSet: true}
}

func (v NullableMonitoringArchiveRequestStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringArchiveRequestStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


