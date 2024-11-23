/*
 * Workload API reference
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

// ApiWatchEvent struct for ApiWatchEvent
type ApiWatchEvent struct {
	Control *ApiWatchControl `json:"control,omitempty"`
	Object *GoogleprotobufAny `json:"object,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewApiWatchEvent instantiates a new ApiWatchEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiWatchEvent() *ApiWatchEvent {
	this := ApiWatchEvent{}
	return &this
}

// NewApiWatchEventWithDefaults instantiates a new ApiWatchEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiWatchEventWithDefaults() *ApiWatchEvent {
	this := ApiWatchEvent{}
	return &this
}

// GetControl returns the Control field value if set, zero value otherwise.
func (o *ApiWatchEvent) GetControl() ApiWatchControl {
	if o == nil || o.Control == nil {
		var ret ApiWatchControl
		return ret
	}
	return *o.Control
}

// GetControlOk returns a tuple with the Control field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiWatchEvent) GetControlOk() (*ApiWatchControl, bool) {
	if o == nil || o.Control == nil {
		return nil, false
	}
	return o.Control, true
}

// HasControl returns a boolean if a field has been set.
func (o *ApiWatchEvent) HasControl() bool {
	if o != nil && o.Control != nil {
		return true
	}

	return false
}

// SetControl gets a reference to the given ApiWatchControl and assigns it to the Control field.
func (o *ApiWatchEvent) SetControl(v ApiWatchControl) {
	o.Control = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *ApiWatchEvent) GetObject() GoogleprotobufAny {
	if o == nil || o.Object == nil {
		var ret GoogleprotobufAny
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiWatchEvent) GetObjectOk() (*GoogleprotobufAny, bool) {
	if o == nil || o.Object == nil {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *ApiWatchEvent) HasObject() bool {
	if o != nil && o.Object != nil {
		return true
	}

	return false
}

// SetObject gets a reference to the given GoogleprotobufAny and assigns it to the Object field.
func (o *ApiWatchEvent) SetObject(v GoogleprotobufAny) {
	o.Object = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ApiWatchEvent) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiWatchEvent) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ApiWatchEvent) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ApiWatchEvent) SetType(v string) {
	o.Type = &v
}

func (o ApiWatchEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Control != nil {
		toSerialize["control"] = o.Control
	}
	if o.Object != nil {
		toSerialize["object"] = o.Object
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableApiWatchEvent struct {
	value *ApiWatchEvent
	isSet bool
}

func (v NullableApiWatchEvent) Get() *ApiWatchEvent {
	return v.value
}

func (v *NullableApiWatchEvent) Set(val *ApiWatchEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableApiWatchEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableApiWatchEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiWatchEvent(val *ApiWatchEvent) *NullableApiWatchEvent {
	return &NullableApiWatchEvent{value: val, isSet: true}
}

func (v NullableApiWatchEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiWatchEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

