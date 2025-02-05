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

// SearchPolicySearchResponse PolicySearchResponse is response to the security/firewall policy search request.
type SearchPolicySearchResponse struct {
	Error *SearchError `json:"error,omitempty"`
	// Result is Map of <Kind, Object name, PolicyMatch Entry>.
	Results *map[string]SearchRulesByPolicyName `json:"results,omitempty"`
	// Status of firewall policy search.
	Status *string `json:"status,omitempty"`
}

// NewSearchPolicySearchResponse instantiates a new SearchPolicySearchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchPolicySearchResponse() *SearchPolicySearchResponse {
	this := SearchPolicySearchResponse{}
	var status string = "match"
	this.Status = &status
	return &this
}

// NewSearchPolicySearchResponseWithDefaults instantiates a new SearchPolicySearchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchPolicySearchResponseWithDefaults() *SearchPolicySearchResponse {
	this := SearchPolicySearchResponse{}
	var status string = "match"
	this.Status = &status
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *SearchPolicySearchResponse) GetError() SearchError {
	if o == nil || o.Error == nil {
		var ret SearchError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchPolicySearchResponse) GetErrorOk() (*SearchError, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *SearchPolicySearchResponse) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given SearchError and assigns it to the Error field.
func (o *SearchPolicySearchResponse) SetError(v SearchError) {
	o.Error = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *SearchPolicySearchResponse) GetResults() map[string]SearchRulesByPolicyName {
	if o == nil || o.Results == nil {
		var ret map[string]SearchRulesByPolicyName
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchPolicySearchResponse) GetResultsOk() (*map[string]SearchRulesByPolicyName, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *SearchPolicySearchResponse) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given map[string]SearchRulesByPolicyName and assigns it to the Results field.
func (o *SearchPolicySearchResponse) SetResults(v map[string]SearchRulesByPolicyName) {
	o.Results = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SearchPolicySearchResponse) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchPolicySearchResponse) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SearchPolicySearchResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SearchPolicySearchResponse) SetStatus(v string) {
	o.Status = &v
}

func (o SearchPolicySearchResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableSearchPolicySearchResponse struct {
	value *SearchPolicySearchResponse
	isSet bool
}

func (v NullableSearchPolicySearchResponse) Get() *SearchPolicySearchResponse {
	return v.value
}

func (v *NullableSearchPolicySearchResponse) Set(val *SearchPolicySearchResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchPolicySearchResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchPolicySearchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchPolicySearchResponse(val *SearchPolicySearchResponse) *NullableSearchPolicySearchResponse {
	return &NullableSearchPolicySearchResponse{value: val, isSet: true}
}

func (v NullableSearchPolicySearchResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchPolicySearchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


