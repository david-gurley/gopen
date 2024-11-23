/*
 * Events API reference
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

// EventsEventSource EventSource has info about the component and host/node that generated the event.
type EventsEventSource struct {
	// Component from which the event is generated.
	Component *string `json:"component,omitempty"`
	// Name of the venice or workload node which is generating the event.
	NodeName *string `json:"node-name,omitempty"`
	// Unit ID on the host which is generating the event.
	UnitId *string `json:"unit-id,omitempty"`
}

// NewEventsEventSource instantiates a new EventsEventSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventsEventSource() *EventsEventSource {
	this := EventsEventSource{}
	return &this
}

// NewEventsEventSourceWithDefaults instantiates a new EventsEventSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventsEventSourceWithDefaults() *EventsEventSource {
	this := EventsEventSource{}
	return &this
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *EventsEventSource) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsEventSource) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *EventsEventSource) HasComponent() bool {
	if o != nil && o.Component != nil {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *EventsEventSource) SetComponent(v string) {
	o.Component = &v
}

// GetNodeName returns the NodeName field value if set, zero value otherwise.
func (o *EventsEventSource) GetNodeName() string {
	if o == nil || o.NodeName == nil {
		var ret string
		return ret
	}
	return *o.NodeName
}

// GetNodeNameOk returns a tuple with the NodeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsEventSource) GetNodeNameOk() (*string, bool) {
	if o == nil || o.NodeName == nil {
		return nil, false
	}
	return o.NodeName, true
}

// HasNodeName returns a boolean if a field has been set.
func (o *EventsEventSource) HasNodeName() bool {
	if o != nil && o.NodeName != nil {
		return true
	}

	return false
}

// SetNodeName gets a reference to the given string and assigns it to the NodeName field.
func (o *EventsEventSource) SetNodeName(v string) {
	o.NodeName = &v
}

// GetUnitId returns the UnitId field value if set, zero value otherwise.
func (o *EventsEventSource) GetUnitId() string {
	if o == nil || o.UnitId == nil {
		var ret string
		return ret
	}
	return *o.UnitId
}

// GetUnitIdOk returns a tuple with the UnitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventsEventSource) GetUnitIdOk() (*string, bool) {
	if o == nil || o.UnitId == nil {
		return nil, false
	}
	return o.UnitId, true
}

// HasUnitId returns a boolean if a field has been set.
func (o *EventsEventSource) HasUnitId() bool {
	if o != nil && o.UnitId != nil {
		return true
	}

	return false
}

// SetUnitId gets a reference to the given string and assigns it to the UnitId field.
func (o *EventsEventSource) SetUnitId(v string) {
	o.UnitId = &v
}

func (o EventsEventSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	if o.NodeName != nil {
		toSerialize["node-name"] = o.NodeName
	}
	if o.UnitId != nil {
		toSerialize["unit-id"] = o.UnitId
	}
	return json.Marshal(toSerialize)
}

type NullableEventsEventSource struct {
	value *EventsEventSource
	isSet bool
}

func (v NullableEventsEventSource) Get() *EventsEventSource {
	return v.value
}

func (v *NullableEventsEventSource) Set(val *EventsEventSource) {
	v.value = val
	v.isSet = true
}

func (v NullableEventsEventSource) IsSet() bool {
	return v.isSet
}

func (v *NullableEventsEventSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventsEventSource(val *EventsEventSource) *NullableEventsEventSource {
	return &NullableEventsEventSource{value: val, isSet: true}
}

func (v NullableEventsEventSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventsEventSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


