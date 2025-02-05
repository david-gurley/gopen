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

// SecurityTLSProtocolSpec struct for SecurityTLSProtocolSpec
type SecurityTLSProtocolSpec struct {
	// The name of the cipher suite in IANA format default is TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384.
	CipherSuite *string `json:"cipher-suite,omitempty"`
	// TLS version: only supported value at present is 1.2.
	Version *string `json:"version,omitempty"`
}

// NewSecurityTLSProtocolSpec instantiates a new SecurityTLSProtocolSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityTLSProtocolSpec() *SecurityTLSProtocolSpec {
	this := SecurityTLSProtocolSpec{}
	return &this
}

// NewSecurityTLSProtocolSpecWithDefaults instantiates a new SecurityTLSProtocolSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityTLSProtocolSpecWithDefaults() *SecurityTLSProtocolSpec {
	this := SecurityTLSProtocolSpec{}
	return &this
}

// GetCipherSuite returns the CipherSuite field value if set, zero value otherwise.
func (o *SecurityTLSProtocolSpec) GetCipherSuite() string {
	if o == nil || o.CipherSuite == nil {
		var ret string
		return ret
	}
	return *o.CipherSuite
}

// GetCipherSuiteOk returns a tuple with the CipherSuite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityTLSProtocolSpec) GetCipherSuiteOk() (*string, bool) {
	if o == nil || o.CipherSuite == nil {
		return nil, false
	}
	return o.CipherSuite, true
}

// HasCipherSuite returns a boolean if a field has been set.
func (o *SecurityTLSProtocolSpec) HasCipherSuite() bool {
	if o != nil && o.CipherSuite != nil {
		return true
	}

	return false
}

// SetCipherSuite gets a reference to the given string and assigns it to the CipherSuite field.
func (o *SecurityTLSProtocolSpec) SetCipherSuite(v string) {
	o.CipherSuite = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SecurityTLSProtocolSpec) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityTLSProtocolSpec) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SecurityTLSProtocolSpec) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *SecurityTLSProtocolSpec) SetVersion(v string) {
	o.Version = &v
}

func (o SecurityTLSProtocolSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CipherSuite != nil {
		toSerialize["cipher-suite"] = o.CipherSuite
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableSecurityTLSProtocolSpec struct {
	value *SecurityTLSProtocolSpec
	isSet bool
}

func (v NullableSecurityTLSProtocolSpec) Get() *SecurityTLSProtocolSpec {
	return v.value
}

func (v *NullableSecurityTLSProtocolSpec) Set(val *SecurityTLSProtocolSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityTLSProtocolSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityTLSProtocolSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityTLSProtocolSpec(val *SecurityTLSProtocolSpec) *NullableSecurityTLSProtocolSpec {
	return &NullableSecurityTLSProtocolSpec{value: val, isSet: true}
}

func (v NullableSecurityTLSProtocolSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityTLSProtocolSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


