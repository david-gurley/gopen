# MonitoringStatsAlertPolicySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destinations** | Pointer to **[]string** | name of the alert destinations to be used to send out notification when an alert gets generated. | [optional] 
**Enable** | Pointer to **bool** | User can disable the policy by setting this field. Disabled policies will not generate any more alerts but the outstanding ones will remain as is. | [optional] [default to true]
**InstanceSelector** | Pointer to [**MonitoringInstanceSelector**](monitoringInstanceSelector.md) |  | [optional] 
**MeasurementCriteria** | Pointer to [**MonitoringMeasurementCriteria**](monitoringMeasurementCriteria.md) |  | [optional] 
**Metric** | Pointer to [**MonitoringMetricIdentifier**](monitoringMetricIdentifier.md) |  | [optional] 
**Thresholds** | Pointer to [**MonitoringThresholds**](monitoringThresholds.md) |  | [optional] 

## Methods

### NewMonitoringStatsAlertPolicySpec

`func NewMonitoringStatsAlertPolicySpec() *MonitoringStatsAlertPolicySpec`

NewMonitoringStatsAlertPolicySpec instantiates a new MonitoringStatsAlertPolicySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringStatsAlertPolicySpecWithDefaults

`func NewMonitoringStatsAlertPolicySpecWithDefaults() *MonitoringStatsAlertPolicySpec`

NewMonitoringStatsAlertPolicySpecWithDefaults instantiates a new MonitoringStatsAlertPolicySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinations

`func (o *MonitoringStatsAlertPolicySpec) GetDestinations() []string`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *MonitoringStatsAlertPolicySpec) GetDestinationsOk() (*[]string, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *MonitoringStatsAlertPolicySpec) SetDestinations(v []string)`

SetDestinations sets Destinations field to given value.

### HasDestinations

`func (o *MonitoringStatsAlertPolicySpec) HasDestinations() bool`

HasDestinations returns a boolean if a field has been set.

### GetEnable

`func (o *MonitoringStatsAlertPolicySpec) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *MonitoringStatsAlertPolicySpec) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *MonitoringStatsAlertPolicySpec) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *MonitoringStatsAlertPolicySpec) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetInstanceSelector

`func (o *MonitoringStatsAlertPolicySpec) GetInstanceSelector() MonitoringInstanceSelector`

GetInstanceSelector returns the InstanceSelector field if non-nil, zero value otherwise.

### GetInstanceSelectorOk

`func (o *MonitoringStatsAlertPolicySpec) GetInstanceSelectorOk() (*MonitoringInstanceSelector, bool)`

GetInstanceSelectorOk returns a tuple with the InstanceSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSelector

`func (o *MonitoringStatsAlertPolicySpec) SetInstanceSelector(v MonitoringInstanceSelector)`

SetInstanceSelector sets InstanceSelector field to given value.

### HasInstanceSelector

`func (o *MonitoringStatsAlertPolicySpec) HasInstanceSelector() bool`

HasInstanceSelector returns a boolean if a field has been set.

### GetMeasurementCriteria

`func (o *MonitoringStatsAlertPolicySpec) GetMeasurementCriteria() MonitoringMeasurementCriteria`

GetMeasurementCriteria returns the MeasurementCriteria field if non-nil, zero value otherwise.

### GetMeasurementCriteriaOk

`func (o *MonitoringStatsAlertPolicySpec) GetMeasurementCriteriaOk() (*MonitoringMeasurementCriteria, bool)`

GetMeasurementCriteriaOk returns a tuple with the MeasurementCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurementCriteria

`func (o *MonitoringStatsAlertPolicySpec) SetMeasurementCriteria(v MonitoringMeasurementCriteria)`

SetMeasurementCriteria sets MeasurementCriteria field to given value.

### HasMeasurementCriteria

`func (o *MonitoringStatsAlertPolicySpec) HasMeasurementCriteria() bool`

HasMeasurementCriteria returns a boolean if a field has been set.

### GetMetric

`func (o *MonitoringStatsAlertPolicySpec) GetMetric() MonitoringMetricIdentifier`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *MonitoringStatsAlertPolicySpec) GetMetricOk() (*MonitoringMetricIdentifier, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *MonitoringStatsAlertPolicySpec) SetMetric(v MonitoringMetricIdentifier)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *MonitoringStatsAlertPolicySpec) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetThresholds

`func (o *MonitoringStatsAlertPolicySpec) GetThresholds() MonitoringThresholds`

GetThresholds returns the Thresholds field if non-nil, zero value otherwise.

### GetThresholdsOk

`func (o *MonitoringStatsAlertPolicySpec) GetThresholdsOk() (*MonitoringThresholds, bool)`

GetThresholdsOk returns a tuple with the Thresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholds

`func (o *MonitoringStatsAlertPolicySpec) SetThresholds(v MonitoringThresholds)`

SetThresholds sets Thresholds field to given value.

### HasThresholds

`func (o *MonitoringStatsAlertPolicySpec) HasThresholds() bool`

HasThresholds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


