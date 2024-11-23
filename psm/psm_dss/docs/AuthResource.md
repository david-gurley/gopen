# AuthResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to **string** | Group of resource to which access is desired. | [optional] 
**Kind** | Pointer to **string** | Kind of resource to which access is desired. | [optional] 
**Name** | Pointer to **string** | Name of a specific resource to which access is desired. | [optional] 
**Namespace** | Pointer to **string** | Namespace of resource within which access to a resource is desired. | [optional] 
**Tenant** | Pointer to **string** | Tenant to which the resource belongs. | [optional] 

## Methods

### NewAuthResource

`func NewAuthResource() *AuthResource`

NewAuthResource instantiates a new AuthResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthResourceWithDefaults

`func NewAuthResourceWithDefaults() *AuthResource`

NewAuthResourceWithDefaults instantiates a new AuthResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *AuthResource) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *AuthResource) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *AuthResource) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *AuthResource) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetKind

`func (o *AuthResource) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AuthResource) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AuthResource) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *AuthResource) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *AuthResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuthResource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *AuthResource) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *AuthResource) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *AuthResource) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *AuthResource) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetTenant

`func (o *AuthResource) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *AuthResource) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *AuthResource) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *AuthResource) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


