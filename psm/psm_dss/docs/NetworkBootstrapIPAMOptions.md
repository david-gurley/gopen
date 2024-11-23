# NetworkBootstrapIPAMOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Controllers** | Pointer to **[]string** |  | [optional] 
**InterfaceIps** | Pointer to [**[]NetworkInterfaceIP**](NetworkInterfaceIP.md) |  | [optional] 

## Methods

### NewNetworkBootstrapIPAMOptions

`func NewNetworkBootstrapIPAMOptions() *NetworkBootstrapIPAMOptions`

NewNetworkBootstrapIPAMOptions instantiates a new NetworkBootstrapIPAMOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkBootstrapIPAMOptionsWithDefaults

`func NewNetworkBootstrapIPAMOptionsWithDefaults() *NetworkBootstrapIPAMOptions`

NewNetworkBootstrapIPAMOptionsWithDefaults instantiates a new NetworkBootstrapIPAMOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetControllers

`func (o *NetworkBootstrapIPAMOptions) GetControllers() []string`

GetControllers returns the Controllers field if non-nil, zero value otherwise.

### GetControllersOk

`func (o *NetworkBootstrapIPAMOptions) GetControllersOk() (*[]string, bool)`

GetControllersOk returns a tuple with the Controllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllers

`func (o *NetworkBootstrapIPAMOptions) SetControllers(v []string)`

SetControllers sets Controllers field to given value.

### HasControllers

`func (o *NetworkBootstrapIPAMOptions) HasControllers() bool`

HasControllers returns a boolean if a field has been set.

### GetInterfaceIps

`func (o *NetworkBootstrapIPAMOptions) GetInterfaceIps() []NetworkInterfaceIP`

GetInterfaceIps returns the InterfaceIps field if non-nil, zero value otherwise.

### GetInterfaceIpsOk

`func (o *NetworkBootstrapIPAMOptions) GetInterfaceIpsOk() (*[]NetworkInterfaceIP, bool)`

GetInterfaceIpsOk returns a tuple with the InterfaceIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceIps

`func (o *NetworkBootstrapIPAMOptions) SetInterfaceIps(v []NetworkInterfaceIP)`

SetInterfaceIps sets InterfaceIps field to given value.

### HasInterfaceIps

`func (o *NetworkBootstrapIPAMOptions) HasInterfaceIps() bool`

HasInterfaceIps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


