# ClusterDSCProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**ApiObjectMeta**](apiObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**ClusterDSCProfileSpec**](clusterDSCProfileSpec.md) |  | [optional] 
**Status** | Pointer to [**ClusterDSCProfileStatus**](clusterDSCProfileStatus.md) |  | [optional] 

## Methods

### NewClusterDSCProfile

`func NewClusterDSCProfile() *ClusterDSCProfile`

NewClusterDSCProfile instantiates a new ClusterDSCProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterDSCProfileWithDefaults

`func NewClusterDSCProfileWithDefaults() *ClusterDSCProfile`

NewClusterDSCProfileWithDefaults instantiates a new ClusterDSCProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ClusterDSCProfile) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ClusterDSCProfile) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ClusterDSCProfile) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ClusterDSCProfile) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *ClusterDSCProfile) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ClusterDSCProfile) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ClusterDSCProfile) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ClusterDSCProfile) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMeta

`func (o *ClusterDSCProfile) GetMeta() ApiObjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ClusterDSCProfile) GetMetaOk() (*ApiObjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ClusterDSCProfile) SetMeta(v ApiObjectMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ClusterDSCProfile) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSpec

`func (o *ClusterDSCProfile) GetSpec() ClusterDSCProfileSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ClusterDSCProfile) GetSpecOk() (*ClusterDSCProfileSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ClusterDSCProfile) SetSpec(v ClusterDSCProfileSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *ClusterDSCProfile) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *ClusterDSCProfile) GetStatus() ClusterDSCProfileStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterDSCProfile) GetStatusOk() (*ClusterDSCProfileStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterDSCProfile) SetStatus(v ClusterDSCProfileStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClusterDSCProfile) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


