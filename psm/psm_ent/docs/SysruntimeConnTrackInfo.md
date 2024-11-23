# SysruntimeConnTrackInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateTimestamp** | Pointer to **time.Time** |  | [optional] 
**FlowBytes** | Pointer to **string** |  | [optional] 
**FlowPackets** | Pointer to **int64** |  | [optional] 

## Methods

### NewSysruntimeConnTrackInfo

`func NewSysruntimeConnTrackInfo() *SysruntimeConnTrackInfo`

NewSysruntimeConnTrackInfo instantiates a new SysruntimeConnTrackInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSysruntimeConnTrackInfoWithDefaults

`func NewSysruntimeConnTrackInfoWithDefaults() *SysruntimeConnTrackInfo`

NewSysruntimeConnTrackInfoWithDefaults instantiates a new SysruntimeConnTrackInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateTimestamp

`func (o *SysruntimeConnTrackInfo) GetCreateTimestamp() time.Time`

GetCreateTimestamp returns the CreateTimestamp field if non-nil, zero value otherwise.

### GetCreateTimestampOk

`func (o *SysruntimeConnTrackInfo) GetCreateTimestampOk() (*time.Time, bool)`

GetCreateTimestampOk returns a tuple with the CreateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimestamp

`func (o *SysruntimeConnTrackInfo) SetCreateTimestamp(v time.Time)`

SetCreateTimestamp sets CreateTimestamp field to given value.

### HasCreateTimestamp

`func (o *SysruntimeConnTrackInfo) HasCreateTimestamp() bool`

HasCreateTimestamp returns a boolean if a field has been set.

### GetFlowBytes

`func (o *SysruntimeConnTrackInfo) GetFlowBytes() string`

GetFlowBytes returns the FlowBytes field if non-nil, zero value otherwise.

### GetFlowBytesOk

`func (o *SysruntimeConnTrackInfo) GetFlowBytesOk() (*string, bool)`

GetFlowBytesOk returns a tuple with the FlowBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowBytes

`func (o *SysruntimeConnTrackInfo) SetFlowBytes(v string)`

SetFlowBytes sets FlowBytes field to given value.

### HasFlowBytes

`func (o *SysruntimeConnTrackInfo) HasFlowBytes() bool`

HasFlowBytes returns a boolean if a field has been set.

### GetFlowPackets

`func (o *SysruntimeConnTrackInfo) GetFlowPackets() int64`

GetFlowPackets returns the FlowPackets field if non-nil, zero value otherwise.

### GetFlowPacketsOk

`func (o *SysruntimeConnTrackInfo) GetFlowPacketsOk() (*int64, bool)`

GetFlowPacketsOk returns a tuple with the FlowPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowPackets

`func (o *SysruntimeConnTrackInfo) SetFlowPackets(v int64)`

SetFlowPackets sets FlowPackets field to given value.

### HasFlowPackets

`func (o *SysruntimeConnTrackInfo) HasFlowPackets() bool`

HasFlowPackets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


