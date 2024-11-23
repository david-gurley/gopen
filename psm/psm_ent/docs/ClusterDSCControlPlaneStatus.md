# ClusterDSCControlPlaneStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BgpStatus** | Pointer to [**[]ClusterPeerStatus**](ClusterPeerStatus.md) | BGPStatus contains the list of BGP peers and their status. | [optional] 
**LastUpdatedTime** | Pointer to **string** | The last time the control plane was updated. | [optional] 
**Message** | Pointer to **string** | A message indicating details about the control plane status. | [optional] 

## Methods

### NewClusterDSCControlPlaneStatus

`func NewClusterDSCControlPlaneStatus() *ClusterDSCControlPlaneStatus`

NewClusterDSCControlPlaneStatus instantiates a new ClusterDSCControlPlaneStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterDSCControlPlaneStatusWithDefaults

`func NewClusterDSCControlPlaneStatusWithDefaults() *ClusterDSCControlPlaneStatus`

NewClusterDSCControlPlaneStatusWithDefaults instantiates a new ClusterDSCControlPlaneStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBgpStatus

`func (o *ClusterDSCControlPlaneStatus) GetBgpStatus() []ClusterPeerStatus`

GetBgpStatus returns the BgpStatus field if non-nil, zero value otherwise.

### GetBgpStatusOk

`func (o *ClusterDSCControlPlaneStatus) GetBgpStatusOk() (*[]ClusterPeerStatus, bool)`

GetBgpStatusOk returns a tuple with the BgpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpStatus

`func (o *ClusterDSCControlPlaneStatus) SetBgpStatus(v []ClusterPeerStatus)`

SetBgpStatus sets BgpStatus field to given value.

### HasBgpStatus

`func (o *ClusterDSCControlPlaneStatus) HasBgpStatus() bool`

HasBgpStatus returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *ClusterDSCControlPlaneStatus) GetLastUpdatedTime() string`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *ClusterDSCControlPlaneStatus) GetLastUpdatedTimeOk() (*string, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *ClusterDSCControlPlaneStatus) SetLastUpdatedTime(v string)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *ClusterDSCControlPlaneStatus) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetMessage

`func (o *ClusterDSCControlPlaneStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ClusterDSCControlPlaneStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ClusterDSCControlPlaneStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ClusterDSCControlPlaneStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


