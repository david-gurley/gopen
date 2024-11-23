# SearchSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregate** | Pointer to **bool** | Indicates whether to perform aggregation on the search results or not. | [optional] [default to true]
**From** | Pointer to **int32** | From represents the start offset (zero based), used in paginated search requests The results returned would be in the range [From ... From+MaxResults-1] This can be specified as URI parameter. Default value is 0 and valid range is 0..8192. Value should be between 0 and 8192. | [optional] 
**MaxResults** | Pointer to **int32** | MaxResults is the max-count of search results This can be specified as URI parameter. Default value is 50 and valid range is 0..8192. Value should be between 0 and 8192. | [optional] [default to 50]
**Mode** | Pointer to **string** | Query Mode. | [optional] [default to "full"]
**Query** | Pointer to [**SearchSearchQuery**](searchSearchQuery.md) |  | [optional] 
**QueryString** | Pointer to **string** | Simple query string This can be specified as URI parameter. For advanced query cases, Users should use specify SearchQuery and pass the SearchRequest in a GET/POST Body The max query-string length is 256 bytes. Length of string should be between 0 and 256. | [optional] 
**SortBy** | Pointer to **string** | SortyBy is an optional parameter and contains the field name to be sorted by, For eg: \&quot;meta.name\&quot; This can be specified as URI parameter. Length of string should be between 0 and 256. | [optional] 
**SortOrder** | Pointer to **string** | SortOrder is an optional parameter and contains whether to sort ascending or descending This can be specified as URI parameter. | [optional] [default to "ascending"]
**Tenants** | Pointer to **[]string** | OR of tenants within the scope of which search needs to be performed. If not specified, it will be set to tenant of the logged in user. | [optional] 

## Methods

### NewSearchSearchRequest

`func NewSearchSearchRequest() *SearchSearchRequest`

NewSearchSearchRequest instantiates a new SearchSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchSearchRequestWithDefaults

`func NewSearchSearchRequestWithDefaults() *SearchSearchRequest`

NewSearchSearchRequestWithDefaults instantiates a new SearchSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregate

`func (o *SearchSearchRequest) GetAggregate() bool`

GetAggregate returns the Aggregate field if non-nil, zero value otherwise.

### GetAggregateOk

`func (o *SearchSearchRequest) GetAggregateOk() (*bool, bool)`

GetAggregateOk returns a tuple with the Aggregate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregate

`func (o *SearchSearchRequest) SetAggregate(v bool)`

SetAggregate sets Aggregate field to given value.

### HasAggregate

`func (o *SearchSearchRequest) HasAggregate() bool`

HasAggregate returns a boolean if a field has been set.

### GetFrom

`func (o *SearchSearchRequest) GetFrom() int32`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *SearchSearchRequest) GetFromOk() (*int32, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *SearchSearchRequest) SetFrom(v int32)`

SetFrom sets From field to given value.

### HasFrom

`func (o *SearchSearchRequest) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetMaxResults

`func (o *SearchSearchRequest) GetMaxResults() int32`

GetMaxResults returns the MaxResults field if non-nil, zero value otherwise.

### GetMaxResultsOk

`func (o *SearchSearchRequest) GetMaxResultsOk() (*int32, bool)`

GetMaxResultsOk returns a tuple with the MaxResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResults

`func (o *SearchSearchRequest) SetMaxResults(v int32)`

SetMaxResults sets MaxResults field to given value.

### HasMaxResults

`func (o *SearchSearchRequest) HasMaxResults() bool`

HasMaxResults returns a boolean if a field has been set.

### GetMode

`func (o *SearchSearchRequest) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *SearchSearchRequest) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *SearchSearchRequest) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *SearchSearchRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetQuery

`func (o *SearchSearchRequest) GetQuery() SearchSearchQuery`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SearchSearchRequest) GetQueryOk() (*SearchSearchQuery, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SearchSearchRequest) SetQuery(v SearchSearchQuery)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *SearchSearchRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetQueryString

`func (o *SearchSearchRequest) GetQueryString() string`

GetQueryString returns the QueryString field if non-nil, zero value otherwise.

### GetQueryStringOk

`func (o *SearchSearchRequest) GetQueryStringOk() (*string, bool)`

GetQueryStringOk returns a tuple with the QueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryString

`func (o *SearchSearchRequest) SetQueryString(v string)`

SetQueryString sets QueryString field to given value.

### HasQueryString

`func (o *SearchSearchRequest) HasQueryString() bool`

HasQueryString returns a boolean if a field has been set.

### GetSortBy

`func (o *SearchSearchRequest) GetSortBy() string`

GetSortBy returns the SortBy field if non-nil, zero value otherwise.

### GetSortByOk

`func (o *SearchSearchRequest) GetSortByOk() (*string, bool)`

GetSortByOk returns a tuple with the SortBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortBy

`func (o *SearchSearchRequest) SetSortBy(v string)`

SetSortBy sets SortBy field to given value.

### HasSortBy

`func (o *SearchSearchRequest) HasSortBy() bool`

HasSortBy returns a boolean if a field has been set.

### GetSortOrder

`func (o *SearchSearchRequest) GetSortOrder() string`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *SearchSearchRequest) GetSortOrderOk() (*string, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *SearchSearchRequest) SetSortOrder(v string)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *SearchSearchRequest) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetTenants

`func (o *SearchSearchRequest) GetTenants() []string`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *SearchSearchRequest) GetTenantsOk() (*[]string, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *SearchSearchRequest) SetTenants(v []string)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *SearchSearchRequest) HasTenants() bool`

HasTenants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


