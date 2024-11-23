# AuthLdapAttributeMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | The name of the attribute for storing the users’ e-mail address. This attribute is primarily used for linked Authentication Server Users. It can also be used to identify users by their e-mail address in certificate authentication. | [optional] 
**Fullname** | Pointer to **string** | The name that the server uses for the Name attribute. | [optional] 
**Group** | Pointer to **string** | The name that the server uses for the Group Member Attribute. By default, the attribute is set to member for standard schema, and sgMember for updated schema. | [optional] 
**GroupObjectClass** | Pointer to **string** | GroupObjectClass is the STRUCTURAL object class for group entry in LDAP. It is used as a filter for group search. | [optional] 
**Tenant** | Pointer to **string** | The tenant the server will use for authentication. | [optional] 
**User** | Pointer to **string** | The name that the server uses for the UserID Attribute. | [optional] 
**UserObjectClass** | Pointer to **string** | UserObjectClass is the STRUCTURAL object class for user entry in LDAP. It is used as a filter for user search. | [optional] 

## Methods

### NewAuthLdapAttributeMapping

`func NewAuthLdapAttributeMapping() *AuthLdapAttributeMapping`

NewAuthLdapAttributeMapping instantiates a new AuthLdapAttributeMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthLdapAttributeMappingWithDefaults

`func NewAuthLdapAttributeMappingWithDefaults() *AuthLdapAttributeMapping`

NewAuthLdapAttributeMappingWithDefaults instantiates a new AuthLdapAttributeMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *AuthLdapAttributeMapping) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AuthLdapAttributeMapping) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AuthLdapAttributeMapping) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AuthLdapAttributeMapping) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFullname

`func (o *AuthLdapAttributeMapping) GetFullname() string`

GetFullname returns the Fullname field if non-nil, zero value otherwise.

### GetFullnameOk

`func (o *AuthLdapAttributeMapping) GetFullnameOk() (*string, bool)`

GetFullnameOk returns a tuple with the Fullname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullname

`func (o *AuthLdapAttributeMapping) SetFullname(v string)`

SetFullname sets Fullname field to given value.

### HasFullname

`func (o *AuthLdapAttributeMapping) HasFullname() bool`

HasFullname returns a boolean if a field has been set.

### GetGroup

`func (o *AuthLdapAttributeMapping) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *AuthLdapAttributeMapping) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *AuthLdapAttributeMapping) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *AuthLdapAttributeMapping) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetGroupObjectClass

`func (o *AuthLdapAttributeMapping) GetGroupObjectClass() string`

GetGroupObjectClass returns the GroupObjectClass field if non-nil, zero value otherwise.

### GetGroupObjectClassOk

`func (o *AuthLdapAttributeMapping) GetGroupObjectClassOk() (*string, bool)`

GetGroupObjectClassOk returns a tuple with the GroupObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupObjectClass

`func (o *AuthLdapAttributeMapping) SetGroupObjectClass(v string)`

SetGroupObjectClass sets GroupObjectClass field to given value.

### HasGroupObjectClass

`func (o *AuthLdapAttributeMapping) HasGroupObjectClass() bool`

HasGroupObjectClass returns a boolean if a field has been set.

### GetTenant

`func (o *AuthLdapAttributeMapping) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *AuthLdapAttributeMapping) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *AuthLdapAttributeMapping) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *AuthLdapAttributeMapping) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetUser

`func (o *AuthLdapAttributeMapping) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AuthLdapAttributeMapping) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AuthLdapAttributeMapping) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *AuthLdapAttributeMapping) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserObjectClass

`func (o *AuthLdapAttributeMapping) GetUserObjectClass() string`

GetUserObjectClass returns the UserObjectClass field if non-nil, zero value otherwise.

### GetUserObjectClassOk

`func (o *AuthLdapAttributeMapping) GetUserObjectClassOk() (*string, bool)`

GetUserObjectClassOk returns a tuple with the UserObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserObjectClass

`func (o *AuthLdapAttributeMapping) SetUserObjectClass(v string)`

SetUserObjectClass sets UserObjectClass field to given value.

### HasUserObjectClass

`func (o *AuthLdapAttributeMapping) HasUserObjectClass() bool`

HasUserObjectClass returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


