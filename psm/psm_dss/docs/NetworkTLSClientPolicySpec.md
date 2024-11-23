# NetworkTLSClientPolicySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TlsClientAllowedPeerId** | Pointer to **[]string** | Valid DNS names or IP addresses that must appear in the server certificate SubjAltName or Common Name (if SAN is not specified). If not specified, client validates the IP address of the server. | [optional] 
**TlsClientCertificatesSelector** | Pointer to **map[string]string** | A map containing the certificate to use for a set of destinations. The key is a selector for workloads that exist either inside or outside the cluster. It can be based on labels, hostnames or \&quot;IP:port\&quot; pairs. The value is the name of the certificate to use for the selected destinations. The certificates \&quot;usage\&quot; field must contain \&quot;client\&quot;. TODO: replace the first \&quot;string\&quot; type with proper selector type when available. A single \&quot;default\&quot; certificate which matches all destinations is allowed. If a destination matches multiple non-default map keys, an error is returned. If a destination does not match any map key (and there is no default), the outbound connection is initiated without TLS. | [optional] 
**TlsClientTrustRoots** | Pointer to **[]string** | The list of root certificates used to validate a trust chain presented by a server. If the list is empty, all roots certificates in the tenant scope are considered. | [optional] 

## Methods

### NewNetworkTLSClientPolicySpec

`func NewNetworkTLSClientPolicySpec() *NetworkTLSClientPolicySpec`

NewNetworkTLSClientPolicySpec instantiates a new NetworkTLSClientPolicySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkTLSClientPolicySpecWithDefaults

`func NewNetworkTLSClientPolicySpecWithDefaults() *NetworkTLSClientPolicySpec`

NewNetworkTLSClientPolicySpecWithDefaults instantiates a new NetworkTLSClientPolicySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTlsClientAllowedPeerId

`func (o *NetworkTLSClientPolicySpec) GetTlsClientAllowedPeerId() []string`

GetTlsClientAllowedPeerId returns the TlsClientAllowedPeerId field if non-nil, zero value otherwise.

### GetTlsClientAllowedPeerIdOk

`func (o *NetworkTLSClientPolicySpec) GetTlsClientAllowedPeerIdOk() (*[]string, bool)`

GetTlsClientAllowedPeerIdOk returns a tuple with the TlsClientAllowedPeerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsClientAllowedPeerId

`func (o *NetworkTLSClientPolicySpec) SetTlsClientAllowedPeerId(v []string)`

SetTlsClientAllowedPeerId sets TlsClientAllowedPeerId field to given value.

### HasTlsClientAllowedPeerId

`func (o *NetworkTLSClientPolicySpec) HasTlsClientAllowedPeerId() bool`

HasTlsClientAllowedPeerId returns a boolean if a field has been set.

### GetTlsClientCertificatesSelector

`func (o *NetworkTLSClientPolicySpec) GetTlsClientCertificatesSelector() map[string]string`

GetTlsClientCertificatesSelector returns the TlsClientCertificatesSelector field if non-nil, zero value otherwise.

### GetTlsClientCertificatesSelectorOk

`func (o *NetworkTLSClientPolicySpec) GetTlsClientCertificatesSelectorOk() (*map[string]string, bool)`

GetTlsClientCertificatesSelectorOk returns a tuple with the TlsClientCertificatesSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsClientCertificatesSelector

`func (o *NetworkTLSClientPolicySpec) SetTlsClientCertificatesSelector(v map[string]string)`

SetTlsClientCertificatesSelector sets TlsClientCertificatesSelector field to given value.

### HasTlsClientCertificatesSelector

`func (o *NetworkTLSClientPolicySpec) HasTlsClientCertificatesSelector() bool`

HasTlsClientCertificatesSelector returns a boolean if a field has been set.

### GetTlsClientTrustRoots

`func (o *NetworkTLSClientPolicySpec) GetTlsClientTrustRoots() []string`

GetTlsClientTrustRoots returns the TlsClientTrustRoots field if non-nil, zero value otherwise.

### GetTlsClientTrustRootsOk

`func (o *NetworkTLSClientPolicySpec) GetTlsClientTrustRootsOk() (*[]string, bool)`

GetTlsClientTrustRootsOk returns a tuple with the TlsClientTrustRoots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsClientTrustRoots

`func (o *NetworkTLSClientPolicySpec) SetTlsClientTrustRoots(v []string)`

SetTlsClientTrustRoots sets TlsClientTrustRoots field to given value.

### HasTlsClientTrustRoots

`func (o *NetworkTLSClientPolicySpec) HasTlsClientTrustRoots() bool`

HasTlsClientTrustRoots returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


