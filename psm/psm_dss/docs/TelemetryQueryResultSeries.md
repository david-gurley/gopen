# TelemetryQueryResultSeries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Columns** | Pointer to **[]string** | columns list all available fields in tsdb. | [optional] 
**Name** | Pointer to **string** | Name of the series. | [optional] 
**Tags** | Pointer to **map[string]string** | Tags are the TSDB tags in the query response. | [optional] 
**Values** | Pointer to [**[]ApiInterfaceSlice**](ApiInterfaceSlice.md) | values contain field values received frpm tsdb, it is in the form of [][]interface{}. | [optional] 

## Methods

### NewTelemetryQueryResultSeries

`func NewTelemetryQueryResultSeries() *TelemetryQueryResultSeries`

NewTelemetryQueryResultSeries instantiates a new TelemetryQueryResultSeries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryQueryResultSeriesWithDefaults

`func NewTelemetryQueryResultSeriesWithDefaults() *TelemetryQueryResultSeries`

NewTelemetryQueryResultSeriesWithDefaults instantiates a new TelemetryQueryResultSeries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumns

`func (o *TelemetryQueryResultSeries) GetColumns() []string`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *TelemetryQueryResultSeries) GetColumnsOk() (*[]string, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *TelemetryQueryResultSeries) SetColumns(v []string)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *TelemetryQueryResultSeries) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetName

`func (o *TelemetryQueryResultSeries) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TelemetryQueryResultSeries) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TelemetryQueryResultSeries) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TelemetryQueryResultSeries) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTags

`func (o *TelemetryQueryResultSeries) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TelemetryQueryResultSeries) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TelemetryQueryResultSeries) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TelemetryQueryResultSeries) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetValues

`func (o *TelemetryQueryResultSeries) GetValues() []ApiInterfaceSlice`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *TelemetryQueryResultSeries) GetValuesOk() (*[]ApiInterfaceSlice, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *TelemetryQueryResultSeries) SetValues(v []ApiInterfaceSlice)`

SetValues sets Values field to given value.

### HasValues

`func (o *TelemetryQueryResultSeries) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


