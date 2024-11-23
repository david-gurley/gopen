/*
 * Orchestration API reference
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

// OrchestrationManagedNamespaceSpec ManagedNamespaceSpec contains namespace specific configuration.
type OrchestrationManagedNamespaceSpec struct {
	DiscoveryOperation *string `json:"discovery-operation,omitempty"`
	DiscoveryProtocol *string `json:"discovery-protocol,omitempty"`
	// MTU, 0 = Use system-default or orchestrator settings.
	Mtu *int64 `json:"mtu,omitempty"`
	MulticastFilter *string `json:"multicast-filter,omitempty"`
	// End of vlan range (inclusive) to be used for overrides. Range should be >= number of vnics expected per host. Value should be between 0 and 4095.
	OverrideVlanEnd *int64 `json:"override-vlan-end,omitempty"`
	// Start of vlan range to be used for overrides. Range should be >= number of vnics expected per host. Value should be between 0 and 4095.
	OverrideVlanStart *int64 `json:"override-vlan-start,omitempty"`
}

// NewOrchestrationManagedNamespaceSpec instantiates a new OrchestrationManagedNamespaceSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrchestrationManagedNamespaceSpec() *OrchestrationManagedNamespaceSpec {
	this := OrchestrationManagedNamespaceSpec{}
	var discoveryOperation string = "disc-op-default"
	this.DiscoveryOperation = &discoveryOperation
	var discoveryProtocol string = "disc-proto-default"
	this.DiscoveryProtocol = &discoveryProtocol
	var mtu int64 = 0
	this.Mtu = &mtu
	var multicastFilter string = "mcast-filter-default"
	this.MulticastFilter = &multicastFilter
	var overrideVlanEnd int64 = 4087
	this.OverrideVlanEnd = &overrideVlanEnd
	var overrideVlanStart int64 = 3832
	this.OverrideVlanStart = &overrideVlanStart
	return &this
}

// NewOrchestrationManagedNamespaceSpecWithDefaults instantiates a new OrchestrationManagedNamespaceSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrchestrationManagedNamespaceSpecWithDefaults() *OrchestrationManagedNamespaceSpec {
	this := OrchestrationManagedNamespaceSpec{}
	var discoveryOperation string = "disc-op-default"
	this.DiscoveryOperation = &discoveryOperation
	var discoveryProtocol string = "disc-proto-default"
	this.DiscoveryProtocol = &discoveryProtocol
	var mtu int64 = 0
	this.Mtu = &mtu
	var multicastFilter string = "mcast-filter-default"
	this.MulticastFilter = &multicastFilter
	var overrideVlanEnd int64 = 4087
	this.OverrideVlanEnd = &overrideVlanEnd
	var overrideVlanStart int64 = 3832
	this.OverrideVlanStart = &overrideVlanStart
	return &this
}

// GetDiscoveryOperation returns the DiscoveryOperation field value if set, zero value otherwise.
func (o *OrchestrationManagedNamespaceSpec) GetDiscoveryOperation() string {
	if o == nil || o.DiscoveryOperation == nil {
		var ret string
		return ret
	}
	return *o.DiscoveryOperation
}

// GetDiscoveryOperationOk returns a tuple with the DiscoveryOperation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrchestrationManagedNamespaceSpec) GetDiscoveryOperationOk() (*string, bool) {
	if o == nil || o.DiscoveryOperation == nil {
		return nil, false
	}
	return o.DiscoveryOperation, true
}

// HasDiscoveryOperation returns a boolean if a field has been set.
func (o *OrchestrationManagedNamespaceSpec) HasDiscoveryOperation() bool {
	if o != nil && o.DiscoveryOperation != nil {
		return true
	}

	return false
}

// SetDiscoveryOperation gets a reference to the given string and assigns it to the DiscoveryOperation field.
func (o *OrchestrationManagedNamespaceSpec) SetDiscoveryOperation(v string) {
	o.DiscoveryOperation = &v
}

// GetDiscoveryProtocol returns the DiscoveryProtocol field value if set, zero value otherwise.
func (o *OrchestrationManagedNamespaceSpec) GetDiscoveryProtocol() string {
	if o == nil || o.DiscoveryProtocol == nil {
		var ret string
		return ret
	}
	return *o.DiscoveryProtocol
}

// GetDiscoveryProtocolOk returns a tuple with the DiscoveryProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrchestrationManagedNamespaceSpec) GetDiscoveryProtocolOk() (*string, bool) {
	if o == nil || o.DiscoveryProtocol == nil {
		return nil, false
	}
	return o.DiscoveryProtocol, true
}

// HasDiscoveryProtocol returns a boolean if a field has been set.
func (o *OrchestrationManagedNamespaceSpec) HasDiscoveryProtocol() bool {
	if o != nil && o.DiscoveryProtocol != nil {
		return true
	}

	return false
}

// SetDiscoveryProtocol gets a reference to the given string and assigns it to the DiscoveryProtocol field.
func (o *OrchestrationManagedNamespaceSpec) SetDiscoveryProtocol(v string) {
	o.DiscoveryProtocol = &v
}

// GetMtu returns the Mtu field value if set, zero value otherwise.
func (o *OrchestrationManagedNamespaceSpec) GetMtu() int64 {
	if o == nil || o.Mtu == nil {
		var ret int64
		return ret
	}
	return *o.Mtu
}

// GetMtuOk returns a tuple with the Mtu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrchestrationManagedNamespaceSpec) GetMtuOk() (*int64, bool) {
	if o == nil || o.Mtu == nil {
		return nil, false
	}
	return o.Mtu, true
}

// HasMtu returns a boolean if a field has been set.
func (o *OrchestrationManagedNamespaceSpec) HasMtu() bool {
	if o != nil && o.Mtu != nil {
		return true
	}

	return false
}

// SetMtu gets a reference to the given int64 and assigns it to the Mtu field.
func (o *OrchestrationManagedNamespaceSpec) SetMtu(v int64) {
	o.Mtu = &v
}

// GetMulticastFilter returns the MulticastFilter field value if set, zero value otherwise.
func (o *OrchestrationManagedNamespaceSpec) GetMulticastFilter() string {
	if o == nil || o.MulticastFilter == nil {
		var ret string
		return ret
	}
	return *o.MulticastFilter
}

// GetMulticastFilterOk returns a tuple with the MulticastFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrchestrationManagedNamespaceSpec) GetMulticastFilterOk() (*string, bool) {
	if o == nil || o.MulticastFilter == nil {
		return nil, false
	}
	return o.MulticastFilter, true
}

// HasMulticastFilter returns a boolean if a field has been set.
func (o *OrchestrationManagedNamespaceSpec) HasMulticastFilter() bool {
	if o != nil && o.MulticastFilter != nil {
		return true
	}

	return false
}

// SetMulticastFilter gets a reference to the given string and assigns it to the MulticastFilter field.
func (o *OrchestrationManagedNamespaceSpec) SetMulticastFilter(v string) {
	o.MulticastFilter = &v
}

// GetOverrideVlanEnd returns the OverrideVlanEnd field value if set, zero value otherwise.
func (o *OrchestrationManagedNamespaceSpec) GetOverrideVlanEnd() int64 {
	if o == nil || o.OverrideVlanEnd == nil {
		var ret int64
		return ret
	}
	return *o.OverrideVlanEnd
}

// GetOverrideVlanEndOk returns a tuple with the OverrideVlanEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrchestrationManagedNamespaceSpec) GetOverrideVlanEndOk() (*int64, bool) {
	if o == nil || o.OverrideVlanEnd == nil {
		return nil, false
	}
	return o.OverrideVlanEnd, true
}

// HasOverrideVlanEnd returns a boolean if a field has been set.
func (o *OrchestrationManagedNamespaceSpec) HasOverrideVlanEnd() bool {
	if o != nil && o.OverrideVlanEnd != nil {
		return true
	}

	return false
}

// SetOverrideVlanEnd gets a reference to the given int64 and assigns it to the OverrideVlanEnd field.
func (o *OrchestrationManagedNamespaceSpec) SetOverrideVlanEnd(v int64) {
	o.OverrideVlanEnd = &v
}

// GetOverrideVlanStart returns the OverrideVlanStart field value if set, zero value otherwise.
func (o *OrchestrationManagedNamespaceSpec) GetOverrideVlanStart() int64 {
	if o == nil || o.OverrideVlanStart == nil {
		var ret int64
		return ret
	}
	return *o.OverrideVlanStart
}

// GetOverrideVlanStartOk returns a tuple with the OverrideVlanStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrchestrationManagedNamespaceSpec) GetOverrideVlanStartOk() (*int64, bool) {
	if o == nil || o.OverrideVlanStart == nil {
		return nil, false
	}
	return o.OverrideVlanStart, true
}

// HasOverrideVlanStart returns a boolean if a field has been set.
func (o *OrchestrationManagedNamespaceSpec) HasOverrideVlanStart() bool {
	if o != nil && o.OverrideVlanStart != nil {
		return true
	}

	return false
}

// SetOverrideVlanStart gets a reference to the given int64 and assigns it to the OverrideVlanStart field.
func (o *OrchestrationManagedNamespaceSpec) SetOverrideVlanStart(v int64) {
	o.OverrideVlanStart = &v
}

func (o OrchestrationManagedNamespaceSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DiscoveryOperation != nil {
		toSerialize["discovery-operation"] = o.DiscoveryOperation
	}
	if o.DiscoveryProtocol != nil {
		toSerialize["discovery-protocol"] = o.DiscoveryProtocol
	}
	if o.Mtu != nil {
		toSerialize["mtu"] = o.Mtu
	}
	if o.MulticastFilter != nil {
		toSerialize["multicast-filter"] = o.MulticastFilter
	}
	if o.OverrideVlanEnd != nil {
		toSerialize["override-vlan-end"] = o.OverrideVlanEnd
	}
	if o.OverrideVlanStart != nil {
		toSerialize["override-vlan-start"] = o.OverrideVlanStart
	}
	return json.Marshal(toSerialize)
}

type NullableOrchestrationManagedNamespaceSpec struct {
	value *OrchestrationManagedNamespaceSpec
	isSet bool
}

func (v NullableOrchestrationManagedNamespaceSpec) Get() *OrchestrationManagedNamespaceSpec {
	return v.value
}

func (v *NullableOrchestrationManagedNamespaceSpec) Set(val *OrchestrationManagedNamespaceSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableOrchestrationManagedNamespaceSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableOrchestrationManagedNamespaceSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrchestrationManagedNamespaceSpec(val *OrchestrationManagedNamespaceSpec) *NullableOrchestrationManagedNamespaceSpec {
	return &NullableOrchestrationManagedNamespaceSpec{value: val, isSet: true}
}

func (v NullableOrchestrationManagedNamespaceSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrchestrationManagedNamespaceSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

