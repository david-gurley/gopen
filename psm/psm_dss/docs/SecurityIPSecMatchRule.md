# SecurityIPSecMatchRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apps** | Pointer to **[]string** | list of apps objects to which the rule applies to. | [optional] 
**DstIpAddresses** | Pointer to **[]string** | outbound rule from a given ip-address/ip-mask/ip-range. Use any to refer to all ipaddresses cli-tags: id&#x3D;to-ip. | [optional] 
**ProtoPorts** | Pointer to [**[]SecurityProtoPort**](SecurityProtoPort.md) | list of (protocol, ports) pairs to which the rule applies to, in addition to apps. | [optional] 
**SrcIpAddresses** | Pointer to **[]string** | inbound rule from a given ip-address/ip-mask/ip-range. Use any to refer to all ipaddresses cli-tags: id&#x3D;from-ip. | [optional] 

## Methods

### NewSecurityIPSecMatchRule

`func NewSecurityIPSecMatchRule() *SecurityIPSecMatchRule`

NewSecurityIPSecMatchRule instantiates a new SecurityIPSecMatchRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityIPSecMatchRuleWithDefaults

`func NewSecurityIPSecMatchRuleWithDefaults() *SecurityIPSecMatchRule`

NewSecurityIPSecMatchRuleWithDefaults instantiates a new SecurityIPSecMatchRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApps

`func (o *SecurityIPSecMatchRule) GetApps() []string`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *SecurityIPSecMatchRule) GetAppsOk() (*[]string, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *SecurityIPSecMatchRule) SetApps(v []string)`

SetApps sets Apps field to given value.

### HasApps

`func (o *SecurityIPSecMatchRule) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetDstIpAddresses

`func (o *SecurityIPSecMatchRule) GetDstIpAddresses() []string`

GetDstIpAddresses returns the DstIpAddresses field if non-nil, zero value otherwise.

### GetDstIpAddressesOk

`func (o *SecurityIPSecMatchRule) GetDstIpAddressesOk() (*[]string, bool)`

GetDstIpAddressesOk returns a tuple with the DstIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstIpAddresses

`func (o *SecurityIPSecMatchRule) SetDstIpAddresses(v []string)`

SetDstIpAddresses sets DstIpAddresses field to given value.

### HasDstIpAddresses

`func (o *SecurityIPSecMatchRule) HasDstIpAddresses() bool`

HasDstIpAddresses returns a boolean if a field has been set.

### GetProtoPorts

`func (o *SecurityIPSecMatchRule) GetProtoPorts() []SecurityProtoPort`

GetProtoPorts returns the ProtoPorts field if non-nil, zero value otherwise.

### GetProtoPortsOk

`func (o *SecurityIPSecMatchRule) GetProtoPortsOk() (*[]SecurityProtoPort, bool)`

GetProtoPortsOk returns a tuple with the ProtoPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtoPorts

`func (o *SecurityIPSecMatchRule) SetProtoPorts(v []SecurityProtoPort)`

SetProtoPorts sets ProtoPorts field to given value.

### HasProtoPorts

`func (o *SecurityIPSecMatchRule) HasProtoPorts() bool`

HasProtoPorts returns a boolean if a field has been set.

### GetSrcIpAddresses

`func (o *SecurityIPSecMatchRule) GetSrcIpAddresses() []string`

GetSrcIpAddresses returns the SrcIpAddresses field if non-nil, zero value otherwise.

### GetSrcIpAddressesOk

`func (o *SecurityIPSecMatchRule) GetSrcIpAddressesOk() (*[]string, bool)`

GetSrcIpAddressesOk returns a tuple with the SrcIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcIpAddresses

`func (o *SecurityIPSecMatchRule) SetSrcIpAddresses(v []string)`

SetSrcIpAddresses sets SrcIpAddresses field to given value.

### HasSrcIpAddresses

`func (o *SecurityIPSecMatchRule) HasSrcIpAddresses() bool`

HasSrcIpAddresses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


