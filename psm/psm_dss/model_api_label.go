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
	"time"
)

// ApiLabel Label is the input for label operations.
type ApiLabel struct {
	ApiVersion *string `json:"api-version,omitempty"`
	CreationTime *time.Time `json:"creation-time,omitempty"`
	GenerationId *string `json:"generation-id,omitempty"`
	Kind *string `json:"kind,omitempty"`
	Labels *map[string]string `json:"labels,omitempty"`
	ModTime *time.Time `json:"mod-time,omitempty"`
	// Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64.
	Name *string `json:"name,omitempty"`
	// Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64.
	Namespace *string `json:"namespace,omitempty"`
	ResourceVersion *string `json:"resource-version,omitempty"`
	SelfLink *string `json:"self-link,omitempty"`
	// Must be alpha-numerics. Length of string should be between 1 and 48.
	Tenant *string `json:"tenant,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
}

// NewApiLabel instantiates a new ApiLabel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiLabel() *ApiLabel {
	this := ApiLabel{}
	return &this
}

// NewApiLabelWithDefaults instantiates a new ApiLabel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiLabelWithDefaults() *ApiLabel {
	this := ApiLabel{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *ApiLabel) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLabel) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *ApiLabel) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *ApiLabel) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *ApiLabel) GetCreationTime() time.Time {
	if o == nil || o.CreationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLabel) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || o.CreationTime == nil {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *ApiLabel) HasCreationTime() bool {
	if o != nil && o.CreationTime != nil {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *ApiLabel) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetGenerationId returns the GenerationId field value if set, zero value otherwise.
func (o *ApiLabel) GetGenerationId() string {
	if o == nil || o.GenerationId == nil {
		var ret string
		return ret
	}
	return *o.GenerationId
}

// GetGenerationIdOk returns a tuple with the GenerationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLabel) GetGenerationIdOk() (*string, bool) {
	if o == nil || o.GenerationId == nil {
		return nil, false
	}
	return o.GenerationId, true
}

// HasGenerationId returns a boolean if a field has been set.
func (o *ApiLabel) HasGenerationId() bool {
	if o != nil && o.GenerationId != nil {
		return true
	}

	return false
}

// SetGenerationId gets a reference to the given string and assigns it to the GenerationId field.
func (o *ApiLabel) SetGenerationId(v string) {
	o.GenerationId = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ApiLabel) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLabel) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ApiLabel) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *ApiLabel) SetKind(v string) {
	o.Kind = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *ApiLabel) GetLabels() map[string]string {
	if o == nil || o.Labels == nil {
		var ret map[string]string
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLabel) GetLabelsOk() (*map[string]string, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *ApiLabel) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]string and assigns it to the Labels field.
func (o *ApiLabel) SetLabels(v map[string]string) {
	o.Labels = &v
}

// GetModTime returns the ModTime field value if set, zero value otherwise.
func (o *ApiLabel) GetModTime() time.Time {
	if o == nil || o.ModTime == nil {
		var ret time.Time
		return ret
	}
	return *o.ModTime
}

// GetModTimeOk returns a tuple with the ModTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLabel) GetModTimeOk() (*time.Time, bool) {
	if o == nil || o.ModTime == nil {
		return nil, false
	}
	return o.ModTime, true
}

// HasModTime returns a boolean if a field has been set.
func (o *ApiLabel) HasModTime() bool {
	if o != nil && o.ModTime != nil {
		return true
	}

	return false
}

// SetModTime gets a reference to the given time.Time and assigns it to the ModTime field.
func (o *ApiLabel) SetModTime(v time.Time) {
	o.ModTime = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiLabel) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLabel) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiLabel) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiLabel) SetName(v string) {
	o.Name = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *ApiLabel) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLabel) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *ApiLabel) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *ApiLabel) SetNamespace(v string) {
	o.Namespace = &v
}

// GetResourceVersion returns the ResourceVersion field value if set, zero value otherwise.
func (o *ApiLabel) GetResourceVersion() string {
	if o == nil || o.ResourceVersion == nil {
		var ret string
		return ret
	}
	return *o.ResourceVersion
}

// GetResourceVersionOk returns a tuple with the ResourceVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLabel) GetResourceVersionOk() (*string, bool) {
	if o == nil || o.ResourceVersion == nil {
		return nil, false
	}
	return o.ResourceVersion, true
}

// HasResourceVersion returns a boolean if a field has been set.
func (o *ApiLabel) HasResourceVersion() bool {
	if o != nil && o.ResourceVersion != nil {
		return true
	}

	return false
}

// SetResourceVersion gets a reference to the given string and assigns it to the ResourceVersion field.
func (o *ApiLabel) SetResourceVersion(v string) {
	o.ResourceVersion = &v
}

// GetSelfLink returns the SelfLink field value if set, zero value otherwise.
func (o *ApiLabel) GetSelfLink() string {
	if o == nil || o.SelfLink == nil {
		var ret string
		return ret
	}
	return *o.SelfLink
}

// GetSelfLinkOk returns a tuple with the SelfLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLabel) GetSelfLinkOk() (*string, bool) {
	if o == nil || o.SelfLink == nil {
		return nil, false
	}
	return o.SelfLink, true
}

// HasSelfLink returns a boolean if a field has been set.
func (o *ApiLabel) HasSelfLink() bool {
	if o != nil && o.SelfLink != nil {
		return true
	}

	return false
}

// SetSelfLink gets a reference to the given string and assigns it to the SelfLink field.
func (o *ApiLabel) SetSelfLink(v string) {
	o.SelfLink = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *ApiLabel) GetTenant() string {
	if o == nil || o.Tenant == nil {
		var ret string
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLabel) GetTenantOk() (*string, bool) {
	if o == nil || o.Tenant == nil {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *ApiLabel) HasTenant() bool {
	if o != nil && o.Tenant != nil {
		return true
	}

	return false
}

// SetTenant gets a reference to the given string and assigns it to the Tenant field.
func (o *ApiLabel) SetTenant(v string) {
	o.Tenant = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *ApiLabel) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLabel) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *ApiLabel) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *ApiLabel) SetUuid(v string) {
	o.Uuid = &v
}

func (o ApiLabel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiVersion != nil {
		toSerialize["api-version"] = o.ApiVersion
	}
	if o.CreationTime != nil {
		toSerialize["creation-time"] = o.CreationTime
	}
	if o.GenerationId != nil {
		toSerialize["generation-id"] = o.GenerationId
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if o.ModTime != nil {
		toSerialize["mod-time"] = o.ModTime
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.ResourceVersion != nil {
		toSerialize["resource-version"] = o.ResourceVersion
	}
	if o.SelfLink != nil {
		toSerialize["self-link"] = o.SelfLink
	}
	if o.Tenant != nil {
		toSerialize["tenant"] = o.Tenant
	}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	return json.Marshal(toSerialize)
}

type NullableApiLabel struct {
	value *ApiLabel
	isSet bool
}

func (v NullableApiLabel) Get() *ApiLabel {
	return v.value
}

func (v *NullableApiLabel) Set(val *ApiLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableApiLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableApiLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiLabel(val *ApiLabel) *NullableApiLabel {
	return &NullableApiLabel{value: val, isSet: true}
}

func (v NullableApiLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


