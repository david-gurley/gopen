# ClusterMemInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size** | Pointer to **string** | Memory size in bytes, eg: 274760318976. | [optional] 
**Type** | Pointer to **string** | Type. | [optional] [default to "unknown"]

## Methods

### NewClusterMemInfo

`func NewClusterMemInfo() *ClusterMemInfo`

NewClusterMemInfo instantiates a new ClusterMemInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterMemInfoWithDefaults

`func NewClusterMemInfoWithDefaults() *ClusterMemInfo`

NewClusterMemInfoWithDefaults instantiates a new ClusterMemInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSize

`func (o *ClusterMemInfo) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ClusterMemInfo) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ClusterMemInfo) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *ClusterMemInfo) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetType

`func (o *ClusterMemInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClusterMemInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClusterMemInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ClusterMemInfo) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


