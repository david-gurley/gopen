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

// EVPNRouteEVPNType2Route struct for EVPNRouteEVPNType2Route
type EVPNRouteEVPNType2Route struct {
	Esi *string `json:"esi,omitempty"`
	EthTag *string `json:"eth-tag,omitempty"`
	IpAddr *string `json:"ip-addr,omitempty"`
	Mac *string `json:"mac,omitempty"`
	MplsLabel1 *string `json:"mpls-label1,omitempty"`
	MplsLabel2 *string `json:"mpls-label2,omitempty"`
	Rd *string `json:"rd,omitempty"`
}

// NewEVPNRouteEVPNType2Route instantiates a new EVPNRouteEVPNType2Route object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEVPNRouteEVPNType2Route() *EVPNRouteEVPNType2Route {
	this := EVPNRouteEVPNType2Route{}
	return &this
}

// NewEVPNRouteEVPNType2RouteWithDefaults instantiates a new EVPNRouteEVPNType2Route object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEVPNRouteEVPNType2RouteWithDefaults() *EVPNRouteEVPNType2Route {
	this := EVPNRouteEVPNType2Route{}
	return &this
}

// GetEsi returns the Esi field value if set, zero value otherwise.
func (o *EVPNRouteEVPNType2Route) GetEsi() string {
	if o == nil || o.Esi == nil {
		var ret string
		return ret
	}
	return *o.Esi
}

// GetEsiOk returns a tuple with the Esi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EVPNRouteEVPNType2Route) GetEsiOk() (*string, bool) {
	if o == nil || o.Esi == nil {
		return nil, false
	}
	return o.Esi, true
}

// HasEsi returns a boolean if a field has been set.
func (o *EVPNRouteEVPNType2Route) HasEsi() bool {
	if o != nil && o.Esi != nil {
		return true
	}

	return false
}

// SetEsi gets a reference to the given string and assigns it to the Esi field.
func (o *EVPNRouteEVPNType2Route) SetEsi(v string) {
	o.Esi = &v
}

// GetEthTag returns the EthTag field value if set, zero value otherwise.
func (o *EVPNRouteEVPNType2Route) GetEthTag() string {
	if o == nil || o.EthTag == nil {
		var ret string
		return ret
	}
	return *o.EthTag
}

// GetEthTagOk returns a tuple with the EthTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EVPNRouteEVPNType2Route) GetEthTagOk() (*string, bool) {
	if o == nil || o.EthTag == nil {
		return nil, false
	}
	return o.EthTag, true
}

// HasEthTag returns a boolean if a field has been set.
func (o *EVPNRouteEVPNType2Route) HasEthTag() bool {
	if o != nil && o.EthTag != nil {
		return true
	}

	return false
}

// SetEthTag gets a reference to the given string and assigns it to the EthTag field.
func (o *EVPNRouteEVPNType2Route) SetEthTag(v string) {
	o.EthTag = &v
}

// GetIpAddr returns the IpAddr field value if set, zero value otherwise.
func (o *EVPNRouteEVPNType2Route) GetIpAddr() string {
	if o == nil || o.IpAddr == nil {
		var ret string
		return ret
	}
	return *o.IpAddr
}

// GetIpAddrOk returns a tuple with the IpAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EVPNRouteEVPNType2Route) GetIpAddrOk() (*string, bool) {
	if o == nil || o.IpAddr == nil {
		return nil, false
	}
	return o.IpAddr, true
}

// HasIpAddr returns a boolean if a field has been set.
func (o *EVPNRouteEVPNType2Route) HasIpAddr() bool {
	if o != nil && o.IpAddr != nil {
		return true
	}

	return false
}

