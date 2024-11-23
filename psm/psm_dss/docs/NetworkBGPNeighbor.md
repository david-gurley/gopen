# NetworkBGPNeighbor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowAsIn** | Pointer to **int64** | AllowASIn allows local ASN to be in received routes. The value indicates how many instances are allowed. 0 disables the feature. Value should be between 0 and 255. | [optional] 
**DscAutoConfig** | Pointer to **bool** | DSCAutoConfig sets the flag that this neighbor config is to be used as a template for auto configuration. | [optional] 
**EnableAddressFamilies** | Pointer to **[]string** | Address families to enable on the neighbor. | [optional] 
**Holdtime** | Pointer to **int64** | Holdtime is time for which not receiving a keepalive message results in declaring the peer as dead. Value should be between 0 and 3600. | [optional] 
**IpAddress** | Pointer to **string** | Neighbor IP Address. Should be a valid v4 or v6 IP address. | [optional] 
**KeepaliveInterval** | Pointer to **int64** | KeepaliveInterval is time interval at which keepalive messages are sent. Value should be between 0 and 3600. | [optional] 
**MultiHop** | Pointer to **int64** | BGP Multihop configuration. Value should be between 1 and 255. | [optional] [default to 64]
**Password** | Pointer to **string** | Enable Password authentication. Disabled if the string is empty. Length of string should be between 1 and 128. | [optional] 
**RemoteAs** | Pointer to [**ApiBgpAsn**](apiBgpAsn.md) |  | [optional] 
**Shutdown** | Pointer to **bool** | Shutdown this neighbor session. | [optional] 

## Methods

### NewNetworkBGPNeighbor

`func NewNetworkBGPNeighbor() *NetworkBGPNeighbor`

NewNetworkBGPNeighbor instantiates a new NetworkBGPNeighbor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkBGPNeighborWithDefaults

`func NewNetworkBGPNeighborWithDefaults() *NetworkBGPNeighbor`

NewNetworkBGPNeighborWithDefaults instantiates a new NetworkBGPNeighbor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowAsIn

`func (o *NetworkBGPNeighbor) GetAllowAsIn() int64`

GetAllowAsIn returns the AllowAsIn field if non-nil, zero value otherwise.

### GetAllowAsInOk

`func (o *NetworkBGPNeighbor) GetAllowAsInOk() (*int64, bool)`

GetAllowAsInOk returns a tuple with the AllowAsIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAsIn

`func (o *NetworkBGPNeighbor) SetAllowAsIn(v int64)`

SetAllowAsIn sets AllowAsIn field to given value.

### HasAllowAsIn

`func (o *NetworkBGPNeighbor) HasAllowAsIn() bool`

HasAllowAsIn returns a boolean if a field has been set.

### GetDscAutoConfig

`func (o *NetworkBGPNeighbor) GetDscAutoConfig() bool`

GetDscAutoConfig returns the DscAutoConfig field if non-nil, zero value otherwise.

### GetDscAutoConfigOk

`func (o *NetworkBGPNeighbor) GetDscAutoConfigOk() (*bool, bool)`

GetDscAutoConfigOk returns a tuple with the DscAutoConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDscAutoConfig

`func (o *NetworkBGPNeighbor) SetDscAutoConfig(v bool)`

SetDscAutoConfig sets DscAutoConfig field to given value.

### HasDscAutoConfig

`func (o *NetworkBGPNeighbor) HasDscAutoConfig() bool`

HasDscAutoConfig returns a boolean if a field has been set.

### GetEnableAddressFamilies

`func (o *NetworkBGPNeighbor) GetEnableAddressFamilies() []string`

GetEnableAddressFamilies returns the EnableAddressFamilies field if non-nil, zero value otherwise.

### GetEnableAddressFamiliesOk

`func (o *NetworkBGPNeighbor) GetEnableAddressFamiliesOk() (*[]string, bool)`

GetEnableAddressFamiliesOk returns a tuple with the EnableAddressFamilies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAddressFamilies

`func (o *NetworkBGPNeighbor) SetEnableAddressFamilies(v []string)`

