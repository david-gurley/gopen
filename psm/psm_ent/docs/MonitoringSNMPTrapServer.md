# MonitoringSNMPTrapServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthConfig** | Pointer to [**MonitoringAuthConfig**](monitoringAuthConfig.md) |  | [optional] 
**CommunityOrUser** | Pointer to **string** | CommunityOrUser contains community string for v2c, user for v3. | [optional] 
**Host** | Pointer to **string** | Host where the trap needs to be sent. | [optional] 
**Port** | Pointer to **string** | Port on the Host where the trap needs to be sent, default is 162. | [optional] [default to "162"]
**PrivacyConfig** | Pointer to [**MonitoringPrivacyConfig**](monitoringPrivacyConfig.md) |  | [optional] 
**Version** | Pointer to **string** |  | [optional] [default to "v2c"]

## Methods

### NewMonitoringSNMPTrapServer

`func NewMonitoringSNMPTrapServer() *MonitoringSNMPTrapServer`

NewMonitoringSNMPTrapServer instantiates a new MonitoringSNMPTrapServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringSNMPTrapServerWithDefaults

`func NewMonitoringSNMPTrapServerWithDefaults() *MonitoringSNMPTrapServer`

NewMonitoringSNMPTrapServerWithDefaults instantiates a new MonitoringSNMPTrapServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthConfig

`func (o *MonitoringSNMPTrapServer) GetAuthConfig() MonitoringAuthConfig`

GetAuthConfig returns the AuthConfig field if non-nil, zero value otherwise.

### GetAuthConfigOk

`func (o *MonitoringSNMPTrapServer) GetAuthConfigOk() (*MonitoringAuthConfig, bool)`

GetAuthConfigOk returns a tuple with the AuthConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthConfig

`func (o *MonitoringSNMPTrapServer) SetAuthConfig(v MonitoringAuthConfig)`

SetAuthConfig sets AuthConfig field to given value.

### HasAuthConfig

`func (o *MonitoringSNMPTrapServer) HasAuthConfig() bool`

HasAuthConfig returns a boolean if a field has been set.

### GetCommunityOrUser

`func (o *MonitoringSNMPTrapServer) GetCommunityOrUser() string`

GetCommunityOrUser returns the CommunityOrUser field if non-nil, zero value otherwise.

### GetCommunityOrUserOk

`func (o *MonitoringSNMPTrapServer) GetCommunityOrUserOk() (*string, bool)`

GetCommunityOrUserOk returns a tuple with the CommunityOrUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityOrUser

`func (o *MonitoringSNMPTrapServer) SetCommunityOrUser(v string)`

SetCommunityOrUser sets CommunityOrUser field to given value.

### HasCommunityOrUser

`func (o *MonitoringSNMPTrapServer) HasCommunityOrUser() bool`

HasCommunityOrUser returns a boolean if a field has been set.

### GetHost

`func (o *MonitoringSNMPTrapServer) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *MonitoringSNMPTrapServer) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *MonitoringSNMPTrapServer) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *MonitoringSNMPTrapServer) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *MonitoringSNMPTrapServer) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *MonitoringSNMPTrapServer) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *MonitoringSNMPTrapServer) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *MonitoringSNMPTrapServer) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPrivacyConfig

`func (o *MonitoringSNMPTrapServer) GetPrivacyConfig() MonitoringPrivacyConfig`

GetPrivacyConfig returns the PrivacyConfig field if non-nil, zero value otherwise.

### GetPrivacyConfigOk

`func (o *MonitoringSNMPTrapServer) GetPrivacyConfigOk() (*MonitoringPrivacyConfig, bool)`

GetPrivacyConfigOk returns a tuple with the PrivacyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyConfig

`func (o *MonitoringSNMPTrapServer) SetPrivacyConfig(v MonitoringPrivacyConfig)`

SetPrivacyConfig sets PrivacyConfig field to given value.

### HasPrivacyConfig

`func (o *MonitoringSNMPTrapServer) HasPrivacyConfig() bool`

HasPrivacyConfig returns a boolean if a field has been set.

### GetVersion

`func (o *MonitoringSNMPTrapServer) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MonitoringSNMPTrapServer) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MonitoringSNMPTrapServer) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MonitoringSNMPTrapServer) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


