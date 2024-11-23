# AuthRoleSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permissions** | Pointer to [**[]AuthPermission**](AuthPermission.md) | Permissions define actions allowed on resources. A resource can be an API Server object or an arbitrary API endpoint. | [optional] 

## Methods

### NewAuthRoleSpec

`func NewAuthRoleSpec() *AuthRoleSpec`

NewAuthRoleSpec instantiates a new AuthRoleSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthRoleSpecWithDefaults

`func NewAuthRoleSpecWithDefaults() *AuthRoleSpec`

NewAuthRoleSpecWithDefaults instantiates a new AuthRoleSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissions

`func (o *AuthRoleSpec) GetPermissions() []AuthPermission`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *AuthRoleSpec) GetPermissionsOk() (*[]AuthPermission, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *AuthRoleSpec) SetPermissions(v []AuthPermission)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *AuthRoleSpec) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


