# RoutingRouteList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Count** | Pointer to **int64** | The number of items returned in the current response. | [optional] 
**Items** | Pointer to [**[]RoutingRoute**](RoutingRoute.md) | List of items being returned. | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**ListMeta** | Pointer to [**ApiListMeta**](apiListMeta.md) |  | [optional] 
**PageNumber** | Pointer to **int64** | The page number of the items being returned. | [optional] 
**Total** | Pointer to **int64** | The total number of items available in the dataplane. | [optional] 

## Methods

### NewRoutingRouteList

`func NewRoutingRouteList() *RoutingRouteList`

NewRoutingRouteList instantiates a new RoutingRouteList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingRouteListWithDefaults

`func NewRoutingRouteListWithDefaults() *RoutingRouteList`

NewRoutingRouteListWithDefaults instantiates a new RoutingRouteList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *RoutingRouteList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *RoutingRouteList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *RoutingRouteList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *RoutingRouteList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetCount

`func (o *RoutingRouteList) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *RoutingRouteList) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *RoutingRouteList) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *RoutingRouteList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetItems

`func (o *RoutingRouteList) GetItems() []RoutingRoute`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *RoutingRouteList) GetItemsOk() (*[]RoutingRoute, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *RoutingRouteList) SetItems(v []RoutingRoute)`

SetItems sets Items field to given value.

### HasItems

`func (o *RoutingRouteList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetKind

`func (o *RoutingRouteList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *RoutingRouteList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *RoutingRouteList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *RoutingRouteList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetListMeta

`func (o *RoutingRouteList) GetListMeta() ApiListMeta`

GetListMeta returns the ListMeta field if non-nil, zero value otherwise.

### GetListMetaOk

`func (o *RoutingRouteList) GetListMetaOk() (*ApiListMeta, bool)`

GetListMetaOk returns a tuple with the ListMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListMeta

`func (o *RoutingRouteList) SetListMeta(v ApiListMeta)`

SetListMeta sets ListMeta field to given value.

### HasListMeta

`func (o *RoutingRouteList) HasListMeta() bool`

HasListMeta returns a boolean if a field has been set.

### GetPageNumber

`func (o *RoutingRouteList) GetPageNumber() int64`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *RoutingRouteList) GetPageNumberOk() (*int64, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *RoutingRouteList) SetPageNumber(v int64)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *RoutingRouteList) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetTotal

`func (o *RoutingRouteList) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *RoutingRouteList) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *RoutingRouteList) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *RoutingRouteList) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


