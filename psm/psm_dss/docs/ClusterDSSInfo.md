# ClusterDSSInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dsms** | Pointer to [**[]ClusterDSM**](ClusterDSM.md) | Distributed service module information. | [optional] 
**FaultInfo** | Pointer to [**ClusterFault**](clusterFault.md) |  | [optional] 
**HaPeer** | Pointer to [**ClusterPeer**](clusterPeer.md) |  | [optional] 
**HostName** | Pointer to **string** | Hostname of the switch. | [optional] 
**LinkInfo** | Pointer to [**[]ClusterNeighborPortInfo**](ClusterNeighborPortInfo.md) | Information of the remote port mac amd local port of a link. | [optional] 
**Version** | Pointer to **string** | switch software version. | [optional] 

## Methods

### NewClusterDSSInfo

`func NewClusterDSSInfo() *ClusterDSSInfo`

NewClusterDSSInfo instantiates a new ClusterDSSInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterDSSInfoWithDefaults

`func NewClusterDSSInfoWithDefaults() *ClusterDSSInfo`

NewClusterDSSInfoWithDefaults instantiates a new ClusterDSSInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDsms

`func (o *ClusterDSSInfo) GetDsms() []ClusterDSM`

GetDsms returns the Dsms field if non-nil, zero value otherwise.

### GetDsmsOk

`func (o *ClusterDSSInfo) GetDsmsOk() (*[]ClusterDSM, bool)`

GetDsmsOk returns a tuple with the Dsms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsms

`func (o *ClusterDSSInfo) SetDsms(v []ClusterDSM)`

SetDsms sets Dsms field to given value.

### HasDsms

`func (o *ClusterDSSInfo) HasDsms() bool`

HasDsms returns a boolean if a field has been set.

### GetFaultInfo

`func (o *ClusterDSSInfo) GetFaultInfo() ClusterFault`

GetFaultInfo returns the FaultInfo field if non-nil, zero value otherwise.

### GetFaultInfoOk

`func (o *ClusterDSSInfo) GetFaultInfoOk() (*ClusterFault, bool)`

GetFaultInfoOk returns a tuple with the FaultInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaultInfo

`func (o *ClusterDSSInfo) SetFaultInfo(v ClusterFault)`

SetFaultInfo sets FaultInfo field to given value.

### HasFaultInfo

`func (o *ClusterDSSInfo) HasFaultInfo() bool`

HasFaultInfo returns a boolean if a field has been set.

### GetHaPeer

`func (o *ClusterDSSInfo) GetHaPeer() ClusterPeer`

GetHaPeer returns the HaPeer field if non-nil, zero value otherwise.

### GetHaPeerOk

`func (o *ClusterDSSInfo) GetHaPeerOk() (*ClusterPeer, bool)`

GetHaPeerOk returns a tuple with the HaPeer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaPeer

`func (o *ClusterDSSInfo) SetHaPeer(v ClusterPeer)`

SetHaPeer sets HaPeer field to given value.

### HasHaPeer

`func (o *ClusterDSSInfo) HasHaPeer() bool`

HasHaPeer returns a boolean if a field has been set.

### GetHostName

`func (o *ClusterDSSInfo) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *ClusterDSSInfo) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *ClusterDSSInfo) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *ClusterDSSInfo) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetLinkInfo

`func (o *ClusterDSSInfo) GetLinkInfo() []ClusterNeighborPortInfo`

GetLinkInfo returns the LinkInfo field if non-nil, zero value otherwise.

### GetLinkInfoOk

`func (o *ClusterDSSInfo) GetLinkInfoOk() (*[]ClusterNeighborPortInfo, bool)`

GetLinkInfoOk returns a tuple with the LinkInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkInfo

`func (o *ClusterDSSInfo) SetLinkInfo(v []ClusterNeighborPortInfo)`

SetLinkInfo sets LinkInfo field to given value.

### HasLinkInfo

`func (o *ClusterDSSInfo) HasLinkInfo() bool`

HasLinkInfo returns a boolean if a field has been set.

### GetVersion

`func (o *ClusterDSSInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ClusterDSSInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ClusterDSSInfo) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ClusterDSSInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