SetEnableAddressFamilies sets EnableAddressFamilies field to given value.

### HasEnableAddressFamilies

`func (o *NetworkBGPNeighbor) HasEnableAddressFamilies() bool`

HasEnableAddressFamilies returns a boolean if a field has been set.

### GetHoldtime

`func (o *NetworkBGPNeighbor) GetHoldtime() int64`

GetHoldtime returns the Holdtime field if non-nil, zero value otherwise.

### GetHoldtimeOk

`func (o *NetworkBGPNeighbor) GetHoldtimeOk() (*int64, bool)`

GetHoldtimeOk returns a tuple with the Holdtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldtime

`func (o *NetworkBGPNeighbor) SetHoldtime(v int64)`

SetHoldtime sets Holdtime field to given value.

### HasHoldtime

`func (o *NetworkBGPNeighbor) HasHoldtime() bool`

HasHoldtime returns a boolean if a field has been set.

### GetIpAddress

`func (o *NetworkBGPNeighbor) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *NetworkBGPNeighbor) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *NetworkBGPNeighbor) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *NetworkBGPNeighbor) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetKeepaliveInterval

`func (o *NetworkBGPNeighbor) GetKeepaliveInterval() int64`

GetKeepaliveInterval returns the KeepaliveInterval field if non-nil, zero value otherwise.

### GetKeepaliveIntervalOk

`func (o *NetworkBGPNeighbor) GetKeepaliveIntervalOk() (*int64, bool)`

GetKeepaliveIntervalOk returns a tuple with the KeepaliveInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepaliveInterval

`func (o *NetworkBGPNeighbor) SetKeepaliveInterval(v int64)`

SetKeepaliveInterval sets KeepaliveInterval field to given value.

### HasKeepaliveInterval

`func (o *NetworkBGPNeighbor) HasKeepaliveInterval() bool`

HasKeepaliveInterval returns a boolean if a field has been set.

### GetMultiHop

`func (o *NetworkBGPNeighbor) GetMultiHop() int64`

GetMultiHop returns the MultiHop field if non-nil, zero value otherwise.

### GetMultiHopOk

`func (o *NetworkBGPNeighbor) GetMultiHopOk() (*int64, bool)`

GetMultiHopOk returns a tuple with the MultiHop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiHop

`func (o *NetworkBGPNeighbor) SetMultiHop(v int64)`

SetMultiHop sets MultiHop field to given value.

### HasMultiHop

`func (o *NetworkBGPNeighbor) HasMultiHop() bool`

HasMultiHop returns a boolean if a field has been set.

### GetPassword

`func (o *NetworkBGPNeighbor) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *NetworkBGPNeighbor) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *NetworkBGPNeighbor) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *NetworkBGPNeighbor) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRemoteAs

`func (o *NetworkBGPNeighbor) GetRemoteAs() ApiBgpAsn`

GetRemoteAs returns the RemoteAs field if non-nil, zero value otherwise.

### GetRemoteAsOk

`func (o *NetworkBGPNeighbor) GetRemoteAsOk() (*ApiBgpAsn, bool)`

GetRemoteAsOk returns a tuple with the RemoteAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAs

`func (o *NetworkBGPNeighbor) SetRemoteAs(v ApiBgpAsn)`

SetRemoteAs sets RemoteAs field to given value.

### HasRemoteAs

`func (o *NetworkBGPNeighbor) HasRemoteAs() bool`

HasRemoteAs returns a boolean if a field has been set.

### GetShutdown

`func (o *NetworkBGPNeighbor) GetShutdown() bool`

GetShutdown returns the Shutdown field if non-nil, zero value otherwise.

### GetShutdownOk

`func (o *NetworkBGPNeighbor) GetShutdownOk() (*bool, bool)`

GetShutdownOk returns a tuple with the Shutdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdown

`func (o *NetworkBGPNeighbor) SetShutdown(v bool)`

SetShutdown sets Shutdown field to given value.

### HasShutdown

`func (o *NetworkBGPNeighbor) HasShutdown() bool`

HasShutdown returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


