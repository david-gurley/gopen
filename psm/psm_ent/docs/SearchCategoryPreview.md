# SearchCategoryPreview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Categories** | Pointer to [**map[string]SearchKindPreview**](searchKindPreview.md) |  | [optional] 

## Methods

### NewSearchCategoryPreview

`func NewSearchCategoryPreview() *SearchCategoryPreview`

NewSearchCategoryPreview instantiates a new SearchCategoryPreview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchCategoryPreviewWithDefaults

`func NewSearchCategoryPreviewWithDefaults() *SearchCategoryPreview`

NewSearchCategoryPreviewWithDefaults instantiates a new SearchCategoryPreview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategories

`func (o *SearchCategoryPreview) GetCategories() map[string]SearchKindPreview`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *SearchCategoryPreview) GetCategoriesOk() (*map[string]SearchKindPreview, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *SearchCategoryPreview) SetCategories(v map[string]SearchKindPreview)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *SearchCategoryPreview) HasCategories() bool`

HasCategories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

