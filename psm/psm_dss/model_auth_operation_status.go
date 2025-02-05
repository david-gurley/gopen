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

// AuthOperationStatus OperationStatus contains the result of the authorization check for a requested operation.
type AuthOperationStatus struct {
	// Allowed indicates if Operation is authorized.
	Allowed *bool `json:"allowed,omitempty"`
	// Message reports error validating Operation.
	Message *string `json:"message,omitempty"`
	Operation *AuthOperation `json:"operation,omitempty"`
}

// NewAuthOperationStatus instantiates a new AuthOperationStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthOperationStatus() *AuthOperationStatus {
	this := AuthOperationStatus{}
	return &this
}

// NewAuthOperationStatusWithDefaults instantiates a new AuthOperationStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthOperationStatusWithDefaults() *AuthOperationStatus {
	this := AuthOperationStatus{}
	return &this
}

// GetAllowed returns the Allowed field value if set, zero value otherwise.
func (o *AuthOperationStatus) GetAllowed() bool {
	if o == nil || o.Allowed == nil {
		var ret bool
		return ret
	}
	return *o.Allowed
}

// GetAllowedOk returns a tuple with the Allowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthOperationStatus) GetAllowedOk() (*bool, bool) {
	if o == nil || o.Allowed == nil {
		return nil, false
	}
	return o.Allowed, true
}

// HasAllowed returns a boolean if a field has been set.
func (o *AuthOperationStatus) HasAllowed() bool {
	if o != nil && o.Allowed != nil {
		return true
	}

	return false
}

// SetAllowed gets a reference to the given bool and assigns it to the Allowed field.
func (o *AuthOperationStatus) SetAllowed(v bool) {
	o.Allowed = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *AuthOperationStatus) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthOperationStatus) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *AuthOperationStatus) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *AuthOperationStatus) SetMessage(v string) {
	o.Message = &v
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *AuthOperationStatus) GetOperation() AuthOperation {
	if o == nil || o.Operation == nil {
		var ret AuthOperation
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthOperationStatus) GetOperationOk() (*AuthOperation, bool) {
	if o == nil || o.Operation == nil {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *AuthOperationStatus) HasOperation() bool {
	if o != nil && o.Operation != nil {
		return true
	}

	return false
}

// SetOperation gets a reference to the given AuthOperation and assigns it to the Operation field.
func (o *AuthOperationStatus) SetOperation(v AuthOperation) {
	o.Operation = &v
}

func (o AuthOperationStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Allowed != nil {
		toSerialize["allowed"] = o.Allowed
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Operation != nil {
		toSerialize["operation"] = o.Operation
	}
	return json.Marshal(toSerialize)
}

type NullableAuthOperationStatus struct {
	value *AuthOperationStatus
	isSet bool
}

func (v NullableAuthOperationStatus) Get() *AuthOperationStatus {
	return v.value
}

func (v *NullableAuthOperationStatus) Set(val *AuthOperationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthOperationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthOperationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthOperationStatus(val *AuthOperationStatus) *NullableAuthOperationStatus {
	return &NullableAuthOperationStatus{value: val, isSet: true}
}

func (v NullableAuthOperationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthOperationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


