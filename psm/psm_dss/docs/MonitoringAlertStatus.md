# MonitoringAlertStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acknowledged** | Pointer to [**MonitoringAuditInfo**](monitoringAuditInfo.md) |  | [optional] 
**EventUri** | Pointer to **string** | Event that triggered the alert. | [optional] 
**Message** | Pointer to **string** | Message from the alert rule that triggered the alert. | [optional] 
**ObjectRef** | Pointer to [**ApiObjectRef**](apiObjectRef.md) |  | [optional] 
**Reason** | Pointer to [**MonitoringAlertReason**](monitoringAlertReason.md) |  | [optional] 
**Resolved** | Pointer to [**MonitoringAuditInfo**](monitoringAuditInfo.md) |  | [optional] 
**Severity** | Pointer to **string** | Severity of an alert. | [optional] [default to "info"]
**Source** | Pointer to [**MonitoringAlertSource**](monitoringAlertSource.md) |  | [optional] 
**TotalHits** | Pointer to **int32** | TotalHits on this alert, If there is an exisiting alert for the condition, we do not re-create the alert instead we update this counter. | [optional] 

## Methods

### NewMonitoringAlertStatus

`func NewMonitoringAlertStatus() *MonitoringAlertStatus`

NewMonitoringAlertStatus instantiates a new MonitoringAlertStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringAlertStatusWithDefaults

`func NewMonitoringAlertStatusWithDefaults() *MonitoringAlertStatus`

NewMonitoringAlertStatusWithDefaults instantiates a new MonitoringAlertStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcknowledged

`func (o *MonitoringAlertStatus) GetAcknowledged() MonitoringAuditInfo`

GetAcknowledged returns the Acknowledged field if non-nil, zero value otherwise.

### GetAcknowledgedOk

`func (o *MonitoringAlertStatus) GetAcknowledgedOk() (*MonitoringAuditInfo, bool)`

GetAcknowledgedOk returns a tuple with the Acknowledged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledged

`func (o *MonitoringAlertStatus) SetAcknowledged(v MonitoringAuditInfo)`

SetAcknowledged sets Acknowledged field to given value.

### HasAcknowledged

`func (o *MonitoringAlertStatus) HasAcknowledged() bool`

HasAcknowledged returns a boolean if a field has been set.

### GetEventUri

`func (o *MonitoringAlertStatus) GetEventUri() string`

GetEventUri returns the EventUri field if non-nil, zero value otherwise.

### GetEventUriOk

`func (o *MonitoringAlertStatus) GetEventUriOk() (*string, bool)`

GetEventUriOk returns a tuple with the EventUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventUri

`func (o *MonitoringAlertStatus) SetEventUri(v string)`

SetEventUri sets EventUri field to given value.

### HasEventUri

`func (o *MonitoringAlertStatus) HasEventUri() bool`

HasEventUri returns a boolean if a field has been set.

### GetMessage

`func (o *MonitoringAlertStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MonitoringAlertStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MonitoringAlertStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *MonitoringAlertStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetObjectRef

`func (o *MonitoringAlertStatus) GetObjectRef() ApiObjectRef`

GetObjectRef returns the ObjectRef field if non-nil, zero value otherwise.

### GetObjectRefOk

`func (o *MonitoringAlertStatus) GetObjectRefOk() (*ApiObjectRef, bool)`

GetObjectRefOk returns a tuple with the ObjectRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectRef

`func (o *MonitoringAlertStatus) SetObjectRef(v ApiObjectRef)`

SetObjectRef sets ObjectRef field to given value.

### HasObjectRef

`func (o *MonitoringAlertStatus) HasObjectRef() bool`

HasObjectRef returns a boolean if a field has been set.

### GetReason

`func (o *MonitoringAlertStatus) GetReason() MonitoringAlertReason`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *MonitoringAlertStatus) GetReasonOk() (*MonitoringAlertReason, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *MonitoringAlertStatus) SetReason(v MonitoringAlertReason)`

SetReason sets Reason field to given value.

### HasReason

`func (o *MonitoringAlertStatus) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetResolved

`func (o *MonitoringAlertStatus) GetResolved() MonitoringAuditInfo`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *MonitoringAlertStatus) GetResolvedOk() (*MonitoringAuditInfo, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *MonitoringAlertStatus) SetResolved(v MonitoringAuditInfo)`

SetResolved sets Resolved field to given value.

### HasResolved

`func (o *MonitoringAlertStatus) HasResolved() bool`

HasResolved returns a boolean if a field has been set.

### GetSeverity

`func (o *MonitoringAlertStatus) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *MonitoringAlertStatus) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *MonitoringAlertStatus) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *MonitoringAlertStatus) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSource

`func (o *MonitoringAlertStatus) GetSource() MonitoringAlertSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *MonitoringAlertStatus) GetSourceOk() (*MonitoringAlertSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *MonitoringAlertStatus) SetSource(v MonitoringAlertSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *MonitoringAlertStatus) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTotalHits

`func (o *MonitoringAlertStatus) GetTotalHits() int32`

GetTotalHits returns the TotalHits field if non-nil, zero value otherwise.

### GetTotalHitsOk

`func (o *MonitoringAlertStatus) GetTotalHitsOk() (*int32, bool)`

GetTotalHitsOk returns a tuple with the TotalHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalHits

`func (o *MonitoringAlertStatus) SetTotalHits(v int32)`

SetTotalHits sets TotalHits field to given value.

### HasTotalHits

`func (o *MonitoringAlertStatus) HasTotalHits() bool`

HasTotalHits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


