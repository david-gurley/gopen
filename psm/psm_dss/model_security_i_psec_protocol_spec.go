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

// SecurityIPsecProtocolSpec struct for SecurityIPsecProtocolSpec
type SecurityIPsecProtocolSpec struct {
	// ESP encryption algorithm. Default is \"aes-256-gcm-128\" (See RFC4106).
	EncryptionTransform *string `json:"encryption-transform,omitempty"`
	// ESP integrity algorithm. Default is \"NULL\" (must be \"NULL\" if AES-GCM is used for encryption).
	IntegrityTransform *string `json:"integrity-transform,omitempty"`
}

// NewSecurityIPsecProtocolSpec instantiates a new SecurityIPsecProtocolSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityIPsecProtocolSpec() *SecurityIPsecProtocolSpec {
	this := SecurityIPsecProtocolSpec{}
	return &this
}

// NewSecurityIPsecProtocolSpecWithDefaults instantiates a new SecurityIPsecProtocolSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityIPsecProtocolSpecWithDefaults() *SecurityIPsecProtocolSpec {
	this := SecurityIPsecProtocolSpec{}
	return &this
}

// GetEncryptionTransform returns the EncryptionTransform field value if set, zero value otherwise.
func (o *SecurityIPsecProtocolSpec) GetEncryptionTransform() string {
	if o == nil || o.EncryptionTransform == nil {
		var ret string
		return ret
	}
	return *o.EncryptionTransform
}

// GetEncryptionTransformOk returns a tuple with the EncryptionTransform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityIPsecProtocolSpec) GetEncryptionTransformOk() (*string, bool) {
	if o == nil || o.EncryptionTransform == nil {
		return nil, false
	}
	return o.EncryptionTransform, true
}

// HasEncryptionTransform returns a boolean if a field has been set.
func (o *SecurityIPsecProtocolSpec) HasEncryptionTransform() bool {
	if o != nil && o.EncryptionTransform != nil {
		return true
	}

	return false
}

// SetEncryptionTransform gets a reference to the given string and assigns it to the EncryptionTransform field.
func (o *SecurityIPsecProtocolSpec) SetEncryptionTransform(v string) {
	o.EncryptionTransform = &v
}

// GetIntegrityTransform returns the IntegrityTransform field value if set, zero value otherwise.
func (o *SecurityIPsecProtocolSpec) GetIntegrityTransform() string {
	if o == nil || o.IntegrityTransform == nil {
		var ret string
		return ret
	}
	return *o.IntegrityTransform
}

// GetIntegrityTransformOk returns a tuple with the IntegrityTransform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityIPsecProtocolSpec) GetIntegrityTransformOk() (*string, bool) {
	if o == nil || o.IntegrityTransform == nil {
		return nil, false
	}
	return o.IntegrityTransform, true
}

// HasIntegrityTransform returns a boolean if a field has been set.
func (o *SecurityIPsecProtocolSpec) HasIntegrityTransform() bool {
	if o != nil && o.IntegrityTransform != nil {
		return true
	}

	return false
}

// SetIntegrityTransform gets a reference to the given string and assigns it to the IntegrityTransform field.
func (o *SecurityIPsecProtocolSpec) SetIntegrityTransform(v string) {
	o.IntegrityTransform = &v
}

func (o SecurityIPsecProtocolSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EncryptionTransform != nil {
		toSerialize["encryption-transform"] = o.EncryptionTransform
	}
	if o.IntegrityTransform != nil {
		toSerialize["integrity-transform"] = o.IntegrityTransform
	}
	return json.Marshal(toSerialize)
}

type NullableSecurityIPsecProtocolSpec struct {
	value *SecurityIPsecProtocolSpec
	isSet bool
}

func (v NullableSecurityIPsecProtocolSpec) Get() *SecurityIPsecProtocolSpec {
	return v.value
}

func (v *NullableSecurityIPsecProtocolSpec) Set(val *SecurityIPsecProtocolSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityIPsecProtocolSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityIPsecProtocolSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityIPsecProtocolSpec(val *SecurityIPsecProtocolSpec) *NullableSecurityIPsecProtocolSpec {
	return &NullableSecurityIPsecProtocolSpec{value: val, isSet: true}
}

func (v NullableSecurityIPsecProtocolSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityIPsecProtocolSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


