/*
 * Workload API reference
 *
 * Service name  
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package psm_ent

import (
	"encoding/json"
)

// WorkloadWorkload Workload represents a VM, container/pod or Baremetal.
type WorkloadWorkload struct {
	ApiVersion *string `json:"api-version,omitempty"`
	Kind *string `json:"kind,omitempty"`
	Meta *ApiObjectMeta `json:"meta,omitempty"`
	Spec *WorkloadWorkloadSpec `json:"spec,omitempty"`
	Status *WorkloadWorkloadStatus `json:"status,omitempty"`
}

// NewWorkloadWorkload instantiates a new WorkloadWorkload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkloadWorkload() *WorkloadWorkload {
	this := WorkloadWorkload{}
	return &this
}

// NewWorkloadWorkloadWithDefaults instantiates a new WorkloadWorkload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkloadWorkloadWithDefaults() *WorkloadWorkload {
	this := WorkloadWorkload{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *WorkloadWorkload) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadWorkload) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *WorkloadWorkload) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *WorkloadWorkload) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *WorkloadWorkload) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadWorkload) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *WorkloadWorkload) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *WorkloadWorkload) SetKind(v string) {
	o.Kind = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *WorkloadWorkload) GetMeta() ApiObjectMeta {
	if o == nil || o.Meta == nil {
		var ret ApiObjectMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadWorkload) GetMetaOk() (*ApiObjectMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *WorkloadWorkload) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ApiObjectMeta and assigns it to the Meta field.
func (o *WorkloadWorkload) SetMeta(v ApiObjectMeta) {
	o.Meta = &v
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *WorkloadWorkload) GetSpec() WorkloadWorkloadSpec {
	if o == nil || o.Spec == nil {
		var ret WorkloadWorkloadSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadWorkload) GetSpecOk() (*WorkloadWorkloadSpec, bool) {
	if o == nil || o.Spec == nil {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *WorkloadWorkload) HasSpec() bool {
	if o != nil && o.Spec != nil {
		return true
	}

	return false
}

// SetSpec gets a reference to the given WorkloadWorkloadSpec and assigns it to the Spec field.
func (o *WorkloadWorkload) SetSpec(v WorkloadWorkloadSpec) {
	o.Spec = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WorkloadWorkload) GetStatus() WorkloadWorkloadStatus {
	if o == nil || o.Status == nil {
		var ret WorkloadWorkloadStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadWorkload) GetStatusOk() (*WorkloadWorkloadStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WorkloadWorkload) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given WorkloadWorkloadStatus and assigns it to the Status field.
func (o *WorkloadWorkload) SetStatus(v WorkloadWorkloadStatus) {
	o.Status = &v
}

func (o WorkloadWorkload) MarshalJSON() ([]byte, error) {
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

type NullableWorkloadWorkload struct {
	value *WorkloadWorkload
	isSet bool
}

func (v NullableWorkloadWorkload) Get() *WorkloadWorkload {
	return v.value
}

func (v *NullableWorkloadWorkload) Set(val *WorkloadWorkload) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkloadWorkload) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkloadWorkload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkloadWorkload(val *WorkloadWorkload) *NullableWorkloadWorkload {
	return &NullableWorkloadWorkload{value: val, isSet: true}
}

func (v NullableWorkloadWorkload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkloadWorkload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


