# ClusterConfigurationSnapshotSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | Pointer to [**ClusterSnapshotDestination**](clusterSnapshotDestination.md) |  | [optional] 

## Methods

### NewClusterConfigurationSnapshotSpec

`func NewClusterConfigurationSnapshotSpec() *ClusterConfigurationSnapshotSpec`

NewClusterConfigurationSnapshotSpec instantiates a new ClusterConfigurationSnapshotSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterConfigurationSnapshotSpecWithDefaults

`func NewClusterConfigurationSnapshotSpecWithDefaults() *ClusterConfigurationSnapshotSpec`

NewClusterConfigurationSnapshotSpecWithDefaults instantiates a new ClusterConfigurationSnapshotSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *ClusterConfigurationSnapshotSpec) GetDestination() ClusterSnapshotDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *ClusterConfigurationSnapshotSpec) GetDestinationOk() (*ClusterSnapshotDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *ClusterConfigurationSnapshotSpec) SetDestination(v ClusterSnapshotDestination)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *ClusterConfigurationSnapshotSpec) HasDestination() bool`

HasDestination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


