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

// SearchSearchQuery struct for SearchSearchQuery
type SearchSearchQuery struct {
	// OR of Categories to be matched, AND and Exclude are not supported for this type The max category string length is 64 bytes. Length of string should be between 0 and 64.
	Categories *[]string `json:"categories,omitempty"`
	Fields *FieldsSelector `json:"fields,omitempty"`
	// OR of Kinds to be matched, AND and Exclude are not supported for this type. Should be a valid object Kind.
	Kinds *[]string `json:"kinds,omitempty"`
	Labels *LabelsSelector `json:"labels,omitempty"`
	// OR of Text-requirements to be matched, Exclude is not supported for Text search.
	Texts *[]SearchTextRequirement `json:"texts,omitempty"`
}

// NewSearchSearchQuery instantiates a new SearchSearchQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchSearchQuery() *SearchSearchQuery {
	this := SearchSearchQuery{}
	return &this
}

// NewSearchSearchQueryWithDefaults instantiates a new SearchSearchQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchSearchQueryWithDefaults() *SearchSearchQuery {
	this := SearchSearchQuery{}
	return &this
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *SearchSearchQuery) GetCategories() []string {
	if o == nil || o.Categories == nil {
		var ret []string
		return ret
	}
	return *o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchSearchQuery) GetCategoriesOk() (*[]string, bool) {
	if o == nil || o.Categories == nil {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *SearchSearchQuery) HasCategories() bool {
	if o != nil && o.Categories != nil {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []string and assigns it to the Categories field.
func (o *SearchSearchQuery) SetCategories(v []string) {
	o.Categories = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *SearchSearchQuery) GetFields() FieldsSelector {
	if o == nil || o.Fields == nil {
		var ret FieldsSelector
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchSearchQuery) GetFieldsOk() (*FieldsSelector, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *SearchSearchQuery) HasFields() bool {
	if o != nil && o.Fields != nil {
		return true
	}

	return false
}

// SetFields gets a reference to the given FieldsSelector and assigns it to the Fields field.
func (o *SearchSearchQuery) SetFields(v FieldsSelector) {
	o.Fields = &v
}

// GetKinds returns the Kinds field value if set, zero value otherwise.
func (o *SearchSearchQuery) GetKinds() []string {
	if o == nil || o.Kinds == nil {
		var ret []string
		return ret
	}
	return *o.Kinds
}

// GetKindsOk returns a tuple with the Kinds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchSearchQuery) GetKindsOk() (*[]string, bool) {
	if o == nil || o.Kinds == nil {
		return nil, false
	}
	return o.Kinds, true
}

// HasKinds returns a boolean if a field has been set.
func (o *SearchSearchQuery) HasKinds() bool {
	if o != nil && o.Kinds != nil {
		return true
	}

	return false
}

// SetKinds gets a reference to the given []string and assigns it to the Kinds field.
func (o *SearchSearchQuery) SetKinds(v []string) {
	o.Kinds = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *SearchSearchQuery) GetLabels() LabelsSelector {
	if o == nil || o.Labels == nil {
		var ret LabelsSelector
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchSearchQuery) GetLabelsOk() (*LabelsSelector, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *SearchSearchQuery) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given LabelsSelector and assigns it to the Labels field.
func (o *SearchSearchQuery) SetLabels(v LabelsSelector) {
	o.Labels = &v
}

// GetTexts returns the Texts field value if set, zero value otherwise.
func (o *SearchSearchQuery) GetTexts() []SearchTextRequirement {
	if o == nil || o.Texts == nil {
		var ret []SearchTextRequirement
		return ret
	}
	return *o.Texts
}

// GetTextsOk returns a tuple with the Texts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchSearchQuery) GetTextsOk() (*[]SearchTextRequirement, bool) {
	if o == nil || o.Texts == nil {
		return nil, false
	}
	return o.Texts, true
}

// HasTexts returns a boolean if a field has been set.
func (o *SearchSearchQuery) HasTexts() bool {
	if o != nil && o.Texts != nil {
		return true
	}

	return false
}

// SetTexts gets a reference to the given []SearchTextRequirement and assigns it to the Texts field.
func (o *SearchSearchQuery) SetTexts(v []SearchTextRequirement) {
	o.Texts = &v
}

func (o SearchSearchQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Categories != nil {
		toSerialize["categories"] = o.Categories
	}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	if o.Kinds != nil {
		toSerialize["kinds"] = o.Kinds
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if o.Texts != nil {
		toSerialize["texts"] = o.Texts
	}
	return json.Marshal(toSerialize)
}

type NullableSearchSearchQuery struct {
	value *SearchSearchQuery
	isSet bool
}

func (v NullableSearchSearchQuery) Get() *SearchSearchQuery {
	return v.value
}

func (v *NullableSearchSearchQuery) Set(val *SearchSearchQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchSearchQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchSearchQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchSearchQuery(val *SearchSearchQuery) *NullableSearchSearchQuery {
	return &NullableSearchSearchQuery{value: val, isSet: true}
}

func (v NullableSearchSearchQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchSearchQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


