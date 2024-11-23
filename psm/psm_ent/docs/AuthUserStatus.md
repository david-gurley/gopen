# AuthUserStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessReview** | Pointer to [**[]AuthOperationStatus**](AuthOperationStatus.md) | Authorization information about requested operations. | [optional] 
**Authenticators** | Pointer to **[]string** | Authenticators used for last successful login. | [optional] 
**LastLogin** | Pointer to **time.Time** | Last login time. | [optional] 
**LastPasswordChange** | Pointer to **time.Time** | Last password change time for local user. | [optional] 
**Roles** | Pointer to **[]string** | Roles assigned to user. | [optional] 
**UserGroups** | Pointer to **[]string** | Groups that external user belongs to. | [optional] 

## Methods

### NewAuthUserStatus

`func NewAuthUserStatus() *AuthUserStatus`

NewAuthUserStatus instantiates a new AuthUserStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthUserStatusWithDefaults

`func NewAuthUserStatusWithDefaults() *AuthUserStatus`

NewAuthUserStatusWithDefaults instantiates a new AuthUserStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessReview

`func (o *AuthUserStatus) GetAccessReview() []AuthOperationStatus`

GetAccessReview returns the AccessReview field if non-nil, zero value otherwise.

### GetAccessReviewOk

`func (o *AuthUserStatus) GetAccessReviewOk() (*[]AuthOperationStatus, bool)`

GetAccessReviewOk returns a tuple with the AccessReview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessReview

`func (o *AuthUserStatus) SetAccessReview(v []AuthOperationStatus)`

SetAccessReview sets AccessReview field to given value.

### HasAccessReview

`func (o *AuthUserStatus) HasAccessReview() bool`

HasAccessReview returns a boolean if a field has been set.

### GetAuthenticators

`func (o *AuthUserStatus) GetAuthenticators() []string`

GetAuthenticators returns the Authenticators field if non-nil, zero value otherwise.

### GetAuthenticatorsOk

`func (o *AuthUserStatus) GetAuthenticatorsOk() (*[]string, bool)`

GetAuthenticatorsOk returns a tuple with the Authenticators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticators

`func (o *AuthUserStatus) SetAuthenticators(v []string)`

SetAuthenticators sets Authenticators field to given value.

### HasAuthenticators

`func (o *AuthUserStatus) HasAuthenticators() bool`

HasAuthenticators returns a boolean if a field has been set.

### GetLastLogin

`func (o *AuthUserStatus) GetLastLogin() time.Time`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *AuthUserStatus) GetLastLoginOk() (*time.Time, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *AuthUserStatus) SetLastLogin(v time.Time)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *AuthUserStatus) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### GetLastPasswordChange

`func (o *AuthUserStatus) GetLastPasswordChange() time.Time`

GetLastPasswordChange returns the LastPasswordChange field if non-nil, zero value otherwise.

### GetLastPasswordChangeOk

`func (o *AuthUserStatus) GetLastPasswordChangeOk() (*time.Time, bool)`

GetLastPasswordChangeOk returns a tuple with the LastPasswordChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPasswordChange

`func (o *AuthUserStatus) SetLastPasswordChange(v time.Time)`

SetLastPasswordChange sets LastPasswordChange field to given value.

### HasLastPasswordChange

`func (o *AuthUserStatus) HasLastPasswordChange() bool`

HasLastPasswordChange returns a boolean if a field has been set.

### GetRoles

`func (o *AuthUserStatus) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *AuthUserStatus) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *AuthUserStatus) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *AuthUserStatus) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetUserGroups

`func (o *AuthUserStatus) GetUserGroups() []string`

GetUserGroups returns the UserGroups field if non-nil, zero value otherwise.

### GetUserGroupsOk

`func (o *AuthUserStatus) GetUserGroupsOk() (*[]string, bool)`

GetUserGroupsOk returns a tuple with the UserGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroups

`func (o *AuthUserStatus) SetUserGroups(v []string)`

SetUserGroups sets UserGroups field to given value.

### HasUserGroups

`func (o *AuthUserStatus) HasUserGroups() bool`

HasUserGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


