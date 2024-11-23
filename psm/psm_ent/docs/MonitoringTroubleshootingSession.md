# MonitoringTroubleshootingSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**ApiObjectMeta**](apiObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**MonitoringTroubleshootingSessionSpec**](monitoringTroubleshootingSessionSpec.md) |  | [optional] 
**Status** | Pointer to [**MonitoringTroubleshootingSessionStatus**](monitoringTroubleshootingSessionStatus.md) |  | [optional] 

## Methods

### NewMonitoringTroubleshootingSession

`func NewMonitoringTroubleshootingSession() *MonitoringTroubleshootingSession`

NewMonitoringTroubleshootingSession instantiates a new MonitoringTroubleshootingSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringTroubleshootingSessionWithDefaults

`func NewMonitoringTroubleshootingSessionWithDefaults() *MonitoringTroubleshootingSession`

NewMonitoringTroubleshootingSessionWithDefaults instantiates a new MonitoringTroubleshootingSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *MonitoringTroubleshootingSession) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *MonitoringTroubleshootingSession) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *MonitoringTroubleshootingSession) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *MonitoringTroubleshootingSession) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *MonitoringTroubleshootingSession) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *MonitoringTroubleshootingSession) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *MonitoringTroubleshootingSession) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *MonitoringTroubleshootingSession) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMeta

`func (o *MonitoringTroubleshootingSession) GetMeta() ApiObjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MonitoringTroubleshootingSession) GetMetaOk() (*ApiObjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MonitoringTroubleshootingSession) SetMeta(v ApiObjectMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *MonitoringTroubleshootingSession) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSpec

`func (o *MonitoringTroubleshootingSession) GetSpec() MonitoringTroubleshootingSessionSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *MonitoringTroubleshootingSession) GetSpecOk() (*MonitoringTroubleshootingSessionSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *MonitoringTroubleshootingSession) SetSpec(v MonitoringTroubleshootingSessionSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *MonitoringTroubleshootingSession) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *MonitoringTroubleshootingSession) GetStatus() MonitoringTroubleshootingSessionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MonitoringTroubleshootingSession) GetStatusOk() (*MonitoringTroubleshootingSessionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MonitoringTroubleshootingSession) SetStatus(v MonitoringTroubleshootingSessionStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MonitoringTroubleshootingSession) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


