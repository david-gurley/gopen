# MonitoringSyslogExport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**MonitoringSyslogExportConfig**](monitoringSyslogExportConfig.md) |  | [optional] 
**Format** | Pointer to **string** | event export format, SYSLOG_BSD default. | [optional] [default to "syslog-bsd"]
**Targets** | Pointer to [**[]MonitoringExportConfig**](MonitoringExportConfig.md) | export target ip/port/protocol. | [optional] 

## Methods

### NewMonitoringSyslogExport

`func NewMonitoringSyslogExport() *MonitoringSyslogExport`

NewMonitoringSyslogExport instantiates a new MonitoringSyslogExport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringSyslogExportWithDefaults

`func NewMonitoringSyslogExportWithDefaults() *MonitoringSyslogExport`

NewMonitoringSyslogExportWithDefaults instantiates a new MonitoringSyslogExport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *MonitoringSyslogExport) GetConfig() MonitoringSyslogExportConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *MonitoringSyslogExport) GetConfigOk() (*MonitoringSyslogExportConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *MonitoringSyslogExport) SetConfig(v MonitoringSyslogExportConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *MonitoringSyslogExport) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetFormat

`func (o *MonitoringSyslogExport) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *MonitoringSyslogExport) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *MonitoringSyslogExport) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *MonitoringSyslogExport) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetTargets

`func (o *MonitoringSyslogExport) GetTargets() []MonitoringExportConfig`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *MonitoringSyslogExport) GetTargetsOk() (*[]MonitoringExportConfig, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *MonitoringSyslogExport) SetTargets(v []MonitoringExportConfig)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *MonitoringSyslogExport) HasTargets() bool`

HasTargets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


