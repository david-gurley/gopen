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

// WorkloadEndpointSpec spec part of Endpoint object.
type WorkloadEndpointSpec struct {
	// IP of the DSC where this endpoint exists.
	HomingHostAddr *string `json:"homing-host-addr,omitempty"`
	// MicroSegmentVlan to be assigned to the endpoint.
	MicroSegmentVlan *int64 `json:"micro-segment-vlan,omitempty"`
	// The DSC Name or MAC where the endpoint should reside.
	NodeUuid *string `json:"node-uuid,omitempty"`
	// NodeUUIDList has the list of DSCs where a EP is supposed to go to.
	NodeUuidList *[]string `json:"node-uuid-list,omitempty"`
	// Type is the type of Endpoint that is being created - L2 or L3.
	Type *string `json:"type,omitempty"`
}

// NewWorkloadEndpointSpec instantiates a new WorkloadEndpointSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkloadEndpointSpec() *WorkloadEndpointSpec {
	this := WorkloadEndpointSpec{}
	var type_ string = "l2"
	this.Type = &type_
	return &this
}

// NewWorkloadEndpointSpecWithDefaults instantiates a new WorkloadEndpointSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkloadEndpointSpecWithDefaults() *WorkloadEndpointSpec {
	this := WorkloadEndpointSpec{}
	var type_ string = "l2"
	this.Type = &type_
	return &this
}

// GetHomingHostAddr returns the HomingHostAddr field value if set, zero value otherwise.
func (o *WorkloadEndpointSpec) GetHomingHostAddr() string {
	if o == nil || o.HomingHostAddr == nil {
		var ret string
		return ret
	}
	return *o.HomingHostAddr
}

// GetHomingHostAddrOk returns a tuple with the HomingHostAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadEndpointSpec) GetHomingHostAddrOk() (*string, bool) {
	if o == nil || o.HomingHostAddr == nil {
		return nil, false
	}
	return o.HomingHostAddr, true
}

// HasHomingHostAddr returns a boolean if a field has been set.
func (o *WorkloadEndpointSpec) HasHomingHostAddr() bool {
	if o != nil && o.HomingHostAddr != nil {
		return true
	}

	return false
}

// SetHomingHostAddr gets a reference to the given string and assigns it to the HomingHostAddr field.
func (o *WorkloadEndpointSpec) SetHomingHostAddr(v string) {
	o.HomingHostAddr = &v
}

// GetMicroSegmentVlan returns the MicroSegmentVlan field value if set, zero value otherwise.
func (o *WorkloadEndpointSpec) GetMicroSegmentVlan() int64 {
	if o == nil || o.MicroSegmentVlan == nil {
		var ret int64
		return ret
	}
	return *o.MicroSegmentVlan
}

// GetMicroSegmentVlanOk returns a tuple with the MicroSegmentVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadEndpointSpec) GetMicroSegmentVlanOk() (*int64, bool) {
	if o == nil || o.MicroSegmentVlan == nil {
		return nil, false
	}
	return o.MicroSegmentVlan, true
}

// HasMicroSegmentVlan returns a boolean if a field has been set.
func (o *WorkloadEndpointSpec) HasMicroSegmentVlan() bool {
	if o != nil && o.MicroSegmentVlan != nil {
		return true
	}

	return false
}

// SetMicroSegmentVlan gets a reference to the given int64 and assigns it to the MicroSegmentVlan field.
func (o *WorkloadEndpointSpec) SetMicroSegmentVlan(v int64) {
	o.MicroSegmentVlan = &v
}

// GetNodeUuid returns the NodeUuid field value if set, zero value otherwise.
func (o *WorkloadEndpointSpec) GetNodeUuid() string {
	if o == nil || o.NodeUuid == nil {
		var ret string
		return ret
	}
	return *o.NodeUuid
}

// GetNodeUuidOk returns a tuple with the NodeUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadEndpointSpec) GetNodeUuidOk() (*string, bool) {
	if o == nil || o.NodeUuid == nil {
		return nil, false
	}
	return o.NodeUuid, true
}

// HasNodeUuid returns a boolean if a field has been set.
func (o *WorkloadEndpointSpec) HasNodeUuid() bool {
	if o != nil && o.NodeUuid != nil {
		return true
	}

	return false
}

// SetNodeUuid gets a reference to the given string and assigns it to the NodeUuid field.
func (o *WorkloadEndpointSpec) SetNodeUuid(v string) {
	o.NodeUuid = &v
}

// GetNodeUuidList returns the NodeUuidList field value if set, zero value otherwise.
func (o *WorkloadEndpointSpec) GetNodeUuidList() []string {
	if o == nil || o.NodeUuidList == nil {
		var ret []string
		return ret
	}
	return *o.NodeUuidList
}

// GetNodeUuidListOk returns a tuple with the NodeUuidList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadEndpointSpec) GetNodeUuidListOk() (*[]string, bool) {
	if o == nil || o.NodeUuidList == nil {
		return nil, false
	}
	return o.NodeUuidList, true
}

// HasNodeUuidList returns a boolean if a field has been set.
func (o *WorkloadEndpointSpec) HasNodeUuidList() bool {
	if o != nil && o.NodeUuidList != nil {
		return true
	}

	return false
}

// SetNodeUuidList gets a reference to the given []string and assigns it to the NodeUuidList field.
func (o *WorkloadEndpointSpec) SetNodeUuidList(v []string) {
	o.NodeUuidList = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WorkloadEndpointSpec) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadEndpointSpec) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WorkloadEndpointSpec) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WorkloadEndpointSpec) SetType(v string) {
	o.Type = &v
}

func (o WorkloadEndpointSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.HomingHostAddr != nil {
		toSerialize["homing-host-addr"] = o.HomingHostAddr
	}
	if o.MicroSegmentVlan != nil {
		toSerialize["micro-segment-vlan"] = o.MicroSegmentVlan
	}
	if o.NodeUuid != nil {
		toSerialize["node-uuid"] = o.NodeUuid
	}
	if o.NodeUuidList != nil {
		toSerialize["node-uuid-list"] = o.NodeUuidList
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableWorkloadEndpointSpec struct {
	value *WorkloadEndpointSpec
	isSet bool
}

func (v NullableWorkloadEndpointSpec) Get() *WorkloadEndpointSpec {
	return v.value
}

func (v *NullableWorkloadEndpointSpec) Set(val *WorkloadEndpointSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkloadEndpointSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkloadEndpointSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkloadEndpointSpec(val *WorkloadEndpointSpec) *NullableWorkloadEndpointSpec {
	return &NullableWorkloadEndpointSpec{value: val, isSet: true}
}

func (v NullableWorkloadEndpointSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkloadEndpointSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


