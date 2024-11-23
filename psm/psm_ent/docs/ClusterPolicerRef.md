# ClusterPolicerRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tenant** | Pointer to **string** | Tenant is the tenant to which policerprofile belongs to. | [optional] [default to "default"]
**TxPolicer** | Pointer to **string** | TxPolicer is the name of the policerprofile to be applied in Tx direction. | [optional] 

## Methods

### NewClusterPolicerRef

`func NewClusterPolicerRef() *ClusterPolicerRef`

NewClusterPolicerRef instantiates a new ClusterPolicerRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterPolicerRefWithDefaults

`func NewClusterPolicerRefWithDefaults() *ClusterPolicerRef`

NewClusterPolicerRefWithDefaults instantiates a new ClusterPolicerRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenant

`func (o *ClusterPolicerRef) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *ClusterPolicerRef) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *ClusterPolicerRef) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *ClusterPolicerRef) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetTxPolicer

`func (o *ClusterPolicerRef) GetTxPolicer() string`

GetTxPolicer returns the TxPolicer field if non-nil, zero value otherwise.

### GetTxPolicerOk

`func (o *ClusterPolicerRef) GetTxPolicerOk() (*string, bool)`

GetTxPolicerOk returns a tuple with the TxPolicer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxPolicer

`func (o *ClusterPolicerRef) SetTxPolicer(v string)`

SetTxPolicer sets TxPolicer field to given value.

### HasTxPolicer

`func (o *ClusterPolicerRef) HasTxPolicer() bool`

HasTxPolicer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


