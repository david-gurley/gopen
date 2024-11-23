# BrowserBrowseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**MaxDepth** | Pointer to **int64** |  | [optional] 
**Meta** | Pointer to [**ApiObjectMeta**](apiObjectMeta.md) |  | [optional] 
**Objects** | Pointer to [**map[string]BrowserObject**](browserObject.md) |  | [optional] 
**QueryType** | Pointer to **string** |  | [optional] 
**RootUri** | Pointer to **string** |  | [optional] 
**TotalCount** | Pointer to **int64** |  | [optional] 

## Methods

### NewBrowserBrowseResponse

`func NewBrowserBrowseResponse() *BrowserBrowseResponse`

NewBrowserBrowseResponse instantiates a new BrowserBrowseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrowserBrowseResponseWithDefaults

`func NewBrowserBrowseResponseWithDefaults() *BrowserBrowseResponse`

NewBrowserBrowseResponseWithDefaults instantiates a new BrowserBrowseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *BrowserBrowseResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *BrowserBrowseResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *BrowserBrowseResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *BrowserBrowseResponse) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *BrowserBrowseResponse) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *BrowserBrowseResponse) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *BrowserBrowseResponse) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *BrowserBrowseResponse) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMaxDepth

`func (o *BrowserBrowseResponse) GetMaxDepth() int64`

GetMaxDepth returns the MaxDepth field if non-nil, zero value otherwise.

### GetMaxDepthOk

`func (o *BrowserBrowseResponse) GetMaxDepthOk() (*int64, bool)`

GetMaxDepthOk returns a tuple with the MaxDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDepth

`func (o *BrowserBrowseResponse) SetMaxDepth(v int64)`

SetMaxDepth sets MaxDepth field to given value.

### HasMaxDepth

`func (o *BrowserBrowseResponse) HasMaxDepth() bool`

HasMaxDepth returns a boolean if a field has been set.

### GetMeta

`func (o *BrowserBrowseResponse) GetMeta() ApiObjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BrowserBrowseResponse) GetMetaOk() (*ApiObjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BrowserBrowseResponse) SetMeta(v ApiObjectMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *BrowserBrowseResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetObjects

`func (o *BrowserBrowseResponse) GetObjects() map[string]BrowserObject`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *BrowserBrowseResponse) GetObjectsOk() (*map[string]BrowserObject, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *BrowserBrowseResponse) SetObjects(v map[string]BrowserObject)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *BrowserBrowseResponse) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### GetQueryType

`func (o *BrowserBrowseResponse) GetQueryType() string`

GetQueryType returns the QueryType field if non-nil, zero value otherwise.

### GetQueryTypeOk

`func (o *BrowserBrowseResponse) GetQueryTypeOk() (*string, bool)`

GetQueryTypeOk returns a tuple with the QueryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryType

`func (o *BrowserBrowseResponse) SetQueryType(v string)`

SetQueryType sets QueryType field to given value.

### HasQueryType

`func (o *BrowserBrowseResponse) HasQueryType() bool`

HasQueryType returns a boolean if a field has been set.

### GetRootUri

`func (o *BrowserBrowseResponse) GetRootUri() string`

GetRootUri returns the RootUri field if non-nil, zero value otherwise.

### GetRootUriOk

`func (o *BrowserBrowseResponse) GetRootUriOk() (*string, bool)`

GetRootUriOk returns a tuple with the RootUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootUri

`func (o *BrowserBrowseResponse) SetRootUri(v string)`

SetRootUri sets RootUri field to given value.

### HasRootUri

`func (o *BrowserBrowseResponse) HasRootUri() bool`

HasRootUri returns a boolean if a field has been set.

### GetTotalCount

`func (o *BrowserBrowseResponse) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *BrowserBrowseResponse) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *BrowserBrowseResponse) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *BrowserBrowseResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


