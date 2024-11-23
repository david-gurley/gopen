# ApiAggWatchOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**ApiObjectMeta**](apiObjectMeta.md) |  | [optional] 
**WatchOptions** | Pointer to [**[]ApiKindWatchOptions**](ApiKindWatchOptions.md) |  | [optional] 

## Methods

### NewApiAggWatchOptions

`func NewApiAggWatchOptions() *ApiAggWatchOptions`

NewApiAggWatchOptions instantiates a new ApiAggWatchOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAggWatchOptionsWithDefaults

`func NewApiAggWatchOptionsWithDefaults() *ApiAggWatchOptions`

NewApiAggWatchOptionsWithDefaults instantiates a new ApiAggWatchOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ApiAggWatchOptions) GetMeta() ApiObjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ApiAggWatchOptions) GetMetaOk() (*ApiObjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ApiAggWatchOptions) SetMeta(v ApiObjectMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ApiAggWatchOptions) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetWatchOptions

`func (o *ApiAggWatchOptions) GetWatchOptions() []ApiKindWatchOptions`

GetWatchOptions returns the WatchOptions field if non-nil, zero value otherwise.

### GetWatchOptionsOk

`func (o *ApiAggWatchOptions) GetWatchOptionsOk() (*[]ApiKindWatchOptions, bool)`

GetWatchOptionsOk returns a tuple with the WatchOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchOptions

`func (o *ApiAggWatchOptions) SetWatchOptions(v []ApiKindWatchOptions)`

SetWatchOptions sets WatchOptions field to given value.

### HasWatchOptions

`func (o *ApiAggWatchOptions) HasWatchOptions() bool`

HasWatchOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


