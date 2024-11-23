# MonitoringAlertPolicySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destinations** | Pointer to **[]string** | name of the alert destinations to be used to send out notification when an alert gets generated. | [optional] 
**Enable** | Pointer to **bool** | User can disable the policy by setting this field. Disabled policies will not generate any more alerts but the outstanding ones will remain as is. | [optional] [default to true]
**Message** | Pointer to **string** | Message to be used while generating the alert. | [optional] 
**Requirements** | Pointer to [**[]FieldsRequirement**](FieldsRequirement.md) | List of requirements that needs to be met to trigger an alert. | [optional] 
**Resource** | Pointer to **string** | Resource type - target resource to run this policy. e.g. Network, Endpoint - object based alert policy Event - event based alert policy EndpointMetrics - metric based alert policy based on the resource type, the policy gets interpreted. | [optional] 
**Severity** | Pointer to **string** | Severity to be set for an alert that gets triggered from this rule. | [optional] [default to "info"]

## Methods

### NewMonitoringAlertPolicySpec

`func NewMonitoringAlertPolicySpec() *MonitoringAlertPolicySpec`

NewMonitoringAlertPolicySpec instantiates a new MonitoringAlertPolicySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringAlertPolicySpecWithDefaults

`func NewMonitoringAlertPolicySpecWithDefaults() *MonitoringAlertPolicySpec`

NewMonitoringAlertPolicySpecWithDefaults instantiates a new MonitoringAlertPolicySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinations

`func (o *MonitoringAlertPolicySpec) GetDestinations() []string`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *MonitoringAlertPolicySpec) GetDestinationsOk() (*[]string, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *MonitoringAlertPolicySpec) SetDestinations(v []string)`

SetDestinations sets Destinations field to given value.

### HasDestinations

`func (o *MonitoringAlertPolicySpec) HasDestinations() bool`

HasDestinations returns a boolean if a field has been set.

### GetEnable

`func (o *MonitoringAlertPolicySpec) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *MonitoringAlertPolicySpec) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *MonitoringAlertPolicySpec) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *MonitoringAlertPolicySpec) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetMessage

`func (o *MonitoringAlertPolicySpec) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MonitoringAlertPolicySpec) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MonitoringAlertPolicySpec) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *MonitoringAlertPolicySpec) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetRequirements

`func (o *MonitoringAlertPolicySpec) GetRequirements() []FieldsRequirement`

GetRequirements returns the Requirements field if non-nil, zero value otherwise.

### GetRequirementsOk

`func (o *MonitoringAlertPolicySpec) GetRequirementsOk() (*[]FieldsRequirement, bool)`

GetRequirementsOk returns a tuple with the Requirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirements

`func (o *MonitoringAlertPolicySpec) SetRequirements(v []FieldsRequirement)`

SetRequirements sets Requirements field to given value.

### HasRequirements

`func (o *MonitoringAlertPolicySpec) HasRequirements() bool`

HasRequirements returns a boolean if a field has been set.

### GetResource

`func (o *MonitoringAlertPolicySpec) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *MonitoringAlertPolicySpec) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *MonitoringAlertPolicySpec) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *MonitoringAlertPolicySpec) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetSeverity

`func (o *MonitoringAlertPolicySpec) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *MonitoringAlertPolicySpec) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *MonitoringAlertPolicySpec) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *MonitoringAlertPolicySpec) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


