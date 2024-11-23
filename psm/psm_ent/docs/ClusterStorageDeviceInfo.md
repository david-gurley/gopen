# ClusterStorageDeviceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capacity** | Pointer to **string** | Capacity in bytes. | [optional] 
**PercentLifeUsedA** | Pointer to **int32** | Used life in percentage. | [optional] 
**PercentLifeUsedB** | Pointer to **int32** |  | [optional] 
**SerialNum** | Pointer to **string** | Serial Number. | [optional] 
**Type** | Pointer to **string** | Storage Type (TBD for Naples) Eg: SATA, SCSI, NVMe  or HDD, SSD, NVMe. | [optional] 
**Vendor** | Pointer to **string** | Vendor info. | [optional] 

## Methods

### NewClusterStorageDeviceInfo

`func NewClusterStorageDeviceInfo() *ClusterStorageDeviceInfo`

NewClusterStorageDeviceInfo instantiates a new ClusterStorageDeviceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterStorageDeviceInfoWithDefaults

`func NewClusterStorageDeviceInfoWithDefaults() *ClusterStorageDeviceInfo`

NewClusterStorageDeviceInfoWithDefaults instantiates a new ClusterStorageDeviceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacity

`func (o *ClusterStorageDeviceInfo) GetCapacity() string`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *ClusterStorageDeviceInfo) GetCapacityOk() (*string, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *ClusterStorageDeviceInfo) SetCapacity(v string)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *ClusterStorageDeviceInfo) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetPercentLifeUsedA

`func (o *ClusterStorageDeviceInfo) GetPercentLifeUsedA() int32`

GetPercentLifeUsedA returns the PercentLifeUsedA field if non-nil, zero value otherwise.

### GetPercentLifeUsedAOk

`func (o *ClusterStorageDeviceInfo) GetPercentLifeUsedAOk() (*int32, bool)`

GetPercentLifeUsedAOk returns a tuple with the PercentLifeUsedA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentLifeUsedA

`func (o *ClusterStorageDeviceInfo) SetPercentLifeUsedA(v int32)`

SetPercentLifeUsedA sets PercentLifeUsedA field to given value.

### HasPercentLifeUsedA

`func (o *ClusterStorageDeviceInfo) HasPercentLifeUsedA() bool`

HasPercentLifeUsedA returns a boolean if a field has been set.

### GetPercentLifeUsedB

`func (o *ClusterStorageDeviceInfo) GetPercentLifeUsedB() int32`

GetPercentLifeUsedB returns the PercentLifeUsedB field if non-nil, zero value otherwise.

### GetPercentLifeUsedBOk

`func (o *ClusterStorageDeviceInfo) GetPercentLifeUsedBOk() (*int32, bool)`

GetPercentLifeUsedBOk returns a tuple with the PercentLifeUsedB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentLifeUsedB

`func (o *ClusterStorageDeviceInfo) SetPercentLifeUsedB(v int32)`

SetPercentLifeUsedB sets PercentLifeUsedB field to given value.

### HasPercentLifeUsedB

`func (o *ClusterStorageDeviceInfo) HasPercentLifeUsedB() bool`

HasPercentLifeUsedB returns a boolean if a field has been set.

### GetSerialNum

`func (o *ClusterStorageDeviceInfo) GetSerialNum() string`

GetSerialNum returns the SerialNum field if non-nil, zero value otherwise.

### GetSerialNumOk

`func (o *ClusterStorageDeviceInfo) GetSerialNumOk() (*string, bool)`

GetSerialNumOk returns a tuple with the SerialNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNum

`func (o *ClusterStorageDeviceInfo) SetSerialNum(v string)`

SetSerialNum sets SerialNum field to given value.

### HasSerialNum

`func (o *ClusterStorageDeviceInfo) HasSerialNum() bool`

HasSerialNum returns a boolean if a field has been set.

### GetType

`func (o *ClusterStorageDeviceInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClusterStorageDeviceInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClusterStorageDeviceInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ClusterStorageDeviceInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVendor

`func (o *ClusterStorageDeviceInfo) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *ClusterStorageDeviceInfo) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *ClusterStorageDeviceInfo) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *ClusterStorageDeviceInfo) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


