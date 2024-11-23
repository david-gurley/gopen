# SecurityCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**ApiObjectMeta**](apiObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**SecurityCertificateSpec**](securityCertificateSpec.md) |  | [optional] 
**Status** | Pointer to [**SecurityCertificateStatus**](securityCertificateStatus.md) |  | [optional] 

## Methods

### NewSecurityCertificate

`func NewSecurityCertificate() *SecurityCertificate`

NewSecurityCertificate instantiates a new SecurityCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityCertificateWithDefaults

`func NewSecurityCertificateWithDefaults() *SecurityCertificate`

NewSecurityCertificateWithDefaults instantiates a new SecurityCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *SecurityCertificate) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *SecurityCertificate) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *SecurityCertificate) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *SecurityCertificate) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *SecurityCertificate) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SecurityCertificate) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SecurityCertificate) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SecurityCertificate) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMeta

`func (o *SecurityCertificate) GetMeta() ApiObjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SecurityCertificate) GetMetaOk() (*ApiObjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SecurityCertificate) SetMeta(v ApiObjectMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SecurityCertificate) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSpec

`func (o *SecurityCertificate) GetSpec() SecurityCertificateSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *SecurityCertificate) GetSpecOk() (*SecurityCertificateSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *SecurityCertificate) SetSpec(v SecurityCertificateSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *SecurityCertificate) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *SecurityCertificate) GetStatus() SecurityCertificateStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SecurityCertificate) GetStatusOk() (*SecurityCertificateStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SecurityCertificate) SetStatus(v SecurityCertificateStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SecurityCertificate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


