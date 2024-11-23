# AuthLdapDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeMapping** | Pointer to [**AuthLdapAttributeMapping**](authLdapAttributeMapping.md) |  | [optional] 
**BaseDn** | Pointer to **string** | The LDAP base DN to be used in a user search. | [optional] 
**BindDn** | Pointer to **string** | The bind DN is the string that Venice uses to log in to the LDAP server. Venice uses this account to validate the remote user attempting to log in. The base DN is the container name and path in the LDAPserver where Venice searches for the remote user account. This is where the password is validated. This contains the user authorization and assigned RBAC roles for use on Venice. Venice requests the attribute from theLDAP server. | [optional] 
**BindPassword** | Pointer to **string** | The password for the LDAP database account specified in the Root DN field. | [optional] 
**Servers** | Pointer to [**[]AuthLdapServer**](AuthLdapServer.md) | Servers is a list that lets you configure multiple LDAP servers for high availability. | [optional] 
**Tag** | Pointer to **string** | Tag to group domains for authentication. | [optional] 

## Methods

### NewAuthLdapDomain

`func NewAuthLdapDomain() *AuthLdapDomain`

NewAuthLdapDomain instantiates a new AuthLdapDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthLdapDomainWithDefaults

`func NewAuthLdapDomainWithDefaults() *AuthLdapDomain`

NewAuthLdapDomainWithDefaults instantiates a new AuthLdapDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeMapping

`func (o *AuthLdapDomain) GetAttributeMapping() AuthLdapAttributeMapping`

GetAttributeMapping returns the AttributeMapping field if non-nil, zero value otherwise.

### GetAttributeMappingOk

`func (o *AuthLdapDomain) GetAttributeMappingOk() (*AuthLdapAttributeMapping, bool)`

GetAttributeMappingOk returns a tuple with the AttributeMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeMapping

`func (o *AuthLdapDomain) SetAttributeMapping(v AuthLdapAttributeMapping)`

SetAttributeMapping sets AttributeMapping field to given value.

### HasAttributeMapping

`func (o *AuthLdapDomain) HasAttributeMapping() bool`

HasAttributeMapping returns a boolean if a field has been set.

### GetBaseDn

`func (o *AuthLdapDomain) GetBaseDn() string`

GetBaseDn returns the BaseDn field if non-nil, zero value otherwise.

### GetBaseDnOk

`func (o *AuthLdapDomain) GetBaseDnOk() (*string, bool)`

GetBaseDnOk returns a tuple with the BaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDn

`func (o *AuthLdapDomain) SetBaseDn(v string)`

SetBaseDn sets BaseDn field to given value.

### HasBaseDn

`func (o *AuthLdapDomain) HasBaseDn() bool`

HasBaseDn returns a boolean if a field has been set.

### GetBindDn

`func (o *AuthLdapDomain) GetBindDn() string`

GetBindDn returns the BindDn field if non-nil, zero value otherwise.

### GetBindDnOk

`func (o *AuthLdapDomain) GetBindDnOk() (*string, bool)`

GetBindDnOk returns a tuple with the BindDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindDn

`func (o *AuthLdapDomain) SetBindDn(v string)`

SetBindDn sets BindDn field to given value.

### HasBindDn

`func (o *AuthLdapDomain) HasBindDn() bool`

HasBindDn returns a boolean if a field has been set.

### GetBindPassword

`func (o *AuthLdapDomain) GetBindPassword() string`

GetBindPassword returns the BindPassword field if non-nil, zero value otherwise.

### GetBindPasswordOk

`func (o *AuthLdapDomain) GetBindPasswordOk() (*string, bool)`

GetBindPasswordOk returns a tuple with the BindPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindPassword

`func (o *AuthLdapDomain) SetBindPassword(v string)`

SetBindPassword sets BindPassword field to given value.

### HasBindPassword

`func (o *AuthLdapDomain) HasBindPassword() bool`

HasBindPassword returns a boolean if a field has been set.

### GetServers

`func (o *AuthLdapDomain) GetServers() []AuthLdapServer`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *AuthLdapDomain) GetServersOk() (*[]AuthLdapServer, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *AuthLdapDomain) SetServers(v []AuthLdapServer)`

SetServers sets Servers field to given value.

### HasServers

`func (o *AuthLdapDomain) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetTag

`func (o *AuthLdapDomain) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *AuthLdapDomain) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *AuthLdapDomain) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *AuthLdapDomain) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


