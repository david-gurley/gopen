/*
 * Auth API reference
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

// AuthOperation Operation represents an user action on a resource about which an authorization decision has to be made.
type AuthOperation struct {
	Action *string `json:"action,omitempty"`
	Resource *AuthResource `json:"resource,omitempty"`
}

// NewAuthOperation instantiates a new AuthOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthOperation() *AuthOperation {
	this := AuthOperation{}
	var action string = "all-actions"
	this.Action = &action
	return &this
}

// NewAuthOperationWithDefaults instantiates a new AuthOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthOperationWithDefaults() *AuthOperation {
	this := AuthOperation{}
	var action string = "all-actions"
	this.Action = &action
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *AuthOperation) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthOperation) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *AuthOperation) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *AuthOperation) SetAction(v string) {
	o.Action = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *AuthOperation) GetResource() AuthResource {
	if o == nil || o.Resource == nil {
		var ret AuthResource
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthOperation) GetResourceOk() (*AuthResource, bool) {
	if o == nil || o.Resource == nil {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *AuthOperation) HasResource() bool {
	if o != nil && o.Resource != nil {
		return true
	}

	return false
}

// SetResource gets a reference to the given AuthResource and assigns it to the Resource field.
func (o *AuthOperation) SetResource(v AuthResource) {
	o.Resource = &v
}

func (o AuthOperation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.Resource != nil {
		toSerialize["resource"] = o.Resource
	}
	return json.Marshal(toSerialize)
}

type NullableAuthOperation struct {
	value *AuthOperation
	isSet bool
}

func (v NullableAuthOperation) Get() *AuthOperation {
	return v.value
}

func (v *NullableAuthOperation) Set(val *AuthOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthOperation(val *AuthOperation) *NullableAuthOperation {
	return &NullableAuthOperation{value: val, isSet: true}
}

func (v NullableAuthOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

