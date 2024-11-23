# NetworkIPAMConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BootstrapIpamOptions** | Pointer to [**NetworkBootstrapIPAMOptions**](networkBootstrapIPAMOptions.md) |  | [optional] 
**IpamOptions** | Pointer to [**NetworkIPAMOptions**](networkIPAMOptions.md) |  | [optional] 
**Ipv4IpamPool** | Pointer to [**[]NetworkIPAMPoolInfo**](NetworkIPAMPoolInfo.md) |  | [optional] 
**Ipv4StaticBindings** | Pointer to [**[]NetworkIPAMBinding**](NetworkIPAMBinding.md) |  | [optional] 

## Methods

### NewNetworkIPAMConfig

`func NewNetworkIPAMConfig() *NetworkIPAMConfig`

NewNetworkIPAMConfig instantiates a new NetworkIPAMConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkIPAMConfigWithDefaults

`func NewNetworkIPAMConfigWithDefaults() *NetworkIPAMConfig`

NewNetworkIPAMConfigWithDefaults instantiates a new NetworkIPAMConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBootstrapIpamOptions

`func (o *NetworkIPAMConfig) GetBootstrapIpamOptions() NetworkBootstrapIPAMOptions`

GetBootstrapIpamOptions returns the BootstrapIpamOptions field if non-nil, zero value otherwise.

### GetBootstrapIpamOptionsOk

`func (o *NetworkIPAMConfig) GetBootstrapIpamOptionsOk() (*NetworkBootstrapIPAMOptions, bool)`

GetBootstrapIpamOptionsOk returns a tuple with the BootstrapIpamOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootstrapIpamOptions

`func (o *NetworkIPAMConfig) SetBootstrapIpamOptions(v NetworkBootstrapIPAMOptions)`

SetBootstrapIpamOptions sets BootstrapIpamOptions field to given value.

### HasBootstrapIpamOptions

`func (o *NetworkIPAMConfig) HasBootstrapIpamOptions() bool`

HasBootstrapIpamOptions returns a boolean if a field has been set.

### GetIpamOptions

`func (o *NetworkIPAMConfig) GetIpamOptions() NetworkIPAMOptions`

GetIpamOptions returns the IpamOptions field if non-nil, zero value otherwise.

### GetIpamOptionsOk

`func (o *NetworkIPAMConfig) GetIpamOptionsOk() (*NetworkIPAMOptions, bool)`

GetIpamOptionsOk returns a tuple with the IpamOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpamOptions

`func (o *NetworkIPAMConfig) SetIpamOptions(v NetworkIPAMOptions)`

SetIpamOptions sets IpamOptions field to given value.

### HasIpamOptions

`func (o *NetworkIPAMConfig) HasIpamOptions() bool`

HasIpamOptions returns a boolean if a field has been set.

### GetIpv4IpamPool

`func (o *NetworkIPAMConfig) GetIpv4IpamPool() []NetworkIPAMPoolInfo`

GetIpv4IpamPool returns the Ipv4IpamPool field if non-nil, zero value otherwise.

### GetIpv4IpamPoolOk

`func (o *NetworkIPAMConfig) GetIpv4IpamPoolOk() (*[]NetworkIPAMPoolInfo, bool)`

GetIpv4IpamPoolOk returns a tuple with the Ipv4IpamPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4IpamPool

`func (o *NetworkIPAMConfig) SetIpv4IpamPool(v []NetworkIPAMPoolInfo)`

SetIpv4IpamPool sets Ipv4IpamPool field to given value.

### HasIpv4IpamPool

`func (o *NetworkIPAMConfig) HasIpv4IpamPool() bool`

HasIpv4IpamPool returns a boolean if a field has been set.

### GetIpv4StaticBindings

`func (o *NetworkIPAMConfig) GetIpv4StaticBindings() []NetworkIPAMBinding`

GetIpv4StaticBindings returns the Ipv4StaticBindings field if non-nil, zero value otherwise.

### GetIpv4StaticBindingsOk

`func (o *NetworkIPAMConfig) GetIpv4StaticBindingsOk() (*[]NetworkIPAMBinding, bool)`

GetIpv4StaticBindingsOk returns a tuple with the Ipv4StaticBindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4StaticBindings

`func (o *NetworkIPAMConfig) SetIpv4StaticBindings(v []NetworkIPAMBinding)`

SetIpv4StaticBindings sets Ipv4StaticBindings field to given value.

### HasIpv4StaticBindings

`func (o *NetworkIPAMConfig) HasIpv4StaticBindings() bool`

HasIpv4StaticBindings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


