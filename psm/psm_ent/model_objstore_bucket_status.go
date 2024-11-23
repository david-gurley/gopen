/*
 * Objstore API reference
 *
 *  
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package psm_ent

import (
	"encoding/json"
)

// ObjstoreBucketStatus struct for ObjstoreBucketStatus
type ObjstoreBucketStatus struct {
	NumObjects *int32 `json:"num-objects,omitempty"`
	TotalSize *int32 `json:"total-size,omitempty"`
}

// NewObjstoreBucketStatus instantiates a new ObjstoreBucketStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjstoreBucketStatus() *ObjstoreBucketStatus {
	this := ObjstoreBucketStatus{}
	return &this
}

// NewObjstoreBucketStatusWithDefaults instantiates a new ObjstoreBucketStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjstoreBucketStatusWithDefaults() *ObjstoreBucketStatus {
	this := ObjstoreBucketStatus{}
	return &this
}

// GetNumObjects returns the NumObjects field value if set, zero value otherwise.
func (o *ObjstoreBucketStatus) GetNumObjects() int32 {
	if o == nil || o.NumObjects == nil {
		var ret int32
		return ret
	}
	return *o.NumObjects
}

// GetNumObjectsOk returns a tuple with the NumObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjstoreBucketStatus) GetNumObjectsOk() (*int32, bool) {
	if o == nil || o.NumObjects == nil {
		return nil, false
	}
	return o.NumObjects, true
}

// HasNumObjects returns a boolean if a field has been set.
func (o *ObjstoreBucketStatus) HasNumObjects() bool {
	if o != nil && o.NumObjects != nil {
		return true
	}

	return false
}

// SetNumObjects gets a reference to the given int32 and assigns it to the NumObjects field.
func (o *ObjstoreBucketStatus) SetNumObjects(v int32) {
	o.NumObjects = &v
}

// GetTotalSize returns the TotalSize field value if set, zero value otherwise.
func (o *ObjstoreBucketStatus) GetTotalSize() int32 {
	if o == nil || o.TotalSize == nil {
		var ret int32
		return ret
	}
	return *o.TotalSize
}

// GetTotalSizeOk returns a tuple with the TotalSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjstoreBucketStatus) GetTotalSizeOk() (*int32, bool) {
	if o == nil || o.TotalSize == nil {
		return nil, false
	}
	return o.TotalSize, true
}

// HasTotalSize returns a boolean if a field has been set.
func (o *ObjstoreBucketStatus) HasTotalSize() bool {
	if o != nil && o.TotalSize != nil {
		return true
	}

	return false
}

// SetTotalSize gets a reference to the given int32 and assigns it to the TotalSize field.
func (o *ObjstoreBucketStatus) SetTotalSize(v int32) {
	o.TotalSize = &v
}

func (o ObjstoreBucketStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NumObjects != nil {
		toSerialize["num-objects"] = o.NumObjects
	}
	if o.TotalSize != nil {
		toSerialize["total-size"] = o.TotalSize
	}
	return json.Marshal(toSerialize)
}

type NullableObjstoreBucketStatus struct {
	value *ObjstoreBucketStatus
	isSet bool
}

func (v NullableObjstoreBucketStatus) Get() *ObjstoreBucketStatus {
	return v.value
}

func (v *NullableObjstoreBucketStatus) Set(val *ObjstoreBucketStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableObjstoreBucketStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableObjstoreBucketStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjstoreBucketStatus(val *ObjstoreBucketStatus) *NullableObjstoreBucketStatus {
	return &NullableObjstoreBucketStatus{value: val, isSet: true}
}

func (v NullableObjstoreBucketStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjstoreBucketStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


