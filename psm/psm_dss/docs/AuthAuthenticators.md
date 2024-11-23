# AuthAuthenticators

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticatorOrder** | Pointer to **[]string** | Order in which authenticators are applied. If an authenticator returns success, others are skipped. | [optional] 
**Ldap** | Pointer to [**AuthLdap**](authLdap.md) |  | [optional] 
**Local** | Pointer to [**AuthLocal**](authLocal.md) |  | [optional] 
**Radius** | Pointer to [**AuthRadius**](authRadius.md) |  | [optional] 

## Methods

### NewAuthAuthenticators

`func NewAuthAuthenticators() *AuthAuthenticators`

NewAuthAuthenticators instantiates a new AuthAuthenticators object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthAuthenticatorsWithDefaults

`func NewAuthAuthenticatorsWithDefaults() *AuthAuthenticators`

NewAuthAuthenticatorsWithDefaults instantiates a new AuthAuthenticators object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticatorOrder

`func (o *AuthAuthenticators) GetAuthenticatorOrder() []string`

GetAuthenticatorOrder returns the AuthenticatorOrder field if non-nil, zero value otherwise.

### GetAuthenticatorOrderOk

`func (o *AuthAuthenticators) GetAuthenticatorOrderOk() (*[]string, bool)`

GetAuthenticatorOrderOk returns a tuple with the AuthenticatorOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticatorOrder

`func (o *AuthAuthenticators) SetAuthenticatorOrder(v []string)`

SetAuthenticatorOrder sets AuthenticatorOrder field to given value.

### HasAuthenticatorOrder

`func (o *AuthAuthenticators) HasAuthenticatorOrder() bool`

HasAuthenticatorOrder returns a boolean if a field has been set.

### GetLdap

`func (o *AuthAuthenticators) GetLdap() AuthLdap`

GetLdap returns the Ldap field if non-nil, zero value otherwise.

### GetLdapOk

`func (o *AuthAuthenticators) GetLdapOk() (*AuthLdap, bool)`

GetLdapOk returns a tuple with the Ldap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdap

`func (o *AuthAuthenticators) SetLdap(v AuthLdap)`

SetLdap sets Ldap field to given value.

### HasLdap

`func (o *AuthAuthenticators) HasLdap() bool`

HasLdap returns a boolean if a field has been set.

### GetLocal

`func (o *AuthAuthenticators) GetLocal() AuthLocal`

GetLocal returns the Local field if non-nil, zero value otherwise.

### GetLocalOk

`func (o *AuthAuthenticators) GetLocalOk() (*AuthLocal, bool)`

GetLocalOk returns a tuple with the Local field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocal

`func (o *AuthAuthenticators) SetLocal(v AuthLocal)`

SetLocal sets Local field to given value.

### HasLocal

`func (o *AuthAuthenticators) HasLocal() bool`

HasLocal returns a boolean if a field has been set.

### GetRadius

`func (o *AuthAuthenticators) GetRadius() AuthRadius`

GetRadius returns the Radius field if non-nil, zero value otherwise.

### GetRadiusOk

`func (o *AuthAuthenticators) GetRadiusOk() (*AuthRadius, bool)`

GetRadiusOk returns a tuple with the Radius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadius

`func (o *AuthAuthenticators) SetRadius(v AuthRadius)`

SetRadius sets Radius field to given value.

### HasRadius

`func (o *AuthAuthenticators) HasRadius() bool`

HasRadius returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


