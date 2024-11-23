/*
 * Security API reference
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

// SecurityALG ALG specifies the application specific configuration for the list of applications mentioned below.
type SecurityALG struct {
	Dns *SecurityDns `json:"dns,omitempty"`
	Ftp *SecurityFtp `json:"ftp,omitempty"`
	Icmp *SecurityIcmp `json:"icmp,omitempty"`
	Msrpc *[]SecurityMsrpc `json:"msrpc,omitempty"`
	Sunrpc *[]SecuritySunrpc `json:"sunrpc,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewSecurityALG instantiates a new SecurityALG object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityALG() *SecurityALG {
	this := SecurityALG{}
	var type_ string = "icmp"
	this.Type = &type_
	return &this
}

// NewSecurityALGWithDefaults instantiates a new SecurityALG object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityALGWithDefaults() *SecurityALG {
	this := SecurityALG{}
	var type_ string = "icmp"
	this.Type = &type_
	return &this
}

// GetDns returns the Dns field value if set, zero value otherwise.
func (o *SecurityALG) GetDns() SecurityDns {
	if o == nil || o.Dns == nil {
		var ret SecurityDns
		return ret
	}
	return *o.Dns
}

// GetDnsOk returns a tuple with the Dns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityALG) GetDnsOk() (*SecurityDns, bool) {
	if o == nil || o.Dns == nil {
		return nil, false
	}
	return o.Dns, true
}

// HasDns returns a boolean if a field has been set.
func (o *SecurityALG) HasDns() bool {
	if o != nil && o.Dns != nil {
		return true
	}

	return false
}

// SetDns gets a reference to the given SecurityDns and assigns it to the Dns field.
func (o *SecurityALG) SetDns(v SecurityDns) {
	o.Dns = &v
}

// GetFtp returns the Ftp field value if set, zero value otherwise.
func (o *SecurityALG) GetFtp() SecurityFtp {
	if o == nil || o.Ftp == nil {
		var ret SecurityFtp
		return ret
	}
	return *o.Ftp
}

// GetFtpOk returns a tuple with the Ftp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityALG) GetFtpOk() (*SecurityFtp, bool) {
	if o == nil || o.Ftp == nil {
		return nil, false
	}
	return o.Ftp, true
}

// HasFtp returns a boolean if a field has been set.
func (o *SecurityALG) HasFtp() bool {
	if o != nil && o.Ftp != nil {
		return true
	}

	return false
}

// SetFtp gets a reference to the given SecurityFtp and assigns it to the Ftp field.
func (o *SecurityALG) SetFtp(v SecurityFtp) {
	o.Ftp = &v
}

// GetIcmp returns the Icmp field value if set, zero value otherwise.
func (o *SecurityALG) GetIcmp() SecurityIcmp {
	if o == nil || o.Icmp == nil {
		var ret SecurityIcmp
		return ret
	}
	return *o.Icmp
}

// GetIcmpOk returns a tuple with the Icmp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityALG) GetIcmpOk() (*SecurityIcmp, bool) {
	if o == nil || o.Icmp == nil {
		return nil, false
	}
	return o.Icmp, true
}

// HasIcmp returns a boolean if a field has been set.
func (o *SecurityALG) HasIcmp() bool {
	if o != nil && o.Icmp != nil {
		return true
	}

	return false
}

// SetIcmp gets a reference to the given SecurityIcmp and assigns it to the Icmp field.
func (o *SecurityALG) SetIcmp(v SecurityIcmp) {
	o.Icmp = &v
}

// GetMsrpc returns the Msrpc field value if set, zero value otherwise.
func (o *SecurityALG) GetMsrpc() []SecurityMsrpc {
	if o == nil || o.Msrpc == nil {
		var ret []SecurityMsrpc
		return ret
	}
	return *o.Msrpc
}

// GetMsrpcOk returns a tuple with the Msrpc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityALG) GetMsrpcOk() (*[]SecurityMsrpc, bool) {
	if o == nil || o.Msrpc == nil {
		return nil, false
	}
	return o.Msrpc, true
}

// HasMsrpc returns a boolean if a field has been set.
func (o *SecurityALG) HasMsrpc() bool {
	if o != nil && o.Msrpc != nil {
		return true
	}

	return false
}

// SetMsrpc gets a reference to the given []SecurityMsrpc and assigns it to the Msrpc field.
func (o *SecurityALG) SetMsrpc(v []SecurityMsrpc) {
	o.Msrpc = &v
}

// GetSunrpc returns the Sunrpc field value if set, zero value otherwise.
func (o *SecurityALG) GetSunrpc() []SecuritySunrpc {
	if o == nil || o.Sunrpc == nil {
		var ret []SecuritySunrpc
		return ret
	}
	return *o.Sunrpc
}

// GetSunrpcOk returns a tuple with the Sunrpc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityALG) GetSunrpcOk() (*[]SecuritySunrpc, bool) {
	if o == nil || o.Sunrpc == nil {
		return nil, false
	}
	return o.Sunrpc, true
}

// HasSunrpc returns a boolean if a field has been set.
func (o *SecurityALG) HasSunrpc() bool {
	if o != nil && o.Sunrpc != nil {
		return true
	}

	return false
}

// SetSunrpc gets a reference to the given []SecuritySunrpc and assigns it to the Sunrpc field.
func (o *SecurityALG) SetSunrpc(v []SecuritySunrpc) {
	o.Sunrpc = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SecurityALG) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityALG) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SecurityALG) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SecurityALG) SetType(v string) {
	o.Type = &v
}

func (o SecurityALG) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Dns != nil {
		toSerialize["dns"] = o.Dns
	}
	if o.Ftp != nil {
		toSerialize["ftp"] = o.Ftp
	}
	if o.Icmp != nil {
		toSerialize["icmp"] = o.Icmp
	}
	if o.Msrpc != nil {
		toSerialize["msrpc"] = o.Msrpc
	}
	if o.Sunrpc != nil {
		toSerialize["sunrpc"] = o.Sunrpc
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableSecurityALG struct {
	value *SecurityALG
	isSet bool
}

func (v NullableSecurityALG) Get() *SecurityALG {
	return v.value
}

func (v *NullableSecurityALG) Set(val *SecurityALG) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityALG) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityALG) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityALG(val *SecurityALG) *NullableSecurityALG {
	return &NullableSecurityALG{value: val, isSet: true}
}

func (v NullableSecurityALG) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityALG) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


