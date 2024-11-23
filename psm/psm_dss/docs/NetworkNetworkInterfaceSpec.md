# NetworkNetworkInterfaceSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminStatus** | Pointer to **string** | desired Admin state of the port. | [optional] [default to "up"]
**AttachNetwork** | Pointer to **string** | AttachNetwork associates the interface with a Network. This is only valid for HOST_PF type. | [optional] 
**AttachTenant** | Pointer to **string** |  | [optional] 
**ConnectionTracking** | Pointer to **bool** | ConnectionTracking enables connection tracking on the interface. This is valid only for HOST_PF type. | [optional] [default to false]
**EnableFwLogging** | Pointer to **bool** | EnableFwLogging enables flow logging on the interface. This is valid only for HOST_PF type. | [optional] 
**IpAllocType** | Pointer to **string** |  | [optional] [default to "none"]
**IpConfig** | Pointer to [**ClusterIPConfig**](clusterIPConfig.md) |  | [optional] 
**MacAddress** | Pointer to **string** | Override system allocated MAC address. Should be a valid MAC address. | [optional] 
**Mtu** | Pointer to **int64** | Mtu of the interface. | [optional] 
**Pause** | Pointer to [**NetworkPauseSpec**](networkPauseSpec.md) |  | [optional] 
**Speed** | Pointer to **string** | Intefaae speed. | [optional] 
**TxPolicer** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | Type specifies the type of interface. | [optional] [default to "none"]
**VnfAttached** | Pointer to **bool** | VNFAttached knob on the interface. This is valid only for HOST_PF type. | [optional] 

## Methods

### NewNetworkNetworkInterfaceSpec

`func NewNetworkNetworkInterfaceSpec() *NetworkNetworkInterfaceSpec`

NewNetworkNetworkInterfaceSpec instantiates a new NetworkNetworkInterfaceSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkNetworkInterfaceSpecWithDefaults

`func NewNetworkNetworkInterfaceSpecWithDefaults() *NetworkNetworkInterfaceSpec`

NewNetworkNetworkInterfaceSpecWithDefaults instantiates a new NetworkNetworkInterfaceSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminStatus

`func (o *NetworkNetworkInterfaceSpec) GetAdminStatus() string`

GetAdminStatus returns the AdminStatus field if non-nil, zero value otherwise.

### GetAdminStatusOk

`func (o *NetworkNetworkInterfaceSpec) GetAdminStatusOk() (*string, bool)`

GetAdminStatusOk returns a tuple with the AdminStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminStatus

`func (o *NetworkNetworkInterfaceSpec) SetAdminStatus(v string)`

SetAdminStatus sets AdminStatus field to given value.

### HasAdminStatus

`func (o *NetworkNetworkInterfaceSpec) HasAdminStatus() bool`

HasAdminStatus returns a boolean if a field has been set.

### GetAttachNetwork

`func (o *NetworkNetworkInterfaceSpec) GetAttachNetwork() string`

GetAttachNetwork returns the AttachNetwork field if non-nil, zero value otherwise.

### GetAttachNetworkOk

`func (o *NetworkNetworkInterfaceSpec) GetAttachNetworkOk() (*string, bool)`

GetAttachNetworkOk returns a tuple with the AttachNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachNetwork

`func (o *NetworkNetworkInterfaceSpec) SetAttachNetwork(v string)`

SetAttachNetwork sets AttachNetwork field to given value.

### HasAttachNetwork

`func (o *NetworkNetworkInterfaceSpec) HasAttachNetwork() bool`

HasAttachNetwork returns a boolean if a field has been set.

### GetAttachTenant

`func (o *NetworkNetworkInterfaceSpec) GetAttachTenant() string`

GetAttachTenant returns the AttachTenant field if non-nil, zero value otherwise.

### GetAttachTenantOk

`func (o *NetworkNetworkInterfaceSpec) GetAttachTenantOk() (*string, bool)`

GetAttachTenantOk returns a tuple with the AttachTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachTenant

`func (o *NetworkNetworkInterfaceSpec) SetAttachTenant(v string)`

SetAttachTenant sets AttachTenant field to given value.

### HasAttachTenant

`func (o *NetworkNetworkInterfaceSpec) HasAttachTenant() bool`

HasAttachTenant returns a boolean if a field has been set.

### GetConnectionTracking

`func (o *NetworkNetworkInterfaceSpec) GetConnectionTracking() bool`

GetConnectionTracking returns the ConnectionTracking field if non-nil, zero value otherwise.

### GetConnectionTrackingOk

`func (o *NetworkNetworkInterfaceSpec) GetConnectionTrackingOk() (*bool, bool)`

GetConnectionTrackingOk returns a tuple with the ConnectionTracking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTracking

`func (o *NetworkNetworkInterfaceSpec) SetConnectionTracking(v bool)`

SetConnectionTracking sets ConnectionTracking field to given value.

### HasConnectionTracking

`func (o *NetworkNetworkInterfaceSpec) HasConnectionTracking() bool`

HasConnectionTracking returns a boolean if a field has been set.

### GetEnableFwLogging

`func (o *NetworkNetworkInterfaceSpec) GetEnableFwLogging() bool`

