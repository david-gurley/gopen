# AuthLdap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domains** | Pointer to [**[]AuthLdapDomain**](AuthLdapDomain.md) |  | [optional] 

## Methods

### NewAuthLdap

`func NewAuthLdap() *AuthLdap`

NewAuthLdap instantiates a new AuthLdap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthLdapWithDefaults

`func NewAuthLdapWithDefaults() *AuthLdap`

NewAuthLdapWithDefaults instantiates a new AuthLdap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomains

`func (o *AuthLdap) GetDomains() []AuthLdapDomain`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *AuthLdap) GetDomainsOk() (*[]AuthLdapDomain, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *AuthLdap) SetDomains(v []AuthLdapDomain)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *AuthLdap) HasDomains() bool`

HasDomains returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

