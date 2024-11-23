/*
 * Preferences API reference
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

// PreferencesUIGlobalSettingsSpec struct for PreferencesUIGlobalSettingsSpec
type PreferencesUIGlobalSettingsSpec struct {
	IdleTimeout *PreferencesIdleTimeout `json:"idle-timeout,omitempty"`
	// Can contain any UI style preferences. Provide typing through UI code.
	StyleOptions *string `json:"style-options,omitempty"`
}

// NewPreferencesUIGlobalSettingsSpec instantiates a new PreferencesUIGlobalSettingsSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreferencesUIGlobalSettingsSpec() *PreferencesUIGlobalSettingsSpec {
	this := PreferencesUIGlobalSettingsSpec{}
	return &this
}

// NewPreferencesUIGlobalSettingsSpecWithDefaults instantiates a new PreferencesUIGlobalSettingsSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreferencesUIGlobalSettingsSpecWithDefaults() *PreferencesUIGlobalSettingsSpec {
	this := PreferencesUIGlobalSettingsSpec{}
	return &this
}

// GetIdleTimeout returns the IdleTimeout field value if set, zero value otherwise.
func (o *PreferencesUIGlobalSettingsSpec) GetIdleTimeout() PreferencesIdleTimeout {
	if o == nil || o.IdleTimeout == nil {
		var ret PreferencesIdleTimeout
		return ret
	}
	return *o.IdleTimeout
}

// GetIdleTimeoutOk returns a tuple with the IdleTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreferencesUIGlobalSettingsSpec) GetIdleTimeoutOk() (*PreferencesIdleTimeout, bool) {
	if o == nil || o.IdleTimeout == nil {
		return nil, false
	}
	return o.IdleTimeout, true
}

// HasIdleTimeout returns a boolean if a field has been set.
func (o *PreferencesUIGlobalSettingsSpec) HasIdleTimeout() bool {
	if o != nil && o.IdleTimeout != nil {
		return true
	}

	return false
}

// SetIdleTimeout gets a reference to the given PreferencesIdleTimeout and assigns it to the IdleTimeout field.
func (o *PreferencesUIGlobalSettingsSpec) SetIdleTimeout(v PreferencesIdleTimeout) {
	o.IdleTimeout = &v
}

// GetStyleOptions returns the StyleOptions field value if set, zero value otherwise.
func (o *PreferencesUIGlobalSettingsSpec) GetStyleOptions() string {
	if o == nil || o.StyleOptions == nil {
		var ret string
		return ret
	}
	return *o.StyleOptions
}

// GetStyleOptionsOk returns a tuple with the StyleOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreferencesUIGlobalSettingsSpec) GetStyleOptionsOk() (*string, bool) {
	if o == nil || o.StyleOptions == nil {
		return nil, false
	}
	return o.StyleOptions, true
}

// HasStyleOptions returns a boolean if a field has been set.
func (o *PreferencesUIGlobalSettingsSpec) HasStyleOptions() bool {
	if o != nil && o.StyleOptions != nil {
		return true
	}

	return false
}

// SetStyleOptions gets a reference to the given string and assigns it to the StyleOptions field.
func (o *PreferencesUIGlobalSettingsSpec) SetStyleOptions(v string) {
	o.StyleOptions = &v
}

func (o PreferencesUIGlobalSettingsSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IdleTimeout != nil {
		toSerialize["idle-timeout"] = o.IdleTimeout
	}
	if o.StyleOptions != nil {
		toSerialize["style-options"] = o.StyleOptions
	}
	return json.Marshal(toSerialize)
}

type NullablePreferencesUIGlobalSettingsSpec struct {
	value *PreferencesUIGlobalSettingsSpec
	isSet bool
}

func (v NullablePreferencesUIGlobalSettingsSpec) Get() *PreferencesUIGlobalSettingsSpec {
	return v.value
}

func (v *NullablePreferencesUIGlobalSettingsSpec) Set(val *PreferencesUIGlobalSettingsSpec) {
	v.value = val
	v.isSet = true
}

func (v NullablePreferencesUIGlobalSettingsSpec) IsSet() bool {
	return v.isSet
}

func (v *NullablePreferencesUIGlobalSettingsSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreferencesUIGlobalSettingsSpec(val *PreferencesUIGlobalSettingsSpec) *NullablePreferencesUIGlobalSettingsSpec {
	return &NullablePreferencesUIGlobalSettingsSpec{value: val, isSet: true}
}

func (v NullablePreferencesUIGlobalSettingsSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreferencesUIGlobalSettingsSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


