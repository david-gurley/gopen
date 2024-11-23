# ClusterDSCProfileSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplyPoliciesToEncapsulatedTraffic** | Pointer to **string** |  | [optional] [default to "disabled"]
**DeploymentTarget** | Pointer to **string** |  | [optional] [default to "host"]
**FeatureSet** | Pointer to **string** |  | [optional] [default to "smartnic"]
**InterfaceProfile** | Pointer to [**DSCProfileSpecInterfaces**](DSCProfileSpecInterfaces.md) |  | [optional] 

## Methods

### NewClusterDSCProfileSpec

`func NewClusterDSCProfileSpec() *ClusterDSCProfileSpec`

NewClusterDSCProfileSpec instantiates a new ClusterDSCProfileSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterDSCProfileSpecWithDefaults

`func NewClusterDSCProfileSpecWithDefaults() *ClusterDSCProfileSpec`

NewClusterDSCProfileSpecWithDefaults instantiates a new ClusterDSCProfileSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplyPoliciesToEncapsulatedTraffic

`func (o *ClusterDSCProfileSpec) GetApplyPoliciesToEncapsulatedTraffic() string`

GetApplyPoliciesToEncapsulatedTraffic returns the ApplyPoliciesToEncapsulatedTraffic field if non-nil, zero value otherwise.

### GetApplyPoliciesToEncapsulatedTrafficOk

`func (o *ClusterDSCProfileSpec) GetApplyPoliciesToEncapsulatedTrafficOk() (*string, bool)`

GetApplyPoliciesToEncapsulatedTrafficOk returns a tuple with the ApplyPoliciesToEncapsulatedTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyPoliciesToEncapsulatedTraffic

`func (o *ClusterDSCProfileSpec) SetApplyPoliciesToEncapsulatedTraffic(v string)`

SetApplyPoliciesToEncapsulatedTraffic sets ApplyPoliciesToEncapsulatedTraffic field to given value.

### HasApplyPoliciesToEncapsulatedTraffic

`func (o *ClusterDSCProfileSpec) HasApplyPoliciesToEncapsulatedTraffic() bool`

HasApplyPoliciesToEncapsulatedTraffic returns a boolean if a field has been set.

### GetDeploymentTarget

`func (o *ClusterDSCProfileSpec) GetDeploymentTarget() string`

GetDeploymentTarget returns the DeploymentTarget field if non-nil, zero value otherwise.

### GetDeploymentTargetOk

`func (o *ClusterDSCProfileSpec) GetDeploymentTargetOk() (*string, bool)`

GetDeploymentTargetOk returns a tuple with the DeploymentTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentTarget

`func (o *ClusterDSCProfileSpec) SetDeploymentTarget(v string)`

SetDeploymentTarget sets DeploymentTarget field to given value.

### HasDeploymentTarget

`func (o *ClusterDSCProfileSpec) HasDeploymentTarget() bool`

HasDeploymentTarget returns a boolean if a field has been set.

### GetFeatureSet

`func (o *ClusterDSCProfileSpec) GetFeatureSet() string`

GetFeatureSet returns the FeatureSet field if non-nil, zero value otherwise.

### GetFeatureSetOk

`func (o *ClusterDSCProfileSpec) GetFeatureSetOk() (*string, bool)`

GetFeatureSetOk returns a tuple with the FeatureSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureSet

`func (o *ClusterDSCProfileSpec) SetFeatureSet(v string)`

SetFeatureSet sets FeatureSet field to given value.

### HasFeatureSet

`func (o *ClusterDSCProfileSpec) HasFeatureSet() bool`

HasFeatureSet returns a boolean if a field has been set.

### GetInterfaceProfile

`func (o *ClusterDSCProfileSpec) GetInterfaceProfile() DSCProfileSpecInterfaces`

GetInterfaceProfile returns the InterfaceProfile field if non-nil, zero value otherwise.

### GetInterfaceProfileOk

`func (o *ClusterDSCProfileSpec) GetInterfaceProfileOk() (*DSCProfileSpecInterfaces, bool)`

GetInterfaceProfileOk returns a tuple with the InterfaceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceProfile

`func (o *ClusterDSCProfileSpec) SetInterfaceProfile(v DSCProfileSpecInterfaces)`

SetInterfaceProfile sets InterfaceProfile field to given value.

### HasInterfaceProfile

`func (o *ClusterDSCProfileSpec) HasInterfaceProfile() bool`

HasInterfaceProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


