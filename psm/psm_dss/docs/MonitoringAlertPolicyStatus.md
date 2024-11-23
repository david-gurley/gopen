# MonitoringAlertPolicyStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcknowledgedAlerts** | Pointer to **int32** | Acknowledged alerts based on this policy. | [optional] 
**OpenAlerts** | Pointer to **int32** | Open alerts based on this policy. | [optional] 
**TotalHits** | Pointer to **int32** | Total hits on this policy. | [optional] 

## Methods

### NewMonitoringAlertPolicyStatus

`func NewMonitoringAlertPolicyStatus() *MonitoringAlertPolicyStatus`

NewMonitoringAlertPolicyStatus instantiates a new MonitoringAlertPolicyStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringAlertPolicyStatusWithDefaults

`func NewMonitoringAlertPolicyStatusWithDefaults() *MonitoringAlertPolicyStatus`

NewMonitoringAlertPolicyStatusWithDefaults instantiates a new MonitoringAlertPolicyStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcknowledgedAlerts

`func (o *MonitoringAlertPolicyStatus) GetAcknowledgedAlerts() int32`

GetAcknowledgedAlerts returns the AcknowledgedAlerts field if non-nil, zero value otherwise.

### GetAcknowledgedAlertsOk

`func (o *MonitoringAlertPolicyStatus) GetAcknowledgedAlertsOk() (*int32, bool)`

GetAcknowledgedAlertsOk returns a tuple with the AcknowledgedAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgedAlerts

`func (o *MonitoringAlertPolicyStatus) SetAcknowledgedAlerts(v int32)`

SetAcknowledgedAlerts sets AcknowledgedAlerts field to given value.

### HasAcknowledgedAlerts

`func (o *MonitoringAlertPolicyStatus) HasAcknowledgedAlerts() bool`

HasAcknowledgedAlerts returns a boolean if a field has been set.

### GetOpenAlerts

`func (o *MonitoringAlertPolicyStatus) GetOpenAlerts() int32`

GetOpenAlerts returns the OpenAlerts field if non-nil, zero value otherwise.

### GetOpenAlertsOk

`func (o *MonitoringAlertPolicyStatus) GetOpenAlertsOk() (*int32, bool)`

GetOpenAlertsOk returns a tuple with the OpenAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenAlerts

`func (o *MonitoringAlertPolicyStatus) SetOpenAlerts(v int32)`

SetOpenAlerts sets OpenAlerts field to given value.

### HasOpenAlerts

`func (o *MonitoringAlertPolicyStatus) HasOpenAlerts() bool`

HasOpenAlerts returns a boolean if a field has been set.

### GetTotalHits

`func (o *MonitoringAlertPolicyStatus) GetTotalHits() int32`

GetTotalHits returns the TotalHits field if non-nil, zero value otherwise.

### GetTotalHitsOk

`func (o *MonitoringAlertPolicyStatus) GetTotalHitsOk() (*int32, bool)`

GetTotalHitsOk returns a tuple with the TotalHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalHits

`func (o *MonitoringAlertPolicyStatus) SetTotalHits(v int32)`

SetTotalHits sets TotalHits field to given value.

### HasTotalHits

`func (o *MonitoringAlertPolicyStatus) HasTotalHits() bool`

HasTotalHits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


