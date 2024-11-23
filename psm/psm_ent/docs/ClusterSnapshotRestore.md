# ClusterSnapshotRestore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**ApiObjectMeta**](apiObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**ClusterSnapshotRestoreSpec**](clusterSnapshotRestoreSpec.md) |  | [optional] 
**Status** | Pointer to [**ClusterSnapshotRestoreStatus**](clusterSnapshotRestoreStatus.md) |  | [optional] 

## Methods

### NewClusterSnapshotRestore

`func NewClusterSnapshotRestore() *ClusterSnapshotRestore`

NewClusterSnapshotRestore instantiates a new ClusterSnapshotRestore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterSnapshotRestoreWithDefaults

`func NewClusterSnapshotRestoreWithDefaults() *ClusterSnapshotRestore`

NewClusterSnapshotRestoreWithDefaults instantiates a new ClusterSnapshotRestore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ClusterSnapshotRestore) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ClusterSnapshotRestore) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ClusterSnapshotRestore) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ClusterSnapshotRestore) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *ClusterSnapshotRestore) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ClusterSnapshotRestore) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ClusterSnapshotRestore) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ClusterSnapshotRestore) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMeta

`func (o *ClusterSnapshotRestore) GetMeta() ApiObjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ClusterSnapshotRestore) GetMetaOk() (*ApiObjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ClusterSnapshotRestore) SetMeta(v ApiObjectMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ClusterSnapshotRestore) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSpec

`func (o *ClusterSnapshotRestore) GetSpec() ClusterSnapshotRestoreSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ClusterSnapshotRestore) GetSpecOk() (*ClusterSnapshotRestoreSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ClusterSnapshotRestore) SetSpec(v ClusterSnapshotRestoreSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *ClusterSnapshotRestore) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *ClusterSnapshotRestore) GetStatus() ClusterSnapshotRestoreStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterSnapshotRestore) GetStatusOk() (*ClusterSnapshotRestoreStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterSnapshotRestore) SetStatus(v ClusterSnapshotRestoreStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClusterSnapshotRestore) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


