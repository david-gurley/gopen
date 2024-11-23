# RoutingRouteStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aspath** | Pointer to **string** |  | [optional] 
**Bestroute** | Pointer to **bool** |  | [optional] 
**Ecmproute** | Pointer to **bool** |  | [optional] 
**Extcomm** | Pointer to **[]string** |  | [optional] 
**FlapStartTime** | Pointer to **int64** |  | [optional] 
**FlapStatsFlapcnt** | Pointer to **int64** |  | [optional] 
**FlapStatsSupprsd** | Pointer to **bool** |  | [optional] 
**Isactive** | Pointer to **string** |  | [optional] 
**Nexthopaddr** | Pointer to **string** |  | [optional] 
**Pathorigid** | Pointer to **string** |  | [optional] 
**Prefix** | Pointer to [**RoutingEVPNRoute**](routingEVPNRoute.md) |  | [optional] 
**Reasonnotbest** | Pointer to **string** |  | [optional] 
**RemoteAddr** | Pointer to **string** |  | [optional] 
**Routesource** | Pointer to **string** |  | [optional] 
**Stale** | Pointer to **bool** |  | [optional] 

## Methods

### NewRoutingRouteStatus

`func NewRoutingRouteStatus() *RoutingRouteStatus`

NewRoutingRouteStatus instantiates a new RoutingRouteStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingRouteStatusWithDefaults

`func NewRoutingRouteStatusWithDefaults() *RoutingRouteStatus`

NewRoutingRouteStatusWithDefaults instantiates a new RoutingRouteStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAspath

`func (o *RoutingRouteStatus) GetAspath() string`

GetAspath returns the Aspath field if non-nil, zero value otherwise.

### GetAspathOk

`func (o *RoutingRouteStatus) GetAspathOk() (*string, bool)`

GetAspathOk returns a tuple with the Aspath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspath

`func (o *RoutingRouteStatus) SetAspath(v string)`

SetAspath sets Aspath field to given value.

### HasAspath

`func (o *RoutingRouteStatus) HasAspath() bool`

HasAspath returns a boolean if a field has been set.

### GetBestroute

`func (o *RoutingRouteStatus) GetBestroute() bool`

GetBestroute returns the Bestroute field if non-nil, zero value otherwise.

### GetBestrouteOk

`func (o *RoutingRouteStatus) GetBestrouteOk() (*bool, bool)`

GetBestrouteOk returns a tuple with the Bestroute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBestroute

`func (o *RoutingRouteStatus) SetBestroute(v bool)`

SetBestroute sets Bestroute field to given value.

### HasBestroute

`func (o *RoutingRouteStatus) HasBestroute() bool`

HasBestroute returns a boolean if a field has been set.

### GetEcmproute

`func (o *RoutingRouteStatus) GetEcmproute() bool`

GetEcmproute returns the Ecmproute field if non-nil, zero value otherwise.

### GetEcmprouteOk

`func (o *RoutingRouteStatus) GetEcmprouteOk() (*bool, bool)`

GetEcmprouteOk returns a tuple with the Ecmproute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcmproute

`func (o *RoutingRouteStatus) SetEcmproute(v bool)`

SetEcmproute sets Ecmproute field to given value.

### HasEcmproute

`func (o *RoutingRouteStatus) HasEcmproute() bool`

HasEcmproute returns a boolean if a field has been set.

### GetExtcomm

`func (o *RoutingRouteStatus) GetExtcomm() []string`

GetExtcomm returns the Extcomm field if non-nil, zero value otherwise.

### GetExtcommOk

`func (o *RoutingRouteStatus) GetExtcommOk() (*[]string, bool)`

GetExtcommOk returns a tuple with the Extcomm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtcomm

`func (o *RoutingRouteStatus) SetExtcomm(v []string)`

SetExtcomm sets Extcomm field to given value.

### HasExtcomm

`func (o *RoutingRouteStatus) HasExtcomm() bool`

HasExtcomm returns a boolean if a field has been set.

### GetFlapStartTime

`func (o *RoutingRouteStatus) GetFlapStartTime() int64`

GetFlapStartTime returns the FlapStartTime field if non-nil, zero value otherwise.

### GetFlapStartTimeOk

`func (o *RoutingRouteStatus) GetFlapStartTimeOk() (*int64, bool)`

GetFlapStartTimeOk returns a tuple with the FlapStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlapStartTime

`func (o *RoutingRouteStatus) SetFlapStartTime(v int64)`

SetFlapStartTime sets FlapStartTime field to given value.

### HasFlapStartTime

`func (o *RoutingRouteStatus) HasFlapStartTime() bool`

HasFlapStartTime returns a boolean if a field has been set.

### GetFlapStatsFlapcnt

`func (o *RoutingRouteStatus) GetFlapStatsFlapcnt() int64`

GetFlapStatsFlapcnt returns the FlapStatsFlapcnt field if non-nil, zero value otherwise.

### GetFlapStatsFlapcntOk

`func (o *RoutingRouteStatus) GetFlapStatsFlapcntOk() (*int64, bool)`

GetFlapStatsFlapcntOk returns a tuple with the FlapStatsFlapcnt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlapStatsFlapcnt

`func (o *RoutingRouteStatus) SetFlapStatsFlapcnt(v int64)`

SetFlapStatsFlapcnt sets FlapStatsFlapcnt field to given value.

### HasFlapStatsFlapcnt

