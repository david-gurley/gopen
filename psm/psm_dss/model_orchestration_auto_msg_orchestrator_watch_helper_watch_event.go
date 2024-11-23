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

// OrchestrationAutoMsgOrchestratorWatchHelperWatchEvent struct for OrchestrationAutoMsgOrchestratorWatchHelperWatchEvent
type OrchestrationAutoMsgOrchestratorWatchHelperWatchEvent struct {
	Object *OrchestrationOrchestrator `json:"object,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewOrchestrationAutoMsgOrchestratorWatchHelperWatchEvent instantiates a new OrchestrationAutoMsgOrchestratorWatchHelperWatchEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrchestrationAutoMsgOrchestratorWatchHelperWatchEvent() *OrchestrationAutoMsgOrchestratorWatchHelperWatchEvent {
	this := OrchestrationAutoMsgOrchestratorWatchHelperWatchEvent{}
	return &this
}

// NewOrchestrationAutoMsgOrchestratorWatchHelperWatchEventWithDefaults instantiates a new OrchestrationAutoMsgOrchestratorWatchHelperWatchEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrchestrationAutoMsgOrchestratorWatchHelperWatchEventWithDefaults() *OrchestrationAutoMsgOrchestratorWatchHelperWatchEvent {
	this := OrchestrationAutoMsgOrchestratorWatchHelperWatchEvent{}
	return &this
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *OrchestrationAutoMsgOrchestratorWatchHelperWatchEvent) GetObject() OrchestrationOrchestrator {
	if o == nil || o.Object == nil {
		var ret OrchestrationOrchestrator
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrchestrationAutoMsgOrchestratorWatchHelperWatchEvent) GetObjectOk() (*OrchestrationOrchestrator, bool) {
	if o == nil || o.Object == nil {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *OrchestrationAutoMsgOrchestratorWatchHelperWatchEvent) HasObject() bool {
	if o != nil && o.Object != nil {
		return true
	}

	return false
}

// SetObject gets a reference to the given OrchestrationOrchestrator and assigns it to the Object field.
func (o *OrchestrationAutoMsgOrchestratorWatchHelperWatchEvent) SetObject(v OrchestrationOrchestrator) {
	o.Object = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OrchestrationAutoMsgOrchestratorWatchHelperWatchEvent) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrchestrationAutoMsgOrchestratorWatchHelperWatchEvent) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OrchestrationAutoMsgOrchestratorWatchHelperWatchEvent) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OrchestrationAutoMsgOrchestratorWatchHelperWatchEvent) SetType(v string) {
	o.Type = &v
}

func (o OrchestrationAutoMsgOrchestratorWatchHelperWatchEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Object != nil {
		toSerialize["object"] = o.Object
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableOrchestrationAutoMsgOrchestratorWatchHelperWatchEvent struct {
	value *OrchestrationAutoMsgOrchestratorWatchHelperWatchEvent
	isSet bool
}

func (v NullableOrchestrationAutoMsgOrchestratorWatchHelperWatchEvent) Get() *OrchestrationAutoMsgOrchestratorWatchHelperWatchEvent {
	return v.value
}

func (v *NullableOrchestrationAutoMsgOrchestratorWatchHelperWatchEvent) Set(val *OrchestrationAutoMsgOrchestratorWatchHelperWatchEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableOrchestrationAutoMsgOrchestratorWatchHelperWatchEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableOrchestrationAutoMsgOrchestratorWatchHelperWatchEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrchestrationAutoMsgOrchestratorWatchHelperWatchEvent(val *OrchestrationAutoMsgOrchestratorWatchHelperWatchEvent) *NullableOrchestrationAutoMsgOrchestratorWatchHelperWatchEvent {
	return &NullableOrchestrationAutoMsgOrchestratorWatchHelperWatchEvent{value: val, isSet: true}
}

func (v NullableOrchestrationAutoMsgOrchestratorWatchHelperWatchEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrchestrationAutoMsgOrchestratorWatchHelperWatchEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


