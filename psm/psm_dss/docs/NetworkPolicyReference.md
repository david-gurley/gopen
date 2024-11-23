# NetworkPolicyReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the policy. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | [optional] 
**Uuid** | Pointer to **string** | UUID of the policy. | [optional] 

## Methods

### NewNetworkPolicyReference

`func NewNetworkPolicyReference() *NetworkPolicyReference`

NewNetworkPolicyReference instantiates a new NetworkPolicyReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkPolicyReferenceWithDefaults

`func NewNetworkPolicyReferenceWithDefaults() *NetworkPolicyReference`

NewNetworkPolicyReferenceWithDefaults instantiates a new NetworkPolicyReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NetworkPolicyReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkPolicyReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkPolicyReference) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkPolicyReference) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUuid

`func (o *NetworkPolicyReference) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *NetworkPolicyReference) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *NetworkPolicyReference) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *NetworkPolicyReference) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


