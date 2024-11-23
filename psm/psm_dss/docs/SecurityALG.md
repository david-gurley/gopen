# SecurityALG

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dns** | Pointer to [**SecurityDns**](securityDns.md) |  | [optional] 
**Ftp** | Pointer to [**SecurityFtp**](securityFtp.md) |  | [optional] 
**Icmp** | Pointer to [**SecurityIcmp**](securityIcmp.md) |  | [optional] 
**Msrpc** | Pointer to [**[]SecurityMsrpc**](SecurityMsrpc.md) |  | [optional] 
**Sunrpc** | Pointer to [**[]SecuritySunrpc**](SecuritySunrpc.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] [default to "icmp"]

## Methods

### NewSecurityALG

`func NewSecurityALG() *SecurityALG`

NewSecurityALG instantiates a new SecurityALG object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityALGWithDefaults

`func NewSecurityALGWithDefaults() *SecurityALG`

NewSecurityALGWithDefaults instantiates a new SecurityALG object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDns

`func (o *SecurityALG) GetDns() SecurityDns`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *SecurityALG) GetDnsOk() (*SecurityDns, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *SecurityALG) SetDns(v SecurityDns)`

SetDns sets Dns field to given value.

### HasDns

`func (o *SecurityALG) HasDns() bool`

HasDns returns a boolean if a field has been set.

### GetFtp

`func (o *SecurityALG) GetFtp() SecurityFtp`

GetFtp returns the Ftp field if non-nil, zero value otherwise.

### GetFtpOk

`func (o *SecurityALG) GetFtpOk() (*SecurityFtp, bool)`

GetFtpOk returns a tuple with the Ftp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtp

`func (o *SecurityALG) SetFtp(v SecurityFtp)`

SetFtp sets Ftp field to given value.

### HasFtp

`func (o *SecurityALG) HasFtp() bool`

HasFtp returns a boolean if a field has been set.

### GetIcmp

`func (o *SecurityALG) GetIcmp() SecurityIcmp`

GetIcmp returns the Icmp field if non-nil, zero value otherwise.

### GetIcmpOk

`func (o *SecurityALG) GetIcmpOk() (*SecurityIcmp, bool)`

GetIcmpOk returns a tuple with the Icmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcmp

`func (o *SecurityALG) SetIcmp(v SecurityIcmp)`

SetIcmp sets Icmp field to given value.

### HasIcmp

`func (o *SecurityALG) HasIcmp() bool`

HasIcmp returns a boolean if a field has been set.

### GetMsrpc

`func (o *SecurityALG) GetMsrpc() []SecurityMsrpc`

GetMsrpc returns the Msrpc field if non-nil, zero value otherwise.

### GetMsrpcOk

`func (o *SecurityALG) GetMsrpcOk() (*[]SecurityMsrpc, bool)`

GetMsrpcOk returns a tuple with the Msrpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsrpc

`func (o *SecurityALG) SetMsrpc(v []SecurityMsrpc)`

SetMsrpc sets Msrpc field to given value.

### HasMsrpc

`func (o *SecurityALG) HasMsrpc() bool`

HasMsrpc returns a boolean if a field has been set.

### GetSunrpc

`func (o *SecurityALG) GetSunrpc() []SecuritySunrpc`

GetSunrpc returns the Sunrpc field if non-nil, zero value otherwise.

### GetSunrpcOk

`func (o *SecurityALG) GetSunrpcOk() (*[]SecuritySunrpc, bool)`

GetSunrpcOk returns a tuple with the Sunrpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSunrpc

`func (o *SecurityALG) SetSunrpc(v []SecuritySunrpc)`

SetSunrpc sets Sunrpc field to given value.

### HasSunrpc

`func (o *SecurityALG) HasSunrpc() bool`

HasSunrpc returns a boolean if a field has been set.

### GetType

`func (o *SecurityALG) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SecurityALG) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SecurityALG) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SecurityALG) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


