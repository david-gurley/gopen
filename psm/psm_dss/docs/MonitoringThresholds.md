# MonitoringThresholds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | Pointer to **string** | Operator to be applied when comparing metric values against the threshold values. | [optional] [default to "less_or_equal_than"]
**Values** | Pointer to [**[]MonitoringThreshold**](MonitoringThreshold.md) | List of threshold values to be acted against. Key should be one of the alert severity. | [optional] 

## Methods

### NewMonitoringThresholds

`func NewMonitoringThresholds() *MonitoringThresholds`

NewMonitoringThresholds instantiates a new MonitoringThresholds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringThresholdsWithDefaults

`func NewMonitoringThresholdsWithDefaults() *MonitoringThresholds`

NewMonitoringThresholdsWithDefaults instantiates a new MonitoringThresholds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *MonitoringThresholds) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *MonitoringThresholds) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *MonitoringThresholds) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *MonitoringThresholds) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValues

`func (o *MonitoringThresholds) GetValues() []MonitoringThreshold`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *MonitoringThresholds) GetValuesOk() (*[]MonitoringThreshold, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *MonitoringThresholds) SetValues(v []MonitoringThreshold)`

SetValues sets Values field to given value.

### HasValues

`func (o *MonitoringThresholds) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


