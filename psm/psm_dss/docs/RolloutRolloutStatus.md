# RolloutRolloutStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletionPercent** | Pointer to **int64** | Heuristic value of percentage completion of the rollout. | [optional] 
**ControllerNodesStatus** | Pointer to [**[]RolloutRolloutPhase**](RolloutRolloutPhase.md) | Rollout status of Controller Node. | [optional] 
**ControllerServicesStatus** | Pointer to [**[]RolloutRolloutPhase**](RolloutRolloutPhase.md) | Rollout status of Various Controller Services. | [optional] 
**DscsStatus** | Pointer to [**[]RolloutRolloutPhase**](RolloutRolloutPhase.md) | Rollout status of DistributedServiceCards in the cluster. Has entries for DistributedServiceCards on Controller nodes as well as workload nodes The entries are group by parallelism based on the order-constraints and max-parallel specified by the user. | [optional] 
**EndTime** | Pointer to **time.Time** | End time of Rollout. | [optional] 
**PrevVersion** | Pointer to **string** | Version of the cluster before the start of rollout. | [optional] 
**Reason** | Pointer to **string** | details the reason for overall Failure or Suspend. | [optional] 
**StartTime** | Pointer to **time.Time** | Start time of Rollout. | [optional] 
**State** | Pointer to **string** |  | [optional] [default to "progressing"]

## Methods

### NewRolloutRolloutStatus

`func NewRolloutRolloutStatus() *RolloutRolloutStatus`

NewRolloutRolloutStatus instantiates a new RolloutRolloutStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRolloutRolloutStatusWithDefaults

`func NewRolloutRolloutStatusWithDefaults() *RolloutRolloutStatus`

NewRolloutRolloutStatusWithDefaults instantiates a new RolloutRolloutStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletionPercent

`func (o *RolloutRolloutStatus) GetCompletionPercent() int64`

GetCompletionPercent returns the CompletionPercent field if non-nil, zero value otherwise.

### GetCompletionPercentOk

`func (o *RolloutRolloutStatus) GetCompletionPercentOk() (*int64, bool)`

GetCompletionPercentOk returns a tuple with the CompletionPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionPercent

`func (o *RolloutRolloutStatus) SetCompletionPercent(v int64)`

SetCompletionPercent sets CompletionPercent field to given value.

### HasCompletionPercent

`func (o *RolloutRolloutStatus) HasCompletionPercent() bool`

HasCompletionPercent returns a boolean if a field has been set.

### GetControllerNodesStatus

`func (o *RolloutRolloutStatus) GetControllerNodesStatus() []RolloutRolloutPhase`

GetControllerNodesStatus returns the ControllerNodesStatus field if non-nil, zero value otherwise.

### GetControllerNodesStatusOk

`func (o *RolloutRolloutStatus) GetControllerNodesStatusOk() (*[]RolloutRolloutPhase, bool)`

GetControllerNodesStatusOk returns a tuple with the ControllerNodesStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerNodesStatus

`func (o *RolloutRolloutStatus) SetControllerNodesStatus(v []RolloutRolloutPhase)`

SetControllerNodesStatus sets ControllerNodesStatus field to given value.

### HasControllerNodesStatus

`func (o *RolloutRolloutStatus) HasControllerNodesStatus() bool`

HasControllerNodesStatus returns a boolean if a field has been set.

### GetControllerServicesStatus

`func (o *RolloutRolloutStatus) GetControllerServicesStatus() []RolloutRolloutPhase`

GetControllerServicesStatus returns the ControllerServicesStatus field if non-nil, zero value otherwise.

### GetControllerServicesStatusOk

`func (o *RolloutRolloutStatus) GetControllerServicesStatusOk() (*[]RolloutRolloutPhase, bool)`

GetControllerServicesStatusOk returns a tuple with the ControllerServicesStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerServicesStatus

`func (o *RolloutRolloutStatus) SetControllerServicesStatus(v []RolloutRolloutPhase)`

SetControllerServicesStatus sets ControllerServicesStatus field to given value.

### HasControllerServicesStatus

`func (o *RolloutRolloutStatus) HasControllerServicesStatus() bool`

HasControllerServicesStatus returns a boolean if a field has been set.

### GetDscsStatus

`func (o *RolloutRolloutStatus) GetDscsStatus() []RolloutRolloutPhase`

GetDscsStatus returns the DscsStatus field if non-nil, zero value otherwise.

### GetDscsStatusOk

`func (o *RolloutRolloutStatus) GetDscsStatusOk() (*[]RolloutRolloutPhase, bool)`

GetDscsStatusOk returns a tuple with the DscsStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDscsStatus

`func (o *RolloutRolloutStatus) SetDscsStatus(v []RolloutRolloutPhase)`

SetDscsStatus sets DscsStatus field to given value.

### HasDscsStatus

`func (o *RolloutRolloutStatus) HasDscsStatus() bool`

HasDscsStatus returns a boolean if a field has been set.

### GetEndTime

`func (o *RolloutRolloutStatus) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *RolloutRolloutStatus) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *RolloutRolloutStatus) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *RolloutRolloutStatus) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetPrevVersion

`func (o *RolloutRolloutStatus) GetPrevVersion() string`

GetPrevVersion returns the PrevVersion field if non-nil, zero value otherwise.

### GetPrevVersionOk

`func (o *RolloutRolloutStatus) GetPrevVersionOk() (*string, bool)`

GetPrevVersionOk returns a tuple with the PrevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevVersion

`func (o *RolloutRolloutStatus) SetPrevVersion(v string)`

SetPrevVersion sets PrevVersion field to given value.

### HasPrevVersion

`func (o *RolloutRolloutStatus) HasPrevVersion() bool`

HasPrevVersion returns a boolean if a field has been set.

### GetReason

`func (o *RolloutRolloutStatus) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *RolloutRolloutStatus) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *RolloutRolloutStatus) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *RolloutRolloutStatus) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStartTime

`func (o *RolloutRolloutStatus) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *RolloutRolloutStatus) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *RolloutRolloutStatus) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *RolloutRolloutStatus) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetState

`func (o *RolloutRolloutStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RolloutRolloutStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RolloutRolloutStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *RolloutRolloutStatus) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


