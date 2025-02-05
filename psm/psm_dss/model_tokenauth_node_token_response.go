/*
 * Tokenauth API reference
 *
 * APIs to generate node auth tokens  
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package psm_dss

import (
	"encoding/json"
)

// TokenauthNodeTokenResponse struct for TokenauthNodeTokenResponse
type TokenauthNodeTokenResponse struct {
	Token *string `json:"Token,omitempty"`
}

// NewTokenauthNodeTokenResponse instantiates a new TokenauthNodeTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenauthNodeTokenResponse() *TokenauthNodeTokenResponse {
	this := TokenauthNodeTokenResponse{}
	return &this
}

// NewTokenauthNodeTokenResponseWithDefaults instantiates a new TokenauthNodeTokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenauthNodeTokenResponseWithDefaults() *TokenauthNodeTokenResponse {
	this := TokenauthNodeTokenResponse{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *TokenauthNodeTokenResponse) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenauthNodeTokenResponse) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *TokenauthNodeTokenResponse) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *TokenauthNodeTokenResponse) SetToken(v string) {
	o.Token = &v
}

func (o TokenauthNodeTokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Token != nil {
		toSerialize["Token"] = o.Token
	}
	return json.Marshal(toSerialize)
}

type NullableTokenauthNodeTokenResponse struct {
	value *TokenauthNodeTokenResponse
	isSet bool
}

func (v NullableTokenauthNodeTokenResponse) Get() *TokenauthNodeTokenResponse {
	return v.value
}

func (v *NullableTokenauthNodeTokenResponse) Set(val *TokenauthNodeTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenauthNodeTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenauthNodeTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenauthNodeTokenResponse(val *TokenauthNodeTokenResponse) *NullableTokenauthNodeTokenResponse {
	return &NullableTokenauthNodeTokenResponse{value: val, isSet: true}
}

func (v NullableTokenauthNodeTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenauthNodeTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


