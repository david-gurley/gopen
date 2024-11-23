# MonitoringExportConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | Pointer to **string** | IP address of the collector/entity to which the data is to be exported. Should be a valid IPv4 address. | [optional] 
**Gateway** | Pointer to **string** | Gateway of the dest IP address to which the data is to be exported. Should be a valid IPv4 address. | [optional] 
**Transport** | Pointer to **string** | protocol and Port number where an external collector is gathering the data example \&quot;UDP/2055\&quot;. Should be a valid layer 3 or layer 4 protocol and port/type (only support UDP currently). | [optional] 

## Methods

### NewMonitoringExportConfig

`func NewMonitoringExportConfig() *MonitoringExportConfig`

NewMonitoringExportConfig instantiates a new MonitoringExportConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringExportConfigWithDefaults

`func NewMonitoringExportConfigWithDefaults() *MonitoringExportConfig`

NewMonitoringExportConfigWithDefaults instantiates a new MonitoringExportConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *MonitoringExportConfig) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *MonitoringExportConfig) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *MonitoringExportConfig) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *MonitoringExportConfig) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetGateway

`func (o *MonitoringExportConfig) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *MonitoringExportConfig) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *MonitoringExportConfig) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *MonitoringExportConfig) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetTransport

`func (o *MonitoringExportConfig) GetTransport() string`

GetTransport returns the Transport field if non-nil, zero value otherwise.

### GetTransportOk

`func (o *MonitoringExportConfig) GetTransportOk() (*string, bool)`

GetTransportOk returns a tuple with the Transport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransport

`func (o *MonitoringExportConfig) SetTransport(v string)`

SetTransport sets Transport field to given value.

### HasTransport

`func (o *MonitoringExportConfig) HasTransport() bool`

HasTransport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


