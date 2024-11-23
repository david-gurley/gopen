/*
 * Diagnostics API reference
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

// DiagnosticsModuleStatus ModuleStatus contains collected diagnostics of a process.
type DiagnosticsModuleStatus struct {
	// Category specifies whether process is part of Venice(controller) or Naples(io) subsystem.
	Category *string `json:"category,omitempty"`
	// Arbitrary string, json, backtrace, etc. offering clues for restart.
	LastRestartReason *string `json:"last-restart-reason,omitempty"`
	// Last start time.
	LastStart *time.Time `json:"last-start,omitempty"`
	// MACAddress of the smart nic on which this module runs.
	MacAddress *string `json:"mac-address,omitempty"`
	// Module is the name of the process/container.
	Module *string `json:"module,omitempty"`
	// Node on which this process is running.
	Node *string `json:"node,omitempty"`
	// Number of times process got restarted. zero if never restarted.
	RestartCount *int32 `json:"restart-count,omitempty"`
	// Service is the name of the service/pod this process is part of.
	Service *string `json:"service,omitempty"`
	// Ports on which this process is listening.
	ServicePorts *[]DiagnosticsServicePort `json:"service-ports,omitempty"`
}

// NewDiagnosticsModuleStatus instantiates a new DiagnosticsModuleStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiagnosticsModuleStatus() *DiagnosticsModuleStatus {
	this := DiagnosticsModuleStatus{}
	var category string = "venice"
	this.Category = &category
	return &this
}

// NewDiagnosticsModuleStatusWithDefaults instantiates a new DiagnosticsModuleStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiagnosticsModuleStatusWithDefaults() *DiagnosticsModuleStatus {
	this := DiagnosticsModuleStatus{}
	var category string = "venice"
	this.Category = &category
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *DiagnosticsModuleStatus) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticsModuleStatus) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *DiagnosticsModuleStatus) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *DiagnosticsModuleStatus) SetCategory(v string) {
	o.Category = &v
}

// GetLastRestartReason returns the LastRestartReason field value if set, zero value otherwise.
func (o *DiagnosticsModuleStatus) GetLastRestartReason() string {
	if o == nil || o.LastRestartReason == nil {
		var ret string
		return ret
	}
	return *o.LastRestartReason
}

// GetLastRestartReasonOk returns a tuple with the LastRestartReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticsModuleStatus) GetLastRestartReasonOk() (*string, bool) {
	if o == nil || o.LastRestartReason == nil {
		return nil, false
	}
	return o.LastRestartReason, true
}

// HasLastRestartReason returns a boolean if a field has been set.
func (o *DiagnosticsModuleStatus) HasLastRestartReason() bool {
	if o != nil && o.LastRestartReason != nil {
		return true
	}

	return false
}

// SetLastRestartReason gets a reference to the given string and assigns it to the LastRestartReason field.
func (o *DiagnosticsModuleStatus) SetLastRestartReason(v string) {
	o.LastRestartReason = &v
}

// GetLastStart returns the LastStart field value if set, zero value otherwise.
func (o *DiagnosticsModuleStatus) GetLastStart() time.Time {
	if o == nil || o.LastStart == nil {
		var ret time.Time
		return ret
	}
	return *o.LastStart
}

// GetLastStartOk returns a tuple with the LastStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticsModuleStatus) GetLastStartOk() (*time.Time, bool) {
	if o == nil || o.LastStart == nil {
		return nil, false
	}
	return o.LastStart, true
}

// HasLastStart returns a boolean if a field has been set.
func (o *DiagnosticsModuleStatus) HasLastStart() bool {
	if o != nil && o.LastStart != nil {
		return true
	}

	return false
}

// SetLastStart gets a reference to the given time.Time and assigns it to the LastStart field.
func (o *DiagnosticsModuleStatus) SetLastStart(v time.Time) {
	o.LastStart = &v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *DiagnosticsModuleStatus) GetMacAddress() string {
	if o == nil || o.MacAddress == nil {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticsModuleStatus) GetMacAddressOk() (*string, bool) {
	if o == nil || o.MacAddress == nil {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *DiagnosticsModuleStatus) HasMacAddress() bool {
	if o != nil && o.MacAddress != nil {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *DiagnosticsModuleStatus) SetMacAddress(v string) {
	o.MacAddress = &v
}

// GetModule returns the Module field value if set, zero value otherwise.
func (o *DiagnosticsModuleStatus) GetModule() string {
	if o == nil || o.Module == nil {
		var ret string
		return ret
	}
	return *o.Module
}

// GetModuleOk returns a tuple with the Module field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticsModuleStatus) GetModuleOk() (*string, bool) {
	if o == nil || o.Module == nil {
		return nil, false
	}
	return o.Module, true
}

// HasModule returns a boolean if a field has been set.
func (o *DiagnosticsModuleStatus) HasModule() bool {
	if o != nil && o.Module != nil {
		return true
	}

	return false
}

// SetModule gets a reference to the given string and assigns it to the Module field.
func (o *DiagnosticsModuleStatus) SetModule(v string) {
	o.Module = &v
}

// GetNode returns the Node field value if set, zero value otherwise.
func (o *DiagnosticsModuleStatus) GetNode() string {
	if o == nil || o.Node == nil {
		var ret string
		return ret
	}
	return *o.Node
}

// GetNodeOk returns a tuple with the Node field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticsModuleStatus) GetNodeOk() (*string, bool) {
	if o == nil || o.Node == nil {
		return nil, false
	}
	return o.Node, true
}

// HasNode returns a boolean if a field has been set.
func (o *DiagnosticsModuleStatus) HasNode() bool {
	if o != nil && o.Node != nil {
		return true
	}

	return false
}

// SetNode gets a reference to the given string and assigns it to the Node field.
func (o *DiagnosticsModuleStatus) SetNode(v string) {
	o.Node = &v
}

// GetRestartCount returns the RestartCount field value if set, zero value otherwise.
func (o *DiagnosticsModuleStatus) GetRestartCount() int32 {
	if o == nil || o.RestartCount == nil {
		var ret int32
		return ret
	}
	return *o.RestartCount
}

// GetRestartCountOk returns a tuple with the RestartCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticsModuleStatus) GetRestartCountOk() (*int32, bool) {
	if o == nil || o.RestartCount == nil {
		return nil, false
	}
	return o.RestartCount, true
}

// HasRestartCount returns a boolean if a field has been set.
func (o *DiagnosticsModuleStatus) HasRestartCount() bool {
	if o != nil && o.RestartCount != nil {
		return true
	}

	return false
}

// SetRestartCount gets a reference to the given int32 and assigns it to the RestartCount field.
func (o *DiagnosticsModuleStatus) SetRestartCount(v int32) {
	o.RestartCount = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *DiagnosticsModuleStatus) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticsModuleStatus) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *DiagnosticsModuleStatus) HasService() bool {
	if o != nil && o.Service != nil {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *DiagnosticsModuleStatus) SetService(v string) {
	o.Service = &v
}

// GetServicePorts returns the ServicePorts field value if set, zero value otherwise.
func (o *DiagnosticsModuleStatus) GetServicePorts() []DiagnosticsServicePort {
	if o == nil || o.ServicePorts == nil {
		var ret []DiagnosticsServicePort
		return ret
	}
	return *o.ServicePorts
}

// GetServicePortsOk returns a tuple with the ServicePorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticsModuleStatus) GetServicePortsOk() (*[]DiagnosticsServicePort, bool) {
	if o == nil || o.ServicePorts == nil {
		return nil, false
	}
	return o.ServicePorts, true
}

// HasServicePorts returns a boolean if a field has been set.
func (o *DiagnosticsModuleStatus) HasServicePorts() bool {
	if o != nil && o.ServicePorts != nil {
		return true
	}

	return false
}

// SetServicePorts gets a reference to the given []DiagnosticsServicePort and assigns it to the ServicePorts field.
func (o *DiagnosticsModuleStatus) SetServicePorts(v []DiagnosticsServicePort) {
	o.ServicePorts = &v
}

func (o DiagnosticsModuleStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.LastRestartReason != nil {
		toSerialize["last-restart-reason"] = o.LastRestartReason
	}
	if o.LastStart != nil {
		toSerialize["last-start"] = o.LastStart
	}
	if o.MacAddress != nil {
		toSerialize["mac-address"] = o.MacAddress
	}
	if o.Module != nil {
		toSerialize["module"] = o.Module
	}
	if o.Node != nil {
		toSerialize["node"] = o.Node
	}
	if o.RestartCount != nil {
		toSerialize["restart-count"] = o.RestartCount
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}
	if o.ServicePorts != nil {
		toSerialize["service-ports"] = o.ServicePorts
	}
	return json.Marshal(toSerialize)
}

type NullableDiagnosticsModuleStatus struct {
	value *DiagnosticsModuleStatus
	isSet bool
}

func (v NullableDiagnosticsModuleStatus) Get() *DiagnosticsModuleStatus {
	return v.value
}

func (v *NullableDiagnosticsModuleStatus) Set(val *DiagnosticsModuleStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableDiagnosticsModuleStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableDiagnosticsModuleStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiagnosticsModuleStatus(val *DiagnosticsModuleStatus) *NullableDiagnosticsModuleStatus {
	return &NullableDiagnosticsModuleStatus{value: val, isSet: true}
}

func (v NullableDiagnosticsModuleStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiagnosticsModuleStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

