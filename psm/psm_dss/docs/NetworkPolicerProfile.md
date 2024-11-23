# NetworkPolicerProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**ApiObjectMeta**](apiObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**NetworkPolicerProfileSpec**](networkPolicerProfileSpec.md) |  | [optional] 
**Status** | Pointer to [**NetworkPolicerProfileStatus**](networkPolicerProfileStatus.md) |  | [optional] 

## Methods

### NewNetworkPolicerProfile

`func NewNetworkPolicerProfile() *NetworkPolicerProfile`

NewNetworkPolicerProfile instantiates a new NetworkPolicerProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkPolicerProfileWithDefaults

`func NewNetworkPolicerProfileWithDefaults() *NetworkPolicerProfile`

NewNetworkPolicerProfileWithDefaults instantiates a new NetworkPolicerProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *NetworkPolicerProfile) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NetworkPolicerProfile) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NetworkPolicerProfile) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *NetworkPolicerProfile) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *NetworkPolicerProfile) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkPolicerProfile) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkPolicerProfile) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *NetworkPolicerProfile) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMeta

`func (o *NetworkPolicerProfile) GetMeta() ApiObjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *NetworkPolicerProfile) GetMetaOk() (*ApiObjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *NetworkPolicerProfile) SetMeta(v ApiObjectMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *NetworkPolicerProfile) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSpec

`func (o *NetworkPolicerProfile) GetSpec() NetworkPolicerProfileSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *NetworkPolicerProfile) GetSpecOk() (*NetworkPolicerProfileSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *NetworkPolicerProfile) SetSpec(v NetworkPolicerProfileSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *NetworkPolicerProfile) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *NetworkPolicerProfile) GetStatus() NetworkPolicerProfileStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkPolicerProfile) GetStatusOk() (*NetworkPolicerProfileStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkPolicerProfile) SetStatus(v NetworkPolicerProfileStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NetworkPolicerProfile) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


