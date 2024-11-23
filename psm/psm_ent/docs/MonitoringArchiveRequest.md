# MonitoringArchiveRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**ApiObjectMeta**](apiObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**MonitoringArchiveRequestSpec**](monitoringArchiveRequestSpec.md) |  | [optional] 
**Status** | Pointer to [**MonitoringArchiveRequestStatus**](monitoringArchiveRequestStatus.md) |  | [optional] 

## Methods

### NewMonitoringArchiveRequest

`func NewMonitoringArchiveRequest() *MonitoringArchiveRequest`

NewMonitoringArchiveRequest instantiates a new MonitoringArchiveRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringArchiveRequestWithDefaults

`func NewMonitoringArchiveRequestWithDefaults() *MonitoringArchiveRequest`

NewMonitoringArchiveRequestWithDefaults instantiates a new MonitoringArchiveRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *MonitoringArchiveRequest) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *MonitoringArchiveRequest) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *MonitoringArchiveRequest) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *MonitoringArchiveRequest) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *MonitoringArchiveRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *MonitoringArchiveRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *MonitoringArchiveRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *MonitoringArchiveRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMeta

`func (o *MonitoringArchiveRequest) GetMeta() ApiObjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MonitoringArchiveRequest) GetMetaOk() (*ApiObjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MonitoringArchiveRequest) SetMeta(v ApiObjectMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *MonitoringArchiveRequest) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSpec

`func (o *MonitoringArchiveRequest) GetSpec() MonitoringArchiveRequestSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *MonitoringArchiveRequest) GetSpecOk() (*MonitoringArchiveRequestSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *MonitoringArchiveRequest) SetSpec(v MonitoringArchiveRequestSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *MonitoringArchiveRequest) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *MonitoringArchiveRequest) GetStatus() MonitoringArchiveRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MonitoringArchiveRequest) GetStatusOk() (*MonitoringArchiveRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MonitoringArchiveRequest) SetStatus(v MonitoringArchiveRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MonitoringArchiveRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


