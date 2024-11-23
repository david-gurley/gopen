# ClusterClusterStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthBootstrapped** | Pointer to **bool** | AuthBootstrapped indicates whether the Cluster has Completed BootStrap of Auth. | [optional] 
**Conditions** | Pointer to [**[]ClusterClusterCondition**](ClusterClusterCondition.md) | List of current cluster conditions. | [optional] 
**CurrentTime** | Pointer to **time.Time** | CurrentTime is current time of a cluster. | [optional] 
**LastLeaderTransitionTime** | Pointer to **time.Time** | LastLeaderTransitionTime is when the leadership changed last time. | [optional] 
**Leader** | Pointer to **string** | Leader contains the node name of the cluster leader. | [optional] 
**QuorumStatus** | Pointer to [**ClusterQuorumStatus**](clusterQuorumStatus.md) |  | [optional] 
**RecoveryKeysDownloaded** | Pointer to **bool** | RecoveryKeysDownloaded indicates whether keys have been downloaded since the cluster has been bootstrapped. | [optional] 

## Methods

### NewClusterClusterStatus

`func NewClusterClusterStatus() *ClusterClusterStatus`

NewClusterClusterStatus instantiates a new ClusterClusterStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterClusterStatusWithDefaults

`func NewClusterClusterStatusWithDefaults() *ClusterClusterStatus`

NewClusterClusterStatusWithDefaults instantiates a new ClusterClusterStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthBootstrapped

`func (o *ClusterClusterStatus) GetAuthBootstrapped() bool`

GetAuthBootstrapped returns the AuthBootstrapped field if non-nil, zero value otherwise.

### GetAuthBootstrappedOk

`func (o *ClusterClusterStatus) GetAuthBootstrappedOk() (*bool, bool)`

GetAuthBootstrappedOk returns a tuple with the AuthBootstrapped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthBootstrapped

`func (o *ClusterClusterStatus) SetAuthBootstrapped(v bool)`

SetAuthBootstrapped sets AuthBootstrapped field to given value.

### HasAuthBootstrapped

`func (o *ClusterClusterStatus) HasAuthBootstrapped() bool`

HasAuthBootstrapped returns a boolean if a field has been set.

### GetConditions

`func (o *ClusterClusterStatus) GetConditions() []ClusterClusterCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *ClusterClusterStatus) GetConditionsOk() (*[]ClusterClusterCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *ClusterClusterStatus) SetConditions(v []ClusterClusterCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *ClusterClusterStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetCurrentTime

`func (o *ClusterClusterStatus) GetCurrentTime() time.Time`

GetCurrentTime returns the CurrentTime field if non-nil, zero value otherwise.

### GetCurrentTimeOk

`func (o *ClusterClusterStatus) GetCurrentTimeOk() (*time.Time, bool)`

GetCurrentTimeOk returns a tuple with the CurrentTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTime

`func (o *ClusterClusterStatus) SetCurrentTime(v time.Time)`

SetCurrentTime sets CurrentTime field to given value.

### HasCurrentTime

`func (o *ClusterClusterStatus) HasCurrentTime() bool`

HasCurrentTime returns a boolean if a field has been set.

### GetLastLeaderTransitionTime

`func (o *ClusterClusterStatus) GetLastLeaderTransitionTime() time.Time`

GetLastLeaderTransitionTime returns the LastLeaderTransitionTime field if non-nil, zero value otherwise.

### GetLastLeaderTransitionTimeOk

`func (o *ClusterClusterStatus) GetLastLeaderTransitionTimeOk() (*time.Time, bool)`

GetLastLeaderTransitionTimeOk returns a tuple with the LastLeaderTransitionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLeaderTransitionTime

`func (o *ClusterClusterStatus) SetLastLeaderTransitionTime(v time.Time)`

SetLastLeaderTransitionTime sets LastLeaderTransitionTime field to given value.

### HasLastLeaderTransitionTime

`func (o *ClusterClusterStatus) HasLastLeaderTransitionTime() bool`

HasLastLeaderTransitionTime returns a boolean if a field has been set.

### GetLeader

`func (o *ClusterClusterStatus) GetLeader() string`

GetLeader returns the Leader field if non-nil, zero value otherwise.

### GetLeaderOk

`func (o *ClusterClusterStatus) GetLeaderOk() (*string, bool)`

GetLeaderOk returns a tuple with the Leader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeader

`func (o *ClusterClusterStatus) SetLeader(v string)`

SetLeader sets Leader field to given value.

### HasLeader

`func (o *ClusterClusterStatus) HasLeader() bool`

HasLeader returns a boolean if a field has been set.

### GetQuorumStatus

`func (o *ClusterClusterStatus) GetQuorumStatus() ClusterQuorumStatus`

GetQuorumStatus returns the QuorumStatus field if non-nil, zero value otherwise.

### GetQuorumStatusOk

`func (o *ClusterClusterStatus) GetQuorumStatusOk() (*ClusterQuorumStatus, bool)`

GetQuorumStatusOk returns a tuple with the QuorumStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuorumStatus

`func (o *ClusterClusterStatus) SetQuorumStatus(v ClusterQuorumStatus)`

SetQuorumStatus sets QuorumStatus field to given value.

### HasQuorumStatus

`func (o *ClusterClusterStatus) HasQuorumStatus() bool`

HasQuorumStatus returns a boolean if a field has been set.

### GetRecoveryKeysDownloaded

`func (o *ClusterClusterStatus) GetRecoveryKeysDownloaded() bool`

GetRecoveryKeysDownloaded returns the RecoveryKeysDownloaded field if non-nil, zero value otherwise.

### GetRecoveryKeysDownloadedOk

`func (o *ClusterClusterStatus) GetRecoveryKeysDownloadedOk() (*bool, bool)`

GetRecoveryKeysDownloadedOk returns a tuple with the RecoveryKeysDownloaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryKeysDownloaded

`func (o *ClusterClusterStatus) SetRecoveryKeysDownloaded(v bool)`

SetRecoveryKeysDownloaded sets RecoveryKeysDownloaded field to given value.

### HasRecoveryKeysDownloaded

`func (o *ClusterClusterStatus) HasRecoveryKeysDownloaded() bool`

HasRecoveryKeysDownloaded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


