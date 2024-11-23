# MonitoringMirrorCollector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExportConfig** | Pointer to [**MonitoringMirrorExportConfig**](monitoringMirrorExportConfig.md) |  | [optional] 
**StripVlanHdr** | Pointer to **bool** | remove vlan from mirror packet. | [optional] 
**Type** | Pointer to **string** | Type of Collector. | [optional] [default to "erspan_type_3"]

## Methods

### NewMonitoringMirrorCollector

`func NewMonitoringMirrorCollector() *MonitoringMirrorCollector`

NewMonitoringMirrorCollector instantiates a new MonitoringMirrorCollector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringMirrorCollectorWithDefaults

`func NewMonitoringMirrorCollectorWithDefaults() *MonitoringMirrorCollector`

NewMonitoringMirrorCollectorWithDefaults instantiates a new MonitoringMirrorCollector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExportConfig

`func (o *MonitoringMirrorCollector) GetExportConfig() MonitoringMirrorExportConfig`

GetExportConfig returns the ExportConfig field if non-nil, zero value otherwise.

### GetExportConfigOk

`func (o *MonitoringMirrorCollector) GetExportConfigOk() (*MonitoringMirrorExportConfig, bool)`

GetExportConfigOk returns a tuple with the ExportConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportConfig

`func (o *MonitoringMirrorCollector) SetExportConfig(v MonitoringMirrorExportConfig)`

SetExportConfig sets ExportConfig field to given value.

### HasExportConfig

`func (o *MonitoringMirrorCollector) HasExportConfig() bool`

HasExportConfig returns a boolean if a field has been set.

### GetStripVlanHdr

`func (o *MonitoringMirrorCollector) GetStripVlanHdr() bool`

GetStripVlanHdr returns the StripVlanHdr field if non-nil, zero value otherwise.

### GetStripVlanHdrOk

`func (o *MonitoringMirrorCollector) GetStripVlanHdrOk() (*bool, bool)`

GetStripVlanHdrOk returns a tuple with the StripVlanHdr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripVlanHdr

`func (o *MonitoringMirrorCollector) SetStripVlanHdr(v bool)`

SetStripVlanHdr sets StripVlanHdr field to given value.

### HasStripVlanHdr

`func (o *MonitoringMirrorCollector) HasStripVlanHdr() bool`

HasStripVlanHdr returns a boolean if a field has been set.

### GetType

`func (o *MonitoringMirrorCollector) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MonitoringMirrorCollector) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MonitoringMirrorCollector) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MonitoringMirrorCollector) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


