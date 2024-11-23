# SecurityIPSecPolicyStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EspParams** | Pointer to **string** | Encryption and Algorithm details used for encrypting data traffic (ESP). | [optional] 
**IkeParams** | Pointer to **string** | Encryption and Algorithm details used for IKEv2 key exchange. | [optional] 
**PropagationStatus** | Pointer to [**SecurityPropagationStatus**](securityPropagationStatus.md) |  | [optional] 
**RuleStatus** | Pointer to [**[]SecurityIPSecRuleStatus**](SecurityIPSecRuleStatus.md) |  | [optional] 

## Methods

### NewSecurityIPSecPolicyStatus

`func NewSecurityIPSecPolicyStatus() *SecurityIPSecPolicyStatus`

NewSecurityIPSecPolicyStatus instantiates a new SecurityIPSecPolicyStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityIPSecPolicyStatusWithDefaults

`func NewSecurityIPSecPolicyStatusWithDefaults() *SecurityIPSecPolicyStatus`

NewSecurityIPSecPolicyStatusWithDefaults instantiates a new SecurityIPSecPolicyStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEspParams

`func (o *SecurityIPSecPolicyStatus) GetEspParams() string`

GetEspParams returns the EspParams field if non-nil, zero value otherwise.

### GetEspParamsOk

`func (o *SecurityIPSecPolicyStatus) GetEspParamsOk() (*string, bool)`

GetEspParamsOk returns a tuple with the EspParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEspParams

`func (o *SecurityIPSecPolicyStatus) SetEspParams(v string)`

SetEspParams sets EspParams field to given value.

### HasEspParams

`func (o *SecurityIPSecPolicyStatus) HasEspParams() bool`

HasEspParams returns a boolean if a field has been set.

### GetIkeParams

`func (o *SecurityIPSecPolicyStatus) GetIkeParams() string`

GetIkeParams returns the IkeParams field if non-nil, zero value otherwise.

### GetIkeParamsOk

`func (o *SecurityIPSecPolicyStatus) GetIkeParamsOk() (*string, bool)`

GetIkeParamsOk returns a tuple with the IkeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeParams

`func (o *SecurityIPSecPolicyStatus) SetIkeParams(v string)`

SetIkeParams sets IkeParams field to given value.

### HasIkeParams

`func (o *SecurityIPSecPolicyStatus) HasIkeParams() bool`

HasIkeParams returns a boolean if a field has been set.

### GetPropagationStatus

`func (o *SecurityIPSecPolicyStatus) GetPropagationStatus() SecurityPropagationStatus`

GetPropagationStatus returns the PropagationStatus field if non-nil, zero value otherwise.

### GetPropagationStatusOk

`func (o *SecurityIPSecPolicyStatus) GetPropagationStatusOk() (*SecurityPropagationStatus, bool)`

GetPropagationStatusOk returns a tuple with the PropagationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropagationStatus

`func (o *SecurityIPSecPolicyStatus) SetPropagationStatus(v SecurityPropagationStatus)`

SetPropagationStatus sets PropagationStatus field to given value.

### HasPropagationStatus

`func (o *SecurityIPSecPolicyStatus) HasPropagationStatus() bool`

HasPropagationStatus returns a boolean if a field has been set.

### GetRuleStatus

`func (o *SecurityIPSecPolicyStatus) GetRuleStatus() []SecurityIPSecRuleStatus`

GetRuleStatus returns the RuleStatus field if non-nil, zero value otherwise.

### GetRuleStatusOk

`func (o *SecurityIPSecPolicyStatus) GetRuleStatusOk() (*[]SecurityIPSecRuleStatus, bool)`

GetRuleStatusOk returns a tuple with the RuleStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleStatus

`func (o *SecurityIPSecPolicyStatus) SetRuleStatus(v []SecurityIPSecRuleStatus)`

SetRuleStatus sets RuleStatus field to given value.

### HasRuleStatus

`func (o *SecurityIPSecPolicyStatus) HasRuleStatus() bool`

HasRuleStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


