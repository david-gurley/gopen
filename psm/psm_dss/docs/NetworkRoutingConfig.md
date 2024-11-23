# NetworkRoutingConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**ApiObjectMeta**](apiObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**NetworkRoutingConfigSpec**](networkRoutingConfigSpec.md) |  | [optional] 
**Status** | Pointer to [**NetworkRoutingConfigStatus**](networkRoutingConfigStatus.md) |  | [optional] 

## Methods

### NewNetworkRoutingConfig

`func NewNetworkRoutingConfig() *NetworkRoutingConfig`

NewNetworkRoutingConfig instantiates a new NetworkRoutingConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkRoutingConfigWithDefaults

`func NewNetworkRoutingConfigWithDefaults() *NetworkRoutingConfig`

NewNetworkRoutingConfigWithDefaults instantiates a new NetworkRoutingConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *NetworkRoutingConfig) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NetworkRoutingConfig) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NetworkRoutingConfig) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *NetworkRoutingConfig) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *NetworkRoutingConfig) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkRoutingConfig) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkRoutingConfig) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *NetworkRoutingConfig) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMeta

`func (o *NetworkRoutingConfig) GetMeta() ApiObjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *NetworkRoutingConfig) GetMetaOk() (*ApiObjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *NetworkRoutingConfig) SetMeta(v ApiObjectMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *NetworkRoutingConfig) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSpec

`func (o *NetworkRoutingConfig) GetSpec() NetworkRoutingConfigSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *NetworkRoutingConfig) GetSpecOk() (*NetworkRoutingConfigSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *NetworkRoutingConfig) SetSpec(v NetworkRoutingConfigSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *NetworkRoutingConfig) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *NetworkRoutingConfig) GetStatus() NetworkRoutingConfigStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkRoutingConfig) GetStatusOk() (*NetworkRoutingConfigStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkRoutingConfig) SetStatus(v NetworkRoutingConfigStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NetworkRoutingConfig) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


