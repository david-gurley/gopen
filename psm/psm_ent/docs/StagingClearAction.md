# StagingClearAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**ApiObjectMeta**](apiObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**StagingClearActionSpec**](stagingClearActionSpec.md) |  | [optional] 
**Status** | Pointer to [**StagingClearActionStatus**](stagingClearActionStatus.md) |  | [optional] 

## Methods

### NewStagingClearAction

`func NewStagingClearAction() *StagingClearAction`

NewStagingClearAction instantiates a new StagingClearAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStagingClearActionWithDefaults

`func NewStagingClearActionWithDefaults() *StagingClearAction`

NewStagingClearActionWithDefaults instantiates a new StagingClearAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *StagingClearAction) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *StagingClearAction) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *StagingClearAction) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *StagingClearAction) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *StagingClearAction) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *StagingClearAction) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *StagingClearAction) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *StagingClearAction) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMeta

`func (o *StagingClearAction) GetMeta() ApiObjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *StagingClearAction) GetMetaOk() (*ApiObjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *StagingClearAction) SetMeta(v ApiObjectMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *StagingClearAction) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSpec

`func (o *StagingClearAction) GetSpec() StagingClearActionSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *StagingClearAction) GetSpecOk() (*StagingClearActionSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *StagingClearAction) SetSpec(v StagingClearActionSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *StagingClearAction) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *StagingClearAction) GetStatus() StagingClearActionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StagingClearAction) GetStatusOk() (*StagingClearActionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StagingClearAction) SetStatus(v StagingClearActionStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StagingClearAction) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


