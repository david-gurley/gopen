/*
 * Monitoring API reference
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

// MonitoringInstanceSelector Helps to select/filter instances of a kind. If instance selector is not given, then the policy will be applied across all DSCs or Nodes based on the metric identifier. If kind was given, then either names or label or both have to be specified. If the user specified both names and label, policies will be applied on the union of both. Allowed list of kinds: Node, DistributedServiceCard.
type MonitoringInstanceSelector struct {
	// Kind of the instances to be selected using names/label.
	Kind *string `json:"kind,omitempty"`
	Labels *LabelsSelector `json:"labels,omitempty"`
	// List of names/reporter IDs.
	Names *[]string `json:"names,omitempty"`
}

// NewMonitoringInstanceSelector instantiates a new MonitoringInstanceSelector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitoringInstanceSelector() *MonitoringInstanceSelector {
	this := MonitoringInstanceSelector{}
	return &this
}

// NewMonitoringInstanceSelectorWithDefaults instantiates a new MonitoringInstanceSelector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitoringInstanceSelectorWithDefaults() *MonitoringInstanceSelector {
	this := MonitoringInstanceSelector{}
	return &this
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *MonitoringInstanceSelector) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringInstanceSelector) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *MonitoringInstanceSelector) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *MonitoringInstanceSelector) SetKind(v string) {
	o.Kind = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *MonitoringInstanceSelector) GetLabels() LabelsSelector {
	if o == nil || o.Labels == nil {
		var ret LabelsSelector
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringInstanceSelector) GetLabelsOk() (*LabelsSelector, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *MonitoringInstanceSelector) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given LabelsSelector and assigns it to the Labels field.
func (o *MonitoringInstanceSelector) SetLabels(v LabelsSelector) {
	o.Labels = &v
}

// GetNames returns the Names field value if set, zero value otherwise.
func (o *MonitoringInstanceSelector) GetNames() []string {
	if o == nil || o.Names == nil {
		var ret []string
		return ret
	}
	return *o.Names
}

// GetNamesOk returns a tuple with the Names field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringInstanceSelector) GetNamesOk() (*[]string, bool) {
	if o == nil || o.Names == nil {
		return nil, false
	}
	return o.Names, true
}

// HasNames returns a boolean if a field has been set.
func (o *MonitoringInstanceSelector) HasNames() bool {
	if o != nil && o.Names != nil {
		return true
	}

	return false
}

// SetNames gets a reference to the given []string and assigns it to the Names field.
func (o *MonitoringInstanceSelector) SetNames(v []string) {
	o.Names = &v
}

func (o MonitoringInstanceSelector) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if o.Names != nil {
		toSerialize["names"] = o.Names
	}
	return json.Marshal(toSerialize)
}

type NullableMonitoringInstanceSelector struct {
	value *MonitoringInstanceSelector
	isSet bool
}

func (v NullableMonitoringInstanceSelector) Get() *MonitoringInstanceSelector {
	return v.value
}

func (v *NullableMonitoringInstanceSelector) Set(val *MonitoringInstanceSelector) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringInstanceSelector) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringInstanceSelector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringInstanceSelector(val *MonitoringInstanceSelector) *NullableMonitoringInstanceSelector {
	return &NullableMonitoringInstanceSelector{value: val, isSet: true}
}

func (v NullableMonitoringInstanceSelector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringInstanceSelector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


