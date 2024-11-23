/*
 * Workload API reference
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

// ApiStatus Status is returned for calls that dont return objects.
type ApiStatus struct {
	ApiVersion *string `json:"api-version,omitempty"`
	// Code is the HTTP status code.
	Code *int32 `json:"code,omitempty"`
	Kind *string `json:"kind,omitempty"`
	// Message contains human readable form of the error.
	Message *[]string `json:"message,omitempty"`
	ObjectRef *ApiObjectRef `json:"object-ref,omitempty"`
	Result *ApiStatusResult `json:"result,omitempty"`
}

// NewApiStatus instantiates a new ApiStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiStatus() *ApiStatus {
	this := ApiStatus{}
	return &this
}

// NewApiStatusWithDefaults instantiates a new ApiStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiStatusWithDefaults() *ApiStatus {
	this := ApiStatus{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *ApiStatus) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStatus) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *ApiStatus) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *ApiStatus) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ApiStatus) GetCode() int32 {
	if o == nil || o.Code == nil {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStatus) GetCodeOk() (*int32, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ApiStatus) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *ApiStatus) SetCode(v int32) {
	o.Code = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ApiStatus) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStatus) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ApiStatus) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *ApiStatus) SetKind(v string) {
	o.Kind = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ApiStatus) GetMessage() []string {
	if o == nil || o.Message == nil {
		var ret []string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStatus) GetMessageOk() (*[]string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ApiStatus) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given []string and assigns it to the Message field.
func (o *ApiStatus) SetMessage(v []string) {
	o.Message = &v
}

// GetObjectRef returns the ObjectRef field value if set, zero value otherwise.
func (o *ApiStatus) GetObjectRef() ApiObjectRef {
	if o == nil || o.ObjectRef == nil {
		var ret ApiObjectRef
		return ret
	}
	return *o.ObjectRef
}

// GetObjectRefOk returns a tuple with the ObjectRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStatus) GetObjectRefOk() (*ApiObjectRef, bool) {
	if o == nil || o.ObjectRef == nil {
		return nil, false
	}
	return o.ObjectRef, true
}

// HasObjectRef returns a boolean if a field has been set.
func (o *ApiStatus) HasObjectRef() bool {
	if o != nil && o.ObjectRef != nil {
		return true
	}

	return false
}

// SetObjectRef gets a reference to the given ApiObjectRef and assigns it to the ObjectRef field.
func (o *ApiStatus) SetObjectRef(v ApiObjectRef) {
	o.ObjectRef = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ApiStatus) GetResult() ApiStatusResult {
	if o == nil || o.Result == nil {
		var ret ApiStatusResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStatus) GetResultOk() (*ApiStatusResult, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ApiStatus) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given ApiStatusResult and assigns it to the Result field.
func (o *ApiStatus) SetResult(v ApiStatusResult) {
	o.Result = &v
}

func (o ApiStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiVersion != nil {
		toSerialize["api-version"] = o.ApiVersion
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.ObjectRef != nil {
		toSerialize["object-ref"] = o.ObjectRef
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	return json.Marshal(toSerialize)
}

type NullableApiStatus struct {
	value *ApiStatus
	isSet bool
}

func (v NullableApiStatus) Get() *ApiStatus {
	return v.value
}

func (v *NullableApiStatus) Set(val *ApiStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableApiStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableApiStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiStatus(val *ApiStatus) *NullableApiStatus {
	return &NullableApiStatus{value: val, isSet: true}
}

func (v NullableApiStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

