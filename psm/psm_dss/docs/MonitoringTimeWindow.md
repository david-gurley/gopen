# MonitoringTimeWindow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTime** | Pointer to **time.Time** | Start/Stop Time - when start time is not specified, it implies start NOW. | [optional] 
**StopTime** | Pointer to **time.Time** | Stop time - when not specified, default will be used. | [optional] 

## Methods

### NewMonitoringTimeWindow

`func NewMonitoringTimeWindow() *MonitoringTimeWindow`

NewMonitoringTimeWindow instantiates a new MonitoringTimeWindow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringTimeWindowWithDefaults

`func NewMonitoringTimeWindowWithDefaults() *MonitoringTimeWindow`

NewMonitoringTimeWindowWithDefaults instantiates a new MonitoringTimeWindow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTime

`func (o *MonitoringTimeWindow) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *MonitoringTimeWindow) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *MonitoringTimeWindow) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *MonitoringTimeWindow) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStopTime

`func (o *MonitoringTimeWindow) GetStopTime() time.Time`

GetStopTime returns the StopTime field if non-nil, zero value otherwise.

### GetStopTimeOk

`func (o *MonitoringTimeWindow) GetStopTimeOk() (*time.Time, bool)`

GetStopTimeOk returns a tuple with the StopTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopTime

`func (o *MonitoringTimeWindow) SetStopTime(v time.Time)`

SetStopTime sets StopTime field to given value.

### HasStopTime

`func (o *MonitoringTimeWindow) HasStopTime() bool`

HasStopTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


