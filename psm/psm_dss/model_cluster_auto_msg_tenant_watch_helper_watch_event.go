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

// ClusterAutoMsgTenantWatchHelperWatchEvent struct for ClusterAutoMsgTenantWatchHelperWatchEvent
type ClusterAutoMsgTenantWatchHelperWatchEvent struct {
	Object *ClusterTenant `json:"object,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewClusterAutoMsgTenantWatchHelperWatchEvent instantiates a new ClusterAutoMsgTenantWatchHelperWatchEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterAutoMsgTenantWatchHelperWatchEvent() *ClusterAutoMsgTenantWatchHelperWatchEvent {
	this := ClusterAutoMsgTenantWatchHelperWatchEvent{}
	return &this
}

// NewClusterAutoMsgTenantWatchHelperWatchEventWithDefaults instantiates a new ClusterAutoMsgTenantWatchHelperWatchEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterAutoMsgTenantWatchHelperWatchEventWithDefaults() *ClusterAutoMsgTenantWatchHelperWatchEvent {
	this := ClusterAutoMsgTenantWatchHelperWatchEvent{}
	return &this
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *ClusterAutoMsgTenantWatchHelperWatchEvent) GetObject() ClusterTenant {
	if o == nil || o.Object == nil {
		var ret ClusterTenant
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAutoMsgTenantWatchHelperWatchEvent) GetObjectOk() (*ClusterTenant, bool) {
	if o == nil || o.Object == nil {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *ClusterAutoMsgTenantWatchHelperWatchEvent) HasObject() bool {
	if o != nil && o.Object != nil {
		return true
	}

	return false
}

// SetObject gets a reference to the given ClusterTenant and assigns it to the Object field.
func (o *ClusterAutoMsgTenantWatchHelperWatchEvent) SetObject(v ClusterTenant) {
	o.Object = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ClusterAutoMsgTenantWatchHelperWatchEvent) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAutoMsgTenantWatchHelperWatchEvent) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ClusterAutoMsgTenantWatchHelperWatchEvent) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ClusterAutoMsgTenantWatchHelperWatchEvent) SetType(v string) {
	o.Type = &v
}

func (o ClusterAutoMsgTenantWatchHelperWatchEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Object != nil {
		toSerialize["object"] = o.Object
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableClusterAutoMsgTenantWatchHelperWatchEvent struct {
	value *ClusterAutoMsgTenantWatchHelperWatchEvent
	isSet bool
}

func (v NullableClusterAutoMsgTenantWatchHelperWatchEvent) Get() *ClusterAutoMsgTenantWatchHelperWatchEvent {
	return v.value
}

func (v *NullableClusterAutoMsgTenantWatchHelperWatchEvent) Set(val *ClusterAutoMsgTenantWatchHelperWatchEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterAutoMsgTenantWatchHelperWatchEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterAutoMsgTenantWatchHelperWatchEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterAutoMsgTenantWatchHelperWatchEvent(val *ClusterAutoMsgTenantWatchHelperWatchEvent) *NullableClusterAutoMsgTenantWatchHelperWatchEvent {
	return &NullableClusterAutoMsgTenantWatchHelperWatchEvent{value: val, isSet: true}
}

func (v NullableClusterAutoMsgTenantWatchHelperWatchEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterAutoMsgTenantWatchHelperWatchEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


