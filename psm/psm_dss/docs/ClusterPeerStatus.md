# ClusterPeerStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressFamilies** | Pointer to **[]string** | BGP Address families. | [optional] 
**PeerAddress** | Pointer to **string** | BGP peer IP address. | [optional] 
**RemoteAsn** | Pointer to **int64** | Remote AS number. | [optional] 
**State** | Pointer to **string** | BGP session status. | [optional] 

## Methods

### NewClusterPeerStatus

`func NewClusterPeerStatus() *ClusterPeerStatus`

NewClusterPeerStatus instantiates a new ClusterPeerStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterPeerStatusWithDefaults

`func NewClusterPeerStatusWithDefaults() *ClusterPeerStatus`

NewClusterPeerStatusWithDefaults instantiates a new ClusterPeerStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressFamilies

`func (o *ClusterPeerStatus) GetAddressFamilies() []string`

GetAddressFamilies returns the AddressFamilies field if non-nil, zero value otherwise.

### GetAddressFamiliesOk

`func (o *ClusterPeerStatus) GetAddressFamiliesOk() (*[]string, bool)`

GetAddressFamiliesOk returns a tuple with the AddressFamilies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFamilies

`func (o *ClusterPeerStatus) SetAddressFamilies(v []string)`

SetAddressFamilies sets AddressFamilies field to given value.

### HasAddressFamilies

`func (o *ClusterPeerStatus) HasAddressFamilies() bool`

HasAddressFamilies returns a boolean if a field has been set.

### GetPeerAddress

`func (o *ClusterPeerStatus) GetPeerAddress() string`

GetPeerAddress returns the PeerAddress field if non-nil, zero value otherwise.

### GetPeerAddressOk

`func (o *ClusterPeerStatus) GetPeerAddressOk() (*string, bool)`

GetPeerAddressOk returns a tuple with the PeerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerAddress

`func (o *ClusterPeerStatus) SetPeerAddress(v string)`

SetPeerAddress sets PeerAddress field to given value.

### HasPeerAddress

`func (o *ClusterPeerStatus) HasPeerAddress() bool`

HasPeerAddress returns a boolean if a field has been set.

### GetRemoteAsn

`func (o *ClusterPeerStatus) GetRemoteAsn() int64`

GetRemoteAsn returns the RemoteAsn field if non-nil, zero value otherwise.

### GetRemoteAsnOk

`func (o *ClusterPeerStatus) GetRemoteAsnOk() (*int64, bool)`

GetRemoteAsnOk returns a tuple with the RemoteAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAsn

`func (o *ClusterPeerStatus) SetRemoteAsn(v int64)`

SetRemoteAsn sets RemoteAsn field to given value.

### HasRemoteAsn

`func (o *ClusterPeerStatus) HasRemoteAsn() bool`

HasRemoteAsn returns a boolean if a field has been set.

### GetState

`func (o *ClusterPeerStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ClusterPeerStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ClusterPeerStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ClusterPeerStatus) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


