# MonitoringSNMPExport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SnmpTrapServers** | Pointer to [**[]MonitoringSNMPTrapServer**](MonitoringSNMPTrapServer.md) | TODO:  format, config SNMP trap destination(s). | [optional] 

## Methods

### NewMonitoringSNMPExport

`func NewMonitoringSNMPExport() *MonitoringSNMPExport`

NewMonitoringSNMPExport instantiates a new MonitoringSNMPExport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringSNMPExportWithDefaults

`func NewMonitoringSNMPExportWithDefaults() *MonitoringSNMPExport`

NewMonitoringSNMPExportWithDefaults instantiates a new MonitoringSNMPExport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnmpTrapServers

`func (o *MonitoringSNMPExport) GetSnmpTrapServers() []MonitoringSNMPTrapServer`

GetSnmpTrapServers returns the SnmpTrapServers field if non-nil, zero value otherwise.

### GetSnmpTrapServersOk

`func (o *MonitoringSNMPExport) GetSnmpTrapServersOk() (*[]MonitoringSNMPTrapServer, bool)`

GetSnmpTrapServersOk returns a tuple with the SnmpTrapServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpTrapServers

`func (o *MonitoringSNMPExport) SetSnmpTrapServers(v []MonitoringSNMPTrapServer)`

SetSnmpTrapServers sets SnmpTrapServers field to given value.

### HasSnmpTrapServers

`func (o *MonitoringSNMPExport) HasSnmpTrapServers() bool`

HasSnmpTrapServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


