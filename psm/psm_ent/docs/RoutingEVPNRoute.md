# RoutingEVPNRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **int64** |  | [optional] 
**Type2Prefix** | Pointer to [**EVPNRouteEVPNType2Route**](EVPNRouteEVPNType2Route.md) |  | [optional] 
**Type5Prefix** | Pointer to [**EVPNRouteEVPNType5Route**](EVPNRouteEVPNType5Route.md) |  | [optional] 

## Methods

### NewRoutingEVPNRoute

`func NewRoutingEVPNRoute() *RoutingEVPNRoute`

NewRoutingEVPNRoute instantiates a new RoutingEVPNRoute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingEVPNRouteWithDefaults

`func NewRoutingEVPNRouteWithDefaults() *RoutingEVPNRoute`

NewRoutingEVPNRouteWithDefaults instantiates a new RoutingEVPNRoute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RoutingEVPNRoute) GetType() int64`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RoutingEVPNRoute) GetTypeOk() (*int64, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RoutingEVPNRoute) SetType(v int64)`

SetType sets Type field to given value.

### HasType

`func (o *RoutingEVPNRoute) HasType() bool`

HasType returns a boolean if a field has been set.

### GetType2Prefix

`func (o *RoutingEVPNRoute) GetType2Prefix() EVPNRouteEVPNType2Route`

GetType2Prefix returns the Type2Prefix field if non-nil, zero value otherwise.

### GetType2PrefixOk

`func (o *RoutingEVPNRoute) GetType2PrefixOk() (*EVPNRouteEVPNType2Route, bool)`

GetType2PrefixOk returns a tuple with the Type2Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType2Prefix

`func (o *RoutingEVPNRoute) SetType2Prefix(v EVPNRouteEVPNType2Route)`

SetType2Prefix sets Type2Prefix field to given value.

### HasType2Prefix

`func (o *RoutingEVPNRoute) HasType2Prefix() bool`

HasType2Prefix returns a boolean if a field has been set.

### GetType5Prefix

`func (o *RoutingEVPNRoute) GetType5Prefix() EVPNRouteEVPNType5Route`

GetType5Prefix returns the Type5Prefix field if non-nil, zero value otherwise.

### GetType5PrefixOk

`func (o *RoutingEVPNRoute) GetType5PrefixOk() (*EVPNRouteEVPNType5Route, bool)`

GetType5PrefixOk returns a tuple with the Type5Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType5Prefix

`func (o *RoutingEVPNRoute) SetType5Prefix(v EVPNRouteEVPNType5Route)`

SetType5Prefix sets Type5Prefix field to given value.

### HasType5Prefix

`func (o *RoutingEVPNRoute) HasType5Prefix() bool`

HasType5Prefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


