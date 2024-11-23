# MonitoringEventPolicySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**MonitoringSyslogExportConfig**](monitoringSyslogExportConfig.md) |  | [optional] 
**Format** | Pointer to **string** | event export format, SYSLOG_BSD default. | [optional] [default to "syslog-bsd"]
**Selector** | Pointer to [**FieldsSelector**](fieldsSelector.md) |  | [optional] 
**Targets** | Pointer to [**[]MonitoringExportConfig**](MonitoringExportConfig.md) | export target ip/port/protocol. | [optional] 

## Methods

### NewMonitoringEventPolicySpec

`func NewMonitoringEventPolicySpec() *MonitoringEventPolicySpec`

NewMonitoringEventPolicySpec instantiates a new MonitoringEventPolicySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringEventPolicySpecWithDefaults

`func NewMonitoringEventPolicySpecWithDefaults() *MonitoringEventPolicySpec`

NewMonitoringEventPolicySpecWithDefaults instantiates a new MonitoringEventPolicySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *MonitoringEventPolicySpec) GetConfig() MonitoringSyslogExportConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *MonitoringEventPolicySpec) GetConfigOk() (*MonitoringSyslogExportConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *MonitoringEventPolicySpec) SetConfig(v MonitoringSyslogExportConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *MonitoringEventPolicySpec) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetFormat

`func (o *MonitoringEventPolicySpec) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *MonitoringEventPolicySpec) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *MonitoringEventPolicySpec) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *MonitoringEventPolicySpec) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetSelector

`func (o *MonitoringEventPolicySpec) GetSelector() FieldsSelector`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *MonitoringEventPolicySpec) GetSelectorOk() (*FieldsSelector, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *MonitoringEventPolicySpec) SetSelector(v FieldsSelector)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *MonitoringEventPolicySpec) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### GetTargets

`func (o *MonitoringEventPolicySpec) GetTargets() []MonitoringExportConfig`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *MonitoringEventPolicySpec) GetTargetsOk() (*[]MonitoringExportConfig, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *MonitoringEventPolicySpec) SetTargets(v []MonitoringExportConfig)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *MonitoringEventPolicySpec) HasTargets() bool`

HasTargets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


