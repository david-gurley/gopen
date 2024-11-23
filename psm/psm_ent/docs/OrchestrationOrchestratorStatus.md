# OrchestrationOrchestratorStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionStatus** | Pointer to **string** |  | [optional] [default to "unknown"]
**DiscoveredNamespaces** | Pointer to **[]string** |  | [optional] 
**IncompatibleDscs** | Pointer to **[]string** |  | [optional] 
**LastTransitionTime** | Pointer to **time.Time** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**OrchId** | Pointer to **int32** |  | [optional] 

## Methods

### NewOrchestrationOrchestratorStatus

`func NewOrchestrationOrchestratorStatus() *OrchestrationOrchestratorStatus`

NewOrchestrationOrchestratorStatus instantiates a new OrchestrationOrchestratorStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrchestrationOrchestratorStatusWithDefaults

`func NewOrchestrationOrchestratorStatusWithDefaults() *OrchestrationOrchestratorStatus`

NewOrchestrationOrchestratorStatusWithDefaults instantiates a new OrchestrationOrchestratorStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionStatus

`func (o *OrchestrationOrchestratorStatus) GetConnectionStatus() string`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *OrchestrationOrchestratorStatus) GetConnectionStatusOk() (*string, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *OrchestrationOrchestratorStatus) SetConnectionStatus(v string)`

SetConnectionStatus sets ConnectionStatus field to given value.

### HasConnectionStatus

`func (o *OrchestrationOrchestratorStatus) HasConnectionStatus() bool`

HasConnectionStatus returns a boolean if a field has been set.

### GetDiscoveredNamespaces

`func (o *OrchestrationOrchestratorStatus) GetDiscoveredNamespaces() []string`

GetDiscoveredNamespaces returns the DiscoveredNamespaces field if non-nil, zero value otherwise.

### GetDiscoveredNamespacesOk

`func (o *OrchestrationOrchestratorStatus) GetDiscoveredNamespacesOk() (*[]string, bool)`

GetDiscoveredNamespacesOk returns a tuple with the DiscoveredNamespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredNamespaces

`func (o *OrchestrationOrchestratorStatus) SetDiscoveredNamespaces(v []string)`

SetDiscoveredNamespaces sets DiscoveredNamespaces field to given value.

### HasDiscoveredNamespaces

`func (o *OrchestrationOrchestratorStatus) HasDiscoveredNamespaces() bool`

HasDiscoveredNamespaces returns a boolean if a field has been set.

### GetIncompatibleDscs

`func (o *OrchestrationOrchestratorStatus) GetIncompatibleDscs() []string`

GetIncompatibleDscs returns the IncompatibleDscs field if non-nil, zero value otherwise.

### GetIncompatibleDscsOk

`func (o *OrchestrationOrchestratorStatus) GetIncompatibleDscsOk() (*[]string, bool)`

GetIncompatibleDscsOk returns a tuple with the IncompatibleDscs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncompatibleDscs

`func (o *OrchestrationOrchestratorStatus) SetIncompatibleDscs(v []string)`

SetIncompatibleDscs sets IncompatibleDscs field to given value.

### HasIncompatibleDscs

`func (o *OrchestrationOrchestratorStatus) HasIncompatibleDscs() bool`

HasIncompatibleDscs returns a boolean if a field has been set.

### GetLastTransitionTime

`func (o *OrchestrationOrchestratorStatus) GetLastTransitionTime() time.Time`

GetLastTransitionTime returns the LastTransitionTime field if non-nil, zero value otherwise.

### GetLastTransitionTimeOk

`func (o *OrchestrationOrchestratorStatus) GetLastTransitionTimeOk() (*time.Time, bool)`

GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransitionTime

`func (o *OrchestrationOrchestratorStatus) SetLastTransitionTime(v time.Time)`

SetLastTransitionTime sets LastTransitionTime field to given value.

### HasLastTransitionTime

`func (o *OrchestrationOrchestratorStatus) HasLastTransitionTime() bool`

HasLastTransitionTime returns a boolean if a field has been set.

### GetMessage

`func (o *OrchestrationOrchestratorStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *OrchestrationOrchestratorStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *OrchestrationOrchestratorStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *OrchestrationOrchestratorStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOrchId

`func (o *OrchestrationOrchestratorStatus) GetOrchId() int32`

GetOrchId returns the OrchId field if non-nil, zero value otherwise.

### GetOrchIdOk

`func (o *OrchestrationOrchestratorStatus) GetOrchIdOk() (*int32, bool)`

GetOrchIdOk returns a tuple with the OrchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrchId

`func (o *OrchestrationOrchestratorStatus) SetOrchId(v int32)`

SetOrchId sets OrchId field to given value.

### HasOrchId

`func (o *OrchestrationOrchestratorStatus) HasOrchId() bool`

HasOrchId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


