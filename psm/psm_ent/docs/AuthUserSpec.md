# AuthUserSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | Must be a valid email. | [optional] 
**Fullname** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** | Password should be atleast 9 characters containing atleast 1 digit, 1 uppercase letter and 1 special character from \&quot;~!@#$%^&amp;*()_+&#x60;-&#x3D;{}|[]\\\\:\\\&quot;&lt;&gt;?,./\&quot;. | [optional] 
**Type** | Pointer to **string** |  | [optional] [default to "local"]

## Methods

### NewAuthUserSpec

`func NewAuthUserSpec() *AuthUserSpec`

NewAuthUserSpec instantiates a new AuthUserSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthUserSpecWithDefaults

`func NewAuthUserSpecWithDefaults() *AuthUserSpec`

NewAuthUserSpecWithDefaults instantiates a new AuthUserSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *AuthUserSpec) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AuthUserSpec) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AuthUserSpec) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AuthUserSpec) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFullname

`func (o *AuthUserSpec) GetFullname() string`

GetFullname returns the Fullname field if non-nil, zero value otherwise.

### GetFullnameOk

`func (o *AuthUserSpec) GetFullnameOk() (*string, bool)`

GetFullnameOk returns a tuple with the Fullname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullname

`func (o *AuthUserSpec) SetFullname(v string)`

SetFullname sets Fullname field to given value.

### HasFullname

`func (o *AuthUserSpec) HasFullname() bool`

HasFullname returns a boolean if a field has been set.

### GetPassword

`func (o *AuthUserSpec) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AuthUserSpec) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AuthUserSpec) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AuthUserSpec) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetType

`func (o *AuthUserSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthUserSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthUserSpec) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AuthUserSpec) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


