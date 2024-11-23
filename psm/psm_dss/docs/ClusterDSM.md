# ClusterDSM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MacAddress** | Pointer to **string** | Base MAC address of. | [optional] 
**UnitId** | Pointer to **int64** | unit id of the dsm. | [optional] 

## Methods

### NewClusterDSM

`func NewClusterDSM() *ClusterDSM`

NewClusterDSM instantiates a new ClusterDSM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterDSMWithDefaults

`func NewClusterDSMWithDefaults() *ClusterDSM`

NewClusterDSMWithDefaults instantiates a new ClusterDSM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMacAddress

`func (o *ClusterDSM) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *ClusterDSM) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *ClusterDSM) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *ClusterDSM) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetUnitId

`func (o *ClusterDSM) GetUnitId() int64`

GetUnitId returns the UnitId field if non-nil, zero value otherwise.

### GetUnitIdOk

`func (o *ClusterDSM) GetUnitIdOk() (*int64, bool)`

GetUnitIdOk returns a tuple with the UnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitId

`func (o *ClusterDSM) SetUnitId(v int64)`

SetUnitId sets UnitId field to given value.

### HasUnitId

`func (o *ClusterDSM) HasUnitId() bool`

HasUnitId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


