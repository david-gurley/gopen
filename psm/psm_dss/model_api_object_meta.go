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

// ApiObjectMeta ObjectMeta contains metadata that all objects stored in kvstore must have.
type ApiObjectMeta struct {
	// System generated and updated, not updatable by user.
	CreationTime *time.Time `json:"creation-time,omitempty"`
	// This is incremented anytime there is an update to the user intent, including Spec update and any update to ObjectMeta. System generated and updated, not updatable by user.
	GenerationId *string `json:"generation-id,omitempty"`
	Labels *map[string]string `json:"labels,omitempty"`
	// System generated and updated, not updatable by user.
	ModTime *time.Time `json:"mod-time,omitempty"`
	// Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64.
	Name *string `json:"name,omitempty"`
	// Namespace of the object, for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64.
	Namespace *string `json:"namespace,omitempty"`
	// This is updated anytime there is any change to the object. System generated and updated, not updatable by user.
	ResourceVersion *string `json:"resource-version,omitempty"`
	// When the object is served from the API-GW it is the URI path. Example: - \"/v1/tenants/tenants/tenant2\" System generated and updated, not updatable by user.
	SelfLink *string `json:"self-link,omitempty"`
	// This can be automatically filled in many cases based on the tenant the user, who created the object, belongs to. Must be alpha-numerics. Length of string should be between 1 and 48.
	Tenant *string `json:"tenant,omitempty"`
	// This is generated on creation of the object. System generated, not updatable by user.
	Uuid *string `json:"uuid,omitempty"`
}

// NewApiObjectMeta instantiates a new ApiObjectMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiObjectMeta() *ApiObjectMeta {
	this := ApiObjectMeta{}
	return &this
}

// NewApiObjectMetaWithDefaults instantiates a new ApiObjectMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiObjectMetaWithDefaults() *ApiObjectMeta {
	this := ApiObjectMeta{}
	return &this
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *ApiObjectMeta) GetCreationTime() time.Time {
	if o == nil || o.CreationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiObjectMeta) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || o.CreationTime == nil {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *ApiObjectMeta) HasCreationTime() bool {
	if o != nil && o.CreationTime != nil {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *ApiObjectMeta) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetGenerationId returns the GenerationId field value if set, zero value otherwise.
func (o *ApiObjectMeta) GetGenerationId() string {
	if o == nil || o.GenerationId == nil {
		var ret string
		return ret
	}
	return *o.GenerationId
}

// GetGenerationIdOk returns a tuple with the GenerationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiObjectMeta) GetGenerationIdOk() (*string, bool) {
	if o == nil || o.GenerationId == nil {
		return nil, false
	}
	return o.GenerationId, true
}

// HasGenerationId returns a boolean if a field has been set.
func (o *ApiObjectMeta) HasGenerationId() bool {
	if o != nil && o.GenerationId != nil {
		return true
	}

	return false
}

// SetGenerationId gets a reference to the given string and assigns it to the GenerationId field.
func (o *ApiObjectMeta) SetGenerationId(v string) {
	o.GenerationId = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *ApiObjectMeta) GetLabels() map[string]string {
	if o == nil || o.Labels == nil {
		var ret map[string]string
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiObjectMeta) GetLabelsOk() (*map[string]string, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *ApiObjectMeta) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]string and assigns it to the Labels field.
func (o *ApiObjectMeta) SetLabels(v map[string]string) {
	o.Labels = &v
}

// GetModTime returns the ModTime field value if set, zero value otherwise.
func (o *ApiObjectMeta) GetModTime() time.Time {
	if o == nil || o.ModTime == nil {
		var ret time.Time
		return ret
	}
	return *o.ModTime
}

// GetModTimeOk returns a tuple with the ModTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiObjectMeta) GetModTimeOk() (*time.Time, bool) {
	if o == nil || o.ModTime == nil {
		return nil, false
	}
	return o.ModTime, true
}

