# ClusterClusterCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastTransitionTime** | Pointer to **string** | The last time the condition transitioned. | [optional] 
**Message** | Pointer to **string** | A detailed message indicating details about the transition. | [optional] 
**Reason** | Pointer to **string** | The reason for the condition&#39;s last transition. | [optional] 
**Status** | Pointer to **string** | Condition Status. | [optional] [default to "unknown"]
**Type** | Pointer to **string** | Type indicates a certain DSC condition. | [optional] [default to "healthy"]

## Methods

### NewClusterClusterCondition

`func NewClusterClusterCondition() *ClusterClusterCondition`

NewClusterClusterCondition instantiates a new ClusterClusterCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterClusterConditionWithDefaults

`func NewClusterClusterConditionWithDefaults() *ClusterClusterCondition`

NewClusterClusterConditionWithDefaults instantiates a new ClusterClusterCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastTransitionTime

`func (o *ClusterClusterCondition) GetLastTransitionTime() string`

GetLastTransitionTime returns the LastTransitionTime field if non-nil, zero value otherwise.

### GetLastTransitionTimeOk

`func (o *ClusterClusterCondition) GetLastTransitionTimeOk() (*string, bool)`

GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransitionTime

`func (o *ClusterClusterCondition) SetLastTransitionTime(v string)`

SetLastTransitionTime sets LastTransitionTime field to given value.

### HasLastTransitionTime

`func (o *ClusterClusterCondition) HasLastTransitionTime() bool`

HasLastTransitionTime returns a boolean if a field has been set.

### GetMessage

`func (o *ClusterClusterCondition) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ClusterClusterCondition) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ClusterClusterCondition) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ClusterClusterCondition) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetReason

`func (o *ClusterClusterCondition) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ClusterClusterCondition) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ClusterClusterCondition) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ClusterClusterCondition) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatus

`func (o *ClusterClusterCondition) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterClusterCondition) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterClusterCondition) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClusterClusterCondition) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *ClusterClusterCondition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClusterClusterCondition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClusterClusterCondition) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ClusterClusterCondition) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


