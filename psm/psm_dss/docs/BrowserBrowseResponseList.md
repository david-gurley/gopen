# BrowserBrowseResponseList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**ApiObjectMeta**](apiObjectMeta.md) |  | [optional] 
**Responselist** | Pointer to [**[]BrowserBrowseResponseObject**](BrowserBrowseResponseObject.md) |  | [optional] 

## Methods

### NewBrowserBrowseResponseList

`func NewBrowserBrowseResponseList() *BrowserBrowseResponseList`

NewBrowserBrowseResponseList instantiates a new BrowserBrowseResponseList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrowserBrowseResponseListWithDefaults

`func NewBrowserBrowseResponseListWithDefaults() *BrowserBrowseResponseList`

NewBrowserBrowseResponseListWithDefaults instantiates a new BrowserBrowseResponseList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *BrowserBrowseResponseList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *BrowserBrowseResponseList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *BrowserBrowseResponseList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *BrowserBrowseResponseList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *BrowserBrowseResponseList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *BrowserBrowseResponseList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *BrowserBrowseResponseList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *BrowserBrowseResponseList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMeta

`func (o *BrowserBrowseResponseList) GetMeta() ApiObjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BrowserBrowseResponseList) GetMetaOk() (*ApiObjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BrowserBrowseResponseList) SetMeta(v ApiObjectMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *BrowserBrowseResponseList) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetResponselist

`func (o *BrowserBrowseResponseList) GetResponselist() []BrowserBrowseResponseObject`

GetResponselist returns the Responselist field if non-nil, zero value otherwise.

### GetResponselistOk

`func (o *BrowserBrowseResponseList) GetResponselistOk() (*[]BrowserBrowseResponseObject, bool)`

GetResponselistOk returns a tuple with the Responselist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponselist

`func (o *BrowserBrowseResponseList) SetResponselist(v []BrowserBrowseResponseObject)`

SetResponselist sets Responselist field to given value.

### HasResponselist

`func (o *BrowserBrowseResponseList) HasResponselist() bool`

HasResponselist returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


