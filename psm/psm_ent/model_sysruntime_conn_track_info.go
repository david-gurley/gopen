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
	"time"
)

// SysruntimeConnTrackInfo ConnTrackInfo contains dynamic information that was recorded due to connection tracking.
type SysruntimeConnTrackInfo struct {
	CreateTimestamp *time.Time `json:"create-timestamp,omitempty"`
	FlowBytes *string `json:"flow-bytes,omitempty"`
	FlowPackets *int64 `json:"flow-packets,omitempty"`
}

// NewSysruntimeConnTrackInfo instantiates a new SysruntimeConnTrackInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSysruntimeConnTrackInfo() *SysruntimeConnTrackInfo {
	this := SysruntimeConnTrackInfo{}
	return &this
}

// NewSysruntimeConnTrackInfoWithDefaults instantiates a new SysruntimeConnTrackInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSysruntimeConnTrackInfoWithDefaults() *SysruntimeConnTrackInfo {
	this := SysruntimeConnTrackInfo{}
	return &this
}

// GetCreateTimestamp returns the CreateTimestamp field value if set, zero value otherwise.
func (o *SysruntimeConnTrackInfo) GetCreateTimestamp() time.Time {
	if o == nil || o.CreateTimestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.CreateTimestamp
}

// GetCreateTimestampOk returns a tuple with the CreateTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SysruntimeConnTrackInfo) GetCreateTimestampOk() (*time.Time, bool) {
	if o == nil || o.CreateTimestamp == nil {
		return nil, false
	}
	return o.CreateTimestamp, true
}

// HasCreateTimestamp returns a boolean if a field has been set.
func (o *SysruntimeConnTrackInfo) HasCreateTimestamp() bool {
	if o != nil && o.CreateTimestamp != nil {
		return true
	}

	return false
}

// SetCreateTimestamp gets a reference to the given time.Time and assigns it to the CreateTimestamp field.
func (o *SysruntimeConnTrackInfo) SetCreateTimestamp(v time.Time) {
	o.CreateTimestamp = &v
}

// GetFlowBytes returns the FlowBytes field value if set, zero value otherwise.
func (o *SysruntimeConnTrackInfo) GetFlowBytes() string {
	if o == nil || o.FlowBytes == nil {
		var ret string
		return ret
	}
	return *o.FlowBytes
}

// GetFlowBytesOk returns a tuple with the FlowBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SysruntimeConnTrackInfo) GetFlowBytesOk() (*string, bool) {
	if o == nil || o.FlowBytes == nil {
		return nil, false
	}
	return o.FlowBytes, true
}

// HasFlowBytes returns a boolean if a field has been set.
func (o *SysruntimeConnTrackInfo) HasFlowBytes() bool {
	if o != nil && o.FlowBytes != nil {
		return true
	}

	return false
}

// SetFlowBytes gets a reference to the given string and assigns it to the FlowBytes field.
func (o *SysruntimeConnTrackInfo) SetFlowBytes(v string) {
	o.FlowBytes = &v
}

// GetFlowPackets returns the FlowPackets field value if set, zero value otherwise.
func (o *SysruntimeConnTrackInfo) GetFlowPackets() int64 {
	if o == nil || o.FlowPackets == nil {
		var ret int64
		return ret
	}
	return *o.FlowPackets
}

// GetFlowPacketsOk returns a tuple with the FlowPackets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SysruntimeConnTrackInfo) GetFlowPacketsOk() (*int64, bool) {
	if o == nil || o.FlowPackets == nil {
		return nil, false
	}
	return o.FlowPackets, true
}

// HasFlowPackets returns a boolean if a field has been set.
func (o *SysruntimeConnTrackInfo) HasFlowPackets() bool {
	if o != nil && o.FlowPackets != nil {
		return true
	}

	return false
}

// SetFlowPackets gets a reference to the given int64 and assigns it to the FlowPackets field.
func (o *SysruntimeConnTrackInfo) SetFlowPackets(v int64) {
	o.FlowPackets = &v
}

func (o SysruntimeConnTrackInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreateTimestamp != nil {
		toSerialize["create-timestamp"] = o.CreateTimestamp
	}
	if o.FlowBytes != nil {
		toSerialize["flow-bytes"] = o.FlowBytes
	}
	if o.FlowPackets != nil {
		toSerialize["flow-packets"] = o.FlowPackets
	}
	return json.Marshal(toSerialize)
}

type NullableSysruntimeConnTrackInfo struct {
	value *SysruntimeConnTrackInfo
	isSet bool
}

func (v NullableSysruntimeConnTrackInfo) Get() *SysruntimeConnTrackInfo {
	return v.value
}

func (v *NullableSysruntimeConnTrackInfo) Set(val *SysruntimeConnTrackInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSysruntimeConnTrackInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSysruntimeConnTrackInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSysruntimeConnTrackInfo(val *SysruntimeConnTrackInfo) *NullableSysruntimeConnTrackInfo {
	return &NullableSysruntimeConnTrackInfo{value: val, isSet: true}
}

func (v NullableSysruntimeConnTrackInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSysruntimeConnTrackInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


