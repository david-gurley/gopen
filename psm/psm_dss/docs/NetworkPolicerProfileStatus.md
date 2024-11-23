# NetworkPolicerProfileStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PropagationStatus** | Pointer to [**SecurityPropagationStatus**](securityPropagationStatus.md) |  | [optional] 

## Methods

### NewNetworkPolicerProfileStatus

`func NewNetworkPolicerProfileStatus() *NetworkPolicerProfileStatus`

NewNetworkPolicerProfileStatus instantiates a new NetworkPolicerProfileStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkPolicerProfileStatusWithDefaults

`func NewNetworkPolicerProfileStatusWithDefaults() *NetworkPolicerProfileStatus`

NewNetworkPolicerProfileStatusWithDefaults instantiates a new NetworkPolicerProfileStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPropagationStatus

`func (o *NetworkPolicerProfileStatus) GetPropagationStatus() SecurityPropagationStatus`

GetPropagationStatus returns the PropagationStatus field if non-nil, zero value otherwise.

### GetPropagationStatusOk

`func (o *NetworkPolicerProfileStatus) GetPropagationStatusOk() (*SecurityPropagationStatus, bool)`

GetPropagationStatusOk returns a tuple with the PropagationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropagationStatus

`func (o *NetworkPolicerProfileStatus) SetPropagationStatus(v SecurityPropagationStatus)`

SetPropagationStatus sets PropagationStatus field to given value.

### HasPropagationStatus

`func (o *NetworkPolicerProfileStatus) HasPropagationStatus() bool`

HasPropagationStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


