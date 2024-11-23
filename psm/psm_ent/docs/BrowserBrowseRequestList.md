# BrowserBrowseRequestList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**ApiObjectMeta**](apiObjectMeta.md) |  | [optional] 
**Requestlist** | Pointer to [**[]BrowserBrowseRequestObject**](BrowserBrowseRequestObject.md) |  | [optional] 

## Methods

### NewBrowserBrowseRequestList

`func NewBrowserBrowseRequestList() *BrowserBrowseRequestList`

NewBrowserBrowseRequestList instantiates a new BrowserBrowseRequestList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrowserBrowseRequestListWithDefaults

`func NewBrowserBrowseRequestListWithDefaults() *BrowserBrowseRequestList`

NewBrowserBrowseRequestListWithDefaults instantiates a new BrowserBrowseRequestList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *BrowserBrowseRequestList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *BrowserBrowseRequestList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *BrowserBrowseRequestList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *BrowserBrowseRequestList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *BrowserBrowseRequestList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *BrowserBrowseRequestList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *BrowserBrowseRequestList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *BrowserBrowseRequestList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMeta

`func (o *BrowserBrowseRequestList) GetMeta() ApiObjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BrowserBrowseRequestList) GetMetaOk() (*ApiObjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BrowserBrowseRequestList) SetMeta(v ApiObjectMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *BrowserBrowseRequestList) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetRequestlist

`func (o *BrowserBrowseRequestList) GetRequestlist() []BrowserBrowseRequestObject`

GetRequestlist returns the Requestlist field if non-nil, zero value otherwise.

### GetRequestlistOk

`func (o *BrowserBrowseRequestList) GetRequestlistOk() (*[]BrowserBrowseRequestObject, bool)`

GetRequestlistOk returns a tuple with the Requestlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestlist

`func (o *BrowserBrowseRequestList) SetRequestlist(v []BrowserBrowseRequestObject)`

SetRequestlist sets Requestlist field to given value.

### HasRequestlist

`func (o *BrowserBrowseRequestList) HasRequestlist() bool`

HasRequestlist returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


