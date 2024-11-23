# ClusterPeer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpConfig** | Pointer to [**ClusterIPConfig**](clusterIPConfig.md) |  | [optional] 
**MacAddress** | Pointer to **string** | MACAddress of the peer DSS. | [optional] 

## Methods

### NewClusterPeer

`func NewClusterPeer() *ClusterPeer`

NewClusterPeer instantiates a new ClusterPeer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterPeerWithDefaults

`func NewClusterPeerWithDefaults() *ClusterPeer`

NewClusterPeerWithDefaults instantiates a new ClusterPeer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpConfig

`func (o *ClusterPeer) GetIpConfig() ClusterIPConfig`

GetIpConfig returns the IpConfig field if non-nil, zero value otherwise.

### GetIpConfigOk

`func (o *ClusterPeer) GetIpConfigOk() (*ClusterIPConfig, bool)`

GetIpConfigOk returns a tuple with the IpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpConfig

`func (o *ClusterPeer) SetIpConfig(v ClusterIPConfig)`

SetIpConfig sets IpConfig field to given value.

### HasIpConfig

`func (o *ClusterPeer) HasIpConfig() bool`

HasIpConfig returns a boolean if a field has been set.

### GetMacAddress

`func (o *ClusterPeer) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *ClusterPeer) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *ClusterPeer) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *ClusterPeer) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


