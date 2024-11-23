# AuthLdapServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TlsOptions** | Pointer to [**AuthTLSOptions**](authTLSOptions.md) |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewAuthLdapServer

`func NewAuthLdapServer() *AuthLdapServer`

NewAuthLdapServer instantiates a new AuthLdapServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthLdapServerWithDefaults

`func NewAuthLdapServerWithDefaults() *AuthLdapServer`

NewAuthLdapServerWithDefaults instantiates a new AuthLdapServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTlsOptions

`func (o *AuthLdapServer) GetTlsOptions() AuthTLSOptions`

GetTlsOptions returns the TlsOptions field if non-nil, zero value otherwise.

### GetTlsOptionsOk

`func (o *AuthLdapServer) GetTlsOptionsOk() (*AuthTLSOptions, bool)`

GetTlsOptionsOk returns a tuple with the TlsOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsOptions

`func (o *AuthLdapServer) SetTlsOptions(v AuthTLSOptions)`

SetTlsOptions sets TlsOptions field to given value.

### HasTlsOptions

`func (o *AuthLdapServer) HasTlsOptions() bool`

HasTlsOptions returns a boolean if a field has been set.

### GetUrl

`func (o *AuthLdapServer) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AuthLdapServer) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AuthLdapServer) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AuthLdapServer) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


