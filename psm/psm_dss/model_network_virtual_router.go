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

// NetworkVirtualRouter VirtualRouter - represents a forwarding instance. Could be multiple per tenant.
type NetworkVirtualRouter struct {
	ApiVersion *string `json:"api-version,omitempty"`
	Kind *string `json:"kind,omitempty"`
	Meta *ApiObjectMeta `json:"meta,omitempty"`
	Spec *NetworkVirtualRouterSpec `json:"spec,omitempty"`
	Status *NetworkVirtualRouterStatus `json:"status,omitempty"`
}

// NewNetworkVirtualRouter instantiates a new NetworkVirtualRouter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkVirtualRouter() *NetworkVirtualRouter {
	this := NetworkVirtualRouter{}
	return &this
}

// NewNetworkVirtualRouterWithDefaults instantiates a new NetworkVirtualRouter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkVirtualRouterWithDefaults() *NetworkVirtualRouter {
	this := NetworkVirtualRouter{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *NetworkVirtualRouter) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkVirtualRouter) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *NetworkVirtualRouter) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *NetworkVirtualRouter) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *NetworkVirtualRouter) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkVirtualRouter) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *NetworkVirtualRouter) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *NetworkVirtualRouter) SetKind(v string) {
	o.Kind = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *NetworkVirtualRouter) GetMeta() ApiObjectMeta {
	if o == nil || o.Meta == nil {
		var ret ApiObjectMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkVirtualRouter) GetMetaOk() (*ApiObjectMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *NetworkVirtualRouter) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ApiObjectMeta and assigns it to the Meta field.
func (o *NetworkVirtualRouter) SetMeta(v ApiObjectMeta) {
	o.Meta = &v
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *NetworkVirtualRouter) GetSpec() NetworkVirtualRouterSpec {
	if o == nil || o.Spec == nil {
		var ret NetworkVirtualRouterSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkVirtualRouter) GetSpecOk() (*NetworkVirtualRouterSpec, bool) {
	if o == nil || o.Spec == nil {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *NetworkVirtualRouter) HasSpec() bool {
	if o != nil && o.Spec != nil {
		return true
	}

	return false
}

// SetSpec gets a reference to the given NetworkVirtualRouterSpec and assigns it to the Spec field.
func (o *NetworkVirtualRouter) SetSpec(v NetworkVirtualRouterSpec) {
	o.Spec = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *NetworkVirtualRouter) GetStatus() NetworkVirtualRouterStatus {
	if o == nil || o.Status == nil {
		var ret NetworkVirtualRouterStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkVirtualRouter) GetStatusOk() (*NetworkVirtualRouterStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *NetworkVirtualRouter) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NetworkVirtualRouterStatus and assigns it to the Status field.
func (o *NetworkVirtualRouter) SetStatus(v NetworkVirtualRouterStatus) {
	o.Status = &v
}

func (o NetworkVirtualRouter) MarshalJSON() ([]byte, error) {
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

type NullableNetworkVirtualRouter struct {
	value *NetworkVirtualRouter
	isSet bool
}

func (v NullableNetworkVirtualRouter) Get() *NetworkVirtualRouter {
	return v.value
}

func (v *NullableNetworkVirtualRouter) Set(val *NetworkVirtualRouter) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkVirtualRouter) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkVirtualRouter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkVirtualRouter(val *NetworkVirtualRouter) *NullableNetworkVirtualRouter {
	return &NullableNetworkVirtualRouter{value: val, isSet: true}
}

func (v NullableNetworkVirtualRouter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkVirtualRouter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


