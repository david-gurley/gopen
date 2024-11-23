# StagingCommitAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**ApiObjectMeta**](apiObjectMeta.md) |  | [optional] 
**Spec** | Pointer to **map[string]interface{}** |  | [optional] 
**Status** | Pointer to [**StagingCommitActionStatus**](stagingCommitActionStatus.md) |  | [optional] 

## Methods

### NewStagingCommitAction

`func NewStagingCommitAction() *StagingCommitAction`

NewStagingCommitAction instantiates a new StagingCommitAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStagingCommitActionWithDefaults

`func NewStagingCommitActionWithDefaults() *StagingCommitAction`

NewStagingCommitActionWithDefaults instantiates a new StagingCommitAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *StagingCommitAction) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *StagingCommitAction) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *StagingCommitAction) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *StagingCommitAction) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *StagingCommitAction) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *StagingCommitAction) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *StagingCommitAction) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *StagingCommitAction) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMeta

`func (o *StagingCommitAction) GetMeta() ApiObjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *StagingCommitAction) GetMetaOk() (*ApiObjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *StagingCommitAction) SetMeta(v ApiObjectMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *StagingCommitAction) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSpec

`func (o *StagingCommitAction) GetSpec() map[string]interface{}`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *StagingCommitAction) GetSpecOk() (*map[string]interface{}, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *StagingCommitAction) SetSpec(v map[string]interface{})`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *StagingCommitAction) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *StagingCommitAction) GetStatus() StagingCommitActionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StagingCommitAction) GetStatusOk() (*StagingCommitActionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StagingCommitAction) SetStatus(v StagingCommitActionStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StagingCommitAction) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


