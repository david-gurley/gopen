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

// AuthAuthenticators struct for AuthAuthenticators
type AuthAuthenticators struct {
	// Order in which authenticators are applied. If an authenticator returns success, others are skipped.
	AuthenticatorOrder *[]string `json:"authenticator-order,omitempty"`
	Ldap *AuthLdap `json:"ldap,omitempty"`
	Local *AuthLocal `json:"local,omitempty"`
	Radius *AuthRadius `json:"radius,omitempty"`
}

// NewAuthAuthenticators instantiates a new AuthAuthenticators object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthAuthenticators() *AuthAuthenticators {
	this := AuthAuthenticators{}
	return &this
}

// NewAuthAuthenticatorsWithDefaults instantiates a new AuthAuthenticators object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthAuthenticatorsWithDefaults() *AuthAuthenticators {
	this := AuthAuthenticators{}
	return &this
}

// GetAuthenticatorOrder returns the AuthenticatorOrder field value if set, zero value otherwise.
func (o *AuthAuthenticators) GetAuthenticatorOrder() []string {
	if o == nil || o.AuthenticatorOrder == nil {
		var ret []string
		return ret
	}
	return *o.AuthenticatorOrder
}

// GetAuthenticatorOrderOk returns a tuple with the AuthenticatorOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthAuthenticators) GetAuthenticatorOrderOk() (*[]string, bool) {
	if o == nil || o.AuthenticatorOrder == nil {
		return nil, false
	}
	return o.AuthenticatorOrder, true
}

// HasAuthenticatorOrder returns a boolean if a field has been set.
func (o *AuthAuthenticators) HasAuthenticatorOrder() bool {
	if o != nil && o.AuthenticatorOrder != nil {
		return true
	}

	return false
}

// SetAuthenticatorOrder gets a reference to the given []string and assigns it to the AuthenticatorOrder field.
func (o *AuthAuthenticators) SetAuthenticatorOrder(v []string) {
	o.AuthenticatorOrder = &v
}

// GetLdap returns the Ldap field value if set, zero value otherwise.
func (o *AuthAuthenticators) GetLdap() AuthLdap {
	if o == nil || o.Ldap == nil {
		var ret AuthLdap
		return ret
	}
	return *o.Ldap
}

// GetLdapOk returns a tuple with the Ldap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthAuthenticators) GetLdapOk() (*AuthLdap, bool) {
	if o == nil || o.Ldap == nil {
		return nil, false
	}
	return o.Ldap, true
}

// HasLdap returns a boolean if a field has been set.
func (o *AuthAuthenticators) HasLdap() bool {
	if o != nil && o.Ldap != nil {
		return true
	}

	return false
}

// SetLdap gets a reference to the given AuthLdap and assigns it to the Ldap field.
func (o *AuthAuthenticators) SetLdap(v AuthLdap) {
	o.Ldap = &v
}

// GetLocal returns the Local field value if set, zero value otherwise.
func (o *AuthAuthenticators) GetLocal() AuthLocal {
	if o == nil || o.Local == nil {
		var ret AuthLocal
		return ret
	}
	return *o.Local
}

// GetLocalOk returns a tuple with the Local field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthAuthenticators) GetLocalOk() (*AuthLocal, bool) {
	if o == nil || o.Local == nil {
		return nil, false
	}
	return o.Local, true
}

// HasLocal returns a boolean if a field has been set.
func (o *AuthAuthenticators) HasLocal() bool {
	if o != nil && o.Local != nil {
		return true
	}

	return false
}

// SetLocal gets a reference to the given AuthLocal and assigns it to the Local field.
func (o *AuthAuthenticators) SetLocal(v AuthLocal) {
	o.Local = &v
}

// GetRadius returns the Radius field value if set, zero value otherwise.
func (o *AuthAuthenticators) GetRadius() AuthRadius {
	if o == nil || o.Radius == nil {
		var ret AuthRadius
		return ret
	}
	return *o.Radius
}

// GetRadiusOk returns a tuple with the Radius field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthAuthenticators) GetRadiusOk() (*AuthRadius, bool) {
	if o == nil || o.Radius == nil {
		return nil, false
	}
	return o.Radius, true
}

// HasRadius returns a boolean if a field has been set.
func (o *AuthAuthenticators) HasRadius() bool {
	if o != nil && o.Radius != nil {
		return true
	}

	return false
}

// SetRadius gets a reference to the given AuthRadius and assigns it to the Radius field.
func (o *AuthAuthenticators) SetRadius(v AuthRadius) {
	o.Radius = &v
}

func (o AuthAuthenticators) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthenticatorOrder != nil {
		toSerialize["authenticator-order"] = o.AuthenticatorOrder
	}
	if o.Ldap != nil {
		toSerialize["ldap"] = o.Ldap
	}
	if o.Local != nil {
		toSerialize["local"] = o.Local
	}
	if o.Radius != nil {
		toSerialize["radius"] = o.Radius
	}
	return json.Marshal(toSerialize)
}

type NullableAuthAuthenticators struct {
	value *AuthAuthenticators
	isSet bool
}

func (v NullableAuthAuthenticators) Get() *AuthAuthenticators {
	return v.value
}

func (v *NullableAuthAuthenticators) Set(val *AuthAuthenticators) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthAuthenticators) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthAuthenticators) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthAuthenticators(val *AuthAuthenticators) *NullableAuthAuthenticators {
	return &NullableAuthAuthenticators{value: val, isSet: true}
}

func (v NullableAuthAuthenticators) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthAuthenticators) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


