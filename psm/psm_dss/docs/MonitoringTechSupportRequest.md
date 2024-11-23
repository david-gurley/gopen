# MonitoringTechSupportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**ApiObjectMeta**](apiObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**MonitoringTechSupportRequestSpec**](monitoringTechSupportRequestSpec.md) |  | [optional] 
**Status** | Pointer to [**MonitoringTechSupportRequestStatus**](monitoringTechSupportRequestStatus.md) |  | [optional] 

## Methods

### NewMonitoringTechSupportRequest

`func NewMonitoringTechSupportRequest() *MonitoringTechSupportRequest`

NewMonitoringTechSupportRequest instantiates a new MonitoringTechSupportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringTechSupportRequestWithDefaults

`func NewMonitoringTechSupportRequestWithDefaults() *MonitoringTechSupportRequest`

NewMonitoringTechSupportRequestWithDefaults instantiates a new MonitoringTechSupportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *MonitoringTechSupportRequest) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *MonitoringTechSupportRequest) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *MonitoringTechSupportRequest) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *MonitoringTechSupportRequest) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *MonitoringTechSupportRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *MonitoringTechSupportRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *MonitoringTechSupportRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *MonitoringTechSupportRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMeta

`func (o *MonitoringTechSupportRequest) GetMeta() ApiObjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MonitoringTechSupportRequest) GetMetaOk() (*ApiObjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MonitoringTechSupportRequest) SetMeta(v ApiObjectMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *MonitoringTechSupportRequest) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSpec

`func (o *MonitoringTechSupportRequest) GetSpec() MonitoringTechSupportRequestSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *MonitoringTechSupportRequest) GetSpecOk() (*MonitoringTechSupportRequestSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *MonitoringTechSupportRequest) SetSpec(v MonitoringTechSupportRequestSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *MonitoringTechSupportRequest) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *MonitoringTechSupportRequest) GetStatus() MonitoringTechSupportRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MonitoringTechSupportRequest) GetStatusOk() (*MonitoringTechSupportRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MonitoringTechSupportRequest) SetStatus(v MonitoringTechSupportRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MonitoringTechSupportRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


