/*
 * Telemetry_query API reference
 *
 * Service name  
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package psm_dss

import (
	"encoding/json"
)

// TelemetryQueryMetricsSQLQuery MetricsSQLQuery contains SQL query to execute.
type TelemetryQueryMetricsSQLQuery struct {
	// Namespace for the request.
	Namespace *string `json:"namespace,omitempty"`
	// Query string.
	Query *string `json:"query,omitempty"`
	// Tenant for the request.
	Tenant *string `json:"tenant,omitempty"`
}

// NewTelemetryQueryMetricsSQLQuery instantiates a new TelemetryQueryMetricsSQLQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryQueryMetricsSQLQuery() *TelemetryQueryMetricsSQLQuery {
	this := TelemetryQueryMetricsSQLQuery{}
	return &this
}

// NewTelemetryQueryMetricsSQLQueryWithDefaults instantiates a new TelemetryQueryMetricsSQLQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryQueryMetricsSQLQueryWithDefaults() *TelemetryQueryMetricsSQLQuery {
	this := TelemetryQueryMetricsSQLQuery{}
	return &this
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *TelemetryQueryMetricsSQLQuery) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryQueryMetricsSQLQuery) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *TelemetryQueryMetricsSQLQuery) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *TelemetryQueryMetricsSQLQuery) SetNamespace(v string) {
	o.Namespace = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *TelemetryQueryMetricsSQLQuery) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryQueryMetricsSQLQuery) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *TelemetryQueryMetricsSQLQuery) HasQuery() bool {
	if o != nil && o.Query != nil {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *TelemetryQueryMetricsSQLQuery) SetQuery(v string) {
	o.Query = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *TelemetryQueryMetricsSQLQuery) GetTenant() string {
	if o == nil || o.Tenant == nil {
		var ret string
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryQueryMetricsSQLQuery) GetTenantOk() (*string, bool) {
	if o == nil || o.Tenant == nil {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *TelemetryQueryMetricsSQLQuery) HasTenant() bool {
	if o != nil && o.Tenant != nil {
		return true
	}

	return false
}

// SetTenant gets a reference to the given string and assigns it to the Tenant field.
func (o *TelemetryQueryMetricsSQLQuery) SetTenant(v string) {
	o.Tenant = &v
}

func (o TelemetryQueryMetricsSQLQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.Tenant != nil {
		toSerialize["tenant"] = o.Tenant
	}
	return json.Marshal(toSerialize)
}

type NullableTelemetryQueryMetricsSQLQuery struct {
	value *TelemetryQueryMetricsSQLQuery
	isSet bool
}

func (v NullableTelemetryQueryMetricsSQLQuery) Get() *TelemetryQueryMetricsSQLQuery {
	return v.value
}

func (v *NullableTelemetryQueryMetricsSQLQuery) Set(val *TelemetryQueryMetricsSQLQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryQueryMetricsSQLQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryQueryMetricsSQLQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryQueryMetricsSQLQuery(val *TelemetryQueryMetricsSQLQuery) *NullableTelemetryQueryMetricsSQLQuery {
	return &NullableTelemetryQueryMetricsSQLQuery{value: val, isSet: true}
}

func (v NullableTelemetryQueryMetricsSQLQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryQueryMetricsSQLQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


