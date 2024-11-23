# AuthRoleBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**ApiObjectMeta**](apiObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**AuthRoleBindingSpec**](authRoleBindingSpec.md) |  | [optional] 
**Status** | Pointer to **map[string]interface{}** | status part of role binding object. | [optional] 

## Methods

### NewAuthRoleBinding

`func NewAuthRoleBinding() *AuthRoleBinding`

NewAuthRoleBinding instantiates a new AuthRoleBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthRoleBindingWithDefaults

`func NewAuthRoleBindingWithDefaults() *AuthRoleBinding`

NewAuthRoleBindingWithDefaults instantiates a new AuthRoleBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *AuthRoleBinding) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AuthRoleBinding) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AuthRoleBinding) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *AuthRoleBinding) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *AuthRoleBinding) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AuthRoleBinding) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AuthRoleBinding) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *AuthRoleBinding) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMeta

`func (o *AuthRoleBinding) GetMeta() ApiObjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AuthRoleBinding) GetMetaOk() (*ApiObjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AuthRoleBinding) SetMeta(v ApiObjectMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AuthRoleBinding) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSpec

`func (o *AuthRoleBinding) GetSpec() AuthRoleBindingSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *AuthRoleBinding) GetSpecOk() (*AuthRoleBindingSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *AuthRoleBinding) SetSpec(v AuthRoleBindingSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *AuthRoleBinding) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *AuthRoleBinding) GetStatus() map[string]interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AuthRoleBinding) GetStatusOk() (*map[string]interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AuthRoleBinding) SetStatus(v map[string]interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AuthRoleBinding) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


