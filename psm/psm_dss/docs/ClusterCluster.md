# ClusterCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**ApiObjectMeta**](apiObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**ClusterClusterSpec**](clusterClusterSpec.md) |  | [optional] 
**Status** | Pointer to [**ClusterClusterStatus**](clusterClusterStatus.md) |  | [optional] 

## Methods

### NewClusterCluster

`func NewClusterCluster() *ClusterCluster`

NewClusterCluster instantiates a new ClusterCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterClusterWithDefaults

`func NewClusterClusterWithDefaults() *ClusterCluster`

NewClusterClusterWithDefaults instantiates a new ClusterCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ClusterCluster) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ClusterCluster) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ClusterCluster) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ClusterCluster) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *ClusterCluster) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ClusterCluster) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ClusterCluster) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ClusterCluster) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMeta

`func (o *ClusterCluster) GetMeta() ApiObjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ClusterCluster) GetMetaOk() (*ApiObjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ClusterCluster) SetMeta(v ApiObjectMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ClusterCluster) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSpec

`func (o *ClusterCluster) GetSpec() ClusterClusterSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ClusterCluster) GetSpecOk() (*ClusterClusterSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ClusterCluster) SetSpec(v ClusterClusterSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *ClusterCluster) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *ClusterCluster) GetStatus() ClusterClusterStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterCluster) GetStatusOk() (*ClusterClusterStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterCluster) SetStatus(v ClusterClusterStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClusterCluster) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


