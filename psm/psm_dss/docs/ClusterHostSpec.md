# ClusterHostSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dscs** | Pointer to [**[]ClusterDistributedServiceCardID**](ClusterDistributedServiceCardID.md) | DSCs contains the information about all DistributedServiceCards on a host. | [optional] 
**PnicInfo** | Pointer to [**[]ClusterPnicInfo**](ClusterPnicInfo.md) | information of the pnics associated with this host. | [optional] 

## Methods

### NewClusterHostSpec

`func NewClusterHostSpec() *ClusterHostSpec`

NewClusterHostSpec instantiates a new ClusterHostSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterHostSpecWithDefaults

`func NewClusterHostSpecWithDefaults() *ClusterHostSpec`

NewClusterHostSpecWithDefaults instantiates a new ClusterHostSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDscs

`func (o *ClusterHostSpec) GetDscs() []ClusterDistributedServiceCardID`

GetDscs returns the Dscs field if non-nil, zero value otherwise.

### GetDscsOk

`func (o *ClusterHostSpec) GetDscsOk() (*[]ClusterDistributedServiceCardID, bool)`

GetDscsOk returns a tuple with the Dscs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDscs

`func (o *ClusterHostSpec) SetDscs(v []ClusterDistributedServiceCardID)`

SetDscs sets Dscs field to given value.

### HasDscs

`func (o *ClusterHostSpec) HasDscs() bool`

HasDscs returns a boolean if a field has been set.

### GetPnicInfo

`func (o *ClusterHostSpec) GetPnicInfo() []ClusterPnicInfo`

GetPnicInfo returns the PnicInfo field if non-nil, zero value otherwise.

### GetPnicInfoOk

`func (o *ClusterHostSpec) GetPnicInfoOk() (*[]ClusterPnicInfo, bool)`

GetPnicInfoOk returns a tuple with the PnicInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPnicInfo

`func (o *ClusterHostSpec) SetPnicInfo(v []ClusterPnicInfo)`

SetPnicInfo sets PnicInfo field to given value.

### HasPnicInfo

`func (o *ClusterHostSpec) HasPnicInfo() bool`

HasPnicInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


