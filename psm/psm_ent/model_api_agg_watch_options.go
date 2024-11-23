/*
 * Workload API reference
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

// ApiAggWatchOptions struct for ApiAggWatchOptions
type ApiAggWatchOptions struct {
	Meta *ApiObjectMeta `json:"meta,omitempty"`
	WatchOptions *[]ApiKindWatchOptions `json:"watch-options,omitempty"`
}

// NewApiAggWatchOptions instantiates a new ApiAggWatchOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAggWatchOptions() *ApiAggWatchOptions {
	this := ApiAggWatchOptions{}
	return &this
}

// NewApiAggWatchOptionsWithDefaults instantiates a new ApiAggWatchOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAggWatchOptionsWithDefaults() *ApiAggWatchOptions {
	this := ApiAggWatchOptions{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ApiAggWatchOptions) GetMeta() ApiObjectMeta {
	if o == nil || o.Meta == nil {
		var ret ApiObjectMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAggWatchOptions) GetMetaOk() (*ApiObjectMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ApiAggWatchOptions) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ApiObjectMeta and assigns it to the Meta field.
func (o *ApiAggWatchOptions) SetMeta(v ApiObjectMeta) {
	o.Meta = &v
}

// GetWatchOptions returns the WatchOptions field value if set, zero value otherwise.
func (o *ApiAggWatchOptions) GetWatchOptions() []ApiKindWatchOptions {
	if o == nil || o.WatchOptions == nil {
		var ret []ApiKindWatchOptions
		return ret
	}
	return *o.WatchOptions
}

// GetWatchOptionsOk returns a tuple with the WatchOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAggWatchOptions) GetWatchOptionsOk() (*[]ApiKindWatchOptions, bool) {
	if o == nil || o.WatchOptions == nil {
		return nil, false
	}
	return o.WatchOptions, true
}

// HasWatchOptions returns a boolean if a field has been set.
func (o *ApiAggWatchOptions) HasWatchOptions() bool {
	if o != nil && o.WatchOptions != nil {
		return true
	}

	return false
}

// SetWatchOptions gets a reference to the given []ApiKindWatchOptions and assigns it to the WatchOptions field.
func (o *ApiAggWatchOptions) SetWatchOptions(v []ApiKindWatchOptions) {
	o.WatchOptions = &v
}

func (o ApiAggWatchOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.WatchOptions != nil {
		toSerialize["watch-options"] = o.WatchOptions
	}
	return json.Marshal(toSerialize)
}

type NullableApiAggWatchOptions struct {
	value *ApiAggWatchOptions
	isSet bool
}

func (v NullableApiAggWatchOptions) Get() *ApiAggWatchOptions {
	return v.value
}

func (v *NullableApiAggWatchOptions) Set(val *ApiAggWatchOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAggWatchOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAggWatchOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAggWatchOptions(val *ApiAggWatchOptions) *NullableApiAggWatchOptions {
	return &NullableApiAggWatchOptions{value: val, isSet: true}
}

func (v NullableApiAggWatchOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAggWatchOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


