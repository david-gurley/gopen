/*
 * Search API reference
 *
 * Service name  
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package psm_ent

import (
	"encoding/json"
)

// SearchSearchResponse SearchResponse is the output provided by the search API Based on the search request, search results would be part of one of the entities : Entries or NestedAggregation. In case of failures, Error would indicate the error status and description.
type SearchSearchResponse struct {
	// ActualHits indicates the actual hits returned in this response.
	ActualHits *string `json:"actual-hits,omitempty"`
	AggregatedEntries *SearchTenantAggregation `json:"aggregated-entries,omitempty"`
	// EntryList is list of all search results with no grouping. This attribute is populated and valid only in Full request-mode.
	Entries *[]SearchEntry `json:"entries,omitempty"`
	Error *SearchError `json:"error,omitempty"`
	PreviewEntries *SearchTenantPreview `json:"preview-entries,omitempty"`
	// TimeTakenMsecs is the time taken for search response in millisecs.
	TimeTakenMsecs *string `json:"time-taken-msecs,omitempty"`
	// TotalHits indicates total number of hits matched.
	TotalHits *string `json:"total-hits,omitempty"`
}

// NewSearchSearchResponse instantiates a new SearchSearchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchSearchResponse() *SearchSearchResponse {
	this := SearchSearchResponse{}
	return &this
}

// NewSearchSearchResponseWithDefaults instantiates a new SearchSearchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchSearchResponseWithDefaults() *SearchSearchResponse {
	this := SearchSearchResponse{}
	return &this
}

// GetActualHits returns the ActualHits field value if set, zero value otherwise.
func (o *SearchSearchResponse) GetActualHits() string {
	if o == nil || o.ActualHits == nil {
		var ret string
		return ret
	}
	return *o.ActualHits
}

// GetActualHitsOk returns a tuple with the ActualHits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchSearchResponse) GetActualHitsOk() (*string, bool) {
	if o == nil || o.ActualHits == nil {
		return nil, false
	}
	return o.ActualHits, true
}

// HasActualHits returns a boolean if a field has been set.
func (o *SearchSearchResponse) HasActualHits() bool {
	if o != nil && o.ActualHits != nil {
		return true
	}

	return false
}

// SetActualHits gets a reference to the given string and assigns it to the ActualHits field.
func (o *SearchSearchResponse) SetActualHits(v string) {
	o.ActualHits = &v
}

// GetAggregatedEntries returns the AggregatedEntries field value if set, zero value otherwise.
func (o *SearchSearchResponse) GetAggregatedEntries() SearchTenantAggregation {
	if o == nil || o.AggregatedEntries == nil {
		var ret SearchTenantAggregation
		return ret
	}
	return *o.AggregatedEntries
}

// GetAggregatedEntriesOk returns a tuple with the AggregatedEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchSearchResponse) GetAggregatedEntriesOk() (*SearchTenantAggregation, bool) {
	if o == nil || o.AggregatedEntries == nil {
		return nil, false
	}
	return o.AggregatedEntries, true
}

// HasAggregatedEntries returns a boolean if a field has been set.
func (o *SearchSearchResponse) HasAggregatedEntries() bool {
	if o != nil && o.AggregatedEntries != nil {
		return true
	}

	return false
}

// SetAggregatedEntries gets a reference to the given SearchTenantAggregation and assigns it to the AggregatedEntries field.
func (o *SearchSearchResponse) SetAggregatedEntries(v SearchTenantAggregation) {
	o.AggregatedEntries = &v
}

// GetEntries returns the Entries field value if set, zero value otherwise.
func (o *SearchSearchResponse) GetEntries() []SearchEntry {
	if o == nil || o.Entries == nil {
		var ret []SearchEntry
		return ret
	}
	return *o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchSearchResponse) GetEntriesOk() (*[]SearchEntry, bool) {
	if o == nil || o.Entries == nil {
		return nil, false
	}
	return o.Entries, true
}

// HasEntries returns a boolean if a field has been set.
func (o *SearchSearchResponse) HasEntries() bool {
	if o != nil && o.Entries != nil {
		return true
	}

	return false
}

// SetEntries gets a reference to the given []SearchEntry and assigns it to the Entries field.
func (o *SearchSearchResponse) SetEntries(v []SearchEntry) {
	o.Entries = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *SearchSearchResponse) GetError() SearchError {
	if o == nil || o.Error == nil {
		var ret SearchError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchSearchResponse) GetErrorOk() (*SearchError, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *SearchSearchResponse) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given SearchError and assigns it to the Error field.
func (o *SearchSearchResponse) SetError(v SearchError) {
	o.Error = &v
}

// GetPreviewEntries returns the PreviewEntries field value if set, zero value otherwise.
func (o *SearchSearchResponse) GetPreviewEntries() SearchTenantPreview {
	if o == nil || o.PreviewEntries == nil {
		var ret SearchTenantPreview
		return ret
	}
	return *o.PreviewEntries
}

// GetPreviewEntriesOk returns a tuple with the PreviewEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchSearchResponse) GetPreviewEntriesOk() (*SearchTenantPreview, bool) {
	if o == nil || o.PreviewEntries == nil {
		return nil, false
	}
	return o.PreviewEntries, true
}

// HasPreviewEntries returns a boolean if a field has been set.
func (o *SearchSearchResponse) HasPreviewEntries() bool {
	if o != nil && o.PreviewEntries != nil {
		return true
	}

	return false
}

// SetPreviewEntries gets a reference to the given SearchTenantPreview and assigns it to the PreviewEntries field.
func (o *SearchSearchResponse) SetPreviewEntries(v SearchTenantPreview) {
	o.PreviewEntries = &v
}

// GetTimeTakenMsecs returns the TimeTakenMsecs field value if set, zero value otherwise.
func (o *SearchSearchResponse) GetTimeTakenMsecs() string {
	if o == nil || o.TimeTakenMsecs == nil {
		var ret string
		return ret
	}
	return *o.TimeTakenMsecs
}

// GetTimeTakenMsecsOk returns a tuple with the TimeTakenMsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchSearchResponse) GetTimeTakenMsecsOk() (*string, bool) {
	if o == nil || o.TimeTakenMsecs == nil {
		return nil, false
	}
	return o.TimeTakenMsecs, true
}

// HasTimeTakenMsecs returns a boolean if a field has been set.
func (o *SearchSearchResponse) HasTimeTakenMsecs() bool {
	if o != nil && o.TimeTakenMsecs != nil {
		return true
	}

	return false
}

// SetTimeTakenMsecs gets a reference to the given string and assigns it to the TimeTakenMsecs field.
func (o *SearchSearchResponse) SetTimeTakenMsecs(v string) {
	o.TimeTakenMsecs = &v
}

// GetTotalHits returns the TotalHits field value if set, zero value otherwise.
func (o *SearchSearchResponse) GetTotalHits() string {
	if o == nil || o.TotalHits == nil {
		var ret string
		return ret
	}
	return *o.TotalHits
}

// GetTotalHitsOk returns a tuple with the TotalHits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchSearchResponse) GetTotalHitsOk() (*string, bool) {
	if o == nil || o.TotalHits == nil {
		return nil, false
	}
	return o.TotalHits, true
}

// HasTotalHits returns a boolean if a field has been set.
func (o *SearchSearchResponse) HasTotalHits() bool {
	if o != nil && o.TotalHits != nil {
		return true
	}

	return false
}

// SetTotalHits gets a reference to the given string and assigns it to the TotalHits field.
func (o *SearchSearchResponse) SetTotalHits(v string) {
	o.TotalHits = &v
}

func (o SearchSearchResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ActualHits != nil {
		toSerialize["actual-hits"] = o.ActualHits
	}
	if o.AggregatedEntries != nil {
		toSerialize["aggregated-entries"] = o.AggregatedEntries
	}
	if o.Entries != nil {
		toSerialize["entries"] = o.Entries
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.PreviewEntries != nil {
		toSerialize["preview-entries"] = o.PreviewEntries
	}
	if o.TimeTakenMsecs != nil {
		toSerialize["time-taken-msecs"] = o.TimeTakenMsecs
	}
	if o.TotalHits != nil {
		toSerialize["total-hits"] = o.TotalHits
	}
	return json.Marshal(toSerialize)
}

type NullableSearchSearchResponse struct {
	value *SearchSearchResponse
	isSet bool
}

func (v NullableSearchSearchResponse) Get() *SearchSearchResponse {
	return v.value
}

func (v *NullableSearchSearchResponse) Set(val *SearchSearchResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchSearchResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchSearchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchSearchResponse(val *SearchSearchResponse) *NullableSearchSearchResponse {
	return &NullableSearchSearchResponse{value: val, isSet: true}
}

func (v NullableSearchSearchResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchSearchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