// HasModTime returns a boolean if a field has been set.
func (o *ApiObjectMeta) HasModTime() bool {
	if o != nil && o.ModTime != nil {
		return true
	}

	return false
}

// SetModTime gets a reference to the given time.Time and assigns it to the ModTime field.
func (o *ApiObjectMeta) SetModTime(v time.Time) {
	o.ModTime = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiObjectMeta) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiObjectMeta) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiObjectMeta) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiObjectMeta) SetName(v string) {
	o.Name = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *ApiObjectMeta) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiObjectMeta) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *ApiObjectMeta) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *ApiObjectMeta) SetNamespace(v string) {
	o.Namespace = &v
}

// GetResourceVersion returns the ResourceVersion field value if set, zero value otherwise.
func (o *ApiObjectMeta) GetResourceVersion() string {
	if o == nil || o.ResourceVersion == nil {
		var ret string
		return ret
	}
	return *o.ResourceVersion
}

// GetResourceVersionOk returns a tuple with the ResourceVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiObjectMeta) GetResourceVersionOk() (*string, bool) {
	if o == nil || o.ResourceVersion == nil {
		return nil, false
	}
	return o.ResourceVersion, true
}

// HasResourceVersion returns a boolean if a field has been set.
func (o *ApiObjectMeta) HasResourceVersion() bool {
	if o != nil && o.ResourceVersion != nil {
		return true
	}

	return false
}

// SetResourceVersion gets a reference to the given string and assigns it to the ResourceVersion field.
func (o *ApiObjectMeta) SetResourceVersion(v string) {
	o.ResourceVersion = &v
}

// GetSelfLink returns the SelfLink field value if set, zero value otherwise.
func (o *ApiObjectMeta) GetSelfLink() string {
	if o == nil || o.SelfLink == nil {
		var ret string
		return ret
	}
	return *o.SelfLink
}

// GetSelfLinkOk returns a tuple with the SelfLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiObjectMeta) GetSelfLinkOk() (*string, bool) {
	if o == nil || o.SelfLink == nil {
		return nil, false
	}
	return o.SelfLink, true
}

// HasSelfLink returns a boolean if a field has been set.
func (o *ApiObjectMeta) HasSelfLink() bool {
	if o != nil && o.SelfLink != nil {
		return true
	}

	return false
}

// SetSelfLink gets a reference to the given string and assigns it to the SelfLink field.
func (o *ApiObjectMeta) SetSelfLink(v string) {
	o.SelfLink = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *ApiObjectMeta) GetTenant() string {
	if o == nil || o.Tenant == nil {
		var ret string
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiObjectMeta) GetTenantOk() (*string, bool) {
	if o == nil || o.Tenant == nil {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *ApiObjectMeta) HasTenant() bool {
	if o != nil && o.Tenant != nil {
		return true
	}

	return false
}

// SetTenant gets a reference to the given string and assigns it to the Tenant field.
func (o *ApiObjectMeta) SetTenant(v string) {
	o.Tenant = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *ApiObjectMeta) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiObjectMeta) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *ApiObjectMeta) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *ApiObjectMeta) SetUuid(v string) {
	o.Uuid = &v
}

func (o ApiObjectMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreationTime != nil {
		toSerialize["creation-time"] = o.CreationTime
	}
	if o.GenerationId != nil {
		toSerialize["generation-id"] = o.GenerationId
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

type NullableApiObjectMeta struct {
	value *ApiObjectMeta
	isSet bool
}

func (v NullableApiObjectMeta) Get() *ApiObjectMeta {
	return v.value
}

func (v *NullableApiObjectMeta) Set(val *ApiObjectMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableApiObjectMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableApiObjectMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiObjectMeta(val *ApiObjectMeta) *NullableApiObjectMeta {
	return &NullableApiObjectMeta{value: val, isSet: true}
}

func (v NullableApiObjectMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiObjectMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


