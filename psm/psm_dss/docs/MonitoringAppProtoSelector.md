# MonitoringAppProtoSelector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applications** | Pointer to **[]string** | Apps - E.g. [\&quot;Redis\&quot;]. | [optional] 
**ProtoPorts** | Pointer to **[]string** | ports - Includes protocol name and port Eg [\&quot;tcp/1234\&quot;, \&quot;udp\&quot;]. Should be a valid layer 3 or layer 4 protocol and port range. any is also allowed. | [optional] 

## Methods

### NewMonitoringAppProtoSelector

`func NewMonitoringAppProtoSelector() *MonitoringAppProtoSelector`

NewMonitoringAppProtoSelector instantiates a new MonitoringAppProtoSelector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringAppProtoSelectorWithDefaults

`func NewMonitoringAppProtoSelectorWithDefaults() *MonitoringAppProtoSelector`

NewMonitoringAppProtoSelectorWithDefaults instantiates a new MonitoringAppProtoSelector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplications

`func (o *MonitoringAppProtoSelector) GetApplications() []string`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *MonitoringAppProtoSelector) GetApplicationsOk() (*[]string, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *MonitoringAppProtoSelector) SetApplications(v []string)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *MonitoringAppProtoSelector) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetProtoPorts

`func (o *MonitoringAppProtoSelector) GetProtoPorts() []string`

GetProtoPorts returns the ProtoPorts field if non-nil, zero value otherwise.

### GetProtoPortsOk

`func (o *MonitoringAppProtoSelector) GetProtoPortsOk() (*[]string, bool)`

GetProtoPortsOk returns a tuple with the ProtoPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtoPorts

`func (o *MonitoringAppProtoSelector) SetProtoPorts(v []string)`

SetProtoPorts sets ProtoPorts field to given value.

### HasProtoPorts

`func (o *MonitoringAppProtoSelector) HasProtoPorts() bool`

HasProtoPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


