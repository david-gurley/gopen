# MonitoringFwlogPolicySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**MonitoringSyslogExportConfig**](monitoringSyslogExportConfig.md) |  | [optional] 
**Filter** | Pointer to **[]string** | filter firewall logs, FWLOG_ALL default. | [optional] 
**Format** | Pointer to **string** | fwlog format, SYSLOG_BSD default. | [optional] [default to "syslog-bsd"]
**PsmTarget** | Pointer to [**MonitoringPSMExportTarget**](monitoringPSMExportTarget.md) |  | [optional] 
**Targets** | Pointer to [**[]MonitoringExportConfig**](MonitoringExportConfig.md) | Target contains ip/port/protocol. | [optional] 
**VrfName** | Pointer to **string** | VrfName specifies the name of the VRF that the current Firewall Log Policy belongs to. | [optional] 

## Methods

### NewMonitoringFwlogPolicySpec

`func NewMonitoringFwlogPolicySpec() *MonitoringFwlogPolicySpec`

NewMonitoringFwlogPolicySpec instantiates a new MonitoringFwlogPolicySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringFwlogPolicySpecWithDefaults

`func NewMonitoringFwlogPolicySpecWithDefaults() *MonitoringFwlogPolicySpec`

NewMonitoringFwlogPolicySpecWithDefaults instantiates a new MonitoringFwlogPolicySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *MonitoringFwlogPolicySpec) GetConfig() MonitoringSyslogExportConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *MonitoringFwlogPolicySpec) GetConfigOk() (*MonitoringSyslogExportConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *MonitoringFwlogPolicySpec) SetConfig(v MonitoringSyslogExportConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *MonitoringFwlogPolicySpec) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetFilter

`func (o *MonitoringFwlogPolicySpec) GetFilter() []string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *MonitoringFwlogPolicySpec) GetFilterOk() (*[]string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *MonitoringFwlogPolicySpec) SetFilter(v []string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *MonitoringFwlogPolicySpec) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetFormat

`func (o *MonitoringFwlogPolicySpec) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *MonitoringFwlogPolicySpec) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *MonitoringFwlogPolicySpec) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *MonitoringFwlogPolicySpec) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetPsmTarget

`func (o *MonitoringFwlogPolicySpec) GetPsmTarget() MonitoringPSMExportTarget`

GetPsmTarget returns the PsmTarget field if non-nil, zero value otherwise.

### GetPsmTargetOk

`func (o *MonitoringFwlogPolicySpec) GetPsmTargetOk() (*MonitoringPSMExportTarget, bool)`

GetPsmTargetOk returns a tuple with the PsmTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsmTarget

`func (o *MonitoringFwlogPolicySpec) SetPsmTarget(v MonitoringPSMExportTarget)`

SetPsmTarget sets PsmTarget field to given value.

### HasPsmTarget

`func (o *MonitoringFwlogPolicySpec) HasPsmTarget() bool`

HasPsmTarget returns a boolean if a field has been set.

### GetTargets

`func (o *MonitoringFwlogPolicySpec) GetTargets() []MonitoringExportConfig`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *MonitoringFwlogPolicySpec) GetTargetsOk() (*[]MonitoringExportConfig, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *MonitoringFwlogPolicySpec) SetTargets(v []MonitoringExportConfig)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *MonitoringFwlogPolicySpec) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetVrfName

`func (o *MonitoringFwlogPolicySpec) GetVrfName() string`

GetVrfName returns the VrfName field if non-nil, zero value otherwise.

### GetVrfNameOk

`func (o *MonitoringFwlogPolicySpec) GetVrfNameOk() (*string, bool)`

GetVrfNameOk returns a tuple with the VrfName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfName

`func (o *MonitoringFwlogPolicySpec) SetVrfName(v string)`

SetVrfName sets VrfName field to given value.

### HasVrfName

`func (o *MonitoringFwlogPolicySpec) HasVrfName() bool`

HasVrfName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


