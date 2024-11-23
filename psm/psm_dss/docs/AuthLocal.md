# AuthLocal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedFailedLoginAttempts** | Pointer to **int32** | Failed login attempts after which user is locked. Value should be at least 1. | [optional] [default to 10]
**FailedLoginAttemptsDuration** | Pointer to **string** | FailedLoginAttemptsDuration is time duration after number of failed login attempts are cleared for a user. Default is 15 min. Minimum value is 1 sec. A duration string is a sequence of decimal number and a unit suffix, such as \&quot;300ms\&quot; or \&quot;2h45m\&quot;. Valid time units are \&quot;ns\&quot;, \&quot;us\&quot; (or \&quot;µs\&quot;), \&quot;ms\&quot;, \&quot;s\&quot;, \&quot;m\&quot;, \&quot;h\&quot;. Should be a valid time duration of at least 1s. | [optional] [default to "15m"]
**PasswordLength** | Pointer to **int32** | Minimum password length. Value should be at least 3. | [optional] [default to 9]

## Methods

### NewAuthLocal

`func NewAuthLocal() *AuthLocal`

NewAuthLocal instantiates a new AuthLocal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthLocalWithDefaults

`func NewAuthLocalWithDefaults() *AuthLocal`

NewAuthLocalWithDefaults instantiates a new AuthLocal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedFailedLoginAttempts

`func (o *AuthLocal) GetAllowedFailedLoginAttempts() int32`

GetAllowedFailedLoginAttempts returns the AllowedFailedLoginAttempts field if non-nil, zero value otherwise.

### GetAllowedFailedLoginAttemptsOk

`func (o *AuthLocal) GetAllowedFailedLoginAttemptsOk() (*int32, bool)`

GetAllowedFailedLoginAttemptsOk returns a tuple with the AllowedFailedLoginAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedFailedLoginAttempts

`func (o *AuthLocal) SetAllowedFailedLoginAttempts(v int32)`

SetAllowedFailedLoginAttempts sets AllowedFailedLoginAttempts field to given value.

### HasAllowedFailedLoginAttempts

`func (o *AuthLocal) HasAllowedFailedLoginAttempts() bool`

HasAllowedFailedLoginAttempts returns a boolean if a field has been set.

### GetFailedLoginAttemptsDuration

`func (o *AuthLocal) GetFailedLoginAttemptsDuration() string`

GetFailedLoginAttemptsDuration returns the FailedLoginAttemptsDuration field if non-nil, zero value otherwise.

### GetFailedLoginAttemptsDurationOk

`func (o *AuthLocal) GetFailedLoginAttemptsDurationOk() (*string, bool)`

GetFailedLoginAttemptsDurationOk returns a tuple with the FailedLoginAttemptsDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedLoginAttemptsDuration

`func (o *AuthLocal) SetFailedLoginAttemptsDuration(v string)`

SetFailedLoginAttemptsDuration sets FailedLoginAttemptsDuration field to given value.

### HasFailedLoginAttemptsDuration

`func (o *AuthLocal) HasFailedLoginAttemptsDuration() bool`

HasFailedLoginAttemptsDuration returns a boolean if a field has been set.

### GetPasswordLength

`func (o *AuthLocal) GetPasswordLength() int32`

GetPasswordLength returns the PasswordLength field if non-nil, zero value otherwise.

### GetPasswordLengthOk

`func (o *AuthLocal) GetPasswordLengthOk() (*int32, bool)`

GetPasswordLengthOk returns a tuple with the PasswordLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLength

`func (o *AuthLocal) SetPasswordLength(v int32)`

SetPasswordLength sets PasswordLength field to given value.

### HasPasswordLength

`func (o *AuthLocal) HasPasswordLength() bool`

HasPasswordLength returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


