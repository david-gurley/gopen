# AuthAuthenticationPolicySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authenticators** | Pointer to [**AuthAuthenticators**](authAuthenticators.md) |  | [optional] 
**Secret** | Pointer to **string** | Secret used to sign JWT token. | [optional] 
**TokenExpiry** | Pointer to **string** | TokenExpiry is time duration after which JWT token expires. Default is 6 days. A duration string is a sequence of decimal number and a unit suffix, such as \&quot;300ms\&quot; or \&quot;2h45m\&quot;. Valid time units are \&quot;ns\&quot;, \&quot;us\&quot; (or \&quot;µs\&quot;), \&quot;ms\&quot;, \&quot;s\&quot;, \&quot;m\&quot;, \&quot;h\&quot;. Should be a valid time duration. | [optional] [default to "144h"]

## Methods

### NewAuthAuthenticationPolicySpec

`func NewAuthAuthenticationPolicySpec() *AuthAuthenticationPolicySpec`

NewAuthAuthenticationPolicySpec instantiates a new AuthAuthenticationPolicySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthAuthenticationPolicySpecWithDefaults

`func NewAuthAuthenticationPolicySpecWithDefaults() *AuthAuthenticationPolicySpec`

NewAuthAuthenticationPolicySpecWithDefaults instantiates a new AuthAuthenticationPolicySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticators

`func (o *AuthAuthenticationPolicySpec) GetAuthenticators() AuthAuthenticators`

GetAuthenticators returns the Authenticators field if non-nil, zero value otherwise.

### GetAuthenticatorsOk

`func (o *AuthAuthenticationPolicySpec) GetAuthenticatorsOk() (*AuthAuthenticators, bool)`

GetAuthenticatorsOk returns a tuple with the Authenticators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticators

`func (o *AuthAuthenticationPolicySpec) SetAuthenticators(v AuthAuthenticators)`

SetAuthenticators sets Authenticators field to given value.

### HasAuthenticators

`func (o *AuthAuthenticationPolicySpec) HasAuthenticators() bool`

HasAuthenticators returns a boolean if a field has been set.

### GetSecret

`func (o *AuthAuthenticationPolicySpec) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *AuthAuthenticationPolicySpec) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *AuthAuthenticationPolicySpec) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *AuthAuthenticationPolicySpec) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetTokenExpiry

`func (o *AuthAuthenticationPolicySpec) GetTokenExpiry() string`

GetTokenExpiry returns the TokenExpiry field if non-nil, zero value otherwise.

### GetTokenExpiryOk

`func (o *AuthAuthenticationPolicySpec) GetTokenExpiryOk() (*string, bool)`

GetTokenExpiryOk returns a tuple with the TokenExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExpiry

`func (o *AuthAuthenticationPolicySpec) SetTokenExpiry(v string)`

SetTokenExpiry sets TokenExpiry field to given value.

### HasTokenExpiry

`func (o *AuthAuthenticationPolicySpec) HasTokenExpiry() bool`

HasTokenExpiry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


