# ClusterQuorumMemberStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to [**[]ClusterQuorumMemberCondition**](ClusterQuorumMemberCondition.md) | Conditions reported by the quorum member. | [optional] 
**Id** | Pointer to **string** | A unique identifier for this quorum member. | [optional] 
**Name** | Pointer to **string** | The name of the quorum member, matching the node name. | [optional] 
**Status** | Pointer to **string** | \&quot;Started\&quot; if the member succesfully joined the quorum, \&quot;Unstarted\&quot; otherwise. | [optional] 
**Term** | Pointer to **string** | The last election term this member has participated in. | [optional] 

## Methods

### NewClusterQuorumMemberStatus

`func NewClusterQuorumMemberStatus() *ClusterQuorumMemberStatus`

NewClusterQuorumMemberStatus instantiates a new ClusterQuorumMemberStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterQuorumMemberStatusWithDefaults

`func NewClusterQuorumMemberStatusWithDefaults() *ClusterQuorumMemberStatus`

NewClusterQuorumMemberStatusWithDefaults instantiates a new ClusterQuorumMemberStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *ClusterQuorumMemberStatus) GetConditions() []ClusterQuorumMemberCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *ClusterQuorumMemberStatus) GetConditionsOk() (*[]ClusterQuorumMemberCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *ClusterQuorumMemberStatus) SetConditions(v []ClusterQuorumMemberCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *ClusterQuorumMemberStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetId

`func (o *ClusterQuorumMemberStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterQuorumMemberStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterQuorumMemberStatus) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ClusterQuorumMemberStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ClusterQuorumMemberStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterQuorumMemberStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterQuorumMemberStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterQuorumMemberStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *ClusterQuorumMemberStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterQuorumMemberStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterQuorumMemberStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClusterQuorumMemberStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTerm

`func (o *ClusterQuorumMemberStatus) GetTerm() string`

GetTerm returns the Term field if non-nil, zero value otherwise.

### GetTermOk

`func (o *ClusterQuorumMemberStatus) GetTermOk() (*string, bool)`

GetTermOk returns a tuple with the Term field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerm

`func (o *ClusterQuorumMemberStatus) SetTerm(v string)`

SetTerm sets Term field to given value.

### HasTerm

`func (o *ClusterQuorumMemberStatus) HasTerm() bool`

HasTerm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


