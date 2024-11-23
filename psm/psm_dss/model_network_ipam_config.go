/*
 * Network API reference
 *
 *  
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package psm_dss

import (
	"encoding/json"
)

// NetworkIPAMConfig struct for NetworkIPAMConfig
type NetworkIPAMConfig struct {
	BootstrapIpamOptions *NetworkBootstrapIPAMOptions `json:"bootstrap-ipam-options,omitempty"`
	IpamOptions *NetworkIPAMOptions `json:"ipam-options,omitempty"`
	Ipv4IpamPool *[]NetworkIPAMPoolInfo `json:"ipv4-ipam-pool,omitempty"`
	Ipv4StaticBindings *[]NetworkIPAMBinding `json:"ipv4-static-bindings,omitempty"`
}

// NewNetworkIPAMConfig instantiates a new NetworkIPAMConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkIPAMConfig() *NetworkIPAMConfig {
	this := NetworkIPAMConfig{}
	return &this
}

// NewNetworkIPAMConfigWithDefaults instantiates a new NetworkIPAMConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkIPAMConfigWithDefaults() *NetworkIPAMConfig {
	this := NetworkIPAMConfig{}
	return &this
}

// GetBootstrapIpamOptions returns the BootstrapIpamOptions field value if set, zero value otherwise.
func (o *NetworkIPAMConfig) GetBootstrapIpamOptions() NetworkBootstrapIPAMOptions {
	if o == nil || o.BootstrapIpamOptions == nil {
		var ret NetworkBootstrapIPAMOptions
		return ret
	}
	return *o.BootstrapIpamOptions
}

// GetBootstrapIpamOptionsOk returns a tuple with the BootstrapIpamOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkIPAMConfig) GetBootstrapIpamOptionsOk() (*NetworkBootstrapIPAMOptions, bool) {
	if o == nil || o.BootstrapIpamOptions == nil {
		return nil, false
	}
	return o.BootstrapIpamOptions, true
}

// HasBootstrapIpamOptions returns a boolean if a field has been set.
func (o *NetworkIPAMConfig) HasBootstrapIpamOptions() bool {
	if o != nil && o.BootstrapIpamOptions != nil {
		return true
	}

	return false
}

// SetBootstrapIpamOptions gets a reference to the given NetworkBootstrapIPAMOptions and assigns it to the BootstrapIpamOptions field.
func (o *NetworkIPAMConfig) SetBootstrapIpamOptions(v NetworkBootstrapIPAMOptions) {
	o.BootstrapIpamOptions = &v
}

// GetIpamOptions returns the IpamOptions field value if set, zero value otherwise.
func (o *NetworkIPAMConfig) GetIpamOptions() NetworkIPAMOptions {
	if o == nil || o.IpamOptions == nil {
		var ret NetworkIPAMOptions
		return ret
	}
	return *o.IpamOptions
}

// GetIpamOptionsOk returns a tuple with the IpamOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkIPAMConfig) GetIpamOptionsOk() (*NetworkIPAMOptions, bool) {
	if o == nil || o.IpamOptions == nil {
		return nil, false
	}
	return o.IpamOptions, true
}

// HasIpamOptions returns a boolean if a field has been set.
func (o *NetworkIPAMConfig) HasIpamOptions() bool {
	if o != nil && o.IpamOptions != nil {
		return true
	}

	return false
}

// SetIpamOptions gets a reference to the given NetworkIPAMOptions and assigns it to the IpamOptions field.
func (o *NetworkIPAMConfig) SetIpamOptions(v NetworkIPAMOptions) {
	o.IpamOptions = &v
}

// GetIpv4IpamPool returns the Ipv4IpamPool field value if set, zero value otherwise.
func (o *NetworkIPAMConfig) GetIpv4IpamPool() []NetworkIPAMPoolInfo {
	if o == nil || o.Ipv4IpamPool == nil {
		var ret []NetworkIPAMPoolInfo
		return ret
	}
	return *o.Ipv4IpamPool
}

// GetIpv4IpamPoolOk returns a tuple with the Ipv4IpamPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkIPAMConfig) GetIpv4IpamPoolOk() (*[]NetworkIPAMPoolInfo, bool) {
	if o == nil || o.Ipv4IpamPool == nil {
		return nil, false
	}
	return o.Ipv4IpamPool, true
}

// HasIpv4IpamPool returns a boolean if a field has been set.
func (o *NetworkIPAMConfig) HasIpv4IpamPool() bool {
	if o != nil && o.Ipv4IpamPool != nil {
		return true
	}

	return false
}

// SetIpv4IpamPool gets a reference to the given []NetworkIPAMPoolInfo and assigns it to the Ipv4IpamPool field.
func (o *NetworkIPAMConfig) SetIpv4IpamPool(v []NetworkIPAMPoolInfo) {
	o.Ipv4IpamPool = &v
}

// GetIpv4StaticBindings returns the Ipv4StaticBindings field value if set, zero value otherwise.
func (o *NetworkIPAMConfig) GetIpv4StaticBindings() []NetworkIPAMBinding {
	if o == nil || o.Ipv4StaticBindings == nil {
		var ret []NetworkIPAMBinding
		return ret
	}
	return *o.Ipv4StaticBindings
}

// GetIpv4StaticBindingsOk returns a tuple with the Ipv4StaticBindings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkIPAMConfig) GetIpv4StaticBindingsOk() (*[]NetworkIPAMBinding, bool) {
	if o == nil || o.Ipv4StaticBindings == nil {
		return nil, false
	}
	return o.Ipv4StaticBindings, true
}

// HasIpv4StaticBindings returns a boolean if a field has been set.
func (o *NetworkIPAMConfig) HasIpv4StaticBindings() bool {
	if o != nil && o.Ipv4StaticBindings != nil {
		return true
	}

	return false
}

// SetIpv4StaticBindings gets a reference to the given []NetworkIPAMBinding and assigns it to the Ipv4StaticBindings field.
func (o *NetworkIPAMConfig) SetIpv4StaticBindings(v []NetworkIPAMBinding) {
	o.Ipv4StaticBindings = &v
}

func (o NetworkIPAMConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BootstrapIpamOptions != nil {
		toSerialize["bootstrap-ipam-options"] = o.BootstrapIpamOptions
	}
	if o.IpamOptions != nil {
		toSerialize["ipam-options"] = o.IpamOptions
	}
	if o.Ipv4IpamPool != nil {
		toSerialize["ipv4-ipam-pool"] = o.Ipv4IpamPool
	}
	if o.Ipv4StaticBindings != nil {
		toSerialize["ipv4-static-bindings"] = o.Ipv4StaticBindings
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkIPAMConfig struct {
	value *NetworkIPAMConfig
	isSet bool
}

func (v NullableNetworkIPAMConfig) Get() *NetworkIPAMConfig {
	return v.value
}

func (v *NullableNetworkIPAMConfig) Set(val *NetworkIPAMConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkIPAMConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkIPAMConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkIPAMConfig(val *NetworkIPAMConfig) *NullableNetworkIPAMConfig {
	return &NullableNetworkIPAMConfig{value: val, isSet: true}
}

func (v NullableNetworkIPAMConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkIPAMConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


