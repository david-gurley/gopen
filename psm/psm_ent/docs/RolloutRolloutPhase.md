# RolloutRolloutPhase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTime** | Pointer to **time.Time** | Time at which rollout completed for this node/service. | [optional] 
**Message** | Pointer to **string** | A detailed message indicating details about the transition. | [optional] 
**Name** | Pointer to **string** | Name of the Node, Service or DistributedServiceCard. | [optional] 
**NumRetries** | Pointer to **int64** | Number of retries rollout performed. | [optional] [default to 0]
**Phase** | Pointer to **string** | Phase indicates a certain rollout phase/condition. | [optional] [default to "pre-check"]
**Reason** | Pointer to **string** | The reason for the Phase last transition, if any. | [optional] 
**StartTime** | Pointer to **time.Time** | The time of starting the rollout for this node/service. This does not include the pre-check which can happen way before the actual rollout. | [optional] 

## Methods

### NewRolloutRolloutPhase

`func NewRolloutRolloutPhase() *RolloutRolloutPhase`

NewRolloutRolloutPhase instantiates a new RolloutRolloutPhase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRolloutRolloutPhaseWithDefaults

`func NewRolloutRolloutPhaseWithDefaults() *RolloutRolloutPhase`

NewRolloutRolloutPhaseWithDefaults instantiates a new RolloutRolloutPhase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTime

`func (o *RolloutRolloutPhase) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *RolloutRolloutPhase) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *RolloutRolloutPhase) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *RolloutRolloutPhase) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetMessage

`func (o *RolloutRolloutPhase) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RolloutRolloutPhase) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RolloutRolloutPhase) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *RolloutRolloutPhase) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetName

`func (o *RolloutRolloutPhase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RolloutRolloutPhase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RolloutRolloutPhase) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RolloutRolloutPhase) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumRetries

`func (o *RolloutRolloutPhase) GetNumRetries() int64`

GetNumRetries returns the NumRetries field if non-nil, zero value otherwise.

### GetNumRetriesOk

`func (o *RolloutRolloutPhase) GetNumRetriesOk() (*int64, bool)`

GetNumRetriesOk returns a tuple with the NumRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumRetries

`func (o *RolloutRolloutPhase) SetNumRetries(v int64)`

SetNumRetries sets NumRetries field to given value.

### HasNumRetries

`func (o *RolloutRolloutPhase) HasNumRetries() bool`

HasNumRetries returns a boolean if a field has been set.

### GetPhase

`func (o *RolloutRolloutPhase) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *RolloutRolloutPhase) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *RolloutRolloutPhase) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *RolloutRolloutPhase) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetReason

`func (o *RolloutRolloutPhase) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *RolloutRolloutPhase) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *RolloutRolloutPhase) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *RolloutRolloutPhase) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStartTime

`func (o *RolloutRolloutPhase) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *RolloutRolloutPhase) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *RolloutRolloutPhase) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *RolloutRolloutPhase) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


