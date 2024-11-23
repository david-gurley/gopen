# SecuritySGRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | SGRule action, either PERMIT, DENY or REJECT. | [optional] [default to "permit"]
**Apps** | Pointer to **[]string** | list of apps objects to which the rule applies to. | [optional] 
**Description** | Pointer to **string** | describes rule. Length of string should be between 0 and 256. | [optional] 
**FromIpAddresses** | Pointer to **[]string** | inbound rule from a given ip-address/ip-mask/ip-range. Use any to refer to all ipaddresses cli-tags: id&#x3D;from-ip. | [optional] 
**FromSecurityGroups** | Pointer to **[]string** | inbound rule from a given security group. | [optional] 
**ProtoPorts** | Pointer to [**[]SecurityProtoPort**](SecurityProtoPort.md) | list of (protocol, ports) pairs to which the rule applies to, in addition to apps. | [optional] 
**ToIpAddresses** | Pointer to **[]string** | outbound rule from a given ip-address/ip-mask/ip-range. Use any to refer to all ipaddresses cli-tags: id&#x3D;to-ip. | [optional] 
**ToSecurityGroups** | Pointer to **[]string** | outbound rule from a given security group. | [optional] 

## Methods

### NewSecuritySGRule

`func NewSecuritySGRule() *SecuritySGRule`

NewSecuritySGRule instantiates a new SecuritySGRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecuritySGRuleWithDefaults

`func NewSecuritySGRuleWithDefaults() *SecuritySGRule`

NewSecuritySGRuleWithDefaults instantiates a new SecuritySGRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *SecuritySGRule) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SecuritySGRule) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SecuritySGRule) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *SecuritySGRule) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetApps

`func (o *SecuritySGRule) GetApps() []string`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *SecuritySGRule) GetAppsOk() (*[]string, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *SecuritySGRule) SetApps(v []string)`

SetApps sets Apps field to given value.

### HasApps

`func (o *SecuritySGRule) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetDescription

`func (o *SecuritySGRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SecuritySGRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SecuritySGRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SecuritySGRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFromIpAddresses

`func (o *SecuritySGRule) GetFromIpAddresses() []string`

GetFromIpAddresses returns the FromIpAddresses field if non-nil, zero value otherwise.

### GetFromIpAddressesOk

`func (o *SecuritySGRule) GetFromIpAddressesOk() (*[]string, bool)`

GetFromIpAddressesOk returns a tuple with the FromIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromIpAddresses

`func (o *SecuritySGRule) SetFromIpAddresses(v []string)`

SetFromIpAddresses sets FromIpAddresses field to given value.

### HasFromIpAddresses

`func (o *SecuritySGRule) HasFromIpAddresses() bool`

HasFromIpAddresses returns a boolean if a field has been set.

### GetFromSecurityGroups

`func (o *SecuritySGRule) GetFromSecurityGroups() []string`

GetFromSecurityGroups returns the FromSecurityGroups field if non-nil, zero value otherwise.

### GetFromSecurityGroupsOk

`func (o *SecuritySGRule) GetFromSecurityGroupsOk() (*[]string, bool)`

GetFromSecurityGroupsOk returns a tuple with the FromSecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromSecurityGroups

`func (o *SecuritySGRule) SetFromSecurityGroups(v []string)`

SetFromSecurityGroups sets FromSecurityGroups field to given value.

### HasFromSecurityGroups

`func (o *SecuritySGRule) HasFromSecurityGroups() bool`

HasFromSecurityGroups returns a boolean if a field has been set.

### GetProtoPorts

`func (o *SecuritySGRule) GetProtoPorts() []SecurityProtoPort`

GetProtoPorts returns the ProtoPorts field if non-nil, zero value otherwise.

### GetProtoPortsOk

`func (o *SecuritySGRule) GetProtoPortsOk() (*[]SecurityProtoPort, bool)`

GetProtoPortsOk returns a tuple with the ProtoPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtoPorts

`func (o *SecuritySGRule) SetProtoPorts(v []SecurityProtoPort)`

SetProtoPorts sets ProtoPorts field to given value.

### HasProtoPorts

`func (o *SecuritySGRule) HasProtoPorts() bool`

HasProtoPorts returns a boolean if a field has been set.

### GetToIpAddresses

`func (o *SecuritySGRule) GetToIpAddresses() []string`

GetToIpAddresses returns the ToIpAddresses field if non-nil, zero value otherwise.

### GetToIpAddressesOk

`func (o *SecuritySGRule) GetToIpAddressesOk() (*[]string, bool)`

GetToIpAddressesOk returns a tuple with the ToIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToIpAddresses

`func (o *SecuritySGRule) SetToIpAddresses(v []string)`

SetToIpAddresses sets ToIpAddresses field to given value.

### HasToIpAddresses

`func (o *SecuritySGRule) HasToIpAddresses() bool`

HasToIpAddresses returns a boolean if a field has been set.

### GetToSecurityGroups

`func (o *SecuritySGRule) GetToSecurityGroups() []string`

GetToSecurityGroups returns the ToSecurityGroups field if non-nil, zero value otherwise.

### GetToSecurityGroupsOk

`func (o *SecuritySGRule) GetToSecurityGroupsOk() (*[]string, bool)`

GetToSecurityGroupsOk returns a tuple with the ToSecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToSecurityGroups

`func (o *SecuritySGRule) SetToSecurityGroups(v []string)`

SetToSecurityGroups sets ToSecurityGroups field to given value.

### HasToSecurityGroups

`func (o *SecuritySGRule) HasToSecurityGroups() bool`

HasToSecurityGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


