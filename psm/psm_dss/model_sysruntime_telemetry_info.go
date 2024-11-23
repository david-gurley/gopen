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

// SysruntimeTelemetryInfo TelemetryInfo captures telemetry session state of the flow.
type SysruntimeTelemetryInfo struct {
	EgressMirrorEnabled *bool `json:"egress-mirror-enabled,omitempty"`
	FlowExportEnabled *bool `json:"flow-export-enabled,omitempty"`
	IngressMirrorEnabled *bool `json:"ingress-mirror-enabled,omitempty"`
}

// NewSysruntimeTelemetryInfo instantiates a new SysruntimeTelemetryInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSysruntimeTelemetryInfo() *SysruntimeTelemetryInfo {
	this := SysruntimeTelemetryInfo{}
	return &this
}

// NewSysruntimeTelemetryInfoWithDefaults instantiates a new SysruntimeTelemetryInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSysruntimeTelemetryInfoWithDefaults() *SysruntimeTelemetryInfo {
	this := SysruntimeTelemetryInfo{}
	return &this
}

// GetEgressMirrorEnabled returns the EgressMirrorEnabled field value if set, zero value otherwise.
func (o *SysruntimeTelemetryInfo) GetEgressMirrorEnabled() bool {
	if o == nil || o.EgressMirrorEnabled == nil {
		var ret bool
		return ret
	}
	return *o.EgressMirrorEnabled
}

// GetEgressMirrorEnabledOk returns a tuple with the EgressMirrorEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SysruntimeTelemetryInfo) GetEgressMirrorEnabledOk() (*bool, bool) {
	if o == nil || o.EgressMirrorEnabled == nil {
		return nil, false
	}
	return o.EgressMirrorEnabled, true
}

// HasEgressMirrorEnabled returns a boolean if a field has been set.
func (o *SysruntimeTelemetryInfo) HasEgressMirrorEnabled() bool {
	if o != nil && o.EgressMirrorEnabled != nil {
		return true
	}

	return false
}

// SetEgressMirrorEnabled gets a reference to the given bool and assigns it to the EgressMirrorEnabled field.
func (o *SysruntimeTelemetryInfo) SetEgressMirrorEnabled(v bool) {
	o.EgressMirrorEnabled = &v
}

// GetFlowExportEnabled returns the FlowExportEnabled field value if set, zero value otherwise.
func (o *SysruntimeTelemetryInfo) GetFlowExportEnabled() bool {
	if o == nil || o.FlowExportEnabled == nil {
		var ret bool
		return ret
	}
	return *o.FlowExportEnabled
}

// GetFlowExportEnabledOk returns a tuple with the FlowExportEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SysruntimeTelemetryInfo) GetFlowExportEnabledOk() (*bool, bool) {
	if o == nil || o.FlowExportEnabled == nil {
		return nil, false
	}
	return o.FlowExportEnabled, true
}

// HasFlowExportEnabled returns a boolean if a field has been set.
func (o *SysruntimeTelemetryInfo) HasFlowExportEnabled() bool {
	if o != nil && o.FlowExportEnabled != nil {
		return true
	}

	return false
}

// SetFlowExportEnabled gets a reference to the given bool and assigns it to the FlowExportEnabled field.
func (o *SysruntimeTelemetryInfo) SetFlowExportEnabled(v bool) {
	o.FlowExportEnabled = &v
}

// GetIngressMirrorEnabled returns the IngressMirrorEnabled field value if set, zero value otherwise.
func (o *SysruntimeTelemetryInfo) GetIngressMirrorEnabled() bool {
	if o == nil || o.IngressMirrorEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IngressMirrorEnabled
}

// GetIngressMirrorEnabledOk returns a tuple with the IngressMirrorEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SysruntimeTelemetryInfo) GetIngressMirrorEnabledOk() (*bool, bool) {
	if o == nil || o.IngressMirrorEnabled == nil {
		return nil, false
	}
	return o.IngressMirrorEnabled, true
}

// HasIngressMirrorEnabled returns a boolean if a field has been set.
func (o *SysruntimeTelemetryInfo) HasIngressMirrorEnabled() bool {
	if o != nil && o.IngressMirrorEnabled != nil {
		return true
	}

	return false
}

// SetIngressMirrorEnabled gets a reference to the given bool and assigns it to the IngressMirrorEnabled field.
func (o *SysruntimeTelemetryInfo) SetIngressMirrorEnabled(v bool) {
	o.IngressMirrorEnabled = &v
}

func (o SysruntimeTelemetryInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EgressMirrorEnabled != nil {
		toSerialize["egress-mirror-enabled"] = o.EgressMirrorEnabled
	}
	if o.FlowExportEnabled != nil {
		toSerialize["flow-export-enabled"] = o.FlowExportEnabled
	}
	if o.IngressMirrorEnabled != nil {
		toSerialize["ingress-mirror-enabled"] = o.IngressMirrorEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableSysruntimeTelemetryInfo struct {
	value *SysruntimeTelemetryInfo
	isSet bool
}

func (v NullableSysruntimeTelemetryInfo) Get() *SysruntimeTelemetryInfo {
	return v.value
}

func (v *NullableSysruntimeTelemetryInfo) Set(val *SysruntimeTelemetryInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSysruntimeTelemetryInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSysruntimeTelemetryInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSysruntimeTelemetryInfo(val *SysruntimeTelemetryInfo) *NullableSysruntimeTelemetryInfo {
	return &NullableSysruntimeTelemetryInfo{value: val, isSet: true}
}

func (v NullableSysruntimeTelemetryInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSysruntimeTelemetryInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

