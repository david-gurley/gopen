# MonitoringMatchRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppProtocolSelectors** | Pointer to [**MonitoringAppProtoSelector**](monitoringAppProtoSelector.md) |  | [optional] 
**Destination** | Pointer to [**MonitoringMatchSelector**](monitoringMatchSelector.md) |  | [optional] 
**Source** | Pointer to [**MonitoringMatchSelector**](monitoringMatchSelector.md) |  | [optional] 

## Methods

### NewMonitoringMatchRule

`func NewMonitoringMatchRule() *MonitoringMatchRule`

NewMonitoringMatchRule instantiates a new MonitoringMatchRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringMatchRuleWithDefaults

`func NewMonitoringMatchRuleWithDefaults() *MonitoringMatchRule`

NewMonitoringMatchRuleWithDefaults instantiates a new MonitoringMatchRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppProtocolSelectors

`func (o *MonitoringMatchRule) GetAppProtocolSelectors() MonitoringAppProtoSelector`

GetAppProtocolSelectors returns the AppProtocolSelectors field if non-nil, zero value otherwise.

### GetAppProtocolSelectorsOk

`func (o *MonitoringMatchRule) GetAppProtocolSelectorsOk() (*MonitoringAppProtoSelector, bool)`

GetAppProtocolSelectorsOk returns a tuple with the AppProtocolSelectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppProtocolSelectors

`func (o *MonitoringMatchRule) SetAppProtocolSelectors(v MonitoringAppProtoSelector)`

SetAppProtocolSelectors sets AppProtocolSelectors field to given value.

### HasAppProtocolSelectors

`func (o *MonitoringMatchRule) HasAppProtocolSelectors() bool`

HasAppProtocolSelectors returns a boolean if a field has been set.

### GetDestination

`func (o *MonitoringMatchRule) GetDestination() MonitoringMatchSelector`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *MonitoringMatchRule) GetDestinationOk() (*MonitoringMatchSelector, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *MonitoringMatchRule) SetDestination(v MonitoringMatchSelector)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *MonitoringMatchRule) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetSource

`func (o *MonitoringMatchRule) GetSource() MonitoringMatchSelector`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *MonitoringMatchRule) GetSourceOk() (*MonitoringMatchSelector, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *MonitoringMatchRule) SetSource(v MonitoringMatchSelector)`

SetSource sets Source field to given value.

### HasSource

`func (o *MonitoringMatchRule) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


