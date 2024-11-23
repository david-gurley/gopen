# MonitoringExternalCred

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthType** | Pointer to **string** | AuthType is the authentication type used in this config. | [optional] [default to "none"]
**BearerToken** | Pointer to **string** | External entity supports bearer tokens for authentication and authorization Token refresh is not supported using OAuth2. | [optional] 
**CaData** | Pointer to **string** | CaData holds PEM-encoded bytes (typically read from a root certificates bundle). CaData is used by client to autheticate external server. This is applicable to all authentication methods. | [optional] 
**CertData** | Pointer to **string** | CertData holds PEM-encoded bytes (typically read from a client certificate file). | [optional] 
**DisableServerAuthentication** | Pointer to **bool** | DisableServerAuthentication flag can be used when a client does not want to authenticate a server. | [optional] 
**KeyData** | Pointer to **string** | KeyData holds PEM-encoded bytes (typically read from a client certificate key file). | [optional] 
**Password** | Pointer to **string** | Password is one time specified, not visibile on read operations Only valid when UserName is defined. | [optional] 
**Username** | Pointer to **string** | UserName is the login id to be used towards the external entity. | [optional] 

## Methods

### NewMonitoringExternalCred

`func NewMonitoringExternalCred() *MonitoringExternalCred`

NewMonitoringExternalCred instantiates a new MonitoringExternalCred object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringExternalCredWithDefaults

`func NewMonitoringExternalCredWithDefaults() *MonitoringExternalCred`

NewMonitoringExternalCredWithDefaults instantiates a new MonitoringExternalCred object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthType

`func (o *MonitoringExternalCred) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *MonitoringExternalCred) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *MonitoringExternalCred) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *MonitoringExternalCred) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetBearerToken

`func (o *MonitoringExternalCred) GetBearerToken() string`

GetBearerToken returns the BearerToken field if non-nil, zero value otherwise.

### GetBearerTokenOk

`func (o *MonitoringExternalCred) GetBearerTokenOk() (*string, bool)`

GetBearerTokenOk returns a tuple with the BearerToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBearerToken

`func (o *MonitoringExternalCred) SetBearerToken(v string)`

SetBearerToken sets BearerToken field to given value.

### HasBearerToken

`func (o *MonitoringExternalCred) HasBearerToken() bool`

HasBearerToken returns a boolean if a field has been set.

### GetCaData

`func (o *MonitoringExternalCred) GetCaData() string`

GetCaData returns the CaData field if non-nil, zero value otherwise.

### GetCaDataOk

`func (o *MonitoringExternalCred) GetCaDataOk() (*string, bool)`

GetCaDataOk returns a tuple with the CaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaData

`func (o *MonitoringExternalCred) SetCaData(v string)`

SetCaData sets CaData field to given value.

### HasCaData

`func (o *MonitoringExternalCred) HasCaData() bool`

HasCaData returns a boolean if a field has been set.

### GetCertData

`func (o *MonitoringExternalCred) GetCertData() string`

GetCertData returns the CertData field if non-nil, zero value otherwise.

### GetCertDataOk

`func (o *MonitoringExternalCred) GetCertDataOk() (*string, bool)`

GetCertDataOk returns a tuple with the CertData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertData

`func (o *MonitoringExternalCred) SetCertData(v string)`

SetCertData sets CertData field to given value.

### HasCertData

`func (o *MonitoringExternalCred) HasCertData() bool`

HasCertData returns a boolean if a field has been set.

### GetDisableServerAuthentication

`func (o *MonitoringExternalCred) GetDisableServerAuthentication() bool`

GetDisableServerAuthentication returns the DisableServerAuthentication field if non-nil, zero value otherwise.

### GetDisableServerAuthenticationOk

`func (o *MonitoringExternalCred) GetDisableServerAuthenticationOk() (*bool, bool)`

GetDisableServerAuthenticationOk returns a tuple with the DisableServerAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableServerAuthentication

`func (o *MonitoringExternalCred) SetDisableServerAuthentication(v bool)`

SetDisableServerAuthentication sets DisableServerAuthentication field to given value.

### HasDisableServerAuthentication

`func (o *MonitoringExternalCred) HasDisableServerAuthentication() bool`

HasDisableServerAuthentication returns a boolean if a field has been set.

### GetKeyData

`func (o *MonitoringExternalCred) GetKeyData() string`

GetKeyData returns the KeyData field if non-nil, zero value otherwise.

### GetKeyDataOk

`func (o *MonitoringExternalCred) GetKeyDataOk() (*string, bool)`

GetKeyDataOk returns a tuple with the KeyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyData

`func (o *MonitoringExternalCred) SetKeyData(v string)`

SetKeyData sets KeyData field to given value.

### HasKeyData

`func (o *MonitoringExternalCred) HasKeyData() bool`

HasKeyData returns a boolean if a field has been set.

### GetPassword

`func (o *MonitoringExternalCred) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *MonitoringExternalCred) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *MonitoringExternalCred) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *MonitoringExternalCred) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUsername

`func (o *MonitoringExternalCred) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *MonitoringExternalCred) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *MonitoringExternalCred) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *MonitoringExternalCred) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


