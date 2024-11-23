# NetworkVirtualRouterPeeringSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ipv4Prefixes** | Pointer to **[]string** | List of destination prefixes located in this Virtual Router exposed as reachable from any other Virtual Router participating in this peering group. Should be a valid v4 or v6 CIDR block. | [optional] 
**VirtualRouter** | Pointer to **string** |  | [optional] 

## Methods

### NewNetworkVirtualRouterPeeringSpec

`func NewNetworkVirtualRouterPeeringSpec() *NetworkVirtualRouterPeeringSpec`

NewNetworkVirtualRouterPeeringSpec instantiates a new NetworkVirtualRouterPeeringSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkVirtualRouterPeeringSpecWithDefaults

`func NewNetworkVirtualRouterPeeringSpecWithDefaults() *NetworkVirtualRouterPeeringSpec`

NewNetworkVirtualRouterPeeringSpecWithDefaults instantiates a new NetworkVirtualRouterPeeringSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpv4Prefixes

`func (o *NetworkVirtualRouterPeeringSpec) GetIpv4Prefixes() []string`

GetIpv4Prefixes returns the Ipv4Prefixes field if non-nil, zero value otherwise.

### GetIpv4PrefixesOk

`func (o *NetworkVirtualRouterPeeringSpec) GetIpv4PrefixesOk() (*[]string, bool)`

GetIpv4PrefixesOk returns a tuple with the Ipv4Prefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Prefixes

`func (o *NetworkVirtualRouterPeeringSpec) SetIpv4Prefixes(v []string)`

SetIpv4Prefixes sets Ipv4Prefixes field to given value.

### HasIpv4Prefixes

`func (o *NetworkVirtualRouterPeeringSpec) HasIpv4Prefixes() bool`

HasIpv4Prefixes returns a boolean if a field has been set.

### GetVirtualRouter

`func (o *NetworkVirtualRouterPeeringSpec) GetVirtualRouter() string`

GetVirtualRouter returns the VirtualRouter field if non-nil, zero value otherwise.

### GetVirtualRouterOk

`func (o *NetworkVirtualRouterPeeringSpec) GetVirtualRouterOk() (*string, bool)`

GetVirtualRouterOk returns a tuple with the VirtualRouter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualRouter

`func (o *NetworkVirtualRouterPeeringSpec) SetVirtualRouter(v string)`

SetVirtualRouter sets VirtualRouter field to given value.

### HasVirtualRouter

`func (o *NetworkVirtualRouterPeeringSpec) HasVirtualRouter() bool`

HasVirtualRouter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


