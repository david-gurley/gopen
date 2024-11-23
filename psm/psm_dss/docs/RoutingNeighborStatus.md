# RoutingNeighborStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capsneg** | Pointer to **string** |  | [optional] 
**Capsrcvd** | Pointer to **string** |  | [optional] 
**Capssent** | Pointer to **string** |  | [optional] 
**Connectretrycount** | Pointer to **int64** |  | [optional] 
**Fsmestablishedtime** | Pointer to **string** |  | [optional] 
**Fsmesttransitions** | Pointer to **int64** |  | [optional] 
**Holdtime** | Pointer to **int64** |  | [optional] 
**Inkeepalives** | Pointer to **int64** |  | [optional] 
**Innotifications** | Pointer to **int64** |  | [optional] 
**Inopens** | Pointer to **int64** |  | [optional] 
**Inprfxes** | Pointer to **int64** |  | [optional] 
**Inprfxesexpwdr** | Pointer to **int64** |  | [optional] 
**Inprfxesimpwdr** | Pointer to **int64** |  | [optional] 
**Inrefreshes** | Pointer to **int64** |  | [optional] 
**Intotalmessages** | Pointer to **int64** |  | [optional] 
**Inupdates** | Pointer to **int64** |  | [optional] 
**Inupdateselpstime** | Pointer to **int64** |  | [optional] 
**Keepalive** | Pointer to **int64** |  | [optional] 
**Lasterrorrcvd** | Pointer to **string** |  | [optional] 
**Lasterrorsent** | Pointer to **string** |  | [optional] 
**Localaddr** | Pointer to **string** | Should be a valid v4 or v6 IP address. | [optional] 
**Orfentrycount** | Pointer to **int64** |  | [optional] 
**Outkeepalives** | Pointer to **int64** |  | [optional] 
**Outnotifications** | Pointer to **int64** |  | [optional] 
**Outopens** | Pointer to **int64** |  | [optional] 
**Outprfxes** | Pointer to **int64** |  | [optional] 
**Outprfxesdenied** | Pointer to **int64** |  | [optional] 
**Outprfxesexpwdr** | Pointer to **int64** |  | [optional] 
**Outprfxesimpwdr** | Pointer to **int64** |  | [optional] 
**Outrefreshes** | Pointer to **int64** |  | [optional] 
**Outtotalmessages** | Pointer to **int64** |  | [optional] 
**Outupdateelpstime** | Pointer to **int64** |  | [optional] 
**Outupdates** | Pointer to **int64** |  | [optional] 
**Peergr** | Pointer to **int64** |  | [optional] 
**Peerindex** | Pointer to **int64** |  | [optional] 
**PrevStatus** | Pointer to **string** |  | [optional] [default to "idle"]
**Rcvdmsgelpstime** | Pointer to **int64** |  | [optional] 
**Receivedholdtime** | Pointer to **int64** |  | [optional] 
**Routerefrrcvd** | Pointer to **int64** |  | [optional] 
**Routerefrsent** | Pointer to **int64** |  | [optional] 
**Sellocaladdrtype** | Pointer to **string** |  | [optional] [default to "other"]
**Stalepathtime** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] [default to "idle"]

## Methods

### NewRoutingNeighborStatus

`func NewRoutingNeighborStatus() *RoutingNeighborStatus`

NewRoutingNeighborStatus instantiates a new RoutingNeighborStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingNeighborStatusWithDefaults

`func NewRoutingNeighborStatusWithDefaults() *RoutingNeighborStatus`

NewRoutingNeighborStatusWithDefaults instantiates a new RoutingNeighborStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapsneg

`func (o *RoutingNeighborStatus) GetCapsneg() string`

GetCapsneg returns the Capsneg field if non-nil, zero value otherwise.

### GetCapsnegOk

`func (o *RoutingNeighborStatus) GetCapsnegOk() (*string, bool)`

GetCapsnegOk returns a tuple with the Capsneg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapsneg

`func (o *RoutingNeighborStatus) SetCapsneg(v string)`

SetCapsneg sets Capsneg field to given value.

### HasCapsneg

`func (o *RoutingNeighborStatus) HasCapsneg() bool`

HasCapsneg returns a boolean if a field has been set.

### GetCapsrcvd