`func (o *RoutingRouteStatus) HasFlapStatsFlapcnt() bool`

HasFlapStatsFlapcnt returns a boolean if a field has been set.

### GetFlapStatsSupprsd

`func (o *RoutingRouteStatus) GetFlapStatsSupprsd() bool`

GetFlapStatsSupprsd returns the FlapStatsSupprsd field if non-nil, zero value otherwise.

### GetFlapStatsSupprsdOk

`func (o *RoutingRouteStatus) GetFlapStatsSupprsdOk() (*bool, bool)`

GetFlapStatsSupprsdOk returns a tuple with the FlapStatsSupprsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlapStatsSupprsd

`func (o *RoutingRouteStatus) SetFlapStatsSupprsd(v bool)`

SetFlapStatsSupprsd sets FlapStatsSupprsd field to given value.

### HasFlapStatsSupprsd

`func (o *RoutingRouteStatus) HasFlapStatsSupprsd() bool`

HasFlapStatsSupprsd returns a boolean if a field has been set.

### GetIsactive

`func (o *RoutingRouteStatus) GetIsactive() string`

GetIsactive returns the Isactive field if non-nil, zero value otherwise.

### GetIsactiveOk

`func (o *RoutingRouteStatus) GetIsactiveOk() (*string, bool)`

GetIsactiveOk returns a tuple with the Isactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsactive

`func (o *RoutingRouteStatus) SetIsactive(v string)`

SetIsactive sets Isactive field to given value.

### HasIsactive

`func (o *RoutingRouteStatus) HasIsactive() bool`

HasIsactive returns a boolean if a field has been set.

### GetNexthopaddr

`func (o *RoutingRouteStatus) GetNexthopaddr() string`

GetNexthopaddr returns the Nexthopaddr field if non-nil, zero value otherwise.

### GetNexthopaddrOk

`func (o *RoutingRouteStatus) GetNexthopaddrOk() (*string, bool)`

GetNexthopaddrOk returns a tuple with the Nexthopaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNexthopaddr

`func (o *RoutingRouteStatus) SetNexthopaddr(v string)`

SetNexthopaddr sets Nexthopaddr field to given value.

### HasNexthopaddr

`func (o *RoutingRouteStatus) HasNexthopaddr() bool`

HasNexthopaddr returns a boolean if a field has been set.

### GetPathorigid

`func (o *RoutingRouteStatus) GetPathorigid() string`

GetPathorigid returns the Pathorigid field if non-nil, zero value otherwise.

### GetPathorigidOk

`func (o *RoutingRouteStatus) GetPathorigidOk() (*string, bool)`

GetPathorigidOk returns a tuple with the Pathorigid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathorigid

`func (o *RoutingRouteStatus) SetPathorigid(v string)`

SetPathorigid sets Pathorigid field to given value.

### HasPathorigid

`func (o *RoutingRouteStatus) HasPathorigid() bool`

HasPathorigid returns a boolean if a field has been set.

### GetPrefix

`func (o *RoutingRouteStatus) GetPrefix() RoutingEVPNRoute`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *RoutingRouteStatus) GetPrefixOk() (*RoutingEVPNRoute, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *RoutingRouteStatus) SetPrefix(v RoutingEVPNRoute)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *RoutingRouteStatus) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetReasonnotbest

`func (o *RoutingRouteStatus) GetReasonnotbest() string`

GetReasonnotbest returns the Reasonnotbest field if non-nil, zero value otherwise.

### GetReasonnotbestOk

`func (o *RoutingRouteStatus) GetReasonnotbestOk() (*string, bool)`

GetReasonnotbestOk returns a tuple with the Reasonnotbest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonnotbest

`func (o *RoutingRouteStatus) SetReasonnotbest(v string)`

SetReasonnotbest sets Reasonnotbest field to given value.

### HasReasonnotbest

`func (o *RoutingRouteStatus) HasReasonnotbest() bool`

HasReasonnotbest returns a boolean if a field has been set.

### GetRemoteAddr

`func (o *RoutingRouteStatus) GetRemoteAddr() string`

GetRemoteAddr returns the RemoteAddr field if non-nil, zero value otherwise.

### GetRemoteAddrOk

`func (o *RoutingRouteStatus) GetRemoteAddrOk() (*string, bool)`

GetRemoteAddrOk returns a tuple with the RemoteAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAddr

`func (o *RoutingRouteStatus) SetRemoteAddr(v string)`

SetRemoteAddr sets RemoteAddr field to given value.

### HasRemoteAddr

`func (o *RoutingRouteStatus) HasRemoteAddr() bool`

HasRemoteAddr returns a boolean if a field has been set.

### GetRoutesource

`func (o *RoutingRouteStatus) GetRoutesource() string`

GetRoutesource returns the Routesource field if non-nil, zero value otherwise.

### GetRoutesourceOk

`func (o *RoutingRouteStatus) GetRoutesourceOk() (*string, bool)`

GetRoutesourceOk returns a tuple with the Routesource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutesource

`func (o *RoutingRouteStatus) SetRoutesource(v string)`

SetRoutesource sets Routesource field to given value.

### HasRoutesource

`func (o *RoutingRouteStatus) HasRoutesource() bool`

HasRoutesource returns a boolean if a field has been set.

### GetStale

`func (o *RoutingRouteStatus) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *RoutingRouteStatus) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *RoutingRouteStatus) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *RoutingRouteStatus) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


