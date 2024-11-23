# NetworkServiceStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Workloads** | Pointer to **[]string** | list of workloads that are backends of this service. | [optional] 

## Methods

### NewNetworkServiceStatus

`func NewNetworkServiceStatus() *NetworkServiceStatus`

NewNetworkServiceStatus instantiates a new NetworkServiceStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkServiceStatusWithDefaults

`func NewNetworkServiceStatusWithDefaults() *NetworkServiceStatus`

NewNetworkServiceStatusWithDefaults instantiates a new NetworkServiceStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkloads

`func (o *NetworkServiceStatus) GetWorkloads() []string`

GetWorkloads returns the Workloads field if non-nil, zero value otherwise.

### GetWorkloadsOk

`func (o *NetworkServiceStatus) GetWorkloadsOk() (*[]string, bool)`

GetWorkloadsOk returns a tuple with the Workloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloads

`func (o *NetworkServiceStatus) SetWorkloads(v []string)`

SetWorkloads sets Workloads field to given value.

### HasWorkloads

`func (o *NetworkServiceStatus) HasWorkloads() bool`

HasWorkloads returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


