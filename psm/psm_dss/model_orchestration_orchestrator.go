/*
 * Orchestration API reference
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

// OrchestrationOrchestrator --------------------------------- ORCHESTRATOR --------------------------------------------- Orchestrator represents the config object which allows Venice to connect to the appropriate orchestrator.
type OrchestrationOrchestrator struct {
	ApiVersion *string `json:"api-version,omitempty"`
	Kind *string `json:"kind,omitempty"`
	Meta *ApiObjectMeta `json:"meta,omitempty"`
	Spec *OrchestrationOrchestratorSpec `json:"spec,omitempty"`
	Status *OrchestrationOrchestratorStatus `json:"status,omitempty"`
}

// NewOrchestrationOrchestrator instantiates a new OrchestrationOrchestrator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrchestrationOrchestrator() *OrchestrationOrchestrator {
	this := OrchestrationOrchestrator{}
	return &this
}

// NewOrchestrationOrchestratorWithDefaults instantiates a new OrchestrationOrchestrator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrchestrationOrchestratorWithDefaults() *OrchestrationOrchestrator {
	this := OrchestrationOrchestrator{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *OrchestrationOrchestrator) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrchestrationOrchestrator) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *OrchestrationOrchestrator) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *OrchestrationOrchestrator) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *OrchestrationOrchestrator) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrchestrationOrchestrator) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *OrchestrationOrchestrator) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *OrchestrationOrchestrator) SetKind(v string) {
	o.Kind = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *OrchestrationOrchestrator) GetMeta() ApiObjectMeta {
	if o == nil || o.Meta == nil {
		var ret ApiObjectMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrchestrationOrchestrator) GetMetaOk() (*ApiObjectMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *OrchestrationOrchestrator) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ApiObjectMeta and assigns it to the Meta field.
func (o *OrchestrationOrchestrator) SetMeta(v ApiObjectMeta) {
	o.Meta = &v
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *OrchestrationOrchestrator) GetSpec() OrchestrationOrchestratorSpec {
	if o == nil || o.Spec == nil {
		var ret OrchestrationOrchestratorSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrchestrationOrchestrator) GetSpecOk() (*OrchestrationOrchestratorSpec, bool) {
	if o == nil || o.Spec == nil {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *OrchestrationOrchestrator) HasSpec() bool {
	if o != nil && o.Spec != nil {
		return true
	}

	return false
}

// SetSpec gets a reference to the given OrchestrationOrchestratorSpec and assigns it to the Spec field.
func (o *OrchestrationOrchestrator) SetSpec(v OrchestrationOrchestratorSpec) {
	o.Spec = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OrchestrationOrchestrator) GetStatus() OrchestrationOrchestratorStatus {
	if o == nil || o.Status == nil {
		var ret OrchestrationOrchestratorStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrchestrationOrchestrator) GetStatusOk() (*OrchestrationOrchestratorStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OrchestrationOrchestrator) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given OrchestrationOrchestratorStatus and assigns it to the Status field.
func (o *OrchestrationOrchestrator) SetStatus(v OrchestrationOrchestratorStatus) {
	o.Status = &v
}

func (o OrchestrationOrchestrator) MarshalJSON() ([]byte, error) {
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

type NullableOrchestrationOrchestrator struct {
	value *OrchestrationOrchestrator
	isSet bool
}

func (v NullableOrchestrationOrchestrator) Get() *OrchestrationOrchestrator {
	return v.value
}

func (v *NullableOrchestrationOrchestrator) Set(val *OrchestrationOrchestrator) {
	v.value = val
	v.isSet = true
}

func (v NullableOrchestrationOrchestrator) IsSet() bool {
	return v.isSet
}

func (v *NullableOrchestrationOrchestrator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrchestrationOrchestrator(val *OrchestrationOrchestrator) *NullableOrchestrationOrchestrator {
	return &NullableOrchestrationOrchestrator{value: val, isSet: true}
}

func (v NullableOrchestrationOrchestrator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrchestrationOrchestrator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


