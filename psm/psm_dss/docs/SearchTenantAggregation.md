# SearchTenantAggregation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tenants** | Pointer to [**map[string]SearchCategoryAggregation**](searchCategoryAggregation.md) |  | [optional] 

## Methods

### NewSearchTenantAggregation

`func NewSearchTenantAggregation() *SearchTenantAggregation`

NewSearchTenantAggregation instantiates a new SearchTenantAggregation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchTenantAggregationWithDefaults

`func NewSearchTenantAggregationWithDefaults() *SearchTenantAggregation`

NewSearchTenantAggregationWithDefaults instantiates a new SearchTenantAggregation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenants

`func (o *SearchTenantAggregation) GetTenants() map[string]SearchCategoryAggregation`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *SearchTenantAggregation) GetTenantsOk() (*map[string]SearchCategoryAggregation, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *SearchTenantAggregation) SetTenants(v map[string]SearchCategoryAggregation)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *SearchTenantAggregation) HasTenants() bool`

HasTenants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


