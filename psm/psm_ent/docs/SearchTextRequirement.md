# SearchTextRequirement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | Pointer to **[]string** | AND of words or phrases to be matched The max text-string length is 256 bytes. Length of string should be between 0 and 256. | [optional] 

## Methods

### NewSearchTextRequirement

`func NewSearchTextRequirement() *SearchTextRequirement`

NewSearchTextRequirement instantiates a new SearchTextRequirement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchTextRequirementWithDefaults

`func NewSearchTextRequirementWithDefaults() *SearchTextRequirement`

NewSearchTextRequirementWithDefaults instantiates a new SearchTextRequirement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *SearchTextRequirement) GetText() []string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *SearchTextRequirement) GetTextOk() (*[]string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *SearchTextRequirement) SetText(v []string)`

SetText sets Text field to given value.

### HasText

`func (o *SearchTextRequirement) HasText() bool`

HasText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


