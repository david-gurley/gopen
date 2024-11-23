/*
 * Telemetry_query API reference
 *
 * Service name  
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package psm_ent

import (
	"encoding/json"
)

// TelemetryQueryMetricsQueryResponse MetricsQueryResponse is the response for Metrics Query.
type TelemetryQueryMetricsQueryResponse struct {
	// Namespace for the request.
	Namespace *string `json:"namespace,omitempty"`
	Results *[]TelemetryQueryMetricsQueryResult `json:"results,omitempty"`
	// Tenant for the request.
	Tenant *string `json:"tenant,omitempty"`
}

// NewTelemetryQueryMetricsQueryResponse instantiates a new TelemetryQueryMetricsQueryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryQueryMetricsQueryResponse() *TelemetryQueryMetricsQueryResponse {
	this := TelemetryQueryMetricsQueryResponse{}
	return &this
}

// NewTelemetryQueryMetricsQueryResponseWithDefaults instantiates a new TelemetryQueryMetricsQueryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryQueryMetricsQueryResponseWithDefaults() *TelemetryQueryMetricsQueryResponse {
	this := TelemetryQueryMetricsQueryResponse{}
	return &this
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *TelemetryQueryMetricsQueryResponse) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryQueryMetricsQueryResponse) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *TelemetryQueryMetricsQueryResponse) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *TelemetryQueryMetricsQueryResponse) SetNamespace(v string) {
	o.Namespace = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *TelemetryQueryMetricsQueryResponse) GetResults() []TelemetryQueryMetricsQueryResult {
	if o == nil || o.Results == nil {
		var ret []TelemetryQueryMetricsQueryResult
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryQueryMetricsQueryResponse) GetResultsOk() (*[]TelemetryQueryMetricsQueryResult, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *TelemetryQueryMetricsQueryResponse) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []TelemetryQueryMetricsQueryResult and assigns it to the Results field.
func (o *TelemetryQueryMetricsQueryResponse) SetResults(v []TelemetryQueryMetricsQueryResult) {
	o.Results = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *TelemetryQueryMetricsQueryResponse) GetTenant() string {
	if o == nil || o.Tenant == nil {
		var ret string
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryQueryMetricsQueryResponse) GetTenantOk() (*string, bool) {
	if o == nil || o.Tenant == nil {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *TelemetryQueryMetricsQueryResponse) HasTenant() bool {
	if o != nil && o.Tenant != nil {
		return true
	}

	return false
}

// SetTenant gets a reference to the given string and assigns it to the Tenant field.
func (o *TelemetryQueryMetricsQueryResponse) SetTenant(v string) {
	o.Tenant = &v
}

func (o TelemetryQueryMetricsQueryResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	if o.Tenant != nil {
		toSerialize["tenant"] = o.Tenant
	}
	return json.Marshal(toSerialize)
}

type NullableTelemetryQueryMetricsQueryResponse struct {
	value *TelemetryQueryMetricsQueryResponse
	isSet bool
}

func (v NullableTelemetryQueryMetricsQueryResponse) Get() *TelemetryQueryMetricsQueryResponse {
	return v.value
}

func (v *NullableTelemetryQueryMetricsQueryResponse) Set(val *TelemetryQueryMetricsQueryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryQueryMetricsQueryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryQueryMetricsQueryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryQueryMetricsQueryResponse(val *TelemetryQueryMetricsQueryResponse) *NullableTelemetryQueryMetricsQueryResponse {
	return &NullableTelemetryQueryMetricsQueryResponse{value: val, isSet: true}
}

func (v NullableTelemetryQueryMetricsQueryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryQueryMetricsQueryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

