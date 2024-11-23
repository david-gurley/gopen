# SecurityProtoPort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ports** | Pointer to **string** | TCP or UDP port number(s): comma separate port numbers, or dash separate port range. | [optional] 
**Protocol** | Pointer to **string** | protocol is ip (v4/v6) protocol name/number; names can be: tcp, udp, icmp. | [optional] 

## Methods

### NewSecurityProtoPort

`func NewSecurityProtoPort() *SecurityProtoPort`

NewSecurityProtoPort instantiates a new SecurityProtoPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityProtoPortWithDefaults

`func NewSecurityProtoPortWithDefaults() *SecurityProtoPort`

NewSecurityProtoPortWithDefaults instantiates a new SecurityProtoPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPorts

`func (o *SecurityProtoPort) GetPorts() string`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *SecurityProtoPort) GetPortsOk() (*string, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *SecurityProtoPort) SetPorts(v string)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *SecurityProtoPort) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetProtocol

`func (o *SecurityProtoPort) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *SecurityProtoPort) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *SecurityProtoPort) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *SecurityProtoPort) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