`func (o *RoutingNeighborStatus) GetCapsrcvd() string`

GetCapsrcvd returns the Capsrcvd field if non-nil, zero value otherwise.

### GetCapsrcvdOk

`func (o *RoutingNeighborStatus) GetCapsrcvdOk() (*string, bool)`

GetCapsrcvdOk returns a tuple with the Capsrcvd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapsrcvd

`func (o *RoutingNeighborStatus) SetCapsrcvd(v string)`

SetCapsrcvd sets Capsrcvd field to given value.

### HasCapsrcvd

`func (o *RoutingNeighborStatus) HasCapsrcvd() bool`

HasCapsrcvd returns a boolean if a field has been set.

### GetCapssent

`func (o *RoutingNeighborStatus) GetCapssent() string`

GetCapssent returns the Capssent field if non-nil, zero value otherwise.

### GetCapssentOk

`func (o *RoutingNeighborStatus) GetCapssentOk() (*string, bool)`

GetCapssentOk returns a tuple with the Capssent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapssent

`func (o *RoutingNeighborStatus) SetCapssent(v string)`

SetCapssent sets Capssent field to given value.

### HasCapssent

`func (o *RoutingNeighborStatus) HasCapssent() bool`

HasCapssent returns a boolean if a field has been set.

### GetConnectretrycount

`func (o *RoutingNeighborStatus) GetConnectretrycount() int64`

GetConnectretrycount returns the Connectretrycount field if non-nil, zero value otherwise.

### GetConnectretrycountOk

`func (o *RoutingNeighborStatus) GetConnectretrycountOk() (*int64, bool)`

GetConnectretrycountOk returns a tuple with the Connectretrycount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectretrycount

`func (o *RoutingNeighborStatus) SetConnectretrycount(v int64)`

SetConnectretrycount sets Connectretrycount field to given value.

### HasConnectretrycount

`func (o *RoutingNeighborStatus) HasConnectretrycount() bool`

HasConnectretrycount returns a boolean if a field has been set.

### GetFsmestablishedtime

`func (o *RoutingNeighborStatus) GetFsmestablishedtime() string`

GetFsmestablishedtime returns the Fsmestablishedtime field if non-nil, zero value otherwise.

### GetFsmestablishedtimeOk

`func (o *RoutingNeighborStatus) GetFsmestablishedtimeOk() (*string, bool)`

GetFsmestablishedtimeOk returns a tuple with the Fsmestablishedtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsmestablishedtime

`func (o *RoutingNeighborStatus) SetFsmestablishedtime(v string)`

SetFsmestablishedtime sets Fsmestablishedtime field to given value.

### HasFsmestablishedtime

`func (o *RoutingNeighborStatus) HasFsmestablishedtime() bool`

HasFsmestablishedtime returns a boolean if a field has been set.

### GetFsmesttransitions

`func (o *RoutingNeighborStatus) GetFsmesttransitions() int64`

GetFsmesttransitions returns the Fsmesttransitions field if non-nil, zero value otherwise.

### GetFsmesttransitionsOk

`func (o *RoutingNeighborStatus) GetFsmesttransitionsOk() (*int64, bool)`

GetFsmesttransitionsOk returns a tuple with the Fsmesttransitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsmesttransitions

`func (o *RoutingNeighborStatus) SetFsmesttransitions(v int64)`

SetFsmesttransitions sets Fsmesttransitions field to given value.

### HasFsmesttransitions

`func (o *RoutingNeighborStatus) HasFsmesttransitions() bool`

HasFsmesttransitions returns a boolean if a field has been set.

### GetHoldtime

`func (o *RoutingNeighborStatus) GetHoldtime() int64`

GetHoldtime returns the Holdtime field if non-nil, zero value otherwise.

### GetHoldtimeOk

`func (o *RoutingNeighborStatus) GetHoldtimeOk() (*int64, bool)`

GetHoldtimeOk returns a tuple with the Holdtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldtime

`func (o *RoutingNeighborStatus) SetHoldtime(v int64)`

SetHoldtime sets Holdtime field to given value.

### HasHoldtime

`func (o *RoutingNeighborStatus) HasHoldtime() bool`

HasHoldtime returns a boolean if a field has been set.

### GetInkeepalives

