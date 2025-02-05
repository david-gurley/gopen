/*
 * Search API reference
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

// SearchPolicySearchRequest PolicySearchRequest is input to the security/firewall policy search request.
type SearchPolicySearchRequest struct {
	// Action, e.g. PERMIT, DENY or REJECT/CLEAR, PROTECT_PERMISSIVE or PROTECT_STRICT.
	Action *string `json:"action,omitempty"`
	// App specification,  predefined apps and alg config.
	App *string `json:"app,omitempty"`
	// Inbound ip-address, use any to refer to all ipaddresses eg: 10.1.1.1, any.
	FromIpAddress *string `json:"from-ip-address,omitempty"`
	// Inbound workload group.
	FromWorkloadGroup *string `json:"from-workload-group,omitempty"`
	// Kind of the policy that this request should act on. It should be either NetworkSecurityPolicy or IPSecPolicy.
	Kinds *[]string `json:"kinds,omitempty"`
	// Name is optional. If provided policy-search will be limited to the specified policy of the given name on the given kind. If empty, then all the policies of the given kind will be searched.
	Name *string `json:"name,omitempty"`
	// Namespace is optional. If provided policy-search will be limited to the specified namespace.
	Namespace *string `json:"namespace,omitempty"`
	// TCP or UDP Port number.
	Port *string `json:"port,omitempty"`
	// Protocol eg: tcp, udp, icmp.
	Protocol *string `json:"protocol,omitempty"`
	// Tenant Name, to perform query within a Tenant's scope. The default tenant is \"default\". In the backend this field gets auto-filled & validated by apigw-hook based on user login context.
	Tenant *string `json:"tenant,omitempty"`
	// Outbound ip-address, use any to refer to all ipaddresses eg: 20.1.1.1, any.
	ToIpAddress *string `json:"to-ip-address,omitempty"`
	// Outbound workload group.
	ToWorkloadGroup *string `json:"to-workload-group,omitempty"`
}

// NewSearchPolicySearchRequest instantiates a new SearchPolicySearchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchPolicySearchRequest() *SearchPolicySearchRequest {
	this := SearchPolicySearchRequest{}
	var namespace string = "default"
	this.Namespace = &namespace
	var tenant string = "default"
	this.Tenant = &tenant
	return &this
}

// NewSearchPolicySearchRequestWithDefaults instantiates a new SearchPolicySearchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchPolicySearchRequestWithDefaults() *SearchPolicySearchRequest {
	this := SearchPolicySearchRequest{}
	var namespace string = "default"
	this.Namespace = &namespace
	var tenant string = "default"
	this.Tenant = &tenant
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *SearchPolicySearchRequest) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchPolicySearchRequest) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *SearchPolicySearchRequest) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *SearchPolicySearchRequest) SetAction(v string) {
	o.Action = &v
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *SearchPolicySearchRequest) GetApp() string {
	if o == nil || o.App == nil {
		var ret string
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchPolicySearchRequest) GetAppOk() (*string, bool) {
	if o == nil || o.App == nil {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *SearchPolicySearchRequest) HasApp() bool {
	if o != nil && o.App != nil {
		return true
	}

	return false
}

// SetApp gets a reference to the given string and assigns it to the App field.
func (o *SearchPolicySearchRequest) SetApp(v string) {
	o.App = &v
}

// GetFromIpAddress returns the FromIpAddress field value if set, zero value otherwise.
func (o *SearchPolicySearchRequest) GetFromIpAddress() string {
	if o == nil || o.FromIpAddress == nil {
		var ret string
		return ret
	}
	return *o.FromIpAddress
}

// GetFromIpAddressOk returns a tuple with the FromIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchPolicySearchRequest) GetFromIpAddressOk() (*string, bool) {
	if o == nil || o.FromIpAddress == nil {
		return nil, false
	}
	return o.FromIpAddress, true
}

// HasFromIpAddress returns a boolean if a field has been set.
func (o *SearchPolicySearchRequest) HasFromIpAddress() bool {
	if o != nil && o.FromIpAddress != nil {
		return true
	}

	return false
}

// SetFromIpAddress gets a reference to the given string and assigns it to the FromIpAddress field.
func (o *SearchPolicySearchRequest) SetFromIpAddress(v string) {
	o.FromIpAddress = &v
}

// GetFromWorkloadGroup returns the FromWorkloadGroup field value if set, zero value otherwise.
func (o *SearchPolicySearchRequest) GetFromWorkloadGroup() string {
	if o == nil || o.FromWorkloadGroup == nil {
		var ret string
		return ret
	}
	return *o.FromWorkloadGroup
}

// GetFromWorkloadGroupOk returns a tuple with the FromWorkloadGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchPolicySearchRequest) GetFromWorkloadGroupOk() (*string, bool) {
	if o == nil || o.FromWorkloadGroup == nil {
		return nil, false
	}
	return o.FromWorkloadGroup, true
}

// HasFromWorkloadGroup returns a boolean if a field has been set.
func (o *SearchPolicySearchRequest) HasFromWorkloadGroup() bool {
	if o != nil && o.FromWorkloadGroup != nil {
		return true
	}

	return false
}

// SetFromWorkloadGroup gets a reference to the given string and assigns it to the FromWorkloadGroup field.
func (o *SearchPolicySearchRequest) SetFromWorkloadGroup(v string) {
	o.FromWorkloadGroup = &v
}

// GetKinds returns the Kinds field value if set, zero value otherwise.
func (o *SearchPolicySearchRequest) GetKinds() []string {
	if o == nil || o.Kinds == nil {
		var ret []string
		return ret
	}
	return *o.Kinds
}

// GetKindsOk returns a tuple with the Kinds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchPolicySearchRequest) GetKindsOk() (*[]string, bool) {
	if o == nil || o.Kinds == nil {
		return nil, false
	}
	return o.Kinds, true
}

// HasKinds returns a boolean if a field has been set.
func (o *SearchPolicySearchRequest) HasKinds() bool {
	if o != nil && o.Kinds != nil {
		return true
	}

	return false
}

// SetKinds gets a reference to the given []string and assigns it to the Kinds field.
func (o *SearchPolicySearchRequest) SetKinds(v []string) {
	o.Kinds = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SearchPolicySearchRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchPolicySearchRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SearchPolicySearchRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SearchPolicySearchRequest) SetName(v string) {
	o.Name = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *SearchPolicySearchRequest) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchPolicySearchRequest) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *SearchPolicySearchRequest) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *SearchPolicySearchRequest) SetNamespace(v string) {
	o.Namespace = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *SearchPolicySearchRequest) GetPort() string {
	if o == nil || o.Port == nil {
		var ret string
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchPolicySearchRequest) GetPortOk() (*string, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *SearchPolicySearchRequest) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given string and assigns it to the Port field.
func (o *SearchPolicySearchRequest) SetPort(v string) {
	o.Port = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *SearchPolicySearchRequest) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchPolicySearchRequest) GetProtocolOk() (*string, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *SearchPolicySearchRequest) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *SearchPolicySearchRequest) SetProtocol(v string) {
	o.Protocol = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *SearchPolicySearchRequest) GetTenant() string {
	if o == nil || o.Tenant == nil {
		var ret string
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchPolicySearchRequest) GetTenantOk() (*string, bool) {
	if o == nil || o.Tenant == nil {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *SearchPolicySearchRequest) HasTenant() bool {
	if o != nil && o.Tenant != nil {
		return true
	}

	return false
}

// SetTenant gets a reference to the given string and assigns it to the Tenant field.
func (o *SearchPolicySearchRequest) SetTenant(v string) {
	o.Tenant = &v
}

// GetToIpAddress returns the ToIpAddress field value if set, zero value otherwise.
func (o *SearchPolicySearchRequest) GetToIpAddress() string {
	if o == nil || o.ToIpAddress == nil {
		var ret string
		return ret
	}
	return *o.ToIpAddress
}

// GetToIpAddressOk returns a tuple with the ToIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchPolicySearchRequest) GetToIpAddressOk() (*string, bool) {
	if o == nil || o.ToIpAddress == nil {
		return nil, false
	}
	return o.ToIpAddress, true
}

// HasToIpAddress returns a boolean if a field has been set.
func (o *SearchPolicySearchRequest) HasToIpAddress() bool {
	if o != nil && o.ToIpAddress != nil {
		return true
	}

	return false
}

// SetToIpAddress gets a reference to the given string and assigns it to the ToIpAddress field.
func (o *SearchPolicySearchRequest) SetToIpAddress(v string) {
	o.ToIpAddress = &v
}

// GetToWorkloadGroup returns the ToWorkloadGroup field value if set, zero value otherwise.
func (o *SearchPolicySearchRequest) GetToWorkloadGroup() string {
	if o == nil || o.ToWorkloadGroup == nil {
		var ret string
		return ret
	}
	return *o.ToWorkloadGroup
}

// GetToWorkloadGroupOk returns a tuple with the ToWorkloadGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchPolicySearchRequest) GetToWorkloadGroupOk() (*string, bool) {
	if o == nil || o.ToWorkloadGroup == nil {
		return nil, false
	}
	return o.ToWorkloadGroup, true
}

// HasToWorkloadGroup returns a boolean if a field has been set.
func (o *SearchPolicySearchRequest) HasToWorkloadGroup() bool {
	if o != nil && o.ToWorkloadGroup != nil {
		return true
	}

	return false
}

// SetToWorkloadGroup gets a reference to the given string and assigns it to the ToWorkloadGroup field.
func (o *SearchPolicySearchRequest) SetToWorkloadGroup(v string) {
	o.ToWorkloadGroup = &v
}

func (o SearchPolicySearchRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.App != nil {
		toSerialize["app"] = o.App
	}
	if o.FromIpAddress != nil {
		toSerialize["from-ip-address"] = o.FromIpAddress
	}
	if o.FromWorkloadGroup != nil {
		toSerialize["from-workload-group"] = o.FromWorkloadGroup
	}
	if o.Kinds != nil {
		toSerialize["kinds"] = o.Kinds
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.Protocol != nil {
		toSerialize["protocol"] = o.Protocol
	}
	if o.Tenant != nil {
		toSerialize["tenant"] = o.Tenant
	}
	if o.ToIpAddress != nil {
		toSerialize["to-ip-address"] = o.ToIpAddress
	}
	if o.ToWorkloadGroup != nil {
		toSerialize["to-workload-group"] = o.ToWorkloadGroup
	}
	return json.Marshal(toSerialize)
}

type NullableSearchPolicySearchRequest struct {
	value *SearchPolicySearchRequest
	isSet bool
}

func (v NullableSearchPolicySearchRequest) Get() *SearchPolicySearchRequest {
	return v.value
}

func (v *NullableSearchPolicySearchRequest) Set(val *SearchPolicySearchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchPolicySearchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchPolicySearchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchPolicySearchRequest(val *SearchPolicySearchRequest) *NullableSearchPolicySearchRequest {
	return &NullableSearchPolicySearchRequest{value: val, isSet: true}
}

func (v NullableSearchPolicySearchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchPolicySearchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


