/*
 * Cluster API reference
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

// ConfigurationSnapshotStatusConfigSaveStatus struct for ConfigurationSnapshotStatusConfigSaveStatus
type ConfigurationSnapshotStatusConfigSaveStatus struct {
	DestType *string `json:"dest-type,omitempty"`
	Uri *string `json:"uri,omitempty"`
}

// NewConfigurationSnapshotStatusConfigSaveStatus instantiates a new ConfigurationSnapshotStatusConfigSaveStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigurationSnapshotStatusConfigSaveStatus() *ConfigurationSnapshotStatusConfigSaveStatus {
	this := ConfigurationSnapshotStatusConfigSaveStatus{}
	var destType string = "objectstore"
	this.DestType = &destType
	return &this
}

// NewConfigurationSnapshotStatusConfigSaveStatusWithDefaults instantiates a new ConfigurationSnapshotStatusConfigSaveStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigurationSnapshotStatusConfigSaveStatusWithDefaults() *ConfigurationSnapshotStatusConfigSaveStatus {
	this := ConfigurationSnapshotStatusConfigSaveStatus{}
	var destType string = "objectstore"
	this.DestType = &destType
	return &this
}

// GetDestType returns the DestType field value if set, zero value otherwise.
func (o *ConfigurationSnapshotStatusConfigSaveStatus) GetDestType() string {
	if o == nil || o.DestType == nil {
		var ret string
		return ret
	}
	return *o.DestType
}

// GetDestTypeOk returns a tuple with the DestType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationSnapshotStatusConfigSaveStatus) GetDestTypeOk() (*string, bool) {
	if o == nil || o.DestType == nil {
		return nil, false
	}
	return o.DestType, true
}

// HasDestType returns a boolean if a field has been set.
func (o *ConfigurationSnapshotStatusConfigSaveStatus) HasDestType() bool {
	if o != nil && o.DestType != nil {
		return true
	}

	return false
}

// SetDestType gets a reference to the given string and assigns it to the DestType field.
func (o *ConfigurationSnapshotStatusConfigSaveStatus) SetDestType(v string) {
	o.DestType = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *ConfigurationSnapshotStatusConfigSaveStatus) GetUri() string {
	if o == nil || o.Uri == nil {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationSnapshotStatusConfigSaveStatus) GetUriOk() (*string, bool) {
	if o == nil || o.Uri == nil {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *ConfigurationSnapshotStatusConfigSaveStatus) HasUri() bool {
	if o != nil && o.Uri != nil {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *ConfigurationSnapshotStatusConfigSaveStatus) SetUri(v string) {
	o.Uri = &v
}

func (o ConfigurationSnapshotStatusConfigSaveStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DestType != nil {
		toSerialize["dest-type"] = o.DestType
	}
	if o.Uri != nil {
		toSerialize["uri"] = o.Uri
	}
	return json.Marshal(toSerialize)
}

type NullableConfigurationSnapshotStatusConfigSaveStatus struct {
	value *ConfigurationSnapshotStatusConfigSaveStatus
	isSet bool
}

func (v NullableConfigurationSnapshotStatusConfigSaveStatus) Get() *ConfigurationSnapshotStatusConfigSaveStatus {
	return v.value
}

func (v *NullableConfigurationSnapshotStatusConfigSaveStatus) Set(val *ConfigurationSnapshotStatusConfigSaveStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigurationSnapshotStatusConfigSaveStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigurationSnapshotStatusConfigSaveStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigurationSnapshotStatusConfigSaveStatus(val *ConfigurationSnapshotStatusConfigSaveStatus) *NullableConfigurationSnapshotStatusConfigSaveStatus {
	return &NullableConfigurationSnapshotStatusConfigSaveStatus{value: val, isSet: true}
}

func (v NullableConfigurationSnapshotStatusConfigSaveStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigurationSnapshotStatusConfigSaveStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


