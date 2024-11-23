# AuthUserPreference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**ApiObjectMeta**](apiObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**AuthUserPreferenceSpec**](authUserPreferenceSpec.md) |  | [optional] 
**Status** | Pointer to **map[string]interface{}** | User Preference Status, currently holds nothing. | [optional] 

## Methods

### NewAuthUserPreference

`func NewAuthUserPreference() *AuthUserPreference`

NewAuthUserPreference instantiates a new AuthUserPreference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthUserPreferenceWithDefaults

`func NewAuthUserPreferenceWithDefaults() *AuthUserPreference`

NewAuthUserPreferenceWithDefaults instantiates a new AuthUserPreference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *AuthUserPreference) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AuthUserPreference) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AuthUserPreference) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *AuthUserPreference) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *AuthUserPreference) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AuthUserPreference) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AuthUserPreference) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *AuthUserPreference) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMeta

`func (o *AuthUserPreference) GetMeta() ApiObjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AuthUserPreference) GetMetaOk() (*ApiObjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AuthUserPreference) SetMeta(v ApiObjectMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AuthUserPreference) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSpec

`func (o *AuthUserPreference) GetSpec() AuthUserPreferenceSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *AuthUserPreference) GetSpecOk() (*AuthUserPreferenceSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *AuthUserPreference) SetSpec(v AuthUserPreferenceSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *AuthUserPreference) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *AuthUserPreference) GetStatus() map[string]interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AuthUserPreference) GetStatusOk() (*map[string]interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AuthUserPreference) SetStatus(v map[string]interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AuthUserPreference) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


