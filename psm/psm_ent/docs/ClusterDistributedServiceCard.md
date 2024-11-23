# ClusterDistributedServiceCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**ApiObjectMeta**](apiObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**ClusterDistributedServiceCardSpec**](clusterDistributedServiceCardSpec.md) |  | [optional] 
**Status** | Pointer to [**ClusterDistributedServiceCardStatus**](clusterDistributedServiceCardStatus.md) |  | [optional] 

## Methods

### NewClusterDistributedServiceCard

`func NewClusterDistributedServiceCard() *ClusterDistributedServiceCard`

NewClusterDistributedServiceCard instantiates a new ClusterDistributedServiceCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterDistributedServiceCardWithDefaults

`func NewClusterDistributedServiceCardWithDefaults() *ClusterDistributedServiceCard`

NewClusterDistributedServiceCardWithDefaults instantiates a new ClusterDistributedServiceCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ClusterDistributedServiceCard) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ClusterDistributedServiceCard) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ClusterDistributedServiceCard) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ClusterDistributedServiceCard) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *ClusterDistributedServiceCard) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ClusterDistributedServiceCard) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ClusterDistributedServiceCard) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ClusterDistributedServiceCard) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMeta

`func (o *ClusterDistributedServiceCard) GetMeta() ApiObjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ClusterDistributedServiceCard) GetMetaOk() (*ApiObjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ClusterDistributedServiceCard) SetMeta(v ApiObjectMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ClusterDistributedServiceCard) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSpec

`func (o *ClusterDistributedServiceCard) GetSpec() ClusterDistributedServiceCardSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ClusterDistributedServiceCard) GetSpecOk() (*ClusterDistributedServiceCardSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ClusterDistributedServiceCard) SetSpec(v ClusterDistributedServiceCardSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *ClusterDistributedServiceCard) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *ClusterDistributedServiceCard) GetStatus() ClusterDistributedServiceCardStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterDistributedServiceCard) GetStatusOk() (*ClusterDistributedServiceCardStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterDistributedServiceCard) SetStatus(v ClusterDistributedServiceCardStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClusterDistributedServiceCard) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


