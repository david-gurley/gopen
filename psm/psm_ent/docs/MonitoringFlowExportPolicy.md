# MonitoringFlowExportPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**ApiObjectMeta**](apiObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**MonitoringFlowExportPolicySpec**](monitoringFlowExportPolicySpec.md) |  | [optional] 
**Status** | Pointer to [**MonitoringFlowExportPolicyStatus**](monitoringFlowExportPolicyStatus.md) |  | [optional] 

## Methods

### NewMonitoringFlowExportPolicy

`func NewMonitoringFlowExportPolicy() *MonitoringFlowExportPolicy`

NewMonitoringFlowExportPolicy instantiates a new MonitoringFlowExportPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringFlowExportPolicyWithDefaults

`func NewMonitoringFlowExportPolicyWithDefaults() *MonitoringFlowExportPolicy`

NewMonitoringFlowExportPolicyWithDefaults instantiates a new MonitoringFlowExportPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *MonitoringFlowExportPolicy) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *MonitoringFlowExportPolicy) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *MonitoringFlowExportPolicy) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *MonitoringFlowExportPolicy) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *MonitoringFlowExportPolicy) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *MonitoringFlowExportPolicy) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *MonitoringFlowExportPolicy) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *MonitoringFlowExportPolicy) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMeta

`func (o *MonitoringFlowExportPolicy) GetMeta() ApiObjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MonitoringFlowExportPolicy) GetMetaOk() (*ApiObjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MonitoringFlowExportPolicy) SetMeta(v ApiObjectMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *MonitoringFlowExportPolicy) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSpec

`func (o *MonitoringFlowExportPolicy) GetSpec() MonitoringFlowExportPolicySpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *MonitoringFlowExportPolicy) GetSpecOk() (*MonitoringFlowExportPolicySpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *MonitoringFlowExportPolicy) SetSpec(v MonitoringFlowExportPolicySpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *MonitoringFlowExportPolicy) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *MonitoringFlowExportPolicy) GetStatus() MonitoringFlowExportPolicyStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MonitoringFlowExportPolicy) GetStatusOk() (*MonitoringFlowExportPolicyStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MonitoringFlowExportPolicy) SetStatus(v MonitoringFlowExportPolicyStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MonitoringFlowExportPolicy) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


