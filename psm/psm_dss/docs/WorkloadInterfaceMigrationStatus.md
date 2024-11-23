# WorkloadInterfaceMigrationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** | The status from the dataplane performing migration. | [optional] [default to "none"]

## Methods

### NewWorkloadInterfaceMigrationStatus

`func NewWorkloadInterfaceMigrationStatus() *WorkloadInterfaceMigrationStatus`

NewWorkloadInterfaceMigrationStatus instantiates a new WorkloadInterfaceMigrationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadInterfaceMigrationStatusWithDefaults

`func NewWorkloadInterfaceMigrationStatusWithDefaults() *WorkloadInterfaceMigrationStatus`

NewWorkloadInterfaceMigrationStatusWithDefaults instantiates a new WorkloadInterfaceMigrationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *WorkloadInterfaceMigrationStatus) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *WorkloadInterfaceMigrationStatus) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *WorkloadInterfaceMigrationStatus) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *WorkloadInterfaceMigrationStatus) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatus

`func (o *WorkloadInterfaceMigrationStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkloadInterfaceMigrationStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkloadInterfaceMigrationStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkloadInterfaceMigrationStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


