# SecurityTrafficEncryptionPolicySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ipsec** | Pointer to [**SecurityIPsecProtocolSpec**](securityIPsecProtocolSpec.md) |  | [optional] 
**KeyRotationIntervalSecs** | Pointer to **int64** | How often the keys should be rotated, in seconds. | [optional] 
**Mode** | Pointer to **string** | Possible values: TLS, IPsec. | [optional] 
**Tls** | Pointer to [**SecurityTLSProtocolSpec**](securityTLSProtocolSpec.md) |  | [optional] 

## Methods

### NewSecurityTrafficEncryptionPolicySpec

`func NewSecurityTrafficEncryptionPolicySpec() *SecurityTrafficEncryptionPolicySpec`

NewSecurityTrafficEncryptionPolicySpec instantiates a new SecurityTrafficEncryptionPolicySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityTrafficEncryptionPolicySpecWithDefaults

`func NewSecurityTrafficEncryptionPolicySpecWithDefaults() *SecurityTrafficEncryptionPolicySpec`

NewSecurityTrafficEncryptionPolicySpecWithDefaults instantiates a new SecurityTrafficEncryptionPolicySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpsec

`func (o *SecurityTrafficEncryptionPolicySpec) GetIpsec() SecurityIPsecProtocolSpec`

GetIpsec returns the Ipsec field if non-nil, zero value otherwise.

### GetIpsecOk

`func (o *SecurityTrafficEncryptionPolicySpec) GetIpsecOk() (*SecurityIPsecProtocolSpec, bool)`

GetIpsecOk returns a tuple with the Ipsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsec

`func (o *SecurityTrafficEncryptionPolicySpec) SetIpsec(v SecurityIPsecProtocolSpec)`

SetIpsec sets Ipsec field to given value.

### HasIpsec

`func (o *SecurityTrafficEncryptionPolicySpec) HasIpsec() bool`

HasIpsec returns a boolean if a field has been set.

### GetKeyRotationIntervalSecs

`func (o *SecurityTrafficEncryptionPolicySpec) GetKeyRotationIntervalSecs() int64`

GetKeyRotationIntervalSecs returns the KeyRotationIntervalSecs field if non-nil, zero value otherwise.

### GetKeyRotationIntervalSecsOk

`func (o *SecurityTrafficEncryptionPolicySpec) GetKeyRotationIntervalSecsOk() (*int64, bool)`

GetKeyRotationIntervalSecsOk returns a tuple with the KeyRotationIntervalSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyRotationIntervalSecs

`func (o *SecurityTrafficEncryptionPolicySpec) SetKeyRotationIntervalSecs(v int64)`

SetKeyRotationIntervalSecs sets KeyRotationIntervalSecs field to given value.

### HasKeyRotationIntervalSecs

`func (o *SecurityTrafficEncryptionPolicySpec) HasKeyRotationIntervalSecs() bool`

HasKeyRotationIntervalSecs returns a boolean if a field has been set.

### GetMode

`func (o *SecurityTrafficEncryptionPolicySpec) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *SecurityTrafficEncryptionPolicySpec) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *SecurityTrafficEncryptionPolicySpec) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *SecurityTrafficEncryptionPolicySpec) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetTls

`func (o *SecurityTrafficEncryptionPolicySpec) GetTls() SecurityTLSProtocolSpec`

GetTls returns the Tls field if non-nil, zero value otherwise.

### GetTlsOk

`func (o *SecurityTrafficEncryptionPolicySpec) GetTlsOk() (*SecurityTLSProtocolSpec, bool)`

GetTlsOk returns a tuple with the Tls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTls

`func (o *SecurityTrafficEncryptionPolicySpec) SetTls(v SecurityTLSProtocolSpec)`

SetTls sets Tls field to given value.

### HasTls

`func (o *SecurityTrafficEncryptionPolicySpec) HasTls() bool`

HasTls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


