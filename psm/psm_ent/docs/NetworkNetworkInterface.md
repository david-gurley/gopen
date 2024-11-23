# NetworkNetworkInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**ApiObjectMeta**](apiObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**NetworkNetworkInterfaceSpec**](networkNetworkInterfaceSpec.md) |  | [optional] 
**Status** | Pointer to [**NetworkNetworkInterfaceStatus**](networkNetworkInterfaceStatus.md) |  | [optional] 

## Methods

### NewNetworkNetworkInterface

`func NewNetworkNetworkInterface() *NetworkNetworkInterface`

NewNetworkNetworkInterface instantiates a new NetworkNetworkInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkNetworkInterfaceWithDefaults

`func NewNetworkNetworkInterfaceWithDefaults() *NetworkNetworkInterface`

NewNetworkNetworkInterfaceWithDefaults instantiates a new NetworkNetworkInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *NetworkNetworkInterface) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NetworkNetworkInterface) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NetworkNetworkInterface) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *NetworkNetworkInterface) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *NetworkNetworkInterface) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkNetworkInterface) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkNetworkInterface) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *NetworkNetworkInterface) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMeta

`func (o *NetworkNetworkInterface) GetMeta() ApiObjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *NetworkNetworkInterface) GetMetaOk() (*ApiObjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *NetworkNetworkInterface) SetMeta(v ApiObjectMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *NetworkNetworkInterface) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSpec

`func (o *NetworkNetworkInterface) GetSpec() NetworkNetworkInterfaceSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *NetworkNetworkInterface) GetSpecOk() (*NetworkNetworkInterfaceSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *NetworkNetworkInterface) SetSpec(v NetworkNetworkInterfaceSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *NetworkNetworkInterface) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *NetworkNetworkInterface) GetStatus() NetworkNetworkInterfaceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkNetworkInterface) GetStatusOk() (*NetworkNetworkInterfaceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkNetworkInterface) SetStatus(v NetworkNetworkInterfaceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NetworkNetworkInterface) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


