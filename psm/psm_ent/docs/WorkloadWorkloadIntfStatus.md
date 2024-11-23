# WorkloadWorkloadIntfStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DscInterfaces** | Pointer to **[]string** | List of all DSC interfaces that can be used. The DSC interface is identified using the MAC address assigned to the DSC port. | [optional] 
**Endpoint** | Pointer to **string** | Endpoint associated with this Workload interface. | [optional] 
**ExternalVlan** | Pointer to **int64** | External vlan assigned for this interface. | [optional] 
**IpAddresses** | Pointer to **[]string** | List of all IP addresses configured and discovered on a Workload Interface. | [optional] 
**MacAddress** | Pointer to **string** | MACAddress contains the MAC address of the interface as seen by the workload. | [optional] 
**MicroSegVlan** | Pointer to **int64** | Micro-segmentation vlan used by this interface. | [optional] 
**Network** | Pointer to **string** | Network this interface belongs to. | [optional] 
**Vni** | Pointer to **int64** | vni is network identifier when the interface uses tunneling protocols, 0 &#x3D; not used. | [optional] 

## Methods

### NewWorkloadWorkloadIntfStatus

`func NewWorkloadWorkloadIntfStatus() *WorkloadWorkloadIntfStatus`

NewWorkloadWorkloadIntfStatus instantiates a new WorkloadWorkloadIntfStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadWorkloadIntfStatusWithDefaults

`func NewWorkloadWorkloadIntfStatusWithDefaults() *WorkloadWorkloadIntfStatus`

NewWorkloadWorkloadIntfStatusWithDefaults instantiates a new WorkloadWorkloadIntfStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDscInterfaces

`func (o *WorkloadWorkloadIntfStatus) GetDscInterfaces() []string`

GetDscInterfaces returns the DscInterfaces field if non-nil, zero value otherwise.

### GetDscInterfacesOk

`func (o *WorkloadWorkloadIntfStatus) GetDscInterfacesOk() (*[]string, bool)`

GetDscInterfacesOk returns a tuple with the DscInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDscInterfaces

`func (o *WorkloadWorkloadIntfStatus) SetDscInterfaces(v []string)`

SetDscInterfaces sets DscInterfaces field to given value.

### HasDscInterfaces

`func (o *WorkloadWorkloadIntfStatus) HasDscInterfaces() bool`

HasDscInterfaces returns a boolean if a field has been set.

### GetEndpoint

`func (o *WorkloadWorkloadIntfStatus) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *WorkloadWorkloadIntfStatus) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *WorkloadWorkloadIntfStatus) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *WorkloadWorkloadIntfStatus) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetExternalVlan

`func (o *WorkloadWorkloadIntfStatus) GetExternalVlan() int64`

GetExternalVlan returns the ExternalVlan field if non-nil, zero value otherwise.

### GetExternalVlanOk

`func (o *WorkloadWorkloadIntfStatus) GetExternalVlanOk() (*int64, bool)`

GetExternalVlanOk returns a tuple with the ExternalVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalVlan

`func (o *WorkloadWorkloadIntfStatus) SetExternalVlan(v int64)`

SetExternalVlan sets ExternalVlan field to given value.

### HasExternalVlan

`func (o *WorkloadWorkloadIntfStatus) HasExternalVlan() bool`

HasExternalVlan returns a boolean if a field has been set.

### GetIpAddresses

`func (o *WorkloadWorkloadIntfStatus) GetIpAddresses() []string`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *WorkloadWorkloadIntfStatus) GetIpAddressesOk() (*[]string, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *WorkloadWorkloadIntfStatus) SetIpAddresses(v []string)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *WorkloadWorkloadIntfStatus) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

### GetMacAddress

`func (o *WorkloadWorkloadIntfStatus) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *WorkloadWorkloadIntfStatus) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *WorkloadWorkloadIntfStatus) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *WorkloadWorkloadIntfStatus) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetMicroSegVlan

`func (o *WorkloadWorkloadIntfStatus) GetMicroSegVlan() int64`

GetMicroSegVlan returns the MicroSegVlan field if non-nil, zero value otherwise.

### GetMicroSegVlanOk

`func (o *WorkloadWorkloadIntfStatus) GetMicroSegVlanOk() (*int64, bool)`

GetMicroSegVlanOk returns a tuple with the MicroSegVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroSegVlan

`func (o *WorkloadWorkloadIntfStatus) SetMicroSegVlan(v int64)`

SetMicroSegVlan sets MicroSegVlan field to given value.

### HasMicroSegVlan

`func (o *WorkloadWorkloadIntfStatus) HasMicroSegVlan() bool`

HasMicroSegVlan returns a boolean if a field has been set.

### GetNetwork

`func (o *WorkloadWorkloadIntfStatus) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *WorkloadWorkloadIntfStatus) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *WorkloadWorkloadIntfStatus) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *WorkloadWorkloadIntfStatus) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetVni

`func (o *WorkloadWorkloadIntfStatus) GetVni() int64`

GetVni returns the Vni field if non-nil, zero value otherwise.

### GetVniOk

`func (o *WorkloadWorkloadIntfStatus) GetVniOk() (*int64, bool)`

GetVniOk returns a tuple with the Vni field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVni

`func (o *WorkloadWorkloadIntfStatus) SetVni(v int64)`

SetVni sets Vni field to given value.

### HasVni

`func (o *WorkloadWorkloadIntfStatus) HasVni() bool`

HasVni returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


