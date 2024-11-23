# AuthOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] [default to "all-actions"]
**Resource** | Pointer to [**AuthResource**](authResource.md) |  | [optional] 

## Methods

### NewAuthOperation

`func NewAuthOperation() *AuthOperation`

NewAuthOperation instantiates a new AuthOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthOperationWithDefaults

`func NewAuthOperationWithDefaults() *AuthOperation`

NewAuthOperationWithDefaults instantiates a new AuthOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *AuthOperation) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AuthOperation) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AuthOperation) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *AuthOperation) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetResource

`func (o *AuthOperation) GetResource() AuthResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *AuthOperation) GetResourceOk() (*AuthResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *AuthOperation) SetResource(v AuthResource)`

SetResource sets Resource field to given value.

### HasResource

`func (o *AuthOperation) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


