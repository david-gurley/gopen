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

// PreferencesUIGlobalSettings UIGlobalSettings holds all global UI customization info.
type PreferencesUIGlobalSettings struct {
	ApiVersion *string `json:"api-version,omitempty"`
	Kind *string `json:"kind,omitempty"`
	Meta *ApiObjectMeta `json:"meta,omitempty"`
	Spec *PreferencesUIGlobalSettingsSpec `json:"spec,omitempty"`
	Status *map[string]interface{} `json:"status,omitempty"`
}

// NewPreferencesUIGlobalSettings instantiates a new PreferencesUIGlobalSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreferencesUIGlobalSettings() *PreferencesUIGlobalSettings {
	this := PreferencesUIGlobalSettings{}
	return &this
}

// NewPreferencesUIGlobalSettingsWithDefaults instantiates a new PreferencesUIGlobalSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreferencesUIGlobalSettingsWithDefaults() *PreferencesUIGlobalSettings {
	this := PreferencesUIGlobalSettings{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *PreferencesUIGlobalSettings) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreferencesUIGlobalSettings) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *PreferencesUIGlobalSettings) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *PreferencesUIGlobalSettings) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *PreferencesUIGlobalSettings) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreferencesUIGlobalSettings) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *PreferencesUIGlobalSettings) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *PreferencesUIGlobalSettings) SetKind(v string) {
	o.Kind = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *PreferencesUIGlobalSettings) GetMeta() ApiObjectMeta {
	if o == nil || o.Meta == nil {
		var ret ApiObjectMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreferencesUIGlobalSettings) GetMetaOk() (*ApiObjectMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *PreferencesUIGlobalSettings) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ApiObjectMeta and assigns it to the Meta field.
func (o *PreferencesUIGlobalSettings) SetMeta(v ApiObjectMeta) {
	o.Meta = &v
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *PreferencesUIGlobalSettings) GetSpec() PreferencesUIGlobalSettingsSpec {
	if o == nil || o.Spec == nil {
		var ret PreferencesUIGlobalSettingsSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreferencesUIGlobalSettings) GetSpecOk() (*PreferencesUIGlobalSettingsSpec, bool) {
	if o == nil || o.Spec == nil {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *PreferencesUIGlobalSettings) HasSpec() bool {
	if o != nil && o.Spec != nil {
		return true
	}

	return false
}

// SetSpec gets a reference to the given PreferencesUIGlobalSettingsSpec and assigns it to the Spec field.
func (o *PreferencesUIGlobalSettings) SetSpec(v PreferencesUIGlobalSettingsSpec) {
	o.Spec = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PreferencesUIGlobalSettings) GetStatus() map[string]interface{} {
	if o == nil || o.Status == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreferencesUIGlobalSettings) GetStatusOk() (*map[string]interface{}, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PreferencesUIGlobalSettings) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given map[string]interface{} and assigns it to the Status field.
func (o *PreferencesUIGlobalSettings) SetStatus(v map[string]interface{}) {
	o.Status = &v
}

func (o PreferencesUIGlobalSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiVersion != nil {
		toSerialize["api-version"] = o.ApiVersion
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.Spec != nil {
		toSerialize["spec"] = o.Spec
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullablePreferencesUIGlobalSettings struct {
	value *PreferencesUIGlobalSettings
	isSet bool
}

func (v NullablePreferencesUIGlobalSettings) Get() *PreferencesUIGlobalSettings {
	return v.value
}

func (v *NullablePreferencesUIGlobalSettings) Set(val *PreferencesUIGlobalSettings) {
	v.value = val
	v.isSet = true
}

func (v NullablePreferencesUIGlobalSettings) IsSet() bool {
	return v.isSet
}

func (v *NullablePreferencesUIGlobalSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreferencesUIGlobalSettings(val *PreferencesUIGlobalSettings) *NullablePreferencesUIGlobalSettings {
	return &NullablePreferencesUIGlobalSettings{value: val, isSet: true}
}

func (v NullablePreferencesUIGlobalSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreferencesUIGlobalSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


