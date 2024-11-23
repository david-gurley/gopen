# MonitoringMatchedRequirement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** |  | [optional] 
**ObservedValue** | Pointer to **string** | The value at which the requirement was met. same as Requirement.value for operator &#x60;Equals&#x60; but could vary for other operators e.g. requirement - CPU;Gt;90 could have a matching value 96. | [optional] 
**Operator** | Pointer to **string** |  | [optional] [default to "equals"]
**Values** | Pointer to **[]string** |  | [optional] 

## Methods

### NewMonitoringMatchedRequirement

`func NewMonitoringMatchedRequirement() *MonitoringMatchedRequirement`

NewMonitoringMatchedRequirement instantiates a new MonitoringMatchedRequirement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringMatchedRequirementWithDefaults

`func NewMonitoringMatchedRequirementWithDefaults() *MonitoringMatchedRequirement`

NewMonitoringMatchedRequirementWithDefaults instantiates a new MonitoringMatchedRequirement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *MonitoringMatchedRequirement) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MonitoringMatchedRequirement) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MonitoringMatchedRequirement) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *MonitoringMatchedRequirement) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetObservedValue

`func (o *MonitoringMatchedRequirement) GetObservedValue() string`

GetObservedValue returns the ObservedValue field if non-nil, zero value otherwise.

### GetObservedValueOk

`func (o *MonitoringMatchedRequirement) GetObservedValueOk() (*string, bool)`

GetObservedValueOk returns a tuple with the ObservedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedValue

`func (o *MonitoringMatchedRequirement) SetObservedValue(v string)`

SetObservedValue sets ObservedValue field to given value.

### HasObservedValue

`func (o *MonitoringMatchedRequirement) HasObservedValue() bool`

HasObservedValue returns a boolean if a field has been set.

### GetOperator

`func (o *MonitoringMatchedRequirement) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *MonitoringMatchedRequirement) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *MonitoringMatchedRequirement) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *MonitoringMatchedRequirement) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValues

`func (o *MonitoringMatchedRequirement) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *MonitoringMatchedRequirement) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *MonitoringMatchedRequirement) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *MonitoringMatchedRequirement) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


