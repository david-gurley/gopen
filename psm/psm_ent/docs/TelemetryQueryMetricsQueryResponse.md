# TelemetryQueryMetricsQueryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | Pointer to **string** | Namespace for the request. | [optional] 
**Results** | Pointer to [**[]TelemetryQueryMetricsQueryResult**](TelemetryQueryMetricsQueryResult.md) |  | [optional] 
**Tenant** | Pointer to **string** | Tenant for the request. | [optional] 

## Methods

### NewTelemetryQueryMetricsQueryResponse

`func NewTelemetryQueryMetricsQueryResponse() *TelemetryQueryMetricsQueryResponse`

NewTelemetryQueryMetricsQueryResponse instantiates a new TelemetryQueryMetricsQueryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryQueryMetricsQueryResponseWithDefaults

`func NewTelemetryQueryMetricsQueryResponseWithDefaults() *TelemetryQueryMetricsQueryResponse`

NewTelemetryQueryMetricsQueryResponseWithDefaults instantiates a new TelemetryQueryMetricsQueryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespace

`func (o *TelemetryQueryMetricsQueryResponse) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *TelemetryQueryMetricsQueryResponse) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *TelemetryQueryMetricsQueryResponse) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *TelemetryQueryMetricsQueryResponse) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetResults

`func (o *TelemetryQueryMetricsQueryResponse) GetResults() []TelemetryQueryMetricsQueryResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *TelemetryQueryMetricsQueryResponse) GetResultsOk() (*[]TelemetryQueryMetricsQueryResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *TelemetryQueryMetricsQueryResponse) SetResults(v []TelemetryQueryMetricsQueryResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *TelemetryQueryMetricsQueryResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetTenant

`func (o *TelemetryQueryMetricsQueryResponse) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *TelemetryQueryMetricsQueryResponse) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *TelemetryQueryMetricsQueryResponse) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *TelemetryQueryMetricsQueryResponse) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


