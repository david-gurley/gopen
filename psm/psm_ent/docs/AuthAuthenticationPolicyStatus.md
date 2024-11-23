# AuthAuthenticationPolicyStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LdapServers** | Pointer to [**[]AuthLdapServerStatus**](AuthLdapServerStatus.md) |  | [optional] 
**RadiusServers** | Pointer to [**[]AuthRadiusServerStatus**](AuthRadiusServerStatus.md) |  | [optional] 

## Methods

### NewAuthAuthenticationPolicyStatus

`func NewAuthAuthenticationPolicyStatus() *AuthAuthenticationPolicyStatus`

NewAuthAuthenticationPolicyStatus instantiates a new AuthAuthenticationPolicyStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthAuthenticationPolicyStatusWithDefaults

`func NewAuthAuthenticationPolicyStatusWithDefaults() *AuthAuthenticationPolicyStatus`

NewAuthAuthenticationPolicyStatusWithDefaults instantiates a new AuthAuthenticationPolicyStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLdapServers

`func (o *AuthAuthenticationPolicyStatus) GetLdapServers() []AuthLdapServerStatus`

GetLdapServers returns the LdapServers field if non-nil, zero value otherwise.

### GetLdapServersOk

`func (o *AuthAuthenticationPolicyStatus) GetLdapServersOk() (*[]AuthLdapServerStatus, bool)`

GetLdapServersOk returns a tuple with the LdapServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapServers

`func (o *AuthAuthenticationPolicyStatus) SetLdapServers(v []AuthLdapServerStatus)`

SetLdapServers sets LdapServers field to given value.

### HasLdapServers

`func (o *AuthAuthenticationPolicyStatus) HasLdapServers() bool`

HasLdapServers returns a boolean if a field has been set.

### GetRadiusServers

`func (o *AuthAuthenticationPolicyStatus) GetRadiusServers() []AuthRadiusServerStatus`

GetRadiusServers returns the RadiusServers field if non-nil, zero value otherwise.

### GetRadiusServersOk

`func (o *AuthAuthenticationPolicyStatus) GetRadiusServersOk() (*[]AuthRadiusServerStatus, bool)`

GetRadiusServersOk returns a tuple with the RadiusServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusServers

`func (o *AuthAuthenticationPolicyStatus) SetRadiusServers(v []AuthRadiusServerStatus)`

SetRadiusServers sets RadiusServers field to given value.

### HasRadiusServers

`func (o *AuthAuthenticationPolicyStatus) HasRadiusServers() bool`

HasRadiusServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


