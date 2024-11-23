# WorkloadWorkloadIntfSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DscInterfaces** | Pointer to **[]string** | List of all DSC interfaces that can be used. The DSC interface is identified using the MAC address assigned to the DSC port. If not specified, DSCs from workload&#39;s host object are used. | [optional] 
**ExternalVlan** | Pointer to **int64** | External vlan assigned for this interface. Value should be between 0 and 4095. | [optional] 
**IpAddresses** | Pointer to **[]string** | List of all IP addresses configured on a Workload Interface. | [optional] 
**MacAddress** | Pointer to **string** | MACAddress contains the MAC address of the interface as seen by the workload. Should be a valid MAC address. | [optional] 
**MicroSegVlan** | Pointer to **int64** | Micro-segmentation vlan assigned for this interface. Value should be between 0 and 4095. | [optional] 
**Network** | Pointer to **string** | Network this interface will belong to. | [optional] 
**Vni** | Pointer to **int64** | vni is network identifier when the interface uses tunneling protocols, 0 &#x3D; not used. | [optional] 

## Methods

### NewWorkloadWorkloadIntfSpec

`func NewWorkloadWorkloadIntfSpec() *WorkloadWorkloadIntfSpec`

NewWorkloadWorkloadIntfSpec instantiates a new WorkloadWorkloadIntfSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadWorkloadIntfSpecWithDefaults

`func NewWorkloadWorkloadIntfSpecWithDefaults() *WorkloadWorkloadIntfSpec`

NewWorkloadWorkloadIntfSpecWithDefaults instantiates a new WorkloadWorkloadIntfSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDscInterfaces

`func (o *WorkloadWorkloadIntfSpec) GetDscInterfaces() []string`

GetDscInterfaces returns the DscInterfaces field if non-nil, zero value otherwise.

### GetDscInterfacesOk

`func (o *WorkloadWorkloadIntfSpec) GetDscInterfacesOk() (*[]string, bool)`

GetDscInterfacesOk returns a tuple with the DscInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDscInterfaces

`func (o *WorkloadWorkloadIntfSpec) SetDscInterfaces(v []string)`

SetDscInterfaces sets DscInterfaces field to given value.

### HasDscInterfaces

`func (o *WorkloadWorkloadIntfSpec) HasDscInterfaces() bool`

HasDscInterfaces returns a boolean if a field has been set.

### GetExternalVlan

`func (o *WorkloadWorkloadIntfSpec) GetExternalVlan() int64`

GetExternalVlan returns the ExternalVlan field if non-nil, zero value otherwise.

### GetExternalVlanOk

`func (o *WorkloadWorkloadIntfSpec) GetExternalVlanOk() (*int64, bool)`

GetExternalVlanOk returns a tuple with the ExternalVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalVlan

`func (o *WorkloadWorkloadIntfSpec) SetExternalVlan(v int64)`

SetExternalVlan sets ExternalVlan field to given value.

### HasExternalVlan

`func (o *WorkloadWorkloadIntfSpec) HasExternalVlan() bool`

HasExternalVlan returns a boolean if a field has been set.

### GetIpAddresses

`func (o *WorkloadWorkloadIntfSpec) GetIpAddresses() []string`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *WorkloadWorkloadIntfSpec) GetIpAddressesOk() (*[]string, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *WorkloadWorkloadIntfSpec) SetIpAddresses(v []string)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *WorkloadWorkloadIntfSpec) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

### GetMacAddress

`func (o *WorkloadWorkloadIntfSpec) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *WorkloadWorkloadIntfSpec) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *WorkloadWorkloadIntfSpec) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *WorkloadWorkloadIntfSpec) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetMicroSegVlan

`func (o *WorkloadWorkloadIntfSpec) GetMicroSegVlan() int64`

GetMicroSegVlan returns the MicroSegVlan field if non-nil, zero value otherwise.

### GetMicroSegVlanOk

`func (o *WorkloadWorkloadIntfSpec) GetMicroSegVlanOk() (*int64, bool)`

GetMicroSegVlanOk returns a tuple with the MicroSegVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroSegVlan

`func (o *WorkloadWorkloadIntfSpec) SetMicroSegVlan(v int64)`

SetMicroSegVlan sets MicroSegVlan field to given value.

### HasMicroSegVlan

`func (o *WorkloadWorkloadIntfSpec) HasMicroSegVlan() bool`

HasMicroSegVlan returns a boolean if a field has been set.

### GetNetwork

`func (o *WorkloadWorkloadIntfSpec) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *WorkloadWorkloadIntfSpec) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *WorkloadWorkloadIntfSpec) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *WorkloadWorkloadIntfSpec) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetVni

`func (o *WorkloadWorkloadIntfSpec) GetVni() int64`

GetVni returns the Vni field if non-nil, zero value otherwise.

### GetVniOk

`func (o *WorkloadWorkloadIntfSpec) GetVniOk() (*int64, bool)`

GetVniOk returns a tuple with the Vni field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVni

`func (o *WorkloadWorkloadIntfSpec) SetVni(v int64)`

SetVni sets Vni field to given value.

### HasVni

`func (o *WorkloadWorkloadIntfSpec) HasVni() bool`

HasVni returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


