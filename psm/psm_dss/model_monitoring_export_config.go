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

// MonitoringExportConfig Export Config specifies server address and user credentials.
type MonitoringExportConfig struct {
	// IP address of the collector/entity to which the data is to be exported. Should be a valid IPv4 address.
	Destination *string `json:"destination,omitempty"`
	// Gateway of the dest IP address to which the data is to be exported. Should be a valid IPv4 address.
	Gateway *string `json:"gateway,omitempty"`
	// protocol and Port number where an external collector is gathering the data example \"UDP/2055\". Should be a valid layer 3 or layer 4 protocol and port/type (only support UDP currently).
	Transport *string `json:"transport,omitempty"`
	// Destination Virtual Router.
	VirtualRouter *string `json:"virtual-router,omitempty"`
}

// NewMonitoringExportConfig instantiates a new MonitoringExportConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitoringExportConfig() *MonitoringExportConfig {
	this := MonitoringExportConfig{}
	return &this
}

// NewMonitoringExportConfigWithDefaults instantiates a new MonitoringExportConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitoringExportConfigWithDefaults() *MonitoringExportConfig {
	this := MonitoringExportConfig{}
	return &this
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *MonitoringExportConfig) GetDestination() string {
	if o == nil || o.Destination == nil {
		var ret string
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringExportConfig) GetDestinationOk() (*string, bool) {
	if o == nil || o.Destination == nil {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *MonitoringExportConfig) HasDestination() bool {
	if o != nil && o.Destination != nil {
		return true
	}

	return false
}

// SetDestination gets a reference to the given string and assigns it to the Destination field.
func (o *MonitoringExportConfig) SetDestination(v string) {
	o.Destination = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *MonitoringExportConfig) GetGateway() string {
	if o == nil || o.Gateway == nil {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringExportConfig) GetGatewayOk() (*string, bool) {
	if o == nil || o.Gateway == nil {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *MonitoringExportConfig) HasGateway() bool {
	if o != nil && o.Gateway != nil {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *MonitoringExportConfig) SetGateway(v string) {
	o.Gateway = &v
}

// GetTransport returns the Transport field value if set, zero value otherwise.
func (o *MonitoringExportConfig) GetTransport() string {
	if o == nil || o.Transport == nil {
		var ret string
		return ret
	}
	return *o.Transport
}

// GetTransportOk returns a tuple with the Transport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringExportConfig) GetTransportOk() (*string, bool) {
	if o == nil || o.Transport == nil {
		return nil, false
	}
	return o.Transport, true
}

// HasTransport returns a boolean if a field has been set.
func (o *MonitoringExportConfig) HasTransport() bool {
	if o != nil && o.Transport != nil {
		return true
	}

	return false
}

// SetTransport gets a reference to the given string and assigns it to the Transport field.
func (o *MonitoringExportConfig) SetTransport(v string) {
	o.Transport = &v
}

// GetVirtualRouter returns the VirtualRouter field value if set, zero value otherwise.
func (o *MonitoringExportConfig) GetVirtualRouter() string {
	if o == nil || o.VirtualRouter == nil {
		var ret string
		return ret
	}
	return *o.VirtualRouter
}

// GetVirtualRouterOk returns a tuple with the VirtualRouter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringExportConfig) GetVirtualRouterOk() (*string, bool) {
	if o == nil || o.VirtualRouter == nil {
		return nil, false
	}
	return o.VirtualRouter, true
}

// HasVirtualRouter returns a boolean if a field has been set.
func (o *MonitoringExportConfig) HasVirtualRouter() bool {
	if o != nil && o.VirtualRouter != nil {
		return true
	}

	return false
}

// SetVirtualRouter gets a reference to the given string and assigns it to the VirtualRouter field.
func (o *MonitoringExportConfig) SetVirtualRouter(v string) {
	o.VirtualRouter = &v
}

func (o MonitoringExportConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Destination != nil {
		toSerialize["destination"] = o.Destination
	}
	if o.Gateway != nil {
		toSerialize["gateway"] = o.Gateway
	}
	if o.Transport != nil {
		toSerialize["transport"] = o.Transport
	}
	if o.VirtualRouter != nil {
		toSerialize["virtual-router"] = o.VirtualRouter
	}
	return json.Marshal(toSerialize)
}

type NullableMonitoringExportConfig struct {
	value *MonitoringExportConfig
	isSet bool
}

func (v NullableMonitoringExportConfig) Get() *MonitoringExportConfig {
	return v.value
}

func (v *NullableMonitoringExportConfig) Set(val *MonitoringExportConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringExportConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringExportConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringExportConfig(val *MonitoringExportConfig) *NullableMonitoringExportConfig {
	return &NullableMonitoringExportConfig{value: val, isSet: true}
}

func (v NullableMonitoringExportConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringExportConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


