# NetworkIPAMPoolInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAddressEnd** | Pointer to **string** | ending IP address of the address pool range. | [optional] 
**IpAddressStart** | Pointer to **string** | starting IP address of the address pool range. | [optional] 
**Subnet** | Pointer to **string** | subnet, in CIDR format, to which the pool belongs. | [optional] 

## Methods

### NewNetworkIPAMPoolInfo

`func NewNetworkIPAMPoolInfo() *NetworkIPAMPoolInfo`

NewNetworkIPAMPoolInfo instantiates a new NetworkIPAMPoolInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkIPAMPoolInfoWithDefaults

`func NewNetworkIPAMPoolInfoWithDefaults() *NetworkIPAMPoolInfo`

NewNetworkIPAMPoolInfoWithDefaults instantiates a new NetworkIPAMPoolInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpAddressEnd

`func (o *NetworkIPAMPoolInfo) GetIpAddressEnd() string`

GetIpAddressEnd returns the IpAddressEnd field if non-nil, zero value otherwise.

### GetIpAddressEndOk

`func (o *NetworkIPAMPoolInfo) GetIpAddressEndOk() (*string, bool)`

GetIpAddressEndOk returns a tuple with the IpAddressEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddressEnd

`func (o *NetworkIPAMPoolInfo) SetIpAddressEnd(v string)`

SetIpAddressEnd sets IpAddressEnd field to given value.

### HasIpAddressEnd

`func (o *NetworkIPAMPoolInfo) HasIpAddressEnd() bool`

HasIpAddressEnd returns a boolean if a field has been set.

### GetIpAddressStart

`func (o *NetworkIPAMPoolInfo) GetIpAddressStart() string`

GetIpAddressStart returns the IpAddressStart field if non-nil, zero value otherwise.

### GetIpAddressStartOk

`func (o *NetworkIPAMPoolInfo) GetIpAddressStartOk() (*string, bool)`

GetIpAddressStartOk returns a tuple with the IpAddressStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddressStart

`func (o *NetworkIPAMPoolInfo) SetIpAddressStart(v string)`

SetIpAddressStart sets IpAddressStart field to given value.

### HasIpAddressStart

`func (o *NetworkIPAMPoolInfo) HasIpAddressStart() bool`

HasIpAddressStart returns a boolean if a field has been set.

### GetSubnet

`func (o *NetworkIPAMPoolInfo) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *NetworkIPAMPoolInfo) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *NetworkIPAMPoolInfo) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *NetworkIPAMPoolInfo) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


