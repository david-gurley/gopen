# NetworkRoutingConfigStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthConfigStatus** | Pointer to [**[]NetworkBGPAuthStatus**](NetworkBGPAuthStatus.md) | Authentication config status. | [optional] 
**PropagationStatus** | Pointer to [**SecurityPropagationStatus**](securityPropagationStatus.md) |  | [optional] 

## Methods

### NewNetworkRoutingConfigStatus

`func NewNetworkRoutingConfigStatus() *NetworkRoutingConfigStatus`

NewNetworkRoutingConfigStatus instantiates a new NetworkRoutingConfigStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkRoutingConfigStatusWithDefaults

`func NewNetworkRoutingConfigStatusWithDefaults() *NetworkRoutingConfigStatus`

NewNetworkRoutingConfigStatusWithDefaults instantiates a new NetworkRoutingConfigStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthConfigStatus

`func (o *NetworkRoutingConfigStatus) GetAuthConfigStatus() []NetworkBGPAuthStatus`

GetAuthConfigStatus returns the AuthConfigStatus field if non-nil, zero value otherwise.

### GetAuthConfigStatusOk

`func (o *NetworkRoutingConfigStatus) GetAuthConfigStatusOk() (*[]NetworkBGPAuthStatus, bool)`

GetAuthConfigStatusOk returns a tuple with the AuthConfigStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthConfigStatus

`func (o *NetworkRoutingConfigStatus) SetAuthConfigStatus(v []NetworkBGPAuthStatus)`

SetAuthConfigStatus sets AuthConfigStatus field to given value.

### HasAuthConfigStatus

`func (o *NetworkRoutingConfigStatus) HasAuthConfigStatus() bool`

HasAuthConfigStatus returns a boolean if a field has been set.

### GetPropagationStatus

`func (o *NetworkRoutingConfigStatus) GetPropagationStatus() SecurityPropagationStatus`

GetPropagationStatus returns the PropagationStatus field if non-nil, zero value otherwise.

### GetPropagationStatusOk

`func (o *NetworkRoutingConfigStatus) GetPropagationStatusOk() (*SecurityPropagationStatus, bool)`

GetPropagationStatusOk returns a tuple with the PropagationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropagationStatus

`func (o *NetworkRoutingConfigStatus) SetPropagationStatus(v SecurityPropagationStatus)`

SetPropagationStatus sets PropagationStatus field to given value.

### HasPropagationStatus

`func (o *NetworkRoutingConfigStatus) HasPropagationStatus() bool`

HasPropagationStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


