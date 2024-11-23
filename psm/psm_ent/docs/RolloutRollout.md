# RolloutRollout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**ApiObjectMeta**](apiObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**RolloutRolloutSpec**](rolloutRolloutSpec.md) |  | [optional] 
**Status** | Pointer to [**RolloutRolloutStatus**](rolloutRolloutStatus.md) |  | [optional] 

## Methods

### NewRolloutRollout

`func NewRolloutRollout() *RolloutRollout`

NewRolloutRollout instantiates a new RolloutRollout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRolloutRolloutWithDefaults

`func NewRolloutRolloutWithDefaults() *RolloutRollout`

NewRolloutRolloutWithDefaults instantiates a new RolloutRollout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *RolloutRollout) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *RolloutRollout) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *RolloutRollout) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *RolloutRollout) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *RolloutRollout) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *RolloutRollout) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *RolloutRollout) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *RolloutRollout) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMeta

`func (o *RolloutRollout) GetMeta() ApiObjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *RolloutRollout) GetMetaOk() (*ApiObjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *RolloutRollout) SetMeta(v ApiObjectMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *RolloutRollout) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSpec

`func (o *RolloutRollout) GetSpec() RolloutRolloutSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *RolloutRollout) GetSpecOk() (*RolloutRolloutSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *RolloutRollout) SetSpec(v RolloutRolloutSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *RolloutRollout) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *RolloutRollout) GetStatus() RolloutRolloutStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RolloutRollout) GetStatusOk() (*RolloutRolloutStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RolloutRollout) SetStatus(v RolloutRolloutStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RolloutRollout) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


