# NetworkLbPolicySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | Pointer to **string** | load balancing algorithm. | [optional] 
**HealthCheck** | Pointer to [**NetworkHealthCheckSpec**](networkHealthCheckSpec.md) |  | [optional] 
**SessionAffinity** | Pointer to **string** | session affinity. | [optional] 
**Type** | Pointer to **string** | load balancing type. | [optional] 

## Methods

### NewNetworkLbPolicySpec

`func NewNetworkLbPolicySpec() *NetworkLbPolicySpec`

NewNetworkLbPolicySpec instantiates a new NetworkLbPolicySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkLbPolicySpecWithDefaults

`func NewNetworkLbPolicySpecWithDefaults() *NetworkLbPolicySpec`

NewNetworkLbPolicySpecWithDefaults instantiates a new NetworkLbPolicySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithm

`func (o *NetworkLbPolicySpec) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *NetworkLbPolicySpec) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *NetworkLbPolicySpec) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *NetworkLbPolicySpec) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetHealthCheck

`func (o *NetworkLbPolicySpec) GetHealthCheck() NetworkHealthCheckSpec`

GetHealthCheck returns the HealthCheck field if non-nil, zero value otherwise.

### GetHealthCheckOk

`func (o *NetworkLbPolicySpec) GetHealthCheckOk() (*NetworkHealthCheckSpec, bool)`

GetHealthCheckOk returns a tuple with the HealthCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheck

`func (o *NetworkLbPolicySpec) SetHealthCheck(v NetworkHealthCheckSpec)`

SetHealthCheck sets HealthCheck field to given value.

### HasHealthCheck

`func (o *NetworkLbPolicySpec) HasHealthCheck() bool`

HasHealthCheck returns a boolean if a field has been set.

### GetSessionAffinity

`func (o *NetworkLbPolicySpec) GetSessionAffinity() string`

GetSessionAffinity returns the SessionAffinity field if non-nil, zero value otherwise.

### GetSessionAffinityOk

`func (o *NetworkLbPolicySpec) GetSessionAffinityOk() (*string, bool)`

GetSessionAffinityOk returns a tuple with the SessionAffinity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionAffinity

`func (o *NetworkLbPolicySpec) SetSessionAffinity(v string)`

SetSessionAffinity sets SessionAffinity field to given value.

### HasSessionAffinity

`func (o *NetworkLbPolicySpec) HasSessionAffinity() bool`

HasSessionAffinity returns a boolean if a field has been set.

### GetType

`func (o *NetworkLbPolicySpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkLbPolicySpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkLbPolicySpec) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NetworkLbPolicySpec) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


