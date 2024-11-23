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

// ClusterConfigurationSnapshotRequest struct for ClusterConfigurationSnapshotRequest
type ClusterConfigurationSnapshotRequest struct {
	ApiVersion *string `json:"api-version,omitempty"`
	Kind *string `json:"kind,omitempty"`
	Meta *ApiObjectMeta `json:"meta,omitempty"`
}

// NewClusterConfigurationSnapshotRequest instantiates a new ClusterConfigurationSnapshotRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterConfigurationSnapshotRequest() *ClusterConfigurationSnapshotRequest {
	this := ClusterConfigurationSnapshotRequest{}
	return &this
}

// NewClusterConfigurationSnapshotRequestWithDefaults instantiates a new ClusterConfigurationSnapshotRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterConfigurationSnapshotRequestWithDefaults() *ClusterConfigurationSnapshotRequest {
	this := ClusterConfigurationSnapshotRequest{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *ClusterConfigurationSnapshotRequest) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterConfigurationSnapshotRequest) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *ClusterConfigurationSnapshotRequest) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *ClusterConfigurationSnapshotRequest) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ClusterConfigurationSnapshotRequest) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterConfigurationSnapshotRequest) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ClusterConfigurationSnapshotRequest) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *ClusterConfigurationSnapshotRequest) SetKind(v string) {
	o.Kind = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ClusterConfigurationSnapshotRequest) GetMeta() ApiObjectMeta {
	if o == nil || o.Meta == nil {
		var ret ApiObjectMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterConfigurationSnapshotRequest) GetMetaOk() (*ApiObjectMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ClusterConfigurationSnapshotRequest) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ApiObjectMeta and assigns it to the Meta field.
func (o *ClusterConfigurationSnapshotRequest) SetMeta(v ApiObjectMeta) {
	o.Meta = &v
}

func (o ClusterConfigurationSnapshotRequest) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullableClusterConfigurationSnapshotRequest struct {
	value *ClusterConfigurationSnapshotRequest
	isSet bool
}

func (v NullableClusterConfigurationSnapshotRequest) Get() *ClusterConfigurationSnapshotRequest {
	return v.value
}

func (v *NullableClusterConfigurationSnapshotRequest) Set(val *ClusterConfigurationSnapshotRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterConfigurationSnapshotRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterConfigurationSnapshotRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterConfigurationSnapshotRequest(val *ClusterConfigurationSnapshotRequest) *NullableClusterConfigurationSnapshotRequest {
	return &NullableClusterConfigurationSnapshotRequest{value: val, isSet: true}
}

func (v NullableClusterConfigurationSnapshotRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterConfigurationSnapshotRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


