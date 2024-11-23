# SecurityIPSecPolicyRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | IPSec action Type for  all traffic that matches the MatchRule. This policy rule, either CLEAR or PROTECT. | [optional] [default to "clear"]
**Description** | Pointer to **string** | Describes rule. Length of string should be between 0 and 256. | [optional] 
**MatchRule** | Pointer to [**SecurityIPSecMatchRule**](securityIPSecMatchRule.md) |  | [optional] 

## Methods

### NewSecurityIPSecPolicyRule

`func NewSecurityIPSecPolicyRule() *SecurityIPSecPolicyRule`

NewSecurityIPSecPolicyRule instantiates a new SecurityIPSecPolicyRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityIPSecPolicyRuleWithDefaults

`func NewSecurityIPSecPolicyRuleWithDefaults() *SecurityIPSecPolicyRule`

NewSecurityIPSecPolicyRuleWithDefaults instantiates a new SecurityIPSecPolicyRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *SecurityIPSecPolicyRule) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SecurityIPSecPolicyRule) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SecurityIPSecPolicyRule) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *SecurityIPSecPolicyRule) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetDescription

`func (o *SecurityIPSecPolicyRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SecurityIPSecPolicyRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SecurityIPSecPolicyRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SecurityIPSecPolicyRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMatchRule

`func (o *SecurityIPSecPolicyRule) GetMatchRule() SecurityIPSecMatchRule`

GetMatchRule returns the MatchRule field if non-nil, zero value otherwise.

### GetMatchRuleOk

`func (o *SecurityIPSecPolicyRule) GetMatchRuleOk() (*SecurityIPSecMatchRule, bool)`

GetMatchRuleOk returns a tuple with the MatchRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchRule

`func (o *SecurityIPSecPolicyRule) SetMatchRule(v SecurityIPSecMatchRule)`

SetMatchRule sets MatchRule field to given value.

### HasMatchRule

`func (o *SecurityIPSecPolicyRule) HasMatchRule() bool`

HasMatchRule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


