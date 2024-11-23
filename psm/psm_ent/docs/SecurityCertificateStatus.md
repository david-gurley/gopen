# SecurityCertificateStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Validity** | Pointer to **string** | Status of the certificate: \&quot;valid\&quot;, \&quot;invalid\&quot;, \&quot;expired\&quot; \&quot;invalid\&quot; means that the signature of the certificate does not match or there are inconsistencies in the trust chain. | [optional] [default to "unknown"]
**Workloads** | Pointer to **[]string** | The workloads where this certificate has been deployed. | [optional] 

## Methods

### NewSecurityCertificateStatus

`func NewSecurityCertificateStatus() *SecurityCertificateStatus`

NewSecurityCertificateStatus instantiates a new SecurityCertificateStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityCertificateStatusWithDefaults

`func NewSecurityCertificateStatusWithDefaults() *SecurityCertificateStatus`

NewSecurityCertificateStatusWithDefaults instantiates a new SecurityCertificateStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValidity

`func (o *SecurityCertificateStatus) GetValidity() string`

GetValidity returns the Validity field if non-nil, zero value otherwise.

### GetValidityOk

`func (o *SecurityCertificateStatus) GetValidityOk() (*string, bool)`

GetValidityOk returns a tuple with the Validity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidity

`func (o *SecurityCertificateStatus) SetValidity(v string)`

SetValidity sets Validity field to given value.

### HasValidity

`func (o *SecurityCertificateStatus) HasValidity() bool`

HasValidity returns a boolean if a field has been set.

### GetWorkloads

`func (o *SecurityCertificateStatus) GetWorkloads() []string`

GetWorkloads returns the Workloads field if non-nil, zero value otherwise.

### GetWorkloadsOk

`func (o *SecurityCertificateStatus) GetWorkloadsOk() (*[]string, bool)`

GetWorkloadsOk returns a tuple with the Workloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloads

`func (o *SecurityCertificateStatus) SetWorkloads(v []string)`

SetWorkloads sets Workloads field to given value.

### HasWorkloads

`func (o *SecurityCertificateStatus) HasWorkloads() bool`

HasWorkloads returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


