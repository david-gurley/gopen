# NetworkBGPAuthStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAddress** | Pointer to **string** | Neighbor IP Address. | [optional] 
**RemoteAs** | Pointer to [**ApiBgpAsn**](apiBgpAsn.md) |  | [optional] 
**Status** | Pointer to **string** | Authentication status. | [optional] [default to "disabled"]

## Methods

### NewNetworkBGPAuthStatus

`func NewNetworkBGPAuthStatus() *NetworkBGPAuthStatus`

NewNetworkBGPAuthStatus instantiates a new NetworkBGPAuthStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkBGPAuthStatusWithDefaults

`func NewNetworkBGPAuthStatusWithDefaults() *NetworkBGPAuthStatus`

NewNetworkBGPAuthStatusWithDefaults instantiates a new NetworkBGPAuthStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpAddress

`func (o *NetworkBGPAuthStatus) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *NetworkBGPAuthStatus) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *NetworkBGPAuthStatus) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *NetworkBGPAuthStatus) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetRemoteAs

`func (o *NetworkBGPAuthStatus) GetRemoteAs() ApiBgpAsn`

GetRemoteAs returns the RemoteAs field if non-nil, zero value otherwise.

### GetRemoteAsOk

`func (o *NetworkBGPAuthStatus) GetRemoteAsOk() (*ApiBgpAsn, bool)`

GetRemoteAsOk returns a tuple with the RemoteAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAs

`func (o *NetworkBGPAuthStatus) SetRemoteAs(v ApiBgpAsn)`

SetRemoteAs sets RemoteAs field to given value.

### HasRemoteAs

`func (o *NetworkBGPAuthStatus) HasRemoteAs() bool`

HasRemoteAs returns a boolean if a field has been set.

### GetStatus

`func (o *NetworkBGPAuthStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkBGPAuthStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkBGPAuthStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NetworkBGPAuthStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


