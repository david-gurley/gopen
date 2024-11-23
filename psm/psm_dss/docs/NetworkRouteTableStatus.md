# NetworkRouteTableStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Routes** | Pointer to [**[]NetworkRoute**](NetworkRoute.md) |  | [optional] 

## Methods

### NewNetworkRouteTableStatus

`func NewNetworkRouteTableStatus() *NetworkRouteTableStatus`

NewNetworkRouteTableStatus instantiates a new NetworkRouteTableStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkRouteTableStatusWithDefaults

`func NewNetworkRouteTableStatusWithDefaults() *NetworkRouteTableStatus`

NewNetworkRouteTableStatusWithDefaults instantiates a new NetworkRouteTableStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoutes

`func (o *NetworkRouteTableStatus) GetRoutes() []NetworkRoute`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *NetworkRouteTableStatus) GetRoutesOk() (*[]NetworkRoute, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *NetworkRouteTableStatus) SetRoutes(v []NetworkRoute)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *NetworkRouteTableStatus) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


