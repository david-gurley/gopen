# ApiTypeMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the version of the API object. This can only be set by the server. | [optional] 
**Kind** | Pointer to **string** | Kind represents the type of the API object. | [optional] 

## Methods

### NewApiTypeMeta

`func NewApiTypeMeta() *ApiTypeMeta`

NewApiTypeMeta instantiates a new ApiTypeMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiTypeMetaWithDefaults

`func NewApiTypeMetaWithDefaults() *ApiTypeMeta`

NewApiTypeMetaWithDefaults instantiates a new ApiTypeMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ApiTypeMeta) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ApiTypeMeta) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ApiTypeMeta) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ApiTypeMeta) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *ApiTypeMeta) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ApiTypeMeta) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ApiTypeMeta) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ApiTypeMeta) HasKind() bool`

HasKind returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


