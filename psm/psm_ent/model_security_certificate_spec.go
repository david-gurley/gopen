/*
 * Security API reference
 *
 *  
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package psm_ent

import (
	"encoding/json"
)

// SecurityCertificateSpec struct for SecurityCertificateSpec
type SecurityCertificateSpec struct {
	// Body of the certificate in PEM encoding.
	Body *string `json:"body,omitempty"`
	// Description of the purpose of this certificate.
	Description *string `json:"description,omitempty"`
	// Trust chain of the certificate in PEM encoding. These certificates are treated opaquely. We do not process them in any way other than decoding them for informational purposes.
	TrustChain *string `json:"trust-chain,omitempty"`
	// Usage can be \"client\", \"server\" or \"trust-root\" in any combination. A \"server\" certificate is used by a server to authenticate itself to the client A \"client\" certificate is used by a client to authenticate itself to a server A \"trust-root\" certificate is self-signed and is only used to validate certificates presented by peers. \"client\" and \"server\" certificates are always accompanied by a private key, whereas \"trust-root\"-only certificates are not.
	Usages *[]string `json:"usages,omitempty"`
}

// NewSecurityCertificateSpec instantiates a new SecurityCertificateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityCertificateSpec() *SecurityCertificateSpec {
	this := SecurityCertificateSpec{}
	return &this
}

// NewSecurityCertificateSpecWithDefaults instantiates a new SecurityCertificateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityCertificateSpecWithDefaults() *SecurityCertificateSpec {
	this := SecurityCertificateSpec{}
	return &this
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *SecurityCertificateSpec) GetBody() string {
	if o == nil || o.Body == nil {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityCertificateSpec) GetBodyOk() (*string, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *SecurityCertificateSpec) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *SecurityCertificateSpec) SetBody(v string) {
	o.Body = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SecurityCertificateSpec) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityCertificateSpec) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SecurityCertificateSpec) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SecurityCertificateSpec) SetDescription(v string) {
	o.Description = &v
}

// GetTrustChain returns the TrustChain field value if set, zero value otherwise.
func (o *SecurityCertificateSpec) GetTrustChain() string {
	if o == nil || o.TrustChain == nil {
		var ret string
		return ret
	}
	return *o.TrustChain
}

// GetTrustChainOk returns a tuple with the TrustChain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityCertificateSpec) GetTrustChainOk() (*string, bool) {
	if o == nil || o.TrustChain == nil {
		return nil, false
	}
	return o.TrustChain, true
}

// HasTrustChain returns a boolean if a field has been set.
func (o *SecurityCertificateSpec) HasTrustChain() bool {
	if o != nil && o.TrustChain != nil {
		return true
	}

	return false
}

// SetTrustChain gets a reference to the given string and assigns it to the TrustChain field.
func (o *SecurityCertificateSpec) SetTrustChain(v string) {
	o.TrustChain = &v
}

// GetUsages returns the Usages field value if set, zero value otherwise.
func (o *SecurityCertificateSpec) GetUsages() []string {
	if o == nil || o.Usages == nil {
		var ret []string
		return ret
	}
	return *o.Usages
}

// GetUsagesOk returns a tuple with the Usages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityCertificateSpec) GetUsagesOk() (*[]string, bool) {
	if o == nil || o.Usages == nil {
		return nil, false
	}
	return o.Usages, true
}

// HasUsages returns a boolean if a field has been set.
func (o *SecurityCertificateSpec) HasUsages() bool {
	if o != nil && o.Usages != nil {
		return true
	}

	return false
}

// SetUsages gets a reference to the given []string and assigns it to the Usages field.
func (o *SecurityCertificateSpec) SetUsages(v []string) {
	o.Usages = &v
}

func (o SecurityCertificateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Body != nil {
		toSerialize["body"] = o.Body
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.TrustChain != nil {
		toSerialize["trust-chain"] = o.TrustChain
	}
	if o.Usages != nil {
		toSerialize["usages"] = o.Usages
	}
	return json.Marshal(toSerialize)
}

type NullableSecurityCertificateSpec struct {
	value *SecurityCertificateSpec
	isSet bool
}

func (v NullableSecurityCertificateSpec) Get() *SecurityCertificateSpec {
	return v.value
}

func (v *NullableSecurityCertificateSpec) Set(val *SecurityCertificateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityCertificateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityCertificateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityCertificateSpec(val *SecurityCertificateSpec) *NullableSecurityCertificateSpec {
	return &NullableSecurityCertificateSpec{value: val, isSet: true}
}

func (v NullableSecurityCertificateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityCertificateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


