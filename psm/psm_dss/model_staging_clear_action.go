/*
 * Staging API reference
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

// StagingClearAction ClearAction deletes objects from the staging buffer. A list of items to be cleared is specified in the Spec. If there are no items are specified then all items are deleted from the staging. buffer.
type StagingClearAction struct {
	ApiVersion *string `json:"api-version,omitempty"`
	Kind *string `json:"kind,omitempty"`
	Meta *ApiObjectMeta `json:"meta,omitempty"`
	Spec *StagingClearActionSpec `json:"spec,omitempty"`
	Status *StagingClearActionStatus `json:"status,omitempty"`
}

// NewStagingClearAction instantiates a new StagingClearAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStagingClearAction() *StagingClearAction {
	this := StagingClearAction{}
	return &this
}

// NewStagingClearActionWithDefaults instantiates a new StagingClearAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStagingClearActionWithDefaults() *StagingClearAction {
	this := StagingClearAction{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *StagingClearAction) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagingClearAction) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *StagingClearAction) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *StagingClearAction) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *StagingClearAction) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagingClearAction) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *StagingClearAction) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *StagingClearAction) SetKind(v string) {
	o.Kind = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *StagingClearAction) GetMeta() ApiObjectMeta {
	if o == nil || o.Meta == nil {
		var ret ApiObjectMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagingClearAction) GetMetaOk() (*ApiObjectMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *StagingClearAction) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ApiObjectMeta and assigns it to the Meta field.
func (o *StagingClearAction) SetMeta(v ApiObjectMeta) {
	o.Meta = &v
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *StagingClearAction) GetSpec() StagingClearActionSpec {
	if o == nil || o.Spec == nil {
		var ret StagingClearActionSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagingClearAction) GetSpecOk() (*StagingClearActionSpec, bool) {
	if o == nil || o.Spec == nil {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *StagingClearAction) HasSpec() bool {
	if o != nil && o.Spec != nil {
		return true
	}

	return false
}

// SetSpec gets a reference to the given StagingClearActionSpec and assigns it to the Spec field.
func (o *StagingClearAction) SetSpec(v StagingClearActionSpec) {
	o.Spec = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *StagingClearAction) GetStatus() StagingClearActionStatus {
	if o == nil || o.Status == nil {
		var ret StagingClearActionStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagingClearAction) GetStatusOk() (*StagingClearActionStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *StagingClearAction) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given StagingClearActionStatus and assigns it to the Status field.
func (o *StagingClearAction) SetStatus(v StagingClearActionStatus) {
	o.Status = &v
}

func (o StagingClearAction) MarshalJSON() ([]byte, error) {
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

type NullableStagingClearAction struct {
	value *StagingClearAction
	isSet bool
}

func (v NullableStagingClearAction) Get() *StagingClearAction {
	return v.value
}

func (v *NullableStagingClearAction) Set(val *StagingClearAction) {
	v.value = val
	v.isSet = true
}

func (v NullableStagingClearAction) IsSet() bool {
	return v.isSet
}

func (v *NullableStagingClearAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStagingClearAction(val *StagingClearAction) *NullableStagingClearAction {
	return &NullableStagingClearAction{value: val, isSet: true}
}

func (v NullableStagingClearAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStagingClearAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


