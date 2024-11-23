# ClusterVersionStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildDate** | Pointer to **string** | Date and Time at which the source code was built. | [optional] 
**BuildVersion** | Pointer to **string** | Human friendly build version. | [optional] 
**RolloutBuildVersion** | Pointer to **string** | RolloutBuildVersion shows in progress rollout version. | [optional] 
**VcsCommit** | Pointer to **string** | Representation of ommit in version control system - e.g: hash in git. | [optional] 

## Methods

### NewClusterVersionStatus

`func NewClusterVersionStatus() *ClusterVersionStatus`

NewClusterVersionStatus instantiates a new ClusterVersionStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterVersionStatusWithDefaults

`func NewClusterVersionStatusWithDefaults() *ClusterVersionStatus`

NewClusterVersionStatusWithDefaults instantiates a new ClusterVersionStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildDate

`func (o *ClusterVersionStatus) GetBuildDate() string`

GetBuildDate returns the BuildDate field if non-nil, zero value otherwise.

### GetBuildDateOk

`func (o *ClusterVersionStatus) GetBuildDateOk() (*string, bool)`

GetBuildDateOk returns a tuple with the BuildDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildDate

`func (o *ClusterVersionStatus) SetBuildDate(v string)`

SetBuildDate sets BuildDate field to given value.

### HasBuildDate

`func (o *ClusterVersionStatus) HasBuildDate() bool`

HasBuildDate returns a boolean if a field has been set.

### GetBuildVersion

`func (o *ClusterVersionStatus) GetBuildVersion() string`

GetBuildVersion returns the BuildVersion field if non-nil, zero value otherwise.

### GetBuildVersionOk

`func (o *ClusterVersionStatus) GetBuildVersionOk() (*string, bool)`

GetBuildVersionOk returns a tuple with the BuildVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildVersion

`func (o *ClusterVersionStatus) SetBuildVersion(v string)`

SetBuildVersion sets BuildVersion field to given value.

### HasBuildVersion

`func (o *ClusterVersionStatus) HasBuildVersion() bool`

HasBuildVersion returns a boolean if a field has been set.

### GetRolloutBuildVersion

`func (o *ClusterVersionStatus) GetRolloutBuildVersion() string`

GetRolloutBuildVersion returns the RolloutBuildVersion field if non-nil, zero value otherwise.

### GetRolloutBuildVersionOk

`func (o *ClusterVersionStatus) GetRolloutBuildVersionOk() (*string, bool)`

GetRolloutBuildVersionOk returns a tuple with the RolloutBuildVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloutBuildVersion

`func (o *ClusterVersionStatus) SetRolloutBuildVersion(v string)`

SetRolloutBuildVersion sets RolloutBuildVersion field to given value.

### HasRolloutBuildVersion

`func (o *ClusterVersionStatus) HasRolloutBuildVersion() bool`

HasRolloutBuildVersion returns a boolean if a field has been set.

### GetVcsCommit

`func (o *ClusterVersionStatus) GetVcsCommit() string`

GetVcsCommit returns the VcsCommit field if non-nil, zero value otherwise.

### GetVcsCommitOk

`func (o *ClusterVersionStatus) GetVcsCommitOk() (*string, bool)`

GetVcsCommitOk returns a tuple with the VcsCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcsCommit

`func (o *ClusterVersionStatus) SetVcsCommit(v string)`

SetVcsCommit sets VcsCommit field to given value.

### HasVcsCommit

`func (o *ClusterVersionStatus) HasVcsCommit() bool`

HasVcsCommit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


