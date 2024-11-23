# NetworkRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextHop** | Pointer to **string** | NextHop for the route. Should be a valid v4 or v6 CIDR block. | [optional] 
**Prefix** | Pointer to **string** | Route Prefix for the route. Should be a valid v4 or v6 CIDR block. | [optional] 
**TargetVirtualRouter** | Pointer to **string** | Target VirtualRouter instance. | [optional] 

## Methods

### NewNetworkRoute

`func NewNetworkRoute() *NetworkRoute`

NewNetworkRoute instantiates a new NetworkRoute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkRouteWithDefaults

`func NewNetworkRouteWithDefaults() *NetworkRoute`

NewNetworkRouteWithDefaults instantiates a new NetworkRoute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextHop

`func (o *NetworkRoute) GetNextHop() string`

GetNextHop returns the NextHop field if non-nil, zero value otherwise.

### GetNextHopOk

`func (o *NetworkRoute) GetNextHopOk() (*string, bool)`

GetNextHopOk returns a tuple with the NextHop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextHop

`func (o *NetworkRoute) SetNextHop(v string)`

SetNextHop sets NextHop field to given value.

### HasNextHop

`func (o *NetworkRoute) HasNextHop() bool`

HasNextHop returns a boolean if a field has been set.

### GetPrefix

`func (o *NetworkRoute) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *NetworkRoute) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *NetworkRoute) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *NetworkRoute) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetTargetVirtualRouter

`func (o *NetworkRoute) GetTargetVirtualRouter() string`

GetTargetVirtualRouter returns the TargetVirtualRouter field if non-nil, zero value otherwise.

### GetTargetVirtualRouterOk

`func (o *NetworkRoute) GetTargetVirtualRouterOk() (*string, bool)`

GetTargetVirtualRouterOk returns a tuple with the TargetVirtualRouter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVirtualRouter

`func (o *NetworkRoute) SetTargetVirtualRouter(v string)`

SetTargetVirtualRouter sets TargetVirtualRouter field to given value.

### HasTargetVirtualRouter

`func (o *NetworkRoute) HasTargetVirtualRouter() bool`

HasTargetVirtualRouter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


