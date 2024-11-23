# ClusterCPUInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumCores** | Pointer to **int32** | Number of physical CPU cores per socket, eg: 36. | [optional] 
**NumSockets** | Pointer to **int32** | Number of CPU sockets, eg: 2, 4. | [optional] 
**NumThreads** | Pointer to **int32** | Number of threads per core, eg: 2. | [optional] 
**Speed** | Pointer to **string** | CPU speed per core, eg: 2099998101. | [optional] 

## Methods

### NewClusterCPUInfo

`func NewClusterCPUInfo() *ClusterCPUInfo`

NewClusterCPUInfo instantiates a new ClusterCPUInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterCPUInfoWithDefaults

`func NewClusterCPUInfoWithDefaults() *ClusterCPUInfo`

NewClusterCPUInfoWithDefaults instantiates a new ClusterCPUInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumCores

`func (o *ClusterCPUInfo) GetNumCores() int32`

GetNumCores returns the NumCores field if non-nil, zero value otherwise.

### GetNumCoresOk

`func (o *ClusterCPUInfo) GetNumCoresOk() (*int32, bool)`

GetNumCoresOk returns a tuple with the NumCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCores

`func (o *ClusterCPUInfo) SetNumCores(v int32)`

SetNumCores sets NumCores field to given value.

### HasNumCores

`func (o *ClusterCPUInfo) HasNumCores() bool`

HasNumCores returns a boolean if a field has been set.

### GetNumSockets

`func (o *ClusterCPUInfo) GetNumSockets() int32`

GetNumSockets returns the NumSockets field if non-nil, zero value otherwise.

### GetNumSocketsOk

`func (o *ClusterCPUInfo) GetNumSocketsOk() (*int32, bool)`

GetNumSocketsOk returns a tuple with the NumSockets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSockets

`func (o *ClusterCPUInfo) SetNumSockets(v int32)`

SetNumSockets sets NumSockets field to given value.

### HasNumSockets

`func (o *ClusterCPUInfo) HasNumSockets() bool`

HasNumSockets returns a boolean if a field has been set.

### GetNumThreads

`func (o *ClusterCPUInfo) GetNumThreads() int32`

GetNumThreads returns the NumThreads field if non-nil, zero value otherwise.

### GetNumThreadsOk

`func (o *ClusterCPUInfo) GetNumThreadsOk() (*int32, bool)`

GetNumThreadsOk returns a tuple with the NumThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumThreads

`func (o *ClusterCPUInfo) SetNumThreads(v int32)`

SetNumThreads sets NumThreads field to given value.

### HasNumThreads

`func (o *ClusterCPUInfo) HasNumThreads() bool`

HasNumThreads returns a boolean if a field has been set.

### GetSpeed

`func (o *ClusterCPUInfo) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *ClusterCPUInfo) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *ClusterCPUInfo) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *ClusterCPUInfo) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


