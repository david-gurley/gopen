# SecurityTLSProtocolSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CipherSuite** | Pointer to **string** | The name of the cipher suite in IANA format default is TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384. | [optional] 
**Version** | Pointer to **string** | TLS version: only supported value at present is 1.2. | [optional] 

## Methods

### NewSecurityTLSProtocolSpec

`func NewSecurityTLSProtocolSpec() *SecurityTLSProtocolSpec`

NewSecurityTLSProtocolSpec instantiates a new SecurityTLSProtocolSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityTLSProtocolSpecWithDefaults

`func NewSecurityTLSProtocolSpecWithDefaults() *SecurityTLSProtocolSpec`

NewSecurityTLSProtocolSpecWithDefaults instantiates a new SecurityTLSProtocolSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCipherSuite

`func (o *SecurityTLSProtocolSpec) GetCipherSuite() string`

GetCipherSuite returns the CipherSuite field if non-nil, zero value otherwise.

### GetCipherSuiteOk

`func (o *SecurityTLSProtocolSpec) GetCipherSuiteOk() (*string, bool)`

GetCipherSuiteOk returns a tuple with the CipherSuite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCipherSuite

`func (o *SecurityTLSProtocolSpec) SetCipherSuite(v string)`

SetCipherSuite sets CipherSuite field to given value.

### HasCipherSuite

`func (o *SecurityTLSProtocolSpec) HasCipherSuite() bool`

HasCipherSuite returns a boolean if a field has been set.

### GetVersion

`func (o *SecurityTLSProtocolSpec) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SecurityTLSProtocolSpec) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SecurityTLSProtocolSpec) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SecurityTLSProtocolSpec) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


