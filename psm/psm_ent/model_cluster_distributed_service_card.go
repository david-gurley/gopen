/*
 * Cluster API reference
 *
 * Service name  
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package psm_ent

import (
	"encoding/json"
)

// ClusterDistributedServiceCard ------------------------------------ Distributed Service Card  ------------------------------------------- DistributedServiceCard represents the Naples I/O subsystem Entity responsible & scenarios involved in managing this object: Create: o CMD - created as part of NIC registration, Admittance Modify: o CMD - update spec attributes - update status attributes Delete: o CMD - aging out stale or rejected NICs (TBD) o NetOps, SecOps - Decomission a NIC (TBD).
type ClusterDistributedServiceCard struct {
	ApiVersion *string `json:"api-version,omitempty"`
	Kind *string `json:"kind,omitempty"`
	Meta *ApiObjectMeta `json:"meta,omitempty"`
	Spec *ClusterDistributedServiceCardSpec `json:"spec,omitempty"`
	Status *ClusterDistributedServiceCardStatus `json:"status,omitempty"`
}

// NewClusterDistributedServiceCard instantiates a new ClusterDistributedServiceCard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterDistributedServiceCard() *ClusterDistributedServiceCard {
	this := ClusterDistributedServiceCard{}
	return &this
}

// NewClusterDistributedServiceCardWithDefaults instantiates a new ClusterDistributedServiceCard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterDistributedServiceCardWithDefaults() *ClusterDistributedServiceCard {
	this := ClusterDistributedServiceCard{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *ClusterDistributedServiceCard) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDistributedServiceCard) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *ClusterDistributedServiceCard) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *ClusterDistributedServiceCard) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ClusterDistributedServiceCard) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDistributedServiceCard) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ClusterDistributedServiceCard) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *ClusterDistributedServiceCard) SetKind(v string) {
	o.Kind = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ClusterDistributedServiceCard) GetMeta() ApiObjectMeta {
	if o == nil || o.Meta == nil {
		var ret ApiObjectMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDistributedServiceCard) GetMetaOk() (*ApiObjectMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ClusterDistributedServiceCard) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ApiObjectMeta and assigns it to the Meta field.
func (o *ClusterDistributedServiceCard) SetMeta(v ApiObjectMeta) {
	o.Meta = &v
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *ClusterDistributedServiceCard) GetSpec() ClusterDistributedServiceCardSpec {
	if o == nil || o.Spec == nil {
		var ret ClusterDistributedServiceCardSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDistributedServiceCard) GetSpecOk() (*ClusterDistributedServiceCardSpec, bool) {
	if o == nil || o.Spec == nil {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *ClusterDistributedServiceCard) HasSpec() bool {
	if o != nil && o.Spec != nil {
		return true
	}

	return false
}

// SetSpec gets a reference to the given ClusterDistributedServiceCardSpec and assigns it to the Spec field.
func (o *ClusterDistributedServiceCard) SetSpec(v ClusterDistributedServiceCardSpec) {
	o.Spec = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ClusterDistributedServiceCard) GetStatus() ClusterDistributedServiceCardStatus {
	if o == nil || o.Status == nil {
		var ret ClusterDistributedServiceCardStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDistributedServiceCard) GetStatusOk() (*ClusterDistributedServiceCardStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ClusterDistributedServiceCard) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ClusterDistributedServiceCardStatus and assigns it to the Status field.
func (o *ClusterDistributedServiceCard) SetStatus(v ClusterDistributedServiceCardStatus) {
	o.Status = &v
}

func (o ClusterDistributedServiceCard) MarshalJSON() ([]byte, error) {
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

type NullableClusterDistributedServiceCard struct {
	value *ClusterDistributedServiceCard
	isSet bool
}

func (v NullableClusterDistributedServiceCard) Get() *ClusterDistributedServiceCard {
	return v.value
}

func (v *NullableClusterDistributedServiceCard) Set(val *ClusterDistributedServiceCard) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterDistributedServiceCard) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterDistributedServiceCard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterDistributedServiceCard(val *ClusterDistributedServiceCard) *NullableClusterDistributedServiceCard {
	return &NullableClusterDistributedServiceCard{value: val, isSet: true}
}

func (v NullableClusterDistributedServiceCard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterDistributedServiceCard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

