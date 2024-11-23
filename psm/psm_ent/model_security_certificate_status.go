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

// SecurityCertificateStatus struct for SecurityCertificateStatus
type SecurityCertificateStatus struct {
	// Status of the certificate: \"valid\", \"invalid\", \"expired\" \"invalid\" means that the signature of the certificate does not match or there are inconsistencies in the trust chain.
	Validity *string `json:"validity,omitempty"`
	// The workloads where this certificate has been deployed.
	Workloads *[]string `json:"workloads,omitempty"`
}

// NewSecurityCertificateStatus instantiates a new SecurityCertificateStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityCertificateStatus() *SecurityCertificateStatus {
	this := SecurityCertificateStatus{}
	var validity string = "unknown"
	this.Validity = &validity
	return &this
}

// NewSecurityCertificateStatusWithDefaults instantiates a new SecurityCertificateStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityCertificateStatusWithDefaults() *SecurityCertificateStatus {
	this := SecurityCertificateStatus{}
	var validity string = "unknown"
	this.Validity = &validity
	return &this
}

// GetValidity returns the Validity field value if set, zero value otherwise.
func (o *SecurityCertificateStatus) GetValidity() string {
	if o == nil || o.Validity == nil {
		var ret string
		return ret
	}
	return *o.Validity
}

// GetValidityOk returns a tuple with the Validity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityCertificateStatus) GetValidityOk() (*string, bool) {
	if o == nil || o.Validity == nil {
		return nil, false
	}
	return o.Validity, true
}

// HasValidity returns a boolean if a field has been set.
func (o *SecurityCertificateStatus) HasValidity() bool {
	if o != nil && o.Validity != nil {
		return true
	}

	return false
}

// SetValidity gets a reference to the given string and assigns it to the Validity field.
func (o *SecurityCertificateStatus) SetValidity(v string) {
	o.Validity = &v
}

// GetWorkloads returns the Workloads field value if set, zero value otherwise.
func (o *SecurityCertificateStatus) GetWorkloads() []string {
	if o == nil || o.Workloads == nil {
		var ret []string
		return ret
	}
	return *o.Workloads
}

// GetWorkloadsOk returns a tuple with the Workloads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityCertificateStatus) GetWorkloadsOk() (*[]string, bool) {
	if o == nil || o.Workloads == nil {
		return nil, false
	}
	return o.Workloads, true
}

// HasWorkloads returns a boolean if a field has been set.
func (o *SecurityCertificateStatus) HasWorkloads() bool {
	if o != nil && o.Workloads != nil {
		return true
	}

	return false
}

// SetWorkloads gets a reference to the given []string and assigns it to the Workloads field.
func (o *SecurityCertificateStatus) SetWorkloads(v []string) {
	o.Workloads = &v
}

func (o SecurityCertificateStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Validity != nil {
		toSerialize["validity"] = o.Validity
	}
	if o.Workloads != nil {
		toSerialize["workloads"] = o.Workloads
	}
	return json.Marshal(toSerialize)
}

type NullableSecurityCertificateStatus struct {
	value *SecurityCertificateStatus
	isSet bool
}

func (v NullableSecurityCertificateStatus) Get() *SecurityCertificateStatus {
	return v.value
}

func (v *NullableSecurityCertificateStatus) Set(val *SecurityCertificateStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityCertificateStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityCertificateStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityCertificateStatus(val *SecurityCertificateStatus) *NullableSecurityCertificateStatus {
	return &NullableSecurityCertificateStatus{value: val, isSet: true}
}

func (v NullableSecurityCertificateStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityCertificateStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


