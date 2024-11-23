# TelemetryQueryMetricsQueryResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Series** | Pointer to [**[]TelemetryQueryResultSeries**](TelemetryQueryResultSeries.md) |  | [optional] 
**StatementId** | Pointer to **int32** |  | [optional] 

## Methods

### NewTelemetryQueryMetricsQueryResult

`func NewTelemetryQueryMetricsQueryResult() *TelemetryQueryMetricsQueryResult`

NewTelemetryQueryMetricsQueryResult instantiates a new TelemetryQueryMetricsQueryResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryQueryMetricsQueryResultWithDefaults

`func NewTelemetryQueryMetricsQueryResultWithDefaults() *TelemetryQueryMetricsQueryResult`

NewTelemetryQueryMetricsQueryResultWithDefaults instantiates a new TelemetryQueryMetricsQueryResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeries

`func (o *TelemetryQueryMetricsQueryResult) GetSeries() []TelemetryQueryResultSeries`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *TelemetryQueryMetricsQueryResult) GetSeriesOk() (*[]TelemetryQueryResultSeries, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *TelemetryQueryMetricsQueryResult) SetSeries(v []TelemetryQueryResultSeries)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *TelemetryQueryMetricsQueryResult) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetStatementId

`func (o *TelemetryQueryMetricsQueryResult) GetStatementId() int32`

GetStatementId returns the StatementId field if non-nil, zero value otherwise.

### GetStatementIdOk

`func (o *TelemetryQueryMetricsQueryResult) GetStatementIdOk() (*int32, bool)`

GetStatementIdOk returns a tuple with the StatementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementId

`func (o *TelemetryQueryMetricsQueryResult) SetStatementId(v int32)`

SetStatementId sets StatementId field to given value.

### HasStatementId

`func (o *TelemetryQueryMetricsQueryResult) HasStatementId() bool`

HasStatementId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


