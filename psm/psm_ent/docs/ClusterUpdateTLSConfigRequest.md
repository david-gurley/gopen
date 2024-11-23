# ClusterUpdateTLSConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Certs** | Pointer to **string** | Certs is the pem encoded certificate bundle used for API Gateway TLS. | [optional] 
**Key** | Pointer to **string** | Key is the pem encoded private key used for API Gateway TLS. We support RSA or ECDSA. | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**ApiObjectMeta**](apiObjectMeta.md) |  | [optional] 

## Methods

### NewClusterUpdateTLSConfigRequest

`func NewClusterUpdateTLSConfigRequest() *ClusterUpdateTLSConfigRequest`

NewClusterUpdateTLSConfigRequest instantiates a new ClusterUpdateTLSConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterUpdateTLSConfigRequestWithDefaults

`func NewClusterUpdateTLSConfigRequestWithDefaults() *ClusterUpdateTLSConfigRequest`

NewClusterUpdateTLSConfigRequestWithDefaults instantiates a new ClusterUpdateTLSConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ClusterUpdateTLSConfigRequest) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ClusterUpdateTLSConfigRequest) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ClusterUpdateTLSConfigRequest) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ClusterUpdateTLSConfigRequest) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetCerts

`func (o *ClusterUpdateTLSConfigRequest) GetCerts() string`

GetCerts returns the Certs field if non-nil, zero value otherwise.

### GetCertsOk

`func (o *ClusterUpdateTLSConfigRequest) GetCertsOk() (*string, bool)`

GetCertsOk returns a tuple with the Certs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCerts

`func (o *ClusterUpdateTLSConfigRequest) SetCerts(v string)`

SetCerts sets Certs field to given value.

### HasCerts

`func (o *ClusterUpdateTLSConfigRequest) HasCerts() bool`

HasCerts returns a boolean if a field has been set.

### GetKey

`func (o *ClusterUpdateTLSConfigRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ClusterUpdateTLSConfigRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ClusterUpdateTLSConfigRequest) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ClusterUpdateTLSConfigRequest) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetKind

`func (o *ClusterUpdateTLSConfigRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ClusterUpdateTLSConfigRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ClusterUpdateTLSConfigRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ClusterUpdateTLSConfigRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMeta

`func (o *ClusterUpdateTLSConfigRequest) GetMeta() ApiObjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ClusterUpdateTLSConfigRequest) GetMetaOk() (*ApiObjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ClusterUpdateTLSConfigRequest) SetMeta(v ApiObjectMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ClusterUpdateTLSConfigRequest) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


