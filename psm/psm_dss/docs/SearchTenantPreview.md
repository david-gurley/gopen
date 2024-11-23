# SearchTenantPreview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tenants** | Pointer to [**map[string]SearchCategoryPreview**](searchCategoryPreview.md) |  | [optional] 

## Methods

### NewSearchTenantPreview

`func NewSearchTenantPreview() *SearchTenantPreview`

NewSearchTenantPreview instantiates a new SearchTenantPreview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchTenantPreviewWithDefaults

`func NewSearchTenantPreviewWithDefaults() *SearchTenantPreview`

NewSearchTenantPreviewWithDefaults instantiates a new SearchTenantPreview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenants

`func (o *SearchTenantPreview) GetTenants() map[string]SearchCategoryPreview`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *SearchTenantPreview) GetTenantsOk() (*map[string]SearchCategoryPreview, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *SearchTenantPreview) SetTenants(v map[string]SearchCategoryPreview)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *SearchTenantPreview) HasTenants() bool`

HasTenants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


