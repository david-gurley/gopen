# StagingClearActionStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] [default to "success"]

## Methods

### NewStagingClearActionStatus

`func NewStagingClearActionStatus() *StagingClearActionStatus`

NewStagingClearActionStatus instantiates a new StagingClearActionStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStagingClearActionStatusWithDefaults

`func NewStagingClearActionStatusWithDefaults() *StagingClearActionStatus`

NewStagingClearActionStatusWithDefaults instantiates a new StagingClearActionStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *StagingClearActionStatus) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *StagingClearActionStatus) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *StagingClearActionStatus) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *StagingClearActionStatus) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatus

`func (o *StagingClearActionStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StagingClearActionStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StagingClearActionStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StagingClearActionStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


