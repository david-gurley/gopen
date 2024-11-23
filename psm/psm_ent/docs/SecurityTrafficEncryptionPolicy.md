# SecurityTrafficEncryptionPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**ApiObjectMeta**](apiObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**SecurityTrafficEncryptionPolicySpec**](securityTrafficEncryptionPolicySpec.md) |  | [optional] 
**Status** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewSecurityTrafficEncryptionPolicy

`func NewSecurityTrafficEncryptionPolicy() *SecurityTrafficEncryptionPolicy`

NewSecurityTrafficEncryptionPolicy instantiates a new SecurityTrafficEncryptionPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityTrafficEncryptionPolicyWithDefaults

`func NewSecurityTrafficEncryptionPolicyWithDefaults() *SecurityTrafficEncryptionPolicy`

NewSecurityTrafficEncryptionPolicyWithDefaults instantiates a new SecurityTrafficEncryptionPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *SecurityTrafficEncryptionPolicy) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *SecurityTrafficEncryptionPolicy) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *SecurityTrafficEncryptionPolicy) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *SecurityTrafficEncryptionPolicy) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *SecurityTrafficEncryptionPolicy) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SecurityTrafficEncryptionPolicy) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SecurityTrafficEncryptionPolicy) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SecurityTrafficEncryptionPolicy) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMeta

`func (o *SecurityTrafficEncryptionPolicy) GetMeta() ApiObjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SecurityTrafficEncryptionPolicy) GetMetaOk() (*ApiObjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SecurityTrafficEncryptionPolicy) SetMeta(v ApiObjectMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SecurityTrafficEncryptionPolicy) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSpec

`func (o *SecurityTrafficEncryptionPolicy) GetSpec() SecurityTrafficEncryptionPolicySpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *SecurityTrafficEncryptionPolicy) GetSpecOk() (*SecurityTrafficEncryptionPolicySpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *SecurityTrafficEncryptionPolicy) SetSpec(v SecurityTrafficEncryptionPolicySpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *SecurityTrafficEncryptionPolicy) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *SecurityTrafficEncryptionPolicy) GetStatus() map[string]interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SecurityTrafficEncryptionPolicy) GetStatusOk() (*map[string]interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SecurityTrafficEncryptionPolicy) SetStatus(v map[string]interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SecurityTrafficEncryptionPolicy) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


