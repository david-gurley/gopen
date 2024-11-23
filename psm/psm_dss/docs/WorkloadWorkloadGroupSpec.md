# WorkloadWorkloadGroupSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpBlock** | Pointer to [**[]WorkloadIPBlock**](WorkloadIPBlock.md) | IP Block selector for entities outside of cluster if used alone When used in conjuction with WorkloadSelector, The Workload selected based on LabelSelector or if IPAddress is part of the IPBlock entries Each IPBlock are ORed. | [optional] 
**WorkloadSelector** | Pointer to [**[]LabelsSelector**](LabelsSelector.md) | Workload selector is a list of label selectors. Each WorkloadSelector are ORed. | [optional] 

## Methods

### NewWorkloadWorkloadGroupSpec

`func NewWorkloadWorkloadGroupSpec() *WorkloadWorkloadGroupSpec`

NewWorkloadWorkloadGroupSpec instantiates a new WorkloadWorkloadGroupSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadWorkloadGroupSpecWithDefaults

`func NewWorkloadWorkloadGroupSpecWithDefaults() *WorkloadWorkloadGroupSpec`

NewWorkloadWorkloadGroupSpecWithDefaults instantiates a new WorkloadWorkloadGroupSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpBlock

`func (o *WorkloadWorkloadGroupSpec) GetIpBlock() []WorkloadIPBlock`

GetIpBlock returns the IpBlock field if non-nil, zero value otherwise.

### GetIpBlockOk

`func (o *WorkloadWorkloadGroupSpec) GetIpBlockOk() (*[]WorkloadIPBlock, bool)`

GetIpBlockOk returns a tuple with the IpBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpBlock

`func (o *WorkloadWorkloadGroupSpec) SetIpBlock(v []WorkloadIPBlock)`

SetIpBlock sets IpBlock field to given value.

### HasIpBlock

`func (o *WorkloadWorkloadGroupSpec) HasIpBlock() bool`

HasIpBlock returns a boolean if a field has been set.

### GetWorkloadSelector

`func (o *WorkloadWorkloadGroupSpec) GetWorkloadSelector() []LabelsSelector`

GetWorkloadSelector returns the WorkloadSelector field if non-nil, zero value otherwise.

### GetWorkloadSelectorOk

`func (o *WorkloadWorkloadGroupSpec) GetWorkloadSelectorOk() (*[]LabelsSelector, bool)`

GetWorkloadSelectorOk returns a tuple with the WorkloadSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadSelector

`func (o *WorkloadWorkloadGroupSpec) SetWorkloadSelector(v []LabelsSelector)`

SetWorkloadSelector sets WorkloadSelector field to given value.

### HasWorkloadSelector

`func (o *WorkloadWorkloadGroupSpec) HasWorkloadSelector() bool`

HasWorkloadSelector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


