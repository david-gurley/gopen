/*
 * Cluster API reference
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

// ClusterDistributedServiceCardList DistributedServiceCardList is a container object for list of DistributedServiceCard objects.
type ClusterDistributedServiceCardList struct {
	ApiVersion *string `json:"api-version,omitempty"`
	// List of DistributedServiceCard objects.
	Items *[]ClusterDistributedServiceCard `json:"items,omitempty"`
	Kind *string `json:"kind,omitempty"`
	ListMeta *ApiListMeta `json:"list-meta,omitempty"`
}

// NewClusterDistributedServiceCardList instantiates a new ClusterDistributedServiceCardList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterDistributedServiceCardList() *ClusterDistributedServiceCardList {
	this := ClusterDistributedServiceCardList{}
	return &this
}

// NewClusterDistributedServiceCardListWithDefaults instantiates a new ClusterDistributedServiceCardList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterDistributedServiceCardListWithDefaults() *ClusterDistributedServiceCardList {
	this := ClusterDistributedServiceCardList{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *ClusterDistributedServiceCardList) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDistributedServiceCardList) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *ClusterDistributedServiceCardList) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *ClusterDistributedServiceCardList) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ClusterDistributedServiceCardList) GetItems() []ClusterDistributedServiceCard {
	if o == nil || o.Items == nil {
		var ret []ClusterDistributedServiceCard
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDistributedServiceCardList) GetItemsOk() (*[]ClusterDistributedServiceCard, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ClusterDistributedServiceCardList) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ClusterDistributedServiceCard and assigns it to the Items field.
func (o *ClusterDistributedServiceCardList) SetItems(v []ClusterDistributedServiceCard) {
	o.Items = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ClusterDistributedServiceCardList) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDistributedServiceCardList) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ClusterDistributedServiceCardList) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *ClusterDistributedServiceCardList) SetKind(v string) {
	o.Kind = &v
}

// GetListMeta returns the ListMeta field value if set, zero value otherwise.
func (o *ClusterDistributedServiceCardList) GetListMeta() ApiListMeta {
	if o == nil || o.ListMeta == nil {
		var ret ApiListMeta
		return ret
	}
	return *o.ListMeta
}

// GetListMetaOk returns a tuple with the ListMeta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDistributedServiceCardList) GetListMetaOk() (*ApiListMeta, bool) {
	if o == nil || o.ListMeta == nil {
		return nil, false
	}
	return o.ListMeta, true
}

// HasListMeta returns a boolean if a field has been set.
func (o *ClusterDistributedServiceCardList) HasListMeta() bool {
	if o != nil && o.ListMeta != nil {
		return true
	}

	return false
}

// SetListMeta gets a reference to the given ApiListMeta and assigns it to the ListMeta field.
func (o *ClusterDistributedServiceCardList) SetListMeta(v ApiListMeta) {
	o.ListMeta = &v
}

func (o ClusterDistributedServiceCardList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiVersion != nil {
		toSerialize["api-version"] = o.ApiVersion
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.ListMeta != nil {
		toSerialize["list-meta"] = o.ListMeta
	}
	return json.Marshal(toSerialize)
}

type NullableClusterDistributedServiceCardList struct {
	value *ClusterDistributedServiceCardList
	isSet bool
}

func (v NullableClusterDistributedServiceCardList) Get() *ClusterDistributedServiceCardList {
	return v.value
}

func (v *NullableClusterDistributedServiceCardList) Set(val *ClusterDistributedServiceCardList) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterDistributedServiceCardList) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterDistributedServiceCardList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterDistributedServiceCardList(val *ClusterDistributedServiceCardList) *NullableClusterDistributedServiceCardList {
	return &NullableClusterDistributedServiceCardList{value: val, isSet: true}
}

func (v NullableClusterDistributedServiceCardList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterDistributedServiceCardList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


