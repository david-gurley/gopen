# SysruntimeConnectionStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**DataTruncated** | Pointer to **bool** | connection data trucated. use filters to avoid truncation. | [optional] 
**Items** | Pointer to [**[]SysruntimeHWConnectionGetResponse**](SysruntimeHWConnectionGetResponse.md) |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**ListMeta** | Pointer to [**ApiListMeta**](apiListMeta.md) |  | [optional] 

## Methods

### NewSysruntimeConnectionStatus

`func NewSysruntimeConnectionStatus() *SysruntimeConnectionStatus`

NewSysruntimeConnectionStatus instantiates a new SysruntimeConnectionStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSysruntimeConnectionStatusWithDefaults

`func NewSysruntimeConnectionStatusWithDefaults() *SysruntimeConnectionStatus`

NewSysruntimeConnectionStatusWithDefaults instantiates a new SysruntimeConnectionStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *SysruntimeConnectionStatus) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *SysruntimeConnectionStatus) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *SysruntimeConnectionStatus) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *SysruntimeConnectionStatus) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetDataTruncated

`func (o *SysruntimeConnectionStatus) GetDataTruncated() bool`

GetDataTruncated returns the DataTruncated field if non-nil, zero value otherwise.

### GetDataTruncatedOk

`func (o *SysruntimeConnectionStatus) GetDataTruncatedOk() (*bool, bool)`

GetDataTruncatedOk returns a tuple with the DataTruncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTruncated

`func (o *SysruntimeConnectionStatus) SetDataTruncated(v bool)`

SetDataTruncated sets DataTruncated field to given value.

### HasDataTruncated

`func (o *SysruntimeConnectionStatus) HasDataTruncated() bool`

HasDataTruncated returns a boolean if a field has been set.

### GetItems

`func (o *SysruntimeConnectionStatus) GetItems() []SysruntimeHWConnectionGetResponse`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SysruntimeConnectionStatus) GetItemsOk() (*[]SysruntimeHWConnectionGetResponse, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SysruntimeConnectionStatus) SetItems(v []SysruntimeHWConnectionGetResponse)`

SetItems sets Items field to given value.

### HasItems

`func (o *SysruntimeConnectionStatus) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetKind

`func (o *SysruntimeConnectionStatus) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SysruntimeConnectionStatus) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SysruntimeConnectionStatus) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SysruntimeConnectionStatus) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetListMeta

`func (o *SysruntimeConnectionStatus) GetListMeta() ApiListMeta`

GetListMeta returns the ListMeta field if non-nil, zero value otherwise.

### GetListMetaOk

`func (o *SysruntimeConnectionStatus) GetListMetaOk() (*ApiListMeta, bool)`

GetListMetaOk returns a tuple with the ListMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListMeta

`func (o *SysruntimeConnectionStatus) SetListMeta(v ApiListMeta)`

SetListMeta sets ListMeta field to given value.

### HasListMeta

`func (o *SysruntimeConnectionStatus) HasListMeta() bool`

HasListMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


