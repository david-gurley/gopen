# MonitoringMatchSelector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAddresses** | Pointer to **[]string** | IP address list, example [\&quot;10.1.1.10\&quot;,\&quot;10.1.1.15\&quot;]. | [optional] 
**MacAddresses** | Pointer to **[]string** | List of MacAddresses - \&quot;aabb.ccdd.eeff\&quot;, \&quot;0001.0203.0405\&quot;. Should be a valid MAC address. | [optional] 

## Methods

### NewMonitoringMatchSelector

`func NewMonitoringMatchSelector() *MonitoringMatchSelector`

NewMonitoringMatchSelector instantiates a new MonitoringMatchSelector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringMatchSelectorWithDefaults

`func NewMonitoringMatchSelectorWithDefaults() *MonitoringMatchSelector`

NewMonitoringMatchSelectorWithDefaults instantiates a new MonitoringMatchSelector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpAddresses

`func (o *MonitoringMatchSelector) GetIpAddresses() []string`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *MonitoringMatchSelector) GetIpAddressesOk() (*[]string, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *MonitoringMatchSelector) SetIpAddresses(v []string)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *MonitoringMatchSelector) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

### GetMacAddresses

`func (o *MonitoringMatchSelector) GetMacAddresses() []string`

GetMacAddresses returns the MacAddresses field if non-nil, zero value otherwise.

### GetMacAddressesOk

`func (o *MonitoringMatchSelector) GetMacAddressesOk() (*[]string, bool)`

GetMacAddressesOk returns a tuple with the MacAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddresses

`func (o *MonitoringMatchSelector) SetMacAddresses(v []string)`

SetMacAddresses sets MacAddresses field to given value.

### HasMacAddresses

`func (o *MonitoringMatchSelector) HasMacAddresses() bool`

HasMacAddresses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


