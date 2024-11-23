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

// NetworkBGPAuthStatus struct for NetworkBGPAuthStatus
type NetworkBGPAuthStatus struct {
	// Neighbor IP Address.
	IpAddress *string `json:"ip-address,omitempty"`
	RemoteAs *ApiBgpAsn `json:"remote-as,omitempty"`
	// Authentication status.
	Status *string `json:"status,omitempty"`
}

// NewNetworkBGPAuthStatus instantiates a new NetworkBGPAuthStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkBGPAuthStatus() *NetworkBGPAuthStatus {
	this := NetworkBGPAuthStatus{}
	var status string = "disabled"
	this.Status = &status
	return &this
}

// NewNetworkBGPAuthStatusWithDefaults instantiates a new NetworkBGPAuthStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkBGPAuthStatusWithDefaults() *NetworkBGPAuthStatus {
	this := NetworkBGPAuthStatus{}
	var status string = "disabled"
	this.Status = &status
	return &this
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *NetworkBGPAuthStatus) GetIpAddress() string {
	if o == nil || o.IpAddress == nil {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkBGPAuthStatus) GetIpAddressOk() (*string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *NetworkBGPAuthStatus) HasIpAddress() bool {
	if o != nil && o.IpAddress != nil {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *NetworkBGPAuthStatus) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetRemoteAs returns the RemoteAs field value if set, zero value otherwise.
func (o *NetworkBGPAuthStatus) GetRemoteAs() ApiBgpAsn {
	if o == nil || o.RemoteAs == nil {
		var ret ApiBgpAsn
		return ret
	}
	return *o.RemoteAs
}

// GetRemoteAsOk returns a tuple with the RemoteAs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkBGPAuthStatus) GetRemoteAsOk() (*ApiBgpAsn, bool) {
	if o == nil || o.RemoteAs == nil {
		return nil, false
	}
	return o.RemoteAs, true
}

// HasRemoteAs returns a boolean if a field has been set.
func (o *NetworkBGPAuthStatus) HasRemoteAs() bool {
	if o != nil && o.RemoteAs != nil {
		return true
	}

	return false
}

// SetRemoteAs gets a reference to the given ApiBgpAsn and assigns it to the RemoteAs field.
func (o *NetworkBGPAuthStatus) SetRemoteAs(v ApiBgpAsn) {
	o.RemoteAs = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *NetworkBGPAuthStatus) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkBGPAuthStatus) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *NetworkBGPAuthStatus) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *NetworkBGPAuthStatus) SetStatus(v string) {
	o.Status = &v
}

func (o NetworkBGPAuthStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IpAddress != nil {
		toSerialize["ip-address"] = o.IpAddress
	}
	if o.RemoteAs != nil {
		toSerialize["remote-as"] = o.RemoteAs
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkBGPAuthStatus struct {
	value *NetworkBGPAuthStatus
	isSet bool
}

func (v NullableNetworkBGPAuthStatus) Get() *NetworkBGPAuthStatus {
	return v.value
}

func (v *NullableNetworkBGPAuthStatus) Set(val *NetworkBGPAuthStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkBGPAuthStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkBGPAuthStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkBGPAuthStatus(val *NetworkBGPAuthStatus) *NullableNetworkBGPAuthStatus {
	return &NullableNetworkBGPAuthStatus{value: val, isSet: true}
}

func (v NullableNetworkBGPAuthStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkBGPAuthStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

