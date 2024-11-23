# TelemetryQueryBottomSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BottomN** | Pointer to **int64** | BottomN defines how many entries returned for bottom aggregation function and by default it is 10. | [optional] 

## Methods

### NewTelemetryQueryBottomSpec

`func NewTelemetryQueryBottomSpec() *TelemetryQueryBottomSpec`

NewTelemetryQueryBottomSpec instantiates a new TelemetryQueryBottomSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryQueryBottomSpecWithDefaults

`func NewTelemetryQueryBottomSpecWithDefaults() *TelemetryQueryBottomSpec`

NewTelemetryQueryBottomSpecWithDefaults instantiates a new TelemetryQueryBottomSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBottomN

`func (o *TelemetryQueryBottomSpec) GetBottomN() int64`

GetBottomN returns the BottomN field if non-nil, zero value otherwise.

### GetBottomNOk

`func (o *TelemetryQueryBottomSpec) GetBottomNOk() (*int64, bool)`

GetBottomNOk returns a tuple with the BottomN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBottomN

`func (o *TelemetryQueryBottomSpec) SetBottomN(v int64)`

SetBottomN sets BottomN field to given value.

### HasBottomN

`func (o *TelemetryQueryBottomSpec) HasBottomN() bool`

HasBottomN returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


