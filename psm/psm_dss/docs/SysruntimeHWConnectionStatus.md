# SysruntimeHWConnectionStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationMac** | Pointer to **string** |  | [optional] 
**FirewallStatus** | Pointer to [**SysruntimeFwStatus**](sysruntimeFwStatus.md) |  | [optional] 
**InitiatorFlowStatus** | Pointer to [**SysruntimeFlowStatus**](sysruntimeFlowStatus.md) |  | [optional] 
**IpsecStatus** | Pointer to [**SysruntimeIPSecStatus**](sysruntimeIPSecStatus.md) |  | [optional] 
**PeerIflowStatus** | Pointer to [**SysruntimeFlowStatus**](sysruntimeFlowStatus.md) |  | [optional] 
**PeerRflowStatus** | Pointer to [**SysruntimeFlowStatus**](sysruntimeFlowStatus.md) |  | [optional] 
**ResponderFlowStatus** | Pointer to [**SysruntimeFlowStatus**](sysruntimeFlowStatus.md) |  | [optional] 
**SessionHandle** | Pointer to **string** |  | [optional] 
**SourceMac** | Pointer to **string** |  | [optional] 

## Methods

### NewSysruntimeHWConnectionStatus

`func NewSysruntimeHWConnectionStatus() *SysruntimeHWConnectionStatus`

NewSysruntimeHWConnectionStatus instantiates a new SysruntimeHWConnectionStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSysruntimeHWConnectionStatusWithDefaults

`func NewSysruntimeHWConnectionStatusWithDefaults() *SysruntimeHWConnectionStatus`

NewSysruntimeHWConnectionStatusWithDefaults instantiates a new SysruntimeHWConnectionStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationMac

`func (o *SysruntimeHWConnectionStatus) GetDestinationMac() string`

GetDestinationMac returns the DestinationMac field if non-nil, zero value otherwise.

### GetDestinationMacOk

`func (o *SysruntimeHWConnectionStatus) GetDestinationMacOk() (*string, bool)`

GetDestinationMacOk returns a tuple with the DestinationMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationMac

`func (o *SysruntimeHWConnectionStatus) SetDestinationMac(v string)`

SetDestinationMac sets DestinationMac field to given value.

### HasDestinationMac

`func (o *SysruntimeHWConnectionStatus) HasDestinationMac() bool`

HasDestinationMac returns a boolean if a field has been set.

### GetFirewallStatus

`func (o *SysruntimeHWConnectionStatus) GetFirewallStatus() SysruntimeFwStatus`

GetFirewallStatus returns the FirewallStatus field if non-nil, zero value otherwise.

### GetFirewallStatusOk

`func (o *SysruntimeHWConnectionStatus) GetFirewallStatusOk() (*SysruntimeFwStatus, bool)`

GetFirewallStatusOk returns a tuple with the FirewallStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallStatus

`func (o *SysruntimeHWConnectionStatus) SetFirewallStatus(v SysruntimeFwStatus)`

SetFirewallStatus sets FirewallStatus field to given value.

### HasFirewallStatus

`func (o *SysruntimeHWConnectionStatus) HasFirewallStatus() bool`

HasFirewallStatus returns a boolean if a field has been set.

### GetInitiatorFlowStatus

`func (o *SysruntimeHWConnectionStatus) GetInitiatorFlowStatus() SysruntimeFlowStatus`

GetInitiatorFlowStatus returns the InitiatorFlowStatus field if non-nil, zero value otherwise.

### GetInitiatorFlowStatusOk

`func (o *SysruntimeHWConnectionStatus) GetInitiatorFlowStatusOk() (*SysruntimeFlowStatus, bool)`

GetInitiatorFlowStatusOk returns a tuple with the InitiatorFlowStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorFlowStatus

`func (o *SysruntimeHWConnectionStatus) SetInitiatorFlowStatus(v SysruntimeFlowStatus)`

SetInitiatorFlowStatus sets InitiatorFlowStatus field to given value.

### HasInitiatorFlowStatus

`func (o *SysruntimeHWConnectionStatus) HasInitiatorFlowStatus() bool`

HasInitiatorFlowStatus returns a boolean if a field has been set.

### GetIpsecStatus

`func (o *SysruntimeHWConnectionStatus) GetIpsecStatus() SysruntimeIPSecStatus`

GetIpsecStatus returns the IpsecStatus field if non-nil, zero value otherwise.

### GetIpsecStatusOk

`func (o *SysruntimeHWConnectionStatus) GetIpsecStatusOk() (*SysruntimeIPSecStatus, bool)`

