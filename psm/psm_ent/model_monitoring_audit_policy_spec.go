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

// MonitoringAuditPolicySpec AuditPolicySpec is the specification of an AuditEvent Policy.
type MonitoringAuditPolicySpec struct {
	SyslogAuditor *MonitoringSyslogAuditor `json:"syslog-auditor,omitempty"`
}

// NewMonitoringAuditPolicySpec instantiates a new MonitoringAuditPolicySpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitoringAuditPolicySpec() *MonitoringAuditPolicySpec {
	this := MonitoringAuditPolicySpec{}
	return &this
}

// NewMonitoringAuditPolicySpecWithDefaults instantiates a new MonitoringAuditPolicySpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitoringAuditPolicySpecWithDefaults() *MonitoringAuditPolicySpec {
	this := MonitoringAuditPolicySpec{}
	return &this
}

// GetSyslogAuditor returns the SyslogAuditor field value if set, zero value otherwise.
func (o *MonitoringAuditPolicySpec) GetSyslogAuditor() MonitoringSyslogAuditor {
	if o == nil || o.SyslogAuditor == nil {
		var ret MonitoringSyslogAuditor
		return ret
	}
	return *o.SyslogAuditor
}

// GetSyslogAuditorOk returns a tuple with the SyslogAuditor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringAuditPolicySpec) GetSyslogAuditorOk() (*MonitoringSyslogAuditor, bool) {
	if o == nil || o.SyslogAuditor == nil {
		return nil, false
	}
	return o.SyslogAuditor, true
}

// HasSyslogAuditor returns a boolean if a field has been set.
func (o *MonitoringAuditPolicySpec) HasSyslogAuditor() bool {
	if o != nil && o.SyslogAuditor != nil {
		return true
	}

	return false
}

// SetSyslogAuditor gets a reference to the given MonitoringSyslogAuditor and assigns it to the SyslogAuditor field.
func (o *MonitoringAuditPolicySpec) SetSyslogAuditor(v MonitoringSyslogAuditor) {
	o.SyslogAuditor = &v
}

func (o MonitoringAuditPolicySpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SyslogAuditor != nil {
		toSerialize["syslog-auditor"] = o.SyslogAuditor
	}
	return json.Marshal(toSerialize)
}

type NullableMonitoringAuditPolicySpec struct {
	value *MonitoringAuditPolicySpec
	isSet bool
}

func (v NullableMonitoringAuditPolicySpec) Get() *MonitoringAuditPolicySpec {
	return v.value
}

func (v *NullableMonitoringAuditPolicySpec) Set(val *MonitoringAuditPolicySpec) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringAuditPolicySpec) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringAuditPolicySpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringAuditPolicySpec(val *MonitoringAuditPolicySpec) *NullableMonitoringAuditPolicySpec {
	return &NullableMonitoringAuditPolicySpec{value: val, isSet: true}
}

func (v NullableMonitoringAuditPolicySpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringAuditPolicySpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


