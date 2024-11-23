# WorkloadEndpointStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndpointState** | Pointer to **string** | endpoint FSM state. | [optional] 
**SecurityGroups** | Pointer to **[]string** | security groups. | [optional] 
**HomingHostAddr** | Pointer to **string** | host address of the host where this endpoint exists. | [optional] 
**HomingHostName** | Pointer to **string** | host name of the host where this endpoint exists. | [optional] 
**Ipv4Address** | Pointer to **string** | IPv4 address of the endpoint. | [optional] 
**Ipv4Addresses** | Pointer to **[]string** | IPv4 addresses of the endpoint. | [optional] 
**Ipv4Gateway** | Pointer to **string** | IPv4 gateway for the endpoint. | [optional] 
**Ipv4Gateways** | Pointer to **[]string** | IPv4 gateways for the endpoint. | [optional] 
**Ipv6Address** | Pointer to **string** | IPv6 address for the endpoint. | [optional] 
**Ipv6Addresses** | Pointer to **[]string** | IPv6 addresses for the endpoint. | [optional] 
**Ipv6Gateway** | Pointer to **string** | IPv6 gateway. | [optional] 
**Ipv6Gateways** | Pointer to **[]string** | IPv6 gateways. | [optional] 
**MacAddress** | Pointer to **string** | Mac address of the endpoint. Should be a valid MAC address. | [optional] 
**MicroSegmentVlan** | Pointer to **int64** | micro-segment VLAN. | [optional] 
**Migration** | Pointer to [**WorkloadEndpointMigrationStatus**](workloadEndpointMigrationStatus.md) |  | [optional] 
**MirrorSessions** | Pointer to **[]string** | MirrorSessions list of mirror sessions enabled on this endpoint. | [optional] 
**Network** | Pointer to **string** | network this endpoint belogs to. | [optional] 
**NodeUuid** | Pointer to **string** | homing host&#39;s UUID. | [optional] 
**NodeUuidList** | Pointer to **[]string** | NodeUUIDList has the list of DSCs where a EP is supposed to go to. | [optional] 
**WorkloadAttributes** | Pointer to **map[string]string** | VM or container attribute/labels. | [optional] 
**WorkloadName** | Pointer to **string** | VM or container name. | [optional] 
**WorkloadNames** | Pointer to **[]string** | WorkloadNames associated with endpoint. | [optional] 

## Methods

### NewWorkloadEndpointStatus

`func NewWorkloadEndpointStatus() *WorkloadEndpointStatus`

NewWorkloadEndpointStatus instantiates a new WorkloadEndpointStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadEndpointStatusWithDefaults

`func NewWorkloadEndpointStatusWithDefaults() *WorkloadEndpointStatus`

NewWorkloadEndpointStatusWithDefaults instantiates a new WorkloadEndpointStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpointState

`func (o *WorkloadEndpointStatus) GetEndpointState() string`

GetEndpointState returns the EndpointState field if non-nil, zero value otherwise.

### GetEndpointStateOk

`func (o *WorkloadEndpointStatus) GetEndpointStateOk() (*string, bool)`

GetEndpointStateOk returns a tuple with the EndpointState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointState

`func (o *WorkloadEndpointStatus) SetEndpointState(v string)`

SetEndpointState sets EndpointState field to given value.

### HasEndpointState

`func (o *WorkloadEndpointStatus) HasEndpointState() bool`

HasEndpointState returns a boolean if a field has been set.

### GetSecurityGroups

`func (o *WorkloadEndpointStatus) GetSecurityGroups() []string`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *WorkloadEndpointStatus) GetSecurityGroupsOk() (*[]string, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *WorkloadEndpointStatus) SetSecurityGroups(v []string)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *WorkloadEndpointStatus) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### GetHomingHostAddr

`func (o *WorkloadEndpointStatus) GetHomingHostAddr() string`

GetHomingHostAddr returns the HomingHostAddr field if non-nil, zero value otherwise.

### GetHomingHostAddrOk

`func (o *WorkloadEndpointStatus) GetHomingHostAddrOk() (*string, bool)`

GetHomingHostAddrOk returns a tuple with the HomingHostAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomingHostAddr

`func (o *WorkloadEndpointStatus) SetHomingHostAddr(v string)`

SetHomingHostAddr sets HomingHostAddr field to given value.

