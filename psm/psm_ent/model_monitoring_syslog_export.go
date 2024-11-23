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

// MonitoringSyslogExport Syslog export configuration.
type MonitoringSyslogExport struct {
	Config *MonitoringSyslogExportConfig `json:"config,omitempty"`
	// event export format, SYSLOG_BSD default.
	Format *string `json:"format,omitempty"`
	// export target ip/port/protocol.
	Targets *[]MonitoringExportConfig `json:"targets,omitempty"`
}

// NewMonitoringSyslogExport instantiates a new MonitoringSyslogExport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitoringSyslogExport() *MonitoringSyslogExport {
	this := MonitoringSyslogExport{}
	var format string = "syslog-bsd"
	this.Format = &format
	return &this
}

// NewMonitoringSyslogExportWithDefaults instantiates a new MonitoringSyslogExport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitoringSyslogExportWithDefaults() *MonitoringSyslogExport {
	this := MonitoringSyslogExport{}
	var format string = "syslog-bsd"
	this.Format = &format
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *MonitoringSyslogExport) GetConfig() MonitoringSyslogExportConfig {
	if o == nil || o.Config == nil {
		var ret MonitoringSyslogExportConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringSyslogExport) GetConfigOk() (*MonitoringSyslogExportConfig, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *MonitoringSyslogExport) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given MonitoringSyslogExportConfig and assigns it to the Config field.
func (o *MonitoringSyslogExport) SetConfig(v MonitoringSyslogExportConfig) {
	o.Config = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *MonitoringSyslogExport) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringSyslogExport) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *MonitoringSyslogExport) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *MonitoringSyslogExport) SetFormat(v string) {
	o.Format = &v
}

// GetTargets returns the Targets field value if set, zero value otherwise.
func (o *MonitoringSyslogExport) GetTargets() []MonitoringExportConfig {
	if o == nil || o.Targets == nil {
		var ret []MonitoringExportConfig
		return ret
	}
	return *o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringSyslogExport) GetTargetsOk() (*[]MonitoringExportConfig, bool) {
	if o == nil || o.Targets == nil {
		return nil, false
	}
	return o.Targets, true
}

// HasTargets returns a boolean if a field has been set.
func (o *MonitoringSyslogExport) HasTargets() bool {
	if o != nil && o.Targets != nil {
		return true
	}

	return false
}

// SetTargets gets a reference to the given []MonitoringExportConfig and assigns it to the Targets field.
func (o *MonitoringSyslogExport) SetTargets(v []MonitoringExportConfig) {
	o.Targets = &v
}

func (o MonitoringSyslogExport) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}
	if o.Targets != nil {
		toSerialize["targets"] = o.Targets
	}
	return json.Marshal(toSerialize)
}

type NullableMonitoringSyslogExport struct {
	value *MonitoringSyslogExport
	isSet bool
}

func (v NullableMonitoringSyslogExport) Get() *MonitoringSyslogExport {
	return v.value
}

func (v *NullableMonitoringSyslogExport) Set(val *MonitoringSyslogExport) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringSyslogExport) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringSyslogExport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringSyslogExport(val *MonitoringSyslogExport) *NullableMonitoringSyslogExport {
	return &NullableMonitoringSyslogExport{value: val, isSet: true}
}

func (v NullableMonitoringSyslogExport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringSyslogExport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


