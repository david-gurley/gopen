# NetworkVirtualRouterPeeringGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**ApiObjectMeta**](apiObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**NetworkVirtualRouterPeeringGroupSpec**](networkVirtualRouterPeeringGroupSpec.md) |  | [optional] 
**Status** | Pointer to [**NetworkVirtualRouterPeeringGroupStatus**](networkVirtualRouterPeeringGroupStatus.md) |  | [optional] 

## Methods

### NewNetworkVirtualRouterPeeringGroup

`func NewNetworkVirtualRouterPeeringGroup() *NetworkVirtualRouterPeeringGroup`

NewNetworkVirtualRouterPeeringGroup instantiates a new NetworkVirtualRouterPeeringGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkVirtualRouterPeeringGroupWithDefaults

`func NewNetworkVirtualRouterPeeringGroupWithDefaults() *NetworkVirtualRouterPeeringGroup`

NewNetworkVirtualRouterPeeringGroupWithDefaults instantiates a new NetworkVirtualRouterPeeringGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *NetworkVirtualRouterPeeringGroup) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NetworkVirtualRouterPeeringGroup) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NetworkVirtualRouterPeeringGroup) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *NetworkVirtualRouterPeeringGroup) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *NetworkVirtualRouterPeeringGroup) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkVirtualRouterPeeringGroup) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkVirtualRouterPeeringGroup) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *NetworkVirtualRouterPeeringGroup) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMeta

`func (o *NetworkVirtualRouterPeeringGroup) GetMeta() ApiObjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *NetworkVirtualRouterPeeringGroup) GetMetaOk() (*ApiObjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *NetworkVirtualRouterPeeringGroup) SetMeta(v ApiObjectMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *NetworkVirtualRouterPeeringGroup) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSpec

`func (o *NetworkVirtualRouterPeeringGroup) GetSpec() NetworkVirtualRouterPeeringGroupSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *NetworkVirtualRouterPeeringGroup) GetSpecOk() (*NetworkVirtualRouterPeeringGroupSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *NetworkVirtualRouterPeeringGroup) SetSpec(v NetworkVirtualRouterPeeringGroupSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *NetworkVirtualRouterPeeringGroup) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *NetworkVirtualRouterPeeringGroup) GetStatus() NetworkVirtualRouterPeeringGroupStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkVirtualRouterPeeringGroup) GetStatusOk() (*NetworkVirtualRouterPeeringGroupStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkVirtualRouterPeeringGroup) SetStatus(v NetworkVirtualRouterPeeringGroupStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NetworkVirtualRouterPeeringGroup) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


