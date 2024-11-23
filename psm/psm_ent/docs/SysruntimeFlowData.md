# SysruntimeFlowData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionTracking** | Pointer to [**SysruntimeConnTrackInfo**](sysruntimeConnTrackInfo.md) |  | [optional] 
**FlowInfo** | Pointer to [**SysruntimeFlowInfo**](sysruntimeFlowInfo.md) |  | [optional] 
**TelemetryInfo** | Pointer to [**SysruntimeTelemetryInfo**](sysruntimeTelemetryInfo.md) |  | [optional] 

## Methods

### NewSysruntimeFlowData

`func NewSysruntimeFlowData() *SysruntimeFlowData`

NewSysruntimeFlowData instantiates a new SysruntimeFlowData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSysruntimeFlowDataWithDefaults

`func NewSysruntimeFlowDataWithDefaults() *SysruntimeFlowData`

NewSysruntimeFlowDataWithDefaults instantiates a new SysruntimeFlowData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionTracking

`func (o *SysruntimeFlowData) GetConnectionTracking() SysruntimeConnTrackInfo`

GetConnectionTracking returns the ConnectionTracking field if non-nil, zero value otherwise.

### GetConnectionTrackingOk

`func (o *SysruntimeFlowData) GetConnectionTrackingOk() (*SysruntimeConnTrackInfo, bool)`

GetConnectionTrackingOk returns a tuple with the ConnectionTracking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTracking

`func (o *SysruntimeFlowData) SetConnectionTracking(v SysruntimeConnTrackInfo)`

SetConnectionTracking sets ConnectionTracking field to given value.

### HasConnectionTracking

`func (o *SysruntimeFlowData) HasConnectionTracking() bool`

HasConnectionTracking returns a boolean if a field has been set.

### GetFlowInfo

`func (o *SysruntimeFlowData) GetFlowInfo() SysruntimeFlowInfo`

GetFlowInfo returns the FlowInfo field if non-nil, zero value otherwise.

### GetFlowInfoOk

`func (o *SysruntimeFlowData) GetFlowInfoOk() (*SysruntimeFlowInfo, bool)`

GetFlowInfoOk returns a tuple with the FlowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowInfo

`func (o *SysruntimeFlowData) SetFlowInfo(v SysruntimeFlowInfo)`

SetFlowInfo sets FlowInfo field to given value.

### HasFlowInfo

`func (o *SysruntimeFlowData) HasFlowInfo() bool`

HasFlowInfo returns a boolean if a field has been set.

### GetTelemetryInfo

`func (o *SysruntimeFlowData) GetTelemetryInfo() SysruntimeTelemetryInfo`

GetTelemetryInfo returns the TelemetryInfo field if non-nil, zero value otherwise.

### GetTelemetryInfoOk

`func (o *SysruntimeFlowData) GetTelemetryInfoOk() (*SysruntimeTelemetryInfo, bool)`

GetTelemetryInfoOk returns a tuple with the TelemetryInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelemetryInfo

`func (o *SysruntimeFlowData) SetTelemetryInfo(v SysruntimeTelemetryInfo)`

SetTelemetryInfo sets TelemetryInfo field to given value.

### HasTelemetryInfo

`func (o *SysruntimeFlowData) HasTelemetryInfo() bool`

HasTelemetryInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