### HasHomingHostAddr

`func (o *WorkloadEndpointStatus) HasHomingHostAddr() bool`

HasHomingHostAddr returns a boolean if a field has been set.

### GetHomingHostName

`func (o *WorkloadEndpointStatus) GetHomingHostName() string`

GetHomingHostName returns the HomingHostName field if non-nil, zero value otherwise.

### GetHomingHostNameOk

`func (o *WorkloadEndpointStatus) GetHomingHostNameOk() (*string, bool)`

GetHomingHostNameOk returns a tuple with the HomingHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomingHostName

`func (o *WorkloadEndpointStatus) SetHomingHostName(v string)`

SetHomingHostName sets HomingHostName field to given value.

### HasHomingHostName

`func (o *WorkloadEndpointStatus) HasHomingHostName() bool`

HasHomingHostName returns a boolean if a field has been set.

### GetIpv4Address

`func (o *WorkloadEndpointStatus) GetIpv4Address() string`

GetIpv4Address returns the Ipv4Address field if non-nil, zero value otherwise.

### GetIpv4AddressOk

`func (o *WorkloadEndpointStatus) GetIpv4AddressOk() (*string, bool)`

GetIpv4AddressOk returns a tuple with the Ipv4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Address

`func (o *WorkloadEndpointStatus) SetIpv4Address(v string)`

SetIpv4Address sets Ipv4Address field to given value.

### HasIpv4Address

`func (o *WorkloadEndpointStatus) HasIpv4Address() bool`

HasIpv4Address returns a boolean if a field has been set.

### GetIpv4Addresses

`func (o *WorkloadEndpointStatus) GetIpv4Addresses() []string`

GetIpv4Addresses returns the Ipv4Addresses field if non-nil, zero value otherwise.

### GetIpv4AddressesOk

`func (o *WorkloadEndpointStatus) GetIpv4AddressesOk() (*[]string, bool)`

GetIpv4AddressesOk returns a tuple with the Ipv4Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Addresses

`func (o *WorkloadEndpointStatus) SetIpv4Addresses(v []string)`

SetIpv4Addresses sets Ipv4Addresses field to given value.

### HasIpv4Addresses

`func (o *WorkloadEndpointStatus) HasIpv4Addresses() bool`

HasIpv4Addresses returns a boolean if a field has been set.

### GetIpv4Gateway

`func (o *WorkloadEndpointStatus) GetIpv4Gateway() string`

GetIpv4Gateway returns the Ipv4Gateway field if non-nil, zero value otherwise.

### GetIpv4GatewayOk

`func (o *WorkloadEndpointStatus) GetIpv4GatewayOk() (*string, bool)`

GetIpv4GatewayOk returns a tuple with the Ipv4Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Gateway

`func (o *WorkloadEndpointStatus) SetIpv4Gateway(v string)`

SetIpv4Gateway sets Ipv4Gateway field to given value.

### HasIpv4Gateway

`func (o *WorkloadEndpointStatus) HasIpv4Gateway() bool`

HasIpv4Gateway returns a boolean if a field has been set.

### GetIpv4Gateways

`func (o *WorkloadEndpointStatus) GetIpv4Gateways() []string`

GetIpv4Gateways returns the Ipv4Gateways field if non-nil, zero value otherwise.

### GetIpv4GatewaysOk

`func (o *WorkloadEndpointStatus) GetIpv4GatewaysOk() (*[]string, bool)`

GetIpv4GatewaysOk returns a tuple with the Ipv4Gateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Gateways

`func (o *WorkloadEndpointStatus) SetIpv4Gateways(v []string)`

SetIpv4Gateways sets Ipv4Gateways field to given value.

### HasIpv4Gateways

`func (o *WorkloadEndpointStatus) HasIpv4Gateways() bool`

HasIpv4Gateways returns a boolean if a field has been set.

### GetIpv6Address

`func (o *WorkloadEndpointStatus) GetIpv6Address() string`

GetIpv6Address returns the Ipv6Address field if non-nil, zero value otherwise.

### GetIpv6AddressOk

`func (o *WorkloadEndpointStatus) GetIpv6AddressOk() (*string, bool)`

GetIpv6AddressOk returns a tuple with the Ipv6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Address

`func (o *WorkloadEndpointStatus) SetIpv6Address(v string)`

SetIpv6Address sets Ipv6Address field to given value.

