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

// ClusterFeatureStatus FeatureStatus is the current operational status of a feature.
type ClusterFeatureStatus struct {
	Expiry *string `json:"expiry,omitempty"`
	FeatureKey *string `json:"feature-key,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewClusterFeatureStatus instantiates a new ClusterFeatureStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterFeatureStatus() *ClusterFeatureStatus {
	this := ClusterFeatureStatus{}
	return &this
}

// NewClusterFeatureStatusWithDefaults instantiates a new ClusterFeatureStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterFeatureStatusWithDefaults() *ClusterFeatureStatus {
	this := ClusterFeatureStatus{}
	return &this
}

// GetExpiry returns the Expiry field value if set, zero value otherwise.
func (o *ClusterFeatureStatus) GetExpiry() string {
	if o == nil || o.Expiry == nil {
		var ret string
		return ret
	}
	return *o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterFeatureStatus) GetExpiryOk() (*string, bool) {
	if o == nil || o.Expiry == nil {
		return nil, false
	}
	return o.Expiry, true
}

// HasExpiry returns a boolean if a field has been set.
func (o *ClusterFeatureStatus) HasExpiry() bool {
	if o != nil && o.Expiry != nil {
		return true
	}

	return false
}

// SetExpiry gets a reference to the given string and assigns it to the Expiry field.
func (o *ClusterFeatureStatus) SetExpiry(v string) {
	o.Expiry = &v
}

// GetFeatureKey returns the FeatureKey field value if set, zero value otherwise.
func (o *ClusterFeatureStatus) GetFeatureKey() string {
	if o == nil || o.FeatureKey == nil {
		var ret string
		return ret
	}
	return *o.FeatureKey
}

// GetFeatureKeyOk returns a tuple with the FeatureKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterFeatureStatus) GetFeatureKeyOk() (*string, bool) {
	if o == nil || o.FeatureKey == nil {
		return nil, false
	}
	return o.FeatureKey, true
}

// HasFeatureKey returns a boolean if a field has been set.
func (o *ClusterFeatureStatus) HasFeatureKey() bool {
	if o != nil && o.FeatureKey != nil {
		return true
	}

	return false
}

// SetFeatureKey gets a reference to the given string and assigns it to the FeatureKey field.
func (o *ClusterFeatureStatus) SetFeatureKey(v string) {
	o.FeatureKey = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ClusterFeatureStatus) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterFeatureStatus) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ClusterFeatureStatus) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ClusterFeatureStatus) SetValue(v string) {
	o.Value = &v
}

func (o ClusterFeatureStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Expiry != nil {
		toSerialize["expiry"] = o.Expiry
	}
	if o.FeatureKey != nil {
		toSerialize["feature-key"] = o.FeatureKey
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableClusterFeatureStatus struct {
	value *ClusterFeatureStatus
	isSet bool
}

func (v NullableClusterFeatureStatus) Get() *ClusterFeatureStatus {
	return v.value
}

func (v *NullableClusterFeatureStatus) Set(val *ClusterFeatureStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterFeatureStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterFeatureStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterFeatureStatus(val *ClusterFeatureStatus) *NullableClusterFeatureStatus {
	return &NullableClusterFeatureStatus{value: val, isSet: true}
}

func (v NullableClusterFeatureStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterFeatureStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


