# SecurityDns

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DropLargeDomainNamePackets** | Pointer to **bool** | Drop if domain name size is &gt; 255 bytes. | [optional] 
**DropLongLabelPackets** | Pointer to **bool** | Drop if label length is 64 bytes or higher. | [optional] 
**DropMultiQuestionPackets** | Pointer to **bool** | Drop packet if number of questions is more than one. | [optional] 
**MaxMessageLength** | Pointer to **int64** | Maximum message length, default value is 512, maximum specified user value is 8129. | [optional] 
**QueryResponseTimeout** | Pointer to **string** | Timeout for DNS Query, default 60s. Should be a valid time duration. | [optional] [default to "60s"]

## Methods

### NewSecurityDns

`func NewSecurityDns() *SecurityDns`

NewSecurityDns instantiates a new SecurityDns object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityDnsWithDefaults

`func NewSecurityDnsWithDefaults() *SecurityDns`

NewSecurityDnsWithDefaults instantiates a new SecurityDns object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDropLargeDomainNamePackets

`func (o *SecurityDns) GetDropLargeDomainNamePackets() bool`

GetDropLargeDomainNamePackets returns the DropLargeDomainNamePackets field if non-nil, zero value otherwise.

### GetDropLargeDomainNamePacketsOk

`func (o *SecurityDns) GetDropLargeDomainNamePacketsOk() (*bool, bool)`

GetDropLargeDomainNamePacketsOk returns a tuple with the DropLargeDomainNamePackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropLargeDomainNamePackets

`func (o *SecurityDns) SetDropLargeDomainNamePackets(v bool)`

SetDropLargeDomainNamePackets sets DropLargeDomainNamePackets field to given value.

### HasDropLargeDomainNamePackets

`func (o *SecurityDns) HasDropLargeDomainNamePackets() bool`

HasDropLargeDomainNamePackets returns a boolean if a field has been set.

### GetDropLongLabelPackets

`func (o *SecurityDns) GetDropLongLabelPackets() bool`

GetDropLongLabelPackets returns the DropLongLabelPackets field if non-nil, zero value otherwise.

### GetDropLongLabelPacketsOk

`func (o *SecurityDns) GetDropLongLabelPacketsOk() (*bool, bool)`

GetDropLongLabelPacketsOk returns a tuple with the DropLongLabelPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropLongLabelPackets

`func (o *SecurityDns) SetDropLongLabelPackets(v bool)`

SetDropLongLabelPackets sets DropLongLabelPackets field to given value.

### HasDropLongLabelPackets

`func (o *SecurityDns) HasDropLongLabelPackets() bool`

HasDropLongLabelPackets returns a boolean if a field has been set.

### GetDropMultiQuestionPackets

`func (o *SecurityDns) GetDropMultiQuestionPackets() bool`

GetDropMultiQuestionPackets returns the DropMultiQuestionPackets field if non-nil, zero value otherwise.

### GetDropMultiQuestionPacketsOk

`func (o *SecurityDns) GetDropMultiQuestionPacketsOk() (*bool, bool)`

GetDropMultiQuestionPacketsOk returns a tuple with the DropMultiQuestionPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropMultiQuestionPackets

`func (o *SecurityDns) SetDropMultiQuestionPackets(v bool)`

SetDropMultiQuestionPackets sets DropMultiQuestionPackets field to given value.

### HasDropMultiQuestionPackets

`func (o *SecurityDns) HasDropMultiQuestionPackets() bool`

HasDropMultiQuestionPackets returns a boolean if a field has been set.

### GetMaxMessageLength

`func (o *SecurityDns) GetMaxMessageLength() int64`

GetMaxMessageLength returns the MaxMessageLength field if non-nil, zero value otherwise.

### GetMaxMessageLengthOk

`func (o *SecurityDns) GetMaxMessageLengthOk() (*int64, bool)`

GetMaxMessageLengthOk returns a tuple with the MaxMessageLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMessageLength

`func (o *SecurityDns) SetMaxMessageLength(v int64)`

SetMaxMessageLength sets MaxMessageLength field to given value.

### HasMaxMessageLength

`func (o *SecurityDns) HasMaxMessageLength() bool`

HasMaxMessageLength returns a boolean if a field has been set.

### GetQueryResponseTimeout

`func (o *SecurityDns) GetQueryResponseTimeout() string`

GetQueryResponseTimeout returns the QueryResponseTimeout field if non-nil, zero value otherwise.

### GetQueryResponseTimeoutOk

`func (o *SecurityDns) GetQueryResponseTimeoutOk() (*string, bool)`

GetQueryResponseTimeoutOk returns a tuple with the QueryResponseTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryResponseTimeout

`func (o *SecurityDns) SetQueryResponseTimeout(v string)`

SetQueryResponseTimeout sets QueryResponseTimeout field to given value.

### HasQueryResponseTimeout

`func (o *SecurityDns) HasQueryResponseTimeout() bool`

HasQueryResponseTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


