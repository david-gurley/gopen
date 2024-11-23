# AuthAuthenticationPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**ApiObjectMeta**](apiObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**AuthAuthenticationPolicySpec**](authAuthenticationPolicySpec.md) |  | [optional] 
**Status** | Pointer to [**AuthAuthenticationPolicyStatus**](authAuthenticationPolicyStatus.md) |  | [optional] 

## Methods

### NewAuthAuthenticationPolicy

`func NewAuthAuthenticationPolicy() *AuthAuthenticationPolicy`

NewAuthAuthenticationPolicy instantiates a new AuthAuthenticationPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthAuthenticationPolicyWithDefaults

`func NewAuthAuthenticationPolicyWithDefaults() *AuthAuthenticationPolicy`

NewAuthAuthenticationPolicyWithDefaults instantiates a new AuthAuthenticationPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *AuthAuthenticationPolicy) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AuthAuthenticationPolicy) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AuthAuthenticationPolicy) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *AuthAuthenticationPolicy) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *AuthAuthenticationPolicy) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AuthAuthenticationPolicy) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AuthAuthenticationPolicy) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *AuthAuthenticationPolicy) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMeta

`func (o *AuthAuthenticationPolicy) GetMeta() ApiObjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AuthAuthenticationPolicy) GetMetaOk() (*ApiObjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AuthAuthenticationPolicy) SetMeta(v ApiObjectMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AuthAuthenticationPolicy) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSpec

`func (o *AuthAuthenticationPolicy) GetSpec() AuthAuthenticationPolicySpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *AuthAuthenticationPolicy) GetSpecOk() (*AuthAuthenticationPolicySpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *AuthAuthenticationPolicy) SetSpec(v AuthAuthenticationPolicySpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *AuthAuthenticationPolicy) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *AuthAuthenticationPolicy) GetStatus() AuthAuthenticationPolicyStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AuthAuthenticationPolicy) GetStatusOk() (*AuthAuthenticationPolicyStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AuthAuthenticationPolicy) SetStatus(v AuthAuthenticationPolicyStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AuthAuthenticationPolicy) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


