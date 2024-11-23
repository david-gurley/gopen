# OrchestrationNamespaceSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ManagedSpec** | Pointer to [**OrchestrationManagedNamespaceSpec**](orchestrationManagedNamespaceSpec.md) |  | [optional] 
**Mode** | Pointer to **string** |  | [optional] [default to "managed"]
**MonitoredSpec** | Pointer to **map[string]interface{}** | MonitoredNamespaceSpec contains namespace specific configuration. | [optional] 
**Name** | Pointer to **string** | Length of string should be at least 1. | [optional] 

## Methods

### NewOrchestrationNamespaceSpec

`func NewOrchestrationNamespaceSpec() *OrchestrationNamespaceSpec`

NewOrchestrationNamespaceSpec instantiates a new OrchestrationNamespaceSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrchestrationNamespaceSpecWithDefaults

`func NewOrchestrationNamespaceSpecWithDefaults() *OrchestrationNamespaceSpec`

NewOrchestrationNamespaceSpecWithDefaults instantiates a new OrchestrationNamespaceSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManagedSpec

`func (o *OrchestrationNamespaceSpec) GetManagedSpec() OrchestrationManagedNamespaceSpec`

GetManagedSpec returns the ManagedSpec field if non-nil, zero value otherwise.

### GetManagedSpecOk

`func (o *OrchestrationNamespaceSpec) GetManagedSpecOk() (*OrchestrationManagedNamespaceSpec, bool)`

GetManagedSpecOk returns a tuple with the ManagedSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedSpec

`func (o *OrchestrationNamespaceSpec) SetManagedSpec(v OrchestrationManagedNamespaceSpec)`

SetManagedSpec sets ManagedSpec field to given value.

### HasManagedSpec

`func (o *OrchestrationNamespaceSpec) HasManagedSpec() bool`

HasManagedSpec returns a boolean if a field has been set.

### GetMode

`func (o *OrchestrationNamespaceSpec) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *OrchestrationNamespaceSpec) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *OrchestrationNamespaceSpec) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *OrchestrationNamespaceSpec) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetMonitoredSpec

`func (o *OrchestrationNamespaceSpec) GetMonitoredSpec() map[string]interface{}`

GetMonitoredSpec returns the MonitoredSpec field if non-nil, zero value otherwise.

### GetMonitoredSpecOk

`func (o *OrchestrationNamespaceSpec) GetMonitoredSpecOk() (*map[string]interface{}, bool)`

GetMonitoredSpecOk returns a tuple with the MonitoredSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoredSpec

`func (o *OrchestrationNamespaceSpec) SetMonitoredSpec(v map[string]interface{})`

SetMonitoredSpec sets MonitoredSpec field to given value.

### HasMonitoredSpec

`func (o *OrchestrationNamespaceSpec) HasMonitoredSpec() bool`

HasMonitoredSpec returns a boolean if a field has been set.

### GetName

`func (o *OrchestrationNamespaceSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrchestrationNamespaceSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrchestrationNamespaceSpec) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrchestrationNamespaceSpec) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


