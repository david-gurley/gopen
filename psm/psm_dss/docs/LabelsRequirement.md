# LabelsRequirement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | The label key that the condition applies to. | [optional] 
**Operator** | Pointer to **string** | Condition checked for the key. | [optional] [default to "equals"]
**Values** | Pointer to **[]string** | Values contains one or more values corresponding to the label key. \&quot;equals\&quot; and \&quot;notEquals\&quot; operators need a single Value. \&quot;in\&quot; and \&quot;notIn\&quot; operators can have one or more values. | [optional] 

## Methods

### NewLabelsRequirement

`func NewLabelsRequirement() *LabelsRequirement`

NewLabelsRequirement instantiates a new LabelsRequirement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabelsRequirementWithDefaults

`func NewLabelsRequirementWithDefaults() *LabelsRequirement`

NewLabelsRequirementWithDefaults instantiates a new LabelsRequirement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *LabelsRequirement) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *LabelsRequirement) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *LabelsRequirement) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *LabelsRequirement) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetOperator

`func (o *LabelsRequirement) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *LabelsRequirement) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *LabelsRequirement) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *LabelsRequirement) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValues

`func (o *LabelsRequirement) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *LabelsRequirement) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *LabelsRequirement) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *LabelsRequirement) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


