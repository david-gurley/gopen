# NetworkTLSServerPolicySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientAuthentication** | Pointer to **string** | Client authentication \&quot;None\&quot; means that server does not request and will not validate a client certificate. \&quot;Mandatory\&quot; means that server requests and validates client certificate. \&quot;Optional\&quot; means that server requests client certificate but proceeds even if client does not present it. Default is \&quot;Mandatory\&quot;. | [optional] [default to "mandatory"]
**TlsServerAllowedPeerId** | Pointer to **[]string** | Valid DNS names or IP addresses that must appear in the client certificate SubjAltName or Common Name (if SAN is not specified). If client auth is enabled and AllowedPeerId is not specified, server accepts any client certificate as long as it is valid  (not expired and with a valid trust chain). | [optional] 
**TlsServerCertificates** | Pointer to **[]string** | List of names of certificates to present to clients. The certificates \&quot;usage\&quot; field must contain \&quot;server\&quot;. If multiple certificates names are provided, system tries to choose the correct one using SNI, otherwise it picks the first one in the list. | [optional] 
**TlsServerTrustRoots** | Pointer to **[]string** | The list of root certificates used to validate a trust chain presented by client. If the list is empty, all roots certificates in the tenant scope are considered. | [optional] 

## Methods

### NewNetworkTLSServerPolicySpec

`func NewNetworkTLSServerPolicySpec() *NetworkTLSServerPolicySpec`

NewNetworkTLSServerPolicySpec instantiates a new NetworkTLSServerPolicySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkTLSServerPolicySpecWithDefaults

`func NewNetworkTLSServerPolicySpecWithDefaults() *NetworkTLSServerPolicySpec`

NewNetworkTLSServerPolicySpecWithDefaults instantiates a new NetworkTLSServerPolicySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientAuthentication

`func (o *NetworkTLSServerPolicySpec) GetClientAuthentication() string`

GetClientAuthentication returns the ClientAuthentication field if non-nil, zero value otherwise.

### GetClientAuthenticationOk

`func (o *NetworkTLSServerPolicySpec) GetClientAuthenticationOk() (*string, bool)`

GetClientAuthenticationOk returns a tuple with the ClientAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAuthentication

`func (o *NetworkTLSServerPolicySpec) SetClientAuthentication(v string)`

SetClientAuthentication sets ClientAuthentication field to given value.

### HasClientAuthentication

`func (o *NetworkTLSServerPolicySpec) HasClientAuthentication() bool`

HasClientAuthentication returns a boolean if a field has been set.

### GetTlsServerAllowedPeerId

`func (o *NetworkTLSServerPolicySpec) GetTlsServerAllowedPeerId() []string`

GetTlsServerAllowedPeerId returns the TlsServerAllowedPeerId field if non-nil, zero value otherwise.

### GetTlsServerAllowedPeerIdOk

`func (o *NetworkTLSServerPolicySpec) GetTlsServerAllowedPeerIdOk() (*[]string, bool)`

GetTlsServerAllowedPeerIdOk returns a tuple with the TlsServerAllowedPeerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsServerAllowedPeerId

`func (o *NetworkTLSServerPolicySpec) SetTlsServerAllowedPeerId(v []string)`

SetTlsServerAllowedPeerId sets TlsServerAllowedPeerId field to given value.

### HasTlsServerAllowedPeerId

`func (o *NetworkTLSServerPolicySpec) HasTlsServerAllowedPeerId() bool`

HasTlsServerAllowedPeerId returns a boolean if a field has been set.

### GetTlsServerCertificates

`func (o *NetworkTLSServerPolicySpec) GetTlsServerCertificates() []string`

GetTlsServerCertificates returns the TlsServerCertificates field if non-nil, zero value otherwise.

### GetTlsServerCertificatesOk

`func (o *NetworkTLSServerPolicySpec) GetTlsServerCertificatesOk() (*[]string, bool)`

GetTlsServerCertificatesOk returns a tuple with the TlsServerCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsServerCertificates

`func (o *NetworkTLSServerPolicySpec) SetTlsServerCertificates(v []string)`

SetTlsServerCertificates sets TlsServerCertificates field to given value.

### HasTlsServerCertificates

`func (o *NetworkTLSServerPolicySpec) HasTlsServerCertificates() bool`

HasTlsServerCertificates returns a boolean if a field has been set.

### GetTlsServerTrustRoots

`func (o *NetworkTLSServerPolicySpec) GetTlsServerTrustRoots() []string`

GetTlsServerTrustRoots returns the TlsServerTrustRoots field if non-nil, zero value otherwise.

### GetTlsServerTrustRootsOk

`func (o *NetworkTLSServerPolicySpec) GetTlsServerTrustRootsOk() (*[]string, bool)`

GetTlsServerTrustRootsOk returns a tuple with the TlsServerTrustRoots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsServerTrustRoots

`func (o *NetworkTLSServerPolicySpec) SetTlsServerTrustRoots(v []string)`

SetTlsServerTrustRoots sets TlsServerTrustRoots field to given value.

### HasTlsServerTrustRoots

`func (o *NetworkTLSServerPolicySpec) HasTlsServerTrustRoots() bool`

HasTlsServerTrustRoots returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


