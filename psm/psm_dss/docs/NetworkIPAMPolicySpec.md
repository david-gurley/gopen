# NetworkIPAMPolicySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DhcpRelay** | Pointer to [**NetworkDHCPRelayPolicy**](networkDHCPRelayPolicy.md) |  | [optional] 
**PsmIpamConfig** | Pointer to [**NetworkIPAMConfig**](networkIPAMConfig.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] [default to "dhcp-relay"]

## Methods

### NewNetworkIPAMPolicySpec

`func NewNetworkIPAMPolicySpec() *NetworkIPAMPolicySpec`

NewNetworkIPAMPolicySpec instantiates a new NetworkIPAMPolicySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkIPAMPolicySpecWithDefaults

`func NewNetworkIPAMPolicySpecWithDefaults() *NetworkIPAMPolicySpec`

NewNetworkIPAMPolicySpecWithDefaults instantiates a new NetworkIPAMPolicySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDhcpRelay

`func (o *NetworkIPAMPolicySpec) GetDhcpRelay() NetworkDHCPRelayPolicy`

GetDhcpRelay returns the DhcpRelay field if non-nil, zero value otherwise.

### GetDhcpRelayOk

`func (o *NetworkIPAMPolicySpec) GetDhcpRelayOk() (*NetworkDHCPRelayPolicy, bool)`

GetDhcpRelayOk returns a tuple with the DhcpRelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpRelay

`func (o *NetworkIPAMPolicySpec) SetDhcpRelay(v NetworkDHCPRelayPolicy)`

SetDhcpRelay sets DhcpRelay field to given value.

### HasDhcpRelay

`func (o *NetworkIPAMPolicySpec) HasDhcpRelay() bool`

HasDhcpRelay returns a boolean if a field has been set.

### GetPsmIpamConfig

`func (o *NetworkIPAMPolicySpec) GetPsmIpamConfig() NetworkIPAMConfig`

GetPsmIpamConfig returns the PsmIpamConfig field if non-nil, zero value otherwise.

### GetPsmIpamConfigOk

`func (o *NetworkIPAMPolicySpec) GetPsmIpamConfigOk() (*NetworkIPAMConfig, bool)`

GetPsmIpamConfigOk returns a tuple with the PsmIpamConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsmIpamConfig

`func (o *NetworkIPAMPolicySpec) SetPsmIpamConfig(v NetworkIPAMConfig)`

SetPsmIpamConfig sets PsmIpamConfig field to given value.

### HasPsmIpamConfig

`func (o *NetworkIPAMPolicySpec) HasPsmIpamConfig() bool`

HasPsmIpamConfig returns a boolean if a field has been set.

### GetType

`func (o *NetworkIPAMPolicySpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkIPAMPolicySpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkIPAMPolicySpec) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NetworkIPAMPolicySpec) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