`func (o *RoutingNeighborStatus) GetInkeepalives() int64`

GetInkeepalives returns the Inkeepalives field if non-nil, zero value otherwise.

### GetInkeepalivesOk

`func (o *RoutingNeighborStatus) GetInkeepalivesOk() (*int64, bool)`

GetInkeepalivesOk returns a tuple with the Inkeepalives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInkeepalives

`func (o *RoutingNeighborStatus) SetInkeepalives(v int64)`

SetInkeepalives sets Inkeepalives field to given value.

### HasInkeepalives

`func (o *RoutingNeighborStatus) HasInkeepalives() bool`

HasInkeepalives returns a boolean if a field has been set.

### GetInnotifications

`func (o *RoutingNeighborStatus) GetInnotifications() int64`

GetInnotifications returns the Innotifications field if non-nil, zero value otherwise.

### GetInnotificationsOk

`func (o *RoutingNeighborStatus) GetInnotificationsOk() (*int64, bool)`

GetInnotificationsOk returns a tuple with the Innotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnotifications

`func (o *RoutingNeighborStatus) SetInnotifications(v int64)`

SetInnotifications sets Innotifications field to given value.

### HasInnotifications

`func (o *RoutingNeighborStatus) HasInnotifications() bool`

HasInnotifications returns a boolean if a field has been set.

### GetInopens

`func (o *RoutingNeighborStatus) GetInopens() int64`

GetInopens returns the Inopens field if non-nil, zero value otherwise.

### GetInopensOk

`func (o *RoutingNeighborStatus) GetInopensOk() (*int64, bool)`

GetInopensOk returns a tuple with the Inopens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInopens

`func (o *RoutingNeighborStatus) SetInopens(v int64)`

SetInopens sets Inopens field to given value.

### HasInopens

`func (o *RoutingNeighborStatus) HasInopens() bool`

HasInopens returns a boolean if a field has been set.

### GetInprfxes

`func (o *RoutingNeighborStatus) GetInprfxes() int64`

GetInprfxes returns the Inprfxes field if non-nil, zero value otherwise.

### GetInprfxesOk

`func (o *RoutingNeighborStatus) GetInprfxesOk() (*int64, bool)`

GetInprfxesOk returns a tuple with the Inprfxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInprfxes

`func (o *RoutingNeighborStatus) SetInprfxes(v int64)`

SetInprfxes sets Inprfxes field to given value.

### HasInprfxes

`func (o *RoutingNeighborStatus) HasInprfxes() bool`

HasInprfxes returns a boolean if a field has been set.

### GetInprfxesexpwdr

`func (o *RoutingNeighborStatus) GetInprfxesexpwdr() int64`

GetInprfxesexpwdr returns the Inprfxesexpwdr field if non-nil, zero value otherwise.

### GetInprfxesexpwdrOk

`func (o *RoutingNeighborStatus) GetInprfxesexpwdrOk() (*int64, bool)`

GetInprfxesexpwdrOk returns a tuple with the Inprfxesexpwdr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInprfxesexpwdr

`func (o *RoutingNeighborStatus) SetInprfxesexpwdr(v int64)`

SetInprfxesexpwdr sets Inprfxesexpwdr field to given value.

### HasInprfxesexpwdr

`func (o *RoutingNeighborStatus) HasInprfxesexpwdr() bool`

HasInprfxesexpwdr returns a boolean if a field has been set.

### GetInprfxesimpwdr

`func (o *RoutingNeighborStatus) GetInprfxesimpwdr() int64`

GetInprfxesimpwdr returns the Inprfxesimpwdr field if non-nil, zero value otherwise.

### GetInprfxesimpwdrOk

`func (o *RoutingNeighborStatus) GetInprfxesimpwdrOk() (*int64, bool)`

GetInprfxesimpwdrOk returns a tuple with the Inprfxesimpwdr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInprfxesimpwdr

`func (o *RoutingNeighborStatus) SetInprfxesimpwdr(v int64)`

SetInprfxesimpwdr sets Inprfxesimpwdr field to given value.

### HasInprfxesimpwdr

`func (o *RoutingNeighborStatus) HasInprfxesimpwdr() bool`

HasInprfxesimpwdr returns a boolean if a field has been set.

### GetInrefreshes

