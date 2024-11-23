# NetworkRouteTableList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Items** | Pointer to [**[]NetworkRouteTable**](NetworkRouteTable.md) | List of RouteTable objects. | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**ListMeta** | Pointer to [**ApiListMeta**](apiListMeta.md) |  | [optional] 

## Methods

### NewNetworkRouteTableList

`func NewNetworkRouteTableList() *NetworkRouteTableList`

NewNetworkRouteTableList instantiates a new NetworkRouteTableList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkRouteTableListWithDefaults

`func NewNetworkRouteTableListWithDefaults() *NetworkRouteTableList`

NewNetworkRouteTableListWithDefaults instantiates a new NetworkRouteTableList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *NetworkRouteTableList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NetworkRouteTableList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NetworkRouteTableList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *NetworkRouteTableList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *NetworkRouteTableList) GetItems() []NetworkRouteTable`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *NetworkRouteTableList) GetItemsOk() (*[]NetworkRouteTable, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *NetworkRouteTableList) SetItems(v []NetworkRouteTable)`

SetItems sets Items field to given value.

### HasItems

`func (o *NetworkRouteTableList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetKind

`func (o *NetworkRouteTableList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkRouteTableList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkRouteTableList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *NetworkRouteTableList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetListMeta

`func (o *NetworkRouteTableList) GetListMeta() ApiListMeta`

GetListMeta returns the ListMeta field if non-nil, zero value otherwise.

### GetListMetaOk

`func (o *NetworkRouteTableList) GetListMetaOk() (*ApiListMeta, bool)`

GetListMetaOk returns a tuple with the ListMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListMeta

`func (o *NetworkRouteTableList) SetListMeta(v ApiListMeta)`

SetListMeta sets ListMeta field to given value.

### HasListMeta

`func (o *NetworkRouteTableList) HasListMeta() bool`

HasListMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


