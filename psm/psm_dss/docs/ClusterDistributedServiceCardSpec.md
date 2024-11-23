# ClusterDistributedServiceCardSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Admit** | Pointer to **bool** | Admit allows a DistributedServiceCard to join the cluster. | [optional] 
**Controllers** | Pointer to **[]string** | Controllers contains the list of remote controllers IP addresses or hostnames. | [optional] 
**Dscprofile** | Pointer to **string** |  | [optional] 
**EnableSecureBoot** | Pointer to **bool** | EnableSecureBoot a true value indicates, set lifecycle fuse to enable secure boot. | [optional] [default to false]
**FlowExportPolicy** | Pointer to [**[]ClusterFlowExportPolicyRef**](ClusterFlowExportPolicyRef.md) | FlowExportPolicy is the configuration for flow export policy. | [optional] 
**FwlogPolicy** | Pointer to [**ClusterFwlogPolicyRef**](clusterFwlogPolicyRef.md) |  | [optional] 
**Id** | Pointer to **string** | ID is used as a user friendly identifier in logs/events. | [optional] 
**IpConfig** | Pointer to [**ClusterIPConfig**](clusterIPConfig.md) |  | [optional] 
**MgmtMode** | Pointer to **string** | MgmtMode defines the management mode of the DistributedServiceCard. | [optional] [default to "host"]
**MgmtVlan** | Pointer to **int64** | MgmtVlan defines the vlan to be used in network managed mode. The default of 0 means we use untagged-vlan for doing inband management. Value should be between 0 and 4095. | [optional] 
**NetworkMode** | Pointer to **string** | MgmtMode defines the management mode of the DistributedServiceCard. | [optional] [default to "oob"]
**Policer** | Pointer to [**ClusterPolicerRef**](clusterPolicerRef.md) |  | [optional] 
**RoutingConfig** | Pointer to **string** | RoutingConfig is the routing configuration for the underlay routed network that this DSC participates in. | [optional] 

## Methods

### NewClusterDistributedServiceCardSpec

`func NewClusterDistributedServiceCardSpec() *ClusterDistributedServiceCardSpec`

NewClusterDistributedServiceCardSpec instantiates a new ClusterDistributedServiceCardSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterDistributedServiceCardSpecWithDefaults

`func NewClusterDistributedServiceCardSpecWithDefaults() *ClusterDistributedServiceCardSpec`

NewClusterDistributedServiceCardSpecWithDefaults instantiates a new ClusterDistributedServiceCardSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdmit

`func (o *ClusterDistributedServiceCardSpec) GetAdmit() bool`

GetAdmit returns the Admit field if non-nil, zero value otherwise.

### GetAdmitOk

`func (o *ClusterDistributedServiceCardSpec) GetAdmitOk() (*bool, bool)`

GetAdmitOk returns a tuple with the Admit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmit

`func (o *ClusterDistributedServiceCardSpec) SetAdmit(v bool)`

SetAdmit sets Admit field to given value.

### HasAdmit

`func (o *ClusterDistributedServiceCardSpec) HasAdmit() bool`

HasAdmit returns a boolean if a field has been set.

### GetControllers

`func (o *ClusterDistributedServiceCardSpec) GetControllers() []string`

GetControllers returns the Controllers field if non-nil, zero value otherwise.

### GetControllersOk

`func (o *ClusterDistributedServiceCardSpec) GetControllersOk() (*[]string, bool)`

GetControllersOk returns a tuple with the Controllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllers

`func (o *ClusterDistributedServiceCardSpec) SetControllers(v []string)`

SetControllers sets Controllers field to given value.

### HasControllers

`func (o *ClusterDistributedServiceCardSpec) HasControllers() bool`

HasControllers returns a boolean if a field has been set.

### GetDscprofile

`func (o *ClusterDistributedServiceCardSpec) GetDscprofile() string`

