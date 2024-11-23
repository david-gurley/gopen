# MonitoringMirrorSessionSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collectors** | Pointer to [**[]MonitoringMirrorCollector**](MonitoringMirrorCollector.md) | Mirrored packet collectors. | [optional] 
**Disabled** | Pointer to **bool** | Enable/disable mirroring. | [optional] 
**Interfaces** | Pointer to [**MonitoringInterfaceMirror**](monitoringInterfaceMirror.md) |  | [optional] 
**MatchRules** | Pointer to [**[]MonitoringMatchRule**](MonitoringMatchRule.md) | Traffic Selection Rules - Matching pakcets are mirrored, based on packet filters and start/stop conditions. | [optional] 
**PacketFilters** | Pointer to **[]string** |  | [optional] 
**PacketSize** | Pointer to **int64** | PacketSize: Max size of a mirrored packet, packet size is not checked by default. Value should be between 64 and 2048. | [optional] 
**Source** | Pointer to [**MonitoringMirrorSource**](monitoringMirrorSource.md) |  | [optional] 
**SpanId** | Pointer to **int64** | Value should be between 1 and 1023. | [optional] [default to 1]
**StartCondition** | Pointer to [**MonitoringMirrorStartConditions**](monitoringMirrorStartConditions.md) |  | [optional] 
**Workloads** | Pointer to [**MonitoringWorkloadMirror**](monitoringWorkloadMirror.md) |  | [optional] 

## Methods

### NewMonitoringMirrorSessionSpec

`func NewMonitoringMirrorSessionSpec() *MonitoringMirrorSessionSpec`

NewMonitoringMirrorSessionSpec instantiates a new MonitoringMirrorSessionSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringMirrorSessionSpecWithDefaults

`func NewMonitoringMirrorSessionSpecWithDefaults() *MonitoringMirrorSessionSpec`

NewMonitoringMirrorSessionSpecWithDefaults instantiates a new MonitoringMirrorSessionSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectors

`func (o *MonitoringMirrorSessionSpec) GetCollectors() []MonitoringMirrorCollector`

GetCollectors returns the Collectors field if non-nil, zero value otherwise.

### GetCollectorsOk

`func (o *MonitoringMirrorSessionSpec) GetCollectorsOk() (*[]MonitoringMirrorCollector, bool)`

GetCollectorsOk returns a tuple with the Collectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectors

`func (o *MonitoringMirrorSessionSpec) SetCollectors(v []MonitoringMirrorCollector)`

SetCollectors sets Collectors field to given value.

### HasCollectors

`func (o *MonitoringMirrorSessionSpec) HasCollectors() bool`

HasCollectors returns a boolean if a field has been set.

### GetDisabled

