# ApiStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **int32** | Code is the HTTP status code. | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **[]string** | Message contains human readable form of the error. | [optional] 
**ObjectRef** | Pointer to [**ApiObjectRef**](apiObjectRef.md) |  | [optional] 
**Result** | Pointer to [**ApiStatusResult**](apiStatusResult.md) |  | [optional] 

## Methods

### NewApiStatus

`func NewApiStatus() *ApiStatus`

NewApiStatus instantiates a new ApiStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiStatusWithDefaults

`func NewApiStatusWithDefaults() *ApiStatus`

NewApiStatusWithDefaults instantiates a new ApiStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ApiStatus) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ApiStatus) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ApiStatus) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ApiStatus) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetCode

`func (o *ApiStatus) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ApiStatus) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ApiStatus) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *ApiStatus) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetKind

`func (o *ApiStatus) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ApiStatus) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ApiStatus) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ApiStatus) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMessage

`func (o *ApiStatus) GetMessage() []string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiStatus) GetMessageOk() (*[]string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiStatus) SetMessage(v []string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ApiStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetObjectRef

`func (o *ApiStatus) GetObjectRef() ApiObjectRef`

GetObjectRef returns the ObjectRef field if non-nil, zero value otherwise.

### GetObjectRefOk

`func (o *ApiStatus) GetObjectRefOk() (*ApiObjectRef, bool)`

GetObjectRefOk returns a tuple with the ObjectRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectRef

`func (o *ApiStatus) SetObjectRef(v ApiObjectRef)`

SetObjectRef sets ObjectRef field to given value.

### HasObjectRef

`func (o *ApiStatus) HasObjectRef() bool`

HasObjectRef returns a boolean if a field has been set.

### GetResult

`func (o *ApiStatus) GetResult() ApiStatusResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ApiStatus) GetResultOk() (*ApiStatusResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ApiStatus) SetResult(v ApiStatusResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *ApiStatus) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


