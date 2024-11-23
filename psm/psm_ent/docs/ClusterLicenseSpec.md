# ClusterLicenseSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Features** | Pointer to [**[]ClusterFeature**](ClusterFeature.md) | List of Feature licences applied. | [optional] 

## Methods

### NewClusterLicenseSpec

`func NewClusterLicenseSpec() *ClusterLicenseSpec`

NewClusterLicenseSpec instantiates a new ClusterLicenseSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterLicenseSpecWithDefaults

`func NewClusterLicenseSpecWithDefaults() *ClusterLicenseSpec`

NewClusterLicenseSpecWithDefaults instantiates a new ClusterLicenseSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatures

`func (o *ClusterLicenseSpec) GetFeatures() []ClusterFeature`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *ClusterLicenseSpec) GetFeaturesOk() (*[]ClusterFeature, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *ClusterLicenseSpec) SetFeatures(v []ClusterFeature)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *ClusterLicenseSpec) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

