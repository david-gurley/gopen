# SecurityFirewallProfileSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionTracking** | Pointer to **bool** | Enable/disable Connection Tracking. | [optional] 
**DetectApp** | Pointer to **bool** | Set the Application Identification Detection config for DSCs. | [optional] [default to false]
**DropTimeout** | Pointer to **string** | Drop Timeout is the period for which a drop entry is installed for a denied non tcp/udp/icmp flow. Should be a valid time duration between 1s and 5m0s. | [optional] [default to "60s"]
**IcmpActiveSessionLimit** | Pointer to **int64** | Icmp active session limit config after which new requests will be dropped. Value should be between 0 and 32768. | [optional] [default to 0]
**IcmpDropTimeout** | Pointer to **string** | ICMP Drop Timeout is the period for which a drop entry is installed for a denied ICMP flow. Should be a valid time duration between 1s and 5m0s. | [optional] [default to "60s"]
**IcmpTimeout** | Pointer to **string** | Icmp Timeout is the period for which a ICMP session is kept alive during inactivity. Should be a valid time duration between 1s and 48h0m0s. | [optional] [default to "6s"]
**SessionIdleTimeout** | Pointer to **string** | Session idle timeout removes/deletes the session/flow if there is inactivity; this value is superceded by any value specified in App object. Should be a valid time duration between 30s and 48h0m0s. | [optional] [default to "90s"]
**TcpCloseTimeout** | Pointer to **string** | TCP Close Timeout is the time for which TCP session is kept after a FIN is seen. Should be a valid time duration between 1s and 5m0s. | [optional] [default to "1s"]
**TcpConnectionSetupTimeout** | Pointer to **string** | TCP Connection Setup Timeout is the period TCP session is kept to see the response of a SYN. Should be a valid time duration between 1s and 1m0s. | [optional] [default to "30s"]
**TcpDropTimeout** | Pointer to **string** | TCP Drop Timeout is the period for which a drop entry is installed for a denied TCP flow. Should be a valid time duration between 1s and 5m0s. | [optional] [default to "90s"]
**TcpHalfClosedTimeout** | Pointer to **string** | TCP Half Closed Timeout is the time for which tCP session is kept when connection is half closed i.e. FIN sent by FIN_Ack not received. Should be a valid time duration between 1s and 48h0m0s. | [optional] [default to "120s"]
**TcpHalfOpenSessionLimit** | Pointer to **int64** | Tcp half open session limit config after which new open requests will be dropped. Value should be between 0 and 32768. | [optional] [default to 0]
**TcpTimeout** | Pointer to **string** | Tcp Timeout is the period for which a TCP session is kept alive during inactivity. Should be a valid time duration between 1s and 48h0m0s. | [optional] [default to "3600s"]
**UdpActiveSessionLimit** | Pointer to **int64** | Udp active session limit config after which new requests will be dropped. Value should be between 0 and 32768. | [optional] [default to 0]
**UdpDropTimeout** | Pointer to **string** | UDP Drop Timeout is the period for which a drop entry is installed for a denied UDP flow. Should be a valid time duration between 1s and 48h0m0s. | [optional] [default to "60s"]
**UdpTimeout** | Pointer to **string** | Udp Timeout is the period for which a UDP session is kept alive during inactivity. Should be a valid time duration between 1s and 48h0m0s. | [optional] [default to "30s"]

## Methods

### NewSecurityFirewallProfileSpec

`func NewSecurityFirewallProfileSpec() *SecurityFirewallProfileSpec`

NewSecurityFirewallProfileSpec instantiates a new SecurityFirewallProfileSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityFirewallProfileSpecWithDefaults

`func NewSecurityFirewallProfileSpecWithDefaults() *SecurityFirewallProfileSpec`

NewSecurityFirewallProfileSpecWithDefaults instantiates a new SecurityFirewallProfileSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionTracking

`func (o *SecurityFirewallProfileSpec) GetConnectionTracking() bool`

GetConnectionTracking returns the ConnectionTracking field if non-nil, zero value otherwise.

### GetConnectionTrackingOk

`func (o *SecurityFirewallProfileSpec) GetConnectionTrackingOk() (*bool, bool)`

GetConnectionTrackingOk returns a tuple with the ConnectionTracking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTracking

`func (o *SecurityFirewallProfileSpec) SetConnectionTracking(v bool)`

