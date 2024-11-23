# WorkloadEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**ApiObjectMeta**](apiObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**WorkloadEndpointSpec**](workloadEndpointSpec.md) |  | [optional] 
**Status** | Pointer to [**WorkloadEndpointStatus**](workloadEndpointStatus.md) |  | [optional] 

## Methods

### NewWorkloadEndpoint

`func NewWorkloadEndpoint() *WorkloadEndpoint`

NewWorkloadEndpoint instantiates a new WorkloadEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadEndpointWithDefaults

`func NewWorkloadEndpointWithDefaults() *WorkloadEndpoint`

NewWorkloadEndpointWithDefaults instantiates a new WorkloadEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *WorkloadEndpoint) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *WorkloadEndpoint) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *WorkloadEndpoint) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *WorkloadEndpoint) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *WorkloadEndpoint) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *WorkloadEndpoint) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *WorkloadEndpoint) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *WorkloadEndpoint) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMeta

`func (o *WorkloadEndpoint) GetMeta() ApiObjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *WorkloadEndpoint) GetMetaOk() (*ApiObjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *WorkloadEndpoint) SetMeta(v ApiObjectMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *WorkloadEndpoint) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSpec

`func (o *WorkloadEndpoint) GetSpec() WorkloadEndpointSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *WorkloadEndpoint) GetSpecOk() (*WorkloadEndpointSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *WorkloadEndpoint) SetSpec(v WorkloadEndpointSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *WorkloadEndpoint) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *WorkloadEndpoint) GetStatus() WorkloadEndpointStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkloadEndpoint) GetStatusOk() (*WorkloadEndpointStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkloadEndpoint) SetStatus(v WorkloadEndpointStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkloadEndpoint) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


