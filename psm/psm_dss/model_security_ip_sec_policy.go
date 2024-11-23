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

// SecurityIPSecPolicy IPSecPolicy represents an IPSec encryption policy.
type SecurityIPSecPolicy struct {
	ApiVersion *string `json:"api-version,omitempty"`
	Kind *string `json:"kind,omitempty"`
	Meta *ApiObjectMeta `json:"meta,omitempty"`
	Spec *SecurityIPSecPolicySpec `json:"spec,omitempty"`
	Status *SecurityIPSecPolicyStatus `json:"status,omitempty"`
}

// NewSecurityIPSecPolicy instantiates a new SecurityIPSecPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityIPSecPolicy() *SecurityIPSecPolicy {
	this := SecurityIPSecPolicy{}
	return &this
}

// NewSecurityIPSecPolicyWithDefaults instantiates a new SecurityIPSecPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityIPSecPolicyWithDefaults() *SecurityIPSecPolicy {
	this := SecurityIPSecPolicy{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *SecurityIPSecPolicy) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityIPSecPolicy) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *SecurityIPSecPolicy) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *SecurityIPSecPolicy) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *SecurityIPSecPolicy) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityIPSecPolicy) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *SecurityIPSecPolicy) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *SecurityIPSecPolicy) SetKind(v string) {
	o.Kind = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SecurityIPSecPolicy) GetMeta() ApiObjectMeta {
	if o == nil || o.Meta == nil {
		var ret ApiObjectMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityIPSecPolicy) GetMetaOk() (*ApiObjectMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SecurityIPSecPolicy) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ApiObjectMeta and assigns it to the Meta field.
func (o *SecurityIPSecPolicy) SetMeta(v ApiObjectMeta) {
	o.Meta = &v
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *SecurityIPSecPolicy) GetSpec() SecurityIPSecPolicySpec {
	if o == nil || o.Spec == nil {
		var ret SecurityIPSecPolicySpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityIPSecPolicy) GetSpecOk() (*SecurityIPSecPolicySpec, bool) {
	if o == nil || o.Spec == nil {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *SecurityIPSecPolicy) HasSpec() bool {
	if o != nil && o.Spec != nil {
		return true
	}

	return false
}

// SetSpec gets a reference to the given SecurityIPSecPolicySpec and assigns it to the Spec field.
func (o *SecurityIPSecPolicy) SetSpec(v SecurityIPSecPolicySpec) {
	o.Spec = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SecurityIPSecPolicy) GetStatus() SecurityIPSecPolicyStatus {
	if o == nil || o.Status == nil {
		var ret SecurityIPSecPolicyStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityIPSecPolicy) GetStatusOk() (*SecurityIPSecPolicyStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SecurityIPSecPolicy) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given SecurityIPSecPolicyStatus and assigns it to the Status field.
func (o *SecurityIPSecPolicy) SetStatus(v SecurityIPSecPolicyStatus) {
	o.Status = &v
}

func (o SecurityIPSecPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiVersion != nil {
		toSerialize["api-version"] = o.ApiVersion
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.Spec != nil {
		toSerialize["spec"] = o.Spec
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableSecurityIPSecPolicy struct {
	value *SecurityIPSecPolicy
	isSet bool
}

func (v NullableSecurityIPSecPolicy) Get() *SecurityIPSecPolicy {
	return v.value
}

func (v *NullableSecurityIPSecPolicy) Set(val *SecurityIPSecPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityIPSecPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityIPSecPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityIPSecPolicy(val *SecurityIPSecPolicy) *NullableSecurityIPSecPolicy {
	return &NullableSecurityIPSecPolicy{value: val, isSet: true}
}

func (v NullableSecurityIPSecPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityIPSecPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

