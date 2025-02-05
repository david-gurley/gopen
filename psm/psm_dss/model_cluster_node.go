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

// ClusterNode Node is representation of a single Venice node in the system.
type ClusterNode struct {
	ApiVersion *string `json:"api-version,omitempty"`
	Kind *string `json:"kind,omitempty"`
	Meta *ApiObjectMeta `json:"meta,omitempty"`
	Spec *ClusterNodeSpec `json:"spec,omitempty"`
	Status *ClusterNodeStatus `json:"status,omitempty"`
}

// NewClusterNode instantiates a new ClusterNode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterNode() *ClusterNode {
	this := ClusterNode{}
	return &this
}

// NewClusterNodeWithDefaults instantiates a new ClusterNode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterNodeWithDefaults() *ClusterNode {
	this := ClusterNode{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *ClusterNode) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterNode) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *ClusterNode) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *ClusterNode) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ClusterNode) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterNode) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ClusterNode) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *ClusterNode) SetKind(v string) {
	o.Kind = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ClusterNode) GetMeta() ApiObjectMeta {
	if o == nil || o.Meta == nil {
		var ret ApiObjectMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterNode) GetMetaOk() (*ApiObjectMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ClusterNode) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ApiObjectMeta and assigns it to the Meta field.
func (o *ClusterNode) SetMeta(v ApiObjectMeta) {
	o.Meta = &v
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *ClusterNode) GetSpec() ClusterNodeSpec {
	if o == nil || o.Spec == nil {
		var ret ClusterNodeSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterNode) GetSpecOk() (*ClusterNodeSpec, bool) {
	if o == nil || o.Spec == nil {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *ClusterNode) HasSpec() bool {
	if o != nil && o.Spec != nil {
		return true
	}

	return false
}

// SetSpec gets a reference to the given ClusterNodeSpec and assigns it to the Spec field.
func (o *ClusterNode) SetSpec(v ClusterNodeSpec) {
	o.Spec = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ClusterNode) GetStatus() ClusterNodeStatus {
	if o == nil || o.Status == nil {
		var ret ClusterNodeStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterNode) GetStatusOk() (*ClusterNodeStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ClusterNode) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ClusterNodeStatus and assigns it to the Status field.
func (o *ClusterNode) SetStatus(v ClusterNodeStatus) {
	o.Status = &v
}

func (o ClusterNode) MarshalJSON() ([]byte, error) {
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

type NullableClusterNode struct {
	value *ClusterNode
	isSet bool
}

func (v NullableClusterNode) Get() *ClusterNode {
	return v.value
}

func (v *NullableClusterNode) Set(val *ClusterNode) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterNode) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterNode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterNode(val *ClusterNode) *NullableClusterNode {
	return &NullableClusterNode{value: val, isSet: true}
}

func (v NullableClusterNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterNode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


