# AuthRoleBindingSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to **string** |  | [optional] 
**UserGroups** | Pointer to **[]string** |  | [optional] 
**Users** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAuthRoleBindingSpec

`func NewAuthRoleBindingSpec() *AuthRoleBindingSpec`

NewAuthRoleBindingSpec instantiates a new AuthRoleBindingSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthRoleBindingSpecWithDefaults

`func NewAuthRoleBindingSpecWithDefaults() *AuthRoleBindingSpec`

NewAuthRoleBindingSpecWithDefaults instantiates a new AuthRoleBindingSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *AuthRoleBindingSpec) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AuthRoleBindingSpec) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AuthRoleBindingSpec) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *AuthRoleBindingSpec) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetUserGroups

`func (o *AuthRoleBindingSpec) GetUserGroups() []string`

GetUserGroups returns the UserGroups field if non-nil, zero value otherwise.

### GetUserGroupsOk

`func (o *AuthRoleBindingSpec) GetUserGroupsOk() (*[]string, bool)`

GetUserGroupsOk returns a tuple with the UserGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroups

`func (o *AuthRoleBindingSpec) SetUserGroups(v []string)`

SetUserGroups sets UserGroups field to given value.

### HasUserGroups

`func (o *AuthRoleBindingSpec) HasUserGroups() bool`

HasUserGroups returns a boolean if a field has been set.

### GetUsers

`func (o *AuthRoleBindingSpec) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *AuthRoleBindingSpec) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *AuthRoleBindingSpec) SetUsers(v []string)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *AuthRoleBindingSpec) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


