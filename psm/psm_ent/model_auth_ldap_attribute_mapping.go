/*
 * Auth API reference
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

// AuthLdapAttributeMapping struct for AuthLdapAttributeMapping
type AuthLdapAttributeMapping struct {
	// The name of the attribute for storing the users’ e-mail address. This attribute is primarily used for linked Authentication Server Users. It can also be used to identify users by their e-mail address in certificate authentication.
	Email *string `json:"email,omitempty"`
	// The name that the server uses for the Name attribute.
	Fullname *string `json:"fullname,omitempty"`
	// The name that the server uses for the Group Member Attribute. By default, the attribute is set to member for standard schema, and sgMember for updated schema.
	Group *string `json:"group,omitempty"`
	// GroupObjectClass is the STRUCTURAL object class for group entry in LDAP. It is used as a filter for group search.
	GroupObjectClass *string `json:"group-object-class,omitempty"`
	// The tenant the server will use for authentication.
	Tenant *string `json:"tenant,omitempty"`
	// The name that the server uses for the UserID Attribute.
	User *string `json:"user,omitempty"`
	// UserObjectClass is the STRUCTURAL object class for user entry in LDAP. It is used as a filter for user search.
	UserObjectClass *string `json:"user-object-class,omitempty"`
}

// NewAuthLdapAttributeMapping instantiates a new AuthLdapAttributeMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthLdapAttributeMapping() *AuthLdapAttributeMapping {
	this := AuthLdapAttributeMapping{}
	return &this
}

// NewAuthLdapAttributeMappingWithDefaults instantiates a new AuthLdapAttributeMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthLdapAttributeMappingWithDefaults() *AuthLdapAttributeMapping {
	this := AuthLdapAttributeMapping{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *AuthLdapAttributeMapping) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthLdapAttributeMapping) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *AuthLdapAttributeMapping) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *AuthLdapAttributeMapping) SetEmail(v string) {
	o.Email = &v
}

// GetFullname returns the Fullname field value if set, zero value otherwise.
func (o *AuthLdapAttributeMapping) GetFullname() string {
	if o == nil || o.Fullname == nil {
		var ret string
		return ret
	}
	return *o.Fullname
}

// GetFullnameOk returns a tuple with the Fullname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthLdapAttributeMapping) GetFullnameOk() (*string, bool) {
	if o == nil || o.Fullname == nil {
		return nil, false
	}
	return o.Fullname, true
}

// HasFullname returns a boolean if a field has been set.
func (o *AuthLdapAttributeMapping) HasFullname() bool {
	if o != nil && o.Fullname != nil {
		return true
	}

	return false
}

// SetFullname gets a reference to the given string and assigns it to the Fullname field.
func (o *AuthLdapAttributeMapping) SetFullname(v string) {
	o.Fullname = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *AuthLdapAttributeMapping) GetGroup() string {
	if o == nil || o.Group == nil {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthLdapAttributeMapping) GetGroupOk() (*string, bool) {
	if o == nil || o.Group == nil {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *AuthLdapAttributeMapping) HasGroup() bool {
	if o != nil && o.Group != nil {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *AuthLdapAttributeMapping) SetGroup(v string) {
	o.Group = &v
}

// GetGroupObjectClass returns the GroupObjectClass field value if set, zero value otherwise.
func (o *AuthLdapAttributeMapping) GetGroupObjectClass() string {
	if o == nil || o.GroupObjectClass == nil {
		var ret string
		return ret
	}
	return *o.GroupObjectClass
}

// GetGroupObjectClassOk returns a tuple with the GroupObjectClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthLdapAttributeMapping) GetGroupObjectClassOk() (*string, bool) {
	if o == nil || o.GroupObjectClass == nil {
		return nil, false
	}
	return o.GroupObjectClass, true
}

// HasGroupObjectClass returns a boolean if a field has been set.
func (o *AuthLdapAttributeMapping) HasGroupObjectClass() bool {
	if o != nil && o.GroupObjectClass != nil {
		return true
	}

	return false
}

// SetGroupObjectClass gets a reference to the given string and assigns it to the GroupObjectClass field.
func (o *AuthLdapAttributeMapping) SetGroupObjectClass(v string) {
	o.GroupObjectClass = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *AuthLdapAttributeMapping) GetTenant() string {
	if o == nil || o.Tenant == nil {
		var ret string
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthLdapAttributeMapping) GetTenantOk() (*string, bool) {
	if o == nil || o.Tenant == nil {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *AuthLdapAttributeMapping) HasTenant() bool {
	if o != nil && o.Tenant != nil {
		return true
	}

	return false
}

// SetTenant gets a reference to the given string and assigns it to the Tenant field.
func (o *AuthLdapAttributeMapping) SetTenant(v string) {
	o.Tenant = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *AuthLdapAttributeMapping) GetUser() string {
	if o == nil || o.User == nil {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthLdapAttributeMapping) GetUserOk() (*string, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *AuthLdapAttributeMapping) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *AuthLdapAttributeMapping) SetUser(v string) {
	o.User = &v
}

// GetUserObjectClass returns the UserObjectClass field value if set, zero value otherwise.
func (o *AuthLdapAttributeMapping) GetUserObjectClass() string {
	if o == nil || o.UserObjectClass == nil {
		var ret string
		return ret
	}
	return *o.UserObjectClass
}

// GetUserObjectClassOk returns a tuple with the UserObjectClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthLdapAttributeMapping) GetUserObjectClassOk() (*string, bool) {
	if o == nil || o.UserObjectClass == nil {
		return nil, false
	}
	return o.UserObjectClass, true
}

// HasUserObjectClass returns a boolean if a field has been set.
func (o *AuthLdapAttributeMapping) HasUserObjectClass() bool {
	if o != nil && o.UserObjectClass != nil {
		return true
	}

	return false
}

// SetUserObjectClass gets a reference to the given string and assigns it to the UserObjectClass field.
func (o *AuthLdapAttributeMapping) SetUserObjectClass(v string) {
	o.UserObjectClass = &v
}

func (o AuthLdapAttributeMapping) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Fullname != nil {
		toSerialize["fullname"] = o.Fullname
	}
	if o.Group != nil {
		toSerialize["group"] = o.Group
	}
	if o.GroupObjectClass != nil {
		toSerialize["group-object-class"] = o.GroupObjectClass
	}
	if o.Tenant != nil {
		toSerialize["tenant"] = o.Tenant
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.UserObjectClass != nil {
		toSerialize["user-object-class"] = o.UserObjectClass
	}
	return json.Marshal(toSerialize)
}

type NullableAuthLdapAttributeMapping struct {
	value *AuthLdapAttributeMapping
	isSet bool
}

func (v NullableAuthLdapAttributeMapping) Get() *AuthLdapAttributeMapping {
	return v.value
}

func (v *NullableAuthLdapAttributeMapping) Set(val *AuthLdapAttributeMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthLdapAttributeMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthLdapAttributeMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthLdapAttributeMapping(val *AuthLdapAttributeMapping) *NullableAuthLdapAttributeMapping {
	return &NullableAuthLdapAttributeMapping{value: val, isSet: true}
}

func (v NullableAuthLdapAttributeMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthLdapAttributeMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