GetEnableFwLogging returns the EnableFwLogging field if non-nil, zero value otherwise.

### GetEnableFwLoggingOk

`func (o *NetworkNetworkInterfaceSpec) GetEnableFwLoggingOk() (*bool, bool)`

GetEnableFwLoggingOk returns a tuple with the EnableFwLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFwLogging

`func (o *NetworkNetworkInterfaceSpec) SetEnableFwLogging(v bool)`

SetEnableFwLogging sets EnableFwLogging field to given value.

### HasEnableFwLogging

`func (o *NetworkNetworkInterfaceSpec) HasEnableFwLogging() bool`

HasEnableFwLogging returns a boolean if a field has been set.

### GetIpAllocType

`func (o *NetworkNetworkInterfaceSpec) GetIpAllocType() string`

GetIpAllocType returns the IpAllocType field if non-nil, zero value otherwise.

### GetIpAllocTypeOk

`func (o *NetworkNetworkInterfaceSpec) GetIpAllocTypeOk() (*string, bool)`

GetIpAllocTypeOk returns a tuple with the IpAllocType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAllocType

`func (o *NetworkNetworkInterfaceSpec) SetIpAllocType(v string)`

SetIpAllocType sets IpAllocType field to given value.

### HasIpAllocType

`func (o *NetworkNetworkInterfaceSpec) HasIpAllocType() bool`

HasIpAllocType returns a boolean if a field has been set.

### GetIpConfig

`func (o *NetworkNetworkInterfaceSpec) GetIpConfig() ClusterIPConfig`

GetIpConfig returns the IpConfig field if non-nil, zero value otherwise.

### GetIpConfigOk

`func (o *NetworkNetworkInterfaceSpec) GetIpConfigOk() (*ClusterIPConfig, bool)`

GetIpConfigOk returns a tuple with the IpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpConfig

`func (o *NetworkNetworkInterfaceSpec) SetIpConfig(v ClusterIPConfig)`

SetIpConfig sets IpConfig field to given value.

### HasIpConfig

`func (o *NetworkNetworkInterfaceSpec) HasIpConfig() bool`

HasIpConfig returns a boolean if a field has been set.

### GetMacAddress

`func (o *NetworkNetworkInterfaceSpec) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *NetworkNetworkInterfaceSpec) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *NetworkNetworkInterfaceSpec) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *NetworkNetworkInterfaceSpec) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetMtu

`func (o *NetworkNetworkInterfaceSpec) GetMtu() int64`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *NetworkNetworkInterfaceSpec) GetMtuOk() (*int64, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *NetworkNetworkInterfaceSpec) SetMtu(v int64)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *NetworkNetworkInterfaceSpec) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetPause

`func (o *NetworkNetworkInterfaceSpec) GetPause() NetworkPauseSpec`

GetPause returns the Pause field if non-nil, zero value otherwise.

### GetPauseOk

`func (o *NetworkNetworkInterfaceSpec) GetPauseOk() (*NetworkPauseSpec, bool)`

GetPauseOk returns a tuple with the Pause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPause

`func (o *NetworkNetworkInterfaceSpec) SetPause(v NetworkPauseSpec)`

SetPause sets Pause field to given value.

### HasPause

`func (o *NetworkNetworkInterfaceSpec) HasPause() bool`

HasPause returns a boolean if a field has been set.

### GetSpeed

`func (o *NetworkNetworkInterfaceSpec) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *NetworkNetworkInterfaceSpec) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *NetworkNetworkInterfaceSpec) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *NetworkNetworkInterfaceSpec) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetTxPolicer

`func (o *NetworkNetworkInterfaceSpec) GetTxPolicer() string`

GetTxPolicer returns the TxPolicer field if non-nil, zero value otherwise.

### GetTxPolicerOk

`func (o *NetworkNetworkInterfaceSpec) GetTxPolicerOk() (*string, bool)`

GetTxPolicerOk returns a tuple with the TxPolicer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxPolicer

`func (o *NetworkNetworkInterfaceSpec) SetTxPolicer(v string)`

SetTxPolicer sets TxPolicer field to given value.

### HasTxPolicer

`func (o *NetworkNetworkInterfaceSpec) HasTxPolicer() bool`

HasTxPolicer returns a boolean if a field has been set.

### GetType

`func (o *NetworkNetworkInterfaceSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkNetworkInterfaceSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkNetworkInterfaceSpec) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NetworkNetworkInterfaceSpec) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVnfAttached

`func (o *NetworkNetworkInterfaceSpec) GetVnfAttached() bool`

GetVnfAttached returns the VnfAttached field if non-nil, zero value otherwise.

### GetVnfAttachedOk

`func (o *NetworkNetworkInterfaceSpec) GetVnfAttachedOk() (*bool, bool)`

GetVnfAttachedOk returns a tuple with the VnfAttached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnfAttached

`func (o *NetworkNetworkInterfaceSpec) SetVnfAttached(v bool)`

SetVnfAttached sets VnfAttached field to given value.

### HasVnfAttached

`func (o *NetworkNetworkInterfaceSpec) HasVnfAttached() bool`

HasVnfAttached returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


