/*
 * Workload API reference
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

// WorkloadWorkloadSpec Spec part of Workload object.
type WorkloadWorkloadSpec struct {
	// Hostname of the server where the workload should be running.
	HostName *string `json:"host-name,omitempty"`
	// Spec of all interfaces in the Workload identified by Primary MAC.
	Interfaces *[]WorkloadWorkloadIntfSpec `json:"interfaces,omitempty"`
	// Should be a valid time duration.
	MigrationTimeout *string `json:"migration-timeout,omitempty"`
}

// NewWorkloadWorkloadSpec instantiates a new WorkloadWorkloadSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkloadWorkloadSpec() *WorkloadWorkloadSpec {
	this := WorkloadWorkloadSpec{}
	var migrationTimeout string = "3m"
	this.MigrationTimeout = &migrationTimeout
	return &this
}

// NewWorkloadWorkloadSpecWithDefaults instantiates a new WorkloadWorkloadSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkloadWorkloadSpecWithDefaults() *WorkloadWorkloadSpec {
	this := WorkloadWorkloadSpec{}
	var migrationTimeout string = "3m"
	this.MigrationTimeout = &migrationTimeout
	return &this
}

// GetHostName returns the HostName field value if set, zero value otherwise.
func (o *WorkloadWorkloadSpec) GetHostName() string {
	if o == nil || o.HostName == nil {
		var ret string
		return ret
	}
	return *o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadWorkloadSpec) GetHostNameOk() (*string, bool) {
	if o == nil || o.HostName == nil {
		return nil, false
	}
	return o.HostName, true
}

// HasHostName returns a boolean if a field has been set.
func (o *WorkloadWorkloadSpec) HasHostName() bool {
	if o != nil && o.HostName != nil {
		return true
	}

	return false
}

// SetHostName gets a reference to the given string and assigns it to the HostName field.
func (o *WorkloadWorkloadSpec) SetHostName(v string) {
	o.HostName = &v
}

// GetInterfaces returns the Interfaces field value if set, zero value otherwise.
func (o *WorkloadWorkloadSpec) GetInterfaces() []WorkloadWorkloadIntfSpec {
	if o == nil || o.Interfaces == nil {
		var ret []WorkloadWorkloadIntfSpec
		return ret
	}
	return *o.Interfaces
}

// GetInterfacesOk returns a tuple with the Interfaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadWorkloadSpec) GetInterfacesOk() (*[]WorkloadWorkloadIntfSpec, bool) {
	if o == nil || o.Interfaces == nil {
		return nil, false
	}
	return o.Interfaces, true
}

// HasInterfaces returns a boolean if a field has been set.
func (o *WorkloadWorkloadSpec) HasInterfaces() bool {
	if o != nil && o.Interfaces != nil {
		return true
	}

	return false
}

// SetInterfaces gets a reference to the given []WorkloadWorkloadIntfSpec and assigns it to the Interfaces field.
func (o *WorkloadWorkloadSpec) SetInterfaces(v []WorkloadWorkloadIntfSpec) {
	o.Interfaces = &v
}

// GetMigrationTimeout returns the MigrationTimeout field value if set, zero value otherwise.
func (o *WorkloadWorkloadSpec) GetMigrationTimeout() string {
	if o == nil || o.MigrationTimeout == nil {
		var ret string
		return ret
	}
	return *o.MigrationTimeout
}

// GetMigrationTimeoutOk returns a tuple with the MigrationTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadWorkloadSpec) GetMigrationTimeoutOk() (*string, bool) {
	if o == nil || o.MigrationTimeout == nil {
		return nil, false
	}
	return o.MigrationTimeout, true
}

// HasMigrationTimeout returns a boolean if a field has been set.
func (o *WorkloadWorkloadSpec) HasMigrationTimeout() bool {
	if o != nil && o.MigrationTimeout != nil {
		return true
	}

	return false
}

// SetMigrationTimeout gets a reference to the given string and assigns it to the MigrationTimeout field.
func (o *WorkloadWorkloadSpec) SetMigrationTimeout(v string) {
	o.MigrationTimeout = &v
}

func (o WorkloadWorkloadSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.HostName != nil {
		toSerialize["host-name"] = o.HostName
	}
	if o.Interfaces != nil {
		toSerialize["interfaces"] = o.Interfaces
	}
	if o.MigrationTimeout != nil {
		toSerialize["migration-timeout"] = o.MigrationTimeout
	}
	return json.Marshal(toSerialize)
}

type NullableWorkloadWorkloadSpec struct {
	value *WorkloadWorkloadSpec
	isSet bool
}

func (v NullableWorkloadWorkloadSpec) Get() *WorkloadWorkloadSpec {
	return v.value
}

func (v *NullableWorkloadWorkloadSpec) Set(val *WorkloadWorkloadSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkloadWorkloadSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkloadWorkloadSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkloadWorkloadSpec(val *WorkloadWorkloadSpec) *NullableWorkloadWorkloadSpec {
	return &NullableWorkloadWorkloadSpec{value: val, isSet: true}
}

func (v NullableWorkloadWorkloadSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkloadWorkloadSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


