/*
 * Auth API reference
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

// AuthRadiusServerStatus struct for AuthRadiusServerStatus
type AuthRadiusServerStatus struct {
	// Message contains error message in case of failed check or a success message.
	Message *string `json:"message,omitempty"`
	// NasID is a string identifying the NAS(API Gw) originating the Access-Request.
	NasId *string `json:"nas-id,omitempty"`
	// Result indicates if radius check was successful.
	Result *string `json:"result,omitempty"`
	Server *AuthRadiusServer `json:"server,omitempty"`
}

// NewAuthRadiusServerStatus instantiates a new AuthRadiusServerStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthRadiusServerStatus() *AuthRadiusServerStatus {
	this := AuthRadiusServerStatus{}
	var result string = "connect-success"
	this.Result = &result
	return &this
}

// NewAuthRadiusServerStatusWithDefaults instantiates a new AuthRadiusServerStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthRadiusServerStatusWithDefaults() *AuthRadiusServerStatus {
	this := AuthRadiusServerStatus{}
	var result string = "connect-success"
	this.Result = &result
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *AuthRadiusServerStatus) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthRadiusServerStatus) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *AuthRadiusServerStatus) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *AuthRadiusServerStatus) SetMessage(v string) {
	o.Message = &v
}

// GetNasId returns the NasId field value if set, zero value otherwise.
func (o *AuthRadiusServerStatus) GetNasId() string {
	if o == nil || o.NasId == nil {
		var ret string
		return ret
	}
	return *o.NasId
}

// GetNasIdOk returns a tuple with the NasId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthRadiusServerStatus) GetNasIdOk() (*string, bool) {
	if o == nil || o.NasId == nil {
		return nil, false
	}
	return o.NasId, true
}

// HasNasId returns a boolean if a field has been set.
func (o *AuthRadiusServerStatus) HasNasId() bool {
	if o != nil && o.NasId != nil {
		return true
	}

	return false
}

// SetNasId gets a reference to the given string and assigns it to the NasId field.
func (o *AuthRadiusServerStatus) SetNasId(v string) {
	o.NasId = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *AuthRadiusServerStatus) GetResult() string {
	if o == nil || o.Result == nil {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthRadiusServerStatus) GetResultOk() (*string, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *AuthRadiusServerStatus) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *AuthRadiusServerStatus) SetResult(v string) {
	o.Result = &v
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *AuthRadiusServerStatus) GetServer() AuthRadiusServer {
	if o == nil || o.Server == nil {
		var ret AuthRadiusServer
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthRadiusServerStatus) GetServerOk() (*AuthRadiusServer, bool) {
	if o == nil || o.Server == nil {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *AuthRadiusServerStatus) HasServer() bool {
	if o != nil && o.Server != nil {
		return true
	}

	return false
}

// SetServer gets a reference to the given AuthRadiusServer and assigns it to the Server field.
func (o *AuthRadiusServerStatus) SetServer(v AuthRadiusServer) {
	o.Server = &v
}

func (o AuthRadiusServerStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.NasId != nil {
		toSerialize["nas-id"] = o.NasId
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	if o.Server != nil {
		toSerialize["server"] = o.Server
	}
	return json.Marshal(toSerialize)
}

type NullableAuthRadiusServerStatus struct {
	value *AuthRadiusServerStatus
	isSet bool
}

func (v NullableAuthRadiusServerStatus) Get() *AuthRadiusServerStatus {
	return v.value
}

func (v *NullableAuthRadiusServerStatus) Set(val *AuthRadiusServerStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthRadiusServerStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthRadiusServerStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthRadiusServerStatus(val *AuthRadiusServerStatus) *NullableAuthRadiusServerStatus {
	return &NullableAuthRadiusServerStatus{value: val, isSet: true}
}

func (v NullableAuthRadiusServerStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthRadiusServerStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

