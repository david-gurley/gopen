# OrchestrationManagedNamespaceSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiscoveryOperation** | Pointer to **string** |  | [optional] [default to "disc-op-default"]
**DiscoveryProtocol** | Pointer to **string** |  | [optional] [default to "disc-proto-default"]
**Mtu** | Pointer to **int64** | MTU, 0 &#x3D; Use system-default or orchestrator settings. | [optional] [default to 0]
**MulticastFilter** | Pointer to **string** |  | [optional] [default to "mcast-filter-default"]
**OverrideVlanEnd** | Pointer to **int64** | End of vlan range (inclusive) to be used for overrides. Range should be &gt;&#x3D; number of vnics expected per host. Value should be between 0 and 4095. | [optional] [default to 4087]
**OverrideVlanStart** | Pointer to **int64** | Start of vlan range to be used for overrides. Range should be &gt;&#x3D; number of vnics expected per host. Value should be between 0 and 4095. | [optional] [default to 3832]

## Methods

### NewOrchestrationManagedNamespaceSpec

`func NewOrchestrationManagedNamespaceSpec() *OrchestrationManagedNamespaceSpec`

NewOrchestrationManagedNamespaceSpec instantiates a new OrchestrationManagedNamespaceSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrchestrationManagedNamespaceSpecWithDefaults

`func NewOrchestrationManagedNamespaceSpecWithDefaults() *OrchestrationManagedNamespaceSpec`

NewOrchestrationManagedNamespaceSpecWithDefaults instantiates a new OrchestrationManagedNamespaceSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscoveryOperation

`func (o *OrchestrationManagedNamespaceSpec) GetDiscoveryOperation() string`

GetDiscoveryOperation returns the DiscoveryOperation field if non-nil, zero value otherwise.

### GetDiscoveryOperationOk

`func (o *OrchestrationManagedNamespaceSpec) GetDiscoveryOperationOk() (*string, bool)`

GetDiscoveryOperationOk returns a tuple with the DiscoveryOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryOperation

`func (o *OrchestrationManagedNamespaceSpec) SetDiscoveryOperation(v string)`

SetDiscoveryOperation sets DiscoveryOperation field to given value.

### HasDiscoveryOperation

`func (o *OrchestrationManagedNamespaceSpec) HasDiscoveryOperation() bool`

HasDiscoveryOperation returns a boolean if a field has been set.

### GetDiscoveryProtocol

`func (o *OrchestrationManagedNamespaceSpec) GetDiscoveryProtocol() string`

GetDiscoveryProtocol returns the DiscoveryProtocol field if non-nil, zero value otherwise.

### GetDiscoveryProtocolOk

`func (o *OrchestrationManagedNamespaceSpec) GetDiscoveryProtocolOk() (*string, bool)`

GetDiscoveryProtocolOk returns a tuple with the DiscoveryProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryProtocol

`func (o *OrchestrationManagedNamespaceSpec) SetDiscoveryProtocol(v string)`

SetDiscoveryProtocol sets DiscoveryProtocol field to given value.

### HasDiscoveryProtocol

`func (o *OrchestrationManagedNamespaceSpec) HasDiscoveryProtocol() bool`

HasDiscoveryProtocol returns a boolean if a field has been set.

### GetMtu

`func (o *OrchestrationManagedNamespaceSpec) GetMtu() int64`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *OrchestrationManagedNamespaceSpec) GetMtuOk() (*int64, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *OrchestrationManagedNamespaceSpec) SetMtu(v int64)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *OrchestrationManagedNamespaceSpec) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetMulticastFilter

`func (o *OrchestrationManagedNamespaceSpec) GetMulticastFilter() string`

GetMulticastFilter returns the MulticastFilter field if non-nil, zero value otherwise.

### GetMulticastFilterOk

`func (o *OrchestrationManagedNamespaceSpec) GetMulticastFilterOk() (*string, bool)`

GetMulticastFilterOk returns a tuple with the MulticastFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMulticastFilter

`func (o *OrchestrationManagedNamespaceSpec) SetMulticastFilter(v string)`

SetMulticastFilter sets MulticastFilter field to given value.

### HasMulticastFilter

`func (o *OrchestrationManagedNamespaceSpec) HasMulticastFilter() bool`

HasMulticastFilter returns a boolean if a field has been set.

### GetOverrideVlanEnd

`func (o *OrchestrationManagedNamespaceSpec) GetOverrideVlanEnd() int64`

GetOverrideVlanEnd returns the OverrideVlanEnd field if non-nil, zero value otherwise.

### GetOverrideVlanEndOk

`func (o *OrchestrationManagedNamespaceSpec) GetOverrideVlanEndOk() (*int64, bool)`

GetOverrideVlanEndOk returns a tuple with the OverrideVlanEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideVlanEnd

`func (o *OrchestrationManagedNamespaceSpec) SetOverrideVlanEnd(v int64)`

SetOverrideVlanEnd sets OverrideVlanEnd field to given value.

### HasOverrideVlanEnd

`func (o *OrchestrationManagedNamespaceSpec) HasOverrideVlanEnd() bool`

HasOverrideVlanEnd returns a boolean if a field has been set.

### GetOverrideVlanStart

`func (o *OrchestrationManagedNamespaceSpec) GetOverrideVlanStart() int64`

GetOverrideVlanStart returns the OverrideVlanStart field if non-nil, zero value otherwise.

### GetOverrideVlanStartOk

`func (o *OrchestrationManagedNamespaceSpec) GetOverrideVlanStartOk() (*int64, bool)`

GetOverrideVlanStartOk returns a tuple with the OverrideVlanStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideVlanStart

`func (o *OrchestrationManagedNamespaceSpec) SetOverrideVlanStart(v int64)`

SetOverrideVlanStart sets OverrideVlanStart field to given value.

### HasOverrideVlanStart

`func (o *OrchestrationManagedNamespaceSpec) HasOverrideVlanStart() bool`

HasOverrideVlanStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


