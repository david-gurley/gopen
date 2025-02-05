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

// SearchKindAggregation KindAggregation contains map of search result entries grouped by Kind.
type SearchKindAggregation struct {
	Kinds *map[string]SearchEntryList `json:"kinds,omitempty"`
}

// NewSearchKindAggregation instantiates a new SearchKindAggregation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchKindAggregation() *SearchKindAggregation {
	this := SearchKindAggregation{}
	return &this
}

// NewSearchKindAggregationWithDefaults instantiates a new SearchKindAggregation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchKindAggregationWithDefaults() *SearchKindAggregation {
	this := SearchKindAggregation{}
	return &this
}

// GetKinds returns the Kinds field value if set, zero value otherwise.
func (o *SearchKindAggregation) GetKinds() map[string]SearchEntryList {
	if o == nil || o.Kinds == nil {
		var ret map[string]SearchEntryList
		return ret
	}
	return *o.Kinds
}

// GetKindsOk returns a tuple with the Kinds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchKindAggregation) GetKindsOk() (*map[string]SearchEntryList, bool) {
	if o == nil || o.Kinds == nil {
		return nil, false
	}
	return o.Kinds, true
}

// HasKinds returns a boolean if a field has been set.
func (o *SearchKindAggregation) HasKinds() bool {
	if o != nil && o.Kinds != nil {
		return true
	}

	return false
}

// SetKinds gets a reference to the given map[string]SearchEntryList and assigns it to the Kinds field.
func (o *SearchKindAggregation) SetKinds(v map[string]SearchEntryList) {
	o.Kinds = &v
}

func (o SearchKindAggregation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Kinds != nil {
		toSerialize["kinds"] = o.Kinds
	}
	return json.Marshal(toSerialize)
}

type NullableSearchKindAggregation struct {
	value *SearchKindAggregation
	isSet bool
}

func (v NullableSearchKindAggregation) Get() *SearchKindAggregation {
	return v.value
}

func (v *NullableSearchKindAggregation) Set(val *SearchKindAggregation) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchKindAggregation) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchKindAggregation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchKindAggregation(val *SearchKindAggregation) *NullableSearchKindAggregation {
	return &NullableSearchKindAggregation{value: val, isSet: true}
}

func (v NullableSearchKindAggregation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchKindAggregation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


