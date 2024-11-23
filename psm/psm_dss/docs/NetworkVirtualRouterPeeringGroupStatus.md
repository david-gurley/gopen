# NetworkVirtualRouterPeeringGroupStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PropagationStatus** | Pointer to [**SecurityPropagationStatus**](securityPropagationStatus.md) |  | [optional] 
**RouteTables** | Pointer to [**map[string]NetworkVirtualRouterPeeringRouteTable**](networkVirtualRouterPeeringRouteTable.md) | VirtualRouter -&gt; Route table derived from the prefixes exposed by other VirtualRouters in this peering group. | [optional] 

## Methods

### NewNetworkVirtualRouterPeeringGroupStatus

`func NewNetworkVirtualRouterPeeringGroupStatus() *NetworkVirtualRouterPeeringGroupStatus`

NewNetworkVirtualRouterPeeringGroupStatus instantiates a new NetworkVirtualRouterPeeringGroupStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkVirtualRouterPeeringGroupStatusWithDefaults

`func NewNetworkVirtualRouterPeeringGroupStatusWithDefaults() *NetworkVirtualRouterPeeringGroupStatus`

NewNetworkVirtualRouterPeeringGroupStatusWithDefaults instantiates a new NetworkVirtualRouterPeeringGroupStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPropagationStatus

`func (o *NetworkVirtualRouterPeeringGroupStatus) GetPropagationStatus() SecurityPropagationStatus`

GetPropagationStatus returns the PropagationStatus field if non-nil, zero value otherwise.

### GetPropagationStatusOk

`func (o *NetworkVirtualRouterPeeringGroupStatus) GetPropagationStatusOk() (*SecurityPropagationStatus, bool)`

GetPropagationStatusOk returns a tuple with the PropagationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropagationStatus

`func (o *NetworkVirtualRouterPeeringGroupStatus) SetPropagationStatus(v SecurityPropagationStatus)`

SetPropagationStatus sets PropagationStatus field to given value.

### HasPropagationStatus

`func (o *NetworkVirtualRouterPeeringGroupStatus) HasPropagationStatus() bool`

HasPropagationStatus returns a boolean if a field has been set.

### GetRouteTables

`func (o *NetworkVirtualRouterPeeringGroupStatus) GetRouteTables() map[string]NetworkVirtualRouterPeeringRouteTable`

GetRouteTables returns the RouteTables field if non-nil, zero value otherwise.

### GetRouteTablesOk

`func (o *NetworkVirtualRouterPeeringGroupStatus) GetRouteTablesOk() (*map[string]NetworkVirtualRouterPeeringRouteTable, bool)`

GetRouteTablesOk returns a tuple with the RouteTables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteTables

`func (o *NetworkVirtualRouterPeeringGroupStatus) SetRouteTables(v map[string]NetworkVirtualRouterPeeringRouteTable)`

SetRouteTables sets RouteTables field to given value.

### HasRouteTables

`func (o *NetworkVirtualRouterPeeringGroupStatus) HasRouteTables() bool`

HasRouteTables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

