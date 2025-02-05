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

// NetworkRoute struct for NetworkRoute
type NetworkRoute struct {
	// NextHop for the route. Should be a valid v4 or v6 CIDR block.
	NextHop *string `json:"next-hop,omitempty"`
	// Route Prefix for the route. Should be a valid v4 or v6 CIDR block.
	Prefix *string `json:"prefix,omitempty"`
	// Target VirtualRouter instance.
	TargetVirtualRouter *string `json:"target-virtual-router,omitempty"`
}

// NewNetworkRoute instantiates a new NetworkRoute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkRoute() *NetworkRoute {
	this := NetworkRoute{}
	return &this
}

// NewNetworkRouteWithDefaults instantiates a new NetworkRoute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkRouteWithDefaults() *NetworkRoute {
	this := NetworkRoute{}
	return &this
}

// GetNextHop returns the NextHop field value if set, zero value otherwise.
func (o *NetworkRoute) GetNextHop() string {
	if o == nil || o.NextHop == nil {
		var ret string
		return ret
	}
	return *o.NextHop
}

// GetNextHopOk returns a tuple with the NextHop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkRoute) GetNextHopOk() (*string, bool) {
	if o == nil || o.NextHop == nil {
		return nil, false
	}
	return o.NextHop, true
}

// HasNextHop returns a boolean if a field has been set.
func (o *NetworkRoute) HasNextHop() bool {
	if o != nil && o.NextHop != nil {
		return true
	}

	return false
}

// SetNextHop gets a reference to the given string and assigns it to the NextHop field.
func (o *NetworkRoute) SetNextHop(v string) {
	o.NextHop = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *NetworkRoute) GetPrefix() string {
	if o == nil || o.Prefix == nil {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkRoute) GetPrefixOk() (*string, bool) {
	if o == nil || o.Prefix == nil {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *NetworkRoute) HasPrefix() bool {
	if o != nil && o.Prefix != nil {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *NetworkRoute) SetPrefix(v string) {
	o.Prefix = &v
}

// GetTargetVirtualRouter returns the TargetVirtualRouter field value if set, zero value otherwise.
func (o *NetworkRoute) GetTargetVirtualRouter() string {
	if o == nil || o.TargetVirtualRouter == nil {
		var ret string
		return ret
	}
	return *o.TargetVirtualRouter
}

// GetTargetVirtualRouterOk returns a tuple with the TargetVirtualRouter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkRoute) GetTargetVirtualRouterOk() (*string, bool) {
	if o == nil || o.TargetVirtualRouter == nil {
		return nil, false
	}
	return o.TargetVirtualRouter, true
}

// HasTargetVirtualRouter returns a boolean if a field has been set.
func (o *NetworkRoute) HasTargetVirtualRouter() bool {
	if o != nil && o.TargetVirtualRouter != nil {
		return true
	}

	return false
}

// SetTargetVirtualRouter gets a reference to the given string and assigns it to the TargetVirtualRouter field.
func (o *NetworkRoute) SetTargetVirtualRouter(v string) {
	o.TargetVirtualRouter = &v
}

func (o NetworkRoute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NextHop != nil {
		toSerialize["next-hop"] = o.NextHop
	}
	if o.Prefix != nil {
		toSerialize["prefix"] = o.Prefix
	}
	if o.TargetVirtualRouter != nil {
		toSerialize["target-virtual-router"] = o.TargetVirtualRouter
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkRoute struct {
	value *NetworkRoute
	isSet bool
}

func (v NullableNetworkRoute) Get() *NetworkRoute {
	return v.value
}

func (v *NullableNetworkRoute) Set(val *NetworkRoute) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkRoute) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkRoute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkRoute(val *NetworkRoute) *NullableNetworkRoute {
	return &NullableNetworkRoute{value: val, isSet: true}
}

func (v NullableNetworkRoute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkRoute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