// SetIpAddr gets a reference to the given string and assigns it to the IpAddr field.
func (o *EVPNRouteEVPNType2Route) SetIpAddr(v string) {
	o.IpAddr = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *EVPNRouteEVPNType2Route) GetMac() string {
	if o == nil || o.Mac == nil {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EVPNRouteEVPNType2Route) GetMacOk() (*string, bool) {
	if o == nil || o.Mac == nil {
		return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *EVPNRouteEVPNType2Route) HasMac() bool {
	if o != nil && o.Mac != nil {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *EVPNRouteEVPNType2Route) SetMac(v string) {
	o.Mac = &v
}

// GetMplsLabel1 returns the MplsLabel1 field value if set, zero value otherwise.
func (o *EVPNRouteEVPNType2Route) GetMplsLabel1() string {
	if o == nil || o.MplsLabel1 == nil {
		var ret string
		return ret
	}
	return *o.MplsLabel1
}

// GetMplsLabel1Ok returns a tuple with the MplsLabel1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EVPNRouteEVPNType2Route) GetMplsLabel1Ok() (*string, bool) {
	if o == nil || o.MplsLabel1 == nil {
		return nil, false
	}
	return o.MplsLabel1, true
}

// HasMplsLabel1 returns a boolean if a field has been set.
func (o *EVPNRouteEVPNType2Route) HasMplsLabel1() bool {
	if o != nil && o.MplsLabel1 != nil {
		return true
	}

	return false
}

// SetMplsLabel1 gets a reference to the given string and assigns it to the MplsLabel1 field.
func (o *EVPNRouteEVPNType2Route) SetMplsLabel1(v string) {
	o.MplsLabel1 = &v
}

// GetMplsLabel2 returns the MplsLabel2 field value if set, zero value otherwise.
func (o *EVPNRouteEVPNType2Route) GetMplsLabel2() string {
	if o == nil || o.MplsLabel2 == nil {
		var ret string
		return ret
	}
	return *o.MplsLabel2
}

// GetMplsLabel2Ok returns a tuple with the MplsLabel2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EVPNRouteEVPNType2Route) GetMplsLabel2Ok() (*string, bool) {
	if o == nil || o.MplsLabel2 == nil {
		return nil, false
	}
	return o.MplsLabel2, true
}

// HasMplsLabel2 returns a boolean if a field has been set.
func (o *EVPNRouteEVPNType2Route) HasMplsLabel2() bool {
	if o != nil && o.MplsLabel2 != nil {
		return true
	}

	return false
}

// SetMplsLabel2 gets a reference to the given string and assigns it to the MplsLabel2 field.
func (o *EVPNRouteEVPNType2Route) SetMplsLabel2(v string) {
	o.MplsLabel2 = &v
}

// GetRd returns the Rd field value if set, zero value otherwise.
func (o *EVPNRouteEVPNType2Route) GetRd() string {
	if o == nil || o.Rd == nil {
		var ret string
		return ret
	}
	return *o.Rd
}

// GetRdOk returns a tuple with the Rd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EVPNRouteEVPNType2Route) GetRdOk() (*string, bool) {
	if o == nil || o.Rd == nil {
		return nil, false
	}
	return o.Rd, true
}

// HasRd returns a boolean if a field has been set.
func (o *EVPNRouteEVPNType2Route) HasRd() bool {
	if o != nil && o.Rd != nil {
		return true
	}

	return false
}

// SetRd gets a reference to the given string and assigns it to the Rd field.
func (o *EVPNRouteEVPNType2Route) SetRd(v string) {
	o.Rd = &v
}

func (o EVPNRouteEVPNType2Route) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Esi != nil {
		toSerialize["esi"] = o.Esi
	}
	if o.EthTag != nil {
		toSerialize["eth-tag"] = o.EthTag
	}
	if o.IpAddr != nil {
		toSerialize["ip-addr"] = o.IpAddr
	}
	if o.Mac != nil {
		toSerialize["mac"] = o.Mac
	}
	if o.MplsLabel1 != nil {
		toSerialize["mpls-label1"] = o.MplsLabel1
	}
	if o.MplsLabel2 != nil {
		toSerialize["mpls-label2"] = o.MplsLabel2
	}
	if o.Rd != nil {
		toSerialize["rd"] = o.Rd
	}
	return json.Marshal(toSerialize)
}

type NullableEVPNRouteEVPNType2Route struct {
	value *EVPNRouteEVPNType2Route
	isSet bool
}

func (v NullableEVPNRouteEVPNType2Route) Get() *EVPNRouteEVPNType2Route {
	return v.value
}

func (v *NullableEVPNRouteEVPNType2Route) Set(val *EVPNRouteEVPNType2Route) {
	v.value = val
	v.isSet = true
}

func (v NullableEVPNRouteEVPNType2Route) IsSet() bool {
	return v.isSet
}

func (v *NullableEVPNRouteEVPNType2Route) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEVPNRouteEVPNType2Route(val *EVPNRouteEVPNType2Route) *NullableEVPNRouteEVPNType2Route {
	return &NullableEVPNRouteEVPNType2Route{value: val, isSet: true}
}

func (v NullableEVPNRouteEVPNType2Route) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEVPNRouteEVPNType2Route) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


