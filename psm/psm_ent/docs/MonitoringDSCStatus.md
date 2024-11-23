# MonitoringDSCStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DscId** | Pointer to **string** | DSC ID for which the agent error or warning is issued. | [optional] 
**DscInfoStatus** | Pointer to **string** | InfoStatus contains agent message the operation is failed or warning is issued. | [optional] 

## Methods

### NewMonitoringDSCStatus

`func NewMonitoringDSCStatus() *MonitoringDSCStatus`

NewMonitoringDSCStatus instantiates a new MonitoringDSCStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringDSCStatusWithDefaults

`func NewMonitoringDSCStatusWithDefaults() *MonitoringDSCStatus`

NewMonitoringDSCStatusWithDefaults instantiates a new MonitoringDSCStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDscId

`func (o *MonitoringDSCStatus) GetDscId() string`

GetDscId returns the DscId field if non-nil, zero value otherwise.

### GetDscIdOk

`func (o *MonitoringDSCStatus) GetDscIdOk() (*string, bool)`

GetDscIdOk returns a tuple with the DscId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDscId

`func (o *MonitoringDSCStatus) SetDscId(v string)`

SetDscId sets DscId field to given value.

### HasDscId

`func (o *MonitoringDSCStatus) HasDscId() bool`

HasDscId returns a boolean if a field has been set.

### GetDscInfoStatus

`func (o *MonitoringDSCStatus) GetDscInfoStatus() string`

GetDscInfoStatus returns the DscInfoStatus field if non-nil, zero value otherwise.

### GetDscInfoStatusOk

`func (o *MonitoringDSCStatus) GetDscInfoStatusOk() (*string, bool)`

GetDscInfoStatusOk returns a tuple with the DscInfoStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDscInfoStatus

`func (o *MonitoringDSCStatus) SetDscInfoStatus(v string)`

SetDscInfoStatus sets DscInfoStatus field to given value.

### HasDscInfoStatus

`func (o *MonitoringDSCStatus) HasDscInfoStatus() bool`

HasDscInfoStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


