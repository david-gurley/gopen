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

// ClusterAutoMsgCredentialsWatchHelperWatchEvent struct for ClusterAutoMsgCredentialsWatchHelperWatchEvent
type ClusterAutoMsgCredentialsWatchHelperWatchEvent struct {
	Object *ClusterCredentials `json:"object,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewClusterAutoMsgCredentialsWatchHelperWatchEvent instantiates a new ClusterAutoMsgCredentialsWatchHelperWatchEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterAutoMsgCredentialsWatchHelperWatchEvent() *ClusterAutoMsgCredentialsWatchHelperWatchEvent {
	this := ClusterAutoMsgCredentialsWatchHelperWatchEvent{}
	return &this
}

// NewClusterAutoMsgCredentialsWatchHelperWatchEventWithDefaults instantiates a new ClusterAutoMsgCredentialsWatchHelperWatchEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterAutoMsgCredentialsWatchHelperWatchEventWithDefaults() *ClusterAutoMsgCredentialsWatchHelperWatchEvent {
	this := ClusterAutoMsgCredentialsWatchHelperWatchEvent{}
	return &this
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *ClusterAutoMsgCredentialsWatchHelperWatchEvent) GetObject() ClusterCredentials {
	if o == nil || o.Object == nil {
		var ret ClusterCredentials
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAutoMsgCredentialsWatchHelperWatchEvent) GetObjectOk() (*ClusterCredentials, bool) {
	if o == nil || o.Object == nil {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *ClusterAutoMsgCredentialsWatchHelperWatchEvent) HasObject() bool {
	if o != nil && o.Object != nil {
		return true
	}

	return false
}

// SetObject gets a reference to the given ClusterCredentials and assigns it to the Object field.
func (o *ClusterAutoMsgCredentialsWatchHelperWatchEvent) SetObject(v ClusterCredentials) {
	o.Object = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ClusterAutoMsgCredentialsWatchHelperWatchEvent) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAutoMsgCredentialsWatchHelperWatchEvent) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ClusterAutoMsgCredentialsWatchHelperWatchEvent) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ClusterAutoMsgCredentialsWatchHelperWatchEvent) SetType(v string) {
	o.Type = &v
}

func (o ClusterAutoMsgCredentialsWatchHelperWatchEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Object != nil {
		toSerialize["object"] = o.Object
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableClusterAutoMsgCredentialsWatchHelperWatchEvent struct {
	value *ClusterAutoMsgCredentialsWatchHelperWatchEvent
	isSet bool
}

func (v NullableClusterAutoMsgCredentialsWatchHelperWatchEvent) Get() *ClusterAutoMsgCredentialsWatchHelperWatchEvent {
	return v.value
}

func (v *NullableClusterAutoMsgCredentialsWatchHelperWatchEvent) Set(val *ClusterAutoMsgCredentialsWatchHelperWatchEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterAutoMsgCredentialsWatchHelperWatchEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterAutoMsgCredentialsWatchHelperWatchEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterAutoMsgCredentialsWatchHelperWatchEvent(val *ClusterAutoMsgCredentialsWatchHelperWatchEvent) *NullableClusterAutoMsgCredentialsWatchHelperWatchEvent {
	return &NullableClusterAutoMsgCredentialsWatchHelperWatchEvent{value: val, isSet: true}
}

func (v NullableClusterAutoMsgCredentialsWatchHelperWatchEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterAutoMsgCredentialsWatchHelperWatchEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

