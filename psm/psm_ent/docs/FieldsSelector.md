# FieldsSelector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Requirements** | Pointer to [**[]FieldsRequirement**](FieldsRequirement.md) | Requirements are ANDed. | [optional] 

## Methods

### NewFieldsSelector

`func NewFieldsSelector() *FieldsSelector`

NewFieldsSelector instantiates a new FieldsSelector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldsSelectorWithDefaults

`func NewFieldsSelectorWithDefaults() *FieldsSelector`

NewFieldsSelectorWithDefaults instantiates a new FieldsSelector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequirements

`func (o *FieldsSelector) GetRequirements() []FieldsRequirement`

GetRequirements returns the Requirements field if non-nil, zero value otherwise.

### GetRequirementsOk

`func (o *FieldsSelector) GetRequirementsOk() (*[]FieldsRequirement, bool)`

GetRequirementsOk returns a tuple with the Requirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirements

`func (o *FieldsSelector) SetRequirements(v []FieldsRequirement)`

SetRequirements sets Requirements field to given value.

### HasRequirements

`func (o *FieldsSelector) HasRequirements() bool`

HasRequirements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


