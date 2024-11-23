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

// ClusterSnapshotRestore struct for ClusterSnapshotRestore
type ClusterSnapshotRestore struct {
	ApiVersion *string `json:"api-version,omitempty"`
	Kind *string `json:"kind,omitempty"`
	Meta *ApiObjectMeta `json:"meta,omitempty"`
	Spec *ClusterSnapshotRestoreSpec `json:"spec,omitempty"`
	Status *ClusterSnapshotRestoreStatus `json:"status,omitempty"`
}

// NewClusterSnapshotRestore instantiates a new ClusterSnapshotRestore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterSnapshotRestore() *ClusterSnapshotRestore {
	this := ClusterSnapshotRestore{}
	return &this
}

// NewClusterSnapshotRestoreWithDefaults instantiates a new ClusterSnapshotRestore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterSnapshotRestoreWithDefaults() *ClusterSnapshotRestore {
	this := ClusterSnapshotRestore{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *ClusterSnapshotRestore) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSnapshotRestore) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *ClusterSnapshotRestore) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *ClusterSnapshotRestore) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ClusterSnapshotRestore) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSnapshotRestore) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ClusterSnapshotRestore) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *ClusterSnapshotRestore) SetKind(v string) {
	o.Kind = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ClusterSnapshotRestore) GetMeta() ApiObjectMeta {
	if o == nil || o.Meta == nil {
		var ret ApiObjectMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSnapshotRestore) GetMetaOk() (*ApiObjectMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ClusterSnapshotRestore) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ApiObjectMeta and assigns it to the Meta field.
func (o *ClusterSnapshotRestore) SetMeta(v ApiObjectMeta) {
	o.Meta = &v
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *ClusterSnapshotRestore) GetSpec() ClusterSnapshotRestoreSpec {
	if o == nil || o.Spec == nil {
		var ret ClusterSnapshotRestoreSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSnapshotRestore) GetSpecOk() (*ClusterSnapshotRestoreSpec, bool) {
	if o == nil || o.Spec == nil {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *ClusterSnapshotRestore) HasSpec() bool {
	if o != nil && o.Spec != nil {
		return true
	}

	return false
}

// SetSpec gets a reference to the given ClusterSnapshotRestoreSpec and assigns it to the Spec field.
func (o *ClusterSnapshotRestore) SetSpec(v ClusterSnapshotRestoreSpec) {
	o.Spec = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ClusterSnapshotRestore) GetStatus() ClusterSnapshotRestoreStatus {
	if o == nil || o.Status == nil {
		var ret ClusterSnapshotRestoreStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSnapshotRestore) GetStatusOk() (*ClusterSnapshotRestoreStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ClusterSnapshotRestore) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ClusterSnapshotRestoreStatus and assigns it to the Status field.
func (o *ClusterSnapshotRestore) SetStatus(v ClusterSnapshotRestoreStatus) {
	o.Status = &v
}

func (o ClusterSnapshotRestore) MarshalJSON() ([]byte, error) {
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

type NullableClusterSnapshotRestore struct {
	value *ClusterSnapshotRestore
	isSet bool
}

func (v NullableClusterSnapshotRestore) Get() *ClusterSnapshotRestore {
	return v.value
}

func (v *NullableClusterSnapshotRestore) Set(val *ClusterSnapshotRestore) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterSnapshotRestore) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterSnapshotRestore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterSnapshotRestore(val *ClusterSnapshotRestore) *NullableClusterSnapshotRestore {
	return &NullableClusterSnapshotRestore{value: val, isSet: true}
}

func (v NullableClusterSnapshotRestore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterSnapshotRestore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


