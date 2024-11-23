# MonitoringMirrorExportConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | Pointer to **string** | IP address of the collector/entity to which the data is to be exported. Should be a valid IPv4 address. | [optional] 
**Gateway** | Pointer to **string** | IP address of the gateway to reach the collector. | [optional] 
**VirtualRouter** | Pointer to **string** | Destination VirtualRouter/VRF for mirror session. If not specified, the value is populated as \&quot;default\&quot;. For underlay Vrf, the VrfName should be \&quot;underlay-vpc\&quot;. | [optional] 

## Methods

### NewMonitoringMirrorExportConfig

`func NewMonitoringMirrorExportConfig() *MonitoringMirrorExportConfig`

NewMonitoringMirrorExportConfig instantiates a new MonitoringMirrorExportConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringMirrorExportConfigWithDefaults

`func NewMonitoringMirrorExportConfigWithDefaults() *MonitoringMirrorExportConfig`

NewMonitoringMirrorExportConfigWithDefaults instantiates a new MonitoringMirrorExportConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *MonitoringMirrorExportConfig) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *MonitoringMirrorExportConfig) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *MonitoringMirrorExportConfig) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *MonitoringMirrorExportConfig) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetGateway

`func (o *MonitoringMirrorExportConfig) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *MonitoringMirrorExportConfig) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *MonitoringMirrorExportConfig) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *MonitoringMirrorExportConfig) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetVirtualRouter

`func (o *MonitoringMirrorExportConfig) GetVirtualRouter() string`

GetVirtualRouter returns the VirtualRouter field if non-nil, zero value otherwise.

### GetVirtualRouterOk

`func (o *MonitoringMirrorExportConfig) GetVirtualRouterOk() (*string, bool)`

GetVirtualRouterOk returns a tuple with the VirtualRouter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualRouter

`func (o *MonitoringMirrorExportConfig) SetVirtualRouter(v string)`

SetVirtualRouter sets VirtualRouter field to given value.

### HasVirtualRouter

`func (o *MonitoringMirrorExportConfig) HasVirtualRouter() bool`

HasVirtualRouter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


