# MonitoringArchiveQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTime** | Pointer to **time.Time** | EndTime selects all logs with timestamp less than the EndTime, example 2018-09-18T00:12:00Z. | [optional] 
**Fields** | Pointer to [**FieldsSelector**](fieldsSelector.md) |  | [optional] 
**Labels** | Pointer to [**LabelsSelector**](labelsSelector.md) |  | [optional] 
**StartTime** | Pointer to **time.Time** | StartTime selects all logs with timestamp greater than the StartTime, example 2018-10-18T00:12:00Z. | [optional] 
**Tenants** | Pointer to **[]string** | OR of tenants within the scope of which archive needs to be performed. If not specified, it will be set to tenant of the logged in user. Also users in non default tenant can archive logs in their tenant scope only. | [optional] 
**Texts** | Pointer to [**[]SearchTextRequirement**](SearchTextRequirement.md) | OR of Text-requirements to be matched, Exclude is not supported for Text search. | [optional] 

## Methods

### NewMonitoringArchiveQuery

`func NewMonitoringArchiveQuery() *MonitoringArchiveQuery`

NewMonitoringArchiveQuery instantiates a new MonitoringArchiveQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringArchiveQueryWithDefaults

`func NewMonitoringArchiveQueryWithDefaults() *MonitoringArchiveQuery`

NewMonitoringArchiveQueryWithDefaults instantiates a new MonitoringArchiveQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTime

`func (o *MonitoringArchiveQuery) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *MonitoringArchiveQuery) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *MonitoringArchiveQuery) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *MonitoringArchiveQuery) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetFields

`func (o *MonitoringArchiveQuery) GetFields() FieldsSelector`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *MonitoringArchiveQuery) GetFieldsOk() (*FieldsSelector, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *MonitoringArchiveQuery) SetFields(v FieldsSelector)`

SetFields sets Fields field to given value.

### HasFields

`func (o *MonitoringArchiveQuery) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetLabels

`func (o *MonitoringArchiveQuery) GetLabels() LabelsSelector`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *MonitoringArchiveQuery) GetLabelsOk() (*LabelsSelector, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *MonitoringArchiveQuery) SetLabels(v LabelsSelector)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *MonitoringArchiveQuery) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetStartTime

`func (o *MonitoringArchiveQuery) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *MonitoringArchiveQuery) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *MonitoringArchiveQuery) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *MonitoringArchiveQuery) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetTenants

`func (o *MonitoringArchiveQuery) GetTenants() []string`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *MonitoringArchiveQuery) GetTenantsOk() (*[]string, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *MonitoringArchiveQuery) SetTenants(v []string)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *MonitoringArchiveQuery) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetTexts

`func (o *MonitoringArchiveQuery) GetTexts() []SearchTextRequirement`

GetTexts returns the Texts field if non-nil, zero value otherwise.

### GetTextsOk

`func (o *MonitoringArchiveQuery) GetTextsOk() (*[]SearchTextRequirement, bool)`

GetTextsOk returns a tuple with the Texts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTexts

`func (o *MonitoringArchiveQuery) SetTexts(v []SearchTextRequirement)`

SetTexts sets Texts field to given value.

### HasTexts

`func (o *MonitoringArchiveQuery) HasTexts() bool`

HasTexts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


