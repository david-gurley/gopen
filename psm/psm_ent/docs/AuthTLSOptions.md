# AuthTLSOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerName** | Pointer to **string** | ServerName is used to verify the hostname on the returned certificates unless SkipServerCertVerification is true. | [optional] 
**SkipServerCertVerification** | Pointer to **bool** | SkipServerCertVerification controls whether a client verifies the server&#39;s certificate chain and host name. If SkipServerCertVerification is true, TLS accepts any certificate presented by the server and any host name in that certificate. In this mode, TLS is susceptible to man-in-the-middle attacks. This should be used only for testing. | [optional] 
**StartTls** | Pointer to **bool** | StartTLS determines if ldap connection uses TLS. | [optional] 
**TrustedCerts** | Pointer to **string** | TrustedCerts defines the set of PEM encoded root certificate authorities that will be used when verifying server certificates. | [optional] 

## Methods

### NewAuthTLSOptions

`func NewAuthTLSOptions() *AuthTLSOptions`

NewAuthTLSOptions instantiates a new AuthTLSOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthTLSOptionsWithDefaults

`func NewAuthTLSOptionsWithDefaults() *AuthTLSOptions`

NewAuthTLSOptionsWithDefaults instantiates a new AuthTLSOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerName

`func (o *AuthTLSOptions) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *AuthTLSOptions) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *AuthTLSOptions) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *AuthTLSOptions) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetSkipServerCertVerification

`func (o *AuthTLSOptions) GetSkipServerCertVerification() bool`

GetSkipServerCertVerification returns the SkipServerCertVerification field if non-nil, zero value otherwise.

### GetSkipServerCertVerificationOk

`func (o *AuthTLSOptions) GetSkipServerCertVerificationOk() (*bool, bool)`

GetSkipServerCertVerificationOk returns a tuple with the SkipServerCertVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipServerCertVerification

`func (o *AuthTLSOptions) SetSkipServerCertVerification(v bool)`

SetSkipServerCertVerification sets SkipServerCertVerification field to given value.

### HasSkipServerCertVerification

`func (o *AuthTLSOptions) HasSkipServerCertVerification() bool`

HasSkipServerCertVerification returns a boolean if a field has been set.

### GetStartTls

`func (o *AuthTLSOptions) GetStartTls() bool`

GetStartTls returns the StartTls field if non-nil, zero value otherwise.

### GetStartTlsOk

`func (o *AuthTLSOptions) GetStartTlsOk() (*bool, bool)`

GetStartTlsOk returns a tuple with the StartTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTls

`func (o *AuthTLSOptions) SetStartTls(v bool)`

SetStartTls sets StartTls field to given value.

### HasStartTls

`func (o *AuthTLSOptions) HasStartTls() bool`

HasStartTls returns a boolean if a field has been set.

### GetTrustedCerts

`func (o *AuthTLSOptions) GetTrustedCerts() string`

GetTrustedCerts returns the TrustedCerts field if non-nil, zero value otherwise.

### GetTrustedCertsOk

`func (o *AuthTLSOptions) GetTrustedCertsOk() (*string, bool)`

GetTrustedCertsOk returns a tuple with the TrustedCerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedCerts

`func (o *AuthTLSOptions) SetTrustedCerts(v string)`

SetTrustedCerts sets TrustedCerts field to given value.

### HasTrustedCerts

`func (o *AuthTLSOptions) HasTrustedCerts() bool`

HasTrustedCerts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


