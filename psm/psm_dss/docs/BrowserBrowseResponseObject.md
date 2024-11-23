# BrowserBrowseResponseObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxDepth** | Pointer to **int64** | MaxDepth that the response explored. Reflects the value specified in the query. | [optional] 
**Objects** | Pointer to [**map[string]BrowserObject**](browserObject.md) | map of results. Key to the map is the URI of the  Object. | [optional] 
**QueryType** | Pointer to **string** | QueryType is the direction of the query. | [optional] 
**RootUri** | Pointer to **string** | RootURI is the root node for the response. | [optional] 
**TotalCount** | Pointer to **int64** | TotalCount of objects in the response. | [optional] 

## Methods

### NewBrowserBrowseResponseObject

`func NewBrowserBrowseResponseObject() *BrowserBrowseResponseObject`

NewBrowserBrowseResponseObject instantiates a new BrowserBrowseResponseObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrowserBrowseResponseObjectWithDefaults

`func NewBrowserBrowseResponseObjectWithDefaults() *BrowserBrowseResponseObject`

NewBrowserBrowseResponseObjectWithDefaults instantiates a new BrowserBrowseResponseObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxDepth

`func (o *BrowserBrowseResponseObject) GetMaxDepth() int64`

GetMaxDepth returns the MaxDepth field if non-nil, zero value otherwise.

### GetMaxDepthOk

`func (o *BrowserBrowseResponseObject) GetMaxDepthOk() (*int64, bool)`

GetMaxDepthOk returns a tuple with the MaxDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDepth

`func (o *BrowserBrowseResponseObject) SetMaxDepth(v int64)`

SetMaxDepth sets MaxDepth field to given value.

### HasMaxDepth

`func (o *BrowserBrowseResponseObject) HasMaxDepth() bool`

HasMaxDepth returns a boolean if a field has been set.

### GetObjects

`func (o *BrowserBrowseResponseObject) GetObjects() map[string]BrowserObject`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *BrowserBrowseResponseObject) GetObjectsOk() (*map[string]BrowserObject, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *BrowserBrowseResponseObject) SetObjects(v map[string]BrowserObject)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *BrowserBrowseResponseObject) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### GetQueryType

`func (o *BrowserBrowseResponseObject) GetQueryType() string`

GetQueryType returns the QueryType field if non-nil, zero value otherwise.

### GetQueryTypeOk

`func (o *BrowserBrowseResponseObject) GetQueryTypeOk() (*string, bool)`

GetQueryTypeOk returns a tuple with the QueryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryType

`func (o *BrowserBrowseResponseObject) SetQueryType(v string)`

SetQueryType sets QueryType field to given value.

### HasQueryType

`func (o *BrowserBrowseResponseObject) HasQueryType() bool`

HasQueryType returns a boolean if a field has been set.

### GetRootUri

`func (o *BrowserBrowseResponseObject) GetRootUri() string`

GetRootUri returns the RootUri field if non-nil, zero value otherwise.

### GetRootUriOk

`func (o *BrowserBrowseResponseObject) GetRootUriOk() (*string, bool)`

GetRootUriOk returns a tuple with the RootUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootUri

`func (o *BrowserBrowseResponseObject) SetRootUri(v string)`

SetRootUri sets RootUri field to given value.

### HasRootUri

`func (o *BrowserBrowseResponseObject) HasRootUri() bool`

HasRootUri returns a boolean if a field has been set.

### GetTotalCount

`func (o *BrowserBrowseResponseObject) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *BrowserBrowseResponseObject) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *BrowserBrowseResponseObject) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *BrowserBrowseResponseObject) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


