/*
 * Auth API reference
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

// AuthLdapServer struct for AuthLdapServer
type AuthLdapServer struct {
	TlsOptions *AuthTLSOptions `json:"tls-options,omitempty"`
	Url *string `json:"url,omitempty"`
}

// NewAuthLdapServer instantiates a new AuthLdapServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthLdapServer() *AuthLdapServer {
	this := AuthLdapServer{}
	return &this
}

// NewAuthLdapServerWithDefaults instantiates a new AuthLdapServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthLdapServerWithDefaults() *AuthLdapServer {
	this := AuthLdapServer{}
	return &this
}

// GetTlsOptions returns the TlsOptions field value if set, zero value otherwise.
func (o *AuthLdapServer) GetTlsOptions() AuthTLSOptions {
	if o == nil || o.TlsOptions == nil {
		var ret AuthTLSOptions
		return ret
	}
	return *o.TlsOptions
}

// GetTlsOptionsOk returns a tuple with the TlsOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthLdapServer) GetTlsOptionsOk() (*AuthTLSOptions, bool) {
	if o == nil || o.TlsOptions == nil {
		return nil, false
	}
	return o.TlsOptions, true
}

// HasTlsOptions returns a boolean if a field has been set.
func (o *AuthLdapServer) HasTlsOptions() bool {
	if o != nil && o.TlsOptions != nil {
		return true
	}

	return false
}

// SetTlsOptions gets a reference to the given AuthTLSOptions and assigns it to the TlsOptions field.
func (o *AuthLdapServer) SetTlsOptions(v AuthTLSOptions) {
	o.TlsOptions = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AuthLdapServer) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthLdapServer) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AuthLdapServer) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AuthLdapServer) SetUrl(v string) {
	o.Url = &v
}

func (o AuthLdapServer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TlsOptions != nil {
		toSerialize["tls-options"] = o.TlsOptions
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableAuthLdapServer struct {
	value *AuthLdapServer
	isSet bool
}

func (v NullableAuthLdapServer) Get() *AuthLdapServer {
	return v.value
}

func (v *NullableAuthLdapServer) Set(val *AuthLdapServer) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthLdapServer) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthLdapServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthLdapServer(val *AuthLdapServer) *NullableAuthLdapServer {
	return &NullableAuthLdapServer{value: val, isSet: true}
}

func (v NullableAuthLdapServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthLdapServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

