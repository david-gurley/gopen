# NetworkVirtualRouter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**ApiObjectMeta**](apiObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**NetworkVirtualRouterSpec**](networkVirtualRouterSpec.md) |  | [optional] 
**Status** | Pointer to [**NetworkVirtualRouterStatus**](networkVirtualRouterStatus.md) |  | [optional] 

## Methods

### NewNetworkVirtualRouter

`func NewNetworkVirtualRouter() *NetworkVirtualRouter`

NewNetworkVirtualRouter instantiates a new NetworkVirtualRouter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkVirtualRouterWithDefaults

`func NewNetworkVirtualRouterWithDefaults() *NetworkVirtualRouter`

NewNetworkVirtualRouterWithDefaults instantiates a new NetworkVirtualRouter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *NetworkVirtualRouter) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NetworkVirtualRouter) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NetworkVirtualRouter) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *NetworkVirtualRouter) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *NetworkVirtualRouter) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkVirtualRouter) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkVirtualRouter) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *NetworkVirtualRouter) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMeta

`func (o *NetworkVirtualRouter) GetMeta() ApiObjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *NetworkVirtualRouter) GetMetaOk() (*ApiObjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *NetworkVirtualRouter) SetMeta(v ApiObjectMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *NetworkVirtualRouter) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSpec

`func (o *NetworkVirtualRouter) GetSpec() NetworkVirtualRouterSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *NetworkVirtualRouter) GetSpecOk() (*NetworkVirtualRouterSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *NetworkVirtualRouter) SetSpec(v NetworkVirtualRouterSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *NetworkVirtualRouter) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *NetworkVirtualRouter) GetStatus() NetworkVirtualRouterStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkVirtualRouter) GetStatusOk() (*NetworkVirtualRouterStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkVirtualRouter) SetStatus(v NetworkVirtualRouterStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NetworkVirtualRouter) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


