# ClusterVersionSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoRolloutDscVersion** | Pointer to **string** | AutoRolloutDSCVersion indicates the version DSCs will be automatically upgraded upon admission. Empty value indicates no auto-rollout. | [optional] 

## Methods

### NewClusterVersionSpec

`func NewClusterVersionSpec() *ClusterVersionSpec`

NewClusterVersionSpec instantiates a new ClusterVersionSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterVersionSpecWithDefaults

`func NewClusterVersionSpecWithDefaults() *ClusterVersionSpec`

NewClusterVersionSpecWithDefaults instantiates a new ClusterVersionSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoRolloutDscVersion

`func (o *ClusterVersionSpec) GetAutoRolloutDscVersion() string`

GetAutoRolloutDscVersion returns the AutoRolloutDscVersion field if non-nil, zero value otherwise.

### GetAutoRolloutDscVersionOk

`func (o *ClusterVersionSpec) GetAutoRolloutDscVersionOk() (*string, bool)`

GetAutoRolloutDscVersionOk returns a tuple with the AutoRolloutDscVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRolloutDscVersion

`func (o *ClusterVersionSpec) SetAutoRolloutDscVersion(v string)`

SetAutoRolloutDscVersion sets AutoRolloutDscVersion field to given value.

### HasAutoRolloutDscVersion

`func (o *ClusterVersionSpec) HasAutoRolloutDscVersion() bool`

HasAutoRolloutDscVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