### HasIpv6Address

`func (o *WorkloadEndpointStatus) HasIpv6Address() bool`

HasIpv6Address returns a boolean if a field has been set.

### GetIpv6Addresses

`func (o *WorkloadEndpointStatus) GetIpv6Addresses() []string`

GetIpv6Addresses returns the Ipv6Addresses field if non-nil, zero value otherwise.

### GetIpv6AddressesOk

`func (o *WorkloadEndpointStatus) GetIpv6AddressesOk() (*[]string, bool)`

GetIpv6AddressesOk returns a tuple with the Ipv6Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Addresses

`func (o *WorkloadEndpointStatus) SetIpv6Addresses(v []string)`

SetIpv6Addresses sets Ipv6Addresses field to given value.

### HasIpv6Addresses

`func (o *WorkloadEndpointStatus) HasIpv6Addresses() bool`

HasIpv6Addresses returns a boolean if a field has been set.

### GetIpv6Gateway

`func (o *WorkloadEndpointStatus) GetIpv6Gateway() string`

GetIpv6Gateway returns the Ipv6Gateway field if non-nil, zero value otherwise.

### GetIpv6GatewayOk

`func (o *WorkloadEndpointStatus) GetIpv6GatewayOk() (*string, bool)`

GetIpv6GatewayOk returns a tuple with the Ipv6Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Gateway

`func (o *WorkloadEndpointStatus) SetIpv6Gateway(v string)`

SetIpv6Gateway sets Ipv6Gateway field to given value.

### HasIpv6Gateway

`func (o *WorkloadEndpointStatus) HasIpv6Gateway() bool`

HasIpv6Gateway returns a boolean if a field has been set.

### GetIpv6Gateways

`func (o *WorkloadEndpointStatus) GetIpv6Gateways() []string`

GetIpv6Gateways returns the Ipv6Gateways field if non-nil, zero value otherwise.

### GetIpv6GatewaysOk

`func (o *WorkloadEndpointStatus) GetIpv6GatewaysOk() (*[]string, bool)`

GetIpv6GatewaysOk returns a tuple with the Ipv6Gateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Gateways

`func (o *WorkloadEndpointStatus) SetIpv6Gateways(v []string)`

SetIpv6Gateways sets Ipv6Gateways field to given value.

### HasIpv6Gateways

`func (o *WorkloadEndpointStatus) HasIpv6Gateways() bool`

HasIpv6Gateways returns a boolean if a field has been set.

### GetMacAddress

`func (o *WorkloadEndpointStatus) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *WorkloadEndpointStatus) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *WorkloadEndpointStatus) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *WorkloadEndpointStatus) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetMicroSegmentVlan

`func (o *WorkloadEndpointStatus) GetMicroSegmentVlan() int64`

GetMicroSegmentVlan returns the MicroSegmentVlan field if non-nil, zero value otherwise.

### GetMicroSegmentVlanOk

`func (o *WorkloadEndpointStatus) GetMicroSegmentVlanOk() (*int64, bool)`

GetMicroSegmentVlanOk returns a tuple with the MicroSegmentVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroSegmentVlan

`func (o *WorkloadEndpointStatus) SetMicroSegmentVlan(v int64)`

SetMicroSegmentVlan sets MicroSegmentVlan field to given value.

### HasMicroSegmentVlan

`func (o *WorkloadEndpointStatus) HasMicroSegmentVlan() bool`

HasMicroSegmentVlan returns a boolean if a field has been set.

### GetMigration

`func (o *WorkloadEndpointStatus) GetMigration() WorkloadEndpointMigrationStatus`

GetMigration returns the Migration field if non-nil, zero value otherwise.

### GetMigrationOk

`func (o *WorkloadEndpointStatus) GetMigrationOk() (*WorkloadEndpointMigrationStatus, bool)`

GetMigrationOk returns a tuple with the Migration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigration

`func (o *WorkloadEndpointStatus) SetMigration(v WorkloadEndpointMigrationStatus)`

SetMigration sets Migration field to given value.

### HasMigration

`func (o *WorkloadEndpointStatus) HasMigration() bool`

HasMigration returns a boolean if a field has been set.

### GetMirrorSessions

`func (o *WorkloadEndpointStatus) GetMirrorSessions() []string`

GetMirrorSessions returns the MirrorSessions field if non-nil, zero value otherwise.

### GetMirrorSessionsOk

`func (o *WorkloadEndpointStatus) GetMirrorSessionsOk() (*[]string, bool)`

GetMirrorSessionsOk returns a tuple with the MirrorSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirrorSessions

`func (o *WorkloadEndpointStatus) SetMirrorSessions(v []string)`

SetMirrorSessions sets MirrorSessions field to given value.

### HasMirrorSessions

`func (o *WorkloadEndpointStatus) HasMirrorSessions() bool`

HasMirrorSessions returns a boolean if a field has been set.

### GetNetwork

`func (o *WorkloadEndpointStatus) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *WorkloadEndpointStatus) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *WorkloadEndpointStatus) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *WorkloadEndpointStatus) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNodeUuid

