# TelemetryQueryTopSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TopN** | Pointer to **int64** | TopN defines how many entries returned for top aggregation function and by default it is 10. | [optional] 

## Methods

### NewTelemetryQueryTopSpec

`func NewTelemetryQueryTopSpec() *TelemetryQueryTopSpec`

NewTelemetryQueryTopSpec instantiates a new TelemetryQueryTopSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryQueryTopSpecWithDefaults

`func NewTelemetryQueryTopSpecWithDefaults() *TelemetryQueryTopSpec`

NewTelemetryQueryTopSpecWithDefaults instantiates a new TelemetryQueryTopSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopN

`func (o *TelemetryQueryTopSpec) GetTopN() int64`

GetTopN returns the TopN field if non-nil, zero value otherwise.

### GetTopNOk

`func (o *TelemetryQueryTopSpec) GetTopNOk() (*int64, bool)`

GetTopNOk returns a tuple with the TopN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopN

`func (o *TelemetryQueryTopSpec) SetTopN(v int64)`

SetTopN sets TopN field to given value.

### HasTopN

`func (o *TelemetryQueryTopSpec) HasTopN() bool`

HasTopN returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


