# ClusterIPConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultGw** | Pointer to **string** | DefaultGW contains the default gateway&#39;s IP address. Should be a valid v4 or v6 IP address. | [optional] 
**DnsServers** | Pointer to **[]string** | DNSServers contains a list of DNS Servers that can be used on DistributedServiceCard. | [optional] 
**IpAddress** | Pointer to **string** | IPAddress contains the Management IP address of the DistributedServiceCard in CIDR format. Should be a valid v4 or v6 CIDR block. | [optional] 

## Methods

### NewClusterIPConfig

`func NewClusterIPConfig() *ClusterIPConfig`

NewClusterIPConfig instantiates a new ClusterIPConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterIPConfigWithDefaults

`func NewClusterIPConfigWithDefaults() *ClusterIPConfig`

NewClusterIPConfigWithDefaults instantiates a new ClusterIPConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultGw

`func (o *ClusterIPConfig) GetDefaultGw() string`

GetDefaultGw returns the DefaultGw field if non-nil, zero value otherwise.

### GetDefaultGwOk

`func (o *ClusterIPConfig) GetDefaultGwOk() (*string, bool)`

GetDefaultGwOk returns a tuple with the DefaultGw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGw

`func (o *ClusterIPConfig) SetDefaultGw(v string)`

SetDefaultGw sets DefaultGw field to given value.

### HasDefaultGw

`func (o *ClusterIPConfig) HasDefaultGw() bool`

HasDefaultGw returns a boolean if a field has been set.

### GetDnsServers

`func (o *ClusterIPConfig) GetDnsServers() []string`

GetDnsServers returns the DnsServers field if non-nil, zero value otherwise.

### GetDnsServersOk

`func (o *ClusterIPConfig) GetDnsServersOk() (*[]string, bool)`

GetDnsServersOk returns a tuple with the DnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServers

`func (o *ClusterIPConfig) SetDnsServers(v []string)`

SetDnsServers sets DnsServers field to given value.

### HasDnsServers

`func (o *ClusterIPConfig) HasDnsServers() bool`

HasDnsServers returns a boolean if a field has been set.

### GetIpAddress

`func (o *ClusterIPConfig) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *ClusterIPConfig) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *ClusterIPConfig) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *ClusterIPConfig) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


