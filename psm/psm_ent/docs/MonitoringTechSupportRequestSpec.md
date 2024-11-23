# MonitoringTechSupportRequestSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionSelector** | Pointer to [**LabelsSelector**](labelsSelector.md) |  | [optional] 
**NodeSelector** | Pointer to [**TechSupportRequestSpecNodeSelectorSpec**](TechSupportRequestSpecNodeSelectorSpec.md) |  | [optional] 
**SkipCores** | Pointer to **bool** | SkipCores if set to true skips the core files when collecting techsupport. | [optional] [default to false]
**Verbosity** | Pointer to **int32** | Verbosity defines the verbosity level. | [optional] 

## Methods

### NewMonitoringTechSupportRequestSpec

`func NewMonitoringTechSupportRequestSpec() *MonitoringTechSupportRequestSpec`

NewMonitoringTechSupportRequestSpec instantiates a new MonitoringTechSupportRequestSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringTechSupportRequestSpecWithDefaults

`func NewMonitoringTechSupportRequestSpecWithDefaults() *MonitoringTechSupportRequestSpec`

NewMonitoringTechSupportRequestSpecWithDefaults instantiates a new MonitoringTechSupportRequestSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionSelector

`func (o *MonitoringTechSupportRequestSpec) GetCollectionSelector() LabelsSelector`

GetCollectionSelector returns the CollectionSelector field if non-nil, zero value otherwise.

### GetCollectionSelectorOk

`func (o *MonitoringTechSupportRequestSpec) GetCollectionSelectorOk() (*LabelsSelector, bool)`

GetCollectionSelectorOk returns a tuple with the CollectionSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionSelector

`func (o *MonitoringTechSupportRequestSpec) SetCollectionSelector(v LabelsSelector)`

SetCollectionSelector sets CollectionSelector field to given value.

### HasCollectionSelector

`func (o *MonitoringTechSupportRequestSpec) HasCollectionSelector() bool`

HasCollectionSelector returns a boolean if a field has been set.

### GetNodeSelector

`func (o *MonitoringTechSupportRequestSpec) GetNodeSelector() TechSupportRequestSpecNodeSelectorSpec`

GetNodeSelector returns the NodeSelector field if non-nil, zero value otherwise.

### GetNodeSelectorOk

`func (o *MonitoringTechSupportRequestSpec) GetNodeSelectorOk() (*TechSupportRequestSpecNodeSelectorSpec, bool)`

GetNodeSelectorOk returns a tuple with the NodeSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSelector

`func (o *MonitoringTechSupportRequestSpec) SetNodeSelector(v TechSupportRequestSpecNodeSelectorSpec)`

SetNodeSelector sets NodeSelector field to given value.

### HasNodeSelector

`func (o *MonitoringTechSupportRequestSpec) HasNodeSelector() bool`

HasNodeSelector returns a boolean if a field has been set.

### GetSkipCores

`func (o *MonitoringTechSupportRequestSpec) GetSkipCores() bool`

GetSkipCores returns the SkipCores field if non-nil, zero value otherwise.

### GetSkipCoresOk

`func (o *MonitoringTechSupportRequestSpec) GetSkipCoresOk() (*bool, bool)`

GetSkipCoresOk returns a tuple with the SkipCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipCores

`func (o *MonitoringTechSupportRequestSpec) SetSkipCores(v bool)`

SetSkipCores sets SkipCores field to given value.

### HasSkipCores

`func (o *MonitoringTechSupportRequestSpec) HasSkipCores() bool`

HasSkipCores returns a boolean if a field has been set.

### GetVerbosity

`func (o *MonitoringTechSupportRequestSpec) GetVerbosity() int32`

GetVerbosity returns the Verbosity field if non-nil, zero value otherwise.

### GetVerbosityOk

`func (o *MonitoringTechSupportRequestSpec) GetVerbosityOk() (*int32, bool)`

GetVerbosityOk returns a tuple with the Verbosity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerbosity

`func (o *MonitoringTechSupportRequestSpec) SetVerbosity(v int32)`

SetVerbosity sets Verbosity field to given value.

### HasVerbosity

`func (o *MonitoringTechSupportRequestSpec) HasVerbosity() bool`

HasVerbosity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


