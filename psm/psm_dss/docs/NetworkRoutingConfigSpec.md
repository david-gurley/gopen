# NetworkRoutingConfigSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BgpConfig** | Pointer to [**NetworkBGPConfig**](networkBGPConfig.md) |  | [optional] 

## Methods

### NewNetworkRoutingConfigSpec

`func NewNetworkRoutingConfigSpec() *NetworkRoutingConfigSpec`

NewNetworkRoutingConfigSpec instantiates a new NetworkRoutingConfigSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkRoutingConfigSpecWithDefaults

`func NewNetworkRoutingConfigSpecWithDefaults() *NetworkRoutingConfigSpec`

NewNetworkRoutingConfigSpecWithDefaults instantiates a new NetworkRoutingConfigSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBgpConfig

`func (o *NetworkRoutingConfigSpec) GetBgpConfig() NetworkBGPConfig`

GetBgpConfig returns the BgpConfig field if non-nil, zero value otherwise.

### GetBgpConfigOk

`func (o *NetworkRoutingConfigSpec) GetBgpConfigOk() (*NetworkBGPConfig, bool)`

GetBgpConfigOk returns a tuple with the BgpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpConfig

`func (o *NetworkRoutingConfigSpec) SetBgpConfig(v NetworkBGPConfig)`

SetBgpConfig sets BgpConfig field to given value.

### HasBgpConfig

`func (o *NetworkRoutingConfigSpec) HasBgpConfig() bool`

HasBgpConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


