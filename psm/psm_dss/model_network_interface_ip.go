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

// NetworkInterfaceIP struct for NetworkInterfaceIP
type NetworkInterfaceIP struct {
	GatewayIp *string `json:"gateway-ip,omitempty"`
	InterfaceId *int64 `json:"interface-id,omitempty"`
	IpAddress *string `json:"ip-address,omitempty"`
}

// NewNetworkInterfaceIP instantiates a new NetworkInterfaceIP object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkInterfaceIP() *NetworkInterfaceIP {
	this := NetworkInterfaceIP{}
	return &this
}

// NewNetworkInterfaceIPWithDefaults instantiates a new NetworkInterfaceIP object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkInterfaceIPWithDefaults() *NetworkInterfaceIP {
	this := NetworkInterfaceIP{}
	return &this
}

// GetGatewayIp returns the GatewayIp field value if set, zero value otherwise.
func (o *NetworkInterfaceIP) GetGatewayIp() string {
	if o == nil || o.GatewayIp == nil {
		var ret string
		return ret
	}
	return *o.GatewayIp
}

// GetGatewayIpOk returns a tuple with the GatewayIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkInterfaceIP) GetGatewayIpOk() (*string, bool) {
	if o == nil || o.GatewayIp == nil {
		return nil, false
	}
	return o.GatewayIp, true
}

// HasGatewayIp returns a boolean if a field has been set.
func (o *NetworkInterfaceIP) HasGatewayIp() bool {
	if o != nil && o.GatewayIp != nil {
		return true
	}

	return false
}

// SetGatewayIp gets a reference to the given string and assigns it to the GatewayIp field.
func (o *NetworkInterfaceIP) SetGatewayIp(v string) {
	o.GatewayIp = &v
}

// GetInterfaceId returns the InterfaceId field value if set, zero value otherwise.
func (o *NetworkInterfaceIP) GetInterfaceId() int64 {
	if o == nil || o.InterfaceId == nil {
		var ret int64
		return ret
	}
	return *o.InterfaceId
}

// GetInterfaceIdOk returns a tuple with the InterfaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkInterfaceIP) GetInterfaceIdOk() (*int64, bool) {
	if o == nil || o.InterfaceId == nil {
		return nil, false
	}
	return o.InterfaceId, true
}

// HasInterfaceId returns a boolean if a field has been set.
func (o *NetworkInterfaceIP) HasInterfaceId() bool {
	if o != nil && o.InterfaceId != nil {
		return true
	}

	return false
}

// SetInterfaceId gets a reference to the given int64 and assigns it to the InterfaceId field.
func (o *NetworkInterfaceIP) SetInterfaceId(v int64) {
	o.InterfaceId = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *NetworkInterfaceIP) GetIpAddress() string {
	if o == nil || o.IpAddress == nil {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkInterfaceIP) GetIpAddressOk() (*string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *NetworkInterfaceIP) HasIpAddress() bool {
	if o != nil && o.IpAddress != nil {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *NetworkInterfaceIP) SetIpAddress(v string) {
	o.IpAddress = &v
}

func (o NetworkInterfaceIP) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GatewayIp != nil {
		toSerialize["gateway-ip"] = o.GatewayIp
	}
	if o.InterfaceId != nil {
		toSerialize["interface-id"] = o.InterfaceId
	}
	if o.IpAddress != nil {
		toSerialize["ip-address"] = o.IpAddress
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkInterfaceIP struct {
	value *NetworkInterfaceIP
	isSet bool
}

func (v NullableNetworkInterfaceIP) Get() *NetworkInterfaceIP {
	return v.value
}

func (v *NullableNetworkInterfaceIP) Set(val *NetworkInterfaceIP) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkInterfaceIP) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkInterfaceIP) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkInterfaceIP(val *NetworkInterfaceIP) *NullableNetworkInterfaceIP {
	return &NullableNetworkInterfaceIP{value: val, isSet: true}
}

func (v NullableNetworkInterfaceIP) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkInterfaceIP) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

