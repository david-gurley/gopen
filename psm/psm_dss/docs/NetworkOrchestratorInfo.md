# NetworkOrchestratorInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | Pointer to **string** | Namespace in the orchestrator in which this network should be created in. Length of string should be at least 1. | [optional] 
**OrchestratorName** | Pointer to **string** | Name of Orchestrator object to which this network should be applied to. Length of string should be at least 1. | [optional] 

## Methods

### NewNetworkOrchestratorInfo

`func NewNetworkOrchestratorInfo() *NetworkOrchestratorInfo`

NewNetworkOrchestratorInfo instantiates a new NetworkOrchestratorInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkOrchestratorInfoWithDefaults

`func NewNetworkOrchestratorInfoWithDefaults() *NetworkOrchestratorInfo`

NewNetworkOrchestratorInfoWithDefaults instantiates a new NetworkOrchestratorInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespace

`func (o *NetworkOrchestratorInfo) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *NetworkOrchestratorInfo) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *NetworkOrchestratorInfo) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *NetworkOrchestratorInfo) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetOrchestratorName

`func (o *NetworkOrchestratorInfo) GetOrchestratorName() string`

GetOrchestratorName returns the OrchestratorName field if non-nil, zero value otherwise.

### GetOrchestratorNameOk

`func (o *NetworkOrchestratorInfo) GetOrchestratorNameOk() (*string, bool)`

GetOrchestratorNameOk returns a tuple with the OrchestratorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrchestratorName

`func (o *NetworkOrchestratorInfo) SetOrchestratorName(v string)`

SetOrchestratorName sets OrchestratorName field to given value.

### HasOrchestratorName

`func (o *NetworkOrchestratorInfo) HasOrchestratorName() bool`

HasOrchestratorName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


