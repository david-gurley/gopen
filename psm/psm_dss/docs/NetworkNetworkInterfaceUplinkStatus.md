# NetworkNetworkInterfaceUplinkStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpConfig** | Pointer to [**ClusterIPConfig**](clusterIPConfig.md) |  | [optional] 
**LinkSpeed** | Pointer to **string** | LinkSpeed auto-negotiated. | [optional] 
**LldpNeighbor** | Pointer to [**NetworkLLDPNeighbor**](networkLLDPNeighbor.md) |  | [optional] 
**TransceiverStatus** | Pointer to [**NetworkTransceiverStatus**](networkTransceiverStatus.md) |  | [optional] 

## Methods

### NewNetworkNetworkInterfaceUplinkStatus

`func NewNetworkNetworkInterfaceUplinkStatus() *NetworkNetworkInterfaceUplinkStatus`

NewNetworkNetworkInterfaceUplinkStatus instantiates a new NetworkNetworkInterfaceUplinkStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkNetworkInterfaceUplinkStatusWithDefaults

`func NewNetworkNetworkInterfaceUplinkStatusWithDefaults() *NetworkNetworkInterfaceUplinkStatus`

NewNetworkNetworkInterfaceUplinkStatusWithDefaults instantiates a new NetworkNetworkInterfaceUplinkStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpConfig

`func (o *NetworkNetworkInterfaceUplinkStatus) GetIpConfig() ClusterIPConfig`

GetIpConfig returns the IpConfig field if non-nil, zero value otherwise.

### GetIpConfigOk

`func (o *NetworkNetworkInterfaceUplinkStatus) GetIpConfigOk() (*ClusterIPConfig, bool)`

GetIpConfigOk returns a tuple with the IpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpConfig

`func (o *NetworkNetworkInterfaceUplinkStatus) SetIpConfig(v ClusterIPConfig)`

SetIpConfig sets IpConfig field to given value.

### HasIpConfig

`func (o *NetworkNetworkInterfaceUplinkStatus) HasIpConfig() bool`

HasIpConfig returns a boolean if a field has been set.

### GetLinkSpeed

`func (o *NetworkNetworkInterfaceUplinkStatus) GetLinkSpeed() string`

GetLinkSpeed returns the LinkSpeed field if non-nil, zero value otherwise.

### GetLinkSpeedOk

`func (o *NetworkNetworkInterfaceUplinkStatus) GetLinkSpeedOk() (*string, bool)`

GetLinkSpeedOk returns a tuple with the LinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeed

`func (o *NetworkNetworkInterfaceUplinkStatus) SetLinkSpeed(v string)`

SetLinkSpeed sets LinkSpeed field to given value.

### HasLinkSpeed

`func (o *NetworkNetworkInterfaceUplinkStatus) HasLinkSpeed() bool`

HasLinkSpeed returns a boolean if a field has been set.

### GetLldpNeighbor

`func (o *NetworkNetworkInterfaceUplinkStatus) GetLldpNeighbor() NetworkLLDPNeighbor`

GetLldpNeighbor returns the LldpNeighbor field if non-nil, zero value otherwise.

### GetLldpNeighborOk

`func (o *NetworkNetworkInterfaceUplinkStatus) GetLldpNeighborOk() (*NetworkLLDPNeighbor, bool)`

GetLldpNeighborOk returns a tuple with the LldpNeighbor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldpNeighbor

`func (o *NetworkNetworkInterfaceUplinkStatus) SetLldpNeighbor(v NetworkLLDPNeighbor)`

SetLldpNeighbor sets LldpNeighbor field to given value.

### HasLldpNeighbor

`func (o *NetworkNetworkInterfaceUplinkStatus) HasLldpNeighbor() bool`

HasLldpNeighbor returns a boolean if a field has been set.

### GetTransceiverStatus

`func (o *NetworkNetworkInterfaceUplinkStatus) GetTransceiverStatus() NetworkTransceiverStatus`

GetTransceiverStatus returns the TransceiverStatus field if non-nil, zero value otherwise.

### GetTransceiverStatusOk

`func (o *NetworkNetworkInterfaceUplinkStatus) GetTransceiverStatusOk() (*NetworkTransceiverStatus, bool)`

GetTransceiverStatusOk returns a tuple with the TransceiverStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransceiverStatus

`func (o *NetworkNetworkInterfaceUplinkStatus) SetTransceiverStatus(v NetworkTransceiverStatus)`

SetTransceiverStatus sets TransceiverStatus field to given value.

### HasTransceiverStatus

`func (o *NetworkNetworkInterfaceUplinkStatus) HasTransceiverStatus() bool`

HasTransceiverStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


