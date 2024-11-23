# MonitoringMirrorSessionStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PropagationStatus** | Pointer to [**MonitoringPropagationStatus**](monitoringPropagationStatus.md) |  | [optional] 
**ScheduleState** | Pointer to **string** |  | [optional] [default to "none"]
**StartedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewMonitoringMirrorSessionStatus

`func NewMonitoringMirrorSessionStatus() *MonitoringMirrorSessionStatus`

NewMonitoringMirrorSessionStatus instantiates a new MonitoringMirrorSessionStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringMirrorSessionStatusWithDefaults

`func NewMonitoringMirrorSessionStatusWithDefaults() *MonitoringMirrorSessionStatus`

NewMonitoringMirrorSessionStatusWithDefaults instantiates a new MonitoringMirrorSessionStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPropagationStatus

`func (o *MonitoringMirrorSessionStatus) GetPropagationStatus() MonitoringPropagationStatus`

GetPropagationStatus returns the PropagationStatus field if non-nil, zero value otherwise.

### GetPropagationStatusOk

`func (o *MonitoringMirrorSessionStatus) GetPropagationStatusOk() (*MonitoringPropagationStatus, bool)`

GetPropagationStatusOk returns a tuple with the PropagationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropagationStatus

`func (o *MonitoringMirrorSessionStatus) SetPropagationStatus(v MonitoringPropagationStatus)`

SetPropagationStatus sets PropagationStatus field to given value.

### HasPropagationStatus

`func (o *MonitoringMirrorSessionStatus) HasPropagationStatus() bool`

HasPropagationStatus returns a boolean if a field has been set.

### GetScheduleState

`func (o *MonitoringMirrorSessionStatus) GetScheduleState() string`

GetScheduleState returns the ScheduleState field if non-nil, zero value otherwise.

### GetScheduleStateOk

`func (o *MonitoringMirrorSessionStatus) GetScheduleStateOk() (*string, bool)`

GetScheduleStateOk returns a tuple with the ScheduleState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleState

`func (o *MonitoringMirrorSessionStatus) SetScheduleState(v string)`

SetScheduleState sets ScheduleState field to given value.

### HasScheduleState

`func (o *MonitoringMirrorSessionStatus) HasScheduleState() bool`

HasScheduleState returns a boolean if a field has been set.

### GetStartedAt

`func (o *MonitoringMirrorSessionStatus) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *MonitoringMirrorSessionStatus) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *MonitoringMirrorSessionStatus) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *MonitoringMirrorSessionStatus) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


