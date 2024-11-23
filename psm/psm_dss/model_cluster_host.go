/*
 * Cluster API reference
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

// ClusterHost Host represents a Baremetal or Hypervisor server.
type ClusterHost struct {
	ApiVersion *string `json:"api-version,omitempty"`
	Kind *string `json:"kind,omitempty"`
	Meta *ApiObjectMeta `json:"meta,omitempty"`
	Spec *ClusterHostSpec `json:"spec,omitempty"`
	Status *ClusterHostStatus `json:"status,omitempty"`
}

// NewClusterHost instantiates a new ClusterHost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterHost() *ClusterHost {
	this := ClusterHost{}
	return &this
}

// NewClusterHostWithDefaults instantiates a new ClusterHost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterHostWithDefaults() *ClusterHost {
	this := ClusterHost{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *ClusterHost) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterHost) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *ClusterHost) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *ClusterHost) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ClusterHost) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterHost) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ClusterHost) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *ClusterHost) SetKind(v string) {
	o.Kind = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ClusterHost) GetMeta() ApiObjectMeta {
	if o == nil || o.Meta == nil {
		var ret ApiObjectMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterHost) GetMetaOk() (*ApiObjectMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ClusterHost) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ApiObjectMeta and assigns it to the Meta field.
func (o *ClusterHost) SetMeta(v ApiObjectMeta) {
	o.Meta = &v
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *ClusterHost) GetSpec() ClusterHostSpec {
	if o == nil || o.Spec == nil {
		var ret ClusterHostSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterHost) GetSpecOk() (*ClusterHostSpec, bool) {
	if o == nil || o.Spec == nil {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *ClusterHost) HasSpec() bool {
	if o != nil && o.Spec != nil {
		return true
	}

	return false
}

// SetSpec gets a reference to the given ClusterHostSpec and assigns it to the Spec field.
func (o *ClusterHost) SetSpec(v ClusterHostSpec) {
	o.Spec = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ClusterHost) GetStatus() ClusterHostStatus {
	if o == nil || o.Status == nil {
		var ret ClusterHostStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterHost) GetStatusOk() (*ClusterHostStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ClusterHost) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ClusterHostStatus and assigns it to the Status field.
func (o *ClusterHost) SetStatus(v ClusterHostStatus) {
	o.Status = &v
}

func (o ClusterHost) MarshalJSON() ([]byte, error) {
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

type NullableClusterHost struct {
	value *ClusterHost
	isSet bool
}

func (v NullableClusterHost) Get() *ClusterHost {
	return v.value
}

func (v *NullableClusterHost) Set(val *ClusterHost) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterHost) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterHost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterHost(val *ClusterHost) *NullableClusterHost {
	return &NullableClusterHost{value: val, isSet: true}
}

func (v NullableClusterHost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterHost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


