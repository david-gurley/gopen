# SearchSearchQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Categories** | Pointer to **[]string** | OR of Categories to be matched, AND and Exclude are not supported for this type The max category string length is 64 bytes. Length of string should be between 0 and 64. | [optional] 
**Fields** | Pointer to [**FieldsSelector**](fieldsSelector.md) |  | [optional] 
**Kinds** | Pointer to **[]string** | OR of Kinds to be matched, AND and Exclude are not supported for this type. Should be a valid object Kind. | [optional] 
**Labels** | Pointer to [**LabelsSelector**](labelsSelector.md) |  | [optional] 
**Texts** | Pointer to [**[]SearchTextRequirement**](SearchTextRequirement.md) | OR of Text-requirements to be matched, Exclude is not supported for Text search. | [optional] 

## Methods

### NewSearchSearchQuery

`func NewSearchSearchQuery() *SearchSearchQuery`

NewSearchSearchQuery instantiates a new SearchSearchQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchSearchQueryWithDefaults

`func NewSearchSearchQueryWithDefaults() *SearchSearchQuery`

NewSearchSearchQueryWithDefaults instantiates a new SearchSearchQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategories

`func (o *SearchSearchQuery) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *SearchSearchQuery) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *SearchSearchQuery) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *SearchSearchQuery) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetFields

`func (o *SearchSearchQuery) GetFields() FieldsSelector`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *SearchSearchQuery) GetFieldsOk() (*FieldsSelector, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *SearchSearchQuery) SetFields(v FieldsSelector)`

SetFields sets Fields field to given value.

### HasFields

`func (o *SearchSearchQuery) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetKinds

`func (o *SearchSearchQuery) GetKinds() []string`

GetKinds returns the Kinds field if non-nil, zero value otherwise.

### GetKindsOk

`func (o *SearchSearchQuery) GetKindsOk() (*[]string, bool)`

GetKindsOk returns a tuple with the Kinds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKinds

`func (o *SearchSearchQuery) SetKinds(v []string)`

SetKinds sets Kinds field to given value.

### HasKinds

`func (o *SearchSearchQuery) HasKinds() bool`

HasKinds returns a boolean if a field has been set.

### GetLabels

`func (o *SearchSearchQuery) GetLabels() LabelsSelector`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *SearchSearchQuery) GetLabelsOk() (*LabelsSelector, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *SearchSearchQuery) SetLabels(v LabelsSelector)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *SearchSearchQuery) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetTexts

`func (o *SearchSearchQuery) GetTexts() []SearchTextRequirement`

GetTexts returns the Texts field if non-nil, zero value otherwise.

### GetTextsOk

`func (o *SearchSearchQuery) GetTextsOk() (*[]SearchTextRequirement, bool)`

GetTextsOk returns a tuple with the Texts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTexts

`func (o *SearchSearchQuery) SetTexts(v []SearchTextRequirement)`

SetTexts sets Texts field to given value.

### HasTexts

`func (o *SearchSearchQuery) HasTexts() bool`

HasTexts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


