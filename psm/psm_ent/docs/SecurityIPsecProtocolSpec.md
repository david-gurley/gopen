# SecurityIPsecProtocolSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncryptionTransform** | Pointer to **string** | ESP encryption algorithm. Default is \&quot;aes-256-gcm-128\&quot; (See RFC4106). | [optional] 
**IntegrityTransform** | Pointer to **string** | ESP integrity algorithm. Default is \&quot;NULL\&quot; (must be \&quot;NULL\&quot; if AES-GCM is used for encryption). | [optional] 

## Methods

### NewSecurityIPsecProtocolSpec

`func NewSecurityIPsecProtocolSpec() *SecurityIPsecProtocolSpec`

NewSecurityIPsecProtocolSpec instantiates a new SecurityIPsecProtocolSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityIPsecProtocolSpecWithDefaults

`func NewSecurityIPsecProtocolSpecWithDefaults() *SecurityIPsecProtocolSpec`

NewSecurityIPsecProtocolSpecWithDefaults instantiates a new SecurityIPsecProtocolSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncryptionTransform

`func (o *SecurityIPsecProtocolSpec) GetEncryptionTransform() string`

GetEncryptionTransform returns the EncryptionTransform field if non-nil, zero value otherwise.

### GetEncryptionTransformOk

`func (o *SecurityIPsecProtocolSpec) GetEncryptionTransformOk() (*string, bool)`

GetEncryptionTransformOk returns a tuple with the EncryptionTransform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionTransform

`func (o *SecurityIPsecProtocolSpec) SetEncryptionTransform(v string)`

SetEncryptionTransform sets EncryptionTransform field to given value.

### HasEncryptionTransform

`func (o *SecurityIPsecProtocolSpec) HasEncryptionTransform() bool`

HasEncryptionTransform returns a boolean if a field has been set.

### GetIntegrityTransform

`func (o *SecurityIPsecProtocolSpec) GetIntegrityTransform() string`

GetIntegrityTransform returns the IntegrityTransform field if non-nil, zero value otherwise.

### GetIntegrityTransformOk

`func (o *SecurityIPsecProtocolSpec) GetIntegrityTransformOk() (*string, bool)`

GetIntegrityTransformOk returns a tuple with the IntegrityTransform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrityTransform

`func (o *SecurityIPsecProtocolSpec) SetIntegrityTransform(v string)`

SetIntegrityTransform sets IntegrityTransform field to given value.

### HasIntegrityTransform

`func (o *SecurityIPsecProtocolSpec) HasIntegrityTransform() bool`

HasIntegrityTransform returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


