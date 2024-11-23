# FwlogFwLogList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**CountRelation** | Pointer to **string** | CountRelation to qualify total count of search results. | [optional] [default to "notsupported"]
**Items** | Pointer to [**[]FwlogFwLog**](FwlogFwLog.md) |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**ListMeta** | Pointer to [**ApiListMeta**](apiListMeta.md) |  | [optional] 
**ScrollId** | Pointer to **string** | ScrollID to scroll through results. | [optional] 

## Methods

### NewFwlogFwLogList

`func NewFwlogFwLogList() *FwlogFwLogList`

NewFwlogFwLogList instantiates a new FwlogFwLogList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFwlogFwLogListWithDefaults

`func NewFwlogFwLogListWithDefaults() *FwlogFwLogList`

NewFwlogFwLogListWithDefaults instantiates a new FwlogFwLogList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *FwlogFwLogList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *FwlogFwLogList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *FwlogFwLogList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *FwlogFwLogList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetCountRelation

`func (o *FwlogFwLogList) GetCountRelation() string`

GetCountRelation returns the CountRelation field if non-nil, zero value otherwise.

### GetCountRelationOk

`func (o *FwlogFwLogList) GetCountRelationOk() (*string, bool)`

GetCountRelationOk returns a tuple with the CountRelation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountRelation

`func (o *FwlogFwLogList) SetCountRelation(v string)`

SetCountRelation sets CountRelation field to given value.

### HasCountRelation

`func (o *FwlogFwLogList) HasCountRelation() bool`

HasCountRelation returns a boolean if a field has been set.

### GetItems

`func (o *FwlogFwLogList) GetItems() []FwlogFwLog`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *FwlogFwLogList) GetItemsOk() (*[]FwlogFwLog, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *FwlogFwLogList) SetItems(v []FwlogFwLog)`

SetItems sets Items field to given value.

### HasItems

`func (o *FwlogFwLogList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetKind

`func (o *FwlogFwLogList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *FwlogFwLogList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *FwlogFwLogList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *FwlogFwLogList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetListMeta

`func (o *FwlogFwLogList) GetListMeta() ApiListMeta`

GetListMeta returns the ListMeta field if non-nil, zero value otherwise.

### GetListMetaOk

`func (o *FwlogFwLogList) GetListMetaOk() (*ApiListMeta, bool)`

GetListMetaOk returns a tuple with the ListMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListMeta

`func (o *FwlogFwLogList) SetListMeta(v ApiListMeta)`

SetListMeta sets ListMeta field to given value.

### HasListMeta

`func (o *FwlogFwLogList) HasListMeta() bool`

HasListMeta returns a boolean if a field has been set.

### GetScrollId

`func (o *FwlogFwLogList) GetScrollId() string`

GetScrollId returns the ScrollId field if non-nil, zero value otherwise.

### GetScrollIdOk

`func (o *FwlogFwLogList) GetScrollIdOk() (*string, bool)`

GetScrollIdOk returns a tuple with the ScrollId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrollId

`func (o *FwlogFwLogList) SetScrollId(v string)`

SetScrollId sets ScrollId field to given value.

### HasScrollId

`func (o *FwlogFwLogList) HasScrollId() bool`

HasScrollId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


