# ClusterOsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KernelRelease** | Pointer to **string** | Kernel release Eg: 3.10.0-514.10.2.el7.x86_64. | [optional] 
**KernelVersion** | Pointer to **string** | Kernel version Eg: #1 SMP Fri Mar 3 00:04:05 UTC 2017. | [optional] 
**Processor** | Pointer to **string** | Processor Info Eg: x86_64. | [optional] 
**Type** | Pointer to **string** | OS Name Eg: GNU/Linux. | [optional] 

## Methods

### NewClusterOsInfo

`func NewClusterOsInfo() *ClusterOsInfo`

NewClusterOsInfo instantiates a new ClusterOsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterOsInfoWithDefaults

`func NewClusterOsInfoWithDefaults() *ClusterOsInfo`

NewClusterOsInfoWithDefaults instantiates a new ClusterOsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKernelRelease

`func (o *ClusterOsInfo) GetKernelRelease() string`

GetKernelRelease returns the KernelRelease field if non-nil, zero value otherwise.

### GetKernelReleaseOk

`func (o *ClusterOsInfo) GetKernelReleaseOk() (*string, bool)`

GetKernelReleaseOk returns a tuple with the KernelRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKernelRelease

`func (o *ClusterOsInfo) SetKernelRelease(v string)`

SetKernelRelease sets KernelRelease field to given value.

### HasKernelRelease

`func (o *ClusterOsInfo) HasKernelRelease() bool`

HasKernelRelease returns a boolean if a field has been set.

### GetKernelVersion

`func (o *ClusterOsInfo) GetKernelVersion() string`

GetKernelVersion returns the KernelVersion field if non-nil, zero value otherwise.

### GetKernelVersionOk

`func (o *ClusterOsInfo) GetKernelVersionOk() (*string, bool)`

GetKernelVersionOk returns a tuple with the KernelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKernelVersion

`func (o *ClusterOsInfo) SetKernelVersion(v string)`

SetKernelVersion sets KernelVersion field to given value.

### HasKernelVersion

`func (o *ClusterOsInfo) HasKernelVersion() bool`

HasKernelVersion returns a boolean if a field has been set.

### GetProcessor

`func (o *ClusterOsInfo) GetProcessor() string`

GetProcessor returns the Processor field if non-nil, zero value otherwise.

### GetProcessorOk

`func (o *ClusterOsInfo) GetProcessorOk() (*string, bool)`

GetProcessorOk returns a tuple with the Processor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessor

`func (o *ClusterOsInfo) SetProcessor(v string)`

SetProcessor sets Processor field to given value.

### HasProcessor

`func (o *ClusterOsInfo) HasProcessor() bool`

HasProcessor returns a boolean if a field has been set.

### GetType

`func (o *ClusterOsInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClusterOsInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClusterOsInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ClusterOsInfo) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


