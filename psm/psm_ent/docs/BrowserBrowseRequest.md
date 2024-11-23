# BrowserBrowseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**CountOnly** | Pointer to **bool** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**MaxDepth** | Pointer to **int64** |  | [optional] [default to 1]
**Meta** | Pointer to [**ApiObjectMeta**](apiObjectMeta.md) |  | [optional] 
**QueryType** | Pointer to **string** |  | [optional] [default to "dependencies"]
**Uri** | Pointer to **string** | Length of string should be between 2 and 512. | [optional] 

## Methods

### NewBrowserBrowseRequest

`func NewBrowserBrowseRequest() *BrowserBrowseRequest`

NewBrowserBrowseRequest instantiates a new BrowserBrowseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrowserBrowseRequestWithDefaults

`func NewBrowserBrowseRequestWithDefaults() *BrowserBrowseRequest`

NewBrowserBrowseRequestWithDefaults instantiates a new BrowserBrowseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *BrowserBrowseRequest) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *BrowserBrowseRequest) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *BrowserBrowseRequest) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *BrowserBrowseRequest) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetCountOnly

`func (o *BrowserBrowseRequest) GetCountOnly() bool`

GetCountOnly returns the CountOnly field if non-nil, zero value otherwise.

### GetCountOnlyOk

`func (o *BrowserBrowseRequest) GetCountOnlyOk() (*bool, bool)`

GetCountOnlyOk returns a tuple with the CountOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountOnly

`func (o *BrowserBrowseRequest) SetCountOnly(v bool)`

SetCountOnly sets CountOnly field to given value.

### HasCountOnly

`func (o *BrowserBrowseRequest) HasCountOnly() bool`

HasCountOnly returns a boolean if a field has been set.

### GetKind

`func (o *BrowserBrowseRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *BrowserBrowseRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *BrowserBrowseRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *BrowserBrowseRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMaxDepth

`func (o *BrowserBrowseRequest) GetMaxDepth() int64`

GetMaxDepth returns the MaxDepth field if non-nil, zero value otherwise.

### GetMaxDepthOk

`func (o *BrowserBrowseRequest) GetMaxDepthOk() (*int64, bool)`

GetMaxDepthOk returns a tuple with the MaxDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDepth

`func (o *BrowserBrowseRequest) SetMaxDepth(v int64)`

SetMaxDepth sets MaxDepth field to given value.

### HasMaxDepth

`func (o *BrowserBrowseRequest) HasMaxDepth() bool`

HasMaxDepth returns a boolean if a field has been set.

### GetMeta

`func (o *BrowserBrowseRequest) GetMeta() ApiObjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BrowserBrowseRequest) GetMetaOk() (*ApiObjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BrowserBrowseRequest) SetMeta(v ApiObjectMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *BrowserBrowseRequest) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetQueryType

`func (o *BrowserBrowseRequest) GetQueryType() string`

GetQueryType returns the QueryType field if non-nil, zero value otherwise.

### GetQueryTypeOk

`func (o *BrowserBrowseRequest) GetQueryTypeOk() (*string, bool)`

GetQueryTypeOk returns a tuple with the QueryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryType

`func (o *BrowserBrowseRequest) SetQueryType(v string)`

SetQueryType sets QueryType field to given value.

### HasQueryType

`func (o *BrowserBrowseRequest) HasQueryType() bool`

HasQueryType returns a boolean if a field has been set.

### GetUri

`func (o *BrowserBrowseRequest) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *BrowserBrowseRequest) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *BrowserBrowseRequest) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *BrowserBrowseRequest) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


