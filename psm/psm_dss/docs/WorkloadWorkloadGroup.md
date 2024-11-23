# WorkloadWorkloadGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**ApiObjectMeta**](apiObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**WorkloadWorkloadGroupSpec**](workloadWorkloadGroupSpec.md) |  | [optional] 
**Status** | Pointer to [**WorkloadWorkloadGroupStatus**](workloadWorkloadGroupStatus.md) |  | [optional] 

## Methods

### NewWorkloadWorkloadGroup

`func NewWorkloadWorkloadGroup() *WorkloadWorkloadGroup`

NewWorkloadWorkloadGroup instantiates a new WorkloadWorkloadGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadWorkloadGroupWithDefaults

`func NewWorkloadWorkloadGroupWithDefaults() *WorkloadWorkloadGroup`

NewWorkloadWorkloadGroupWithDefaults instantiates a new WorkloadWorkloadGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *WorkloadWorkloadGroup) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *WorkloadWorkloadGroup) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *WorkloadWorkloadGroup) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *WorkloadWorkloadGroup) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *WorkloadWorkloadGroup) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *WorkloadWorkloadGroup) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *WorkloadWorkloadGroup) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *WorkloadWorkloadGroup) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMeta

`func (o *WorkloadWorkloadGroup) GetMeta() ApiObjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *WorkloadWorkloadGroup) GetMetaOk() (*ApiObjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *WorkloadWorkloadGroup) SetMeta(v ApiObjectMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *WorkloadWorkloadGroup) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSpec

`func (o *WorkloadWorkloadGroup) GetSpec() WorkloadWorkloadGroupSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *WorkloadWorkloadGroup) GetSpecOk() (*WorkloadWorkloadGroupSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *WorkloadWorkloadGroup) SetSpec(v WorkloadWorkloadGroupSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *WorkloadWorkloadGroup) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *WorkloadWorkloadGroup) GetStatus() WorkloadWorkloadGroupStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkloadWorkloadGroup) GetStatusOk() (*WorkloadWorkloadGroupStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkloadWorkloadGroup) SetStatus(v WorkloadWorkloadGroupStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkloadWorkloadGroup) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


