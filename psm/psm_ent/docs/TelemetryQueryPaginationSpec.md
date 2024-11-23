# TelemetryQueryPaginationSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Count specifies the number of points to include. Value should be at least 1. | [optional] 
**Offset** | Pointer to **int32** | Offset specifies the starting point when using Count. Value should be at least 0. | [optional] [default to 0]

## Methods

### NewTelemetryQueryPaginationSpec

`func NewTelemetryQueryPaginationSpec() *TelemetryQueryPaginationSpec`

NewTelemetryQueryPaginationSpec instantiates a new TelemetryQueryPaginationSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryQueryPaginationSpecWithDefaults

`func NewTelemetryQueryPaginationSpecWithDefaults() *TelemetryQueryPaginationSpec`

NewTelemetryQueryPaginationSpecWithDefaults instantiates a new TelemetryQueryPaginationSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *TelemetryQueryPaginationSpec) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TelemetryQueryPaginationSpec) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *TelemetryQueryPaginationSpec) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *TelemetryQueryPaginationSpec) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetOffset

`func (o *TelemetryQueryPaginationSpec) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *TelemetryQueryPaginationSpec) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *TelemetryQueryPaginationSpec) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *TelemetryQueryPaginationSpec) HasOffset() bool`

HasOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


