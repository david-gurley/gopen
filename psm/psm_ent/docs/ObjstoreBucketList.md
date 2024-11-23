# ObjstoreBucketList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Items** | Pointer to [**[]ObjstoreBucket**](ObjstoreBucket.md) | List of Bucket objects. | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**ListMeta** | Pointer to [**ApiListMeta**](apiListMeta.md) |  | [optional] 

## Methods

### NewObjstoreBucketList

`func NewObjstoreBucketList() *ObjstoreBucketList`

NewObjstoreBucketList instantiates a new ObjstoreBucketList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjstoreBucketListWithDefaults

`func NewObjstoreBucketListWithDefaults() *ObjstoreBucketList`

NewObjstoreBucketListWithDefaults instantiates a new ObjstoreBucketList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ObjstoreBucketList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ObjstoreBucketList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ObjstoreBucketList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ObjstoreBucketList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *ObjstoreBucketList) GetItems() []ObjstoreBucket`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ObjstoreBucketList) GetItemsOk() (*[]ObjstoreBucket, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ObjstoreBucketList) SetItems(v []ObjstoreBucket)`

SetItems sets Items field to given value.

### HasItems

`func (o *ObjstoreBucketList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetKind

`func (o *ObjstoreBucketList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ObjstoreBucketList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ObjstoreBucketList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ObjstoreBucketList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetListMeta

`func (o *ObjstoreBucketList) GetListMeta() ApiListMeta`

GetListMeta returns the ListMeta field if non-nil, zero value otherwise.

### GetListMetaOk

`func (o *ObjstoreBucketList) GetListMetaOk() (*ApiListMeta, bool)`

GetListMetaOk returns a tuple with the ListMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListMeta

`func (o *ObjstoreBucketList) SetListMeta(v ApiListMeta)`

SetListMeta sets ListMeta field to given value.

### HasListMeta

`func (o *ObjstoreBucketList) HasListMeta() bool`

HasListMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


