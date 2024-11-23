# NetworkIPAMBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BootstrapIpamOptions** | Pointer to [**NetworkBootstrapIPAMOptions**](networkBootstrapIPAMOptions.md) |  | [optional] 
**IpAddress** | Pointer to **string** |  | [optional] 
**IpamOptions** | Pointer to [**NetworkIPAMOptions**](networkIPAMOptions.md) |  | [optional] 
**MacAddress** | Pointer to **string** |  | [optional] 

## Methods

### NewNetworkIPAMBinding

`func NewNetworkIPAMBinding() *NetworkIPAMBinding`

NewNetworkIPAMBinding instantiates a new NetworkIPAMBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkIPAMBindingWithDefaults

`func NewNetworkIPAMBindingWithDefaults() *NetworkIPAMBinding`

NewNetworkIPAMBindingWithDefaults instantiates a new NetworkIPAMBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBootstrapIpamOptions

`func (o *NetworkIPAMBinding) GetBootstrapIpamOptions() NetworkBootstrapIPAMOptions`

GetBootstrapIpamOptions returns the BootstrapIpamOptions field if non-nil, zero value otherwise.

### GetBootstrapIpamOptionsOk

`func (o *NetworkIPAMBinding) GetBootstrapIpamOptionsOk() (*NetworkBootstrapIPAMOptions, bool)`

GetBootstrapIpamOptionsOk returns a tuple with the BootstrapIpamOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootstrapIpamOptions

`func (o *NetworkIPAMBinding) SetBootstrapIpamOptions(v NetworkBootstrapIPAMOptions)`

SetBootstrapIpamOptions sets BootstrapIpamOptions field to given value.

### HasBootstrapIpamOptions

`func (o *NetworkIPAMBinding) HasBootstrapIpamOptions() bool`

HasBootstrapIpamOptions returns a boolean if a field has been set.

### GetIpAddress

`func (o *NetworkIPAMBinding) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *NetworkIPAMBinding) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *NetworkIPAMBinding) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *NetworkIPAMBinding) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetIpamOptions

`func (o *NetworkIPAMBinding) GetIpamOptions() NetworkIPAMOptions`

GetIpamOptions returns the IpamOptions field if non-nil, zero value otherwise.

### GetIpamOptionsOk

`func (o *NetworkIPAMBinding) GetIpamOptionsOk() (*NetworkIPAMOptions, bool)`

GetIpamOptionsOk returns a tuple with the IpamOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpamOptions

`func (o *NetworkIPAMBinding) SetIpamOptions(v NetworkIPAMOptions)`

SetIpamOptions sets IpamOptions field to given value.

### HasIpamOptions

`func (o *NetworkIPAMBinding) HasIpamOptions() bool`

HasIpamOptions returns a boolean if a field has been set.

### GetMacAddress

`func (o *NetworkIPAMBinding) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *NetworkIPAMBinding) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *NetworkIPAMBinding) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *NetworkIPAMBinding) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


