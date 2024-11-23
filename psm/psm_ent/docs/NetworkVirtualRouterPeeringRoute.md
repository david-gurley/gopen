# NetworkVirtualRouterPeeringRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestVirtualRouter** | Pointer to **string** | Destination VirtualRouter this prefix is reachable on. | [optional] 
**Ipv4Prefix** | Pointer to **string** | Should be a valid v4 or v6 CIDR block. | [optional] 

## Methods

### NewNetworkVirtualRouterPeeringRoute

`func NewNetworkVirtualRouterPeeringRoute() *NetworkVirtualRouterPeeringRoute`

NewNetworkVirtualRouterPeeringRoute instantiates a new NetworkVirtualRouterPeeringRoute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkVirtualRouterPeeringRouteWithDefaults

`func NewNetworkVirtualRouterPeeringRouteWithDefaults() *NetworkVirtualRouterPeeringRoute`

NewNetworkVirtualRouterPeeringRouteWithDefaults instantiates a new NetworkVirtualRouterPeeringRoute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestVirtualRouter

`func (o *NetworkVirtualRouterPeeringRoute) GetDestVirtualRouter() string`

GetDestVirtualRouter returns the DestVirtualRouter field if non-nil, zero value otherwise.

### GetDestVirtualRouterOk

`func (o *NetworkVirtualRouterPeeringRoute) GetDestVirtualRouterOk() (*string, bool)`

GetDestVirtualRouterOk returns a tuple with the DestVirtualRouter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestVirtualRouter

`func (o *NetworkVirtualRouterPeeringRoute) SetDestVirtualRouter(v string)`

SetDestVirtualRouter sets DestVirtualRouter field to given value.

### HasDestVirtualRouter

`func (o *NetworkVirtualRouterPeeringRoute) HasDestVirtualRouter() bool`

HasDestVirtualRouter returns a boolean if a field has been set.

### GetIpv4Prefix

`func (o *NetworkVirtualRouterPeeringRoute) GetIpv4Prefix() string`

GetIpv4Prefix returns the Ipv4Prefix field if non-nil, zero value otherwise.

### GetIpv4PrefixOk

`func (o *NetworkVirtualRouterPeeringRoute) GetIpv4PrefixOk() (*string, bool)`

GetIpv4PrefixOk returns a tuple with the Ipv4Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Prefix

`func (o *NetworkVirtualRouterPeeringRoute) SetIpv4Prefix(v string)`

SetIpv4Prefix sets Ipv4Prefix field to given value.

### HasIpv4Prefix

`func (o *NetworkVirtualRouterPeeringRoute) HasIpv4Prefix() bool`

HasIpv4Prefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


