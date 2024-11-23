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

// NetworkRouteTableStatus struct for NetworkRouteTableStatus
type NetworkRouteTableStatus struct {
	Routes *[]NetworkRoute `json:"routes,omitempty"`
}

// NewNetworkRouteTableStatus instantiates a new NetworkRouteTableStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkRouteTableStatus() *NetworkRouteTableStatus {
	this := NetworkRouteTableStatus{}
	return &this
}

// NewNetworkRouteTableStatusWithDefaults instantiates a new NetworkRouteTableStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkRouteTableStatusWithDefaults() *NetworkRouteTableStatus {
	this := NetworkRouteTableStatus{}
	return &this
}

// GetRoutes returns the Routes field value if set, zero value otherwise.
func (o *NetworkRouteTableStatus) GetRoutes() []NetworkRoute {
	if o == nil || o.Routes == nil {
		var ret []NetworkRoute
		return ret
	}
	return *o.Routes
}

// GetRoutesOk returns a tuple with the Routes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkRouteTableStatus) GetRoutesOk() (*[]NetworkRoute, bool) {
	if o == nil || o.Routes == nil {
		return nil, false
	}
	return o.Routes, true
}

// HasRoutes returns a boolean if a field has been set.
func (o *NetworkRouteTableStatus) HasRoutes() bool {
	if o != nil && o.Routes != nil {
		return true
	}

	return false
}

// SetRoutes gets a reference to the given []NetworkRoute and assigns it to the Routes field.
func (o *NetworkRouteTableStatus) SetRoutes(v []NetworkRoute) {
	o.Routes = &v
}

func (o NetworkRouteTableStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Routes != nil {
		toSerialize["routes"] = o.Routes
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkRouteTableStatus struct {
	value *NetworkRouteTableStatus
	isSet bool
}

func (v NullableNetworkRouteTableStatus) Get() *NetworkRouteTableStatus {
	return v.value
}

func (v *NullableNetworkRouteTableStatus) Set(val *NetworkRouteTableStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkRouteTableStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkRouteTableStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkRouteTableStatus(val *NetworkRouteTableStatus) *NullableNetworkRouteTableStatus {
	return &NullableNetworkRouteTableStatus{value: val, isSet: true}
}

func (v NullableNetworkRouteTableStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkRouteTableStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


