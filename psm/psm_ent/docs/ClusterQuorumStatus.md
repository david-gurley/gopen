# ClusterQuorumStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Members** | Pointer to [**[]ClusterQuorumMemberStatus**](ClusterQuorumMemberStatus.md) |  | [optional] 

## Methods

### NewClusterQuorumStatus

`func NewClusterQuorumStatus() *ClusterQuorumStatus`

NewClusterQuorumStatus instantiates a new ClusterQuorumStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterQuorumStatusWithDefaults

`func NewClusterQuorumStatusWithDefaults() *ClusterQuorumStatus`

NewClusterQuorumStatusWithDefaults instantiates a new ClusterQuorumStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMembers

`func (o *ClusterQuorumStatus) GetMembers() []ClusterQuorumMemberStatus`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *ClusterQuorumStatus) GetMembersOk() (*[]ClusterQuorumMemberStatus, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *ClusterQuorumStatus) SetMembers(v []ClusterQuorumMemberStatus)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *ClusterQuorumStatus) HasMembers() bool`

HasMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