`func (o *RoutingNeighborStatus) GetInrefreshes() int64`

GetInrefreshes returns the Inrefreshes field if non-nil, zero value otherwise.

### GetInrefreshesOk

`func (o *RoutingNeighborStatus) GetInrefreshesOk() (*int64, bool)`

GetInrefreshesOk returns a tuple with the Inrefreshes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInrefreshes

`func (o *RoutingNeighborStatus) SetInrefreshes(v int64)`

SetInrefreshes sets Inrefreshes field to given value.

### HasInrefreshes

`func (o *RoutingNeighborStatus) HasInrefreshes() bool`

HasInrefreshes returns a boolean if a field has been set.

### GetIntotalmessages

`func (o *RoutingNeighborStatus) GetIntotalmessages() int64`

GetIntotalmessages returns the Intotalmessages field if non-nil, zero value otherwise.

### GetIntotalmessagesOk

`func (o *RoutingNeighborStatus) GetIntotalmessagesOk() (*int64, bool)`

GetIntotalmessagesOk returns a tuple with the Intotalmessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntotalmessages

`func (o *RoutingNeighborStatus) SetIntotalmessages(v int64)`

SetIntotalmessages sets Intotalmessages field to given value.

### HasIntotalmessages

`func (o *RoutingNeighborStatus) HasIntotalmessages() bool`

HasIntotalmessages returns a boolean if a field has been set.

### GetInupdates

`func (o *RoutingNeighborStatus) GetInupdates() int64`

GetInupdates returns the Inupdates field if non-nil, zero value otherwise.

### GetInupdatesOk

`func (o *RoutingNeighborStatus) GetInupdatesOk() (*int64, bool)`

GetInupdatesOk returns a tuple with the Inupdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInupdates

`func (o *RoutingNeighborStatus) SetInupdates(v int64)`

SetInupdates sets Inupdates field to given value.

### HasInupdates

`func (o *RoutingNeighborStatus) HasInupdates() bool`

HasInupdates returns a boolean if a field has been set.

### GetInupdateselpstime

`func (o *RoutingNeighborStatus) GetInupdateselpstime() int64`

GetInupdateselpstime returns the Inupdateselpstime field if non-nil, zero value otherwise.

### GetInupdateselpstimeOk

`func (o *RoutingNeighborStatus) GetInupdateselpstimeOk() (*int64, bool)`

GetInupdateselpstimeOk returns a tuple with the Inupdateselpstime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInupdateselpstime

`func (o *RoutingNeighborStatus) SetInupdateselpstime(v int64)`

SetInupdateselpstime sets Inupdateselpstime field to given value.

### HasInupdateselpstime

`func (o *RoutingNeighborStatus) HasInupdateselpstime() bool`

HasInupdateselpstime returns a boolean if a field has been set.

### GetKeepalive

`func (o *RoutingNeighborStatus) GetKeepalive() int64`

GetKeepalive returns the Keepalive field if non-nil, zero value otherwise.

### GetKeepaliveOk

`func (o *RoutingNeighborStatus) GetKeepaliveOk() (*int64, bool)`

GetKeepaliveOk returns a tuple with the Keepalive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepalive

`func (o *RoutingNeighborStatus) SetKeepalive(v int64)`

SetKeepalive sets Keepalive field to given value.

### HasKeepalive

`func (o *RoutingNeighborStatus) HasKeepalive() bool`

HasKeepalive returns a boolean if a field has been set.

### GetLasterrorrcvd

`func (o *RoutingNeighborStatus) GetLasterrorrcvd() string`

GetLasterrorrcvd returns the Lasterrorrcvd field if non-nil, zero value otherwise.

### GetLasterrorrcvdOk

`func (o *RoutingNeighborStatus) GetLasterrorrcvdOk() (*string, bool)`

GetLasterrorrcvdOk returns a tuple with the Lasterrorrcvd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLasterrorrcvd

`func (o *RoutingNeighborStatus) SetLasterrorrcvd(v string)`

SetLasterrorrcvd sets Lasterrorrcvd field to given value.

### HasLasterrorrcvd

`func (o *RoutingNeighborStatus) HasLasterrorrcvd() bool`

HasLasterrorrcvd returns a boolean if a field has been set.

