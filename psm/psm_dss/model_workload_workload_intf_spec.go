/*
 * Workload API reference
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

// WorkloadWorkloadIntfSpec Spec of a Workload interface.
type WorkloadWorkloadIntfSpec struct {
	// List of all DSC interfaces that can be used. The DSC interface is identified using the MAC address assigned to the DSC port. If not specified, DSCs from workload's host object are used.
	DscInterfaces *[]string `json:"dsc-interfaces,omitempty"`
	// External vlan assigned for this interface. Value should be between 0 and 4095.
	ExternalVlan *int64 `json:"external-vlan,omitempty"`
	// List of all IP addresses configured on a Workload Interface.
	IpAddresses *[]string `json:"ip-addresses,omitempty"`
	// MACAddress contains the MAC address of the interface as seen by the workload. Should be a valid MAC address.
	MacAddress *string `json:"mac-address,omitempty"`
	// Micro-segmentation vlan assigned for this interface. Value should be between 0 and 4095.
	MicroSegVlan *int64 `json:"micro-seg-vlan,omitempty"`
	// Network this interface will belong to.
	Network *string `json:"network,omitempty"`
	// vni is network identifier when the interface uses tunneling protocols, 0 = not used.
	Vni *int64 `json:"vni,omitempty"`
}

// NewWorkloadWorkloadIntfSpec instantiates a new WorkloadWorkloadIntfSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkloadWorkloadIntfSpec() *WorkloadWorkloadIntfSpec {
	this := WorkloadWorkloadIntfSpec{}
	return &this
}

// NewWorkloadWorkloadIntfSpecWithDefaults instantiates a new WorkloadWorkloadIntfSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkloadWorkloadIntfSpecWithDefaults() *WorkloadWorkloadIntfSpec {
	this := WorkloadWorkloadIntfSpec{}
	return &this
}

// GetDscInterfaces returns the DscInterfaces field value if set, zero value otherwise.
func (o *WorkloadWorkloadIntfSpec) GetDscInterfaces() []string {
	if o == nil || o.DscInterfaces == nil {
		var ret []string
		return ret
	}
	return *o.DscInterfaces
}

// GetDscInterfacesOk returns a tuple with the DscInterfaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadWorkloadIntfSpec) GetDscInterfacesOk() (*[]string, bool) {
	if o == nil || o.DscInterfaces == nil {
		return nil, false
	}
	return o.DscInterfaces, true
}

// HasDscInterfaces returns a boolean if a field has been set.
func (o *WorkloadWorkloadIntfSpec) HasDscInterfaces() bool {
	if o != nil && o.DscInterfaces != nil {
		return true
	}

	return false
}

// SetDscInterfaces gets a reference to the given []string and assigns it to the DscInterfaces field.
func (o *WorkloadWorkloadIntfSpec) SetDscInterfaces(v []string) {
	o.DscInterfaces = &v
}

// GetExternalVlan returns the ExternalVlan field value if set, zero value otherwise.
func (o *WorkloadWorkloadIntfSpec) GetExternalVlan() int64 {
	if o == nil || o.ExternalVlan == nil {
		var ret int64
		return ret
	}
	return *o.ExternalVlan
}

// GetExternalVlanOk returns a tuple with the ExternalVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadWorkloadIntfSpec) GetExternalVlanOk() (*int64, bool) {
	if o == nil || o.ExternalVlan == nil {
		return nil, false
	}
	return o.ExternalVlan, true
}

// HasExternalVlan returns a boolean if a field has been set.
func (o *WorkloadWorkloadIntfSpec) HasExternalVlan() bool {
	if o != nil && o.ExternalVlan != nil {
		return true
	}

	return false
}

// SetExternalVlan gets a reference to the given int64 and assigns it to the ExternalVlan field.
func (o *WorkloadWorkloadIntfSpec) SetExternalVlan(v int64) {
	o.ExternalVlan = &v
}

// GetIpAddresses returns the IpAddresses field value if set, zero value otherwise.
func (o *WorkloadWorkloadIntfSpec) GetIpAddresses() []string {
	if o == nil || o.IpAddresses == nil {
		var ret []string
		return ret
	}
	return *o.IpAddresses
}

// GetIpAddressesOk returns a tuple with the IpAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadWorkloadIntfSpec) GetIpAddressesOk() (*[]string, bool) {
	if o == nil || o.IpAddresses == nil {
		return nil, false
	}
	return o.IpAddresses, true
}

// HasIpAddresses returns a boolean if a field has been set.
func (o *WorkloadWorkloadIntfSpec) HasIpAddresses() bool {
	if o != nil && o.IpAddresses != nil {
		return true
	}

	return false
}

// SetIpAddresses gets a reference to the given []string and assigns it to the IpAddresses field.
func (o *WorkloadWorkloadIntfSpec) SetIpAddresses(v []string) {
	o.IpAddresses = &v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *WorkloadWorkloadIntfSpec) GetMacAddress() string {
	if o == nil || o.MacAddress == nil {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadWorkloadIntfSpec) GetMacAddressOk() (*string, bool) {
	if o == nil || o.MacAddress == nil {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *WorkloadWorkloadIntfSpec) HasMacAddress() bool {
	if o != nil && o.MacAddress != nil {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *WorkloadWorkloadIntfSpec) SetMacAddress(v string) {
	o.MacAddress = &v
}

// GetMicroSegVlan returns the MicroSegVlan field value if set, zero value otherwise.
func (o *WorkloadWorkloadIntfSpec) GetMicroSegVlan() int64 {
	if o == nil || o.MicroSegVlan == nil {
		var ret int64
		return ret
	}
	return *o.MicroSegVlan
}

// GetMicroSegVlanOk returns a tuple with the MicroSegVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadWorkloadIntfSpec) GetMicroSegVlanOk() (*int64, bool) {
	if o == nil || o.MicroSegVlan == nil {
		return nil, false
	}
	return o.MicroSegVlan, true
}

// HasMicroSegVlan returns a boolean if a field has been set.
func (o *WorkloadWorkloadIntfSpec) HasMicroSegVlan() bool {
	if o != nil && o.MicroSegVlan != nil {
		return true
	}

	return false
}

// SetMicroSegVlan gets a reference to the given int64 and assigns it to the MicroSegVlan field.
func (o *WorkloadWorkloadIntfSpec) SetMicroSegVlan(v int64) {
	o.MicroSegVlan = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *WorkloadWorkloadIntfSpec) GetNetwork() string {
	if o == nil || o.Network == nil {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadWorkloadIntfSpec) GetNetworkOk() (*string, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *WorkloadWorkloadIntfSpec) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *WorkloadWorkloadIntfSpec) SetNetwork(v string) {
	o.Network = &v
}

// GetVni returns the Vni field value if set, zero value otherwise.
func (o *WorkloadWorkloadIntfSpec) GetVni() int64 {
	if o == nil || o.Vni == nil {
		var ret int64
		return ret
	}
	return *o.Vni
}

// GetVniOk returns a tuple with the Vni field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadWorkloadIntfSpec) GetVniOk() (*int64, bool) {
	if o == nil || o.Vni == nil {
		return nil, false
	}
	return o.Vni, true
}

// HasVni returns a boolean if a field has been set.
func (o *WorkloadWorkloadIntfSpec) HasVni() bool {
	if o != nil && o.Vni != nil {
		return true
	}

	return false
}

// SetVni gets a reference to the given int64 and assigns it to the Vni field.
func (o *WorkloadWorkloadIntfSpec) SetVni(v int64) {
	o.Vni = &v
}

func (o WorkloadWorkloadIntfSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DscInterfaces != nil {
		toSerialize["dsc-interfaces"] = o.DscInterfaces
	}
	if o.ExternalVlan != nil {
		toSerialize["external-vlan"] = o.ExternalVlan
	}
	if o.IpAddresses != nil {
		toSerialize["ip-addresses"] = o.IpAddresses
	}
	if o.MacAddress != nil {
		toSerialize["mac-address"] = o.MacAddress
	}
	if o.MicroSegVlan != nil {
		toSerialize["micro-seg-vlan"] = o.MicroSegVlan
	}
	if o.Network != nil {
		toSerialize["network"] = o.Network
	}
	if o.Vni != nil {
		toSerialize["vni"] = o.Vni
	}
	return json.Marshal(toSerialize)
}

type NullableWorkloadWorkloadIntfSpec struct {
	value *WorkloadWorkloadIntfSpec
	isSet bool
}

func (v NullableWorkloadWorkloadIntfSpec) Get() *WorkloadWorkloadIntfSpec {
	return v.value
}

func (v *NullableWorkloadWorkloadIntfSpec) Set(val *WorkloadWorkloadIntfSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkloadWorkloadIntfSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkloadWorkloadIntfSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkloadWorkloadIntfSpec(val *WorkloadWorkloadIntfSpec) *NullableWorkloadWorkloadIntfSpec {
	return &NullableWorkloadWorkloadIntfSpec{value: val, isSet: true}
}

func (v NullableWorkloadWorkloadIntfSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkloadWorkloadIntfSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


