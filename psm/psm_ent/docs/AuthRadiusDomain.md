# AuthRadiusDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NasId** | Pointer to **string** | NasID is a string identifying the NAS(API Gw) originating the Access-Request. | [optional] 
**Servers** | Pointer to [**[]AuthRadiusServer**](AuthRadiusServer.md) |  | [optional] 
**Tag** | Pointer to **string** | Tag to group domains for authentication. | [optional] 

## Methods

### NewAuthRadiusDomain

`func NewAuthRadiusDomain() *AuthRadiusDomain`

NewAuthRadiusDomain instantiates a new AuthRadiusDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthRadiusDomainWithDefaults

`func NewAuthRadiusDomainWithDefaults() *AuthRadiusDomain`

NewAuthRadiusDomainWithDefaults instantiates a new AuthRadiusDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNasId

`func (o *AuthRadiusDomain) GetNasId() string`

GetNasId returns the NasId field if non-nil, zero value otherwise.

### GetNasIdOk

`func (o *AuthRadiusDomain) GetNasIdOk() (*string, bool)`

GetNasIdOk returns a tuple with the NasId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNasId

`func (o *AuthRadiusDomain) SetNasId(v string)`

SetNasId sets NasId field to given value.

### HasNasId

`func (o *AuthRadiusDomain) HasNasId() bool`

HasNasId returns a boolean if a field has been set.

### GetServers

`func (o *AuthRadiusDomain) GetServers() []AuthRadiusServer`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *AuthRadiusDomain) GetServersOk() (*[]AuthRadiusServer, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *AuthRadiusDomain) SetServers(v []AuthRadiusServer)`

SetServers sets Servers field to given value.

### HasServers

`func (o *AuthRadiusDomain) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetTag

`func (o *AuthRadiusDomain) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *AuthRadiusDomain) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *AuthRadiusDomain) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *AuthRadiusDomain) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


