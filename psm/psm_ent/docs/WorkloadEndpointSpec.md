# WorkloadEndpointSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HomingHostAddr** | Pointer to **string** | IP of the DSC where this endpoint exists. | [optional] 
**MicroSegmentVlan** | Pointer to **int64** | MicroSegmentVlan to be assigned to the endpoint. | [optional] 
**NodeUuid** | Pointer to **string** | The DSC Name or MAC where the endpoint should reside. | [optional] 
**NodeUuidList** | Pointer to **[]string** | NodeUUIDList has the list of DSCs where a EP is supposed to go to. | [optional] 
**Type** | Pointer to **string** | Type is the type of Endpoint that is being created - L2 or L3. | [optional] [default to "l2"]

## Methods

### NewWorkloadEndpointSpec

`func NewWorkloadEndpointSpec() *WorkloadEndpointSpec`

NewWorkloadEndpointSpec instantiates a new WorkloadEndpointSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadEndpointSpecWithDefaults

`func NewWorkloadEndpointSpecWithDefaults() *WorkloadEndpointSpec`

NewWorkloadEndpointSpecWithDefaults instantiates a new WorkloadEndpointSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHomingHostAddr

`func (o *WorkloadEndpointSpec) GetHomingHostAddr() string`

GetHomingHostAddr returns the HomingHostAddr field if non-nil, zero value otherwise.

### GetHomingHostAddrOk

`func (o *WorkloadEndpointSpec) GetHomingHostAddrOk() (*string, bool)`

GetHomingHostAddrOk returns a tuple with the HomingHostAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomingHostAddr

`func (o *WorkloadEndpointSpec) SetHomingHostAddr(v string)`

SetHomingHostAddr sets HomingHostAddr field to given value.

### HasHomingHostAddr

`func (o *WorkloadEndpointSpec) HasHomingHostAddr() bool`

HasHomingHostAddr returns a boolean if a field has been set.

### GetMicroSegmentVlan

`func (o *WorkloadEndpointSpec) GetMicroSegmentVlan() int64`

GetMicroSegmentVlan returns the MicroSegmentVlan field if non-nil, zero value otherwise.

### GetMicroSegmentVlanOk

`func (o *WorkloadEndpointSpec) GetMicroSegmentVlanOk() (*int64, bool)`

GetMicroSegmentVlanOk returns a tuple with the MicroSegmentVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroSegmentVlan

`func (o *WorkloadEndpointSpec) SetMicroSegmentVlan(v int64)`

SetMicroSegmentVlan sets MicroSegmentVlan field to given value.

### HasMicroSegmentVlan

`func (o *WorkloadEndpointSpec) HasMicroSegmentVlan() bool`

HasMicroSegmentVlan returns a boolean if a field has been set.

### GetNodeUuid

`func (o *WorkloadEndpointSpec) GetNodeUuid() string`

GetNodeUuid returns the NodeUuid field if non-nil, zero value otherwise.

### GetNodeUuidOk

`func (o *WorkloadEndpointSpec) GetNodeUuidOk() (*string, bool)`

GetNodeUuidOk returns a tuple with the NodeUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeUuid

`func (o *WorkloadEndpointSpec) SetNodeUuid(v string)`

SetNodeUuid sets NodeUuid field to given value.

### HasNodeUuid

`func (o *WorkloadEndpointSpec) HasNodeUuid() bool`

HasNodeUuid returns a boolean if a field has been set.

### GetNodeUuidList

`func (o *WorkloadEndpointSpec) GetNodeUuidList() []string`

GetNodeUuidList returns the NodeUuidList field if non-nil, zero value otherwise.

### GetNodeUuidListOk

`func (o *WorkloadEndpointSpec) GetNodeUuidListOk() (*[]string, bool)`

GetNodeUuidListOk returns a tuple with the NodeUuidList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeUuidList

`func (o *WorkloadEndpointSpec) SetNodeUuidList(v []string)`

SetNodeUuidList sets NodeUuidList field to given value.

### HasNodeUuidList

`func (o *WorkloadEndpointSpec) HasNodeUuidList() bool`

HasNodeUuidList returns a boolean if a field has been set.

### GetType

`func (o *WorkloadEndpointSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkloadEndpointSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkloadEndpointSpec) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WorkloadEndpointSpec) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


