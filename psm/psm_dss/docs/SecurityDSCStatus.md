# SecurityDSCStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DscId** | Pointer to **string** | DSC ID for which the agent error or warning is issued. | [optional] 
**DscInfoStatus** | Pointer to **string** | InfoStatus contains agent message the operation is failed or warning is issued. | [optional] 
**PolicyEntriesConsumed** | Pointer to **string** | PolicyEntriesConsumed shows number of policy entries consumed in the DSC by this policy. | [optional] 
**RuleEntriesConsumed** | Pointer to **string** | RuleEntriesConsumed shows number of rule entries consumed in the DSE by this policy. | [optional] 

## Methods

### NewSecurityDSCStatus

`func NewSecurityDSCStatus() *SecurityDSCStatus`

NewSecurityDSCStatus instantiates a new SecurityDSCStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityDSCStatusWithDefaults

`func NewSecurityDSCStatusWithDefaults() *SecurityDSCStatus`

NewSecurityDSCStatusWithDefaults instantiates a new SecurityDSCStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDscId

`func (o *SecurityDSCStatus) GetDscId() string`

GetDscId returns the DscId field if non-nil, zero value otherwise.

### GetDscIdOk

`func (o *SecurityDSCStatus) GetDscIdOk() (*string, bool)`

GetDscIdOk returns a tuple with the DscId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDscId

`func (o *SecurityDSCStatus) SetDscId(v string)`

SetDscId sets DscId field to given value.

### HasDscId

`func (o *SecurityDSCStatus) HasDscId() bool`

HasDscId returns a boolean if a field has been set.

### GetDscInfoStatus

`func (o *SecurityDSCStatus) GetDscInfoStatus() string`

GetDscInfoStatus returns the DscInfoStatus field if non-nil, zero value otherwise.

### GetDscInfoStatusOk

`func (o *SecurityDSCStatus) GetDscInfoStatusOk() (*string, bool)`

GetDscInfoStatusOk returns a tuple with the DscInfoStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDscInfoStatus

`func (o *SecurityDSCStatus) SetDscInfoStatus(v string)`

SetDscInfoStatus sets DscInfoStatus field to given value.

### HasDscInfoStatus

`func (o *SecurityDSCStatus) HasDscInfoStatus() bool`

HasDscInfoStatus returns a boolean if a field has been set.

### GetPolicyEntriesConsumed

`func (o *SecurityDSCStatus) GetPolicyEntriesConsumed() string`

GetPolicyEntriesConsumed returns the PolicyEntriesConsumed field if non-nil, zero value otherwise.

### GetPolicyEntriesConsumedOk

`func (o *SecurityDSCStatus) GetPolicyEntriesConsumedOk() (*string, bool)`

GetPolicyEntriesConsumedOk returns a tuple with the PolicyEntriesConsumed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyEntriesConsumed

`func (o *SecurityDSCStatus) SetPolicyEntriesConsumed(v string)`

SetPolicyEntriesConsumed sets PolicyEntriesConsumed field to given value.

### HasPolicyEntriesConsumed

`func (o *SecurityDSCStatus) HasPolicyEntriesConsumed() bool`

HasPolicyEntriesConsumed returns a boolean if a field has been set.

### GetRuleEntriesConsumed

`func (o *SecurityDSCStatus) GetRuleEntriesConsumed() string`

GetRuleEntriesConsumed returns the RuleEntriesConsumed field if non-nil, zero value otherwise.

### GetRuleEntriesConsumedOk

`func (o *SecurityDSCStatus) GetRuleEntriesConsumedOk() (*string, bool)`

GetRuleEntriesConsumedOk returns a tuple with the RuleEntriesConsumed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleEntriesConsumed

`func (o *SecurityDSCStatus) SetRuleEntriesConsumed(v string)`

SetRuleEntriesConsumed sets RuleEntriesConsumed field to given value.

### HasRuleEntriesConsumed

`func (o *SecurityDSCStatus) HasRuleEntriesConsumed() bool`

HasRuleEntriesConsumed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


