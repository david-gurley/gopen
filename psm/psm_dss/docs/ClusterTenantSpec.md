# ClusterTenantSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminUser** | Pointer to **string** | Tenant admin user. | [optional] 

## Methods

### NewClusterTenantSpec

`func NewClusterTenantSpec() *ClusterTenantSpec`

NewClusterTenantSpec instantiates a new ClusterTenantSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterTenantSpecWithDefaults

`func NewClusterTenantSpecWithDefaults() *ClusterTenantSpec`

NewClusterTenantSpecWithDefaults instantiates a new ClusterTenantSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminUser

`func (o *ClusterTenantSpec) GetAdminUser() string`

GetAdminUser returns the AdminUser field if non-nil, zero value otherwise.

### GetAdminUserOk

`func (o *ClusterTenantSpec) GetAdminUserOk() (*string, bool)`

GetAdminUserOk returns a tuple with the AdminUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminUser

`func (o *ClusterTenantSpec) SetAdminUser(v string)`

SetAdminUser sets AdminUser field to given value.

### HasAdminUser

`func (o *ClusterTenantSpec) HasAdminUser() bool`

HasAdminUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


