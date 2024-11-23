# SecurityCertificateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** | Body of the certificate in PEM encoding. | [optional] 
**Description** | Pointer to **string** | Description of the purpose of this certificate. | [optional] 
**TrustChain** | Pointer to **string** | Trust chain of the certificate in PEM encoding. These certificates are treated opaquely. We do not process them in any way other than decoding them for informational purposes. | [optional] 
**Usages** | Pointer to **[]string** | Usage can be \&quot;client\&quot;, \&quot;server\&quot; or \&quot;trust-root\&quot; in any combination. A \&quot;server\&quot; certificate is used by a server to authenticate itself to the client A \&quot;client\&quot; certificate is used by a client to authenticate itself to a server A \&quot;trust-root\&quot; certificate is self-signed and is only used to validate certificates presented by peers. \&quot;client\&quot; and \&quot;server\&quot; certificates are always accompanied by a private key, whereas \&quot;trust-root\&quot;-only certificates are not. | [optional] 

## Methods

### NewSecurityCertificateSpec

`func NewSecurityCertificateSpec() *SecurityCertificateSpec`

NewSecurityCertificateSpec instantiates a new SecurityCertificateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityCertificateSpecWithDefaults

`func NewSecurityCertificateSpecWithDefaults() *SecurityCertificateSpec`

NewSecurityCertificateSpecWithDefaults instantiates a new SecurityCertificateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *SecurityCertificateSpec) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *SecurityCertificateSpec) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *SecurityCertificateSpec) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *SecurityCertificateSpec) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetDescription

`func (o *SecurityCertificateSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SecurityCertificateSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SecurityCertificateSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SecurityCertificateSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTrustChain

`func (o *SecurityCertificateSpec) GetTrustChain() string`

GetTrustChain returns the TrustChain field if non-nil, zero value otherwise.

### GetTrustChainOk

`func (o *SecurityCertificateSpec) GetTrustChainOk() (*string, bool)`

GetTrustChainOk returns a tuple with the TrustChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustChain

`func (o *SecurityCertificateSpec) SetTrustChain(v string)`

SetTrustChain sets TrustChain field to given value.

### HasTrustChain

`func (o *SecurityCertificateSpec) HasTrustChain() bool`

HasTrustChain returns a boolean if a field has been set.

### GetUsages

`func (o *SecurityCertificateSpec) GetUsages() []string`

GetUsages returns the Usages field if non-nil, zero value otherwise.

### GetUsagesOk

`func (o *SecurityCertificateSpec) GetUsagesOk() (*[]string, bool)`

GetUsagesOk returns a tuple with the Usages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsages

`func (o *SecurityCertificateSpec) SetUsages(v []string)`

SetUsages sets Usages field to given value.

### HasUsages

`func (o *SecurityCertificateSpec) HasUsages() bool`

HasUsages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


