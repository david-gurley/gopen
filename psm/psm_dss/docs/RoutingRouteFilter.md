# RoutingRouteFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllRoutes** | Pointer to **bool** | Fetch all routes rather than just the best routes selected by BGP. | [optional] 
**ApiVersion** | Pointer to **string** |  | [optional] 
**Extcomm** | Pointer to **string** |  | [optional] 
**Instance** | Pointer to **string** |  | [optional] 
**Ipaddress** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**ApiObjectMeta**](apiObjectMeta.md) |  | [optional] 
**Nhop** | Pointer to **string** |  | [optional] 
**PageNumber** | Pointer to **int64** |  | [optional] 
**Rtype** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Vnid** | Pointer to **string** |  | [optional] 

## Methods

### NewRoutingRouteFilter

`func NewRoutingRouteFilter() *RoutingRouteFilter`

NewRoutingRouteFilter instantiates a new RoutingRouteFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingRouteFilterWithDefaults

`func NewRoutingRouteFilterWithDefaults() *RoutingRouteFilter`

NewRoutingRouteFilterWithDefaults instantiates a new RoutingRouteFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllRoutes

`func (o *RoutingRouteFilter) GetAllRoutes() bool`

GetAllRoutes returns the AllRoutes field if non-nil, zero value otherwise.

### GetAllRoutesOk

`func (o *RoutingRouteFilter) GetAllRoutesOk() (*bool, bool)`

GetAllRoutesOk returns a tuple with the AllRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllRoutes

`func (o *RoutingRouteFilter) SetAllRoutes(v bool)`

SetAllRoutes sets AllRoutes field to given value.

### HasAllRoutes

`func (o *RoutingRouteFilter) HasAllRoutes() bool`

HasAllRoutes returns a boolean if a field has been set.

### GetApiVersion

`func (o *RoutingRouteFilter) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *RoutingRouteFilter) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *RoutingRouteFilter) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *RoutingRouteFilter) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetExtcomm

`func (o *RoutingRouteFilter) GetExtcomm() string`

GetExtcomm returns the Extcomm field if non-nil, zero value otherwise.

### GetExtcommOk

`func (o *RoutingRouteFilter) GetExtcommOk() (*string, bool)`

GetExtcommOk returns a tuple with the Extcomm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtcomm

`func (o *RoutingRouteFilter) SetExtcomm(v string)`

SetExtcomm sets Extcomm field to given value.

### HasExtcomm

`func (o *RoutingRouteFilter) HasExtcomm() bool`

HasExtcomm returns a boolean if a field has been set.

### GetInstance

`func (o *RoutingRouteFilter) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *RoutingRouteFilter) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *RoutingRouteFilter) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *RoutingRouteFilter) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetIpaddress

`func (o *RoutingRouteFilter) GetIpaddress() string`

GetIpaddress returns the Ipaddress field if non-nil, zero value otherwise.

### GetIpaddressOk

`func (o *RoutingRouteFilter) GetIpaddressOk() (*string, bool)`

GetIpaddressOk returns a tuple with the Ipaddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpaddress

`func (o *RoutingRouteFilter) SetIpaddress(v string)`

SetIpaddress sets Ipaddress field to given value.

### HasIpaddress

`func (o *RoutingRouteFilter) HasIpaddress() bool`

HasIpaddress returns a boolean if a field has been set.

### GetKind

`func (o *RoutingRouteFilter) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *RoutingRouteFilter) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *RoutingRouteFilter) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *RoutingRouteFilter) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMeta

`func (o *RoutingRouteFilter) GetMeta() ApiObjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *RoutingRouteFilter) GetMetaOk() (*ApiObjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *RoutingRouteFilter) SetMeta(v ApiObjectMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *RoutingRouteFilter) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetNhop

`func (o *RoutingRouteFilter) GetNhop() string`

GetNhop returns the Nhop field if non-nil, zero value otherwise.

### GetNhopOk

`func (o *RoutingRouteFilter) GetNhopOk() (*string, bool)`

GetNhopOk returns a tuple with the Nhop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNhop

`func (o *RoutingRouteFilter) SetNhop(v string)`

SetNhop sets Nhop field to given value.

### HasNhop

`func (o *RoutingRouteFilter) HasNhop() bool`

HasNhop returns a boolean if a field has been set.

### GetPageNumber

`func (o *RoutingRouteFilter) GetPageNumber() int64`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *RoutingRouteFilter) GetPageNumberOk() (*int64, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *RoutingRouteFilter) SetPageNumber(v int64)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *RoutingRouteFilter) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetRtype

`func (o *RoutingRouteFilter) GetRtype() string`

GetRtype returns the Rtype field if non-nil, zero value otherwise.

### GetRtypeOk

`func (o *RoutingRouteFilter) GetRtypeOk() (*string, bool)`

GetRtypeOk returns a tuple with the Rtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtype

`func (o *RoutingRouteFilter) SetRtype(v string)`

SetRtype sets Rtype field to given value.

### HasRtype

`func (o *RoutingRouteFilter) HasRtype() bool`

HasRtype returns a boolean if a field has been set.

### GetType

`func (o *RoutingRouteFilter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RoutingRouteFilter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RoutingRouteFilter) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RoutingRouteFilter) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVnid

`func (o *RoutingRouteFilter) GetVnid() string`

GetVnid returns the Vnid field if non-nil, zero value otherwise.

### GetVnidOk

`func (o *RoutingRouteFilter) GetVnidOk() (*string, bool)`

GetVnidOk returns a tuple with the Vnid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnid

`func (o *RoutingRouteFilter) SetVnid(v string)`

SetVnid sets Vnid field to given value.

### HasVnid

`func (o *RoutingRouteFilter) HasVnid() bool`

HasVnid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


