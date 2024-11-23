# AuthLdapServerStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseDn** | Pointer to **string** | The LDAP base DN to be used in a user search. | [optional] 
**BindDn** | Pointer to **string** | The bind DN is the string that Venice uses to log in to the LDAP server. Venice uses this account to validate the remote user attempting to log in. The base DN is the container name and path in the LDAPserver where Venice searches for the remote user account. This is where the password is validated. This contains the user authorization and assigned RBAC roles for use on Venice. Venice requests the attribute from theLDAP server. | [optional] 
**BindPassword** | Pointer to **string** | The password for the LDAP database account specified in the Root DN field. | [optional] 
**Message** | Pointer to **string** | Message contains error message in case of failed check or a success message. | [optional] 
**Result** | Pointer to **string** | Result indicates if ldap check was successful. | [optional] [default to "connect-success"]
**Server** | Pointer to [**AuthLdapServer**](authLdapServer.md) |  | [optional] 

## Methods

### NewAuthLdapServerStatus

`func NewAuthLdapServerStatus() *AuthLdapServerStatus`

NewAuthLdapServerStatus instantiates a new AuthLdapServerStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthLdapServerStatusWithDefaults

`func NewAuthLdapServerStatusWithDefaults() *AuthLdapServerStatus`

NewAuthLdapServerStatusWithDefaults instantiates a new AuthLdapServerStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseDn

`func (o *AuthLdapServerStatus) GetBaseDn() string`

GetBaseDn returns the BaseDn field if non-nil, zero value otherwise.

### GetBaseDnOk

`func (o *AuthLdapServerStatus) GetBaseDnOk() (*string, bool)`

GetBaseDnOk returns a tuple with the BaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDn

`func (o *AuthLdapServerStatus) SetBaseDn(v string)`

SetBaseDn sets BaseDn field to given value.

### HasBaseDn

`func (o *AuthLdapServerStatus) HasBaseDn() bool`

HasBaseDn returns a boolean if a field has been set.

### GetBindDn

`func (o *AuthLdapServerStatus) GetBindDn() string`

GetBindDn returns the BindDn field if non-nil, zero value otherwise.

### GetBindDnOk

`func (o *AuthLdapServerStatus) GetBindDnOk() (*string, bool)`

GetBindDnOk returns a tuple with the BindDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindDn

`func (o *AuthLdapServerStatus) SetBindDn(v string)`

SetBindDn sets BindDn field to given value.

### HasBindDn

`func (o *AuthLdapServerStatus) HasBindDn() bool`

HasBindDn returns a boolean if a field has been set.

### GetBindPassword

`func (o *AuthLdapServerStatus) GetBindPassword() string`

GetBindPassword returns the BindPassword field if non-nil, zero value otherwise.

### GetBindPasswordOk

`func (o *AuthLdapServerStatus) GetBindPasswordOk() (*string, bool)`

GetBindPasswordOk returns a tuple with the BindPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindPassword

`func (o *AuthLdapServerStatus) SetBindPassword(v string)`

SetBindPassword sets BindPassword field to given value.

### HasBindPassword

`func (o *AuthLdapServerStatus) HasBindPassword() bool`

HasBindPassword returns a boolean if a field has been set.

### GetMessage

`func (o *AuthLdapServerStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AuthLdapServerStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AuthLdapServerStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AuthLdapServerStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetResult

`func (o *AuthLdapServerStatus) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *AuthLdapServerStatus) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *AuthLdapServerStatus) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *AuthLdapServerStatus) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetServer

`func (o *AuthLdapServerStatus) GetServer() AuthLdapServer`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *AuthLdapServerStatus) GetServerOk() (*AuthLdapServer, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *AuthLdapServerStatus) SetServer(v AuthLdapServer)`

SetServer sets Server field to given value.

### HasServer

`func (o *AuthLdapServerStatus) HasServer() bool`

HasServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