SetConnectionTracking sets ConnectionTracking field to given value.

### HasConnectionTracking

`func (o *SecurityFirewallProfileSpec) HasConnectionTracking() bool`

HasConnectionTracking returns a boolean if a field has been set.

### GetDetectApp

`func (o *SecurityFirewallProfileSpec) GetDetectApp() bool`

GetDetectApp returns the DetectApp field if non-nil, zero value otherwise.

### GetDetectAppOk

`func (o *SecurityFirewallProfileSpec) GetDetectAppOk() (*bool, bool)`

GetDetectAppOk returns a tuple with the DetectApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectApp

`func (o *SecurityFirewallProfileSpec) SetDetectApp(v bool)`

SetDetectApp sets DetectApp field to given value.

### HasDetectApp

`func (o *SecurityFirewallProfileSpec) HasDetectApp() bool`

HasDetectApp returns a boolean if a field has been set.

### GetDropTimeout

`func (o *SecurityFirewallProfileSpec) GetDropTimeout() string`

GetDropTimeout returns the DropTimeout field if non-nil, zero value otherwise.

### GetDropTimeoutOk

`func (o *SecurityFirewallProfileSpec) GetDropTimeoutOk() (*string, bool)`

GetDropTimeoutOk returns a tuple with the DropTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropTimeout

`func (o *SecurityFirewallProfileSpec) SetDropTimeout(v string)`

SetDropTimeout sets DropTimeout field to given value.

### HasDropTimeout

`func (o *SecurityFirewallProfileSpec) HasDropTimeout() bool`

HasDropTimeout returns a boolean if a field has been set.

### GetIcmpActiveSessionLimit

`func (o *SecurityFirewallProfileSpec) GetIcmpActiveSessionLimit() int64`

GetIcmpActiveSessionLimit returns the IcmpActiveSessionLimit field if non-nil, zero value otherwise.

### GetIcmpActiveSessionLimitOk

`func (o *SecurityFirewallProfileSpec) GetIcmpActiveSessionLimitOk() (*int64, bool)`

GetIcmpActiveSessionLimitOk returns a tuple with the IcmpActiveSessionLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcmpActiveSessionLimit

`func (o *SecurityFirewallProfileSpec) SetIcmpActiveSessionLimit(v int64)`

SetIcmpActiveSessionLimit sets IcmpActiveSessionLimit field to given value.

### HasIcmpActiveSessionLimit

`func (o *SecurityFirewallProfileSpec) HasIcmpActiveSessionLimit() bool`

HasIcmpActiveSessionLimit returns a boolean if a field has been set.

### GetIcmpDropTimeout

`func (o *SecurityFirewallProfileSpec) GetIcmpDropTimeout() string`

GetIcmpDropTimeout returns the IcmpDropTimeout field if non-nil, zero value otherwise.

### GetIcmpDropTimeoutOk

`func (o *SecurityFirewallProfileSpec) GetIcmpDropTimeoutOk() (*string, bool)`

GetIcmpDropTimeoutOk returns a tuple with the IcmpDropTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcmpDropTimeout

`func (o *SecurityFirewallProfileSpec) SetIcmpDropTimeout(v string)`

SetIcmpDropTimeout sets IcmpDropTimeout field to given value.

### HasIcmpDropTimeout

`func (o *SecurityFirewallProfileSpec) HasIcmpDropTimeout() bool`

HasIcmpDropTimeout returns a boolean if a field has been set.

### GetIcmpTimeout

`func (o *SecurityFirewallProfileSpec) GetIcmpTimeout() string`

GetIcmpTimeout returns the IcmpTimeout field if non-nil, zero value otherwise.

### GetIcmpTimeoutOk

`func (o *SecurityFirewallProfileSpec) GetIcmpTimeoutOk() (*string, bool)`

GetIcmpTimeoutOk returns a tuple with the IcmpTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcmpTimeout

`func (o *SecurityFirewallProfileSpec) SetIcmpTimeout(v string)`

SetIcmpTimeout sets IcmpTimeout field to given value.

### HasIcmpTimeout

`func (o *SecurityFirewallProfileSpec) HasIcmpTimeout() bool`

HasIcmpTimeout returns a boolean if a field has been set.

### GetSessionIdleTimeout

