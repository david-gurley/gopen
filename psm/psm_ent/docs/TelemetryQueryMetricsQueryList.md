# TelemetryQueryMetricsQueryList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | Pointer to **string** | Namespace for the request. | [optional] 
**Queries** | Pointer to [**[]TelemetryQueryMetricsQuerySpec**](TelemetryQueryMetricsQuerySpec.md) | List of queries to execute. | [optional] 
**Tenant** | Pointer to **string** | Tenant for the request. | [optional] 

## Methods

### NewTelemetryQueryMetricsQueryList

`func NewTelemetryQueryMetricsQueryList() *TelemetryQueryMetricsQueryList`

NewTelemetryQueryMetricsQueryList instantiates a new TelemetryQueryMetricsQueryList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryQueryMetricsQueryListWithDefaults

`func NewTelemetryQueryMetricsQueryListWithDefaults() *TelemetryQueryMetricsQueryList`

NewTelemetryQueryMetricsQueryListWithDefaults instantiates a new TelemetryQueryMetricsQueryList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespace

`func (o *TelemetryQueryMetricsQueryList) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *TelemetryQueryMetricsQueryList) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *TelemetryQueryMetricsQueryList) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *TelemetryQueryMetricsQueryList) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetQueries

`func (o *TelemetryQueryMetricsQueryList) GetQueries() []TelemetryQueryMetricsQuerySpec`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *TelemetryQueryMetricsQueryList) GetQueriesOk() (*[]TelemetryQueryMetricsQuerySpec, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *TelemetryQueryMetricsQueryList) SetQueries(v []TelemetryQueryMetricsQuerySpec)`

SetQueries sets Queries field to given value.

### HasQueries

`func (o *TelemetryQueryMetricsQueryList) HasQueries() bool`

HasQueries returns a boolean if a field has been set.

### GetTenant

`func (o *TelemetryQueryMetricsQueryList) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *TelemetryQueryMetricsQueryList) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *TelemetryQueryMetricsQueryList) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *TelemetryQueryMetricsQueryList) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


