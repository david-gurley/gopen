# SecurityAppStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttachedPolicies** | Pointer to **[]string** | List of security group policies attached to the app. | [optional] 

## Methods

### NewSecurityAppStatus

`func NewSecurityAppStatus() *SecurityAppStatus`

NewSecurityAppStatus instantiates a new SecurityAppStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityAppStatusWithDefaults

`func NewSecurityAppStatusWithDefaults() *SecurityAppStatus`

NewSecurityAppStatusWithDefaults instantiates a new SecurityAppStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachedPolicies

`func (o *SecurityAppStatus) GetAttachedPolicies() []string`

GetAttachedPolicies returns the AttachedPolicies field if non-nil, zero value otherwise.

### GetAttachedPoliciesOk

`func (o *SecurityAppStatus) GetAttachedPoliciesOk() (*[]string, bool)`

GetAttachedPoliciesOk returns a tuple with the AttachedPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedPolicies

`func (o *SecurityAppStatus) SetAttachedPolicies(v []string)`

SetAttachedPolicies sets AttachedPolicies field to given value.

### HasAttachedPolicies

`func (o *SecurityAppStatus) HasAttachedPolicies() bool`

HasAttachedPolicies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


