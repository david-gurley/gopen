/*
 * Routing API reference
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

// EVPNRouteEVPNType5Route struct for EVPNRouteEVPNType5Route
type EVPNRouteEVPNType5Route struct {
	Esi *string `json:"esi,omitempty"`
	EthTag *string `json:"eth-tag,omitempty"`
	GwIpAddr *string `json:"gw-ip-addr,omitempty"`
	IpPrefix *string `json:"ip-prefix,omitempty"`
	MplsLabel1 *string `json:"mpls-label1,omitempty"`
	Rd *string `json:"rd,omitempty"`
}

// NewEVPNRouteEVPNType5Route instantiates a new EVPNRouteEVPNType5Route object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEVPNRouteEVPNType5Route() *EVPNRouteEVPNType5Route {
	this := EVPNRouteEVPNType5Route{}
	return &this
}

// NewEVPNRouteEVPNType5RouteWithDefaults instantiates a new EVPNRouteEVPNType5Route object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEVPNRouteEVPNType5RouteWithDefaults() *EVPNRouteEVPNType5Route {
	this := EVPNRouteEVPNType5Route{}
	return &this
}

// GetEsi returns the Esi field value if set, zero value otherwise.
func (o *EVPNRouteEVPNType5Route) GetEsi() string {
	if o == nil || o.Esi == nil {
		var ret string
		return ret
	}
	return *o.Esi
}

// GetEsiOk returns a tuple with the Esi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EVPNRouteEVPNType5Route) GetEsiOk() (*string, bool) {
	if o == nil || o.Esi == nil {
		return nil, false
	}
	return o.Esi, true
}

// HasEsi returns a boolean if a field has been set.
func (o *EVPNRouteEVPNType5Route) HasEsi() bool {
	if o != nil && o.Esi != nil {
		return true
	}

	return false
}

// SetEsi gets a reference to the given string and assigns it to the Esi field.
func (o *EVPNRouteEVPNType5Route) SetEsi(v string) {
	o.Esi = &v
}

// GetEthTag returns the EthTag field value if set, zero value otherwise.
func (o *EVPNRouteEVPNType5Route) GetEthTag() string {
	if o == nil || o.EthTag == nil {
		var ret string
		return ret
	}
	return *o.EthTag
}

// GetEthTagOk returns a tuple with the EthTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EVPNRouteEVPNType5Route) GetEthTagOk() (*string, bool) {
	if o == nil || o.EthTag == nil {
		return nil, false
	}
	return o.EthTag, true
}

// HasEthTag returns a boolean if a field has been set.
func (o *EVPNRouteEVPNType5Route) HasEthTag() bool {
	if o != nil && o.EthTag != nil {
		return true
	}

	return false
}

// SetEthTag gets a reference to the given string and assigns it to the EthTag field.
func (o *EVPNRouteEVPNType5Route) SetEthTag(v string) {
	o.EthTag = &v
}

// GetGwIpAddr returns the GwIpAddr field value if set, zero value otherwise.
func (o *EVPNRouteEVPNType5Route) GetGwIpAddr() string {
	if o == nil || o.GwIpAddr == nil {
		var ret string
		return ret
	}
	return *o.GwIpAddr
}

// GetGwIpAddrOk returns a tuple with the GwIpAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EVPNRouteEVPNType5Route) GetGwIpAddrOk() (*string, bool) {
	if o == nil || o.GwIpAddr == nil {
		return nil, false
	}
	return o.GwIpAddr, true
}

// HasGwIpAddr returns a boolean if a field has been set.
func (o *EVPNRouteEVPNType5Route) HasGwIpAddr() bool {
	if o != nil && o.GwIpAddr != nil {
		return true
	}

	return false
}

// SetGwIpAddr gets a reference to the given string and assigns it to the GwIpAddr field.
func (o *EVPNRouteEVPNType5Route) SetGwIpAddr(v string) {
	o.GwIpAddr = &v
}

// GetIpPrefix returns the IpPrefix field value if set, zero value otherwise.
func (o *EVPNRouteEVPNType5Route) GetIpPrefix() string {
	if o == nil || o.IpPrefix == nil {
		var ret string
		return ret
	}
	return *o.IpPrefix
}

// GetIpPrefixOk returns a tuple with the IpPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EVPNRouteEVPNType5Route) GetIpPrefixOk() (*string, bool) {
	if o == nil || o.IpPrefix == nil {
		return nil, false
	}
	return o.IpPrefix, true
}

// HasIpPrefix returns a boolean if a field has been set.
func (o *EVPNRouteEVPNType5Route) HasIpPrefix() bool {
	if o != nil && o.IpPrefix != nil {
		return true
	}

	return false
}

// SetIpPrefix gets a reference to the given string and assigns it to the IpPrefix field.
func (o *EVPNRouteEVPNType5Route) SetIpPrefix(v string) {
	o.IpPrefix = &v
}

// GetMplsLabel1 returns the MplsLabel1 field value if set, zero value otherwise.
func (o *EVPNRouteEVPNType5Route) GetMplsLabel1() string {
	if o == nil || o.MplsLabel1 == nil {
		var ret string
		return ret
	}
	return *o.MplsLabel1
}

// GetMplsLabel1Ok returns a tuple with the MplsLabel1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EVPNRouteEVPNType5Route) GetMplsLabel1Ok() (*string, bool) {
	if o == nil || o.MplsLabel1 == nil {
		return nil, false
	}
	return o.MplsLabel1, true
}

// HasMplsLabel1 returns a boolean if a field has been set.
func (o *EVPNRouteEVPNType5Route) HasMplsLabel1() bool {
	if o != nil && o.MplsLabel1 != nil {
		return true
	}

	return false
}

// SetMplsLabel1 gets a reference to the given string and assigns it to the MplsLabel1 field.
func (o *EVPNRouteEVPNType5Route) SetMplsLabel1(v string) {
	o.MplsLabel1 = &v
}

// GetRd returns the Rd field value if set, zero value otherwise.
func (o *EVPNRouteEVPNType5Route) GetRd() string {
	if o == nil || o.Rd == nil {
		var ret string
		return ret
	}
	return *o.Rd
}

// GetRdOk returns a tuple with the Rd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EVPNRouteEVPNType5Route) GetRdOk() (*string, bool) {
	if o == nil || o.Rd == nil {
		return nil, false
	}
	return o.Rd, true
}

// HasRd returns a boolean if a field has been set.
func (o *EVPNRouteEVPNType5Route) HasRd() bool {
	if o != nil && o.Rd != nil {
		return true
	}

	return false
}

// SetRd gets a reference to the given string and assigns it to the Rd field.
func (o *EVPNRouteEVPNType5Route) SetRd(v string) {
	o.Rd = &v
}

func (o EVPNRouteEVPNType5Route) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Esi != nil {
		toSerialize["esi"] = o.Esi
	}
	if o.EthTag != nil {
		toSerialize["eth-tag"] = o.EthTag
	}
	if o.GwIpAddr != nil {
		toSerialize["gw-ip-addr"] = o.GwIpAddr
	}
	if o.IpPrefix != nil {
		toSerialize["ip-prefix"] = o.IpPrefix
	}
	if o.MplsLabel1 != nil {
		toSerialize["mpls-label1"] = o.MplsLabel1
	}
	if o.Rd != nil {
		toSerialize["rd"] = o.Rd
	}
	return json.Marshal(toSerialize)
}

type NullableEVPNRouteEVPNType5Route struct {
	value *EVPNRouteEVPNType5Route
	isSet bool
}

func (v NullableEVPNRouteEVPNType5Route) Get() *EVPNRouteEVPNType5Route {
	return v.value
}

func (v *NullableEVPNRouteEVPNType5Route) Set(val *EVPNRouteEVPNType5Route) {
	v.value = val
	v.isSet = true
}

func (v NullableEVPNRouteEVPNType5Route) IsSet() bool {
	return v.isSet
}

func (v *NullableEVPNRouteEVPNType5Route) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEVPNRouteEVPNType5Route(val *EVPNRouteEVPNType5Route) *NullableEVPNRouteEVPNType5Route {
	return &NullableEVPNRouteEVPNType5Route{value: val, isSet: true}
}

func (v NullableEVPNRouteEVPNType5Route) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEVPNRouteEVPNType5Route) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

