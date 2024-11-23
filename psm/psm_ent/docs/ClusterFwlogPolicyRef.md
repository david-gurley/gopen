# ClusterFwlogPolicyRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of FwlogPolicy. | [optional] 
**Tenant** | Pointer to **string** | Tenant of FwlogPolicy. | [optional] [default to "default"]

## Methods

### NewClusterFwlogPolicyRef

`func NewClusterFwlogPolicyRef() *ClusterFwlogPolicyRef`

NewClusterFwlogPolicyRef instantiates a new ClusterFwlogPolicyRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterFwlogPolicyRefWithDefaults

`func NewClusterFwlogPolicyRefWithDefaults() *ClusterFwlogPolicyRef`

NewClusterFwlogPolicyRefWithDefaults instantiates a new ClusterFwlogPolicyRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ClusterFwlogPolicyRef) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterFwlogPolicyRef) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterFwlogPolicyRef) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterFwlogPolicyRef) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTenant

`func (o *ClusterFwlogPolicyRef) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *ClusterFwlogPolicyRef) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *ClusterFwlogPolicyRef) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *ClusterFwlogPolicyRef) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


