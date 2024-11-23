/*
 * Monitoring API reference
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

// MonitoringAlertReason AlertReason captures all the requirements with matched value from the alert policy rule at the time of creating an alert. e.g. \"matched-requirements\": [{\"field\": \"cpu\", \"operator\": \"Gt\", \"values\": [90], \"observed-value\": 95}].
type MonitoringAlertReason struct {
	// Alert Policy ID that matched.
	AlertPolicyId *string `json:"alert-policy-id,omitempty"`
	// List of requirements from the alert policy with it's matched value.
	MatchedRequirements *[]MonitoringMatchedRequirement `json:"matched-requirements,omitempty"`
}

// NewMonitoringAlertReason instantiates a new MonitoringAlertReason object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitoringAlertReason() *MonitoringAlertReason {
	this := MonitoringAlertReason{}
	return &this
}

// NewMonitoringAlertReasonWithDefaults instantiates a new MonitoringAlertReason object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitoringAlertReasonWithDefaults() *MonitoringAlertReason {
	this := MonitoringAlertReason{}
	return &this
}

// GetAlertPolicyId returns the AlertPolicyId field value if set, zero value otherwise.
func (o *MonitoringAlertReason) GetAlertPolicyId() string {
	if o == nil || o.AlertPolicyId == nil {
		var ret string
		return ret
	}
	return *o.AlertPolicyId
}

// GetAlertPolicyIdOk returns a tuple with the AlertPolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringAlertReason) GetAlertPolicyIdOk() (*string, bool) {
	if o == nil || o.AlertPolicyId == nil {
		return nil, false
	}
	return o.AlertPolicyId, true
}

// HasAlertPolicyId returns a boolean if a field has been set.
func (o *MonitoringAlertReason) HasAlertPolicyId() bool {
	if o != nil && o.AlertPolicyId != nil {
		return true
	}

	return false
}

// SetAlertPolicyId gets a reference to the given string and assigns it to the AlertPolicyId field.
func (o *MonitoringAlertReason) SetAlertPolicyId(v string) {
	o.AlertPolicyId = &v
}

// GetMatchedRequirements returns the MatchedRequirements field value if set, zero value otherwise.
func (o *MonitoringAlertReason) GetMatchedRequirements() []MonitoringMatchedRequirement {
	if o == nil || o.MatchedRequirements == nil {
		var ret []MonitoringMatchedRequirement
		return ret
	}
	return *o.MatchedRequirements
}

// GetMatchedRequirementsOk returns a tuple with the MatchedRequirements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringAlertReason) GetMatchedRequirementsOk() (*[]MonitoringMatchedRequirement, bool) {
	if o == nil || o.MatchedRequirements == nil {
		return nil, false
	}
	return o.MatchedRequirements, true
}

// HasMatchedRequirements returns a boolean if a field has been set.
func (o *MonitoringAlertReason) HasMatchedRequirements() bool {
	if o != nil && o.MatchedRequirements != nil {
		return true
	}

	return false
}

// SetMatchedRequirements gets a reference to the given []MonitoringMatchedRequirement and assigns it to the MatchedRequirements field.
func (o *MonitoringAlertReason) SetMatchedRequirements(v []MonitoringMatchedRequirement) {
	o.MatchedRequirements = &v
}

func (o MonitoringAlertReason) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AlertPolicyId != nil {
		toSerialize["alert-policy-id"] = o.AlertPolicyId
	}
	if o.MatchedRequirements != nil {
		toSerialize["matched-requirements"] = o.MatchedRequirements
	}
	return json.Marshal(toSerialize)
}

type NullableMonitoringAlertReason struct {
	value *MonitoringAlertReason
	isSet bool
}

func (v NullableMonitoringAlertReason) Get() *MonitoringAlertReason {
	return v.value
}

func (v *NullableMonitoringAlertReason) Set(val *MonitoringAlertReason) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringAlertReason) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringAlertReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringAlertReason(val *MonitoringAlertReason) *NullableMonitoringAlertReason {
	return &NullableMonitoringAlertReason{value: val, isSet: true}
}

func (v NullableMonitoringAlertReason) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringAlertReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


