# NetworkNetworkSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EgressSecurityPolicy** | Pointer to **[]string** | Security Policy to apply in the egress direction. | [optional] 
**FirewallProfile** | Pointer to [**NetworkNetworkFirewallProfile**](networkNetworkFirewallProfile.md) |  | [optional] 
**IngressSecurityPolicy** | Pointer to **[]string** | Security Policy to apply in the ingress direction. | [optional] 
**IpamPolicy** | Pointer to **string** | Relay Configuration if any. | [optional] 
**Ipv4Gateway** | Pointer to **string** | IPv4 gateway for this subnet. Should be a valid v4 or v6 IP address. | [optional] 
**Ipv4Subnet** | Pointer to **string** | IPv4 subnet CIDR. Should be a valid v4 or v6 CIDR block. | [optional] 
**Ipv6Gateway** | Pointer to **string** | IPv6 gateway. | [optional] 
**Ipv6Subnet** | Pointer to **string** | IPv6 subnet CIDR. | [optional] 
**Orchestrators** | Pointer to [**[]NetworkOrchestratorInfo**](NetworkOrchestratorInfo.md) | If supplied, this network will only be applied to the orchestrators specified. | [optional] 
**RouteImportExport** | Pointer to [**NetworkRDSpec**](networkRDSpec.md) |  | [optional] 
**Type** | Pointer to **string** | type of network. (vlan/vxlan/routed etc). | [optional] [default to "bridged"]
**VirtualRouter** | Pointer to **string** | VirtualRouter specifies the VRF this network belongs to. | [optional] 
**VlanId** | Pointer to **int64** | Vlan ID for the network. Value should be between 0 and 4095. | [optional] 
**VxlanVni** | Pointer to **int64** | Vxlan VNI for the network. Value should be between 0 and 16777215. | [optional] 

## Methods

### NewNetworkNetworkSpec

`func NewNetworkNetworkSpec() *NetworkNetworkSpec`

NewNetworkNetworkSpec instantiates a new NetworkNetworkSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkNetworkSpecWithDefaults

`func NewNetworkNetworkSpecWithDefaults() *NetworkNetworkSpec`

NewNetworkNetworkSpecWithDefaults instantiates a new NetworkNetworkSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEgressSecurityPolicy

`func (o *NetworkNetworkSpec) GetEgressSecurityPolicy() []string`

GetEgressSecurityPolicy returns the EgressSecurityPolicy field if non-nil, zero value otherwise.

### GetEgressSecurityPolicyOk

`func (o *NetworkNetworkSpec) GetEgressSecurityPolicyOk() (*[]string, bool)`

GetEgressSecurityPolicyOk returns a tuple with the EgressSecurityPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgressSecurityPolicy

`func (o *NetworkNetworkSpec) SetEgressSecurityPolicy(v []string)`

SetEgressSecurityPolicy sets EgressSecurityPolicy field to given value.

### HasEgressSecurityPolicy

`func (o *NetworkNetworkSpec) HasEgressSecurityPolicy() bool`

HasEgressSecurityPolicy returns a boolean if a field has been set.

### GetFirewallProfile

`func (o *NetworkNetworkSpec) GetFirewallProfile() NetworkNetworkFirewallProfile`

GetFirewallProfile returns the FirewallProfile field if non-nil, zero value otherwise.

### GetFirewallProfileOk

`func (o *NetworkNetworkSpec) GetFirewallProfileOk() (*NetworkNetworkFirewallProfile, bool)`

GetFirewallProfileOk returns a tuple with the FirewallProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallProfile

`func (o *NetworkNetworkSpec) SetFirewallProfile(v NetworkNetworkFirewallProfile)`

SetFirewallProfile sets FirewallProfile field to given value.

### HasFirewallProfile

`func (o *NetworkNetworkSpec) HasFirewallProfile() bool`

HasFirewallProfile returns a boolean if a field has been set.

