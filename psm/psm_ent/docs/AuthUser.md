# AuthUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**ApiObjectMeta**](apiObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**AuthUserSpec**](authUserSpec.md) |  | [optional] 
**Status** | Pointer to [**AuthUserStatus**](authUserStatus.md) |  | [optional] 

## Methods

### NewAuthUser

`func NewAuthUser() *AuthUser`

NewAuthUser instantiates a new AuthUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthUserWithDefaults

`func NewAuthUserWithDefaults() *AuthUser`

NewAuthUserWithDefaults instantiates a new AuthUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *AuthUser) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AuthUser) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AuthUser) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *AuthUser) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *AuthUser) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AuthUser) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AuthUser) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *AuthUser) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMeta

`func (o *AuthUser) GetMeta() ApiObjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AuthUser) GetMetaOk() (*ApiObjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AuthUser) SetMeta(v ApiObjectMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AuthUser) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSpec

`func (o *AuthUser) GetSpec() AuthUserSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *AuthUser) GetSpecOk() (*AuthUserSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *AuthUser) SetSpec(v AuthUserSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *AuthUser) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *AuthUser) GetStatus() AuthUserStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AuthUser) GetStatusOk() (*AuthUserStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AuthUser) SetStatus(v AuthUserStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AuthUser) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


