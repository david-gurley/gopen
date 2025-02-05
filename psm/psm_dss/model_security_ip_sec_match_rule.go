/*
 * Security API reference
 *
 *  
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package psm_dss

import (
	"encoding/json"
)

// SecurityIPSecMatchRule IPSecMatchRule : This is used to select packets to match the IPSec Rule.
type SecurityIPSecMatchRule struct {
	// list of apps objects to which the rule applies to.
	Apps *[]string `json:"apps,omitempty"`
	// outbound rule from a given ip-address/ip-mask/ip-range. Use any to refer to all ipaddresses cli-tags: id=to-ip.
	DstIpAddresses *[]string `json:"dst-ip-addresses,omitempty"`
	// list of (protocol, ports) pairs to which the rule applies to, in addition to apps.
	ProtoPorts *[]SecurityProtoPort `json:"proto-ports,omitempty"`
	// inbound rule from a given ip-address/ip-mask/ip-range. Use any to refer to all ipaddresses cli-tags: id=from-ip.
	SrcIpAddresses *[]string `json:"src-ip-addresses,omitempty"`
}

// NewSecurityIPSecMatchRule instantiates a new SecurityIPSecMatchRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityIPSecMatchRule() *SecurityIPSecMatchRule {
	this := SecurityIPSecMatchRule{}
	return &this
}

// NewSecurityIPSecMatchRuleWithDefaults instantiates a new SecurityIPSecMatchRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityIPSecMatchRuleWithDefaults() *SecurityIPSecMatchRule {
	this := SecurityIPSecMatchRule{}
	return &this
}

// GetApps returns the Apps field value if set, zero value otherwise.
func (o *SecurityIPSecMatchRule) GetApps() []string {
	if o == nil || o.Apps == nil {
		var ret []string
		return ret
	}
	return *o.Apps
}

// GetAppsOk returns a tuple with the Apps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityIPSecMatchRule) GetAppsOk() (*[]string, bool) {
	if o == nil || o.Apps == nil {
		return nil, false
	}
	return o.Apps, true
}

// HasApps returns a boolean if a field has been set.
func (o *SecurityIPSecMatchRule) HasApps() bool {
	if o != nil && o.Apps != nil {
		return true
	}

	return false
}

// SetApps gets a reference to the given []string and assigns it to the Apps field.
func (o *SecurityIPSecMatchRule) SetApps(v []string) {
	o.Apps = &v
}

// GetDstIpAddresses returns the DstIpAddresses field value if set, zero value otherwise.
func (o *SecurityIPSecMatchRule) GetDstIpAddresses() []string {
	if o == nil || o.DstIpAddresses == nil {
		var ret []string
		return ret
	}
	return *o.DstIpAddresses
}

// GetDstIpAddressesOk returns a tuple with the DstIpAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityIPSecMatchRule) GetDstIpAddressesOk() (*[]string, bool) {
	if o == nil || o.DstIpAddresses == nil {
		return nil, false
	}
	return o.DstIpAddresses, true
}

// HasDstIpAddresses returns a boolean if a field has been set.
func (o *SecurityIPSecMatchRule) HasDstIpAddresses() bool {
	if o != nil && o.DstIpAddresses != nil {
		return true
	}

	return false
}

// SetDstIpAddresses gets a reference to the given []string and assigns it to the DstIpAddresses field.
func (o *SecurityIPSecMatchRule) SetDstIpAddresses(v []string) {
	o.DstIpAddresses = &v
}

// GetProtoPorts returns the ProtoPorts field value if set, zero value otherwise.
func (o *SecurityIPSecMatchRule) GetProtoPorts() []SecurityProtoPort {
	if o == nil || o.ProtoPorts == nil {
		var ret []SecurityProtoPort
		return ret
	}
	return *o.ProtoPorts
}

// GetProtoPortsOk returns a tuple with the ProtoPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityIPSecMatchRule) GetProtoPortsOk() (*[]SecurityProtoPort, bool) {
	if o == nil || o.ProtoPorts == nil {
		return nil, false
	}
	return o.ProtoPorts, true
}

// HasProtoPorts returns a boolean if a field has been set.
func (o *SecurityIPSecMatchRule) HasProtoPorts() bool {
	if o != nil && o.ProtoPorts != nil {
		return true
	}

	return false
}

// SetProtoPorts gets a reference to the given []SecurityProtoPort and assigns it to the ProtoPorts field.
func (o *SecurityIPSecMatchRule) SetProtoPorts(v []SecurityProtoPort) {
	o.ProtoPorts = &v
}

// GetSrcIpAddresses returns the SrcIpAddresses field value if set, zero value otherwise.
func (o *SecurityIPSecMatchRule) GetSrcIpAddresses() []string {
	if o == nil || o.SrcIpAddresses == nil {
		var ret []string
		return ret
	}
	return *o.SrcIpAddresses
}

// GetSrcIpAddressesOk returns a tuple with the SrcIpAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityIPSecMatchRule) GetSrcIpAddressesOk() (*[]string, bool) {
	if o == nil || o.SrcIpAddresses == nil {
		return nil, false
	}
	return o.SrcIpAddresses, true
}

// HasSrcIpAddresses returns a boolean if a field has been set.
func (o *SecurityIPSecMatchRule) HasSrcIpAddresses() bool {
	if o != nil && o.SrcIpAddresses != nil {
		return true
	}

	return false
}

// SetSrcIpAddresses gets a reference to the given []string and assigns it to the SrcIpAddresses field.
func (o *SecurityIPSecMatchRule) SetSrcIpAddresses(v []string) {
	o.SrcIpAddresses = &v
}

func (o SecurityIPSecMatchRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Apps != nil {
		toSerialize["apps"] = o.Apps
	}
	if o.DstIpAddresses != nil {
		toSerialize["dst-ip-addresses"] = o.DstIpAddresses
	}
	if o.ProtoPorts != nil {
		toSerialize["proto-ports"] = o.ProtoPorts
	}
	if o.SrcIpAddresses != nil {
		toSerialize["src-ip-addresses"] = o.SrcIpAddresses
	}
	return json.Marshal(toSerialize)
}

type NullableSecurityIPSecMatchRule struct {
	value *SecurityIPSecMatchRule
	isSet bool
}

func (v NullableSecurityIPSecMatchRule) Get() *SecurityIPSecMatchRule {
	return v.value
}

func (v *NullableSecurityIPSecMatchRule) Set(val *SecurityIPSecMatchRule) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityIPSecMatchRule) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityIPSecMatchRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityIPSecMatchRule(val *SecurityIPSecMatchRule) *NullableSecurityIPSecMatchRule {
	return &NullableSecurityIPSecMatchRule{value: val, isSet: true}
}

func (v NullableSecurityIPSecMatchRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityIPSecMatchRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


