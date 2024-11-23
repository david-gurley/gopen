# MonitoringStatsAlertPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**ApiObjectMeta**](apiObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**MonitoringStatsAlertPolicySpec**](monitoringStatsAlertPolicySpec.md) |  | [optional] 
**Status** | Pointer to [**MonitoringStatsAlertPolicyStatus**](monitoringStatsAlertPolicyStatus.md) |  | [optional] 

## Methods

### NewMonitoringStatsAlertPolicy

`func NewMonitoringStatsAlertPolicy() *MonitoringStatsAlertPolicy`

NewMonitoringStatsAlertPolicy instantiates a new MonitoringStatsAlertPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringStatsAlertPolicyWithDefaults

`func NewMonitoringStatsAlertPolicyWithDefaults() *MonitoringStatsAlertPolicy`

NewMonitoringStatsAlertPolicyWithDefaults instantiates a new MonitoringStatsAlertPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *MonitoringStatsAlertPolicy) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *MonitoringStatsAlertPolicy) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *MonitoringStatsAlertPolicy) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *MonitoringStatsAlertPolicy) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *MonitoringStatsAlertPolicy) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *MonitoringStatsAlertPolicy) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *MonitoringStatsAlertPolicy) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *MonitoringStatsAlertPolicy) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMeta

`func (o *MonitoringStatsAlertPolicy) GetMeta() ApiObjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MonitoringStatsAlertPolicy) GetMetaOk() (*ApiObjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MonitoringStatsAlertPolicy) SetMeta(v ApiObjectMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *MonitoringStatsAlertPolicy) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSpec

`func (o *MonitoringStatsAlertPolicy) GetSpec() MonitoringStatsAlertPolicySpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *MonitoringStatsAlertPolicy) GetSpecOk() (*MonitoringStatsAlertPolicySpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *MonitoringStatsAlertPolicy) SetSpec(v MonitoringStatsAlertPolicySpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *MonitoringStatsAlertPolicy) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *MonitoringStatsAlertPolicy) GetStatus() MonitoringStatsAlertPolicyStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MonitoringStatsAlertPolicy) GetStatusOk() (*MonitoringStatsAlertPolicyStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MonitoringStatsAlertPolicy) SetStatus(v MonitoringStatsAlertPolicyStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MonitoringStatsAlertPolicy) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


