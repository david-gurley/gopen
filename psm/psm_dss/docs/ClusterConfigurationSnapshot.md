# ClusterConfigurationSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**ApiObjectMeta**](apiObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**ClusterConfigurationSnapshotSpec**](clusterConfigurationSnapshotSpec.md) |  | [optional] 
**Status** | Pointer to [**ClusterConfigurationSnapshotStatus**](clusterConfigurationSnapshotStatus.md) |  | [optional] 

## Methods

### NewClusterConfigurationSnapshot

`func NewClusterConfigurationSnapshot() *ClusterConfigurationSnapshot`

NewClusterConfigurationSnapshot instantiates a new ClusterConfigurationSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterConfigurationSnapshotWithDefaults

`func NewClusterConfigurationSnapshotWithDefaults() *ClusterConfigurationSnapshot`

NewClusterConfigurationSnapshotWithDefaults instantiates a new ClusterConfigurationSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ClusterConfigurationSnapshot) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ClusterConfigurationSnapshot) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ClusterConfigurationSnapshot) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ClusterConfigurationSnapshot) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *ClusterConfigurationSnapshot) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ClusterConfigurationSnapshot) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ClusterConfigurationSnapshot) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ClusterConfigurationSnapshot) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMeta

`func (o *ClusterConfigurationSnapshot) GetMeta() ApiObjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ClusterConfigurationSnapshot) GetMetaOk() (*ApiObjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ClusterConfigurationSnapshot) SetMeta(v ApiObjectMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ClusterConfigurationSnapshot) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSpec

`func (o *ClusterConfigurationSnapshot) GetSpec() ClusterConfigurationSnapshotSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ClusterConfigurationSnapshot) GetSpecOk() (*ClusterConfigurationSnapshotSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ClusterConfigurationSnapshot) SetSpec(v ClusterConfigurationSnapshotSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *ClusterConfigurationSnapshot) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *ClusterConfigurationSnapshot) GetStatus() ClusterConfigurationSnapshotStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterConfigurationSnapshot) GetStatusOk() (*ClusterConfigurationSnapshotStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterConfigurationSnapshot) SetStatus(v ClusterConfigurationSnapshotStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClusterConfigurationSnapshot) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


