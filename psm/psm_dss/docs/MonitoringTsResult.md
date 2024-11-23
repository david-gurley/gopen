# MonitoringTsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportUrl** | Pointer to **string** |  | [optional] 
**TimeWindow** | Pointer to [**MonitoringTimeWindow**](monitoringTimeWindow.md) |  | [optional] 

## Methods

### NewMonitoringTsResult

`func NewMonitoringTsResult() *MonitoringTsResult`

NewMonitoringTsResult instantiates a new MonitoringTsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringTsResultWithDefaults

`func NewMonitoringTsResultWithDefaults() *MonitoringTsResult`

NewMonitoringTsResultWithDefaults instantiates a new MonitoringTsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReportUrl

`func (o *MonitoringTsResult) GetReportUrl() string`

GetReportUrl returns the ReportUrl field if non-nil, zero value otherwise.

### GetReportUrlOk

`func (o *MonitoringTsResult) GetReportUrlOk() (*string, bool)`

GetReportUrlOk returns a tuple with the ReportUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportUrl

`func (o *MonitoringTsResult) SetReportUrl(v string)`

SetReportUrl sets ReportUrl field to given value.

### HasReportUrl

`func (o *MonitoringTsResult) HasReportUrl() bool`

HasReportUrl returns a boolean if a field has been set.

### GetTimeWindow

`func (o *MonitoringTsResult) GetTimeWindow() MonitoringTimeWindow`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *MonitoringTsResult) GetTimeWindowOk() (*MonitoringTimeWindow, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *MonitoringTsResult) SetTimeWindow(v MonitoringTimeWindow)`

SetTimeWindow sets TimeWindow field to given value.

### HasTimeWindow

`func (o *MonitoringTsResult) HasTimeWindow() bool`

HasTimeWindow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


