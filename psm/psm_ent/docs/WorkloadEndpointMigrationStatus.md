# WorkloadEndpointMigrationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Status of migration. | [optional] [default to "none"]

## Methods

### NewWorkloadEndpointMigrationStatus

`func NewWorkloadEndpointMigrationStatus() *WorkloadEndpointMigrationStatus`

NewWorkloadEndpointMigrationStatus instantiates a new WorkloadEndpointMigrationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadEndpointMigrationStatusWithDefaults

`func NewWorkloadEndpointMigrationStatusWithDefaults() *WorkloadEndpointMigrationStatus`

NewWorkloadEndpointMigrationStatusWithDefaults instantiates a new WorkloadEndpointMigrationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *WorkloadEndpointMigrationStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkloadEndpointMigrationStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkloadEndpointMigrationStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkloadEndpointMigrationStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