GetDscprofile returns the Dscprofile field if non-nil, zero value otherwise.

### GetDscprofileOk

`func (o *ClusterDistributedServiceCardSpec) GetDscprofileOk() (*string, bool)`

GetDscprofileOk returns a tuple with the Dscprofile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDscprofile

`func (o *ClusterDistributedServiceCardSpec) SetDscprofile(v string)`

SetDscprofile sets Dscprofile field to given value.

### HasDscprofile

`func (o *ClusterDistributedServiceCardSpec) HasDscprofile() bool`

HasDscprofile returns a boolean if a field has been set.

### GetEnableSecureBoot

`func (o *ClusterDistributedServiceCardSpec) GetEnableSecureBoot() bool`

GetEnableSecureBoot returns the EnableSecureBoot field if non-nil, zero value otherwise.

### GetEnableSecureBootOk

`func (o *ClusterDistributedServiceCardSpec) GetEnableSecureBootOk() (*bool, bool)`

GetEnableSecureBootOk returns a tuple with the EnableSecureBoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSecureBoot

`func (o *ClusterDistributedServiceCardSpec) SetEnableSecureBoot(v bool)`

SetEnableSecureBoot sets EnableSecureBoot field to given value.

### HasEnableSecureBoot

`func (o *ClusterDistributedServiceCardSpec) HasEnableSecureBoot() bool`

HasEnableSecureBoot returns a boolean if a field has been set.

### GetFlowExportPolicy

`func (o *ClusterDistributedServiceCardSpec) GetFlowExportPolicy() []ClusterFlowExportPolicyRef`

GetFlowExportPolicy returns the FlowExportPolicy field if non-nil, zero value otherwise.

### GetFlowExportPolicyOk

`func (o *ClusterDistributedServiceCardSpec) GetFlowExportPolicyOk() (*[]ClusterFlowExportPolicyRef, bool)`

GetFlowExportPolicyOk returns a tuple with the FlowExportPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowExportPolicy

`func (o *ClusterDistributedServiceCardSpec) SetFlowExportPolicy(v []ClusterFlowExportPolicyRef)`

SetFlowExportPolicy sets FlowExportPolicy field to given value.

### HasFlowExportPolicy

`func (o *ClusterDistributedServiceCardSpec) HasFlowExportPolicy() bool`

HasFlowExportPolicy returns a boolean if a field has been set.

### GetFwlogPolicy

`func (o *ClusterDistributedServiceCardSpec) GetFwlogPolicy() ClusterFwlogPolicyRef`

GetFwlogPolicy returns the FwlogPolicy field if non-nil, zero value otherwise.

### GetFwlogPolicyOk

`func (o *ClusterDistributedServiceCardSpec) GetFwlogPolicyOk() (*ClusterFwlogPolicyRef, bool)`

GetFwlogPolicyOk returns a tuple with the FwlogPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFwlogPolicy

`func (o *ClusterDistributedServiceCardSpec) SetFwlogPolicy(v ClusterFwlogPolicyRef)`

SetFwlogPolicy sets FwlogPolicy field to given value.

### HasFwlogPolicy

`func (o *ClusterDistributedServiceCardSpec) HasFwlogPolicy() bool`

HasFwlogPolicy returns a boolean if a field has been set.

### GetId

