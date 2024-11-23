# NetworkNetworkFirewallProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaximumCpsPerDistributedServicesEntity** | Pointer to **int32** | Maximum Connections Per Second supported for the Network within a Distributed Services Entity. The value configured here overrides the MaxCPSPerNetworkPerDistributedServicesEntity configuration in the Virtual Router that the Network belongs to. Value 0 means the CPS limit is not enforced and the CPS is limited only by the system capacity. Connections exceeding the CPS limit are dropped. Value should be between -1 and 409599. | [optional] [default to -1]
**MaximumSessionsPerDistributedServicesEntity** | Pointer to **int32** | Maximum sessions supported in the Network within a Distributed Services Entity. The value configured here overrides the MaxSessionsPerNetworkPerDistributedServicesEntity configuration in the Virtual Router that the Network belongs to. Value 0 means the sessions limit is not enforced and the number of sessions is limited only by the system capacity. Sessions exceeding the sessions limit are dropped. Value should be between -1 and 16777215. | [optional] [default to -1]

## Methods

### NewNetworkNetworkFirewallProfile

`func NewNetworkNetworkFirewallProfile() *NetworkNetworkFirewallProfile`

NewNetworkNetworkFirewallProfile instantiates a new NetworkNetworkFirewallProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkNetworkFirewallProfileWithDefaults

`func NewNetworkNetworkFirewallProfileWithDefaults() *NetworkNetworkFirewallProfile`

NewNetworkNetworkFirewallProfileWithDefaults instantiates a new NetworkNetworkFirewallProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaximumCpsPerDistributedServicesEntity

`func (o *NetworkNetworkFirewallProfile) GetMaximumCpsPerDistributedServicesEntity() int32`

GetMaximumCpsPerDistributedServicesEntity returns the MaximumCpsPerDistributedServicesEntity field if non-nil, zero value otherwise.

### GetMaximumCpsPerDistributedServicesEntityOk

`func (o *NetworkNetworkFirewallProfile) GetMaximumCpsPerDistributedServicesEntityOk() (*int32, bool)`

GetMaximumCpsPerDistributedServicesEntityOk returns a tuple with the MaximumCpsPerDistributedServicesEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumCpsPerDistributedServicesEntity

`func (o *NetworkNetworkFirewallProfile) SetMaximumCpsPerDistributedServicesEntity(v int32)`

SetMaximumCpsPerDistributedServicesEntity sets MaximumCpsPerDistributedServicesEntity field to given value.

### HasMaximumCpsPerDistributedServicesEntity

`func (o *NetworkNetworkFirewallProfile) HasMaximumCpsPerDistributedServicesEntity() bool`

HasMaximumCpsPerDistributedServicesEntity returns a boolean if a field has been set.

### GetMaximumSessionsPerDistributedServicesEntity

`func (o *NetworkNetworkFirewallProfile) GetMaximumSessionsPerDistributedServicesEntity() int32`

GetMaximumSessionsPerDistributedServicesEntity returns the MaximumSessionsPerDistributedServicesEntity field if non-nil, zero value otherwise.

### GetMaximumSessionsPerDistributedServicesEntityOk

`func (o *NetworkNetworkFirewallProfile) GetMaximumSessionsPerDistributedServicesEntityOk() (*int32, bool)`

GetMaximumSessionsPerDistributedServicesEntityOk returns a tuple with the MaximumSessionsPerDistributedServicesEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumSessionsPerDistributedServicesEntity

`func (o *NetworkNetworkFirewallProfile) SetMaximumSessionsPerDistributedServicesEntity(v int32)`

SetMaximumSessionsPerDistributedServicesEntity sets MaximumSessionsPerDistributedServicesEntity field to given value.

### HasMaximumSessionsPerDistributedServicesEntity

`func (o *NetworkNetworkFirewallProfile) HasMaximumSessionsPerDistributedServicesEntity() bool`

HasMaximumSessionsPerDistributedServicesEntity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


