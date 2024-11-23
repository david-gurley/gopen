# NetworkSecurityPolicyStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EpochId** | Pointer to **int64** | Epoch ID which is monotonically incremented for each policy update. | [optional] 
**PolicyReferences** | Pointer to [**[]NetworkPolicyReference**](NetworkPolicyReference.md) | List of policy refernces for this network. | [optional] 

## Methods

### NewNetworkSecurityPolicyStatus

`func NewNetworkSecurityPolicyStatus() *NetworkSecurityPolicyStatus`

NewNetworkSecurityPolicyStatus instantiates a new NetworkSecurityPolicyStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkSecurityPolicyStatusWithDefaults

`func NewNetworkSecurityPolicyStatusWithDefaults() *NetworkSecurityPolicyStatus`

NewNetworkSecurityPolicyStatusWithDefaults instantiates a new NetworkSecurityPolicyStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEpochId

`func (o *NetworkSecurityPolicyStatus) GetEpochId() int64`

GetEpochId returns the EpochId field if non-nil, zero value otherwise.

### GetEpochIdOk

`func (o *NetworkSecurityPolicyStatus) GetEpochIdOk() (*int64, bool)`

GetEpochIdOk returns a tuple with the EpochId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpochId

`func (o *NetworkSecurityPolicyStatus) SetEpochId(v int64)`

SetEpochId sets EpochId field to given value.

### HasEpochId

`func (o *NetworkSecurityPolicyStatus) HasEpochId() bool`

HasEpochId returns a boolean if a field has been set.

### GetPolicyReferences

`func (o *NetworkSecurityPolicyStatus) GetPolicyReferences() []NetworkPolicyReference`

GetPolicyReferences returns the PolicyReferences field if non-nil, zero value otherwise.

### GetPolicyReferencesOk

`func (o *NetworkSecurityPolicyStatus) GetPolicyReferencesOk() (*[]NetworkPolicyReference, bool)`

GetPolicyReferencesOk returns a tuple with the PolicyReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyReferences

`func (o *NetworkSecurityPolicyStatus) SetPolicyReferences(v []NetworkPolicyReference)`

SetPolicyReferences sets PolicyReferences field to given value.

### HasPolicyReferences

`func (o *NetworkSecurityPolicyStatus) HasPolicyReferences() bool`

HasPolicyReferences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