### GetIngressSecurityPolicy

`func (o *NetworkNetworkSpec) GetIngressSecurityPolicy() []string`

GetIngressSecurityPolicy returns the IngressSecurityPolicy field if non-nil, zero value otherwise.

### GetIngressSecurityPolicyOk

`func (o *NetworkNetworkSpec) GetIngressSecurityPolicyOk() (*[]string, bool)`

GetIngressSecurityPolicyOk returns a tuple with the IngressSecurityPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressSecurityPolicy

`func (o *NetworkNetworkSpec) SetIngressSecurityPolicy(v []string)`

SetIngressSecurityPolicy sets IngressSecurityPolicy field to given value.

### HasIngressSecurityPolicy

`func (o *NetworkNetworkSpec) HasIngressSecurityPolicy() bool`

HasIngressSecurityPolicy returns a boolean if a field has been set.

### GetIpamPolicy

`func (o *NetworkNetworkSpec) GetIpamPolicy() string`

GetIpamPolicy returns the IpamPolicy field if non-nil, zero value otherwise.

### GetIpamPolicyOk

`func (o *NetworkNetworkSpec) GetIpamPolicyOk() (*string, bool)`

GetIpamPolicyOk returns a tuple with the IpamPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpamPolicy

`func (o *NetworkNetworkSpec) SetIpamPolicy(v string)`

SetIpamPolicy sets IpamPolicy field to given value.

### HasIpamPolicy

`func (o *NetworkNetworkSpec) HasIpamPolicy() bool`

HasIpamPolicy returns a boolean if a field has been set.

### GetIpv4Gateway

`func (o *NetworkNetworkSpec) GetIpv4Gateway() string`

GetIpv4Gateway returns the Ipv4Gateway field if non-nil, zero value otherwise.

### GetIpv4GatewayOk

`func (o *NetworkNetworkSpec) GetIpv4GatewayOk() (*string, bool)`

GetIpv4GatewayOk returns a tuple with the Ipv4Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Gateway

`func (o *NetworkNetworkSpec) SetIpv4Gateway(v string)`

SetIpv4Gateway sets Ipv4Gateway field to given value.

### HasIpv4Gateway

`func (o *NetworkNetworkSpec) HasIpv4Gateway() bool`

HasIpv4Gateway returns a boolean if a field has been set.

### GetIpv4Subnet

`func (o *NetworkNetworkSpec) GetIpv4Subnet() string`

GetIpv4Subnet returns the Ipv4Subnet field if non-nil, zero value otherwise.

### GetIpv4SubnetOk

`func (o *NetworkNetworkSpec) GetIpv4SubnetOk() (*string, bool)`

GetIpv4SubnetOk returns a tuple with the Ipv4Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Subnet

`func (o *NetworkNetworkSpec) SetIpv4Subnet(v string)`

SetIpv4Subnet sets Ipv4Subnet field to given value.

### HasIpv4Subnet

`func (o *NetworkNetworkSpec) HasIpv4Subnet() bool`

HasIpv4Subnet returns a boolean if a field has been set.

### GetIpv6Gateway

`func (o *NetworkNetworkSpec) GetIpv6Gateway() string`

GetIpv6Gateway returns the Ipv6Gateway field if non-nil, zero value otherwise.

### GetIpv6GatewayOk

`func (o *NetworkNetworkSpec) GetIpv6GatewayOk() (*string, bool)`

GetIpv6GatewayOk returns a tuple with the Ipv6Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Gateway

`func (o *NetworkNetworkSpec) SetIpv6Gateway(v string)`

SetIpv6Gateway sets Ipv6Gateway field to given value.

### HasIpv6Gateway

`func (o *NetworkNetworkSpec) HasIpv6Gateway() bool`

HasIpv6Gateway returns a boolean if a field has been set.

### GetIpv6Subnet

`func (o *NetworkNetworkSpec) GetIpv6Subnet() string`

