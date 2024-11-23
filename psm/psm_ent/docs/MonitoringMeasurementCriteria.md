# MonitoringMeasurementCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Function** | Pointer to **string** | Aggregate function to be applied on the metric values that were monitored over a window/interval. | [optional] [default to "none_function"]
**Window** | Pointer to **string** | The length of time the metric will be monitored/observed before running the values against thresholds for alert creation/resolution. ui-hint: Allowed values - 5m, 10m, 30m, 1h. | [optional] 

## Methods

### NewMonitoringMeasurementCriteria

`func NewMonitoringMeasurementCriteria() *MonitoringMeasurementCriteria`

NewMonitoringMeasurementCriteria instantiates a new MonitoringMeasurementCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringMeasurementCriteriaWithDefaults

`func NewMonitoringMeasurementCriteriaWithDefaults() *MonitoringMeasurementCriteria`

NewMonitoringMeasurementCriteriaWithDefaults instantiates a new MonitoringMeasurementCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunction

`func (o *MonitoringMeasurementCriteria) GetFunction() string`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *MonitoringMeasurementCriteria) GetFunctionOk() (*string, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *MonitoringMeasurementCriteria) SetFunction(v string)`

SetFunction sets Function field to given value.

### HasFunction

`func (o *MonitoringMeasurementCriteria) HasFunction() bool`

HasFunction returns a boolean if a field has been set.

### GetWindow

`func (o *MonitoringMeasurementCriteria) GetWindow() string`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *MonitoringMeasurementCriteria) GetWindowOk() (*string, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *MonitoringMeasurementCriteria) SetWindow(v string)`

SetWindow sets Window field to given value.

### HasWindow

`func (o *MonitoringMeasurementCriteria) HasWindow() bool`

HasWindow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


