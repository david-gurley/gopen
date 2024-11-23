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

// ClusterAutoMsgDistributedServiceCardWatchHelper AutoMsgDistributedServiceCardWatchHelper is a wrapper object for watch events for DistributedServiceCard objects.
type ClusterAutoMsgDistributedServiceCardWatchHelper struct {
	Events *[]ClusterAutoMsgDistributedServiceCardWatchHelperWatchEvent `json:"events,omitempty"`
}

// NewClusterAutoMsgDistributedServiceCardWatchHelper instantiates a new ClusterAutoMsgDistributedServiceCardWatchHelper object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterAutoMsgDistributedServiceCardWatchHelper() *ClusterAutoMsgDistributedServiceCardWatchHelper {
	this := ClusterAutoMsgDistributedServiceCardWatchHelper{}
	return &this
}

// NewClusterAutoMsgDistributedServiceCardWatchHelperWithDefaults instantiates a new ClusterAutoMsgDistributedServiceCardWatchHelper object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterAutoMsgDistributedServiceCardWatchHelperWithDefaults() *ClusterAutoMsgDistributedServiceCardWatchHelper {
	this := ClusterAutoMsgDistributedServiceCardWatchHelper{}
	return &this
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *ClusterAutoMsgDistributedServiceCardWatchHelper) GetEvents() []ClusterAutoMsgDistributedServiceCardWatchHelperWatchEvent {
	if o == nil || o.Events == nil {
		var ret []ClusterAutoMsgDistributedServiceCardWatchHelperWatchEvent
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAutoMsgDistributedServiceCardWatchHelper) GetEventsOk() (*[]ClusterAutoMsgDistributedServiceCardWatchHelperWatchEvent, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *ClusterAutoMsgDistributedServiceCardWatchHelper) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []ClusterAutoMsgDistributedServiceCardWatchHelperWatchEvent and assigns it to the Events field.
func (o *ClusterAutoMsgDistributedServiceCardWatchHelper) SetEvents(v []ClusterAutoMsgDistributedServiceCardWatchHelperWatchEvent) {
	o.Events = &v
}

func (o ClusterAutoMsgDistributedServiceCardWatchHelper) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableClusterAutoMsgDistributedServiceCardWatchHelper struct {
	value *ClusterAutoMsgDistributedServiceCardWatchHelper
	isSet bool
}

func (v NullableClusterAutoMsgDistributedServiceCardWatchHelper) Get() *ClusterAutoMsgDistributedServiceCardWatchHelper {
	return v.value
}

func (v *NullableClusterAutoMsgDistributedServiceCardWatchHelper) Set(val *ClusterAutoMsgDistributedServiceCardWatchHelper) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterAutoMsgDistributedServiceCardWatchHelper) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterAutoMsgDistributedServiceCardWatchHelper) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterAutoMsgDistributedServiceCardWatchHelper(val *ClusterAutoMsgDistributedServiceCardWatchHelper) *NullableClusterAutoMsgDistributedServiceCardWatchHelper {
	return &NullableClusterAutoMsgDistributedServiceCardWatchHelper{value: val, isSet: true}
}

func (v NullableClusterAutoMsgDistributedServiceCardWatchHelper) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterAutoMsgDistributedServiceCardWatchHelper) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


