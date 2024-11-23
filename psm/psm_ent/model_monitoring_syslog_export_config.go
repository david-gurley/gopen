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

// MonitoringSyslogExportConfig syslog export configurations.
type MonitoringSyslogExportConfig struct {
	// override default facility with this in exported logs.
	FacilityOverride *string `json:"facility-override,omitempty"`
	// add prefix in exported logs.
	Prefix *string `json:"prefix,omitempty"`
}

// NewMonitoringSyslogExportConfig instantiates a new MonitoringSyslogExportConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitoringSyslogExportConfig() *MonitoringSyslogExportConfig {
	this := MonitoringSyslogExportConfig{}
	var facilityOverride string = "user"
	this.FacilityOverride = &facilityOverride
	return &this
}

// NewMonitoringSyslogExportConfigWithDefaults instantiates a new MonitoringSyslogExportConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitoringSyslogExportConfigWithDefaults() *MonitoringSyslogExportConfig {
	this := MonitoringSyslogExportConfig{}
	var facilityOverride string = "user"
	this.FacilityOverride = &facilityOverride
	return &this
}

// GetFacilityOverride returns the FacilityOverride field value if set, zero value otherwise.
func (o *MonitoringSyslogExportConfig) GetFacilityOverride() string {
	if o == nil || o.FacilityOverride == nil {
		var ret string
		return ret
	}
	return *o.FacilityOverride
}

// GetFacilityOverrideOk returns a tuple with the FacilityOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringSyslogExportConfig) GetFacilityOverrideOk() (*string, bool) {
	if o == nil || o.FacilityOverride == nil {
		return nil, false
	}
	return o.FacilityOverride, true
}

// HasFacilityOverride returns a boolean if a field has been set.
func (o *MonitoringSyslogExportConfig) HasFacilityOverride() bool {
	if o != nil && o.FacilityOverride != nil {
		return true
	}

	return false
}

// SetFacilityOverride gets a reference to the given string and assigns it to the FacilityOverride field.
func (o *MonitoringSyslogExportConfig) SetFacilityOverride(v string) {
	o.FacilityOverride = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *MonitoringSyslogExportConfig) GetPrefix() string {
	if o == nil || o.Prefix == nil {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringSyslogExportConfig) GetPrefixOk() (*string, bool) {
	if o == nil || o.Prefix == nil {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *MonitoringSyslogExportConfig) HasPrefix() bool {
	if o != nil && o.Prefix != nil {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *MonitoringSyslogExportConfig) SetPrefix(v string) {
	o.Prefix = &v
}

func (o MonitoringSyslogExportConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FacilityOverride != nil {
		toSerialize["facility-override"] = o.FacilityOverride
	}
	if o.Prefix != nil {
		toSerialize["prefix"] = o.Prefix
	}
	return json.Marshal(toSerialize)
}

type NullableMonitoringSyslogExportConfig struct {
	value *MonitoringSyslogExportConfig
	isSet bool
}

func (v NullableMonitoringSyslogExportConfig) Get() *MonitoringSyslogExportConfig {
	return v.value
}

func (v *NullableMonitoringSyslogExportConfig) Set(val *MonitoringSyslogExportConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringSyslogExportConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringSyslogExportConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringSyslogExportConfig(val *MonitoringSyslogExportConfig) *NullableMonitoringSyslogExportConfig {
	return &NullableMonitoringSyslogExportConfig{value: val, isSet: true}
}

func (v NullableMonitoringSyslogExportConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringSyslogExportConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


