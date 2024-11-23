# MonitoringArchiveRequestStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] [default to "scheduled"]
**Uri** | Pointer to **string** |  | [optional] 

## Methods

### NewMonitoringArchiveRequestStatus

`func NewMonitoringArchiveRequestStatus() *MonitoringArchiveRequestStatus`

NewMonitoringArchiveRequestStatus instantiates a new MonitoringArchiveRequestStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringArchiveRequestStatusWithDefaults

`func NewMonitoringArchiveRequestStatusWithDefaults() *MonitoringArchiveRequestStatus`

NewMonitoringArchiveRequestStatusWithDefaults instantiates a new MonitoringArchiveRequestStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *MonitoringArchiveRequestStatus) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *MonitoringArchiveRequestStatus) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *MonitoringArchiveRequestStatus) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *MonitoringArchiveRequestStatus) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatus

`func (o *MonitoringArchiveRequestStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MonitoringArchiveRequestStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MonitoringArchiveRequestStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MonitoringArchiveRequestStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUri

`func (o *MonitoringArchiveRequestStatus) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *MonitoringArchiveRequestStatus) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *MonitoringArchiveRequestStatus) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *MonitoringArchiveRequestStatus) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


