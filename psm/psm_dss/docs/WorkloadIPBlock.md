# WorkloadIPBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cidr** | Pointer to **string** | CIDRs prefix networks. Should be a valid v4 or v6 CIDR block. | [optional] 

## Methods

### NewWorkloadIPBlock

`func NewWorkloadIPBlock() *WorkloadIPBlock`

NewWorkloadIPBlock instantiates a new WorkloadIPBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadIPBlockWithDefaults

`func NewWorkloadIPBlockWithDefaults() *WorkloadIPBlock`

NewWorkloadIPBlockWithDefaults instantiates a new WorkloadIPBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCidr

`func (o *WorkloadIPBlock) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *WorkloadIPBlock) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *WorkloadIPBlock) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *WorkloadIPBlock) HasCidr() bool`

HasCidr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


