# OrchestrationOrchestratorSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to [**MonitoringExternalCred**](monitoringExternalCred.md) |  | [optional] 
**Namespaces** | Pointer to [**[]OrchestrationNamespaceSpec**](OrchestrationNamespaceSpec.md) | Namespaces are used to provide namespace specific information. From Rel-C this will be the only means to pass namespace information \&quot;all_namespaces\&quot; will be treated as a special namespace, which will apply the same configuration for all the namespaces discovered by the orchestrator. | [optional] 
**Type** | Pointer to **string** | Type of orchestrator. | [optional] [default to "vcenter"]
**Uri** | Pointer to **string** | URI of the orchestrator. Length of string should be at least 1. | [optional] 

## Methods

### NewOrchestrationOrchestratorSpec

`func NewOrchestrationOrchestratorSpec() *OrchestrationOrchestratorSpec`

NewOrchestrationOrchestratorSpec instantiates a new OrchestrationOrchestratorSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrchestrationOrchestratorSpecWithDefaults

`func NewOrchestrationOrchestratorSpecWithDefaults() *OrchestrationOrchestratorSpec`

NewOrchestrationOrchestratorSpecWithDefaults instantiates a new OrchestrationOrchestratorSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *OrchestrationOrchestratorSpec) GetCredentials() MonitoringExternalCred`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *OrchestrationOrchestratorSpec) GetCredentialsOk() (*MonitoringExternalCred, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *OrchestrationOrchestratorSpec) SetCredentials(v MonitoringExternalCred)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *OrchestrationOrchestratorSpec) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetNamespaces

`func (o *OrchestrationOrchestratorSpec) GetNamespaces() []OrchestrationNamespaceSpec`

GetNamespaces returns the Namespaces field if non-nil, zero value otherwise.

### GetNamespacesOk

`func (o *OrchestrationOrchestratorSpec) GetNamespacesOk() (*[]OrchestrationNamespaceSpec, bool)`

GetNamespacesOk returns a tuple with the Namespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaces

`func (o *OrchestrationOrchestratorSpec) SetNamespaces(v []OrchestrationNamespaceSpec)`

SetNamespaces sets Namespaces field to given value.

### HasNamespaces

`func (o *OrchestrationOrchestratorSpec) HasNamespaces() bool`

HasNamespaces returns a boolean if a field has been set.

### GetType

`func (o *OrchestrationOrchestratorSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrchestrationOrchestratorSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrchestrationOrchestratorSpec) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OrchestrationOrchestratorSpec) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUri

`func (o *OrchestrationOrchestratorSpec) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *OrchestrationOrchestratorSpec) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *OrchestrationOrchestratorSpec) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *OrchestrationOrchestratorSpec) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


