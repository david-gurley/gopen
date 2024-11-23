# MonitoringTechSupportRequestStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CtrlrNodeResults** | Pointer to [**map[string]MonitoringTechSupportNodeResult**](monitoringTechSupportNodeResult.md) |  | [optional] 
**DscResults** | Pointer to [**map[string]MonitoringTechSupportNodeResult**](monitoringTechSupportNodeResult.md) |  | [optional] 
**InstanceId** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] [default to "scheduled"]

## Methods

### NewMonitoringTechSupportRequestStatus

`func NewMonitoringTechSupportRequestStatus() *MonitoringTechSupportRequestStatus`

NewMonitoringTechSupportRequestStatus instantiates a new MonitoringTechSupportRequestStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringTechSupportRequestStatusWithDefaults

`func NewMonitoringTechSupportRequestStatusWithDefaults() *MonitoringTechSupportRequestStatus`

NewMonitoringTechSupportRequestStatusWithDefaults instantiates a new MonitoringTechSupportRequestStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCtrlrNodeResults

`func (o *MonitoringTechSupportRequestStatus) GetCtrlrNodeResults() map[string]MonitoringTechSupportNodeResult`

GetCtrlrNodeResults returns the CtrlrNodeResults field if non-nil, zero value otherwise.

### GetCtrlrNodeResultsOk

`func (o *MonitoringTechSupportRequestStatus) GetCtrlrNodeResultsOk() (*map[string]MonitoringTechSupportNodeResult, bool)`

GetCtrlrNodeResultsOk returns a tuple with the CtrlrNodeResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtrlrNodeResults

`func (o *MonitoringTechSupportRequestStatus) SetCtrlrNodeResults(v map[string]MonitoringTechSupportNodeResult)`

SetCtrlrNodeResults sets CtrlrNodeResults field to given value.

### HasCtrlrNodeResults

`func (o *MonitoringTechSupportRequestStatus) HasCtrlrNodeResults() bool`

HasCtrlrNodeResults returns a boolean if a field has been set.

### GetDscResults

`func (o *MonitoringTechSupportRequestStatus) GetDscResults() map[string]MonitoringTechSupportNodeResult`

GetDscResults returns the DscResults field if non-nil, zero value otherwise.

### GetDscResultsOk

`func (o *MonitoringTechSupportRequestStatus) GetDscResultsOk() (*map[string]MonitoringTechSupportNodeResult, bool)`

GetDscResultsOk returns a tuple with the DscResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDscResults

`func (o *MonitoringTechSupportRequestStatus) SetDscResults(v map[string]MonitoringTechSupportNodeResult)`

SetDscResults sets DscResults field to given value.

### HasDscResults

`func (o *MonitoringTechSupportRequestStatus) HasDscResults() bool`

HasDscResults returns a boolean if a field has been set.

### GetInstanceId

`func (o *MonitoringTechSupportRequestStatus) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *MonitoringTechSupportRequestStatus) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *MonitoringTechSupportRequestStatus) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *MonitoringTechSupportRequestStatus) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetReason

`func (o *MonitoringTechSupportRequestStatus) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *MonitoringTechSupportRequestStatus) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *MonitoringTechSupportRequestStatus) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *MonitoringTechSupportRequestStatus) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatus

`func (o *MonitoringTechSupportRequestStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MonitoringTechSupportRequestStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MonitoringTechSupportRequestStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MonitoringTechSupportRequestStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


