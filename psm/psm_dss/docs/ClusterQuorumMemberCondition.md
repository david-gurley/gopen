# ClusterQuorumMemberCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastTransitionTime** | Pointer to **time.Time** | The last time the condition transitioned. | [optional] 
**Status** | Pointer to **string** | Condition Status. | [optional] [default to "unknown"]
**Type** | Pointer to **string** | Type indicates a certain node condition. | [optional] [default to "healthy"]

## Methods

### NewClusterQuorumMemberCondition

`func NewClusterQuorumMemberCondition() *ClusterQuorumMemberCondition`

NewClusterQuorumMemberCondition instantiates a new ClusterQuorumMemberCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterQuorumMemberConditionWithDefaults

`func NewClusterQuorumMemberConditionWithDefaults() *ClusterQuorumMemberCondition`

NewClusterQuorumMemberConditionWithDefaults instantiates a new ClusterQuorumMemberCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastTransitionTime

`func (o *ClusterQuorumMemberCondition) GetLastTransitionTime() time.Time`

GetLastTransitionTime returns the LastTransitionTime field if non-nil, zero value otherwise.

### GetLastTransitionTimeOk

`func (o *ClusterQuorumMemberCondition) GetLastTransitionTimeOk() (*time.Time, bool)`

GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransitionTime

`func (o *ClusterQuorumMemberCondition) SetLastTransitionTime(v time.Time)`

SetLastTransitionTime sets LastTransitionTime field to given value.

### HasLastTransitionTime

`func (o *ClusterQuorumMemberCondition) HasLastTransitionTime() bool`

HasLastTransitionTime returns a boolean if a field has been set.

### GetStatus

`func (o *ClusterQuorumMemberCondition) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterQuorumMemberCondition) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterQuorumMemberCondition) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClusterQuorumMemberCondition) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *ClusterQuorumMemberCondition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClusterQuorumMemberCondition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClusterQuorumMemberCondition) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ClusterQuorumMemberCondition) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


