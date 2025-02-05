/*
 * Orchestration API reference
 *
 * Service name  
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package psm_dss

import (
	"encoding/json"
	"time"
)

// OrchestrationOrchestratorStatus OrchestratorStatus contains the current state of connection with the orchestrator.
type OrchestrationOrchestratorStatus struct {
	ConnectionStatus *string `json:"connection-status,omitempty"`
	DiscoveredNamespaces *[]string `json:"discovered-namespaces,omitempty"`
	IncompatibleDscs *[]string `json:"incompatible-dscs,omitempty"`
	LastTransitionTime *time.Time `json:"last-transition-time,omitempty"`
	Message *string `json:"message,omitempty"`
	OrchId *int32 `json:"orch-id,omitempty"`
}

// NewOrchestrationOrchestratorStatus instantiates a new OrchestrationOrchestratorStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrchestrationOrchestratorStatus() *OrchestrationOrchestratorStatus {
	this := OrchestrationOrchestratorStatus{}
	var connectionStatus string = "unknown"
	this.ConnectionStatus = &connectionStatus
	return &this
}

// NewOrchestrationOrchestratorStatusWithDefaults instantiates a new OrchestrationOrchestratorStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrchestrationOrchestratorStatusWithDefaults() *OrchestrationOrchestratorStatus {
	this := OrchestrationOrchestratorStatus{}
	var connectionStatus string = "unknown"
	this.ConnectionStatus = &connectionStatus
	return &this
}

// GetConnectionStatus returns the ConnectionStatus field value if set, zero value otherwise.
func (o *OrchestrationOrchestratorStatus) GetConnectionStatus() string {
	if o == nil || o.ConnectionStatus == nil {
		var ret string
		return ret
	}
	return *o.ConnectionStatus
}

// GetConnectionStatusOk returns a tuple with the ConnectionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrchestrationOrchestratorStatus) GetConnectionStatusOk() (*string, bool) {
	if o == nil || o.ConnectionStatus == nil {
		return nil, false
	}
	return o.ConnectionStatus, true
}

// HasConnectionStatus returns a boolean if a field has been set.
func (o *OrchestrationOrchestratorStatus) HasConnectionStatus() bool {
	if o != nil && o.ConnectionStatus != nil {
		return true
	}

	return false
}

// SetConnectionStatus gets a reference to the given string and assigns it to the ConnectionStatus field.
func (o *OrchestrationOrchestratorStatus) SetConnectionStatus(v string) {
	o.ConnectionStatus = &v
}

// GetDiscoveredNamespaces returns the DiscoveredNamespaces field value if set, zero value otherwise.
func (o *OrchestrationOrchestratorStatus) GetDiscoveredNamespaces() []string {
	if o == nil || o.DiscoveredNamespaces == nil {
		var ret []string
		return ret
	}
	return *o.DiscoveredNamespaces
}

// GetDiscoveredNamespacesOk returns a tuple with the DiscoveredNamespaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrchestrationOrchestratorStatus) GetDiscoveredNamespacesOk() (*[]string, bool) {
	if o == nil || o.DiscoveredNamespaces == nil {
		return nil, false
	}
	return o.DiscoveredNamespaces, true
}

// HasDiscoveredNamespaces returns a boolean if a field has been set.
func (o *OrchestrationOrchestratorStatus) HasDiscoveredNamespaces() bool {
	if o != nil && o.DiscoveredNamespaces != nil {
		return true
	}

	return false
}

// SetDiscoveredNamespaces gets a reference to the given []string and assigns it to the DiscoveredNamespaces field.
func (o *OrchestrationOrchestratorStatus) SetDiscoveredNamespaces(v []string) {
	o.DiscoveredNamespaces = &v
}

// GetIncompatibleDscs returns the IncompatibleDscs field value if set, zero value otherwise.
func (o *OrchestrationOrchestratorStatus) GetIncompatibleDscs() []string {
	if o == nil || o.IncompatibleDscs == nil {
		var ret []string
		return ret
	}
	return *o.IncompatibleDscs
}

// GetIncompatibleDscsOk returns a tuple with the IncompatibleDscs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrchestrationOrchestratorStatus) GetIncompatibleDscsOk() (*[]string, bool) {
	if o == nil || o.IncompatibleDscs == nil {
		return nil, false
	}
	return o.IncompatibleDscs, true
}

// HasIncompatibleDscs returns a boolean if a field has been set.
func (o *OrchestrationOrchestratorStatus) HasIncompatibleDscs() bool {
	if o != nil && o.IncompatibleDscs != nil {
		return true
	}

	return false
}

// SetIncompatibleDscs gets a reference to the given []string and assigns it to the IncompatibleDscs field.
func (o *OrchestrationOrchestratorStatus) SetIncompatibleDscs(v []string) {
	o.IncompatibleDscs = &v
}

// GetLastTransitionTime returns the LastTransitionTime field value if set, zero value otherwise.
func (o *OrchestrationOrchestratorStatus) GetLastTransitionTime() time.Time {
	if o == nil || o.LastTransitionTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastTransitionTime
}

// GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrchestrationOrchestratorStatus) GetLastTransitionTimeOk() (*time.Time, bool) {
	if o == nil || o.LastTransitionTime == nil {
		return nil, false
	}
	return o.LastTransitionTime, true
}

// HasLastTransitionTime returns a boolean if a field has been set.
func (o *OrchestrationOrchestratorStatus) HasLastTransitionTime() bool {
	if o != nil && o.LastTransitionTime != nil {
		return true
	}

	return false
}

// SetLastTransitionTime gets a reference to the given time.Time and assigns it to the LastTransitionTime field.
func (o *OrchestrationOrchestratorStatus) SetLastTransitionTime(v time.Time) {
	o.LastTransitionTime = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *OrchestrationOrchestratorStatus) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrchestrationOrchestratorStatus) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *OrchestrationOrchestratorStatus) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *OrchestrationOrchestratorStatus) SetMessage(v string) {
	o.Message = &v
}

// GetOrchId returns the OrchId field value if set, zero value otherwise.
func (o *OrchestrationOrchestratorStatus) GetOrchId() int32 {
	if o == nil || o.OrchId == nil {
		var ret int32
		return ret
	}
	return *o.OrchId
}

// GetOrchIdOk returns a tuple with the OrchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrchestrationOrchestratorStatus) GetOrchIdOk() (*int32, bool) {
	if o == nil || o.OrchId == nil {
		return nil, false
	}
	return o.OrchId, true
}

// HasOrchId returns a boolean if a field has been set.
func (o *OrchestrationOrchestratorStatus) HasOrchId() bool {
	if o != nil && o.OrchId != nil {
		return true
	}

	return false
}

// SetOrchId gets a reference to the given int32 and assigns it to the OrchId field.
func (o *OrchestrationOrchestratorStatus) SetOrchId(v int32) {
	o.OrchId = &v
}

func (o OrchestrationOrchestratorStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ConnectionStatus != nil {
		toSerialize["connection-status"] = o.ConnectionStatus
	}
	if o.DiscoveredNamespaces != nil {
		toSerialize["discovered-namespaces"] = o.DiscoveredNamespaces
	}
	if o.IncompatibleDscs != nil {
		toSerialize["incompatible-dscs"] = o.IncompatibleDscs
	}
	if o.LastTransitionTime != nil {
		toSerialize["last-transition-time"] = o.LastTransitionTime
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.OrchId != nil {
		toSerialize["orch-id"] = o.OrchId
	}
	return json.Marshal(toSerialize)
}

type NullableOrchestrationOrchestratorStatus struct {
	value *OrchestrationOrchestratorStatus
	isSet bool
}

func (v NullableOrchestrationOrchestratorStatus) Get() *OrchestrationOrchestratorStatus {
	return v.value
}

func (v *NullableOrchestrationOrchestratorStatus) Set(val *OrchestrationOrchestratorStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableOrchestrationOrchestratorStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableOrchestrationOrchestratorStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrchestrationOrchestratorStatus(val *OrchestrationOrchestratorStatus) *NullableOrchestrationOrchestratorStatus {
	return &NullableOrchestrationOrchestratorStatus{value: val, isSet: true}
}

func (v NullableOrchestrationOrchestratorStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrchestrationOrchestratorStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


