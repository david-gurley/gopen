/*
 * Routing API reference
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

// RoutingNeighbor struct for RoutingNeighbor
type RoutingNeighbor struct {
	ApiVersion *string `json:"api-version,omitempty"`
	Kind *string `json:"kind,omitempty"`
	Meta *ApiObjectMeta `json:"meta,omitempty"`
	Spec *NetworkBGPNeighbor `json:"spec,omitempty"`
	Status *RoutingNeighborStatus `json:"status,omitempty"`
}

// NewRoutingNeighbor instantiates a new RoutingNeighbor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoutingNeighbor() *RoutingNeighbor {
	this := RoutingNeighbor{}
	return &this
}

// NewRoutingNeighborWithDefaults instantiates a new RoutingNeighbor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoutingNeighborWithDefaults() *RoutingNeighbor {
	this := RoutingNeighbor{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *RoutingNeighbor) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingNeighbor) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *RoutingNeighbor) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *RoutingNeighbor) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *RoutingNeighbor) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingNeighbor) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *RoutingNeighbor) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *RoutingNeighbor) SetKind(v string) {
	o.Kind = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *RoutingNeighbor) GetMeta() ApiObjectMeta {
	if o == nil || o.Meta == nil {
		var ret ApiObjectMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingNeighbor) GetMetaOk() (*ApiObjectMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *RoutingNeighbor) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ApiObjectMeta and assigns it to the Meta field.
func (o *RoutingNeighbor) SetMeta(v ApiObjectMeta) {
	o.Meta = &v
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *RoutingNeighbor) GetSpec() NetworkBGPNeighbor {
	if o == nil || o.Spec == nil {
		var ret NetworkBGPNeighbor
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingNeighbor) GetSpecOk() (*NetworkBGPNeighbor, bool) {
	if o == nil || o.Spec == nil {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *RoutingNeighbor) HasSpec() bool {
	if o != nil && o.Spec != nil {
		return true
	}

	return false
}

// SetSpec gets a reference to the given NetworkBGPNeighbor and assigns it to the Spec field.
func (o *RoutingNeighbor) SetSpec(v NetworkBGPNeighbor) {
	o.Spec = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RoutingNeighbor) GetStatus() RoutingNeighborStatus {
	if o == nil || o.Status == nil {
		var ret RoutingNeighborStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingNeighbor) GetStatusOk() (*RoutingNeighborStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RoutingNeighbor) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given RoutingNeighborStatus and assigns it to the Status field.
func (o *RoutingNeighbor) SetStatus(v RoutingNeighborStatus) {
	o.Status = &v
}

func (o RoutingNeighbor) MarshalJSON() ([]byte, error) {
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

type NullableRoutingNeighbor struct {
	value *RoutingNeighbor
	isSet bool
}

func (v NullableRoutingNeighbor) Get() *RoutingNeighbor {
	return v.value
}

func (v *NullableRoutingNeighbor) Set(val *RoutingNeighbor) {
	v.value = val
	v.isSet = true
}

func (v NullableRoutingNeighbor) IsSet() bool {
	return v.isSet
}

func (v *NullableRoutingNeighbor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoutingNeighbor(val *RoutingNeighbor) *NullableRoutingNeighbor {
	return &NullableRoutingNeighbor{value: val, isSet: true}
}

func (v NullableRoutingNeighbor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoutingNeighbor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