`func (o *SecurityFirewallProfileSpec) GetSessionIdleTimeout() string`

GetSessionIdleTimeout returns the SessionIdleTimeout field if non-nil, zero value otherwise.

### GetSessionIdleTimeoutOk

`func (o *SecurityFirewallProfileSpec) GetSessionIdleTimeoutOk() (*string, bool)`

GetSessionIdleTimeoutOk returns a tuple with the SessionIdleTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionIdleTimeout

`func (o *SecurityFirewallProfileSpec) SetSessionIdleTimeout(v string)`

SetSessionIdleTimeout sets SessionIdleTimeout field to given value.

### HasSessionIdleTimeout

`func (o *SecurityFirewallProfileSpec) HasSessionIdleTimeout() bool`

HasSessionIdleTimeout returns a boolean if a field has been set.

### GetTcpCloseTimeout

`func (o *SecurityFirewallProfileSpec) GetTcpCloseTimeout() string`

GetTcpCloseTimeout returns the TcpCloseTimeout field if non-nil, zero value otherwise.

### GetTcpCloseTimeoutOk

`func (o *SecurityFirewallProfileSpec) GetTcpCloseTimeoutOk() (*string, bool)`

GetTcpCloseTimeoutOk returns a tuple with the TcpCloseTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpCloseTimeout

`func (o *SecurityFirewallProfileSpec) SetTcpCloseTimeout(v string)`

SetTcpCloseTimeout sets TcpCloseTimeout field to given value.

### HasTcpCloseTimeout

`func (o *SecurityFirewallProfileSpec) HasTcpCloseTimeout() bool`

HasTcpCloseTimeout returns a boolean if a field has been set.

### GetTcpConnectionSetupTimeout

`func (o *SecurityFirewallProfileSpec) GetTcpConnectionSetupTimeout() string`

GetTcpConnectionSetupTimeout returns the TcpConnectionSetupTimeout field if non-nil, zero value otherwise.

### GetTcpConnectionSetupTimeoutOk

`func (o *SecurityFirewallProfileSpec) GetTcpConnectionSetupTimeoutOk() (*string, bool)`

GetTcpConnectionSetupTimeoutOk returns a tuple with the TcpConnectionSetupTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpConnectionSetupTimeout

`func (o *SecurityFirewallProfileSpec) SetTcpConnectionSetupTimeout(v string)`

SetTcpConnectionSetupTimeout sets TcpConnectionSetupTimeout field to given value.

### HasTcpConnectionSetupTimeout

`func (o *SecurityFirewallProfileSpec) HasTcpConnectionSetupTimeout() bool`

HasTcpConnectionSetupTimeout returns a boolean if a field has been set.

### GetTcpDropTimeout

`func (o *SecurityFirewallProfileSpec) GetTcpDropTimeout() string`

GetTcpDropTimeout returns the TcpDropTimeout field if non-nil, zero value otherwise.

### GetTcpDropTimeoutOk

`func (o *SecurityFirewallProfileSpec) GetTcpDropTimeoutOk() (*string, bool)`

GetTcpDropTimeoutOk returns a tuple with the TcpDropTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpDropTimeout

`func (o *SecurityFirewallProfileSpec) SetTcpDropTimeout(v string)`

SetTcpDropTimeout sets TcpDropTimeout field to given value.

### HasTcpDropTimeout

`func (o *SecurityFirewallProfileSpec) HasTcpDropTimeout() bool`

HasTcpDropTimeout returns a boolean if a field has been set.

### GetTcpHalfClosedTimeout

`func (o *SecurityFirewallProfileSpec) GetTcpHalfClosedTimeout() string`

GetTcpHalfClosedTimeout returns the TcpHalfClosedTimeout field if non-nil, zero value otherwise.

### GetTcpHalfClosedTimeoutOk

`func (o *SecurityFirewallProfileSpec) GetTcpHalfClosedTimeoutOk() (*string, bool)`

GetTcpHalfClosedTimeoutOk returns a tuple with the TcpHalfClosedTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpHalfClosedTimeout

`func (o *SecurityFirewallProfileSpec) SetTcpHalfClosedTimeout(v string)`

SetTcpHalfClosedTimeout sets TcpHalfClosedTimeout field to given value.

### HasTcpHalfClosedTimeout

`func (o *SecurityFirewallProfileSpec) HasTcpHalfClosedTimeout() bool`

