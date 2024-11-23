# SysruntimeFlowKeyV4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | Pointer to **string** |  | [optional] 
**Esp** | Pointer to [**SysruntimeFlowKeyESPInfo**](sysruntimeFlowKeyESPInfo.md) |  | [optional] 
**Icmp** | Pointer to [**SysruntimeFlowKeyICMPInfo**](sysruntimeFlowKeyICMPInfo.md) |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**TcpUdp** | Pointer to [**SysruntimeFlowKeyTcpUdpInfo**](sysruntimeFlowKeyTcpUdpInfo.md) |  | [optional] 

## Methods

### NewSysruntimeFlowKeyV4

`func NewSysruntimeFlowKeyV4() *SysruntimeFlowKeyV4`

NewSysruntimeFlowKeyV4 instantiates a new SysruntimeFlowKeyV4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSysruntimeFlowKeyV4WithDefaults

`func NewSysruntimeFlowKeyV4WithDefaults() *SysruntimeFlowKeyV4`

NewSysruntimeFlowKeyV4WithDefaults instantiates a new SysruntimeFlowKeyV4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *SysruntimeFlowKeyV4) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *SysruntimeFlowKeyV4) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *SysruntimeFlowKeyV4) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *SysruntimeFlowKeyV4) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetEsp

`func (o *SysruntimeFlowKeyV4) GetEsp() SysruntimeFlowKeyESPInfo`

GetEsp returns the Esp field if non-nil, zero value otherwise.

### GetEspOk

`func (o *SysruntimeFlowKeyV4) GetEspOk() (*SysruntimeFlowKeyESPInfo, bool)`

GetEspOk returns a tuple with the Esp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsp

`func (o *SysruntimeFlowKeyV4) SetEsp(v SysruntimeFlowKeyESPInfo)`

SetEsp sets Esp field to given value.

### HasEsp

`func (o *SysruntimeFlowKeyV4) HasEsp() bool`

HasEsp returns a boolean if a field has been set.

### GetIcmp

`func (o *SysruntimeFlowKeyV4) GetIcmp() SysruntimeFlowKeyICMPInfo`

GetIcmp returns the Icmp field if non-nil, zero value otherwise.

### GetIcmpOk

`func (o *SysruntimeFlowKeyV4) GetIcmpOk() (*SysruntimeFlowKeyICMPInfo, bool)`

GetIcmpOk returns a tuple with the Icmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcmp

`func (o *SysruntimeFlowKeyV4) SetIcmp(v SysruntimeFlowKeyICMPInfo)`

SetIcmp sets Icmp field to given value.

### HasIcmp

`func (o *SysruntimeFlowKeyV4) HasIcmp() bool`

HasIcmp returns a boolean if a field has been set.

### GetProtocol

`func (o *SysruntimeFlowKeyV4) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *SysruntimeFlowKeyV4) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *SysruntimeFlowKeyV4) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *SysruntimeFlowKeyV4) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSource

`func (o *SysruntimeFlowKeyV4) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SysruntimeFlowKeyV4) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SysruntimeFlowKeyV4) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *SysruntimeFlowKeyV4) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTcpUdp

`func (o *SysruntimeFlowKeyV4) GetTcpUdp() SysruntimeFlowKeyTcpUdpInfo`

GetTcpUdp returns the TcpUdp field if non-nil, zero value otherwise.

### GetTcpUdpOk

`func (o *SysruntimeFlowKeyV4) GetTcpUdpOk() (*SysruntimeFlowKeyTcpUdpInfo, bool)`

GetTcpUdpOk returns a tuple with the TcpUdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpUdp

`func (o *SysruntimeFlowKeyV4) SetTcpUdp(v SysruntimeFlowKeyTcpUdpInfo)`

SetTcpUdp sets TcpUdp field to given value.

### HasTcpUdp

`func (o *SysruntimeFlowKeyV4) HasTcpUdp() bool`

HasTcpUdp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


