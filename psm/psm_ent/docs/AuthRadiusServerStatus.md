# AuthRadiusServerStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Message contains error message in case of failed check or a success message. | [optional] 
**NasId** | Pointer to **string** | NasID is a string identifying the NAS(API Gw) originating the Access-Request. | [optional] 
**Result** | Pointer to **string** | Result indicates if radius check was successful. | [optional] [default to "connect-success"]
**Server** | Pointer to [**AuthRadiusServer**](authRadiusServer.md) |  | [optional] 

## Methods

### NewAuthRadiusServerStatus

`func NewAuthRadiusServerStatus() *AuthRadiusServerStatus`

NewAuthRadiusServerStatus instantiates a new AuthRadiusServerStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthRadiusServerStatusWithDefaults

`func NewAuthRadiusServerStatusWithDefaults() *AuthRadiusServerStatus`

NewAuthRadiusServerStatusWithDefaults instantiates a new AuthRadiusServerStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *AuthRadiusServerStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AuthRadiusServerStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AuthRadiusServerStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AuthRadiusServerStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetNasId

`func (o *AuthRadiusServerStatus) GetNasId() string`

GetNasId returns the NasId field if non-nil, zero value otherwise.

### GetNasIdOk

`func (o *AuthRadiusServerStatus) GetNasIdOk() (*string, bool)`

GetNasIdOk returns a tuple with the NasId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNasId

`func (o *AuthRadiusServerStatus) SetNasId(v string)`

SetNasId sets NasId field to given value.

### HasNasId

`func (o *AuthRadiusServerStatus) HasNasId() bool`

HasNasId returns a boolean if a field has been set.

### GetResult

`func (o *AuthRadiusServerStatus) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *AuthRadiusServerStatus) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *AuthRadiusServerStatus) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *AuthRadiusServerStatus) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetServer

`func (o *AuthRadiusServerStatus) GetServer() AuthRadiusServer`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *AuthRadiusServerStatus) GetServerOk() (*AuthRadiusServer, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *AuthRadiusServerStatus) SetServer(v AuthRadiusServer)`

SetServer sets Server field to given value.

### HasServer

`func (o *AuthRadiusServerStatus) HasServer() bool`

HasServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


