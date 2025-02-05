/*
 * Search API reference
 *
 * Service name  
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package psm_dss

import (
	"encoding/json"
)

// SearchCategoryAggregation CategoryAggregation contains map of search result entries grouped by two levels: first by Category and then by Kind.
type SearchCategoryAggregation struct {
	Categories *map[string]SearchKindAggregation `json:"categories,omitempty"`
}

// NewSearchCategoryAggregation instantiates a new SearchCategoryAggregation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchCategoryAggregation() *SearchCategoryAggregation {
	this := SearchCategoryAggregation{}
	return &this
}

// NewSearchCategoryAggregationWithDefaults instantiates a new SearchCategoryAggregation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchCategoryAggregationWithDefaults() *SearchCategoryAggregation {
	this := SearchCategoryAggregation{}
	return &this
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *SearchCategoryAggregation) GetCategories() map[string]SearchKindAggregation {
	if o == nil || o.Categories == nil {
		var ret map[string]SearchKindAggregation
		return ret
	}
	return *o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCategoryAggregation) GetCategoriesOk() (*map[string]SearchKindAggregation, bool) {
	if o == nil || o.Categories == nil {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *SearchCategoryAggregation) HasCategories() bool {
	if o != nil && o.Categories != nil {
		return true
	}

	return false
}

// SetCategories gets a reference to the given map[string]SearchKindAggregation and assigns it to the Categories field.
func (o *SearchCategoryAggregation) SetCategories(v map[string]SearchKindAggregation) {
	o.Categories = &v
}

func (o SearchCategoryAggregation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Categories != nil {
		toSerialize["categories"] = o.Categories
	}
	return json.Marshal(toSerialize)
}

type NullableSearchCategoryAggregation struct {
	value *SearchCategoryAggregation
	isSet bool
}

func (v NullableSearchCategoryAggregation) Get() *SearchCategoryAggregation {
	return v.value
}

func (v *NullableSearchCategoryAggregation) Set(val *SearchCategoryAggregation) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchCategoryAggregation) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchCategoryAggregation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchCategoryAggregation(val *SearchCategoryAggregation) *NullableSearchCategoryAggregation {
	return &NullableSearchCategoryAggregation{value: val, isSet: true}
}

func (v NullableSearchCategoryAggregation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchCategoryAggregation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


