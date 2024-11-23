# WorkloadWorkloadSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostName** | Pointer to **string** | Hostname of the server where the workload should be running. | [optional] 
**Interfaces** | Pointer to [**[]WorkloadWorkloadIntfSpec**](WorkloadWorkloadIntfSpec.md) | Spec of all interfaces in the Workload identified by Primary MAC. | [optional] 
**MigrationTimeout** | Pointer to **string** | Should be a valid time duration. | [optional] [default to "3m"]

## Methods

### NewWorkloadWorkloadSpec

`func NewWorkloadWorkloadSpec() *WorkloadWorkloadSpec`

NewWorkloadWorkloadSpec instantiates a new WorkloadWorkloadSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadWorkloadSpecWithDefaults

`func NewWorkloadWorkloadSpecWithDefaults() *WorkloadWorkloadSpec`

NewWorkloadWorkloadSpecWithDefaults instantiates a new WorkloadWorkloadSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostName

`func (o *WorkloadWorkloadSpec) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *WorkloadWorkloadSpec) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *WorkloadWorkloadSpec) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *WorkloadWorkloadSpec) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetInterfaces

`func (o *WorkloadWorkloadSpec) GetInterfaces() []WorkloadWorkloadIntfSpec`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *WorkloadWorkloadSpec) GetInterfacesOk() (*[]WorkloadWorkloadIntfSpec, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *WorkloadWorkloadSpec) SetInterfaces(v []WorkloadWorkloadIntfSpec)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *WorkloadWorkloadSpec) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetMigrationTimeout

`func (o *WorkloadWorkloadSpec) GetMigrationTimeout() string`

GetMigrationTimeout returns the MigrationTimeout field if non-nil, zero value otherwise.

### GetMigrationTimeoutOk

`func (o *WorkloadWorkloadSpec) GetMigrationTimeoutOk() (*string, bool)`

GetMigrationTimeoutOk returns a tuple with the MigrationTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationTimeout

`func (o *WorkloadWorkloadSpec) SetMigrationTimeout(v string)`

SetMigrationTimeout sets MigrationTimeout field to given value.

### HasMigrationTimeout

`func (o *WorkloadWorkloadSpec) HasMigrationTimeout() bool`

HasMigrationTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


