# ClusterFlowExportPolicyRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of FlowExportPolicy. | [optional] 
**Tenant** | Pointer to **string** | Tenant of FlowExportPolicy. | [optional] [default to "default"]

## Methods

### NewClusterFlowExportPolicyRef

`func NewClusterFlowExportPolicyRef() *ClusterFlowExportPolicyRef`

NewClusterFlowExportPolicyRef instantiates a new ClusterFlowExportPolicyRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterFlowExportPolicyRefWithDefaults

`func NewClusterFlowExportPolicyRefWithDefaults() *ClusterFlowExportPolicyRef`

NewClusterFlowExportPolicyRefWithDefaults instantiates a new ClusterFlowExportPolicyRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ClusterFlowExportPolicyRef) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterFlowExportPolicyRef) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterFlowExportPolicyRef) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterFlowExportPolicyRef) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTenant

`func (o *ClusterFlowExportPolicyRef) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *ClusterFlowExportPolicyRef) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *ClusterFlowExportPolicyRef) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *ClusterFlowExportPolicyRef) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


