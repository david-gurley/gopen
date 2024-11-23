# SecurityNetworkSecurityPolicySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttachGroups** | Pointer to **[]string** | list of security groups this policy is attached to. | [optional] 
**AttachTenant** | Pointer to **bool** | specifies if the set of rules need to be attached globally to a tenant. | [optional] 
**Priority** | Pointer to **int64** | Policy priority. If not specified, it will be assigned automatically in increments of 1000. | [optional] 
**Rules** | Pointer to [**[]SecuritySGRule**](SecuritySGRule.md) | list of rules. | [optional] 

## Methods

### NewSecurityNetworkSecurityPolicySpec

`func NewSecurityNetworkSecurityPolicySpec() *SecurityNetworkSecurityPolicySpec`

NewSecurityNetworkSecurityPolicySpec instantiates a new SecurityNetworkSecurityPolicySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityNetworkSecurityPolicySpecWithDefaults

`func NewSecurityNetworkSecurityPolicySpecWithDefaults() *SecurityNetworkSecurityPolicySpec`

NewSecurityNetworkSecurityPolicySpecWithDefaults instantiates a new SecurityNetworkSecurityPolicySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachGroups

`func (o *SecurityNetworkSecurityPolicySpec) GetAttachGroups() []string`

GetAttachGroups returns the AttachGroups field if non-nil, zero value otherwise.

### GetAttachGroupsOk

`func (o *SecurityNetworkSecurityPolicySpec) GetAttachGroupsOk() (*[]string, bool)`

GetAttachGroupsOk returns a tuple with the AttachGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachGroups

`func (o *SecurityNetworkSecurityPolicySpec) SetAttachGroups(v []string)`

SetAttachGroups sets AttachGroups field to given value.

### HasAttachGroups

`func (o *SecurityNetworkSecurityPolicySpec) HasAttachGroups() bool`

HasAttachGroups returns a boolean if a field has been set.

### GetAttachTenant

`func (o *SecurityNetworkSecurityPolicySpec) GetAttachTenant() bool`

GetAttachTenant returns the AttachTenant field if non-nil, zero value otherwise.

### GetAttachTenantOk

`func (o *SecurityNetworkSecurityPolicySpec) GetAttachTenantOk() (*bool, bool)`

GetAttachTenantOk returns a tuple with the AttachTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachTenant

`func (o *SecurityNetworkSecurityPolicySpec) SetAttachTenant(v bool)`

SetAttachTenant sets AttachTenant field to given value.

### HasAttachTenant

`func (o *SecurityNetworkSecurityPolicySpec) HasAttachTenant() bool`

HasAttachTenant returns a boolean if a field has been set.

### GetPriority

`func (o *SecurityNetworkSecurityPolicySpec) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *SecurityNetworkSecurityPolicySpec) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *SecurityNetworkSecurityPolicySpec) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *SecurityNetworkSecurityPolicySpec) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetRules

`func (o *SecurityNetworkSecurityPolicySpec) GetRules() []SecuritySGRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *SecurityNetworkSecurityPolicySpec) GetRulesOk() (*[]SecuritySGRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *SecurityNetworkSecurityPolicySpec) SetRules(v []SecuritySGRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *SecurityNetworkSecurityPolicySpec) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


