# AuthOperationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allowed** | Pointer to **bool** | Allowed indicates if Operation is authorized. | [optional] 
**Message** | Pointer to **string** | Message reports error validating Operation. | [optional] 
**Operation** | Pointer to [**AuthOperation**](authOperation.md) |  | [optional] 

## Methods

### NewAuthOperationStatus

`func NewAuthOperationStatus() *AuthOperationStatus`

NewAuthOperationStatus instantiates a new AuthOperationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthOperationStatusWithDefaults

`func NewAuthOperationStatusWithDefaults() *AuthOperationStatus`

NewAuthOperationStatusWithDefaults instantiates a new AuthOperationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowed

`func (o *AuthOperationStatus) GetAllowed() bool`

GetAllowed returns the Allowed field if non-nil, zero value otherwise.

### GetAllowedOk

`func (o *AuthOperationStatus) GetAllowedOk() (*bool, bool)`

GetAllowedOk returns a tuple with the Allowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowed

`func (o *AuthOperationStatus) SetAllowed(v bool)`

SetAllowed sets Allowed field to given value.

### HasAllowed

`func (o *AuthOperationStatus) HasAllowed() bool`

HasAllowed returns a boolean if a field has been set.

### GetMessage

`func (o *AuthOperationStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AuthOperationStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AuthOperationStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AuthOperationStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOperation

`func (o *AuthOperationStatus) GetOperation() AuthOperation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *AuthOperationStatus) GetOperationOk() (*AuthOperation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *AuthOperationStatus) SetOperation(v AuthOperation)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *AuthOperationStatus) HasOperation() bool`

HasOperation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


