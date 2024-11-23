# NetworkDHCPRelayPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RelayServers** | Pointer to [**[]NetworkDHCPServer**](NetworkDHCPServer.md) |  | [optional] 

## Methods

### NewNetworkDHCPRelayPolicy

`func NewNetworkDHCPRelayPolicy() *NetworkDHCPRelayPolicy`

NewNetworkDHCPRelayPolicy instantiates a new NetworkDHCPRelayPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDHCPRelayPolicyWithDefaults

`func NewNetworkDHCPRelayPolicyWithDefaults() *NetworkDHCPRelayPolicy`

NewNetworkDHCPRelayPolicyWithDefaults instantiates a new NetworkDHCPRelayPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRelayServers

`func (o *NetworkDHCPRelayPolicy) GetRelayServers() []NetworkDHCPServer`

GetRelayServers returns the RelayServers field if non-nil, zero value otherwise.

### GetRelayServersOk

`func (o *NetworkDHCPRelayPolicy) GetRelayServersOk() (*[]NetworkDHCPServer, bool)`

GetRelayServersOk returns a tuple with the RelayServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayServers

`func (o *NetworkDHCPRelayPolicy) SetRelayServers(v []NetworkDHCPServer)`

SetRelayServers sets RelayServers field to given value.

### HasRelayServers

`func (o *NetworkDHCPRelayPolicy) HasRelayServers() bool`

HasRelayServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


