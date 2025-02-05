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

// ClusterNodeSpec NodeSpec contains the configuration of the node.
type ClusterNodeSpec struct {
	// RoutingConfig the routing configuration.
	RoutingConfig *string `json:"routing-config,omitempty"`
}

// NewClusterNodeSpec instantiates a new ClusterNodeSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterNodeSpec() *ClusterNodeSpec {
	this := ClusterNodeSpec{}
	return &this
}

// NewClusterNodeSpecWithDefaults instantiates a new ClusterNodeSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterNodeSpecWithDefaults() *ClusterNodeSpec {
	this := ClusterNodeSpec{}
	return &this
}

// GetRoutingConfig returns the RoutingConfig field value if set, zero value otherwise.
func (o *ClusterNodeSpec) GetRoutingConfig() string {
	if o == nil || o.RoutingConfig == nil {
		var ret string
		return ret
	}
	return *o.RoutingConfig
}

// GetRoutingConfigOk returns a tuple with the RoutingConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterNodeSpec) GetRoutingConfigOk() (*string, bool) {
	if o == nil || o.RoutingConfig == nil {
		return nil, false
	}
	return o.RoutingConfig, true
}

// HasRoutingConfig returns a boolean if a field has been set.
func (o *ClusterNodeSpec) HasRoutingConfig() bool {
	if o != nil && o.RoutingConfig != nil {
		return true
	}

	return false
}

// SetRoutingConfig gets a reference to the given string and assigns it to the RoutingConfig field.
func (o *ClusterNodeSpec) SetRoutingConfig(v string) {
	o.RoutingConfig = &v
}

func (o ClusterNodeSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RoutingConfig != nil {
		toSerialize["routing-config"] = o.RoutingConfig
	}
	return json.Marshal(toSerialize)
}

type NullableClusterNodeSpec struct {
	value *ClusterNodeSpec
	isSet bool
}

func (v NullableClusterNodeSpec) Get() *ClusterNodeSpec {
	return v.value
}

func (v *NullableClusterNodeSpec) Set(val *ClusterNodeSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterNodeSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterNodeSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterNodeSpec(val *ClusterNodeSpec) *NullableClusterNodeSpec {
	return &NullableClusterNodeSpec{value: val, isSet: true}
}

func (v NullableClusterNodeSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterNodeSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


