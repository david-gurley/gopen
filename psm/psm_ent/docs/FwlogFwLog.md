# FwlogFwLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Action. | [optional] [default to "allow"]
**Alg** | Pointer to **string** | Appliction Layer Gateway. | [optional] 
**ApiVersion** | Pointer to **string** |  | [optional] 
**AppId** | Pointer to **string** | Application ID. | [optional] 
**DestinationIp** | Pointer to **string** | Destination IP. | [optional] 
**DestinationPort** | Pointer to **int64** | Destination Port. | [optional] 
**DestinationVrf** | Pointer to **string** | Destination VRF,. | [optional] 
**Direction** | Pointer to **string** | Flow Direction. | [optional] [default to "from-host"]
**FlowAction** | Pointer to **string** | Flow action. | [optional] 
**IcmpCode** | Pointer to **int64** | icmp code. | [optional] 
**IcmpId** | Pointer to **int64** | icmp ID. | [optional] 
**IcmpType** | Pointer to **int64** | icmp type. | [optional] 
**IpsecProtected** | Pointer to **bool** | If IPsec protected. | [optional] 
**IpsecRuleId** | Pointer to **string** | IPsec policy rule ID. | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**ApiObjectMeta**](apiObjectMeta.md) |  | [optional] 
**PolicyName** | Pointer to **string** | policy name. | [optional] 
**Protocol** | Pointer to **string** | Protocol,. | [optional] 
**ReporterId** | Pointer to **string** | Reporter ID. | [optional] 
**RuleId** | Pointer to **string** | Rule ID. | [optional] 
**SessionId** | Pointer to **string** | Session ID. | [optional] 
**SourceIp** | Pointer to **string** | Source IP,. | [optional] 
**SourcePort** | Pointer to **int64** | Source Port. | [optional] 
**SourceVrf** | Pointer to **string** | Source VRF,. | [optional] 
**Vnid** | Pointer to **string** | VXLAN ID. | [optional] 

## Methods

### NewFwlogFwLog

`func NewFwlogFwLog() *FwlogFwLog`

NewFwlogFwLog instantiates a new FwlogFwLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFwlogFwLogWithDefaults

`func NewFwlogFwLogWithDefaults() *FwlogFwLog`

NewFwlogFwLogWithDefaults instantiates a new FwlogFwLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *FwlogFwLog) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *FwlogFwLog) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *FwlogFwLog) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *FwlogFwLog) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetAlg

`func (o *FwlogFwLog) GetAlg() string`

GetAlg returns the Alg field if non-nil, zero value otherwise.

### GetAlgOk

`func (o *FwlogFwLog) GetAlgOk() (*string, bool)`

GetAlgOk returns a tuple with the Alg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlg

`func (o *FwlogFwLog) SetAlg(v string)`

SetAlg sets Alg field to given value.

### HasAlg

`func (o *FwlogFwLog) HasAlg() bool`

HasAlg returns a boolean if a field has been set.

### GetApiVersion

`func (o *FwlogFwLog) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *FwlogFwLog) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *FwlogFwLog) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *FwlogFwLog) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetAppId

`func (o *FwlogFwLog) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *FwlogFwLog) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *FwlogFwLog) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *FwlogFwLog) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetDestinationIp

`func (o *FwlogFwLog) GetDestinationIp() string`

GetDestinationIp returns the DestinationIp field if non-nil, zero value otherwise.

### GetDestinationIpOk

`func (o *FwlogFwLog) GetDestinationIpOk() (*string, bool)`

GetDestinationIpOk returns a tuple with the DestinationIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationIp

`func (o *FwlogFwLog) SetDestinationIp(v string)`

SetDestinationIp sets DestinationIp field to given value.

### HasDestinationIp

`func (o *FwlogFwLog) HasDestinationIp() bool`

HasDestinationIp returns a boolean if a field has been set.

### GetDestinationPort

`func (o *FwlogFwLog) GetDestinationPort() int64`

GetDestinationPort returns the DestinationPort field if non-nil, zero value otherwise.

### GetDestinationPortOk

`func (o *FwlogFwLog) GetDestinationPortOk() (*int64, bool)`

GetDestinationPortOk returns a tuple with the DestinationPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPort

`func (o *FwlogFwLog) SetDestinationPort(v int64)`

SetDestinationPort sets DestinationPort field to given value.