`func (o *MonitoringMirrorSessionSpec) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *MonitoringMirrorSessionSpec) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *MonitoringMirrorSessionSpec) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *MonitoringMirrorSessionSpec) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetInterfaces

`func (o *MonitoringMirrorSessionSpec) GetInterfaces() MonitoringInterfaceMirror`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *MonitoringMirrorSessionSpec) GetInterfacesOk() (*MonitoringInterfaceMirror, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *MonitoringMirrorSessionSpec) SetInterfaces(v MonitoringInterfaceMirror)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *MonitoringMirrorSessionSpec) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetMatchRules

`func (o *MonitoringMirrorSessionSpec) GetMatchRules() []MonitoringMatchRule`

GetMatchRules returns the MatchRules field if non-nil, zero value otherwise.

### GetMatchRulesOk

`func (o *MonitoringMirrorSessionSpec) GetMatchRulesOk() (*[]MonitoringMatchRule, bool)`

GetMatchRulesOk returns a tuple with the MatchRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchRules

`func (o *MonitoringMirrorSessionSpec) SetMatchRules(v []MonitoringMatchRule)`

SetMatchRules sets MatchRules field to given value.

### HasMatchRules

`func (o *MonitoringMirrorSessionSpec) HasMatchRules() bool`

HasMatchRules returns a boolean if a field has been set.

### GetPacketFilters

`func (o *MonitoringMirrorSessionSpec) GetPacketFilters() []string`

GetPacketFilters returns the PacketFilters field if non-nil, zero value otherwise.

### GetPacketFiltersOk

`func (o *MonitoringMirrorSessionSpec) GetPacketFiltersOk() (*[]string, bool)`

GetPacketFiltersOk returns a tuple with the PacketFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketFilters

`func (o *MonitoringMirrorSessionSpec) SetPacketFilters(v []string)`

SetPacketFilters sets PacketFilters field to given value.

### HasPacketFilters

`func (o *MonitoringMirrorSessionSpec) HasPacketFilters() bool`

HasPacketFilters returns a boolean if a field has been set.

### GetPacketSize

`func (o *MonitoringMirrorSessionSpec) GetPacketSize() int64`

GetPacketSize returns the PacketSize field if non-nil, zero value otherwise.

### GetPacketSizeOk

`func (o *MonitoringMirrorSessionSpec) GetPacketSizeOk() (*int64, bool)`

GetPacketSizeOk returns a tuple with the PacketSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketSize

`func (o *MonitoringMirrorSessionSpec) SetPacketSize(v int64)`

SetPacketSize sets PacketSize field to given value.

### HasPacketSize

`func (o *MonitoringMirrorSessionSpec) HasPacketSize() bool`

HasPacketSize returns a boolean if a field has been set.

### GetSource

`func (o *MonitoringMirrorSessionSpec) GetSource() MonitoringMirrorSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *MonitoringMirrorSessionSpec) GetSourceOk() (*MonitoringMirrorSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *MonitoringMirrorSessionSpec) SetSource(v MonitoringMirrorSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *MonitoringMirrorSessionSpec) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSpanId

`func (o *MonitoringMirrorSessionSpec) GetSpanId() int64`

GetSpanId returns the SpanId field if non-nil, zero value otherwise.

### GetSpanIdOk

`func (o *MonitoringMirrorSessionSpec) GetSpanIdOk() (*int64, bool)`

GetSpanIdOk returns a tuple with the SpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanId

`func (o *MonitoringMirrorSessionSpec) SetSpanId(v int64)`

SetSpanId sets SpanId field to given value.

### HasSpanId

`func (o *MonitoringMirrorSessionSpec) HasSpanId() bool`

HasSpanId returns a boolean if a field has been set.

### GetStartCondition

`func (o *MonitoringMirrorSessionSpec) GetStartCondition() MonitoringMirrorStartConditions`

GetStartCondition returns the StartCondition field if non-nil, zero value otherwise.

### GetStartConditionOk

`func (o *MonitoringMirrorSessionSpec) GetStartConditionOk() (*MonitoringMirrorStartConditions, bool)`

GetStartConditionOk returns a tuple with the StartCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartCondition

`func (o *MonitoringMirrorSessionSpec) SetStartCondition(v MonitoringMirrorStartConditions)`

SetStartCondition sets StartCondition field to given value.

### HasStartCondition

`func (o *MonitoringMirrorSessionSpec) HasStartCondition() bool`

HasStartCondition returns a boolean if a field has been set.

### GetWorkloads

`func (o *MonitoringMirrorSessionSpec) GetWorkloads() MonitoringWorkloadMirror`

GetWorkloads returns the Workloads field if non-nil, zero value otherwise.

### GetWorkloadsOk

`func (o *MonitoringMirrorSessionSpec) GetWorkloadsOk() (*MonitoringWorkloadMirror, bool)`

GetWorkloadsOk returns a tuple with the Workloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloads

`func (o *MonitoringMirrorSessionSpec) SetWorkloads(v MonitoringWorkloadMirror)`

SetWorkloads sets Workloads field to given value.

### HasWorkloads

`func (o *MonitoringMirrorSessionSpec) HasWorkloads() bool`

HasWorkloads returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


