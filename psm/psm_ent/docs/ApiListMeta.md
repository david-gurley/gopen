# ApiListMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceVersion** | Pointer to **string** | Resource version of object store at the time of list generation. | [optional] 
**TotalCount** | Pointer to **int32** | TotalCount is the total count of results (non paginated) that the server holds. | [optional] 

## Methods

### NewApiListMeta

`func NewApiListMeta() *ApiListMeta`

NewApiListMeta instantiates a new ApiListMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiListMetaWithDefaults

`func NewApiListMetaWithDefaults() *ApiListMeta`

NewApiListMetaWithDefaults instantiates a new ApiListMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceVersion

`func (o *ApiListMeta) GetResourceVersion() string`

GetResourceVersion returns the ResourceVersion field if non-nil, zero value otherwise.

### GetResourceVersionOk

`func (o *ApiListMeta) GetResourceVersionOk() (*string, bool)`

GetResourceVersionOk returns a tuple with the ResourceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceVersion

`func (o *ApiListMeta) SetResourceVersion(v string)`

SetResourceVersion sets ResourceVersion field to given value.

### HasResourceVersion

`func (o *ApiListMeta) HasResourceVersion() bool`

HasResourceVersion returns a boolean if a field has been set.

### GetTotalCount

`func (o *ApiListMeta) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ApiListMeta) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ApiListMeta) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ApiListMeta) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


