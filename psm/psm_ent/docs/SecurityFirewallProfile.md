# SecurityFirewallProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**ApiObjectMeta**](apiObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**SecurityFirewallProfileSpec**](securityFirewallProfileSpec.md) |  | [optional] 
**Status** | Pointer to [**SecurityFirewallProfileStatus**](securityFirewallProfileStatus.md) |  | [optional] 

## Methods

### NewSecurityFirewallProfile

`func NewSecurityFirewallProfile() *SecurityFirewallProfile`

NewSecurityFirewallProfile instantiates a new SecurityFirewallProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityFirewallProfileWithDefaults

`func NewSecurityFirewallProfileWithDefaults() *SecurityFirewallProfile`

NewSecurityFirewallProfileWithDefaults instantiates a new SecurityFirewallProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *SecurityFirewallProfile) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *SecurityFirewallProfile) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *SecurityFirewallProfile) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *SecurityFirewallProfile) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *SecurityFirewallProfile) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SecurityFirewallProfile) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SecurityFirewallProfile) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SecurityFirewallProfile) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMeta

`func (o *SecurityFirewallProfile) GetMeta() ApiObjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SecurityFirewallProfile) GetMetaOk() (*ApiObjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SecurityFirewallProfile) SetMeta(v ApiObjectMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SecurityFirewallProfile) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSpec

`func (o *SecurityFirewallProfile) GetSpec() SecurityFirewallProfileSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *SecurityFirewallProfile) GetSpecOk() (*SecurityFirewallProfileSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *SecurityFirewallProfile) SetSpec(v SecurityFirewallProfileSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *SecurityFirewallProfile) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *SecurityFirewallProfile) GetStatus() SecurityFirewallProfileStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SecurityFirewallProfile) GetStatusOk() (*SecurityFirewallProfileStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SecurityFirewallProfile) SetStatus(v SecurityFirewallProfileStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SecurityFirewallProfile) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


