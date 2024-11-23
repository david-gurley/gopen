/*
 * Network API reference
 *
 *  
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package psm_ent

import (
	"encoding/json"
)

// NetworkRoutingConfig struct for NetworkRoutingConfig
type NetworkRoutingConfig struct {
	ApiVersion *string `json:"api-version,omitempty"`
	Kind *string `json:"kind,omitempty"`
	Meta *ApiObjectMeta `json:"meta,omitempty"`
	Spec *NetworkRoutingConfigSpec `json:"spec,omitempty"`
	Status *NetworkRoutingConfigStatus `json:"status,omitempty"`
}

// NewNetworkRoutingConfig instantiates a new NetworkRoutingConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkRoutingConfig() *NetworkRoutingConfig {
	this := NetworkRoutingConfig{}
	return &this
}

// NewNetworkRoutingConfigWithDefaults instantiates a new NetworkRoutingConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkRoutingConfigWithDefaults() *NetworkRoutingConfig {
	this := NetworkRoutingConfig{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *NetworkRoutingConfig) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkRoutingConfig) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *NetworkRoutingConfig) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *NetworkRoutingConfig) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *NetworkRoutingConfig) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkRoutingConfig) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *NetworkRoutingConfig) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *NetworkRoutingConfig) SetKind(v string) {
	o.Kind = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *NetworkRoutingConfig) GetMeta() ApiObjectMeta {
	if o == nil || o.Meta == nil {
		var ret ApiObjectMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkRoutingConfig) GetMetaOk() (*ApiObjectMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *NetworkRoutingConfig) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ApiObjectMeta and assigns it to the Meta field.
func (o *NetworkRoutingConfig) SetMeta(v ApiObjectMeta) {
	o.Meta = &v
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *NetworkRoutingConfig) GetSpec() NetworkRoutingConfigSpec {
	if o == nil || o.Spec == nil {
		var ret NetworkRoutingConfigSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkRoutingConfig) GetSpecOk() (*NetworkRoutingConfigSpec, bool) {
	if o == nil || o.Spec == nil {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *NetworkRoutingConfig) HasSpec() bool {
	if o != nil && o.Spec != nil {
		return true
	}

	return false
}

// SetSpec gets a reference to the given NetworkRoutingConfigSpec and assigns it to the Spec field.
func (o *NetworkRoutingConfig) SetSpec(v NetworkRoutingConfigSpec) {
	o.Spec = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *NetworkRoutingConfig) GetStatus() NetworkRoutingConfigStatus {
	if o == nil || o.Status == nil {
		var ret NetworkRoutingConfigStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkRoutingConfig) GetStatusOk() (*NetworkRoutingConfigStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *NetworkRoutingConfig) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NetworkRoutingConfigStatus and assigns it to the Status field.
func (o *NetworkRoutingConfig) SetStatus(v NetworkRoutingConfigStatus) {
	o.Status = &v
}

func (o NetworkRoutingConfig) MarshalJSON() ([]byte, error) {
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

type NullableNetworkRoutingConfig struct {
	value *NetworkRoutingConfig
	isSet bool
}

func (v NullableNetworkRoutingConfig) Get() *NetworkRoutingConfig {
	return v.value
}

func (v *NullableNetworkRoutingConfig) Set(val *NetworkRoutingConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkRoutingConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkRoutingConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkRoutingConfig(val *NetworkRoutingConfig) *NullableNetworkRoutingConfig {
	return &NullableNetworkRoutingConfig{value: val, isSet: true}
}

func (v NullableNetworkRoutingConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkRoutingConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


