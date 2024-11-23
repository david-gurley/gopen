# WorkloadWorkloadStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostName** | Pointer to **string** | Hostname of the server where the workload is currently running. | [optional] 
**Interfaces** | Pointer to [**[]WorkloadWorkloadIntfStatus**](WorkloadWorkloadIntfStatus.md) | Status of all interfaces in the Workload identified by Primary MAC. | [optional] 
**MigrationStatus** | Pointer to [**WorkloadWorkloadMigrationStatus**](workloadWorkloadMigrationStatus.md) |  | [optional] 
**MirrorSessions** | Pointer to **[]string** | MirrorSessions list of mirror sessions enabled on this workload. | [optional] 
**PropagationStatus** | Pointer to [**SecurityPropagationStatus**](securityPropagationStatus.md) |  | [optional] 

## Methods

### NewWorkloadWorkloadStatus

`func NewWorkloadWorkloadStatus() *WorkloadWorkloadStatus`

NewWorkloadWorkloadStatus instantiates a new WorkloadWorkloadStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadWorkloadStatusWithDefaults

`func NewWorkloadWorkloadStatusWithDefaults() *WorkloadWorkloadStatus`

NewWorkloadWorkloadStatusWithDefaults instantiates a new WorkloadWorkloadStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostName

`func (o *WorkloadWorkloadStatus) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *WorkloadWorkloadStatus) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *WorkloadWorkloadStatus) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *WorkloadWorkloadStatus) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetInterfaces

`func (o *WorkloadWorkloadStatus) GetInterfaces() []WorkloadWorkloadIntfStatus`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *WorkloadWorkloadStatus) GetInterfacesOk() (*[]WorkloadWorkloadIntfStatus, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *WorkloadWorkloadStatus) SetInterfaces(v []WorkloadWorkloadIntfStatus)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *WorkloadWorkloadStatus) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetMigrationStatus

`func (o *WorkloadWorkloadStatus) GetMigrationStatus() WorkloadWorkloadMigrationStatus`

GetMigrationStatus returns the MigrationStatus field if non-nil, zero value otherwise.

### GetMigrationStatusOk

`func (o *WorkloadWorkloadStatus) GetMigrationStatusOk() (*WorkloadWorkloadMigrationStatus, bool)`

GetMigrationStatusOk returns a tuple with the MigrationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationStatus

`func (o *WorkloadWorkloadStatus) SetMigrationStatus(v WorkloadWorkloadMigrationStatus)`

SetMigrationStatus sets MigrationStatus field to given value.

### HasMigrationStatus

`func (o *WorkloadWorkloadStatus) HasMigrationStatus() bool`

HasMigrationStatus returns a boolean if a field has been set.

### GetMirrorSessions

`func (o *WorkloadWorkloadStatus) GetMirrorSessions() []string`

GetMirrorSessions returns the MirrorSessions field if non-nil, zero value otherwise.

### GetMirrorSessionsOk

`func (o *WorkloadWorkloadStatus) GetMirrorSessionsOk() (*[]string, bool)`

GetMirrorSessionsOk returns a tuple with the MirrorSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirrorSessions

`func (o *WorkloadWorkloadStatus) SetMirrorSessions(v []string)`

SetMirrorSessions sets MirrorSessions field to given value.

### HasMirrorSessions

`func (o *WorkloadWorkloadStatus) HasMirrorSessions() bool`

HasMirrorSessions returns a boolean if a field has been set.

### GetPropagationStatus

`func (o *WorkloadWorkloadStatus) GetPropagationStatus() SecurityPropagationStatus`

GetPropagationStatus returns the PropagationStatus field if non-nil, zero value otherwise.

### GetPropagationStatusOk

`func (o *WorkloadWorkloadStatus) GetPropagationStatusOk() (*SecurityPropagationStatus, bool)`

GetPropagationStatusOk returns a tuple with the PropagationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropagationStatus

`func (o *WorkloadWorkloadStatus) SetPropagationStatus(v SecurityPropagationStatus)`

SetPropagationStatus sets PropagationStatus field to given value.

### HasPropagationStatus

`func (o *WorkloadWorkloadStatus) HasPropagationStatus() bool`

HasPropagationStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


