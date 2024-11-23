# AuthRadiusServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthMethod** | Pointer to **string** | AuthMethod is authentication method to use with the RADIUS server. | [optional] [default to "pap"]
**Secret** | Pointer to **string** | Secret is the shared secret between API Gw and RADIUS server. | [optional] 
**TrustedCerts** | Pointer to **string** | TrustedCerts defines the set of PEM encoded root certificate authorities that will be used when verifying server certificates. It is used in PEAP and EAP_TTLS auth methods. | [optional] 
**Url** | Pointer to **string** | &lt;IP address&gt;:&lt;Port&gt; of the RADIUS server. | [optional] 

## Methods

### NewAuthRadiusServer

`func NewAuthRadiusServer() *AuthRadiusServer`

NewAuthRadiusServer instantiates a new AuthRadiusServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthRadiusServerWithDefaults

`func NewAuthRadiusServerWithDefaults() *AuthRadiusServer`

NewAuthRadiusServerWithDefaults instantiates a new AuthRadiusServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthMethod

`func (o *AuthRadiusServer) GetAuthMethod() string`

GetAuthMethod returns the AuthMethod field if non-nil, zero value otherwise.

### GetAuthMethodOk

`func (o *AuthRadiusServer) GetAuthMethodOk() (*string, bool)`

GetAuthMethodOk returns a tuple with the AuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethod

`func (o *AuthRadiusServer) SetAuthMethod(v string)`

SetAuthMethod sets AuthMethod field to given value.

### HasAuthMethod

`func (o *AuthRadiusServer) HasAuthMethod() bool`

HasAuthMethod returns a boolean if a field has been set.

### GetSecret

`func (o *AuthRadiusServer) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *AuthRadiusServer) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *AuthRadiusServer) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *AuthRadiusServer) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetTrustedCerts

`func (o *AuthRadiusServer) GetTrustedCerts() string`

GetTrustedCerts returns the TrustedCerts field if non-nil, zero value otherwise.

### GetTrustedCertsOk

`func (o *AuthRadiusServer) GetTrustedCertsOk() (*string, bool)`

GetTrustedCertsOk returns a tuple with the TrustedCerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedCerts

`func (o *AuthRadiusServer) SetTrustedCerts(v string)`

SetTrustedCerts sets TrustedCerts field to given value.

### HasTrustedCerts

`func (o *AuthRadiusServer) HasTrustedCerts() bool`

HasTrustedCerts returns a boolean if a field has been set.

### GetUrl

`func (o *AuthRadiusServer) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AuthRadiusServer) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AuthRadiusServer) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AuthRadiusServer) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


