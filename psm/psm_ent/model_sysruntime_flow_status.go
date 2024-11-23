/*
 * Sysruntime API reference
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

// SysruntimeFlowStatus FlowStatus captures the operational status of flows.
type SysruntimeFlowStatus struct {
	ConnectionTracking *SysruntimeConnTrackInfo `json:"connection-tracking,omitempty"`
	FlowDirection *string `json:"flow-direction,omitempty"`
	FlowInstance *string `json:"flow-instance,omitempty"`
}

// NewSysruntimeFlowStatus instantiates a new SysruntimeFlowStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSysruntimeFlowStatus() *SysruntimeFlowStatus {
	this := SysruntimeFlowStatus{}
	return &this
}

// NewSysruntimeFlowStatusWithDefaults instantiates a new SysruntimeFlowStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSysruntimeFlowStatusWithDefaults() *SysruntimeFlowStatus {
	this := SysruntimeFlowStatus{}
	return &this
}

// GetConnectionTracking returns the ConnectionTracking field value if set, zero value otherwise.
func (o *SysruntimeFlowStatus) GetConnectionTracking() SysruntimeConnTrackInfo {
	if o == nil || o.ConnectionTracking == nil {
		var ret SysruntimeConnTrackInfo
		return ret
	}
	return *o.ConnectionTracking
}

// GetConnectionTrackingOk returns a tuple with the ConnectionTracking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SysruntimeFlowStatus) GetConnectionTrackingOk() (*SysruntimeConnTrackInfo, bool) {
	if o == nil || o.ConnectionTracking == nil {
		return nil, false
	}
	return o.ConnectionTracking, true
}

// HasConnectionTracking returns a boolean if a field has been set.
func (o *SysruntimeFlowStatus) HasConnectionTracking() bool {
	if o != nil && o.ConnectionTracking != nil {
		return true
	}

	return false
}

// SetConnectionTracking gets a reference to the given SysruntimeConnTrackInfo and assigns it to the ConnectionTracking field.
func (o *SysruntimeFlowStatus) SetConnectionTracking(v SysruntimeConnTrackInfo) {
	o.ConnectionTracking = &v
}

// GetFlowDirection returns the FlowDirection field value if set, zero value otherwise.
func (o *SysruntimeFlowStatus) GetFlowDirection() string {
	if o == nil || o.FlowDirection == nil {
		var ret string
		return ret
	}
	return *o.FlowDirection
}

// GetFlowDirectionOk returns a tuple with the FlowDirection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SysruntimeFlowStatus) GetFlowDirectionOk() (*string, bool) {
	if o == nil || o.FlowDirection == nil {
		return nil, false
	}
	return o.FlowDirection, true
}

// HasFlowDirection returns a boolean if a field has been set.
func (o *SysruntimeFlowStatus) HasFlowDirection() bool {
	if o != nil && o.FlowDirection != nil {
		return true
	}

	return false
}

// SetFlowDirection gets a reference to the given string and assigns it to the FlowDirection field.
func (o *SysruntimeFlowStatus) SetFlowDirection(v string) {
	o.FlowDirection = &v
}

// GetFlowInstance returns the FlowInstance field value if set, zero value otherwise.
func (o *SysruntimeFlowStatus) GetFlowInstance() string {
	if o == nil || o.FlowInstance == nil {
		var ret string
		return ret
	}
	return *o.FlowInstance
}

// GetFlowInstanceOk returns a tuple with the FlowInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SysruntimeFlowStatus) GetFlowInstanceOk() (*string, bool) {
	if o == nil || o.FlowInstance == nil {
		return nil, false
	}
	return o.FlowInstance, true
}

// HasFlowInstance returns a boolean if a field has been set.
func (o *SysruntimeFlowStatus) HasFlowInstance() bool {
	if o != nil && o.FlowInstance != nil {
		return true
	}

	return false
}

// SetFlowInstance gets a reference to the given string and assigns it to the FlowInstance field.
func (o *SysruntimeFlowStatus) SetFlowInstance(v string) {
	o.FlowInstance = &v
}

func (o SysruntimeFlowStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ConnectionTracking != nil {
		toSerialize["connection-tracking"] = o.ConnectionTracking
	}
	if o.FlowDirection != nil {
		toSerialize["flow-direction"] = o.FlowDirection
	}
	if o.FlowInstance != nil {
		toSerialize["flow-instance"] = o.FlowInstance
	}
	return json.Marshal(toSerialize)
}

type NullableSysruntimeFlowStatus struct {
	value *SysruntimeFlowStatus
	isSet bool
}

func (v NullableSysruntimeFlowStatus) Get() *SysruntimeFlowStatus {
	return v.value
}

func (v *NullableSysruntimeFlowStatus) Set(val *SysruntimeFlowStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSysruntimeFlowStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSysruntimeFlowStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSysruntimeFlowStatus(val *SysruntimeFlowStatus) *NullableSysruntimeFlowStatus {
	return &NullableSysruntimeFlowStatus{value: val, isSet: true}
}

func (v NullableSysruntimeFlowStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSysruntimeFlowStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


