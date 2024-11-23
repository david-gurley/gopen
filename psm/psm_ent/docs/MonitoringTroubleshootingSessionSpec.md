# MonitoringTroubleshootingSessionSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableMirroring** | Pointer to **bool** | If packet capture is enabled, a mirror-session will be internally created. | [optional] 
**FlowSelector** | Pointer to [**MonitoringMatchRule**](monitoringMatchRule.md) |  | [optional] 
**RepeatEvery** | Pointer to **string** |  | [optional] 
**TimeWindow** | Pointer to [**MonitoringTimeWindow**](monitoringTimeWindow.md) |  | [optional] 

## Methods

### NewMonitoringTroubleshootingSessionSpec

`func NewMonitoringTroubleshootingSessionSpec() *MonitoringTroubleshootingSessionSpec`

NewMonitoringTroubleshootingSessionSpec instantiates a new MonitoringTroubleshootingSessionSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringTroubleshootingSessionSpecWithDefaults

`func NewMonitoringTroubleshootingSessionSpecWithDefaults() *MonitoringTroubleshootingSessionSpec`

NewMonitoringTroubleshootingSessionSpecWithDefaults instantiates a new MonitoringTroubleshootingSessionSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableMirroring

`func (o *MonitoringTroubleshootingSessionSpec) GetEnableMirroring() bool`

GetEnableMirroring returns the EnableMirroring field if non-nil, zero value otherwise.

### GetEnableMirroringOk

`func (o *MonitoringTroubleshootingSessionSpec) GetEnableMirroringOk() (*bool, bool)`

GetEnableMirroringOk returns a tuple with the EnableMirroring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMirroring

`func (o *MonitoringTroubleshootingSessionSpec) SetEnableMirroring(v bool)`

SetEnableMirroring sets EnableMirroring field to given value.

### HasEnableMirroring

`func (o *MonitoringTroubleshootingSessionSpec) HasEnableMirroring() bool`

HasEnableMirroring returns a boolean if a field has been set.

### GetFlowSelector

`func (o *MonitoringTroubleshootingSessionSpec) GetFlowSelector() MonitoringMatchRule`

GetFlowSelector returns the FlowSelector field if non-nil, zero value otherwise.

### GetFlowSelectorOk

`func (o *MonitoringTroubleshootingSessionSpec) GetFlowSelectorOk() (*MonitoringMatchRule, bool)`

GetFlowSelectorOk returns a tuple with the FlowSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSelector

`func (o *MonitoringTroubleshootingSessionSpec) SetFlowSelector(v MonitoringMatchRule)`

SetFlowSelector sets FlowSelector field to given value.

### HasFlowSelector

`func (o *MonitoringTroubleshootingSessionSpec) HasFlowSelector() bool`

HasFlowSelector returns a boolean if a field has been set.

### GetRepeatEvery

`func (o *MonitoringTroubleshootingSessionSpec) GetRepeatEvery() string`

GetRepeatEvery returns the RepeatEvery field if non-nil, zero value otherwise.

### GetRepeatEveryOk

`func (o *MonitoringTroubleshootingSessionSpec) GetRepeatEveryOk() (*string, bool)`

GetRepeatEveryOk returns a tuple with the RepeatEvery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatEvery

`func (o *MonitoringTroubleshootingSessionSpec) SetRepeatEvery(v string)`

SetRepeatEvery sets RepeatEvery field to given value.

### HasRepeatEvery

`func (o *MonitoringTroubleshootingSessionSpec) HasRepeatEvery() bool`

HasRepeatEvery returns a boolean if a field has been set.

### GetTimeWindow

`func (o *MonitoringTroubleshootingSessionSpec) GetTimeWindow() MonitoringTimeWindow`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *MonitoringTroubleshootingSessionSpec) GetTimeWindowOk() (*MonitoringTimeWindow, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *MonitoringTroubleshootingSessionSpec) SetTimeWindow(v MonitoringTimeWindow)`

SetTimeWindow sets TimeWindow field to given value.

### HasTimeWindow

`func (o *MonitoringTroubleshootingSessionSpec) HasTimeWindow() bool`

HasTimeWindow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


