# MonitoringSyslogAuditor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**MonitoringSyslogExportConfig**](monitoringSyslogExportConfig.md) |  | [optional] 
**Enabled** | Pointer to **bool** | flag to enable sending audit events to syslog. | [optional] 
**Format** | Pointer to **string** | audit event export format, SYSLOG_BSD default. | [optional] [default to "syslog-bsd"]
**Targets** | Pointer to [**[]MonitoringExportConfig**](MonitoringExportConfig.md) | export target ip/port/protocol. | [optional] 

## Methods

### NewMonitoringSyslogAuditor

`func NewMonitoringSyslogAuditor() *MonitoringSyslogAuditor`

NewMonitoringSyslogAuditor instantiates a new MonitoringSyslogAuditor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringSyslogAuditorWithDefaults

`func NewMonitoringSyslogAuditorWithDefaults() *MonitoringSyslogAuditor`

NewMonitoringSyslogAuditorWithDefaults instantiates a new MonitoringSyslogAuditor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *MonitoringSyslogAuditor) GetConfig() MonitoringSyslogExportConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *MonitoringSyslogAuditor) GetConfigOk() (*MonitoringSyslogExportConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *MonitoringSyslogAuditor) SetConfig(v MonitoringSyslogExportConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *MonitoringSyslogAuditor) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetEnabled

`func (o *MonitoringSyslogAuditor) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MonitoringSyslogAuditor) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MonitoringSyslogAuditor) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MonitoringSyslogAuditor) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFormat

`func (o *MonitoringSyslogAuditor) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *MonitoringSyslogAuditor) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *MonitoringSyslogAuditor) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *MonitoringSyslogAuditor) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetTargets

`func (o *MonitoringSyslogAuditor) GetTargets() []MonitoringExportConfig`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *MonitoringSyslogAuditor) GetTargetsOk() (*[]MonitoringExportConfig, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *MonitoringSyslogAuditor) SetTargets(v []MonitoringExportConfig)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *MonitoringSyslogAuditor) HasTargets() bool`

HasTargets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


