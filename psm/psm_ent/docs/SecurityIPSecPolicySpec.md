# SecurityIPSecPolicySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**SecurityIPSecConfig**](securityIPSecConfig.md) |  | [optional] 
**Rules** | Pointer to [**[]SecurityIPSecPolicyRule**](SecurityIPSecPolicyRule.md) | list of rules. | [optional] 

## Methods

### NewSecurityIPSecPolicySpec

`func NewSecurityIPSecPolicySpec() *SecurityIPSecPolicySpec`

NewSecurityIPSecPolicySpec instantiates a new SecurityIPSecPolicySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityIPSecPolicySpecWithDefaults

`func NewSecurityIPSecPolicySpecWithDefaults() *SecurityIPSecPolicySpec`

NewSecurityIPSecPolicySpecWithDefaults instantiates a new SecurityIPSecPolicySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *SecurityIPSecPolicySpec) GetConfig() SecurityIPSecConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *SecurityIPSecPolicySpec) GetConfigOk() (*SecurityIPSecConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *SecurityIPSecPolicySpec) SetConfig(v SecurityIPSecConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *SecurityIPSecPolicySpec) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetRules

`func (o *SecurityIPSecPolicySpec) GetRules() []SecurityIPSecPolicyRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *SecurityIPSecPolicySpec) GetRulesOk() (*[]SecurityIPSecPolicyRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *SecurityIPSecPolicySpec) SetRules(v []SecurityIPSecPolicyRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *SecurityIPSecPolicySpec) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