`func (o *ClusterDistributedServiceCardSpec) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterDistributedServiceCardSpec) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterDistributedServiceCardSpec) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ClusterDistributedServiceCardSpec) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpConfig

`func (o *ClusterDistributedServiceCardSpec) GetIpConfig() ClusterIPConfig`

GetIpConfig returns the IpConfig field if non-nil, zero value otherwise.

### GetIpConfigOk

`func (o *ClusterDistributedServiceCardSpec) GetIpConfigOk() (*ClusterIPConfig, bool)`

GetIpConfigOk returns a tuple with the IpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpConfig

`func (o *ClusterDistributedServiceCardSpec) SetIpConfig(v ClusterIPConfig)`

SetIpConfig sets IpConfig field to given value.

### HasIpConfig

`func (o *ClusterDistributedServiceCardSpec) HasIpConfig() bool`

HasIpConfig returns a boolean if a field has been set.

### GetMgmtMode

`func (o *ClusterDistributedServiceCardSpec) GetMgmtMode() string`

GetMgmtMode returns the MgmtMode field if non-nil, zero value otherwise.

### GetMgmtModeOk

`func (o *ClusterDistributedServiceCardSpec) GetMgmtModeOk() (*string, bool)`

GetMgmtModeOk returns a tuple with the MgmtMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtMode

`func (o *ClusterDistributedServiceCardSpec) SetMgmtMode(v string)`

SetMgmtMode sets MgmtMode field to given value.

### HasMgmtMode

`func (o *ClusterDistributedServiceCardSpec) HasMgmtMode() bool`

HasMgmtMode returns a boolean if a field has been set.

### GetMgmtVlan

`func (o *ClusterDistributedServiceCardSpec) GetMgmtVlan() int64`

GetMgmtVlan returns the MgmtVlan field if non-nil, zero value otherwise.

### GetMgmtVlanOk

`func (o *ClusterDistributedServiceCardSpec) GetMgmtVlanOk() (*int64, bool)`

GetMgmtVlanOk returns a tuple with the MgmtVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtVlan

`func (o *ClusterDistributedServiceCardSpec) SetMgmtVlan(v int64)`

SetMgmtVlan sets MgmtVlan field to given value.

### HasMgmtVlan

`func (o *ClusterDistributedServiceCardSpec) HasMgmtVlan() bool`

HasMgmtVlan returns a boolean if a field has been set.

### GetNetworkMode

`func (o *ClusterDistributedServiceCardSpec) GetNetworkMode() string`

GetNetworkMode returns the NetworkMode field if non-nil, zero value otherwise.

### GetNetworkModeOk

`func (o *ClusterDistributedServiceCardSpec) GetNetworkModeOk() (*string, bool)`

GetNetworkModeOk returns a tuple with the NetworkMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkMode

`func (o *ClusterDistributedServiceCardSpec) SetNetworkMode(v string)`

SetNetworkMode sets NetworkMode field to given value.

### HasNetworkMode

`func (o *ClusterDistributedServiceCardSpec) HasNetworkMode() bool`

HasNetworkMode returns a boolean if a field has been set.

### GetPolicer

`func (o *ClusterDistributedServiceCardSpec) GetPolicer() ClusterPolicerRef`

GetPolicer returns the Policer field if non-nil, zero value otherwise.

### GetPolicerOk

`func (o *ClusterDistributedServiceCardSpec) GetPolicerOk() (*ClusterPolicerRef, bool)`

GetPolicerOk returns a tuple with the Policer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicer

`func (o *ClusterDistributedServiceCardSpec) SetPolicer(v ClusterPolicerRef)`

SetPolicer sets Policer field to given value.

### HasPolicer

`func (o *ClusterDistributedServiceCardSpec) HasPolicer() bool`

HasPolicer returns a boolean if a field has been set.

### GetRoutingConfig

`func (o *ClusterDistributedServiceCardSpec) GetRoutingConfig() string`

GetRoutingConfig returns the RoutingConfig field if non-nil, zero value otherwise.

### GetRoutingConfigOk

`func (o *ClusterDistributedServiceCardSpec) GetRoutingConfigOk() (*string, bool)`

GetRoutingConfigOk returns a tuple with the RoutingConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingConfig

`func (o *ClusterDistributedServiceCardSpec) SetRoutingConfig(v string)`

SetRoutingConfig sets RoutingConfig field to given value.

### HasRoutingConfig

`func (o *ClusterDistributedServiceCardSpec) HasRoutingConfig() bool`

HasRoutingConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


