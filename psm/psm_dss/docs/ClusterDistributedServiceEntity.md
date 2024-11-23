# ClusterDistributedServiceEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**ApiObjectMeta**](apiObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**ClusterDistributedServiceEntitySpec**](clusterDistributedServiceEntitySpec.md) |  | [optional] 
**Status** | Pointer to [**ClusterDistributedServiceEntityStatus**](clusterDistributedServiceEntityStatus.md) |  | [optional] 

## Methods

### NewClusterDistributedServiceEntity

`func NewClusterDistributedServiceEntity() *ClusterDistributedServiceEntity`

NewClusterDistributedServiceEntity instantiates a new ClusterDistributedServiceEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterDistributedServiceEntityWithDefaults

`func NewClusterDistributedServiceEntityWithDefaults() *ClusterDistributedServiceEntity`

NewClusterDistributedServiceEntityWithDefaults instantiates a new ClusterDistributedServiceEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ClusterDistributedServiceEntity) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ClusterDistributedServiceEntity) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ClusterDistributedServiceEntity) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ClusterDistributedServiceEntity) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *ClusterDistributedServiceEntity) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ClusterDistributedServiceEntity) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ClusterDistributedServiceEntity) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ClusterDistributedServiceEntity) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMeta

`func (o *ClusterDistributedServiceEntity) GetMeta() ApiObjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ClusterDistributedServiceEntity) GetMetaOk() (*ApiObjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ClusterDistributedServiceEntity) SetMeta(v ApiObjectMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ClusterDistributedServiceEntity) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSpec

`func (o *ClusterDistributedServiceEntity) GetSpec() ClusterDistributedServiceEntitySpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ClusterDistributedServiceEntity) GetSpecOk() (*ClusterDistributedServiceEntitySpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ClusterDistributedServiceEntity) SetSpec(v ClusterDistributedServiceEntitySpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *ClusterDistributedServiceEntity) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *ClusterDistributedServiceEntity) GetStatus() ClusterDistributedServiceEntityStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterDistributedServiceEntity) GetStatusOk() (*ClusterDistributedServiceEntityStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterDistributedServiceEntity) SetStatus(v ClusterDistributedServiceEntityStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClusterDistributedServiceEntity) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


