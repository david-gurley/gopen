# ClusterSearchOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableFwlogSearch** | Pointer to **bool** | EnableFwlogSearch when enabled peforms reverse indexing of fwlogs getting ingested into PSM. The reverse indexed logs can be queried using the fwlog query/search API. | [optional] [default to true]

## Methods

### NewClusterSearchOptions

`func NewClusterSearchOptions() *ClusterSearchOptions`

NewClusterSearchOptions instantiates a new ClusterSearchOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterSearchOptionsWithDefaults

`func NewClusterSearchOptionsWithDefaults() *ClusterSearchOptions`

NewClusterSearchOptionsWithDefaults instantiates a new ClusterSearchOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableFwlogSearch

`func (o *ClusterSearchOptions) GetEnableFwlogSearch() bool`

GetEnableFwlogSearch returns the EnableFwlogSearch field if non-nil, zero value otherwise.

### GetEnableFwlogSearchOk

`func (o *ClusterSearchOptions) GetEnableFwlogSearchOk() (*bool, bool)`

GetEnableFwlogSearchOk returns a tuple with the EnableFwlogSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFwlogSearch

`func (o *ClusterSearchOptions) SetEnableFwlogSearch(v bool)`

SetEnableFwlogSearch sets EnableFwlogSearch field to given value.

### HasEnableFwlogSearch

`func (o *ClusterSearchOptions) HasEnableFwlogSearch() bool`

HasEnableFwlogSearch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


