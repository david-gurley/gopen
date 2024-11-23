# ClusterNodeStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to [**[]ClusterNodeCondition**](ClusterNodeCondition.md) | List of current node conditions. | [optional] 
**Phase** | Pointer to **string** | Current lifecycle phase of the node. | [optional] [default to "unknown"]
**Quorum** | Pointer to **bool** | Quorum node or not. | [optional] 

## Methods

### NewClusterNodeStatus

`func NewClusterNodeStatus() *ClusterNodeStatus`

NewClusterNodeStatus instantiates a new ClusterNodeStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterNodeStatusWithDefaults

`func NewClusterNodeStatusWithDefaults() *ClusterNodeStatus`

NewClusterNodeStatusWithDefaults instantiates a new ClusterNodeStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *ClusterNodeStatus) GetConditions() []ClusterNodeCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *ClusterNodeStatus) GetConditionsOk() (*[]ClusterNodeCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *ClusterNodeStatus) SetConditions(v []ClusterNodeCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *ClusterNodeStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetPhase

`func (o *ClusterNodeStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *ClusterNodeStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *ClusterNodeStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *ClusterNodeStatus) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetQuorum

`func (o *ClusterNodeStatus) GetQuorum() bool`

GetQuorum returns the Quorum field if non-nil, zero value otherwise.

### GetQuorumOk

`func (o *ClusterNodeStatus) GetQuorumOk() (*bool, bool)`

GetQuorumOk returns a tuple with the Quorum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuorum

`func (o *ClusterNodeStatus) SetQuorum(v bool)`

SetQuorum sets Quorum field to given value.

### HasQuorum

`func (o *ClusterNodeStatus) HasQuorum() bool`

HasQuorum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


