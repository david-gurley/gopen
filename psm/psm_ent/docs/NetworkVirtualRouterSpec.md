# NetworkVirtualRouterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultIpamPolicy** | Pointer to **string** | Default IPAM policy for networks belonging to this Virtual Router. Any IPAM Policy specified in the Network overrides this. | [optional] 
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


