# ApiKindWatchOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Options** | Pointer to [**ApiListWatchOptions**](apiListWatchOptions.md) |  | [optional] 

## Methods

### NewApiKindWatchOptions

`func NewApiKindWatchOptions() *ApiKindWatchOptions`

NewApiKindWatchOptions instantiates a new ApiKindWatchOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKindWatchOptionsWithDefaults

`func NewApiKindWatchOptionsWithDefaults() *ApiKindWatchOptions`

NewApiKindWatchOptionsWithDefaults instantiates a new ApiKindWatchOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *ApiKindWatchOptions) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ApiKindWatchOptions) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ApiKindWatchOptions) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *ApiKindWatchOptions) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetKind

`func (o *ApiKindWatchOptions) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ApiKindWatchOptions) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ApiKindWatchOptions) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ApiKindWatchOptions) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetOptions

`func (o *ApiKindWatchOptions) GetOptions() ApiListWatchOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ApiKindWatchOptions) GetOptionsOk() (*ApiListWatchOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ApiKindWatchOptions) SetOptions(v ApiListWatchOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ApiKindWatchOptions) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


