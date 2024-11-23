# StagingBufferStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]StagingValidationError**](StagingValidationError.md) |  | [optional] 
**Items** | Pointer to [**[]StagingItem**](StagingItem.md) |  | [optional] 
**ValidationResult** | Pointer to **string** |  | [optional] [default to "success"]

## Methods

### NewStagingBufferStatus

`func NewStagingBufferStatus() *StagingBufferStatus`

NewStagingBufferStatus instantiates a new StagingBufferStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStagingBufferStatusWithDefaults

`func NewStagingBufferStatusWithDefaults() *StagingBufferStatus`

NewStagingBufferStatusWithDefaults instantiates a new StagingBufferStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *StagingBufferStatus) GetErrors() []StagingValidationError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *StagingBufferStatus) GetErrorsOk() (*[]StagingValidationError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *StagingBufferStatus) SetErrors(v []StagingValidationError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *StagingBufferStatus) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetItems

`func (o *StagingBufferStatus) GetItems() []StagingItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *StagingBufferStatus) GetItemsOk() (*[]StagingItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *StagingBufferStatus) SetItems(v []StagingItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *StagingBufferStatus) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetValidationResult

`func (o *StagingBufferStatus) GetValidationResult() string`

GetValidationResult returns the ValidationResult field if non-nil, zero value otherwise.

### GetValidationResultOk

`func (o *StagingBufferStatus) GetValidationResultOk() (*string, bool)`

GetValidationResultOk returns a tuple with the ValidationResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationResult

`func (o *StagingBufferStatus) SetValidationResult(v string)`

SetValidationResult sets ValidationResult field to given value.

### HasValidationResult

`func (o *StagingBufferStatus) HasValidationResult() bool`

HasValidationResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


