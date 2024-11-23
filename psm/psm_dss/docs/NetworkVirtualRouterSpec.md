# NetworkVirtualRouterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultIpamPolicy** | Pointer to **string** | Default IPAM policy for networks belonging to this Virtual Router. Any IPAM Policy specified in the Network overrides this. | [optional] 
**EgressSecurityPolicy** | Pointer to **[]string** | Security Policy to apply in the egress direction. | [optional] 
**FlowExportPolicy** | Pointer to **[]string** | FlowExportPolicy is the flow export policy associated to this virtual router. | [optional] 
**IngressSecurityPolicy** | Pointer to **[]string** | Security Policy to apply in the ingress direction. | [optional] 
**MaximumCpsPerNetworkPerDistributedServicesEntity** | Pointer to **int32** | Maximum Connections Per Second supported for any Network belonging to the Virtual Router within a Distributed Services Entity. The value configured here is the CPS limit enforced per Network within a Distributed Services Entity and is the same for all Networks within the Virtual Router. Value 0 means the CPS limit is not enforced and the CPS is limited only by the system capacity. Connections exceeding the CPS limit are dropped. Value should be between 0 and 409599. | [optional] 
**MaximumSessionsPerNetworkPerDistributedServicesEntity** | Pointer to **int32** | Maximum sessions supported in any Network belonging to the Virtual Router within a Distributed Services Entity. The value configured here is the sessions limit enforced per Network within a Distributed Services Entity and is the same for all Networks within the Virtual Router. Value 0 means the sessions limit is not enforced and the number of sessions is limited only by the system capacity. Sessions exceeding the sessions limit are dropped. Value should be between 0 and 16777215. | [optional] 
**RouteImportExport** | Pointer to [**NetworkRDSpec**](networkRDSpec.md) |  | [optional] 
**RouterMacAddress** | Pointer to **string** | Default Router MAC Address to use for this Virtual Router. Should be a valid MAC address. | [optional] 
**Type** | Pointer to **string** |  | [optional] [default to "unknown"]
**VxlanVni** | Pointer to **int64** | VxlAN VNI for the Virtual Router. Value should be between 0 and 16777215. | [optional] 

## Methods

### NewNetworkVirtualRouterSpec

`func NewNetworkVirtualRouterSpec() *NetworkVirtualRouterSpec`

NewNetworkVirtualRouterSpec instantiates a new NetworkVirtualRouterSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkVirtualRouterSpecWithDefaults

`func NewNetworkVirtualRouterSpecWithDefaults() *NetworkVirtualRouterSpec`

NewNetworkVirtualRouterSpecWithDefaults instantiates a new NetworkVirtualRouterSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultIpamPolicy

`func (o *NetworkVirtualRouterSpec) GetDefaultIpamPolicy() string`

GetDefaultIpamPolicy returns the DefaultIpamPolicy field if non-nil, zero value otherwise.

### GetDefaultIpamPolicyOk

`func (o *NetworkVirtualRouterSpec) GetDefaultIpamPolicyOk() (*string, bool)`

GetDefaultIpamPolicyOk returns a tuple with the DefaultIpamPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultIpamPolicy

`func (o *NetworkVirtualRouterSpec) SetDefaultIpamPolicy(v string)`

SetDefaultIpamPolicy sets DefaultIpamPolicy field to given value.

### HasDefaultIpamPolicy

`func (o *NetworkVirtualRouterSpec) HasDefaultIpamPolicy() bool`

HasDefaultIpamPolicy returns a boolean if a field has been set.

### GetEgressSecurityPolicy

`func (o *NetworkVirtualRouterSpec) GetEgressSecurityPolicy() []string`

GetEgressSecurityPolicy returns the EgressSecurityPolicy field if non-nil, zero value otherwise.

### GetEgressSecurityPolicyOk

`func (o *NetworkVirtualRouterSpec) GetEgressSecurityPolicyOk() (*[]string, bool)`

GetEgressSecurityPolicyOk returns a tuple with the EgressSecurityPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgressSecurityPolicy

`func (o *NetworkVirtualRouterSpec) SetEgressSecurityPolicy(v []string)`

SetEgressSecurityPolicy sets EgressSecurityPolicy field to given value.

### HasEgressSecurityPolicy

`func (o *NetworkVirtualRouterSpec) HasEgressSecurityPolicy() bool`

HasEgressSecurityPolicy returns a boolean if a field has been set.

### GetFlowExportPolicy

`func (o *NetworkVirtualRouterSpec) GetFlowExportPolicy() []string`

GetFlowExportPolicy returns the FlowExportPolicy field if non-nil, zero value otherwise.

### GetFlowExportPolicyOk

`func (o *NetworkVirtualRouterSpec) GetFlowExportPolicyOk() (*[]string, bool)`

GetFlowExportPolicyOk returns a tuple with the FlowExportPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowExportPolicy

`func (o *NetworkVirtualRouterSpec) SetFlowExportPolicy(v []string)`

SetFlowExportPolicy sets FlowExportPolicy field to given value.

### HasFlowExportPolicy

`func (o *NetworkVirtualRouterSpec) HasFlowExportPolicy() bool`

HasFlowExportPolicy returns a boolean if a field has been set.

### GetIngressSecurityPolicy

`func (o *NetworkVirtualRouterSpec) GetIngressSecurityPolicy() []string`

GetIngressSecurityPolicy returns the IngressSecurityPolicy field if non-nil, zero value otherwise.

### GetIngressSecurityPolicyOk

`func (o *NetworkVirtualRouterSpec) GetIngressSecurityPolicyOk() (*[]string, bool)`

GetIngressSecurityPolicyOk returns a tuple with the IngressSecurityPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressSecurityPolicy

