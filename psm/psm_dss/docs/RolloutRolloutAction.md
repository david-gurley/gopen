# RolloutRolloutAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**ApiObjectMeta**](apiObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**RolloutRolloutSpec**](rolloutRolloutSpec.md) |  | [optional] 
**Status** | Pointer to [**RolloutRolloutActionStatus**](rolloutRolloutActionStatus.md) |  | [optional] 

## Methods

### NewRolloutRolloutAction

`func NewRolloutRolloutAction() *RolloutRolloutAction`

NewRolloutRolloutAction instantiates a new RolloutRolloutAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRolloutRolloutActionWithDefaults

`func NewRolloutRolloutActionWithDefaults() *RolloutRolloutAction`

NewRolloutRolloutActionWithDefaults instantiates a new RolloutRolloutAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *RolloutRolloutAction) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *RolloutRolloutAction) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *RolloutRolloutAction) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *RolloutRolloutAction) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *RolloutRolloutAction) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *RolloutRolloutAction) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *RolloutRolloutAction) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *RolloutRolloutAction) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMeta

`func (o *RolloutRolloutAction) GetMeta() ApiObjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *RolloutRolloutAction) GetMetaOk() (*ApiObjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *RolloutRolloutAction) SetMeta(v ApiObjectMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *RolloutRolloutAction) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSpec

`func (o *RolloutRolloutAction) GetSpec() RolloutRolloutSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *RolloutRolloutAction) GetSpecOk() (*RolloutRolloutSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *RolloutRolloutAction) SetSpec(v RolloutRolloutSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *RolloutRolloutAction) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *RolloutRolloutAction) GetStatus() RolloutRolloutActionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RolloutRolloutAction) GetStatusOk() (*RolloutRolloutActionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RolloutRolloutAction) SetStatus(v RolloutRolloutActionStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RolloutRolloutAction) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


