# MonitoringAuditInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | Pointer to **time.Time** | Time at which the action was performed. | [optional] 
**User** | Pointer to **string** | Name of the user performed some action. | [optional] 

## Methods

### NewMonitoringAuditInfo

`func NewMonitoringAuditInfo() *MonitoringAuditInfo`

NewMonitoringAuditInfo instantiates a new MonitoringAuditInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringAuditInfoWithDefaults

`func NewMonitoringAuditInfoWithDefaults() *MonitoringAuditInfo`

NewMonitoringAuditInfoWithDefaults instantiates a new MonitoringAuditInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *MonitoringAuditInfo) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *MonitoringAuditInfo) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *MonitoringAuditInfo) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *MonitoringAuditInfo) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetUser

`func (o *MonitoringAuditInfo) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *MonitoringAuditInfo) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *MonitoringAuditInfo) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *MonitoringAuditInfo) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


