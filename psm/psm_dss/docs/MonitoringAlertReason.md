# MonitoringAlertReason

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertPolicyId** | Pointer to **string** | Alert Policy ID that matched. | [optional] 
**MatchedRequirements** | Pointer to [**[]MonitoringMatchedRequirement**](MonitoringMatchedRequirement.md) | List of requirements from the alert policy with it&#39;s matched value. | [optional] 

## Methods

### NewMonitoringAlertReason

`func NewMonitoringAlertReason() *MonitoringAlertReason`

NewMonitoringAlertReason instantiates a new MonitoringAlertReason object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringAlertReasonWithDefaults

`func NewMonitoringAlertReasonWithDefaults() *MonitoringAlertReason`

NewMonitoringAlertReasonWithDefaults instantiates a new MonitoringAlertReason object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertPolicyId

`func (o *MonitoringAlertReason) GetAlertPolicyId() string`

GetAlertPolicyId returns the AlertPolicyId field if non-nil, zero value otherwise.

### GetAlertPolicyIdOk

`func (o *MonitoringAlertReason) GetAlertPolicyIdOk() (*string, bool)`

GetAlertPolicyIdOk returns a tuple with the AlertPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertPolicyId

`func (o *MonitoringAlertReason) SetAlertPolicyId(v string)`

SetAlertPolicyId sets AlertPolicyId field to given value.

### HasAlertPolicyId

`func (o *MonitoringAlertReason) HasAlertPolicyId() bool`

HasAlertPolicyId returns a boolean if a field has been set.

### GetMatchedRequirements

`func (o *MonitoringAlertReason) GetMatchedRequirements() []MonitoringMatchedRequirement`

GetMatchedRequirements returns the MatchedRequirements field if non-nil, zero value otherwise.

### GetMatchedRequirementsOk

`func (o *MonitoringAlertReason) GetMatchedRequirementsOk() (*[]MonitoringMatchedRequirement, bool)`

GetMatchedRequirementsOk returns a tuple with the MatchedRequirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedRequirements

`func (o *MonitoringAlertReason) SetMatchedRequirements(v []MonitoringMatchedRequirement)`

SetMatchedRequirements sets MatchedRequirements field to given value.

### HasMatchedRequirements

`func (o *MonitoringAlertReason) HasMatchedRequirements() bool`

HasMatchedRequirements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


