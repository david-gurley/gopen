# WorkloadWorkloadMigrationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletedAt** | Pointer to **time.Time** | migration completion time. | [optional] 
**Stage** | Pointer to **string** | Controller&#39;s migration stage. | [optional] [default to "migration-none"]
**StartedAt** | Pointer to **time.Time** | migration start time. | [optional] 
**Status** | Pointer to **string** | The status from the dataplane performing migration. | [optional] [default to "none"]

## Methods

### NewWorkloadWorkloadMigrationStatus

`func NewWorkloadWorkloadMigrationStatus() *WorkloadWorkloadMigrationStatus`

NewWorkloadWorkloadMigrationStatus instantiates a new WorkloadWorkloadMigrationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadWorkloadMigrationStatusWithDefaults

`func NewWorkloadWorkloadMigrationStatusWithDefaults() *WorkloadWorkloadMigrationStatus`

NewWorkloadWorkloadMigrationStatusWithDefaults instantiates a new WorkloadWorkloadMigrationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletedAt

`func (o *WorkloadWorkloadMigrationStatus) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *WorkloadWorkloadMigrationStatus) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *WorkloadWorkloadMigrationStatus) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *WorkloadWorkloadMigrationStatus) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetStage

`func (o *WorkloadWorkloadMigrationStatus) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *WorkloadWorkloadMigrationStatus) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *WorkloadWorkloadMigrationStatus) SetStage(v string)`

SetStage sets Stage field to given value.

### HasStage

`func (o *WorkloadWorkloadMigrationStatus) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetStartedAt

`func (o *WorkloadWorkloadMigrationStatus) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *WorkloadWorkloadMigrationStatus) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *WorkloadWorkloadMigrationStatus) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *WorkloadWorkloadMigrationStatus) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStatus

`func (o *WorkloadWorkloadMigrationStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkloadWorkloadMigrationStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkloadWorkloadMigrationStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkloadWorkloadMigrationStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


