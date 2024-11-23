# MonitoringThreshold

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RaiseValue** | Pointer to **string** | Raise/Create an alert when the threshold reaches this value. | [optional] 
**Severity** | Pointer to **string** | Severity of the alert to be created. | [optional] [default to "info"]

## Methods

### NewMonitoringThreshold

`func NewMonitoringThreshold() *MonitoringThreshold`

NewMonitoringThreshold instantiates a new MonitoringThreshold object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringThresholdWithDefaults

`func NewMonitoringThresholdWithDefaults() *MonitoringThreshold`

NewMonitoringThresholdWithDefaults instantiates a new MonitoringThreshold object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRaiseValue

`func (o *MonitoringThreshold) GetRaiseValue() string`

GetRaiseValue returns the RaiseValue field if non-nil, zero value otherwise.

### GetRaiseValueOk

`func (o *MonitoringThreshold) GetRaiseValueOk() (*string, bool)`

GetRaiseValueOk returns a tuple with the RaiseValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaiseValue

`func (o *MonitoringThreshold) SetRaiseValue(v string)`

SetRaiseValue sets RaiseValue field to given value.

### HasRaiseValue

`func (o *MonitoringThreshold) HasRaiseValue() bool`

HasRaiseValue returns a boolean if a field has been set.

### GetSeverity

`func (o *MonitoringThreshold) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *MonitoringThreshold) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *MonitoringThreshold) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *MonitoringThreshold) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


