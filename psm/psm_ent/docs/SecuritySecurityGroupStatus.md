# SecuritySecurityGroupStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policies** | Pointer to **[]string** | list of all policies attached to this security group. | [optional] 
**Workloads** | Pointer to **[]string** | list of workloads that are part of this security group. | [optional] 

## Methods

### NewSecuritySecurityGroupStatus

`func NewSecuritySecurityGroupStatus() *SecuritySecurityGroupStatus`

NewSecuritySecurityGroupStatus instantiates a new SecuritySecurityGroupStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecuritySecurityGroupStatusWithDefaults

`func NewSecuritySecurityGroupStatusWithDefaults() *SecuritySecurityGroupStatus`

NewSecuritySecurityGroupStatusWithDefaults instantiates a new SecuritySecurityGroupStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicies

`func (o *SecuritySecurityGroupStatus) GetPolicies() []string`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *SecuritySecurityGroupStatus) GetPoliciesOk() (*[]string, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *SecuritySecurityGroupStatus) SetPolicies(v []string)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *SecuritySecurityGroupStatus) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetWorkloads

`func (o *SecuritySecurityGroupStatus) GetWorkloads() []string`

GetWorkloads returns the Workloads field if non-nil, zero value otherwise.

### GetWorkloadsOk

`func (o *SecuritySecurityGroupStatus) GetWorkloadsOk() (*[]string, bool)`

GetWorkloadsOk returns a tuple with the Workloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloads

`func (o *SecuritySecurityGroupStatus) SetWorkloads(v []string)`

SetWorkloads sets Workloads field to given value.

### HasWorkloads

`func (o *SecuritySecurityGroupStatus) HasWorkloads() bool`

HasWorkloads returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


