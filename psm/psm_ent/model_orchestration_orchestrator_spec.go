/*
 * Orchestration API reference
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

// OrchestrationOrchestratorSpec OrchestratorSpec contains the configuration of the cluster.
type OrchestrationOrchestratorSpec struct {
	Credentials *MonitoringExternalCred `json:"credentials,omitempty"`
	// Namespaces are used to provide namespace specific information. From Rel-C this will be the only means to pass namespace information \"all_namespaces\" will be treated as a special namespace, which will apply the same configuration for all the namespaces discovered by the orchestrator.
	Namespaces *[]OrchestrationNamespaceSpec `json:"namespaces,omitempty"`
	// Type of orchestrator.
	Type *string `json:"type,omitempty"`
	// URI of the orchestrator. Length of string should be at least 1.
	Uri *string `json:"uri,omitempty"`
}

// NewOrchestrationOrchestratorSpec instantiates a new OrchestrationOrchestratorSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrchestrationOrchestratorSpec() *OrchestrationOrchestratorSpec {
	this := OrchestrationOrchestratorSpec{}
	var type_ string = "vcenter"
	this.Type = &type_
	return &this
}

// NewOrchestrationOrchestratorSpecWithDefaults instantiates a new OrchestrationOrchestratorSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrchestrationOrchestratorSpecWithDefaults() *OrchestrationOrchestratorSpec {
	this := OrchestrationOrchestratorSpec{}
	var type_ string = "vcenter"
	this.Type = &type_
	return &this
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *OrchestrationOrchestratorSpec) GetCredentials() MonitoringExternalCred {
	if o == nil || o.Credentials == nil {
		var ret MonitoringExternalCred
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrchestrationOrchestratorSpec) GetCredentialsOk() (*MonitoringExternalCred, bool) {
	if o == nil || o.Credentials == nil {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *OrchestrationOrchestratorSpec) HasCredentials() bool {
	if o != nil && o.Credentials != nil {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given MonitoringExternalCred and assigns it to the Credentials field.
func (o *OrchestrationOrchestratorSpec) SetCredentials(v MonitoringExternalCred) {
	o.Credentials = &v
}

// GetNamespaces returns the Namespaces field value if set, zero value otherwise.
func (o *OrchestrationOrchestratorSpec) GetNamespaces() []OrchestrationNamespaceSpec {
	if o == nil || o.Namespaces == nil {
		var ret []OrchestrationNamespaceSpec
		return ret
	}
	return *o.Namespaces
}

// GetNamespacesOk returns a tuple with the Namespaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrchestrationOrchestratorSpec) GetNamespacesOk() (*[]OrchestrationNamespaceSpec, bool) {
	if o == nil || o.Namespaces == nil {
		return nil, false
	}
	return o.Namespaces, true
}

// HasNamespaces returns a boolean if a field has been set.
func (o *OrchestrationOrchestratorSpec) HasNamespaces() bool {
	if o != nil && o.Namespaces != nil {
		return true
	}

	return false
}

// SetNamespaces gets a reference to the given []OrchestrationNamespaceSpec and assigns it to the Namespaces field.
func (o *OrchestrationOrchestratorSpec) SetNamespaces(v []OrchestrationNamespaceSpec) {
	o.Namespaces = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OrchestrationOrchestratorSpec) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrchestrationOrchestratorSpec) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OrchestrationOrchestratorSpec) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OrchestrationOrchestratorSpec) SetType(v string) {
	o.Type = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *OrchestrationOrchestratorSpec) GetUri() string {
	if o == nil || o.Uri == nil {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrchestrationOrchestratorSpec) GetUriOk() (*string, bool) {
	if o == nil || o.Uri == nil {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *OrchestrationOrchestratorSpec) HasUri() bool {
	if o != nil && o.Uri != nil {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *OrchestrationOrchestratorSpec) SetUri(v string) {
	o.Uri = &v
}

func (o OrchestrationOrchestratorSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Credentials != nil {
		toSerialize["credentials"] = o.Credentials
	}
	if o.Namespaces != nil {
		toSerialize["namespaces"] = o.Namespaces
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Uri != nil {
		toSerialize["uri"] = o.Uri
	}
	return json.Marshal(toSerialize)
}

type NullableOrchestrationOrchestratorSpec struct {
	value *OrchestrationOrchestratorSpec
	isSet bool
}

func (v NullableOrchestrationOrchestratorSpec) Get() *OrchestrationOrchestratorSpec {
	return v.value
}

func (v *NullableOrchestrationOrchestratorSpec) Set(val *OrchestrationOrchestratorSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableOrchestrationOrchestratorSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableOrchestrationOrchestratorSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrchestrationOrchestratorSpec(val *OrchestrationOrchestratorSpec) *NullableOrchestrationOrchestratorSpec {
	return &NullableOrchestrationOrchestratorSpec{value: val, isSet: true}
}

func (v NullableOrchestrationOrchestratorSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrchestrationOrchestratorSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

