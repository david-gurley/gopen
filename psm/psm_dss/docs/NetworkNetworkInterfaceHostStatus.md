# NetworkNetworkInterfaceHostStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **string** | PCIE Device ID. | [optional] 
**HostIfname** | Pointer to **string** | interface name seen by the host driver. | [optional] 
**MacAddress** | Pointer to **string** | mac address of the interface. | [optional] 

## Methods

### NewNetworkNetworkInterfaceHostStatus

`func NewNetworkNetworkInterfaceHostStatus() *NetworkNetworkInterfaceHostStatus`

NewNetworkNetworkInterfaceHostStatus instantiates a new NetworkNetworkInterfaceHostStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkNetworkInterfaceHostStatusWithDefaults

`func NewNetworkNetworkInterfaceHostStatusWithDefaults() *NetworkNetworkInterfaceHostStatus`

NewNetworkNetworkInterfaceHostStatusWithDefaults instantiates a new NetworkNetworkInterfaceHostStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *NetworkNetworkInterfaceHostStatus) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *NetworkNetworkInterfaceHostStatus) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *NetworkNetworkInterfaceHostStatus) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *NetworkNetworkInterfaceHostStatus) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetHostIfname

`func (o *NetworkNetworkInterfaceHostStatus) GetHostIfname() string`

GetHostIfname returns the HostIfname field if non-nil, zero value otherwise.

### GetHostIfnameOk

`func (o *NetworkNetworkInterfaceHostStatus) GetHostIfnameOk() (*string, bool)`

GetHostIfnameOk returns a tuple with the HostIfname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostIfname

`func (o *NetworkNetworkInterfaceHostStatus) SetHostIfname(v string)`

SetHostIfname sets HostIfname field to given value.

### HasHostIfname

`func (o *NetworkNetworkInterfaceHostStatus) HasHostIfname() bool`

HasHostIfname returns a boolean if a field has been set.

### GetMacAddress

`func (o *NetworkNetworkInterfaceHostStatus) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *NetworkNetworkInterfaceHostStatus) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *NetworkNetworkInterfaceHostStatus) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *NetworkNetworkInterfaceHostStatus) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


