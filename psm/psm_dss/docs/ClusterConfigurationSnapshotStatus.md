# ClusterConfigurationSnapshotStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastSnapshot** | Pointer to [**ConfigurationSnapshotStatusConfigSaveStatus**](ConfigurationSnapshotStatusConfigSaveStatus.md) |  | [optional] 

## Methods

### NewClusterConfigurationSnapshotStatus

`func NewClusterConfigurationSnapshotStatus() *ClusterConfigurationSnapshotStatus`

NewClusterConfigurationSnapshotStatus instantiates a new ClusterConfigurationSnapshotStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterConfigurationSnapshotStatusWithDefaults

`func NewClusterConfigurationSnapshotStatusWithDefaults() *ClusterConfigurationSnapshotStatus`

NewClusterConfigurationSnapshotStatusWithDefaults instantiates a new ClusterConfigurationSnapshotStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastSnapshot

`func (o *ClusterConfigurationSnapshotStatus) GetLastSnapshot() ConfigurationSnapshotStatusConfigSaveStatus`

GetLastSnapshot returns the LastSnapshot field if non-nil, zero value otherwise.

### GetLastSnapshotOk

`func (o *ClusterConfigurationSnapshotStatus) GetLastSnapshotOk() (*ConfigurationSnapshotStatusConfigSaveStatus, bool)`

GetLastSnapshotOk returns a tuple with the LastSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSnapshot

`func (o *ClusterConfigurationSnapshotStatus) SetLastSnapshot(v ConfigurationSnapshotStatusConfigSaveStatus)`

SetLastSnapshot sets LastSnapshot field to given value.

### HasLastSnapshot

`func (o *ClusterConfigurationSnapshotStatus) HasLastSnapshot() bool`

HasLastSnapshot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


