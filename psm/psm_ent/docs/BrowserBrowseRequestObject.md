# BrowserBrowseRequestObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountOnly** | Pointer to **bool** | When CountOnly is set the response only contains counts and not the actual objects. | [optional] 
**MaxDepth** | Pointer to **int64** | Max-Depth specifies how deep the query should explore. By default depth is set to 1 which means immediate relations 0 means to maximum depth. | [optional] [default to 1]
**QueryType** | Pointer to **string** | QueryType is the direction of the query. | [optional] [default to "dependencies"]
**Uri** | Pointer to **string** | URI is the root node from where to query. Length of string should be between 2 and 512. | [optional] 

## Methods

### NewBrowserBrowseRequestObject

`func NewBrowserBrowseRequestObject() *BrowserBrowseRequestObject`

NewBrowserBrowseRequestObject instantiates a new BrowserBrowseRequestObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrowserBrowseRequestObjectWithDefaults

`func NewBrowserBrowseRequestObjectWithDefaults() *BrowserBrowseRequestObject`

NewBrowserBrowseRequestObjectWithDefaults instantiates a new BrowserBrowseRequestObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountOnly

`func (o *BrowserBrowseRequestObject) GetCountOnly() bool`

GetCountOnly returns the CountOnly field if non-nil, zero value otherwise.

### GetCountOnlyOk

`func (o *BrowserBrowseRequestObject) GetCountOnlyOk() (*bool, bool)`

GetCountOnlyOk returns a tuple with the CountOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountOnly

`func (o *BrowserBrowseRequestObject) SetCountOnly(v bool)`

SetCountOnly sets CountOnly field to given value.

### HasCountOnly

`func (o *BrowserBrowseRequestObject) HasCountOnly() bool`

HasCountOnly returns a boolean if a field has been set.

### GetMaxDepth

`func (o *BrowserBrowseRequestObject) GetMaxDepth() int64`

GetMaxDepth returns the MaxDepth field if non-nil, zero value otherwise.

### GetMaxDepthOk

`func (o *BrowserBrowseRequestObject) GetMaxDepthOk() (*int64, bool)`

GetMaxDepthOk returns a tuple with the MaxDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDepth

`func (o *BrowserBrowseRequestObject) SetMaxDepth(v int64)`

SetMaxDepth sets MaxDepth field to given value.

### HasMaxDepth

`func (o *BrowserBrowseRequestObject) HasMaxDepth() bool`

HasMaxDepth returns a boolean if a field has been set.

### GetQueryType

`func (o *BrowserBrowseRequestObject) GetQueryType() string`

GetQueryType returns the QueryType field if non-nil, zero value otherwise.

### GetQueryTypeOk

`func (o *BrowserBrowseRequestObject) GetQueryTypeOk() (*string, bool)`

GetQueryTypeOk returns a tuple with the QueryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryType

`func (o *BrowserBrowseRequestObject) SetQueryType(v string)`

SetQueryType sets QueryType field to given value.

### HasQueryType

`func (o *BrowserBrowseRequestObject) HasQueryType() bool`

HasQueryType returns a boolean if a field has been set.

### GetUri

`func (o *BrowserBrowseRequestObject) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *BrowserBrowseRequestObject) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *BrowserBrowseRequestObject) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *BrowserBrowseRequestObject) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


