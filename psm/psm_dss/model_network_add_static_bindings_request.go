/*
 * Network API reference
 *
 *  
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package psm_dss

import (
	"encoding/json"
)

// NetworkAddStaticBindingsRequest struct for NetworkAddStaticBindingsRequest
type NetworkAddStaticBindingsRequest struct {
	ApiVersion *string `json:"api-version,omitempty"`
	Ipv4StaticBindings *[]NetworkIPAMBinding `json:"ipv4-static-bindings,omitempty"`
	Kind *string `json:"kind,omitempty"`
	Meta *ApiObjectMeta `json:"meta,omitempty"`
}

// NewNetworkAddStaticBindingsRequest instantiates a new NetworkAddStaticBindingsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkAddStaticBindingsRequest() *NetworkAddStaticBindingsRequest {
	this := NetworkAddStaticBindingsRequest{}
	return &this
}

// NewNetworkAddStaticBindingsRequestWithDefaults instantiates a new NetworkAddStaticBindingsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkAddStaticBindingsRequestWithDefaults() *NetworkAddStaticBindingsRequest {
	this := NetworkAddStaticBindingsRequest{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *NetworkAddStaticBindingsRequest) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkAddStaticBindingsRequest) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *NetworkAddStaticBindingsRequest) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *NetworkAddStaticBindingsRequest) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetIpv4StaticBindings returns the Ipv4StaticBindings field value if set, zero value otherwise.
func (o *NetworkAddStaticBindingsRequest) GetIpv4StaticBindings() []NetworkIPAMBinding {
	if o == nil || o.Ipv4StaticBindings == nil {
		var ret []NetworkIPAMBinding
		return ret
	}
	return *o.Ipv4StaticBindings
}

// GetIpv4StaticBindingsOk returns a tuple with the Ipv4StaticBindings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkAddStaticBindingsRequest) GetIpv4StaticBindingsOk() (*[]NetworkIPAMBinding, bool) {
	if o == nil || o.Ipv4StaticBindings == nil {
		return nil, false
	}
	return o.Ipv4StaticBindings, true
}

// HasIpv4StaticBindings returns a boolean if a field has been set.
func (o *NetworkAddStaticBindingsRequest) HasIpv4StaticBindings() bool {
	if o != nil && o.Ipv4StaticBindings != nil {
		return true
	}

	return false
}

// SetIpv4StaticBindings gets a reference to the given []NetworkIPAMBinding and assigns it to the Ipv4StaticBindings field.
func (o *NetworkAddStaticBindingsRequest) SetIpv4StaticBindings(v []NetworkIPAMBinding) {
	o.Ipv4StaticBindings = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *NetworkAddStaticBindingsRequest) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkAddStaticBindingsRequest) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *NetworkAddStaticBindingsRequest) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *NetworkAddStaticBindingsRequest) SetKind(v string) {
	o.Kind = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *NetworkAddStaticBindingsRequest) GetMeta() ApiObjectMeta {
	if o == nil || o.Meta == nil {
		var ret ApiObjectMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkAddStaticBindingsRequest) GetMetaOk() (*ApiObjectMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *NetworkAddStaticBindingsRequest) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ApiObjectMeta and assigns it to the Meta field.
func (o *NetworkAddStaticBindingsRequest) SetMeta(v ApiObjectMeta) {
	o.Meta = &v
}

func (o NetworkAddStaticBindingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiVersion != nil {
		toSerialize["api-version"] = o.ApiVersion
	}
	if o.Ipv4StaticBindings != nil {
		toSerialize["ipv4-static-bindings"] = o.Ipv4StaticBindings
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkAddStaticBindingsRequest struct {
	value *NetworkAddStaticBindingsRequest
	isSet bool
}

func (v NullableNetworkAddStaticBindingsRequest) Get() *NetworkAddStaticBindingsRequest {
	return v.value
}

func (v *NullableNetworkAddStaticBindingsRequest) Set(val *NetworkAddStaticBindingsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkAddStaticBindingsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkAddStaticBindingsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkAddStaticBindingsRequest(val *NetworkAddStaticBindingsRequest) *NullableNetworkAddStaticBindingsRequest {
	return &NullableNetworkAddStaticBindingsRequest{value: val, isSet: true}
}

func (v NullableNetworkAddStaticBindingsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkAddStaticBindingsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


