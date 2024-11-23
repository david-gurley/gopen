# MonitoringMirrorSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**ApiObjectMeta**](apiObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**MonitoringMirrorSessionSpec**](monitoringMirrorSessionSpec.md) |  | [optional] 
**Status** | Pointer to [**MonitoringMirrorSessionStatus**](monitoringMirrorSessionStatus.md) |  | [optional] 

## Methods

### NewMonitoringMirrorSession

`func NewMonitoringMirrorSession() *MonitoringMirrorSession`

NewMonitoringMirrorSession instantiates a new MonitoringMirrorSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringMirrorSessionWithDefaults

`func NewMonitoringMirrorSessionWithDefaults() *MonitoringMirrorSession`

NewMonitoringMirrorSessionWithDefaults instantiates a new MonitoringMirrorSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *MonitoringMirrorSession) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *MonitoringMirrorSession) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *MonitoringMirrorSession) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *MonitoringMirrorSession) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *MonitoringMirrorSession) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *MonitoringMirrorSession) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *MonitoringMirrorSession) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *MonitoringMirrorSession) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMeta

`func (o *MonitoringMirrorSession) GetMeta() ApiObjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MonitoringMirrorSession) GetMetaOk() (*ApiObjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MonitoringMirrorSession) SetMeta(v ApiObjectMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *MonitoringMirrorSession) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSpec

`func (o *MonitoringMirrorSession) GetSpec() MonitoringMirrorSessionSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *MonitoringMirrorSession) GetSpecOk() (*MonitoringMirrorSessionSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *MonitoringMirrorSession) SetSpec(v MonitoringMirrorSessionSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *MonitoringMirrorSession) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *MonitoringMirrorSession) GetStatus() MonitoringMirrorSessionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MonitoringMirrorSession) GetStatusOk() (*MonitoringMirrorSessionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MonitoringMirrorSession) SetStatus(v MonitoringMirrorSessionStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MonitoringMirrorSession) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


