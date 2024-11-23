# ClusterNeighborPortInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MacAddress** | Pointer to **string** | mac address of the local port. | [optional] 
**Name** | Pointer to **string** | local port identifier. | [optional] 
**RemoteMac** | Pointer to **string** | mac address of the remote port. | [optional] 

## Methods

### NewClusterNeighborPortInfo

`func NewClusterNeighborPortInfo() *ClusterNeighborPortInfo`

NewClusterNeighborPortInfo instantiates a new ClusterNeighborPortInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterNeighborPortInfoWithDefaults

`func NewClusterNeighborPortInfoWithDefaults() *ClusterNeighborPortInfo`

NewClusterNeighborPortInfoWithDefaults instantiates a new ClusterNeighborPortInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMacAddress

`func (o *ClusterNeighborPortInfo) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *ClusterNeighborPortInfo) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *ClusterNeighborPortInfo) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *ClusterNeighborPortInfo) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetName

`func (o *ClusterNeighborPortInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterNeighborPortInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterNeighborPortInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterNeighborPortInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRemoteMac

`func (o *ClusterNeighborPortInfo) GetRemoteMac() string`

GetRemoteMac returns the RemoteMac field if non-nil, zero value otherwise.

### GetRemoteMacOk

`func (o *ClusterNeighborPortInfo) GetRemoteMacOk() (*string, bool)`

GetRemoteMacOk returns a tuple with the RemoteMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteMac

`func (o *ClusterNeighborPortInfo) SetRemoteMac(v string)`

SetRemoteMac sets RemoteMac field to given value.

### HasRemoteMac

`func (o *ClusterNeighborPortInfo) HasRemoteMac() bool`

HasRemoteMac returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


