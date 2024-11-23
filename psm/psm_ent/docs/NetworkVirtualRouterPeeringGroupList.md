# NetworkVirtualRouterPeeringGroupList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Items** | Pointer to [**[]NetworkVirtualRouterPeeringGroup**](NetworkVirtualRouterPeeringGroup.md) | List of VirtualRouterPeeringGroup objects. | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**ListMeta** | Pointer to [**ApiListMeta**](apiListMeta.md) |  | [optional] 

## Methods

### NewNetworkVirtualRouterPeeringGroupList

`func NewNetworkVirtualRouterPeeringGroupList() *NetworkVirtualRouterPeeringGroupList`

NewNetworkVirtualRouterPeeringGroupList instantiates a new NetworkVirtualRouterPeeringGroupList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkVirtualRouterPeeringGroupListWithDefaults

`func NewNetworkVirtualRouterPeeringGroupListWithDefaults() *NetworkVirtualRouterPeeringGroupList`

NewNetworkVirtualRouterPeeringGroupListWithDefaults instantiates a new NetworkVirtualRouterPeeringGroupList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *NetworkVirtualRouterPeeringGroupList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NetworkVirtualRouterPeeringGroupList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NetworkVirtualRouterPeeringGroupList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *NetworkVirtualRouterPeeringGroupList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *NetworkVirtualRouterPeeringGroupList) GetItems() []NetworkVirtualRouterPeeringGroup`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *NetworkVirtualRouterPeeringGroupList) GetItemsOk() (*[]NetworkVirtualRouterPeeringGroup, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *NetworkVirtualRouterPeeringGroupList) SetItems(v []NetworkVirtualRouterPeeringGroup)`

SetItems sets Items field to given value.

### HasItems

`func (o *NetworkVirtualRouterPeeringGroupList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetKind

`func (o *NetworkVirtualRouterPeeringGroupList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkVirtualRouterPeeringGroupList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkVirtualRouterPeeringGroupList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *NetworkVirtualRouterPeeringGroupList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetListMeta

`func (o *NetworkVirtualRouterPeeringGroupList) GetListMeta() ApiListMeta`

GetListMeta returns the ListMeta field if non-nil, zero value otherwise.

### GetListMetaOk

`func (o *NetworkVirtualRouterPeeringGroupList) GetListMetaOk() (*ApiListMeta, bool)`

GetListMetaOk returns a tuple with the ListMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListMeta

`func (o *NetworkVirtualRouterPeeringGroupList) SetListMeta(v ApiListMeta)`

SetListMeta sets ListMeta field to given value.

### HasListMeta

`func (o *NetworkVirtualRouterPeeringGroupList) HasListMeta() bool`

HasListMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