### GetLasterrorsent

`func (o *RoutingNeighborStatus) GetLasterrorsent() string`

GetLasterrorsent returns the Lasterrorsent field if non-nil, zero value otherwise.

### GetLasterrorsentOk

`func (o *RoutingNeighborStatus) GetLasterrorsentOk() (*string, bool)`

GetLasterrorsentOk returns a tuple with the Lasterrorsent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLasterrorsent

`func (o *RoutingNeighborStatus) SetLasterrorsent(v string)`

SetLasterrorsent sets Lasterrorsent field to given value.

### HasLasterrorsent

`func (o *RoutingNeighborStatus) HasLasterrorsent() bool`

HasLasterrorsent returns a boolean if a field has been set.

### GetLocaladdr

`func (o *RoutingNeighborStatus) GetLocaladdr() string`

GetLocaladdr returns the Localaddr field if non-nil, zero value otherwise.

### GetLocaladdrOk

`func (o *RoutingNeighborStatus) GetLocaladdrOk() (*string, bool)`

GetLocaladdrOk returns a tuple with the Localaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocaladdr

`func (o *RoutingNeighborStatus) SetLocaladdr(v string)`

SetLocaladdr sets Localaddr field to given value.

### HasLocaladdr

`func (o *RoutingNeighborStatus) HasLocaladdr() bool`

HasLocaladdr returns a boolean if a field has been set.

### GetOrfentrycount

`func (o *RoutingNeighborStatus) GetOrfentrycount() int64`

GetOrfentrycount returns the Orfentrycount field if non-nil, zero value otherwise.

### GetOrfentrycountOk

`func (o *RoutingNeighborStatus) GetOrfentrycountOk() (*int64, bool)`

GetOrfentrycountOk returns a tuple with the Orfentrycount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrfentrycount

`func (o *RoutingNeighborStatus) SetOrfentrycount(v int64)`

SetOrfentrycount sets Orfentrycount field to given value.

### HasOrfentrycount

`func (o *RoutingNeighborStatus) HasOrfentrycount() bool`

HasOrfentrycount returns a boolean if a field has been set.

### GetOutkeepalives

`func (o *RoutingNeighborStatus) GetOutkeepalives() int64`

GetOutkeepalives returns the Outkeepalives field if non-nil, zero value otherwise.

### GetOutkeepalivesOk

`func (o *RoutingNeighborStatus) GetOutkeepalivesOk() (*int64, bool)`

GetOutkeepalivesOk returns a tuple with the Outkeepalives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutkeepalives

`func (o *RoutingNeighborStatus) SetOutkeepalives(v int64)`

SetOutkeepalives sets Outkeepalives field to given value.

### HasOutkeepalives

`func (o *RoutingNeighborStatus) HasOutkeepalives() bool`

HasOutkeepalives returns a boolean if a field has been set.

### GetOutnotifications

`func (o *RoutingNeighborStatus) GetOutnotifications() int64`

GetOutnotifications returns the Outnotifications field if non-nil, zero value otherwise.

### GetOutnotificationsOk

`func (o *RoutingNeighborStatus) GetOutnotificationsOk() (*int64, bool)`

GetOutnotificationsOk returns a tuple with the Outnotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutnotifications

`func (o *RoutingNeighborStatus) SetOutnotifications(v int64)`

SetOutnotifications sets Outnotifications field to given value.

### HasOutnotifications

`func (o *RoutingNeighborStatus) HasOutnotifications() bool`

HasOutnotifications returns a boolean if a field has been set.

### GetOutopens

`func (o *RoutingNeighborStatus) GetOutopens() int64`

GetOutopens returns the Outopens field if non-nil, zero value otherwise.

### GetOutopensOk

`func (o *RoutingNeighborStatus) GetOutopensOk() (*int64, bool)`

GetOutopensOk returns a tuple with the Outopens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutopens

`func (o *RoutingNeighborStatus) SetOutopens(v int64)`

SetOutopens sets Outopens field to given value.

### HasOutopens

`func (o *RoutingNeighborStatus) HasOutopens() bool`

HasOutopens returns a boolean if a field has been set.

### GetOutprfxes

`func (o *RoutingNeighborStatus) GetOutprfxes() int64`

