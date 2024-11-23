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

// WorkloadWorkloadIntfStatus Status of a Workload interface.
type WorkloadWorkloadIntfStatus struct {
	// List of all DSC interfaces that can be used. The DSC interface is identified using the MAC address assigned to the DSC port.
	DscInterfaces *[]string `json:"dsc-interfaces,omitempty"`
	// Endpoint associated with this Workload interface.
	Endpoint *string `json:"endpoint,omitempty"`
	// External vlan assigned for this interface.
	ExternalVlan *int64 `json:"external-vlan,omitempty"`
	InterfaceMigrationStatus *WorkloadInterfaceMigrationStatus `json:"interface-migration-status,omitempty"`
	// List of all IP addresses configured and discovered on a Workload Interface.
	IpAddresses *[]string `json:"ip-addresses,omitempty"`
	// MACAddress contains the MAC address of the interface as seen by the workload.
	MacAddress *string `json:"mac-address,omitempty"`
	// Micro-segmentation vlan used by this interface.
	MicroSegVlan *int64 `json:"micro-seg-vlan,omitempty"`
	// Network this interface belongs to.
	Network *string `json:"network,omitempty"`
	// vni is network identifier when the interface uses tunneling protocols, 0 = not used.
	Vni *int64 `json:"vni,omitempty"`
}

// NewWorkloadWorkloadIntfStatus instantiates a new WorkloadWorkloadIntfStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkloadWorkloadIntfStatus() *WorkloadWorkloadIntfStatus {
	this := WorkloadWorkloadIntfStatus{}
	return &this
}

// NewWorkloadWorkloadIntfStatusWithDefaults instantiates a new WorkloadWorkloadIntfStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkloadWorkloadIntfStatusWithDefaults() *WorkloadWorkloadIntfStatus {
	this := WorkloadWorkloadIntfStatus{}
	return &this
}

// GetDscInterfaces returns the DscInterfaces field value if set, zero value otherwise.
func (o *WorkloadWorkloadIntfStatus) GetDscInterfaces() []string {
	if o == nil || o.DscInterfaces == nil {
		var ret []string
		return ret
	}
	return *o.DscInterfaces
}

// GetDscInterfacesOk returns a tuple with the DscInterfaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadWorkloadIntfStatus) GetDscInterfacesOk() (*[]string, bool) {
	if o == nil || o.DscInterfaces == nil {
		return nil, false
	}
	return o.DscInterfaces, true
}

// HasDscInterfaces returns a boolean if a field has been set.
func (o *WorkloadWorkloadIntfStatus) HasDscInterfaces() bool {
	if o != nil && o.DscInterfaces != nil {
		return true
	}

	return false
}

// SetDscInterfaces gets a reference to the given []string and assigns it to the DscInterfaces field.
func (o *WorkloadWorkloadIntfStatus) SetDscInterfaces(v []string) {
	o.DscInterfaces = &v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *WorkloadWorkloadIntfStatus) GetEndpoint() string {
	if o == nil || o.Endpoint == nil {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadWorkloadIntfStatus) GetEndpointOk() (*string, bool) {
	if o == nil || o.Endpoint == nil {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *WorkloadWorkloadIntfStatus) HasEndpoint() bool {
	if o != nil && o.Endpoint != nil {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *WorkloadWorkloadIntfStatus) SetEndpoint(v string) {
	o.Endpoint = &v
}

// GetExternalVlan returns the ExternalVlan field value if set, zero value otherwise.
func (o *WorkloadWorkloadIntfStatus) GetExternalVlan() int64 {
	if o == nil || o.ExternalVlan == nil {
		var ret int64
		return ret
	}
	return *o.ExternalVlan
}

// GetExternalVlanOk returns a tuple with the ExternalVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadWorkloadIntfStatus) GetExternalVlanOk() (*int64, bool) {
	if o == nil || o.ExternalVlan == nil {
		return nil, false
	}
	return o.ExternalVlan, true
}

// HasExternalVlan returns a boolean if a field has been set.
func (o *WorkloadWorkloadIntfStatus) HasExternalVlan() bool {
	if o != nil && o.ExternalVlan != nil {
		return true
	}

	return false
}

// SetExternalVlan gets a reference to the given int64 and assigns it to the ExternalVlan field.
func (o *WorkloadWorkloadIntfStatus) SetExternalVlan(v int64) {
	o.ExternalVlan = &v
}

// GetInterfaceMigrationStatus returns the InterfaceMigrationStatus field value if set, zero value otherwise.
func (o *WorkloadWorkloadIntfStatus) GetInterfaceMigrationStatus() WorkloadInterfaceMigrationStatus {
	if o == nil || o.InterfaceMigrationStatus == nil {
		var ret WorkloadInterfaceMigrationStatus
		return ret
	}
	return *o.InterfaceMigrationStatus
}

// GetInterfaceMigrationStatusOk returns a tuple with the InterfaceMigrationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadWorkloadIntfStatus) GetInterfaceMigrationStatusOk() (*WorkloadInterfaceMigrationStatus, bool) {
	if o == nil || o.InterfaceMigrationStatus == nil {
		return nil, false
	}
	return o.InterfaceMigrationStatus, true
}

// HasInterfaceMigrationStatus returns a boolean if a field has been set.
func (o *WorkloadWorkloadIntfStatus) HasInterfaceMigrationStatus() bool {
	if o != nil && o.InterfaceMigrationStatus != nil {
		return true
	}

	return false
}

// SetInterfaceMigrationStatus gets a reference to the given WorkloadInterfaceMigrationStatus and assigns it to the InterfaceMigrationStatus field.
func (o *WorkloadWorkloadIntfStatus) SetInterfaceMigrationStatus(v WorkloadInterfaceMigrationStatus) {
	o.InterfaceMigrationStatus = &v
}

// GetIpAddresses returns the IpAddresses field value if set, zero value otherwise.
func (o *WorkloadWorkloadIntfStatus) GetIpAddresses() []string {
	if o == nil || o.IpAddresses == nil {
		var ret []string
		return ret
	}
	return *o.IpAddresses
}

// GetIpAddressesOk returns a tuple with the IpAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadWorkloadIntfStatus) GetIpAddressesOk() (*[]string, bool) {
	if o == nil || o.IpAddresses == nil {
		return nil, false
	}
	return o.IpAddresses, true
}