### HasDestinationPort

`func (o *FwlogFwLog) HasDestinationPort() bool`

HasDestinationPort returns a boolean if a field has been set.

### GetDestinationVrf

`func (o *FwlogFwLog) GetDestinationVrf() string`

GetDestinationVrf returns the DestinationVrf field if non-nil, zero value otherwise.

### GetDestinationVrfOk

`func (o *FwlogFwLog) GetDestinationVrfOk() (*string, bool)`

GetDestinationVrfOk returns a tuple with the DestinationVrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationVrf

`func (o *FwlogFwLog) SetDestinationVrf(v string)`

SetDestinationVrf sets DestinationVrf field to given value.

### HasDestinationVrf

`func (o *FwlogFwLog) HasDestinationVrf() bool`

HasDestinationVrf returns a boolean if a field has been set.

### GetDirection

`func (o *FwlogFwLog) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *FwlogFwLog) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *FwlogFwLog) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *FwlogFwLog) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetFlowAction

`func (o *FwlogFwLog) GetFlowAction() string`

GetFlowAction returns the FlowAction field if non-nil, zero value otherwise.

### GetFlowActionOk

`func (o *FwlogFwLog) GetFlowActionOk() (*string, bool)`

GetFlowActionOk returns a tuple with the FlowAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowAction

`func (o *FwlogFwLog) SetFlowAction(v string)`

SetFlowAction sets FlowAction field to given value.

### HasFlowAction

`func (o *FwlogFwLog) HasFlowAction() bool`

HasFlowAction returns a boolean if a field has been set.

### GetIcmpCode

`func (o *FwlogFwLog) GetIcmpCode() int64`

GetIcmpCode returns the IcmpCode field if non-nil, zero value otherwise.

### GetIcmpCodeOk

`func (o *FwlogFwLog) GetIcmpCodeOk() (*int64, bool)`

GetIcmpCodeOk returns a tuple with the IcmpCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcmpCode

`func (o *FwlogFwLog) SetIcmpCode(v int64)`

SetIcmpCode sets IcmpCode field to given value.

### HasIcmpCode

`func (o *FwlogFwLog) HasIcmpCode() bool`

HasIcmpCode returns a boolean if a field has been set.

### GetIcmpId

`func (o *FwlogFwLog) GetIcmpId() int64`

GetIcmpId returns the IcmpId field if non-nil, zero value otherwise.

### GetIcmpIdOk

`func (o *FwlogFwLog) GetIcmpIdOk() (*int64, bool)`

GetIcmpIdOk returns a tuple with the IcmpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcmpId

`func (o *FwlogFwLog) SetIcmpId(v int64)`

SetIcmpId sets IcmpId field to given value.

### HasIcmpId

`func (o *FwlogFwLog) HasIcmpId() bool`

HasIcmpId returns a boolean if a field has been set.

### GetIcmpType

`func (o *FwlogFwLog) GetIcmpType() int64`

GetIcmpType returns the IcmpType field if non-nil, zero value otherwise.

### GetIcmpTypeOk

`func (o *FwlogFwLog) GetIcmpTypeOk() (*int64, bool)`

GetIcmpTypeOk returns a tuple with the IcmpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcmpType

`func (o *FwlogFwLog) SetIcmpType(v int64)`

SetIcmpType sets IcmpType field to given value.

### HasIcmpType

`func (o *FwlogFwLog) HasIcmpType() bool`

HasIcmpType returns a boolean if a field has been set.

### GetIpsecProtected

`func (o *FwlogFwLog) GetIpsecProtected() bool`

GetIpsecProtected returns the IpsecProtected field if non-nil, zero value otherwise.

### GetIpsecProtectedOk

`func (o *FwlogFwLog) GetIpsecProtectedOk() (*bool, bool)`

GetIpsecProtectedOk returns a tuple with the IpsecProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecProtected

`func (o *FwlogFwLog) SetIpsecProtected(v bool)`

SetIpsecProtected sets IpsecProtected field to given value.

### HasIpsecProtected

`func (o *FwlogFwLog) HasIpsecProtected() bool`

HasIpsecProtected returns a boolean if a field has been set.

### GetIpsecRuleId

`func (o *FwlogFwLog) GetIpsecRuleId() string`

GetIpsecRuleId returns the IpsecRuleId field if non-nil, zero value otherwise.

### GetIpsecRuleIdOk

`func (o *FwlogFwLog) GetIpsecRuleIdOk() (*string, bool)`

GetIpsecRuleIdOk returns a tuple with the IpsecRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecRuleId

`func (o *FwlogFwLog) SetIpsecRuleId(v string)`

SetIpsecRuleId sets IpsecRuleId field to given value.

### HasIpsecRuleId

`func (o *FwlogFwLog) HasIpsecRuleId() bool`

HasIpsecRuleId returns a boolean if a field has been set.

### GetKind

`func (o *FwlogFwLog) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *FwlogFwLog) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *FwlogFwLog) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *FwlogFwLog) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMeta

`func (o *FwlogFwLog) GetMeta() ApiObjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FwlogFwLog) GetMetaOk() (*ApiObjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FwlogFwLog) SetMeta(v ApiObjectMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *FwlogFwLog) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetPolicyName

`func (o *FwlogFwLog) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *FwlogFwLog) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *FwlogFwLog) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *FwlogFwLog) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### GetProtocol

`func (o *FwlogFwLog) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *FwlogFwLog) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *FwlogFwLog) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *FwlogFwLog) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetReporterId

`func (o *FwlogFwLog) GetReporterId() string`

GetReporterId returns the ReporterId field if non-nil, zero value otherwise.

### GetReporterIdOk

`func (o *FwlogFwLog) GetReporterIdOk() (*string, bool)`

GetReporterIdOk returns a tuple with the ReporterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReporterId

`func (o *FwlogFwLog) SetReporterId(v string)`

SetReporterId sets ReporterId field to given value.

### HasReporterId

`func (o *FwlogFwLog) HasReporterId() bool`

HasReporterId returns a boolean if a field has been set.

### GetRuleId

`func (o *FwlogFwLog) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *FwlogFwLog) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *FwlogFwLog) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *FwlogFwLog) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetSessionId

`func (o *FwlogFwLog) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *FwlogFwLog) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *FwlogFwLog) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *FwlogFwLog) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetSourceIp

`func (o *FwlogFwLog) GetSourceIp() string`

GetSourceIp returns the SourceIp field if non-nil, zero value otherwise.

### GetSourceIpOk

`func (o *FwlogFwLog) GetSourceIpOk() (*string, bool)`

GetSourceIpOk returns a tuple with the SourceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIp

`func (o *FwlogFwLog) SetSourceIp(v string)`

SetSourceIp sets SourceIp field to given value.

### HasSourceIp

`func (o *FwlogFwLog) HasSourceIp() bool`

HasSourceIp returns a boolean if a field has been set.

### GetSourcePort

`func (o *FwlogFwLog) GetSourcePort() int64`

GetSourcePort returns the SourcePort field if non-nil, zero value otherwise.

### GetSourcePortOk

`func (o *FwlogFwLog) GetSourcePortOk() (*int64, bool)`

GetSourcePortOk returns a tuple with the SourcePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePort

`func (o *FwlogFwLog) SetSourcePort(v int64)`

SetSourcePort sets SourcePort field to given value.

### HasSourcePort

`func (o *FwlogFwLog) HasSourcePort() bool`

HasSourcePort returns a boolean if a field has been set.

### GetSourceVrf

`func (o *FwlogFwLog) GetSourceVrf() string`

GetSourceVrf returns the SourceVrf field if non-nil, zero value otherwise.

### GetSourceVrfOk

`func (o *FwlogFwLog) GetSourceVrfOk() (*string, bool)`

GetSourceVrfOk returns a tuple with the SourceVrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVrf

`func (o *FwlogFwLog) SetSourceVrf(v string)`

SetSourceVrf sets SourceVrf field to given value.

### HasSourceVrf

`func (o *FwlogFwLog) HasSourceVrf() bool`

HasSourceVrf returns a boolean if a field has been set.

### GetVnid

`func (o *FwlogFwLog) GetVnid() string`

GetVnid returns the Vnid field if non-nil, zero value otherwise.

### GetVnidOk

`func (o *FwlogFwLog) GetVnidOk() (*string, bool)`

GetVnidOk returns a tuple with the Vnid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnid

`func (o *FwlogFwLog) SetVnid(v string)`

SetVnid sets Vnid field to given value.

### HasVnid

`func (o *FwlogFwLog) HasVnid() bool`

HasVnid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


