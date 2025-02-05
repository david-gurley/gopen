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

// AuthRadius struct for AuthRadius
type AuthRadius struct {
	Domains *[]AuthRadiusDomain `json:"domains,omitempty"`
}

// NewAuthRadius instantiates a new AuthRadius object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthRadius() *AuthRadius {
	this := AuthRadius{}
	return &this
}

// NewAuthRadiusWithDefaults instantiates a new AuthRadius object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthRadiusWithDefaults() *AuthRadius {
	this := AuthRadius{}
	return &this
}

// GetDomains returns the Domains field value if set, zero value otherwise.
func (o *AuthRadius) GetDomains() []AuthRadiusDomain {
	if o == nil || o.Domains == nil {
		var ret []AuthRadiusDomain
		return ret
	}
	return *o.Domains
}

// GetDomainsOk returns a tuple with the Domains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthRadius) GetDomainsOk() (*[]AuthRadiusDomain, bool) {
	if o == nil || o.Domains == nil {
		return nil, false
	}
	return o.Domains, true
}

// HasDomains returns a boolean if a field has been set.
func (o *AuthRadius) HasDomains() bool {
	if o != nil && o.Domains != nil {
		return true
	}

	return false
}

// SetDomains gets a reference to the given []AuthRadiusDomain and assigns it to the Domains field.
func (o *AuthRadius) SetDomains(v []AuthRadiusDomain) {
	o.Domains = &v
}

func (o AuthRadius) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Domains != nil {
		toSerialize["domains"] = o.Domains
	}
	return json.Marshal(toSerialize)
}

type NullableAuthRadius struct {
	value *AuthRadius
	isSet bool
}

func (v NullableAuthRadius) Get() *AuthRadius {
	return v.value
}

func (v *NullableAuthRadius) Set(val *AuthRadius) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthRadius) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthRadius) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthRadius(val *AuthRadius) *NullableAuthRadius {
	return &NullableAuthRadius{value: val, isSet: true}
}

func (v NullableAuthRadius) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthRadius) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


