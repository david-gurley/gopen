# StagingItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | Pointer to **string** |  | [optional] 
**Object** | Pointer to [**ApiAny**](apiAny.md) |  | [optional] 
**Uri** | Pointer to **string** |  | [optional] 

## Methods

### NewStagingItem

`func NewStagingItem() *StagingItem`

NewStagingItem instantiates a new StagingItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStagingItemWithDefaults

`func NewStagingItemWithDefaults() *StagingItem`

NewStagingItemWithDefaults instantiates a new StagingItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *StagingItem) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *StagingItem) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *StagingItem) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *StagingItem) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetObject

`func (o *StagingItem) GetObject() ApiAny`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *StagingItem) GetObjectOk() (*ApiAny, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *StagingItem) SetObject(v ApiAny)`

SetObject sets Object field to given value.

### HasObject

`func (o *StagingItem) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetUri

`func (o *StagingItem) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *StagingItem) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *StagingItem) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *StagingItem) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


