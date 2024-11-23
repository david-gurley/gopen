# ClusterCredentialsSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyValuePairs** | Pointer to [**[]ClusterKeyValue**](ClusterKeyValue.md) | KeyValuePairs contains all key,value pairs that constitute credentials to access a service. | [optional] 

## Methods

### NewClusterCredentialsSpec

`func NewClusterCredentialsSpec() *ClusterCredentialsSpec`

NewClusterCredentialsSpec instantiates a new ClusterCredentialsSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterCredentialsSpecWithDefaults

`func NewClusterCredentialsSpecWithDefaults() *ClusterCredentialsSpec`

NewClusterCredentialsSpecWithDefaults instantiates a new ClusterCredentialsSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyValuePairs

`func (o *ClusterCredentialsSpec) GetKeyValuePairs() []ClusterKeyValue`

GetKeyValuePairs returns the KeyValuePairs field if non-nil, zero value otherwise.

### GetKeyValuePairsOk

`func (o *ClusterCredentialsSpec) GetKeyValuePairsOk() (*[]ClusterKeyValue, bool)`

GetKeyValuePairsOk returns a tuple with the KeyValuePairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyValuePairs

`func (o *ClusterCredentialsSpec) SetKeyValuePairs(v []ClusterKeyValue)`

SetKeyValuePairs sets KeyValuePairs field to given value.

### HasKeyValuePairs

`func (o *ClusterCredentialsSpec) HasKeyValuePairs() bool`

HasKeyValuePairs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


