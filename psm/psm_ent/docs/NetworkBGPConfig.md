# NetworkBGPConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsNumber** | Pointer to [**ApiBgpAsn**](apiBgpAsn.md) |  | [optional] 
**DscAutoConfig** | Pointer to **bool** | DSCAutoConfig sets the flag that this config is to be used as a template for auto configuration. | [optional] 
**Holdtime** | Pointer to **int64** | Holdtime is time for which not receiving a keepalive message results in declaring the peer as dead. Value should be between 0 and 3600. | [optional] [default to 180]
**KeepaliveInterval** | Pointer to **int64** | KeepaliveInterval is time interval at which keepalive messages are sent. Value should be between 0 and 3600. | [optional] [default to 60]
**Neighbors** | Pointer to [**[]NetworkBGPNeighbor**](NetworkBGPNeighbor.md) | List of all neighbors. | [optional] 
**RouterId** | Pointer to **string** | Router ID for the BGP Instance. Should be a valid v4 or v6 IP address. | [optional] 
**SuppressDefaultResolution** | Pointer to **bool** | SuppressDefaultResolution excludes default route from being used to resolve nexthop reachability in the underlay. WARNING: modifying this has network-wide data traffic impact as it temporarily deactivates and then re-activates all underlay and overlay routes on every node where this config is applied. | [optional] 

## Methods

### NewNetworkBGPConfig

`func NewNetworkBGPConfig() *NetworkBGPConfig`

NewNetworkBGPConfig instantiates a new NetworkBGPConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkBGPConfigWithDefaults

`func NewNetworkBGPConfigWithDefaults() *NetworkBGPConfig`

NewNetworkBGPConfigWithDefaults instantiates a new NetworkBGPConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsNumber

`func (o *NetworkBGPConfig) GetAsNumber() ApiBgpAsn`

GetAsNumber returns the AsNumber field if non-nil, zero value otherwise.

### GetAsNumberOk

`func (o *NetworkBGPConfig) GetAsNumberOk() (*ApiBgpAsn, bool)`

GetAsNumberOk returns a tuple with the AsNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsNumber

`func (o *NetworkBGPConfig) SetAsNumber(v ApiBgpAsn)`

SetAsNumber sets AsNumber field to given value.

### HasAsNumber

`func (o *NetworkBGPConfig) HasAsNumber() bool`

HasAsNumber returns a boolean if a field has been set.

### GetDscAutoConfig

`func (o *NetworkBGPConfig) GetDscAutoConfig() bool`

GetDscAutoConfig returns the DscAutoConfig field if non-nil, zero value otherwise.

### GetDscAutoConfigOk

`func (o *NetworkBGPConfig) GetDscAutoConfigOk() (*bool, bool)`

GetDscAutoConfigOk returns a tuple with the DscAutoConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDscAutoConfig

`func (o *NetworkBGPConfig) SetDscAutoConfig(v bool)`

SetDscAutoConfig sets DscAutoConfig field to given value.

### HasDscAutoConfig

`func (o *NetworkBGPConfig) HasDscAutoConfig() bool`

HasDscAutoConfig returns a boolean if a field has been set.

### GetHoldtime

`func (o *NetworkBGPConfig) GetHoldtime() int64`

GetHoldtime returns the Holdtime field if non-nil, zero value otherwise.

### GetHoldtimeOk

`func (o *NetworkBGPConfig) GetHoldtimeOk() (*int64, bool)`

GetHoldtimeOk returns a tuple with the Holdtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldtime

`func (o *NetworkBGPConfig) SetHoldtime(v int64)`

SetHoldtime sets Holdtime field to given value.

### HasHoldtime

`func (o *NetworkBGPConfig) HasHoldtime() bool`

HasHoldtime returns a boolean if a field has been set.

### GetKeepaliveInterval

`func (o *NetworkBGPConfig) GetKeepaliveInterval() int64`

GetKeepaliveInterval returns the KeepaliveInterval field if non-nil, zero value otherwise.

### GetKeepaliveIntervalOk

`func (o *NetworkBGPConfig) GetKeepaliveIntervalOk() (*int64, bool)`

GetKeepaliveIntervalOk returns a tuple with the KeepaliveInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepaliveInterval

`func (o *NetworkBGPConfig) SetKeepaliveInterval(v int64)`

SetKeepaliveInterval sets KeepaliveInterval field to given value.

### HasKeepaliveInterval

`func (o *NetworkBGPConfig) HasKeepaliveInterval() bool`

HasKeepaliveInterval returns a boolean if a field has been set.

### GetNeighbors

`func (o *NetworkBGPConfig) GetNeighbors() []NetworkBGPNeighbor`

GetNeighbors returns the Neighbors field if non-nil, zero value otherwise.

### GetNeighborsOk

`func (o *NetworkBGPConfig) GetNeighborsOk() (*[]NetworkBGPNeighbor, bool)`

GetNeighborsOk returns a tuple with the Neighbors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighbors

`func (o *NetworkBGPConfig) SetNeighbors(v []NetworkBGPNeighbor)`

SetNeighbors sets Neighbors field to given value.

### HasNeighbors

`func (o *NetworkBGPConfig) HasNeighbors() bool`

HasNeighbors returns a boolean if a field has been set.

### GetRouterId

`func (o *NetworkBGPConfig) GetRouterId() string`

GetRouterId returns the RouterId field if non-nil, zero value otherwise.

### GetRouterIdOk

`func (o *NetworkBGPConfig) GetRouterIdOk() (*string, bool)`

GetRouterIdOk returns a tuple with the RouterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterId

`func (o *NetworkBGPConfig) SetRouterId(v string)`

SetRouterId sets RouterId field to given value.

### HasRouterId

`func (o *NetworkBGPConfig) HasRouterId() bool`

HasRouterId returns a boolean if a field has been set.

### GetSuppressDefaultResolution

`func (o *NetworkBGPConfig) GetSuppressDefaultResolution() bool`

GetSuppressDefaultResolution returns the SuppressDefaultResolution field if non-nil, zero value otherwise.

### GetSuppressDefaultResolutionOk

`func (o *NetworkBGPConfig) GetSuppressDefaultResolutionOk() (*bool, bool)`

GetSuppressDefaultResolutionOk returns a tuple with the SuppressDefaultResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressDefaultResolution

`func (o *NetworkBGPConfig) SetSuppressDefaultResolution(v bool)`

SetSuppressDefaultResolution sets SuppressDefaultResolution field to given value.

### HasSuppressDefaultResolution

`func (o *NetworkBGPConfig) HasSuppressDefaultResolution() bool`

HasSuppressDefaultResolution returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


