# NetworkAddStaticBindingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Ipv4StaticBindings** | Pointer to [**[]NetworkIPAMBinding**](NetworkIPAMBinding.md) |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**ApiObjectMeta**](apiObjectMeta.md) |  | [optional] 

## Methods

### NewNetworkAddStaticBindingsRequest

`func NewNetworkAddStaticBindingsRequest() *NetworkAddStaticBindingsRequest`

NewNetworkAddStaticBindingsRequest instantiates a new NetworkAddStaticBindingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkAddStaticBindingsRequestWithDefaults

`func NewNetworkAddStaticBindingsRequestWithDefaults() *NetworkAddStaticBindingsRequest`

NewNetworkAddStaticBindingsRequestWithDefaults instantiates a new NetworkAddStaticBindingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *NetworkAddStaticBindingsRequest) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NetworkAddStaticBindingsRequest) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NetworkAddStaticBindingsRequest) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *NetworkAddStaticBindingsRequest) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetIpv4StaticBindings

`func (o *NetworkAddStaticBindingsRequest) GetIpv4StaticBindings() []NetworkIPAMBinding`

GetIpv4StaticBindings returns the Ipv4StaticBindings field if non-nil, zero value otherwise.

### GetIpv4StaticBindingsOk

`func (o *NetworkAddStaticBindingsRequest) GetIpv4StaticBindingsOk() (*[]NetworkIPAMBinding, bool)`

GetIpv4StaticBindingsOk returns a tuple with the Ipv4StaticBindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4StaticBindings

`func (o *NetworkAddStaticBindingsRequest) SetIpv4StaticBindings(v []NetworkIPAMBinding)`

SetIpv4StaticBindings sets Ipv4StaticBindings field to given value.

### HasIpv4StaticBindings

`func (o *NetworkAddStaticBindingsRequest) HasIpv4StaticBindings() bool`

HasIpv4StaticBindings returns a boolean if a field has been set.

### GetKind

`func (o *NetworkAddStaticBindingsRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkAddStaticBindingsRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkAddStaticBindingsRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *NetworkAddStaticBindingsRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMeta

`func (o *NetworkAddStaticBindingsRequest) GetMeta() ApiObjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *NetworkAddStaticBindingsRequest) GetMetaOk() (*ApiObjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *NetworkAddStaticBindingsRequest) SetMeta(v ApiObjectMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *NetworkAddStaticBindingsRequest) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


