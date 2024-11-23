# LabelsSelector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Requirements** | Pointer to [**[]LabelsRequirement**](LabelsRequirement.md) | Requirements are ANDed. | [optional] 

## Methods

### NewLabelsSelector

`func NewLabelsSelector() *LabelsSelector`

NewLabelsSelector instantiates a new LabelsSelector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabelsSelectorWithDefaults

`func NewLabelsSelectorWithDefaults() *LabelsSelector`

NewLabelsSelectorWithDefaults instantiates a new LabelsSelector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequirements

`func (o *LabelsSelector) GetRequirements() []LabelsRequirement`

GetRequirements returns the Requirements field if non-nil, zero value otherwise.

### GetRequirementsOk

`func (o *LabelsSelector) GetRequirementsOk() (*[]LabelsRequirement, bool)`

GetRequirementsOk returns a tuple with the Requirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirements

`func (o *LabelsSelector) SetRequirements(v []LabelsRequirement)`

SetRequirements sets Requirements field to given value.

### HasRequirements

`func (o *LabelsSelector) HasRequirements() bool`

HasRequirements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


