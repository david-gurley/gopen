# NetworkRouteTable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**ApiObjectMeta**](apiObjectMeta.md) |  | [optional] 
**Spec** | Pointer to **map[string]interface{}** |  | [optional] 
**Status** | Pointer to [**NetworkRouteTableStatus**](networkRouteTableStatus.md) |  | [optional] 

## Methods

### NewNetworkRouteTable

`func NewNetworkRouteTable() *NetworkRouteTable`

NewNetworkRouteTable instantiates a new NetworkRouteTable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkRouteTableWithDefaults

`func NewNetworkRouteTableWithDefaults() *NetworkRouteTable`

NewNetworkRouteTableWithDefaults instantiates a new NetworkRouteTable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *NetworkRouteTable) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NetworkRouteTable) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NetworkRouteTable) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *NetworkRouteTable) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *NetworkRouteTable) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkRouteTable) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkRouteTable) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *NetworkRouteTable) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMeta

`func (o *NetworkRouteTable) GetMeta() ApiObjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *NetworkRouteTable) GetMetaOk() (*ApiObjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *NetworkRouteTable) SetMeta(v ApiObjectMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *NetworkRouteTable) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSpec

`func (o *NetworkRouteTable) GetSpec() map[string]interface{}`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *NetworkRouteTable) GetSpecOk() (*map[string]interface{}, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *NetworkRouteTable) SetSpec(v map[string]interface{})`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *NetworkRouteTable) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *NetworkRouteTable) GetStatus() NetworkRouteTableStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkRouteTable) GetStatusOk() (*NetworkRouteTableStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkRouteTable) SetStatus(v NetworkRouteTableStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NetworkRouteTable) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

