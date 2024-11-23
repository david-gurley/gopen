# MonitoringTroubleshootingSessionStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** |  | [optional] [default to "running"]
**TroubleshootingResults** | Pointer to [**[]MonitoringTsResult**](MonitoringTsResult.md) | report is generated each time troubleshooting session is activated i.e time-window. | [optional] 

## Methods

### NewMonitoringTroubleshootingSessionStatus

`func NewMonitoringTroubleshootingSessionStatus() *MonitoringTroubleshootingSessionStatus`

NewMonitoringTroubleshootingSessionStatus instantiates a new MonitoringTroubleshootingSessionStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringTroubleshootingSessionStatusWithDefaults

`func NewMonitoringTroubleshootingSessionStatusWithDefaults() *MonitoringTroubleshootingSessionStatus`

NewMonitoringTroubleshootingSessionStatusWithDefaults instantiates a new MonitoringTroubleshootingSessionStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *MonitoringTroubleshootingSessionStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MonitoringTroubleshootingSessionStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MonitoringTroubleshootingSessionStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *MonitoringTroubleshootingSessionStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTroubleshootingResults

`func (o *MonitoringTroubleshootingSessionStatus) GetTroubleshootingResults() []MonitoringTsResult`

GetTroubleshootingResults returns the TroubleshootingResults field if non-nil, zero value otherwise.

### GetTroubleshootingResultsOk

`func (o *MonitoringTroubleshootingSessionStatus) GetTroubleshootingResultsOk() (*[]MonitoringTsResult, bool)`

GetTroubleshootingResultsOk returns a tuple with the TroubleshootingResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTroubleshootingResults

`func (o *MonitoringTroubleshootingSessionStatus) SetTroubleshootingResults(v []MonitoringTsResult)`

SetTroubleshootingResults sets TroubleshootingResults field to given value.

### HasTroubleshootingResults

`func (o *MonitoringTroubleshootingSessionStatus) HasTroubleshootingResults() bool`

HasTroubleshootingResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


