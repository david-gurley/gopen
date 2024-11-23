# AuthSubjectAccessReviewRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**ApiObjectMeta**](apiObjectMeta.md) |  | [optional] 
**Operations** | Pointer to [**[]AuthOperation**](AuthOperation.md) |  | [optional] 

## Methods

### NewAuthSubjectAccessReviewRequest

`func NewAuthSubjectAccessReviewRequest() *AuthSubjectAccessReviewRequest`

NewAuthSubjectAccessReviewRequest instantiates a new AuthSubjectAccessReviewRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthSubjectAccessReviewRequestWithDefaults

`func NewAuthSubjectAccessReviewRequestWithDefaults() *AuthSubjectAccessReviewRequest`

NewAuthSubjectAccessReviewRequestWithDefaults instantiates a new AuthSubjectAccessReviewRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *AuthSubjectAccessReviewRequest) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AuthSubjectAccessReviewRequest) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AuthSubjectAccessReviewRequest) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *AuthSubjectAccessReviewRequest) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *AuthSubjectAccessReviewRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AuthSubjectAccessReviewRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AuthSubjectAccessReviewRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *AuthSubjectAccessReviewRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMeta

`func (o *AuthSubjectAccessReviewRequest) GetMeta() ApiObjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AuthSubjectAccessReviewRequest) GetMetaOk() (*ApiObjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AuthSubjectAccessReviewRequest) SetMeta(v ApiObjectMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AuthSubjectAccessReviewRequest) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetOperations

`func (o *AuthSubjectAccessReviewRequest) GetOperations() []AuthOperation`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *AuthSubjectAccessReviewRequest) GetOperationsOk() (*[]AuthOperation, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *AuthSubjectAccessReviewRequest) SetOperations(v []AuthOperation)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *AuthSubjectAccessReviewRequest) HasOperations() bool`

HasOperations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


