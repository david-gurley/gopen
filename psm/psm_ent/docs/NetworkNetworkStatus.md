# NetworkNetworkStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllocatedIpv4Addrs** | Pointer to **string** | allocated IPv4 addresses (bitmap). | [optional] 
**Id** | Pointer to **string** | Handle is the internal Handle allocated to this network. | [optional] 
**OperState** | Pointer to **string** |  | [optional] [default to "active"]
**PropagationStatus** | Pointer to [**SecurityPropagationStatus**](securityPropagationStatus.md) |  | [optional] 
**Workloads** | Pointer to **[]string** | list of all workloads in this network. | [optional] 

## Methods

### NewNetworkNetworkStatus

`func NewNetworkNetworkStatus() *NetworkNetworkStatus`

NewNetworkNetworkStatus instantiates a new NetworkNetworkStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkNetworkStatusWithDefaults

`func NewNetworkNetworkStatusWithDefaults() *NetworkNetworkStatus`

NewNetworkNetworkStatusWithDefaults instantiates a new NetworkNetworkStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllocatedIpv4Addrs

`func (o *NetworkNetworkStatus) GetAllocatedIpv4Addrs() string`

GetAllocatedIpv4Addrs returns the AllocatedIpv4Addrs field if non-nil, zero value otherwise.

### GetAllocatedIpv4AddrsOk

`func (o *NetworkNetworkStatus) GetAllocatedIpv4AddrsOk() (*string, bool)`

GetAllocatedIpv4AddrsOk returns a tuple with the AllocatedIpv4Addrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedIpv4Addrs

`func (o *NetworkNetworkStatus) SetAllocatedIpv4Addrs(v string)`

SetAllocatedIpv4Addrs sets AllocatedIpv4Addrs field to given value.

### HasAllocatedIpv4Addrs

`func (o *NetworkNetworkStatus) HasAllocatedIpv4Addrs() bool`

HasAllocatedIpv4Addrs returns a boolean if a field has been set.

### GetId

`func (o *NetworkNetworkStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkNetworkStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkNetworkStatus) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkNetworkStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOperState

`func (o *NetworkNetworkStatus) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *NetworkNetworkStatus) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *NetworkNetworkStatus) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *NetworkNetworkStatus) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetPropagationStatus

`func (o *NetworkNetworkStatus) GetPropagationStatus() SecurityPropagationStatus`

GetPropagationStatus returns the PropagationStatus field if non-nil, zero value otherwise.

### GetPropagationStatusOk

`func (o *NetworkNetworkStatus) GetPropagationStatusOk() (*SecurityPropagationStatus, bool)`

GetPropagationStatusOk returns a tuple with the PropagationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropagationStatus

`func (o *NetworkNetworkStatus) SetPropagationStatus(v SecurityPropagationStatus)`

SetPropagationStatus sets PropagationStatus field to given value.

### HasPropagationStatus

`func (o *NetworkNetworkStatus) HasPropagationStatus() bool`

HasPropagationStatus returns a boolean if a field has been set.

### GetWorkloads

`func (o *NetworkNetworkStatus) GetWorkloads() []string`

GetWorkloads returns the Workloads field if non-nil, zero value otherwise.

### GetWorkloadsOk

`func (o *NetworkNetworkStatus) GetWorkloadsOk() (*[]string, bool)`

GetWorkloadsOk returns a tuple with the Workloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloads

`func (o *NetworkNetworkStatus) SetWorkloads(v []string)`

SetWorkloads sets Workloads field to given value.

### HasWorkloads

`func (o *NetworkNetworkStatus) HasWorkloads() bool`

HasWorkloads returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