GetOutprfxes returns the Outprfxes field if non-nil, zero value otherwise.

### GetOutprfxesOk

`func (o *RoutingNeighborStatus) GetOutprfxesOk() (*int64, bool)`

GetOutprfxesOk returns a tuple with the Outprfxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutprfxes

`func (o *RoutingNeighborStatus) SetOutprfxes(v int64)`

SetOutprfxes sets Outprfxes field to given value.

### HasOutprfxes

`func (o *RoutingNeighborStatus) HasOutprfxes() bool`

HasOutprfxes returns a boolean if a field has been set.

### GetOutprfxesdenied

`func (o *RoutingNeighborStatus) GetOutprfxesdenied() int64`

GetOutprfxesdenied returns the Outprfxesdenied field if non-nil, zero value otherwise.

### GetOutprfxesdeniedOk

`func (o *RoutingNeighborStatus) GetOutprfxesdeniedOk() (*int64, bool)`

GetOutprfxesdeniedOk returns a tuple with the Outprfxesdenied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutprfxesdenied

`func (o *RoutingNeighborStatus) SetOutprfxesdenied(v int64)`

SetOutprfxesdenied sets Outprfxesdenied field to given value.

### HasOutprfxesdenied

`func (o *RoutingNeighborStatus) HasOutprfxesdenied() bool`

HasOutprfxesdenied returns a boolean if a field has been set.

### GetOutprfxesexpwdr

`func (o *RoutingNeighborStatus) GetOutprfxesexpwdr() int64`

GetOutprfxesexpwdr returns the Outprfxesexpwdr field if non-nil, zero value otherwise.

### GetOutprfxesexpwdrOk

`func (o *RoutingNeighborStatus) GetOutprfxesexpwdrOk() (*int64, bool)`

GetOutprfxesexpwdrOk returns a tuple with the Outprfxesexpwdr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutprfxesexpwdr

`func (o *RoutingNeighborStatus) SetOutprfxesexpwdr(v int64)`

SetOutprfxesexpwdr sets Outprfxesexpwdr field to given value.

### HasOutprfxesexpwdr

`func (o *RoutingNeighborStatus) HasOutprfxesexpwdr() bool`

HasOutprfxesexpwdr returns a boolean if a field has been set.

### GetOutprfxesimpwdr

`func (o *RoutingNeighborStatus) GetOutprfxesimpwdr() int64`

GetOutprfxesimpwdr returns the Outprfxesimpwdr field if non-nil, zero value otherwise.

### GetOutprfxesimpwdrOk

`func (o *RoutingNeighborStatus) GetOutprfxesimpwdrOk() (*int64, bool)`

GetOutprfxesimpwdrOk returns a tuple with the Outprfxesimpwdr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutprfxesimpwdr

`func (o *RoutingNeighborStatus) SetOutprfxesimpwdr(v int64)`

SetOutprfxesimpwdr sets Outprfxesimpwdr field to given value.

### HasOutprfxesimpwdr

`func (o *RoutingNeighborStatus) HasOutprfxesimpwdr() bool`

HasOutprfxesimpwdr returns a boolean if a field has been set.

### GetOutrefreshes

`func (o *RoutingNeighborStatus) GetOutrefreshes() int64`

GetOutrefreshes returns the Outrefreshes field if non-nil, zero value otherwise.

### GetOutrefreshesOk

`func (o *RoutingNeighborStatus) GetOutrefreshesOk() (*int64, bool)`

GetOutrefreshesOk returns a tuple with the Outrefreshes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutrefreshes

`func (o *RoutingNeighborStatus) SetOutrefreshes(v int64)`

SetOutrefreshes sets Outrefreshes field to given value.

### HasOutrefreshes

`func (o *RoutingNeighborStatus) HasOutrefreshes() bool`

HasOutrefreshes returns a boolean if a field has been set.

### GetOuttotalmessages

`func (o *RoutingNeighborStatus) GetOuttotalmessages() int64`

GetOuttotalmessages returns the Outtotalmessages field if non-nil, zero value otherwise.

### GetOuttotalmessagesOk

`func (o *RoutingNeighborStatus) GetOuttotalmessagesOk() (*int64, bool)`

GetOuttotalmessagesOk returns a tuple with the Outtotalmessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOuttotalmessages

