# OrchestrationOrchestrator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**ApiObjectMeta**](apiObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**OrchestrationOrchestratorSpec**](orchestrationOrchestratorSpec.md) |  | [optional] 
**Status** | Pointer to [**OrchestrationOrchestratorStatus**](orchestrationOrchestratorStatus.md) |  | [optional] 

## Methods

### NewOrchestrationOrchestrator

`func NewOrchestrationOrchestrator() *OrchestrationOrchestrator`

NewOrchestrationOrchestrator instantiates a new OrchestrationOrchestrator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrchestrationOrchestratorWithDefaults

`func NewOrchestrationOrchestratorWithDefaults() *OrchestrationOrchestrator`

NewOrchestrationOrchestratorWithDefaults instantiates a new OrchestrationOrchestrator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *OrchestrationOrchestrator) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *OrchestrationOrchestrator) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *OrchestrationOrchestrator) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *OrchestrationOrchestrator) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *OrchestrationOrchestrator) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *OrchestrationOrchestrator) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *OrchestrationOrchestrator) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *OrchestrationOrchestrator) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMeta

`func (o *OrchestrationOrchestrator) GetMeta() ApiObjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *OrchestrationOrchestrator) GetMetaOk() (*ApiObjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *OrchestrationOrchestrator) SetMeta(v ApiObjectMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *OrchestrationOrchestrator) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSpec

`func (o *OrchestrationOrchestrator) GetSpec() OrchestrationOrchestratorSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *OrchestrationOrchestrator) GetSpecOk() (*OrchestrationOrchestratorSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *OrchestrationOrchestrator) SetSpec(v OrchestrationOrchestratorSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *OrchestrationOrchestrator) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *OrchestrationOrchestrator) GetStatus() OrchestrationOrchestratorStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrchestrationOrchestrator) GetStatusOk() (*OrchestrationOrchestratorStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrchestrationOrchestrator) SetStatus(v OrchestrationOrchestratorStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OrchestrationOrchestrator) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


