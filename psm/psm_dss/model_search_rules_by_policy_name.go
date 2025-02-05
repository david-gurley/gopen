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

// SearchRulesByPolicyName RulesByPolicyName map of object/policy names and the matching rules.
type SearchRulesByPolicyName struct {
	Policies *map[string]SearchPolicyMatchEntries `json:"policies,omitempty"`
}

// NewSearchRulesByPolicyName instantiates a new SearchRulesByPolicyName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchRulesByPolicyName() *SearchRulesByPolicyName {
	this := SearchRulesByPolicyName{}
	return &this
}

// NewSearchRulesByPolicyNameWithDefaults instantiates a new SearchRulesByPolicyName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchRulesByPolicyNameWithDefaults() *SearchRulesByPolicyName {
	this := SearchRulesByPolicyName{}
	return &this
}

// GetPolicies returns the Policies field value if set, zero value otherwise.
func (o *SearchRulesByPolicyName) GetPolicies() map[string]SearchPolicyMatchEntries {
	if o == nil || o.Policies == nil {
		var ret map[string]SearchPolicyMatchEntries
		return ret
	}
	return *o.Policies
}

// GetPoliciesOk returns a tuple with the Policies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchRulesByPolicyName) GetPoliciesOk() (*map[string]SearchPolicyMatchEntries, bool) {
	if o == nil || o.Policies == nil {
		return nil, false
	}
	return o.Policies, true
}

// HasPolicies returns a boolean if a field has been set.
func (o *SearchRulesByPolicyName) HasPolicies() bool {
	if o != nil && o.Policies != nil {
		return true
	}

	return false
}

// SetPolicies gets a reference to the given map[string]SearchPolicyMatchEntries and assigns it to the Policies field.
func (o *SearchRulesByPolicyName) SetPolicies(v map[string]SearchPolicyMatchEntries) {
	o.Policies = &v
}

func (o SearchRulesByPolicyName) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Policies != nil {
		toSerialize["policies"] = o.Policies
	}
	return json.Marshal(toSerialize)
}

type NullableSearchRulesByPolicyName struct {
	value *SearchRulesByPolicyName
	isSet bool
}

func (v NullableSearchRulesByPolicyName) Get() *SearchRulesByPolicyName {
	return v.value
}

func (v *NullableSearchRulesByPolicyName) Set(val *SearchRulesByPolicyName) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchRulesByPolicyName) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchRulesByPolicyName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchRulesByPolicyName(val *SearchRulesByPolicyName) *NullableSearchRulesByPolicyName {
	return &NullableSearchRulesByPolicyName{value: val, isSet: true}
}

func (v NullableSearchRulesByPolicyName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchRulesByPolicyName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