`func (o *NetworkVirtualRouterSpec) SetIngressSecurityPolicy(v []string)`

SetIngressSecurityPolicy sets IngressSecurityPolicy field to given value.

### HasIngressSecurityPolicy

`func (o *NetworkVirtualRouterSpec) HasIngressSecurityPolicy() bool`

HasIngressSecurityPolicy returns a boolean if a field has been set.

### GetMaximumCpsPerNetworkPerDistributedServicesEntity

`func (o *NetworkVirtualRouterSpec) GetMaximumCpsPerNetworkPerDistributedServicesEntity() int32`

GetMaximumCpsPerNetworkPerDistributedServicesEntity returns the MaximumCpsPerNetworkPerDistributedServicesEntity field if non-nil, zero value otherwise.

### GetMaximumCpsPerNetworkPerDistributedServicesEntityOk

`func (o *NetworkVirtualRouterSpec) GetMaximumCpsPerNetworkPerDistributedServicesEntityOk() (*int32, bool)`

GetMaximumCpsPerNetworkPerDistributedServicesEntityOk returns a tuple with the MaximumCpsPerNetworkPerDistributedServicesEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumCpsPerNetworkPerDistributedServicesEntity

`func (o *NetworkVirtualRouterSpec) SetMaximumCpsPerNetworkPerDistributedServicesEntity(v int32)`

SetMaximumCpsPerNetworkPerDistributedServicesEntity sets MaximumCpsPerNetworkPerDistributedServicesEntity field to given value.

### HasMaximumCpsPerNetworkPerDistributedServicesEntity

`func (o *NetworkVirtualRouterSpec) HasMaximumCpsPerNetworkPerDistributedServicesEntity() bool`

HasMaximumCpsPerNetworkPerDistributedServicesEntity returns a boolean if a field has been set.

### GetMaximumSessionsPerNetworkPerDistributedServicesEntity

`func (o *NetworkVirtualRouterSpec) GetMaximumSessionsPerNetworkPerDistributedServicesEntity() int32`

GetMaximumSessionsPerNetworkPerDistributedServicesEntity returns the MaximumSessionsPerNetworkPerDistributedServicesEntity field if non-nil, zero value otherwise.

### GetMaximumSessionsPerNetworkPerDistributedServicesEntityOk

`func (o *NetworkVirtualRouterSpec) GetMaximumSessionsPerNetworkPerDistributedServicesEntityOk() (*int32, bool)`

GetMaximumSessionsPerNetworkPerDistributedServicesEntityOk returns a tuple with the MaximumSessionsPerNetworkPerDistributedServicesEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumSessionsPerNetworkPerDistributedServicesEntity

`func (o *NetworkVirtualRouterSpec) SetMaximumSessionsPerNetworkPerDistributedServicesEntity(v int32)`

SetMaximumSessionsPerNetworkPerDistributedServicesEntity sets MaximumSessionsPerNetworkPerDistributedServicesEntity field to given value.

### HasMaximumSessionsPerNetworkPerDistributedServicesEntity

`func (o *NetworkVirtualRouterSpec) HasMaximumSessionsPerNetworkPerDistributedServicesEntity() bool`

HasMaximumSessionsPerNetworkPerDistributedServicesEntity returns a boolean if a field has been set.

### GetRouteImportExport

`func (o *NetworkVirtualRouterSpec) GetRouteImportExport() NetworkRDSpec`

GetRouteImportExport returns the RouteImportExport field if non-nil, zero value otherwise.

### GetRouteImportExportOk

`func (o *NetworkVirtualRouterSpec) GetRouteImportExportOk() (*NetworkRDSpec, bool)`

GetRouteImportExportOk returns a tuple with the RouteImportExport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteImportExport

`func (o *NetworkVirtualRouterSpec) SetRouteImportExport(v NetworkRDSpec)`

SetRouteImportExport sets RouteImportExport field to given value.

### HasRouteImportExport

`func (o *NetworkVirtualRouterSpec) HasRouteImportExport() bool`

HasRouteImportExport returns a boolean if a field has been set.

### GetRouterMacAddress

`func (o *NetworkVirtualRouterSpec) GetRouterMacAddress() string`

GetRouterMacAddress returns the RouterMacAddress field if non-nil, zero value otherwise.

### GetRouterMacAddressOk

`func (o *NetworkVirtualRouterSpec) GetRouterMacAddressOk() (*string, bool)`

GetRouterMacAddressOk returns a tuple with the RouterMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterMacAddress

`func (o *NetworkVirtualRouterSpec) SetRouterMacAddress(v string)`

SetRouterMacAddress sets RouterMacAddress field to given value.

### HasRouterMacAddress

`func (o *NetworkVirtualRouterSpec) HasRouterMacAddress() bool`

HasRouterMacAddress returns a boolean if a field has been set.

### GetType

`func (o *NetworkVirtualRouterSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkVirtualRouterSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkVirtualRouterSpec) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NetworkVirtualRouterSpec) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVxlanVni

`func (o *NetworkVirtualRouterSpec) GetVxlanVni() int64`

GetVxlanVni returns the VxlanVni field if non-nil, zero value otherwise.

### GetVxlanVniOk

`func (o *NetworkVirtualRouterSpec) GetVxlanVniOk() (*int64, bool)`

GetVxlanVniOk returns a tuple with the VxlanVni field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVxlanVni

`func (o *NetworkVirtualRouterSpec) SetVxlanVni(v int64)`

SetVxlanVni sets VxlanVni field to given value.

### HasVxlanVni

`func (o *NetworkVirtualRouterSpec) HasVxlanVni() bool`

HasVxlanVni returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