`func (o *RoutingNeighborStatus) SetOuttotalmessages(v int64)`

SetOuttotalmessages sets Outtotalmessages field to given value.

### HasOuttotalmessages

`func (o *RoutingNeighborStatus) HasOuttotalmessages() bool`

HasOuttotalmessages returns a boolean if a field has been set.

### GetOutupdateelpstime

`func (o *RoutingNeighborStatus) GetOutupdateelpstime() int64`

GetOutupdateelpstime returns the Outupdateelpstime field if non-nil, zero value otherwise.

### GetOutupdateelpstimeOk

`func (o *RoutingNeighborStatus) GetOutupdateelpstimeOk() (*int64, bool)`

GetOutupdateelpstimeOk returns a tuple with the Outupdateelpstime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutupdateelpstime

`func (o *RoutingNeighborStatus) SetOutupdateelpstime(v int64)`

SetOutupdateelpstime sets Outupdateelpstime field to given value.

### HasOutupdateelpstime

`func (o *RoutingNeighborStatus) HasOutupdateelpstime() bool`

HasOutupdateelpstime returns a boolean if a field has been set.

### GetOutupdates

`func (o *RoutingNeighborStatus) GetOutupdates() int64`

GetOutupdates returns the Outupdates field if non-nil, zero value otherwise.

### GetOutupdatesOk

`func (o *RoutingNeighborStatus) GetOutupdatesOk() (*int64, bool)`

GetOutupdatesOk returns a tuple with the Outupdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutupdates

`func (o *RoutingNeighborStatus) SetOutupdates(v int64)`

SetOutupdates sets Outupdates field to given value.

### HasOutupdates

`func (o *RoutingNeighborStatus) HasOutupdates() bool`

HasOutupdates returns a boolean if a field has been set.

### GetPeergr

`func (o *RoutingNeighborStatus) GetPeergr() int64`

GetPeergr returns the Peergr field if non-nil, zero value otherwise.

### GetPeergrOk

`func (o *RoutingNeighborStatus) GetPeergrOk() (*int64, bool)`

GetPeergrOk returns a tuple with the Peergr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeergr

`func (o *RoutingNeighborStatus) SetPeergr(v int64)`

SetPeergr sets Peergr field to given value.

### HasPeergr

`func (o *RoutingNeighborStatus) HasPeergr() bool`

HasPeergr returns a boolean if a field has been set.

### GetPeerindex

`func (o *RoutingNeighborStatus) GetPeerindex() int64`

GetPeerindex returns the Peerindex field if non-nil, zero value otherwise.

### GetPeerindexOk

`func (o *RoutingNeighborStatus) GetPeerindexOk() (*int64, bool)`

GetPeerindexOk returns a tuple with the Peerindex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerindex

`func (o *RoutingNeighborStatus) SetPeerindex(v int64)`

SetPeerindex sets Peerindex field to given value.

### HasPeerindex

`func (o *RoutingNeighborStatus) HasPeerindex() bool`

HasPeerindex returns a boolean if a field has been set.

### GetPrevStatus

`func (o *RoutingNeighborStatus) GetPrevStatus() string`

GetPrevStatus returns the PrevStatus field if non-nil, zero value otherwise.

### GetPrevStatusOk

`func (o *RoutingNeighborStatus) GetPrevStatusOk() (*string, bool)`

GetPrevStatusOk returns a tuple with the PrevStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevStatus

`func (o *RoutingNeighborStatus) SetPrevStatus(v string)`

SetPrevStatus sets PrevStatus field to given value.

### HasPrevStatus

`func (o *RoutingNeighborStatus) HasPrevStatus() bool`

HasPrevStatus returns a boolean if a field has been set.

### GetRcvdmsgelpstime

`func (o *RoutingNeighborStatus) GetRcvdmsgelpstime() int64`

GetRcvdmsgelpstime returns the Rcvdmsgelpstime field if non-nil, zero value otherwise.

### GetRcvdmsgelpstimeOk

`func (o *RoutingNeighborStatus) GetRcvdmsgelpstimeOk() (*int64, bool)`

GetRcvdmsgelpstimeOk returns a tuple with the Rcvdmsgelpstime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRcvdmsgelpstime

