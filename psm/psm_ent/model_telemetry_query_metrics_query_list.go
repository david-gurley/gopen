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

// TelemetryQueryMetricsQueryList MetricsQueryList contains a list of queries to execute.
type TelemetryQueryMetricsQueryList struct {
	// Namespace for the request.
	Namespace *string `json:"namespace,omitempty"`
	// List of queries to execute.
	Queries *[]TelemetryQueryMetricsQuerySpec `json:"queries,omitempty"`
	// Tenant for the request.
	Tenant *string `json:"tenant,omitempty"`
}

// NewTelemetryQueryMetricsQueryList instantiates a new TelemetryQueryMetricsQueryList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryQueryMetricsQueryList() *TelemetryQueryMetricsQueryList {
	this := TelemetryQueryMetricsQueryList{}
	return &this
}

// NewTelemetryQueryMetricsQueryListWithDefaults instantiates a new TelemetryQueryMetricsQueryList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryQueryMetricsQueryListWithDefaults() *TelemetryQueryMetricsQueryList {
	this := TelemetryQueryMetricsQueryList{}
	return &this
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *TelemetryQueryMetricsQueryList) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryQueryMetricsQueryList) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *TelemetryQueryMetricsQueryList) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *TelemetryQueryMetricsQueryList) SetNamespace(v string) {
	o.Namespace = &v
}

// GetQueries returns the Queries field value if set, zero value otherwise.
func (o *TelemetryQueryMetricsQueryList) GetQueries() []TelemetryQueryMetricsQuerySpec {
	if o == nil || o.Queries == nil {
		var ret []TelemetryQueryMetricsQuerySpec
		return ret
	}
	return *o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryQueryMetricsQueryList) GetQueriesOk() (*[]TelemetryQueryMetricsQuerySpec, bool) {
	if o == nil || o.Queries == nil {
		return nil, false
	}
	return o.Queries, true
}

// HasQueries returns a boolean if a field has been set.
func (o *TelemetryQueryMetricsQueryList) HasQueries() bool {
	if o != nil && o.Queries != nil {
		return true
	}

	return false
}

// SetQueries gets a reference to the given []TelemetryQueryMetricsQuerySpec and assigns it to the Queries field.
func (o *TelemetryQueryMetricsQueryList) SetQueries(v []TelemetryQueryMetricsQuerySpec) {
	o.Queries = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *TelemetryQueryMetricsQueryList) GetTenant() string {
	if o == nil || o.Tenant == nil {
		var ret string
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryQueryMetricsQueryList) GetTenantOk() (*string, bool) {
	if o == nil || o.Tenant == nil {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *TelemetryQueryMetricsQueryList) HasTenant() bool {
	if o != nil && o.Tenant != nil {
		return true
	}

	return false
}

// SetTenant gets a reference to the given string and assigns it to the Tenant field.
func (o *TelemetryQueryMetricsQueryList) SetTenant(v string) {
	o.Tenant = &v
}

func (o TelemetryQueryMetricsQueryList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.Queries != nil {
		toSerialize["queries"] = o.Queries
	}
	if o.Tenant != nil {
		toSerialize["tenant"] = o.Tenant
	}
	return json.Marshal(toSerialize)
}

type NullableTelemetryQueryMetricsQueryList struct {
	value *TelemetryQueryMetricsQueryList
	isSet bool
}

func (v NullableTelemetryQueryMetricsQueryList) Get() *TelemetryQueryMetricsQueryList {
	return v.value
}

func (v *NullableTelemetryQueryMetricsQueryList) Set(val *TelemetryQueryMetricsQueryList) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryQueryMetricsQueryList) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryQueryMetricsQueryList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryQueryMetricsQueryList(val *TelemetryQueryMetricsQueryList) *NullableTelemetryQueryMetricsQueryList {
	return &NullableTelemetryQueryMetricsQueryList{value: val, isSet: true}
}

func (v NullableTelemetryQueryMetricsQueryList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryQueryMetricsQueryList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


