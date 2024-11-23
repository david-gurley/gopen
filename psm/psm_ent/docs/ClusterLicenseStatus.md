# ClusterLicenseStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Features** | Pointer to [**[]ClusterFeatureStatus**](ClusterFeatureStatus.md) | Status of current Licenced features. | [optional] 
**Unknown** | Pointer to **[]string** | Licenses that are not understood by the current running version of software. | [optional] 

## Methods

### NewClusterLicenseStatus

`func NewClusterLicenseStatus() *ClusterLicenseStatus`

NewClusterLicenseStatus instantiates a new ClusterLicenseStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterLicenseStatusWithDefaults

`func NewClusterLicenseStatusWithDefaults() *ClusterLicenseStatus`

NewClusterLicenseStatusWithDefaults instantiates a new ClusterLicenseStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatures

`func (o *ClusterLicenseStatus) GetFeatures() []ClusterFeatureStatus`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *ClusterLicenseStatus) GetFeaturesOk() (*[]ClusterFeatureStatus, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *ClusterLicenseStatus) SetFeatures(v []ClusterFeatureStatus)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *ClusterLicenseStatus) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetUnknown

`func (o *ClusterLicenseStatus) GetUnknown() []string`

GetUnknown returns the Unknown field if non-nil, zero value otherwise.

### GetUnknownOk

`func (o *ClusterLicenseStatus) GetUnknownOk() (*[]string, bool)`

GetUnknownOk returns a tuple with the Unknown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknown

`func (o *ClusterLicenseStatus) SetUnknown(v []string)`

SetUnknown sets Unknown field to given value.

### HasUnknown

`func (o *ClusterLicenseStatus) HasUnknown() bool`

HasUnknown returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


