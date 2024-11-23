# RoutingHealthStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalPeers** | Pointer to [**HealthStatusPeeringStatus**](HealthStatusPeeringStatus.md) |  | [optional] 
**InternalPeers** | Pointer to [**HealthStatusPeeringStatus**](HealthStatusPeeringStatus.md) |  | [optional] 
**PendingConfigPeers** | Pointer to **[]string** |  | [optional] 
**RouterId** | Pointer to **string** |  | [optional] 
**UnexpectedPeers** | Pointer to **int32** |  | [optional] 

## Methods

### NewRoutingHealthStatus

`func NewRoutingHealthStatus() *RoutingHealthStatus`

NewRoutingHealthStatus instantiates a new RoutingHealthStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingHealthStatusWithDefaults

`func NewRoutingHealthStatusWithDefaults() *RoutingHealthStatus`

NewRoutingHealthStatusWithDefaults instantiates a new RoutingHealthStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalPeers

`func (o *RoutingHealthStatus) GetExternalPeers() HealthStatusPeeringStatus`

GetExternalPeers returns the ExternalPeers field if non-nil, zero value otherwise.

### GetExternalPeersOk

`func (o *RoutingHealthStatus) GetExternalPeersOk() (*HealthStatusPeeringStatus, bool)`

GetExternalPeersOk returns a tuple with the ExternalPeers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPeers

`func (o *RoutingHealthStatus) SetExternalPeers(v HealthStatusPeeringStatus)`

SetExternalPeers sets ExternalPeers field to given value.

### HasExternalPeers

`func (o *RoutingHealthStatus) HasExternalPeers() bool`

HasExternalPeers returns a boolean if a field has been set.

### GetInternalPeers

`func (o *RoutingHealthStatus) GetInternalPeers() HealthStatusPeeringStatus`

GetInternalPeers returns the InternalPeers field if non-nil, zero value otherwise.

### GetInternalPeersOk

`func (o *RoutingHealthStatus) GetInternalPeersOk() (*HealthStatusPeeringStatus, bool)`

GetInternalPeersOk returns a tuple with the InternalPeers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalPeers

`func (o *RoutingHealthStatus) SetInternalPeers(v HealthStatusPeeringStatus)`

SetInternalPeers sets InternalPeers field to given value.

### HasInternalPeers

`func (o *RoutingHealthStatus) HasInternalPeers() bool`

HasInternalPeers returns a boolean if a field has been set.

### GetPendingConfigPeers

`func (o *RoutingHealthStatus) GetPendingConfigPeers() []string`

GetPendingConfigPeers returns the PendingConfigPeers field if non-nil, zero value otherwise.

### GetPendingConfigPeersOk

`func (o *RoutingHealthStatus) GetPendingConfigPeersOk() (*[]string, bool)`

GetPendingConfigPeersOk returns a tuple with the PendingConfigPeers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingConfigPeers

`func (o *RoutingHealthStatus) SetPendingConfigPeers(v []string)`

SetPendingConfigPeers sets PendingConfigPeers field to given value.

### HasPendingConfigPeers

`func (o *RoutingHealthStatus) HasPendingConfigPeers() bool`

HasPendingConfigPeers returns a boolean if a field has been set.

### GetRouterId

`func (o *RoutingHealthStatus) GetRouterId() string`

GetRouterId returns the RouterId field if non-nil, zero value otherwise.

### GetRouterIdOk

`func (o *RoutingHealthStatus) GetRouterIdOk() (*string, bool)`

GetRouterIdOk returns a tuple with the RouterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterId

`func (o *RoutingHealthStatus) SetRouterId(v string)`

SetRouterId sets RouterId field to given value.

### HasRouterId

`func (o *RoutingHealthStatus) HasRouterId() bool`

HasRouterId returns a boolean if a field has been set.

### GetUnexpectedPeers

`func (o *RoutingHealthStatus) GetUnexpectedPeers() int32`

GetUnexpectedPeers returns the UnexpectedPeers field if non-nil, zero value otherwise.

### GetUnexpectedPeersOk

`func (o *RoutingHealthStatus) GetUnexpectedPeersOk() (*int32, bool)`

GetUnexpectedPeersOk returns a tuple with the UnexpectedPeers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnexpectedPeers

`func (o *RoutingHealthStatus) SetUnexpectedPeers(v int32)`

SetUnexpectedPeers sets UnexpectedPeers field to given value.

### HasUnexpectedPeers

`func (o *RoutingHealthStatus) HasUnexpectedPeers() bool`

HasUnexpectedPeers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


