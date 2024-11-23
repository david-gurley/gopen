# NetworkDHCPServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAddress** | Pointer to **string** | IP Address of the server. Should be a valid v4 or v6 IP address. | [optional] 
**VirtualRouter** | Pointer to **string** | Destination VRF where the server is connected. An empty value specifies that the server is reachable in the same vrf as the one where the policy is attached. | [optional] 

## Methods

### NewNetworkDHCPServer

`func NewNetworkDHCPServer() *NetworkDHCPServer`

NewNetworkDHCPServer instantiates a new NetworkDHCPServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDHCPServerWithDefaults

`func NewNetworkDHCPServerWithDefaults() *NetworkDHCPServer`

NewNetworkDHCPServerWithDefaults instantiates a new NetworkDHCPServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpAddress

`func (o *NetworkDHCPServer) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *NetworkDHCPServer) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *NetworkDHCPServer) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *NetworkDHCPServer) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetVirtualRouter

`func (o *NetworkDHCPServer) GetVirtualRouter() string`

GetVirtualRouter returns the VirtualRouter field if non-nil, zero value otherwise.

### GetVirtualRouterOk

`func (o *NetworkDHCPServer) GetVirtualRouterOk() (*string, bool)`

GetVirtualRouterOk returns a tuple with the VirtualRouter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualRouter

`func (o *NetworkDHCPServer) SetVirtualRouter(v string)`

SetVirtualRouter sets VirtualRouter field to given value.

### HasVirtualRouter

`func (o *NetworkDHCPServer) HasVirtualRouter() bool`

HasVirtualRouter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