GetIpsecStatusOk returns a tuple with the IpsecStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecStatus

`func (o *SysruntimeHWConnectionStatus) SetIpsecStatus(v SysruntimeIPSecStatus)`

SetIpsecStatus sets IpsecStatus field to given value.

### HasIpsecStatus

`func (o *SysruntimeHWConnectionStatus) HasIpsecStatus() bool`

HasIpsecStatus returns a boolean if a field has been set.

### GetPeerIflowStatus

`func (o *SysruntimeHWConnectionStatus) GetPeerIflowStatus() SysruntimeFlowStatus`

GetPeerIflowStatus returns the PeerIflowStatus field if non-nil, zero value otherwise.

### GetPeerIflowStatusOk

`func (o *SysruntimeHWConnectionStatus) GetPeerIflowStatusOk() (*SysruntimeFlowStatus, bool)`

GetPeerIflowStatusOk returns a tuple with the PeerIflowStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerIflowStatus

`func (o *SysruntimeHWConnectionStatus) SetPeerIflowStatus(v SysruntimeFlowStatus)`

SetPeerIflowStatus sets PeerIflowStatus field to given value.

### HasPeerIflowStatus

`func (o *SysruntimeHWConnectionStatus) HasPeerIflowStatus() bool`

HasPeerIflowStatus returns a boolean if a field has been set.

### GetPeerRflowStatus

`func (o *SysruntimeHWConnectionStatus) GetPeerRflowStatus() SysruntimeFlowStatus`

GetPeerRflowStatus returns the PeerRflowStatus field if non-nil, zero value otherwise.

### GetPeerRflowStatusOk

`func (o *SysruntimeHWConnectionStatus) GetPeerRflowStatusOk() (*SysruntimeFlowStatus, bool)`

GetPeerRflowStatusOk returns a tuple with the PeerRflowStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerRflowStatus

`func (o *SysruntimeHWConnectionStatus) SetPeerRflowStatus(v SysruntimeFlowStatus)`

SetPeerRflowStatus sets PeerRflowStatus field to given value.

### HasPeerRflowStatus

`func (o *SysruntimeHWConnectionStatus) HasPeerRflowStatus() bool`

HasPeerRflowStatus returns a boolean if a field has been set.

### GetResponderFlowStatus

`func (o *SysruntimeHWConnectionStatus) GetResponderFlowStatus() SysruntimeFlowStatus`

GetResponderFlowStatus returns the ResponderFlowStatus field if non-nil, zero value otherwise.

### GetResponderFlowStatusOk

`func (o *SysruntimeHWConnectionStatus) GetResponderFlowStatusOk() (*SysruntimeFlowStatus, bool)`

GetResponderFlowStatusOk returns a tuple with the ResponderFlowStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponderFlowStatus

`func (o *SysruntimeHWConnectionStatus) SetResponderFlowStatus(v SysruntimeFlowStatus)`

SetResponderFlowStatus sets ResponderFlowStatus field to given value.

### HasResponderFlowStatus

`func (o *SysruntimeHWConnectionStatus) HasResponderFlowStatus() bool`

HasResponderFlowStatus returns a boolean if a field has been set.

### GetSessionHandle

`func (o *SysruntimeHWConnectionStatus) GetSessionHandle() string`

GetSessionHandle returns the SessionHandle field if non-nil, zero value otherwise.

### GetSessionHandleOk

`func (o *SysruntimeHWConnectionStatus) GetSessionHandleOk() (*string, bool)`

GetSessionHandleOk returns a tuple with the SessionHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionHandle

`func (o *SysruntimeHWConnectionStatus) SetSessionHandle(v string)`

SetSessionHandle sets SessionHandle field to given value.

### HasSessionHandle

`func (o *SysruntimeHWConnectionStatus) HasSessionHandle() bool`

HasSessionHandle returns a boolean if a field has been set.

### GetSourceMac

`func (o *SysruntimeHWConnectionStatus) GetSourceMac() string`

GetSourceMac returns the SourceMac field if non-nil, zero value otherwise.

### GetSourceMacOk

`func (o *SysruntimeHWConnectionStatus) GetSourceMacOk() (*string, bool)`

GetSourceMacOk returns a tuple with the SourceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceMac

`func (o *SysruntimeHWConnectionStatus) SetSourceMac(v string)`

SetSourceMac sets SourceMac field to given value.

### HasSourceMac

`func (o *SysruntimeHWConnectionStatus) HasSourceMac() bool`

HasSourceMac returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


