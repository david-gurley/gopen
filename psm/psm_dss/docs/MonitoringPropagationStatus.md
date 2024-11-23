# MonitoringPropagationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DscStatus** | Pointer to [**[]MonitoringDSCStatus**](MonitoringDSCStatus.md) | list of DSCs status. | [optional] 
**GenerationId** | Pointer to **string** | The Generation ID this status is for. | [optional] 
**MinVersion** | Pointer to **string** | The Version running on the slowest DSC. | [optional] 
**Pending** | Pointer to **int32** | Number of Naples pending. If this is 0 it can be assumed that everything is up to date. | [optional] 
**PendingDscs** | Pointer to **[]string** | list of DSCs where propagation did not complete. | [optional] 
**Status** | Pointer to **string** | Textual description of propagation status. | [optional] 
**Updated** | Pointer to **int32** | The number of Naples that this version has already been pushed to. | [optional] 

## Methods

### NewMonitoringPropagationStatus

`func NewMonitoringPropagationStatus() *MonitoringPropagationStatus`

NewMonitoringPropagationStatus instantiates a new MonitoringPropagationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringPropagationStatusWithDefaults

`func NewMonitoringPropagationStatusWithDefaults() *MonitoringPropagationStatus`

NewMonitoringPropagationStatusWithDefaults instantiates a new MonitoringPropagationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDscStatus

`func (o *MonitoringPropagationStatus) GetDscStatus() []MonitoringDSCStatus`

GetDscStatus returns the DscStatus field if non-nil, zero value otherwise.

### GetDscStatusOk

`func (o *MonitoringPropagationStatus) GetDscStatusOk() (*[]MonitoringDSCStatus, bool)`

GetDscStatusOk returns a tuple with the DscStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDscStatus

`func (o *MonitoringPropagationStatus) SetDscStatus(v []MonitoringDSCStatus)`

SetDscStatus sets DscStatus field to given value.

### HasDscStatus

`func (o *MonitoringPropagationStatus) HasDscStatus() bool`

HasDscStatus returns a boolean if a field has been set.

### GetGenerationId

`func (o *MonitoringPropagationStatus) GetGenerationId() string`

GetGenerationId returns the GenerationId field if non-nil, zero value otherwise.

### GetGenerationIdOk

`func (o *MonitoringPropagationStatus) GetGenerationIdOk() (*string, bool)`

GetGenerationIdOk returns a tuple with the GenerationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerationId

`func (o *MonitoringPropagationStatus) SetGenerationId(v string)`

SetGenerationId sets GenerationId field to given value.

### HasGenerationId

`func (o *MonitoringPropagationStatus) HasGenerationId() bool`

HasGenerationId returns a boolean if a field has been set.

### GetMinVersion

`func (o *MonitoringPropagationStatus) GetMinVersion() string`

GetMinVersion returns the MinVersion field if non-nil, zero value otherwise.

### GetMinVersionOk

`func (o *MonitoringPropagationStatus) GetMinVersionOk() (*string, bool)`

GetMinVersionOk returns a tuple with the MinVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinVersion

`func (o *MonitoringPropagationStatus) SetMinVersion(v string)`

SetMinVersion sets MinVersion field to given value.

### HasMinVersion

`func (o *MonitoringPropagationStatus) HasMinVersion() bool`

HasMinVersion returns a boolean if a field has been set.

### GetPending

`func (o *MonitoringPropagationStatus) GetPending() int32`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *MonitoringPropagationStatus) GetPendingOk() (*int32, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *MonitoringPropagationStatus) SetPending(v int32)`

SetPending sets Pending field to given value.

### HasPending

`func (o *MonitoringPropagationStatus) HasPending() bool`

HasPending returns a boolean if a field has been set.

### GetPendingDscs

`func (o *MonitoringPropagationStatus) GetPendingDscs() []string`

GetPendingDscs returns the PendingDscs field if non-nil, zero value otherwise.

### GetPendingDscsOk

`func (o *MonitoringPropagationStatus) GetPendingDscsOk() (*[]string, bool)`

GetPendingDscsOk returns a tuple with the PendingDscs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingDscs

`func (o *MonitoringPropagationStatus) SetPendingDscs(v []string)`

SetPendingDscs sets PendingDscs field to given value.

### HasPendingDscs

`func (o *MonitoringPropagationStatus) HasPendingDscs() bool`

HasPendingDscs returns a boolean if a field has been set.

### GetStatus

`func (o *MonitoringPropagationStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MonitoringPropagationStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MonitoringPropagationStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MonitoringPropagationStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdated

`func (o *MonitoringPropagationStatus) GetUpdated() int32`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *MonitoringPropagationStatus) GetUpdatedOk() (*int32, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *MonitoringPropagationStatus) SetUpdated(v int32)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *MonitoringPropagationStatus) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


