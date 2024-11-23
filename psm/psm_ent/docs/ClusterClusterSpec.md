# ClusterClusterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoAdmitDscs** | Pointer to **bool** | AutoAdmitDSCs when enabled auto-admits DSCs that are validated into Venice Cluster. When it is disabled, DSCs validated by CMD are set to Pending state and it requires Manual approval to be admitted into the cluster. | [optional] 
**Certs** | Pointer to **string** | Certs is the pem encoded certificate bundle used for API Gateway TLS. | [optional] 
**Key** | Pointer to **string** | Key is the pem encoded private key used for API Gateway TLS. We support RSA or ECDSA. | [optional] 
**NtpServers** | Pointer to **[]string** | NTPServers contains the list of NTP servers for the cluster. | [optional] 
**QuorumNodes** | Pointer to **[]string** | QuorumNodes contains the list of hostnames for nodes configured to be quorum nodes in the cluster. | [optional] 
**RecoveryKeys** | Pointer to [**RecoverykeysRecoveryKeys**](recoverykeysRecoveryKeys.md) |  | [optional] 
**VirtualIp** | Pointer to **string** | VirtualIP is the IP address for managing the cluster. It will be hosted by the winner of election between quorum nodes. | [optional] 

## Methods

### NewClusterClusterSpec

`func NewClusterClusterSpec() *ClusterClusterSpec`

NewClusterClusterSpec instantiates a new ClusterClusterSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterClusterSpecWithDefaults

`func NewClusterClusterSpecWithDefaults() *ClusterClusterSpec`

NewClusterClusterSpecWithDefaults instantiates a new ClusterClusterSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoAdmitDscs

`func (o *ClusterClusterSpec) GetAutoAdmitDscs() bool`

GetAutoAdmitDscs returns the AutoAdmitDscs field if non-nil, zero value otherwise.

### GetAutoAdmitDscsOk

`func (o *ClusterClusterSpec) GetAutoAdmitDscsOk() (*bool, bool)`

GetAutoAdmitDscsOk returns a tuple with the AutoAdmitDscs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoAdmitDscs

`func (o *ClusterClusterSpec) SetAutoAdmitDscs(v bool)`

SetAutoAdmitDscs sets AutoAdmitDscs field to given value.

### HasAutoAdmitDscs

`func (o *ClusterClusterSpec) HasAutoAdmitDscs() bool`

HasAutoAdmitDscs returns a boolean if a field has been set.

### GetCerts

`func (o *ClusterClusterSpec) GetCerts() string`

GetCerts returns the Certs field if non-nil, zero value otherwise.

### GetCertsOk

`func (o *ClusterClusterSpec) GetCertsOk() (*string, bool)`

GetCertsOk returns a tuple with the Certs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCerts

`func (o *ClusterClusterSpec) SetCerts(v string)`

SetCerts sets Certs field to given value.

### HasCerts

`func (o *ClusterClusterSpec) HasCerts() bool`

HasCerts returns a boolean if a field has been set.

### GetKey

`func (o *ClusterClusterSpec) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ClusterClusterSpec) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ClusterClusterSpec) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ClusterClusterSpec) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetNtpServers

`func (o *ClusterClusterSpec) GetNtpServers() []string`

GetNtpServers returns the NtpServers field if non-nil, zero value otherwise.

### GetNtpServersOk

`func (o *ClusterClusterSpec) GetNtpServersOk() (*[]string, bool)`

GetNtpServersOk returns a tuple with the NtpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpServers

`func (o *ClusterClusterSpec) SetNtpServers(v []string)`

SetNtpServers sets NtpServers field to given value.

### HasNtpServers

`func (o *ClusterClusterSpec) HasNtpServers() bool`

HasNtpServers returns a boolean if a field has been set.

### GetQuorumNodes

`func (o *ClusterClusterSpec) GetQuorumNodes() []string`

GetQuorumNodes returns the QuorumNodes field if non-nil, zero value otherwise.

### GetQuorumNodesOk

`func (o *ClusterClusterSpec) GetQuorumNodesOk() (*[]string, bool)`

GetQuorumNodesOk returns a tuple with the QuorumNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuorumNodes

`func (o *ClusterClusterSpec) SetQuorumNodes(v []string)`

SetQuorumNodes sets QuorumNodes field to given value.

### HasQuorumNodes

`func (o *ClusterClusterSpec) HasQuorumNodes() bool`

HasQuorumNodes returns a boolean if a field has been set.

### GetRecoveryKeys

`func (o *ClusterClusterSpec) GetRecoveryKeys() RecoverykeysRecoveryKeys`

GetRecoveryKeys returns the RecoveryKeys field if non-nil, zero value otherwise.

### GetRecoveryKeysOk

`func (o *ClusterClusterSpec) GetRecoveryKeysOk() (*RecoverykeysRecoveryKeys, bool)`

GetRecoveryKeysOk returns a tuple with the RecoveryKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryKeys

`func (o *ClusterClusterSpec) SetRecoveryKeys(v RecoverykeysRecoveryKeys)`

SetRecoveryKeys sets RecoveryKeys field to given value.

### HasRecoveryKeys

`func (o *ClusterClusterSpec) HasRecoveryKeys() bool`

HasRecoveryKeys returns a boolean if a field has been set.

### GetVirtualIp

`func (o *ClusterClusterSpec) GetVirtualIp() string`

GetVirtualIp returns the VirtualIp field if non-nil, zero value otherwise.

### GetVirtualIpOk

`func (o *ClusterClusterSpec) GetVirtualIpOk() (*string, bool)`

GetVirtualIpOk returns a tuple with the VirtualIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualIp

`func (o *ClusterClusterSpec) SetVirtualIp(v string)`

SetVirtualIp sets VirtualIp field to given value.

### HasVirtualIp

`func (o *ClusterClusterSpec) HasVirtualIp() bool`

HasVirtualIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