`func (o *RoutingNeighborStatus) SetRcvdmsgelpstime(v int64)`

SetRcvdmsgelpstime sets Rcvdmsgelpstime field to given value.

### HasRcvdmsgelpstime

`func (o *RoutingNeighborStatus) HasRcvdmsgelpstime() bool`

HasRcvdmsgelpstime returns a boolean if a field has been set.

### GetReceivedholdtime

`func (o *RoutingNeighborStatus) GetReceivedholdtime() int64`

GetReceivedholdtime returns the Receivedholdtime field if non-nil, zero value otherwise.

### GetReceivedholdtimeOk

`func (o *RoutingNeighborStatus) GetReceivedholdtimeOk() (*int64, bool)`

GetReceivedholdtimeOk returns a tuple with the Receivedholdtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedholdtime

`func (o *RoutingNeighborStatus) SetReceivedholdtime(v int64)`

SetReceivedholdtime sets Receivedholdtime field to given value.

### HasReceivedholdtime

`func (o *RoutingNeighborStatus) HasReceivedholdtime() bool`

HasReceivedholdtime returns a boolean if a field has been set.

### GetRouterefrrcvd

`func (o *RoutingNeighborStatus) GetRouterefrrcvd() int64`

GetRouterefrrcvd returns the Routerefrrcvd field if non-nil, zero value otherwise.

### GetRouterefrrcvdOk

`func (o *RoutingNeighborStatus) GetRouterefrrcvdOk() (*int64, bool)`

GetRouterefrrcvdOk returns a tuple with the Routerefrrcvd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterefrrcvd

`func (o *RoutingNeighborStatus) SetRouterefrrcvd(v int64)`

SetRouterefrrcvd sets Routerefrrcvd field to given value.

### HasRouterefrrcvd

`func (o *RoutingNeighborStatus) HasRouterefrrcvd() bool`

HasRouterefrrcvd returns a boolean if a field has been set.

### GetRouterefrsent

`func (o *RoutingNeighborStatus) GetRouterefrsent() int64`

GetRouterefrsent returns the Routerefrsent field if non-nil, zero value otherwise.

### GetRouterefrsentOk

`func (o *RoutingNeighborStatus) GetRouterefrsentOk() (*int64, bool)`

GetRouterefrsentOk returns a tuple with the Routerefrsent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterefrsent

`func (o *RoutingNeighborStatus) SetRouterefrsent(v int64)`

SetRouterefrsent sets Routerefrsent field to given value.

### HasRouterefrsent

`func (o *RoutingNeighborStatus) HasRouterefrsent() bool`

HasRouterefrsent returns a boolean if a field has been set.

### GetSellocaladdrtype

`func (o *RoutingNeighborStatus) GetSellocaladdrtype() string`

GetSellocaladdrtype returns the Sellocaladdrtype field if non-nil, zero value otherwise.

### GetSellocaladdrtypeOk

`func (o *RoutingNeighborStatus) GetSellocaladdrtypeOk() (*string, bool)`

GetSellocaladdrtypeOk returns a tuple with the Sellocaladdrtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellocaladdrtype

`func (o *RoutingNeighborStatus) SetSellocaladdrtype(v string)`

SetSellocaladdrtype sets Sellocaladdrtype field to given value.

### HasSellocaladdrtype

`func (o *RoutingNeighborStatus) HasSellocaladdrtype() bool`

HasSellocaladdrtype returns a boolean if a field has been set.

### GetStalepathtime

`func (o *RoutingNeighborStatus) GetStalepathtime() int64`

GetStalepathtime returns the Stalepathtime field if non-nil, zero value otherwise.

### GetStalepathtimeOk

`func (o *RoutingNeighborStatus) GetStalepathtimeOk() (*int64, bool)`

GetStalepathtimeOk returns a tuple with the Stalepathtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStalepathtime

`func (o *RoutingNeighborStatus) SetStalepathtime(v int64)`

SetStalepathtime sets Stalepathtime field to given value.

### HasStalepathtime

`func (o *RoutingNeighborStatus) HasStalepathtime() bool`

HasStalepathtime returns a boolean if a field has been set.

### GetStatus

`func (o *RoutingNeighborStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RoutingNeighborStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RoutingNeighborStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RoutingNeighborStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


