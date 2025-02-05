/*
 * Objstore API reference
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

// ObjstoreAutoMsgBucketWatchHelper AutoMsgBucketWatchHelper is a wrapper object for watch events for Bucket objects.
type ObjstoreAutoMsgBucketWatchHelper struct {
	Events *[]ObjstoreAutoMsgBucketWatchHelperWatchEvent `json:"events,omitempty"`
}

// NewObjstoreAutoMsgBucketWatchHelper instantiates a new ObjstoreAutoMsgBucketWatchHelper object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjstoreAutoMsgBucketWatchHelper() *ObjstoreAutoMsgBucketWatchHelper {
	this := ObjstoreAutoMsgBucketWatchHelper{}
	return &this
}

// NewObjstoreAutoMsgBucketWatchHelperWithDefaults instantiates a new ObjstoreAutoMsgBucketWatchHelper object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjstoreAutoMsgBucketWatchHelperWithDefaults() *ObjstoreAutoMsgBucketWatchHelper {
	this := ObjstoreAutoMsgBucketWatchHelper{}
	return &this
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *ObjstoreAutoMsgBucketWatchHelper) GetEvents() []ObjstoreAutoMsgBucketWatchHelperWatchEvent {
	if o == nil || o.Events == nil {
		var ret []ObjstoreAutoMsgBucketWatchHelperWatchEvent
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjstoreAutoMsgBucketWatchHelper) GetEventsOk() (*[]ObjstoreAutoMsgBucketWatchHelperWatchEvent, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *ObjstoreAutoMsgBucketWatchHelper) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []ObjstoreAutoMsgBucketWatchHelperWatchEvent and assigns it to the Events field.
func (o *ObjstoreAutoMsgBucketWatchHelper) SetEvents(v []ObjstoreAutoMsgBucketWatchHelperWatchEvent) {
	o.Events = &v
}

func (o ObjstoreAutoMsgBucketWatchHelper) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableObjstoreAutoMsgBucketWatchHelper struct {
	value *ObjstoreAutoMsgBucketWatchHelper
	isSet bool
}

func (v NullableObjstoreAutoMsgBucketWatchHelper) Get() *ObjstoreAutoMsgBucketWatchHelper {
	return v.value
}

func (v *NullableObjstoreAutoMsgBucketWatchHelper) Set(val *ObjstoreAutoMsgBucketWatchHelper) {
	v.value = val
	v.isSet = true
}

func (v NullableObjstoreAutoMsgBucketWatchHelper) IsSet() bool {
	return v.isSet
}

func (v *NullableObjstoreAutoMsgBucketWatchHelper) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjstoreAutoMsgBucketWatchHelper(val *ObjstoreAutoMsgBucketWatchHelper) *NullableObjstoreAutoMsgBucketWatchHelper {
	return &NullableObjstoreAutoMsgBucketWatchHelper{value: val, isSet: true}
}

func (v NullableObjstoreAutoMsgBucketWatchHelper) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjstoreAutoMsgBucketWatchHelper) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


