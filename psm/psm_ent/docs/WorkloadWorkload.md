# WorkloadWorkload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**ApiObjectMeta**](apiObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**WorkloadWorkloadSpec**](workloadWorkloadSpec.md) |  | [optional] 
**Status** | Pointer to [**WorkloadWorkloadStatus**](workloadWorkloadStatus.md) |  | [optional] 

## Methods

### NewWorkloadWorkload

`func NewWorkloadWorkload() *WorkloadWorkload`

NewWorkloadWorkload instantiates a new WorkloadWorkload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadWorkloadWithDefaults

`func NewWorkloadWorkloadWithDefaults() *WorkloadWorkload`

NewWorkloadWorkloadWithDefaults instantiates a new WorkloadWorkload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *WorkloadWorkload) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *WorkloadWorkload) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *WorkloadWorkload) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *WorkloadWorkload) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *WorkloadWorkload) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *WorkloadWorkload) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *WorkloadWorkload) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *WorkloadWorkload) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMeta

`func (o *WorkloadWorkload) GetMeta() ApiObjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *WorkloadWorkload) GetMetaOk() (*ApiObjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *WorkloadWorkload) SetMeta(v ApiObjectMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *WorkloadWorkload) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSpec

`func (o *WorkloadWorkload) GetSpec() WorkloadWorkloadSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *WorkloadWorkload) GetSpecOk() (*WorkloadWorkloadSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *WorkloadWorkload) SetSpec(v WorkloadWorkloadSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *WorkloadWorkload) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *WorkloadWorkload) GetStatus() WorkloadWorkloadStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkloadWorkload) GetStatusOk() (*WorkloadWorkloadStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkloadWorkload) SetStatus(v WorkloadWorkloadStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkloadWorkload) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


