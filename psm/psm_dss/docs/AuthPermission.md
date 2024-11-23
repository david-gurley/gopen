# AuthPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to **[]string** | Actions are actions on a resource that a permission allows. | [optional] 
**ResourceGroup** | Pointer to **string** | ResourceGroup is grouping of resource types for which a permission is defined. It is empty for Search, Event, MetricsQuery and non-api server endpoint. Specifying \&quot;_All_\&quot; will match all api groups including empty group for non-api server endpoints like those defined in ResrcKind enum. | [optional] 
**ResourceKind** | Pointer to **string** | ResourceKind is a resource kind for which permission is defined. It can be an API Server object kind or kinds defined in ResrcKind enum. Specifying \&quot;_All_\&quot; will match all resource kinds. | [optional] 
**ResourceNames** | Pointer to **[]string** | ResourceNames identify specific objects on which this permission applies. If resource name is not specified in permission then it implies all resources of the specified kind. | [optional] 
**ResourceNamespace** | Pointer to **string** | ResourceNamespace is a namespace to which a resource (API Server object) belongs. Default value is \&quot;_All_\&quot; which matches all namespaces. | [optional] [default to "_All_"]
**ResourceTenant** | Pointer to **string** | ResourceTenant is the tenant to which resource belongs. It will be automatically set to the tenant to which role object belongs. Exception are roles in \&quot;default\&quot; tenant. Role in \&quot;default\&quot; tenant can include permissions for resources in other tenants. Specifying \&quot;_All_\&quot; will match all tenants. | [optional] 

## Methods

### NewAuthPermission

`func NewAuthPermission() *AuthPermission`

NewAuthPermission instantiates a new AuthPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthPermissionWithDefaults

`func NewAuthPermissionWithDefaults() *AuthPermission`

NewAuthPermissionWithDefaults instantiates a new AuthPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *AuthPermission) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *AuthPermission) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *AuthPermission) SetActions(v []string)`

SetActions sets Actions field to given value.

### HasActions

`func (o *AuthPermission) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetResourceGroup

`func (o *AuthPermission) GetResourceGroup() string`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *AuthPermission) GetResourceGroupOk() (*string, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *AuthPermission) SetResourceGroup(v string)`

SetResourceGroup sets ResourceGroup field to given value.

### HasResourceGroup

`func (o *AuthPermission) HasResourceGroup() bool`

HasResourceGroup returns a boolean if a field has been set.

### GetResourceKind

`func (o *AuthPermission) GetResourceKind() string`

GetResourceKind returns the ResourceKind field if non-nil, zero value otherwise.

### GetResourceKindOk

`func (o *AuthPermission) GetResourceKindOk() (*string, bool)`

GetResourceKindOk returns a tuple with the ResourceKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceKind

`func (o *AuthPermission) SetResourceKind(v string)`

SetResourceKind sets ResourceKind field to given value.

### HasResourceKind

`func (o *AuthPermission) HasResourceKind() bool`

HasResourceKind returns a boolean if a field has been set.

### GetResourceNames

`func (o *AuthPermission) GetResourceNames() []string`

GetResourceNames returns the ResourceNames field if non-nil, zero value otherwise.

### GetResourceNamesOk

`func (o *AuthPermission) GetResourceNamesOk() (*[]string, bool)`

GetResourceNamesOk returns a tuple with the ResourceNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceNames

`func (o *AuthPermission) SetResourceNames(v []string)`

SetResourceNames sets ResourceNames field to given value.

### HasResourceNames

`func (o *AuthPermission) HasResourceNames() bool`

HasResourceNames returns a boolean if a field has been set.

### GetResourceNamespace

`func (o *AuthPermission) GetResourceNamespace() string`

GetResourceNamespace returns the ResourceNamespace field if non-nil, zero value otherwise.

### GetResourceNamespaceOk

`func (o *AuthPermission) GetResourceNamespaceOk() (*string, bool)`

GetResourceNamespaceOk returns a tuple with the ResourceNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceNamespace

`func (o *AuthPermission) SetResourceNamespace(v string)`

SetResourceNamespace sets ResourceNamespace field to given value.

### HasResourceNamespace

`func (o *AuthPermission) HasResourceNamespace() bool`

HasResourceNamespace returns a boolean if a field has been set.

### GetResourceTenant

`func (o *AuthPermission) GetResourceTenant() string`

GetResourceTenant returns the ResourceTenant field if non-nil, zero value otherwise.

### GetResourceTenantOk

`func (o *AuthPermission) GetResourceTenantOk() (*string, bool)`

GetResourceTenantOk returns a tuple with the ResourceTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceTenant

`func (o *AuthPermission) SetResourceTenant(v string)`

SetResourceTenant sets ResourceTenant field to given value.

### HasResourceTenant

`func (o *AuthPermission) HasResourceTenant() bool`

HasResourceTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


