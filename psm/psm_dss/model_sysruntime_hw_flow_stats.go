/*
 * Sysruntime API reference
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

// SysruntimeHWFlowStats HWFlowStats captures the stats of the flow.
type SysruntimeHWFlowStats struct {
	FlowDroppedBytes *string `json:"flow-dropped-bytes,omitempty"`
	FlowDroppedPackets *string `json:"flow-dropped-packets,omitempty"`
	FlowPermittedBytes *string `json:"flow-permitted-bytes,omitempty"`
	FlowPermittedPackets *string `json:"flow-permitted-packets,omitempty"`
	NumTcpRstSent *string `json:"num-tcp-rst-sent,omitempty"`
	NumTcpTicklesSent *string `json:"num-tcp-tickles-sent,omitempty"`
}

// NewSysruntimeHWFlowStats instantiates a new SysruntimeHWFlowStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSysruntimeHWFlowStats() *SysruntimeHWFlowStats {
	this := SysruntimeHWFlowStats{}
	return &this
}

// NewSysruntimeHWFlowStatsWithDefaults instantiates a new SysruntimeHWFlowStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSysruntimeHWFlowStatsWithDefaults() *SysruntimeHWFlowStats {
	this := SysruntimeHWFlowStats{}
	return &this
}

// GetFlowDroppedBytes returns the FlowDroppedBytes field value if set, zero value otherwise.
func (o *SysruntimeHWFlowStats) GetFlowDroppedBytes() string {
	if o == nil || o.FlowDroppedBytes == nil {
		var ret string
		return ret
	}
	return *o.FlowDroppedBytes
}

// GetFlowDroppedBytesOk returns a tuple with the FlowDroppedBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SysruntimeHWFlowStats) GetFlowDroppedBytesOk() (*string, bool) {
	if o == nil || o.FlowDroppedBytes == nil {
		return nil, false
	}
	return o.FlowDroppedBytes, true
}

// HasFlowDroppedBytes returns a boolean if a field has been set.
func (o *SysruntimeHWFlowStats) HasFlowDroppedBytes() bool {
	if o != nil && o.FlowDroppedBytes != nil {
		return true
	}

	return false
}

// SetFlowDroppedBytes gets a reference to the given string and assigns it to the FlowDroppedBytes field.
func (o *SysruntimeHWFlowStats) SetFlowDroppedBytes(v string) {
	o.FlowDroppedBytes = &v
}

// GetFlowDroppedPackets returns the FlowDroppedPackets field value if set, zero value otherwise.
func (o *SysruntimeHWFlowStats) GetFlowDroppedPackets() string {
	if o == nil || o.FlowDroppedPackets == nil {
		var ret string
		return ret
	}
	return *o.FlowDroppedPackets
}

// GetFlowDroppedPacketsOk returns a tuple with the FlowDroppedPackets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SysruntimeHWFlowStats) GetFlowDroppedPacketsOk() (*string, bool) {
	if o == nil || o.FlowDroppedPackets == nil {
		return nil, false
	}
	return o.FlowDroppedPackets, true
}

// HasFlowDroppedPackets returns a boolean if a field has been set.
func (o *SysruntimeHWFlowStats) HasFlowDroppedPackets() bool {
	if o != nil && o.FlowDroppedPackets != nil {
		return true
	}

	return false
}

// SetFlowDroppedPackets gets a reference to the given string and assigns it to the FlowDroppedPackets field.
func (o *SysruntimeHWFlowStats) SetFlowDroppedPackets(v string) {
	o.FlowDroppedPackets = &v
}

// GetFlowPermittedBytes returns the FlowPermittedBytes field value if set, zero value otherwise.
func (o *SysruntimeHWFlowStats) GetFlowPermittedBytes() string {
	if o == nil || o.FlowPermittedBytes == nil {
		var ret string
		return ret
	}
	return *o.FlowPermittedBytes
}

// GetFlowPermittedBytesOk returns a tuple with the FlowPermittedBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SysruntimeHWFlowStats) GetFlowPermittedBytesOk() (*string, bool) {
	if o == nil || o.FlowPermittedBytes == nil {
		return nil, false
	}
	return o.FlowPermittedBytes, true
}

// HasFlowPermittedBytes returns a boolean if a field has been set.
func (o *SysruntimeHWFlowStats) HasFlowPermittedBytes() bool {
	if o != nil && o.FlowPermittedBytes != nil {
		return true
	}

	return false
}

// SetFlowPermittedBytes gets a reference to the given string and assigns it to the FlowPermittedBytes field.
func (o *SysruntimeHWFlowStats) SetFlowPermittedBytes(v string) {
	o.FlowPermittedBytes = &v
}

// GetFlowPermittedPackets returns the FlowPermittedPackets field value if set, zero value otherwise.
func (o *SysruntimeHWFlowStats) GetFlowPermittedPackets() string {
	if o == nil || o.FlowPermittedPackets == nil {
		var ret string
		return ret
	}
	return *o.FlowPermittedPackets
}

// GetFlowPermittedPacketsOk returns a tuple with the FlowPermittedPackets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SysruntimeHWFlowStats) GetFlowPermittedPacketsOk() (*string, bool) {
	if o == nil || o.FlowPermittedPackets == nil {
		return nil, false
	}
	return o.FlowPermittedPackets, true
}

// HasFlowPermittedPackets returns a boolean if a field has been set.
func (o *SysruntimeHWFlowStats) HasFlowPermittedPackets() bool {
	if o != nil && o.FlowPermittedPackets != nil {
		return true
	}

	return false
}

// SetFlowPermittedPackets gets a reference to the given string and assigns it to the FlowPermittedPackets field.
func (o *SysruntimeHWFlowStats) SetFlowPermittedPackets(v string) {
	o.FlowPermittedPackets = &v
}

// GetNumTcpRstSent returns the NumTcpRstSent field value if set, zero value otherwise.
func (o *SysruntimeHWFlowStats) GetNumTcpRstSent() string {
	if o == nil || o.NumTcpRstSent == nil {
		var ret string
		return ret
	}
	return *o.NumTcpRstSent
}

// GetNumTcpRstSentOk returns a tuple with the NumTcpRstSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SysruntimeHWFlowStats) GetNumTcpRstSentOk() (*string, bool) {
	if o == nil || o.NumTcpRstSent == nil {
		return nil, false
	}
	return o.NumTcpRstSent, true
}

// HasNumTcpRstSent returns a boolean if a field has been set.
func (o *SysruntimeHWFlowStats) HasNumTcpRstSent() bool {
	if o != nil && o.NumTcpRstSent != nil {
		return true
	}

	return false
}

// SetNumTcpRstSent gets a reference to the given string and assigns it to the NumTcpRstSent field.
func (o *SysruntimeHWFlowStats) SetNumTcpRstSent(v string) {
	o.NumTcpRstSent = &v
}

// GetNumTcpTicklesSent returns the NumTcpTicklesSent field value if set, zero value otherwise.
func (o *SysruntimeHWFlowStats) GetNumTcpTicklesSent() string {
	if o == nil || o.NumTcpTicklesSent == nil {
		var ret string
		return ret
	}
	return *o.NumTcpTicklesSent
}

// GetNumTcpTicklesSentOk returns a tuple with the NumTcpTicklesSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SysruntimeHWFlowStats) GetNumTcpTicklesSentOk() (*string, bool) {
	if o == nil || o.NumTcpTicklesSent == nil {
		return nil, false
	}
	return o.NumTcpTicklesSent, true
}

// HasNumTcpTicklesSent returns a boolean if a field has been set.
func (o *SysruntimeHWFlowStats) HasNumTcpTicklesSent() bool {
	if o != nil && o.NumTcpTicklesSent != nil {
		return true
	}

	return false
}

// SetNumTcpTicklesSent gets a reference to the given string and assigns it to the NumTcpTicklesSent field.
func (o *SysruntimeHWFlowStats) SetNumTcpTicklesSent(v string) {
	o.NumTcpTicklesSent = &v
}

func (o SysruntimeHWFlowStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FlowDroppedBytes != nil {
		toSerialize["flow-dropped-bytes"] = o.FlowDroppedBytes
	}
	if o.FlowDroppedPackets != nil {
		toSerialize["flow-dropped-packets"] = o.FlowDroppedPackets
	}
	if o.FlowPermittedBytes != nil {
		toSerialize["flow-permitted-bytes"] = o.FlowPermittedBytes
	}
	if o.FlowPermittedPackets != nil {
		toSerialize["flow-permitted-packets"] = o.FlowPermittedPackets
	}
	if o.NumTcpRstSent != nil {
		toSerialize["num-tcp-rst-sent"] = o.NumTcpRstSent
	}
	if o.NumTcpTicklesSent != nil {
		toSerialize["num-tcp-tickles-sent"] = o.NumTcpTicklesSent
	}
	return json.Marshal(toSerialize)
}

type NullableSysruntimeHWFlowStats struct {
	value *SysruntimeHWFlowStats
	isSet bool
}

func (v NullableSysruntimeHWFlowStats) Get() *SysruntimeHWFlowStats {
	return v.value
}

func (v *NullableSysruntimeHWFlowStats) Set(val *SysruntimeHWFlowStats) {
	v.value = val
	v.isSet = true
}

func (v NullableSysruntimeHWFlowStats) IsSet() bool {
	return v.isSet
}

func (v *NullableSysruntimeHWFlowStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSysruntimeHWFlowStats(val *SysruntimeHWFlowStats) *NullableSysruntimeHWFlowStats {
	return &NullableSysruntimeHWFlowStats{value: val, isSet: true}
}

func (v NullableSysruntimeHWFlowStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSysruntimeHWFlowStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


