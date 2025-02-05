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

// MonitoringPrivacyConfig PrivacyConfig contains the configuration for SNMP Trap encryption.
type MonitoringPrivacyConfig struct {
	Algo *string `json:"algo,omitempty"`
	// Password contains the privacy password.
	Password *string `json:"password,omitempty"`
}

// NewMonitoringPrivacyConfig instantiates a new MonitoringPrivacyConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitoringPrivacyConfig() *MonitoringPrivacyConfig {
	this := MonitoringPrivacyConfig{}
	var algo string = "des56"
	this.Algo = &algo
	return &this
}

// NewMonitoringPrivacyConfigWithDefaults instantiates a new MonitoringPrivacyConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitoringPrivacyConfigWithDefaults() *MonitoringPrivacyConfig {
	this := MonitoringPrivacyConfig{}
	var algo string = "des56"
	this.Algo = &algo
	return &this
}

// GetAlgo returns the Algo field value if set, zero value otherwise.
func (o *MonitoringPrivacyConfig) GetAlgo() string {
	if o == nil || o.Algo == nil {
		var ret string
		return ret
	}
	return *o.Algo
}

// GetAlgoOk returns a tuple with the Algo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringPrivacyConfig) GetAlgoOk() (*string, bool) {
	if o == nil || o.Algo == nil {
		return nil, false
	}
	return o.Algo, true
}

// HasAlgo returns a boolean if a field has been set.
func (o *MonitoringPrivacyConfig) HasAlgo() bool {
	if o != nil && o.Algo != nil {
		return true
	}

	return false
}

// SetAlgo gets a reference to the given string and assigns it to the Algo field.
func (o *MonitoringPrivacyConfig) SetAlgo(v string) {
	o.Algo = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *MonitoringPrivacyConfig) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringPrivacyConfig) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *MonitoringPrivacyConfig) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *MonitoringPrivacyConfig) SetPassword(v string) {
	o.Password = &v
}

func (o MonitoringPrivacyConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Algo != nil {
		toSerialize["algo"] = o.Algo
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableMonitoringPrivacyConfig struct {
	value *MonitoringPrivacyConfig
	isSet bool
}

func (v NullableMonitoringPrivacyConfig) Get() *MonitoringPrivacyConfig {
	return v.value
}

func (v *NullableMonitoringPrivacyConfig) Set(val *MonitoringPrivacyConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringPrivacyConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringPrivacyConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringPrivacyConfig(val *MonitoringPrivacyConfig) *NullableMonitoringPrivacyConfig {
	return &NullableMonitoringPrivacyConfig{value: val, isSet: true}
}

func (v NullableMonitoringPrivacyConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringPrivacyConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