HasTcpHalfClosedTimeout returns a boolean if a field has been set.

### GetTcpHalfOpenSessionLimit

`func (o *SecurityFirewallProfileSpec) GetTcpHalfOpenSessionLimit() int64`

GetTcpHalfOpenSessionLimit returns the TcpHalfOpenSessionLimit field if non-nil, zero value otherwise.

### GetTcpHalfOpenSessionLimitOk

`func (o *SecurityFirewallProfileSpec) GetTcpHalfOpenSessionLimitOk() (*int64, bool)`

GetTcpHalfOpenSessionLimitOk returns a tuple with the TcpHalfOpenSessionLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpHalfOpenSessionLimit

`func (o *SecurityFirewallProfileSpec) SetTcpHalfOpenSessionLimit(v int64)`

SetTcpHalfOpenSessionLimit sets TcpHalfOpenSessionLimit field to given value.

### HasTcpHalfOpenSessionLimit

`func (o *SecurityFirewallProfileSpec) HasTcpHalfOpenSessionLimit() bool`

HasTcpHalfOpenSessionLimit returns a boolean if a field has been set.

### GetTcpTimeout

`func (o *SecurityFirewallProfileSpec) GetTcpTimeout() string`

GetTcpTimeout returns the TcpTimeout field if non-nil, zero value otherwise.

### GetTcpTimeoutOk

`func (o *SecurityFirewallProfileSpec) GetTcpTimeoutOk() (*string, bool)`

GetTcpTimeoutOk returns a tuple with the TcpTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpTimeout

`func (o *SecurityFirewallProfileSpec) SetTcpTimeout(v string)`

SetTcpTimeout sets TcpTimeout field to given value.

### HasTcpTimeout

`func (o *SecurityFirewallProfileSpec) HasTcpTimeout() bool`

HasTcpTimeout returns a boolean if a field has been set.

### GetUdpActiveSessionLimit

`func (o *SecurityFirewallProfileSpec) GetUdpActiveSessionLimit() int64`

GetUdpActiveSessionLimit returns the UdpActiveSessionLimit field if non-nil, zero value otherwise.

### GetUdpActiveSessionLimitOk

`func (o *SecurityFirewallProfileSpec) GetUdpActiveSessionLimitOk() (*int64, bool)`

GetUdpActiveSessionLimitOk returns a tuple with the UdpActiveSessionLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdpActiveSessionLimit

`func (o *SecurityFirewallProfileSpec) SetUdpActiveSessionLimit(v int64)`

SetUdpActiveSessionLimit sets UdpActiveSessionLimit field to given value.

### HasUdpActiveSessionLimit

`func (o *SecurityFirewallProfileSpec) HasUdpActiveSessionLimit() bool`

HasUdpActiveSessionLimit returns a boolean if a field has been set.

### GetUdpDropTimeout

`func (o *SecurityFirewallProfileSpec) GetUdpDropTimeout() string`

GetUdpDropTimeout returns the UdpDropTimeout field if non-nil, zero value otherwise.

### GetUdpDropTimeoutOk

`func (o *SecurityFirewallProfileSpec) GetUdpDropTimeoutOk() (*string, bool)`

GetUdpDropTimeoutOk returns a tuple with the UdpDropTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdpDropTimeout

`func (o *SecurityFirewallProfileSpec) SetUdpDropTimeout(v string)`

SetUdpDropTimeout sets UdpDropTimeout field to given value.

### HasUdpDropTimeout

`func (o *SecurityFirewallProfileSpec) HasUdpDropTimeout() bool`

HasUdpDropTimeout returns a boolean if a field has been set.

### GetUdpTimeout

`func (o *SecurityFirewallProfileSpec) GetUdpTimeout() string`

GetUdpTimeout returns the UdpTimeout field if non-nil, zero value otherwise.

### GetUdpTimeoutOk

`func (o *SecurityFirewallProfileSpec) GetUdpTimeoutOk() (*string, bool)`

GetUdpTimeoutOk returns a tuple with the UdpTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdpTimeout

`func (o *SecurityFirewallProfileSpec) SetUdpTimeout(v string)`

SetUdpTimeout sets UdpTimeout field to given value.

### HasUdpTimeout

`func (o *SecurityFirewallProfileSpec) HasUdpTimeout() bool`

HasUdpTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


