# BrowserObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Links** | Pointer to [**map[string]ObjectURIs**](ObjectURIs.md) | Links points to the relations of the object. The key for the map is the path to the filed which is causing the relation. | [optional] 
**Meta** | Pointer to [**ApiObjectMeta**](apiObjectMeta.md) |  | [optional] 
**QueryType** | Pointer to **string** | QueryType specifies the direction of the relations in Links. | [optional] [default to "dependencies"]
**Reverse** | Pointer to **string** | Reverse is the view from the object looking back in the reverse direction of the dependency tree. | [optional] 
**Uri** | Pointer to **string** | URI is the Browser URI for this object. | [optional] 

## Methods

### NewBrowserObject

`func NewBrowserObject() *BrowserObject`

NewBrowserObject instantiates a new BrowserObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrowserObjectWithDefaults

`func NewBrowserObjectWithDefaults() *BrowserObject`

NewBrowserObjectWithDefaults instantiates a new BrowserObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *BrowserObject) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *BrowserObject) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *BrowserObject) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *BrowserObject) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *BrowserObject) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *BrowserObject) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *BrowserObject) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *BrowserObject) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLinks

`func (o *BrowserObject) GetLinks() map[string]ObjectURIs`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BrowserObject) GetLinksOk() (*map[string]ObjectURIs, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BrowserObject) SetLinks(v map[string]ObjectURIs)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BrowserObject) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *BrowserObject) GetMeta() ApiObjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BrowserObject) GetMetaOk() (*ApiObjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BrowserObject) SetMeta(v ApiObjectMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *BrowserObject) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetQueryType

`func (o *BrowserObject) GetQueryType() string`

GetQueryType returns the QueryType field if non-nil, zero value otherwise.

### GetQueryTypeOk

`func (o *BrowserObject) GetQueryTypeOk() (*string, bool)`

GetQueryTypeOk returns a tuple with the QueryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryType

`func (o *BrowserObject) SetQueryType(v string)`

SetQueryType sets QueryType field to given value.

### HasQueryType

`func (o *BrowserObject) HasQueryType() bool`

HasQueryType returns a boolean if a field has been set.

### GetReverse

`func (o *BrowserObject) GetReverse() string`

GetReverse returns the Reverse field if non-nil, zero value otherwise.

### GetReverseOk

`func (o *BrowserObject) GetReverseOk() (*string, bool)`

GetReverseOk returns a tuple with the Reverse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverse

`func (o *BrowserObject) SetReverse(v string)`

SetReverse sets Reverse field to given value.

### HasReverse

`func (o *BrowserObject) HasReverse() bool`

HasReverse returns a boolean if a field has been set.

### GetUri

`func (o *BrowserObject) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *BrowserObject) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *BrowserObject) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *BrowserObject) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