GetIpv6Subnet returns the Ipv6Subnet field if non-nil, zero value otherwise.

### GetIpv6SubnetOk

`func (o *NetworkNetworkSpec) GetIpv6SubnetOk() (*string, bool)`

GetIpv6SubnetOk returns a tuple with the Ipv6Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Subnet

`func (o *NetworkNetworkSpec) SetIpv6Subnet(v string)`

SetIpv6Subnet sets Ipv6Subnet field to given value.

### HasIpv6Subnet

`func (o *NetworkNetworkSpec) HasIpv6Subnet() bool`

HasIpv6Subnet returns a boolean if a field has been set.

### GetOrchestrators

`func (o *NetworkNetworkSpec) GetOrchestrators() []NetworkOrchestratorInfo`

GetOrchestrators returns the Orchestrators field if non-nil, zero value otherwise.

### GetOrchestratorsOk

`func (o *NetworkNetworkSpec) GetOrchestratorsOk() (*[]NetworkOrchestratorInfo, bool)`

GetOrchestratorsOk returns a tuple with the Orchestrators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrchestrators

`func (o *NetworkNetworkSpec) SetOrchestrators(v []NetworkOrchestratorInfo)`

SetOrchestrators sets Orchestrators field to given value.

### HasOrchestrators

`func (o *NetworkNetworkSpec) HasOrchestrators() bool`

HasOrchestrators returns a boolean if a field has been set.

### GetRouteImportExport

`func (o *NetworkNetworkSpec) GetRouteImportExport() NetworkRDSpec`

GetRouteImportExport returns the RouteImportExport field if non-nil, zero value otherwise.

### GetRouteImportExportOk

`func (o *NetworkNetworkSpec) GetRouteImportExportOk() (*NetworkRDSpec, bool)`

GetRouteImportExportOk returns a tuple with the RouteImportExport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteImportExport

`func (o *NetworkNetworkSpec) SetRouteImportExport(v NetworkRDSpec)`

SetRouteImportExport sets RouteImportExport field to given value.

### HasRouteImportExport

`func (o *NetworkNetworkSpec) HasRouteImportExport() bool`

HasRouteImportExport returns a boolean if a field has been set.

### GetType

`func (o *NetworkNetworkSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkNetworkSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkNetworkSpec) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NetworkNetworkSpec) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVirtualRouter

`func (o *NetworkNetworkSpec) GetVirtualRouter() string`

GetVirtualRouter returns the VirtualRouter field if non-nil, zero value otherwise.

### GetVirtualRouterOk

`func (o *NetworkNetworkSpec) GetVirtualRouterOk() (*string, bool)`

GetVirtualRouterOk returns a tuple with the VirtualRouter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualRouter

`func (o *NetworkNetworkSpec) SetVirtualRouter(v string)`

SetVirtualRouter sets VirtualRouter field to given value.

### HasVirtualRouter

`func (o *NetworkNetworkSpec) HasVirtualRouter() bool`

HasVirtualRouter returns a boolean if a field has been set.

### GetVlanId

`func (o *NetworkNetworkSpec) GetVlanId() int64`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *NetworkNetworkSpec) GetVlanIdOk() (*int64, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *NetworkNetworkSpec) SetVlanId(v int64)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *NetworkNetworkSpec) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetVxlanVni

`func (o *NetworkNetworkSpec) GetVxlanVni() int64`

GetVxlanVni returns the VxlanVni field if non-nil, zero value otherwise.

### GetVxlanVniOk

`func (o *NetworkNetworkSpec) GetVxlanVniOk() (*int64, bool)`

GetVxlanVniOk returns a tuple with the VxlanVni field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVxlanVni

`func (o *NetworkNetworkSpec) SetVxlanVni(v int64)`

SetVxlanVni sets VxlanVni field to given value.

### HasVxlanVni

`func (o *NetworkNetworkSpec) HasVxlanVni() bool`

HasVxlanVni returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