// HasIpAddresses returns a boolean if a field has been set.
func (o *WorkloadWorkloadIntfStatus) HasIpAddresses() bool {
	if o != nil && o.IpAddresses != nil {
		return true
	}

	return false
}

// SetIpAddresses gets a reference to the given []string and assigns it to the IpAddresses field.
func (o *WorkloadWorkloadIntfStatus) SetIpAddresses(v []string) {
	o.IpAddresses = &v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *WorkloadWorkloadIntfStatus) GetMacAddress() string {
	if o == nil || o.MacAddress == nil {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadWorkloadIntfStatus) GetMacAddressOk() (*string, bool) {
	if o == nil || o.MacAddress == nil {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *WorkloadWorkloadIntfStatus) HasMacAddress() bool {
	if o != nil && o.MacAddress != nil {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *WorkloadWorkloadIntfStatus) SetMacAddress(v string) {
	o.MacAddress = &v
}

// GetMicroSegVlan returns the MicroSegVlan field value if set, zero value otherwise.
func (o *WorkloadWorkloadIntfStatus) GetMicroSegVlan() int64 {
	if o == nil || o.MicroSegVlan == nil {
		var ret int64
		return ret
	}
	return *o.MicroSegVlan
}

// GetMicroSegVlanOk returns a tuple with the MicroSegVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadWorkloadIntfStatus) GetMicroSegVlanOk() (*int64, bool) {
	if o == nil || o.MicroSegVlan == nil {
		return nil, false
	}
	return o.MicroSegVlan, true
}

// HasMicroSegVlan returns a boolean if a field has been set.
func (o *WorkloadWorkloadIntfStatus) HasMicroSegVlan() bool {
	if o != nil && o.MicroSegVlan != nil {
		return true
	}

	return false
}

// SetMicroSegVlan gets a reference to the given int64 and assigns it to the MicroSegVlan field.
func (o *WorkloadWorkloadIntfStatus) SetMicroSegVlan(v int64) {
	o.MicroSegVlan = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *WorkloadWorkloadIntfStatus) GetNetwork() string {
	if o == nil || o.Network == nil {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadWorkloadIntfStatus) GetNetworkOk() (*string, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *WorkloadWorkloadIntfStatus) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *WorkloadWorkloadIntfStatus) SetNetwork(v string) {
	o.Network = &v
}

// GetVni returns the Vni field value if set, zero value otherwise.
func (o *WorkloadWorkloadIntfStatus) GetVni() int64 {
	if o == nil || o.Vni == nil {
		var ret int64
		return ret
	}
	return *o.Vni
}

// GetVniOk returns a tuple with the Vni field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadWorkloadIntfStatus) GetVniOk() (*int64, bool) {
	if o == nil || o.Vni == nil {
		return nil, false
	}
	return o.Vni, true
}

// HasVni returns a boolean if a field has been set.
func (o *WorkloadWorkloadIntfStatus) HasVni() bool {
	if o != nil && o.Vni != nil {
		return true
	}

	return false
}

// SetVni gets a reference to the given int64 and assigns it to the Vni field.
func (o *WorkloadWorkloadIntfStatus) SetVni(v int64) {
	o.Vni = &v
}

func (o WorkloadWorkloadIntfStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DscInterfaces != nil {
		toSerialize["dsc-interfaces"] = o.DscInterfaces
	}
	if o.Endpoint != nil {
		toSerialize["endpoint"] = o.Endpoint
	}
	if o.ExternalVlan != nil {
		toSerialize["external-vlan"] = o.ExternalVlan
	}
	if o.InterfaceMigrationStatus != nil {
		toSerialize["interface-migration-status"] = o.InterfaceMigrationStatus
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

type NullableWorkloadWorkloadIntfStatus struct {
	value *WorkloadWorkloadIntfStatus
	isSet bool
}

func (v NullableWorkloadWorkloadIntfStatus) Get() *WorkloadWorkloadIntfStatus {
	return v.value
}

func (v *NullableWorkloadWorkloadIntfStatus) Set(val *WorkloadWorkloadIntfStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkloadWorkloadIntfStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkloadWorkloadIntfStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkloadWorkloadIntfStatus(val *WorkloadWorkloadIntfStatus) *NullableWorkloadWorkloadIntfStatus {
	return &NullableWorkloadWorkloadIntfStatus{value: val, isSet: true}
}

func (v NullableWorkloadWorkloadIntfStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkloadWorkloadIntfStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


