# ClusterDistributedServiceCardID

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Name contains the name of the DistributedServiceCard on a host. | [optional] 
**MacAddress** | Pointer to **string** | MACAddress contains the primary MAC address of a DistributedServiceCard. Should be a valid MAC address. | [optional] 

## Methods

### NewClusterDistributedServiceCardID

`func NewClusterDistributedServiceCardID() *ClusterDistributedServiceCardID`

NewClusterDistributedServiceCardID instantiates a new ClusterDistributedServiceCardID object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterDistributedServiceCardIDWithDefaults

`func NewClusterDistributedServiceCardIDWithDefaults() *ClusterDistributedServiceCardID`

NewClusterDistributedServiceCardIDWithDefaults instantiates a new ClusterDistributedServiceCardID object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClusterDistributedServiceCardID) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterDistributedServiceCardID) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterDistributedServiceCardID) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ClusterDistributedServiceCardID) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMacAddress

`func (o *ClusterDistributedServiceCardID) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *ClusterDistributedServiceCardID) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *ClusterDistributedServiceCardID) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *ClusterDistributedServiceCardID) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


