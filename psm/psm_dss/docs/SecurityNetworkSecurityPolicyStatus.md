# SecurityNetworkSecurityPolicyStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PropagationStatus** | Pointer to [**SecurityPropagationStatus**](securityPropagationStatus.md) |  | [optional] 
**RuleStatus** | Pointer to [**[]SecuritySGRuleStatus**](SecuritySGRuleStatus.md) |  | [optional] 

## Methods

### NewSecurityNetworkSecurityPolicyStatus

`func NewSecurityNetworkSecurityPolicyStatus() *SecurityNetworkSecurityPolicyStatus`

NewSecurityNetworkSecurityPolicyStatus instantiates a new SecurityNetworkSecurityPolicyStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityNetworkSecurityPolicyStatusWithDefaults

`func NewSecurityNetworkSecurityPolicyStatusWithDefaults() *SecurityNetworkSecurityPolicyStatus`

NewSecurityNetworkSecurityPolicyStatusWithDefaults instantiates a new SecurityNetworkSecurityPolicyStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPropagationStatus

`func (o *SecurityNetworkSecurityPolicyStatus) GetPropagationStatus() SecurityPropagationStatus`

GetPropagationStatus returns the PropagationStatus field if non-nil, zero value otherwise.

### GetPropagationStatusOk

`func (o *SecurityNetworkSecurityPolicyStatus) GetPropagationStatusOk() (*SecurityPropagationStatus, bool)`

GetPropagationStatusOk returns a tuple with the PropagationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropagationStatus

`func (o *SecurityNetworkSecurityPolicyStatus) SetPropagationStatus(v SecurityPropagationStatus)`

SetPropagationStatus sets PropagationStatus field to given value.

### HasPropagationStatus

`func (o *SecurityNetworkSecurityPolicyStatus) HasPropagationStatus() bool`

HasPropagationStatus returns a boolean if a field has been set.

### GetRuleStatus

`func (o *SecurityNetworkSecurityPolicyStatus) GetRuleStatus() []SecuritySGRuleStatus`

GetRuleStatus returns the RuleStatus field if non-nil, zero value otherwise.

### GetRuleStatusOk

`func (o *SecurityNetworkSecurityPolicyStatus) GetRuleStatusOk() (*[]SecuritySGRuleStatus, bool)`

GetRuleStatusOk returns a tuple with the RuleStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleStatus

`func (o *SecurityNetworkSecurityPolicyStatus) SetRuleStatus(v []SecuritySGRuleStatus)`

SetRuleStatus sets RuleStatus field to given value.

### HasRuleStatus

`func (o *SecurityNetworkSecurityPolicyStatus) HasRuleStatus() bool`

HasRuleStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


