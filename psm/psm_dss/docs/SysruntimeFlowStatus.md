# SysruntimeFlowStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionTracking** | Pointer to [**SysruntimeConnTrackInfo**](sysruntimeConnTrackInfo.md) |  | [optional] 
**FlowDirection** | Pointer to **string** |  | [optional] 
**FlowInstance** | Pointer to **string** |  | [optional] 

## Methods

### NewSysruntimeFlowStatus

`func NewSysruntimeFlowStatus() *SysruntimeFlowStatus`

NewSysruntimeFlowStatus instantiates a new SysruntimeFlowStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSysruntimeFlowStatusWithDefaults

`func NewSysruntimeFlowStatusWithDefaults() *SysruntimeFlowStatus`

NewSysruntimeFlowStatusWithDefaults instantiates a new SysruntimeFlowStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionTracking

`func (o *SysruntimeFlowStatus) GetConnectionTracking() SysruntimeConnTrackInfo`

GetConnectionTracking returns the ConnectionTracking field if non-nil, zero value otherwise.

### GetConnectionTrackingOk

`func (o *SysruntimeFlowStatus) GetConnectionTrackingOk() (*SysruntimeConnTrackInfo, bool)`

GetConnectionTrackingOk returns a tuple with the ConnectionTracking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTracking

`func (o *SysruntimeFlowStatus) SetConnectionTracking(v SysruntimeConnTrackInfo)`

SetConnectionTracking sets ConnectionTracking field to given value.

### HasConnectionTracking

`func (o *SysruntimeFlowStatus) HasConnectionTracking() bool`

HasConnectionTracking returns a boolean if a field has been set.

### GetFlowDirection

`func (o *SysruntimeFlowStatus) GetFlowDirection() string`

GetFlowDirection returns the FlowDirection field if non-nil, zero value otherwise.

### GetFlowDirectionOk

`func (o *SysruntimeFlowStatus) GetFlowDirectionOk() (*string, bool)`

GetFlowDirectionOk returns a tuple with the FlowDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowDirection

`func (o *SysruntimeFlowStatus) SetFlowDirection(v string)`

SetFlowDirection sets FlowDirection field to given value.

### HasFlowDirection

`func (o *SysruntimeFlowStatus) HasFlowDirection() bool`

HasFlowDirection returns a boolean if a field has been set.

### GetFlowInstance

`func (o *SysruntimeFlowStatus) GetFlowInstance() string`

GetFlowInstance returns the FlowInstance field if non-nil, zero value otherwise.

### GetFlowInstanceOk

`func (o *SysruntimeFlowStatus) GetFlowInstanceOk() (*string, bool)`

GetFlowInstanceOk returns a tuple with the FlowInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowInstance

`func (o *SysruntimeFlowStatus) SetFlowInstance(v string)`

SetFlowInstance sets FlowInstance field to given value.

### HasFlowInstance

`func (o *SysruntimeFlowStatus) HasFlowInstance() bool`

HasFlowInstance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


