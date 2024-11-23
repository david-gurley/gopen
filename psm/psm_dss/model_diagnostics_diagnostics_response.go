/*
 * Diagnostics API reference
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

// DiagnosticsDiagnosticsResponse DiagnosticsResponse contains the response returned by service for the diagnostics query.
type DiagnosticsDiagnosticsResponse struct {
	Diagnostics *ApiAny `json:"diagnostics,omitempty"`
}

// NewDiagnosticsDiagnosticsResponse instantiates a new DiagnosticsDiagnosticsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiagnosticsDiagnosticsResponse() *DiagnosticsDiagnosticsResponse {
	this := DiagnosticsDiagnosticsResponse{}
	return &this
}

// NewDiagnosticsDiagnosticsResponseWithDefaults instantiates a new DiagnosticsDiagnosticsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiagnosticsDiagnosticsResponseWithDefaults() *DiagnosticsDiagnosticsResponse {
	this := DiagnosticsDiagnosticsResponse{}
	return &this
}

// GetDiagnostics returns the Diagnostics field value if set, zero value otherwise.
func (o *DiagnosticsDiagnosticsResponse) GetDiagnostics() ApiAny {
	if o == nil || o.Diagnostics == nil {
		var ret ApiAny
		return ret
	}
	return *o.Diagnostics
}

// GetDiagnosticsOk returns a tuple with the Diagnostics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticsDiagnosticsResponse) GetDiagnosticsOk() (*ApiAny, bool) {
	if o == nil || o.Diagnostics == nil {
		return nil, false
	}
	return o.Diagnostics, true
}

// HasDiagnostics returns a boolean if a field has been set.
func (o *DiagnosticsDiagnosticsResponse) HasDiagnostics() bool {
	if o != nil && o.Diagnostics != nil {
		return true
	}

	return false
}

// SetDiagnostics gets a reference to the given ApiAny and assigns it to the Diagnostics field.
func (o *DiagnosticsDiagnosticsResponse) SetDiagnostics(v ApiAny) {
	o.Diagnostics = &v
}

func (o DiagnosticsDiagnosticsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Diagnostics != nil {
		toSerialize["diagnostics"] = o.Diagnostics
	}
	return json.Marshal(toSerialize)
}

type NullableDiagnosticsDiagnosticsResponse struct {
	value *DiagnosticsDiagnosticsResponse
	isSet bool
}

func (v NullableDiagnosticsDiagnosticsResponse) Get() *DiagnosticsDiagnosticsResponse {
	return v.value
}

func (v *NullableDiagnosticsDiagnosticsResponse) Set(val *DiagnosticsDiagnosticsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDiagnosticsDiagnosticsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDiagnosticsDiagnosticsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiagnosticsDiagnosticsResponse(val *DiagnosticsDiagnosticsResponse) *NullableDiagnosticsDiagnosticsResponse {
	return &NullableDiagnosticsDiagnosticsResponse{value: val, isSet: true}
}

func (v NullableDiagnosticsDiagnosticsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiagnosticsDiagnosticsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


