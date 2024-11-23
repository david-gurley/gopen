# RolloutRolloutActionStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletionPercent** | Pointer to **int64** | Heuristic value of percentage completion of the rollout. | [optional] 
**EndTime** | Pointer to **time.Time** | End time of Rollout. | [optional] 
**PrevVersion** | Pointer to **string** | Version of the cluster before the start of rollout. | [optional] 
**StartTime** | Pointer to **time.Time** | Start time of Rollout. | [optional] 
**State** | Pointer to **string** |  | [optional] [default to "progressing"]

## Methods

### NewRolloutRolloutActionStatus

`func NewRolloutRolloutActionStatus() *RolloutRolloutActionStatus`

NewRolloutRolloutActionStatus instantiates a new RolloutRolloutActionStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRolloutRolloutActionStatusWithDefaults

`func NewRolloutRolloutActionStatusWithDefaults() *RolloutRolloutActionStatus`

NewRolloutRolloutActionStatusWithDefaults instantiates a new RolloutRolloutActionStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletionPercent

`func (o *RolloutRolloutActionStatus) GetCompletionPercent() int64`

GetCompletionPercent returns the CompletionPercent field if non-nil, zero value otherwise.

### GetCompletionPercentOk

`func (o *RolloutRolloutActionStatus) GetCompletionPercentOk() (*int64, bool)`

GetCompletionPercentOk returns a tuple with the CompletionPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionPercent

`func (o *RolloutRolloutActionStatus) SetCompletionPercent(v int64)`

SetCompletionPercent sets CompletionPercent field to given value.

### HasCompletionPercent

`func (o *RolloutRolloutActionStatus) HasCompletionPercent() bool`

HasCompletionPercent returns a boolean if a field has been set.

### GetEndTime

`func (o *RolloutRolloutActionStatus) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *RolloutRolloutActionStatus) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *RolloutRolloutActionStatus) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *RolloutRolloutActionStatus) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetPrevVersion

`func (o *RolloutRolloutActionStatus) GetPrevVersion() string`

GetPrevVersion returns the PrevVersion field if non-nil, zero value otherwise.

### GetPrevVersionOk

`func (o *RolloutRolloutActionStatus) GetPrevVersionOk() (*string, bool)`

GetPrevVersionOk returns a tuple with the PrevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevVersion

`func (o *RolloutRolloutActionStatus) SetPrevVersion(v string)`

SetPrevVersion sets PrevVersion field to given value.

### HasPrevVersion

`func (o *RolloutRolloutActionStatus) HasPrevVersion() bool`

HasPrevVersion returns a boolean if a field has been set.

### GetStartTime

`func (o *RolloutRolloutActionStatus) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *RolloutRolloutActionStatus) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *RolloutRolloutActionStatus) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *RolloutRolloutActionStatus) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetState

`func (o *RolloutRolloutActionStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RolloutRolloutActionStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RolloutRolloutActionStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *RolloutRolloutActionStatus) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