`func (o *WorkloadEndpointStatus) GetNodeUuid() string`

GetNodeUuid returns the NodeUuid field if non-nil, zero value otherwise.

### GetNodeUuidOk

`func (o *WorkloadEndpointStatus) GetNodeUuidOk() (*string, bool)`

GetNodeUuidOk returns a tuple with the NodeUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeUuid

`func (o *WorkloadEndpointStatus) SetNodeUuid(v string)`

SetNodeUuid sets NodeUuid field to given value.

### HasNodeUuid

`func (o *WorkloadEndpointStatus) HasNodeUuid() bool`

HasNodeUuid returns a boolean if a field has been set.

### GetNodeUuidList

`func (o *WorkloadEndpointStatus) GetNodeUuidList() []string`

GetNodeUuidList returns the NodeUuidList field if non-nil, zero value otherwise.

### GetNodeUuidListOk

`func (o *WorkloadEndpointStatus) GetNodeUuidListOk() (*[]string, bool)`

GetNodeUuidListOk returns a tuple with the NodeUuidList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeUuidList

`func (o *WorkloadEndpointStatus) SetNodeUuidList(v []string)`

SetNodeUuidList sets NodeUuidList field to given value.

### HasNodeUuidList

`func (o *WorkloadEndpointStatus) HasNodeUuidList() bool`

HasNodeUuidList returns a boolean if a field has been set.

### GetWorkloadAttributes

`func (o *WorkloadEndpointStatus) GetWorkloadAttributes() map[string]string`

GetWorkloadAttributes returns the WorkloadAttributes field if non-nil, zero value otherwise.

### GetWorkloadAttributesOk

`func (o *WorkloadEndpointStatus) GetWorkloadAttributesOk() (*map[string]string, bool)`

GetWorkloadAttributesOk returns a tuple with the WorkloadAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadAttributes

`func (o *WorkloadEndpointStatus) SetWorkloadAttributes(v map[string]string)`

SetWorkloadAttributes sets WorkloadAttributes field to given value.

### HasWorkloadAttributes

`func (o *WorkloadEndpointStatus) HasWorkloadAttributes() bool`

HasWorkloadAttributes returns a boolean if a field has been set.

### GetWorkloadName

`func (o *WorkloadEndpointStatus) GetWorkloadName() string`

GetWorkloadName returns the WorkloadName field if non-nil, zero value otherwise.

### GetWorkloadNameOk

`func (o *WorkloadEndpointStatus) GetWorkloadNameOk() (*string, bool)`

GetWorkloadNameOk returns a tuple with the WorkloadName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadName

`func (o *WorkloadEndpointStatus) SetWorkloadName(v string)`

SetWorkloadName sets WorkloadName field to given value.

### HasWorkloadName

`func (o *WorkloadEndpointStatus) HasWorkloadName() bool`

HasWorkloadName returns a boolean if a field has been set.

### GetWorkloadNames

`func (o *WorkloadEndpointStatus) GetWorkloadNames() []string`

GetWorkloadNames returns the WorkloadNames field if non-nil, zero value otherwise.

### GetWorkloadNamesOk

`func (o *WorkloadEndpointStatus) GetWorkloadNamesOk() (*[]string, bool)`

GetWorkloadNamesOk returns a tuple with the WorkloadNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadNames

`func (o *WorkloadEndpointStatus) SetWorkloadNames(v []string)`

SetWorkloadNames sets WorkloadNames field to given value.

### HasWorkloadNames

`func (o *WorkloadEndpointStatus) HasWorkloadNames() bool`

HasWorkloadNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


