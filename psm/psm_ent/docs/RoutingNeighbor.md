# RoutingNeighbor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**ApiObjectMeta**](apiObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**NetworkBGPNeighbor**](networkBGPNeighbor.md) |  | [optional] 
**Status** | Pointer to [**RoutingNeighborStatus**](routingNeighborStatus.md) |  | [optional] 

## Methods

### NewRoutingNeighbor

`func NewRoutingNeighbor() *RoutingNeighbor`

NewRoutingNeighbor instantiates a new RoutingNeighbor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingNeighborWithDefaults

`func NewRoutingNeighborWithDefaults() *RoutingNeighbor`

NewRoutingNeighborWithDefaults instantiates a new RoutingNeighbor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *RoutingNeighbor) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *RoutingNeighbor) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *RoutingNeighbor) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *RoutingNeighbor) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *RoutingNeighbor) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *RoutingNeighbor) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *RoutingNeighbor) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *RoutingNeighbor) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMeta

`func (o *RoutingNeighbor) GetMeta() ApiObjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *RoutingNeighbor) GetMetaOk() (*ApiObjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *RoutingNeighbor) SetMeta(v ApiObjectMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *RoutingNeighbor) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSpec

`func (o *RoutingNeighbor) GetSpec() NetworkBGPNeighbor`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *RoutingNeighbor) GetSpecOk() (*NetworkBGPNeighbor, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *RoutingNeighbor) SetSpec(v NetworkBGPNeighbor)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *RoutingNeighbor) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *RoutingNeighbor) GetStatus() RoutingNeighborStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RoutingNeighbor) GetStatusOk() (*RoutingNeighborStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RoutingNeighbor) SetStatus(v RoutingNeighborStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RoutingNeighbor) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


