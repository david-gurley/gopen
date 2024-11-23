# ClusterDistributedServiceEntitySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Admit** | Pointer to **bool** |  | [optional] 
**Controllers** | Pointer to **[]string** |  | [optional] 
**Dscprofile** | Pointer to **string** |  | [optional] 
**EnableSecureBoot** | Pointer to **bool** |  | [optional] [default to false]
**FlowExportPolicy** | Pointer to [**[]ClusterFlowExportPolicyRef**](ClusterFlowExportPolicyRef.md) |  | [optional] 
**FwlogPolicy** | Pointer to [**ClusterFwlogPolicyRef**](clusterFwlogPolicyRef.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IpConfig** | Pointer to [**ClusterIPConfig**](clusterIPConfig.md) |  | [optional] 
**MgmtMode** | Pointer to **string** |  | [optional] [default to "host"]
**MgmtVlan** | Pointer to **int64** | Value should be between 0 and 4095. | [optional] 
**NetworkMode** | Pointer to **string** |  | [optional] [default to "oob"]
**Policer** | Pointer to [**ClusterPolicerRef**](clusterPolicerRef.md) |  | [optional] 
**RoutingConfig** | Pointer to **string** |  | [optional] 

## Methods

### NewClusterDistributedServiceEntitySpec

`func NewClusterDistributedServiceEntitySpec() *ClusterDistributedServiceEntitySpec`

NewClusterDistributedServiceEntitySpec instantiates a new ClusterDistributedServiceEntitySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterDistributedServiceEntitySpecWithDefaults

`func NewClusterDistributedServiceEntitySpecWithDefaults() *ClusterDistributedServiceEntitySpec`

NewClusterDistributedServiceEntitySpecWithDefaults instantiates a new ClusterDistributedServiceEntitySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdmit

`func (o *ClusterDistributedServiceEntitySpec) GetAdmit() bool`

GetAdmit returns the Admit field if non-nil, zero value otherwise.

### GetAdmitOk

`func (o *ClusterDistributedServiceEntitySpec) GetAdmitOk() (*bool, bool)`

GetAdmitOk returns a tuple with the Admit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmit

`func (o *ClusterDistributedServiceEntitySpec) SetAdmit(v bool)`

SetAdmit sets Admit field to given value.

### HasAdmit

`func (o *ClusterDistributedServiceEntitySpec) HasAdmit() bool`

HasAdmit returns a boolean if a field has been set.

### GetControllers

`func (o *ClusterDistributedServiceEntitySpec) GetControllers() []string`

GetControllers returns the Controllers field if non-nil, zero value otherwise.

### GetControllersOk

`func (o *ClusterDistributedServiceEntitySpec) GetControllersOk() (*[]string, bool)`

GetControllersOk returns a tuple with the Controllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllers

`func (o *ClusterDistributedServiceEntitySpec) SetControllers(v []string)`

SetControllers sets Controllers field to given value.

### HasControllers

`func (o *ClusterDistributedServiceEntitySpec) HasControllers() bool`

HasControllers returns a boolean if a field has been set.

### GetDscprofile

`func (o *ClusterDistributedServiceEntitySpec) GetDscprofile() string`

GetDscprofile returns the Dscprofile field if non-nil, zero value otherwise.

### GetDscprofileOk

`func (o *ClusterDistributedServiceEntitySpec) GetDscprofileOk() (*string, bool)`

GetDscprofileOk returns a tuple with the Dscprofile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDscprofile

`func (o *ClusterDistributedServiceEntitySpec) SetDscprofile(v string)`

SetDscprofile sets Dscprofile field to given value.

### HasDscprofile

`func (o *ClusterDistributedServiceEntitySpec) HasDscprofile() bool`

HasDscprofile returns a boolean if a field has been set.

### GetEnableSecureBoot

`func (o *ClusterDistributedServiceEntitySpec) GetEnableSecureBoot() bool`

GetEnableSecureBoot returns the EnableSecureBoot field if non-nil, zero value otherwise.

### GetEnableSecureBootOk

`func (o *ClusterDistributedServiceEntitySpec) GetEnableSecureBootOk() (*bool, bool)`

GetEnableSecureBootOk returns a tuple with the EnableSecureBoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSecureBoot

`func (o *ClusterDistributedServiceEntitySpec) SetEnableSecureBoot(v bool)`

SetEnableSecureBoot sets EnableSecureBoot field to given value.

### HasEnableSecureBoot

`func (o *ClusterDistributedServiceEntitySpec) HasEnableSecureBoot() bool`

HasEnableSecureBoot returns a boolean if a field has been set.

### GetFlowExportPolicy

`func (o *ClusterDistributedServiceEntitySpec) GetFlowExportPolicy() []ClusterFlowExportPolicyRef`

GetFlowExportPolicy returns the FlowExportPolicy field if non-nil, zero value otherwise.

### GetFlowExportPolicyOk

`func (o *ClusterDistributedServiceEntitySpec) GetFlowExportPolicyOk() (*[]ClusterFlowExportPolicyRef, bool)`

GetFlowExportPolicyOk returns a tuple with the FlowExportPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowExportPolicy

`func (o *ClusterDistributedServiceEntitySpec) SetFlowExportPolicy(v []ClusterFlowExportPolicyRef)`

SetFlowExportPolicy sets FlowExportPolicy field to given value.

### HasFlowExportPolicy

`func (o *ClusterDistributedServiceEntitySpec) HasFlowExportPolicy() bool`

HasFlowExportPolicy returns a boolean if a field has been set.

### GetFwlogPolicy

`func (o *ClusterDistributedServiceEntitySpec) GetFwlogPolicy() ClusterFwlogPolicyRef`

GetFwlogPolicy returns the FwlogPolicy field if non-nil, zero value otherwise.

### GetFwlogPolicyOk

`func (o *ClusterDistributedServiceEntitySpec) GetFwlogPolicyOk() (*ClusterFwlogPolicyRef, bool)`

GetFwlogPolicyOk returns a tuple with the FwlogPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFwlogPolicy

`func (o *ClusterDistributedServiceEntitySpec) SetFwlogPolicy(v ClusterFwlogPolicyRef)`

SetFwlogPolicy sets FwlogPolicy field to given value.

### HasFwlogPolicy

`func (o *ClusterDistributedServiceEntitySpec) HasFwlogPolicy() bool`

HasFwlogPolicy returns a boolean if a field has been set.

### GetId

`func (o *ClusterDistributedServiceEntitySpec) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterDistributedServiceEntitySpec) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterDistributedServiceEntitySpec) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ClusterDistributedServiceEntitySpec) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpConfig

`func (o *ClusterDistributedServiceEntitySpec) GetIpConfig() ClusterIPConfig`

GetIpConfig returns the IpConfig field if non-nil, zero value otherwise.

### GetIpConfigOk

`func (o *ClusterDistributedServiceEntitySpec) GetIpConfigOk() (*ClusterIPConfig, bool)`

GetIpConfigOk returns a tuple with the IpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpConfig

`func (o *ClusterDistributedServiceEntitySpec) SetIpConfig(v ClusterIPConfig)`

SetIpConfig sets IpConfig field to given value.

### HasIpConfig

`func (o *ClusterDistributedServiceEntitySpec) HasIpConfig() bool`

HasIpConfig returns a boolean if a field has been set.

### GetMgmtMode

`func (o *ClusterDistributedServiceEntitySpec) GetMgmtMode() string`

GetMgmtMode returns the MgmtMode field if non-nil, zero value otherwise.

### GetMgmtModeOk

`func (o *ClusterDistributedServiceEntitySpec) GetMgmtModeOk() (*string, bool)`

GetMgmtModeOk returns a tuple with the MgmtMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtMode

`func (o *ClusterDistributedServiceEntitySpec) SetMgmtMode(v string)`

SetMgmtMode sets MgmtMode field to given value.

### HasMgmtMode

`func (o *ClusterDistributedServiceEntitySpec) HasMgmtMode() bool`

HasMgmtMode returns a boolean if a field has been set.

### GetMgmtVlan

`func (o *ClusterDistributedServiceEntitySpec) GetMgmtVlan() int64`

GetMgmtVlan returns the MgmtVlan field if non-nil, zero value otherwise.

### GetMgmtVlanOk

`func (o *ClusterDistributedServiceEntitySpec) GetMgmtVlanOk() (*int64, bool)`

GetMgmtVlanOk returns a tuple with the MgmtVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtVlan

`func (o *ClusterDistributedServiceEntitySpec) SetMgmtVlan(v int64)`

SetMgmtVlan sets MgmtVlan field to given value.

### HasMgmtVlan

`func (o *ClusterDistributedServiceEntitySpec) HasMgmtVlan() bool`

HasMgmtVlan returns a boolean if a field has been set.

### GetNetworkMode

`func (o *ClusterDistributedServiceEntitySpec) GetNetworkMode() string`

GetNetworkMode returns the NetworkMode field if non-nil, zero value otherwise.

### GetNetworkModeOk

`func (o *ClusterDistributedServiceEntitySpec) GetNetworkModeOk() (*string, bool)`

GetNetworkModeOk returns a tuple with the NetworkMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkMode

`func (o *ClusterDistributedServiceEntitySpec) SetNetworkMode(v string)`

SetNetworkMode sets NetworkMode field to given value.

### HasNetworkMode

`func (o *ClusterDistributedServiceEntitySpec) HasNetworkMode() bool`

HasNetworkMode returns a boolean if a field has been set.

### GetPolicer

`func (o *ClusterDistributedServiceEntitySpec) GetPolicer() ClusterPolicerRef`

GetPolicer returns the Policer field if non-nil, zero value otherwise.

### GetPolicerOk

`func (o *ClusterDistributedServiceEntitySpec) GetPolicerOk() (*ClusterPolicerRef, bool)`

GetPolicerOk returns a tuple with the Policer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicer

`func (o *ClusterDistributedServiceEntitySpec) SetPolicer(v ClusterPolicerRef)`

SetPolicer sets Policer field to given value.

### HasPolicer

`func (o *ClusterDistributedServiceEntitySpec) HasPolicer() bool`

HasPolicer returns a boolean if a field has been set.

### GetRoutingConfig

`func (o *ClusterDistributedServiceEntitySpec) GetRoutingConfig() string`

GetRoutingConfig returns the RoutingConfig field if non-nil, zero value otherwise.

### GetRoutingConfigOk

`func (o *ClusterDistributedServiceEntitySpec) GetRoutingConfigOk() (*string, bool)`

GetRoutingConfigOk returns a tuple with the RoutingConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingConfig

`func (o *ClusterDistributedServiceEntitySpec) SetRoutingConfig(v string)`

SetRoutingConfig sets RoutingConfig field to given value.

### HasRoutingConfig

`func (o *ClusterDistributedServiceEntitySpec) HasRoutingConfig() bool`

HasRoutingConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


