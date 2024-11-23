# SearchSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActualHits** | Pointer to **string** | ActualHits indicates the actual hits returned in this response. | [optional] 
**AggregatedEntries** | Pointer to [**SearchTenantAggregation**](searchTenantAggregation.md) |  | [optional] 
**Entries** | Pointer to [**[]SearchEntry**](SearchEntry.md) | EntryList is list of all search results with no grouping. This attribute is populated and valid only in Full request-mode. | [optional] 
**Error** | Pointer to [**SearchError**](searchError.md) |  | [optional] 
**PreviewEntries** | Pointer to [**SearchTenantPreview**](searchTenantPreview.md) |  | [optional] 
**TimeTakenMsecs** | Pointer to **string** | TimeTakenMsecs is the time taken for search response in millisecs. | [optional] 
**TotalHits** | Pointer to **string** | TotalHits indicates total number of hits matched. | [optional] 

## Methods

### NewSearchSearchResponse

`func NewSearchSearchResponse() *SearchSearchResponse`

NewSearchSearchResponse instantiates a new SearchSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchSearchResponseWithDefaults

`func NewSearchSearchResponseWithDefaults() *SearchSearchResponse`

NewSearchSearchResponseWithDefaults instantiates a new SearchSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActualHits

`func (o *SearchSearchResponse) GetActualHits() string`

GetActualHits returns the ActualHits field if non-nil, zero value otherwise.

### GetActualHitsOk

`func (o *SearchSearchResponse) GetActualHitsOk() (*string, bool)`

GetActualHitsOk returns a tuple with the ActualHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualHits

`func (o *SearchSearchResponse) SetActualHits(v string)`

SetActualHits sets ActualHits field to given value.

### HasActualHits

`func (o *SearchSearchResponse) HasActualHits() bool`

HasActualHits returns a boolean if a field has been set.

### GetAggregatedEntries

`func (o *SearchSearchResponse) GetAggregatedEntries() SearchTenantAggregation`

GetAggregatedEntries returns the AggregatedEntries field if non-nil, zero value otherwise.

### GetAggregatedEntriesOk

`func (o *SearchSearchResponse) GetAggregatedEntriesOk() (*SearchTenantAggregation, bool)`

GetAggregatedEntriesOk returns a tuple with the AggregatedEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregatedEntries

`func (o *SearchSearchResponse) SetAggregatedEntries(v SearchTenantAggregation)`

SetAggregatedEntries sets AggregatedEntries field to given value.

### HasAggregatedEntries

`func (o *SearchSearchResponse) HasAggregatedEntries() bool`

HasAggregatedEntries returns a boolean if a field has been set.

### GetEntries

`func (o *SearchSearchResponse) GetEntries() []SearchEntry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *SearchSearchResponse) GetEntriesOk() (*[]SearchEntry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *SearchSearchResponse) SetEntries(v []SearchEntry)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *SearchSearchResponse) HasEntries() bool`

HasEntries returns a boolean if a field has been set.

### GetError

`func (o *SearchSearchResponse) GetError() SearchError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SearchSearchResponse) GetErrorOk() (*SearchError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SearchSearchResponse) SetError(v SearchError)`

SetError sets Error field to given value.

### HasError

`func (o *SearchSearchResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetPreviewEntries

`func (o *SearchSearchResponse) GetPreviewEntries() SearchTenantPreview`

GetPreviewEntries returns the PreviewEntries field if non-nil, zero value otherwise.

### GetPreviewEntriesOk

`func (o *SearchSearchResponse) GetPreviewEntriesOk() (*SearchTenantPreview, bool)`

GetPreviewEntriesOk returns a tuple with the PreviewEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewEntries

`func (o *SearchSearchResponse) SetPreviewEntries(v SearchTenantPreview)`

SetPreviewEntries sets PreviewEntries field to given value.

### HasPreviewEntries

`func (o *SearchSearchResponse) HasPreviewEntries() bool`

HasPreviewEntries returns a boolean if a field has been set.

### GetTimeTakenMsecs

`func (o *SearchSearchResponse) GetTimeTakenMsecs() string`

GetTimeTakenMsecs returns the TimeTakenMsecs field if non-nil, zero value otherwise.

### GetTimeTakenMsecsOk

`func (o *SearchSearchResponse) GetTimeTakenMsecsOk() (*string, bool)`

GetTimeTakenMsecsOk returns a tuple with the TimeTakenMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeTakenMsecs

`func (o *SearchSearchResponse) SetTimeTakenMsecs(v string)`

SetTimeTakenMsecs sets TimeTakenMsecs field to given value.

### HasTimeTakenMsecs

`func (o *SearchSearchResponse) HasTimeTakenMsecs() bool`

HasTimeTakenMsecs returns a boolean if a field has been set.

### GetTotalHits

`func (o *SearchSearchResponse) GetTotalHits() string`

GetTotalHits returns the TotalHits field if non-nil, zero value otherwise.

### GetTotalHitsOk

`func (o *SearchSearchResponse) GetTotalHitsOk() (*string, bool)`

GetTotalHitsOk returns a tuple with the TotalHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalHits

`func (o *SearchSearchResponse) SetTotalHits(v string)`

SetTotalHits sets TotalHits field to given value.

### HasTotalHits

`func (o *SearchSearchResponse) HasTotalHits() bool`

HasTotalHits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


