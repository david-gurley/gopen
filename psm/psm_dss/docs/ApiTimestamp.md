# ApiTimestamp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nanos** | Pointer to **int32** |  | [optional] 
**Seconds** | Pointer to **string** |  | [optional] 

## Methods

### NewApiTimestamp

`func NewApiTimestamp() *ApiTimestamp`

NewApiTimestamp instantiates a new ApiTimestamp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiTimestampWithDefaults

`func NewApiTimestampWithDefaults() *ApiTimestamp`

NewApiTimestampWithDefaults instantiates a new ApiTimestamp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNanos

`func (o *ApiTimestamp) GetNanos() int32`

GetNanos returns the Nanos field if non-nil, zero value otherwise.

### GetNanosOk

`func (o *ApiTimestamp) GetNanosOk() (*int32, bool)`

GetNanosOk returns a tuple with the Nanos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNanos

`func (o *ApiTimestamp) SetNanos(v int32)`

SetNanos sets Nanos field to given value.

### HasNanos

`func (o *ApiTimestamp) HasNanos() bool`

HasNanos returns a boolean if a field has been set.

### GetSeconds

`func (o *ApiTimestamp) GetSeconds() string`

GetSeconds returns the Seconds field if non-nil, zero value otherwise.

### GetSecondsOk

`func (o *ApiTimestamp) GetSecondsOk() (*string, bool)`

GetSecondsOk returns a tuple with the Seconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeconds

`func (o *ApiTimestamp) SetSeconds(v string)`

SetSeconds sets Seconds field to given value.

### HasSeconds

`func (o *ApiTimestamp) HasSeconds() bool`

HasSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


