# ClusterLicense

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**ApiObjectMeta**](apiObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**ClusterLicenseSpec**](clusterLicenseSpec.md) |  | [optional] 
**Status** | Pointer to [**ClusterLicenseStatus**](clusterLicenseStatus.md) |  | [optional] 

## Methods

### NewClusterLicense

`func NewClusterLicense() *ClusterLicense`

NewClusterLicense instantiates a new ClusterLicense object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterLicenseWithDefaults

`func NewClusterLicenseWithDefaults() *ClusterLicense`

NewClusterLicenseWithDefaults instantiates a new ClusterLicense object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ClusterLicense) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ClusterLicense) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ClusterLicense) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ClusterLicense) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *ClusterLicense) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ClusterLicense) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ClusterLicense) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ClusterLicense) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMeta

`func (o *ClusterLicense) GetMeta() ApiObjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ClusterLicense) GetMetaOk() (*ApiObjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ClusterLicense) SetMeta(v ApiObjectMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ClusterLicense) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSpec

`func (o *ClusterLicense) GetSpec() ClusterLicenseSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ClusterLicense) GetSpecOk() (*ClusterLicenseSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ClusterLicense) SetSpec(v ClusterLicenseSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *ClusterLicense) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *ClusterLicense) GetStatus() ClusterLicenseStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterLicense) GetStatusOk() (*ClusterLicenseStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterLicense) SetStatus(v ClusterLicenseStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClusterLicense) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


