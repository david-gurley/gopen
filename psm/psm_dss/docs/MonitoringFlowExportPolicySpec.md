# MonitoringFlowExportPolicySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disabled** | Pointer to **bool** | Enable/disable flow export. | [optional] 
**Exports** | Pointer to [**[]MonitoringExportConfig**](MonitoringExportConfig.md) | Export contains export parameters. | [optional] 
**Format** | Pointer to **string** |  | [optional] [default to "ipfix"]
**Interval** | Pointer to **string** | Interval defines how often to push the records to an external collector The value is specified as a string format, &#39;10s&#39;, &#39;20m&#39;. Should be a valid time duration between 1s and 24h0m0s. | [optional] [default to "10s"]
**MatchRules** | Pointer to [**[]MonitoringMatchRule**](MonitoringMatchRule.md) |  | [optional] 
**TemplateInterval** | Pointer to **string** | TemplateInterval defines how often to send ipfix templates to an external collector The value is specified as a string format, &#39;1m&#39;, &#39;10m&#39;. Should be a valid time duration between 1m0s and 30m0s. | [optional] [default to "5m"]

## Methods

### NewMonitoringFlowExportPolicySpec

`func NewMonitoringFlowExportPolicySpec() *MonitoringFlowExportPolicySpec`

NewMonitoringFlowExportPolicySpec instantiates a new MonitoringFlowExportPolicySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringFlowExportPolicySpecWithDefaults

`func NewMonitoringFlowExportPolicySpecWithDefaults() *MonitoringFlowExportPolicySpec`

NewMonitoringFlowExportPolicySpecWithDefaults instantiates a new MonitoringFlowExportPolicySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisabled

`func (o *MonitoringFlowExportPolicySpec) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *MonitoringFlowExportPolicySpec) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *MonitoringFlowExportPolicySpec) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *MonitoringFlowExportPolicySpec) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetExports

`func (o *MonitoringFlowExportPolicySpec) GetExports() []MonitoringExportConfig`

GetExports returns the Exports field if non-nil, zero value otherwise.

### GetExportsOk

`func (o *MonitoringFlowExportPolicySpec) GetExportsOk() (*[]MonitoringExportConfig, bool)`

GetExportsOk returns a tuple with the Exports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExports

`func (o *MonitoringFlowExportPolicySpec) SetExports(v []MonitoringExportConfig)`

SetExports sets Exports field to given value.

### HasExports

`func (o *MonitoringFlowExportPolicySpec) HasExports() bool`

HasExports returns a boolean if a field has been set.

### GetFormat

`func (o *MonitoringFlowExportPolicySpec) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *MonitoringFlowExportPolicySpec) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *MonitoringFlowExportPolicySpec) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *MonitoringFlowExportPolicySpec) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetInterval

`func (o *MonitoringFlowExportPolicySpec) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *MonitoringFlowExportPolicySpec) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *MonitoringFlowExportPolicySpec) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *MonitoringFlowExportPolicySpec) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetMatchRules

`func (o *MonitoringFlowExportPolicySpec) GetMatchRules() []MonitoringMatchRule`

GetMatchRules returns the MatchRules field if non-nil, zero value otherwise.

### GetMatchRulesOk

`func (o *MonitoringFlowExportPolicySpec) GetMatchRulesOk() (*[]MonitoringMatchRule, bool)`

GetMatchRulesOk returns a tuple with the MatchRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchRules

`func (o *MonitoringFlowExportPolicySpec) SetMatchRules(v []MonitoringMatchRule)`

SetMatchRules sets MatchRules field to given value.

### HasMatchRules

`func (o *MonitoringFlowExportPolicySpec) HasMatchRules() bool`

HasMatchRules returns a boolean if a field has been set.

### GetTemplateInterval

`func (o *MonitoringFlowExportPolicySpec) GetTemplateInterval() string`

GetTemplateInterval returns the TemplateInterval field if non-nil, zero value otherwise.

### GetTemplateIntervalOk

`func (o *MonitoringFlowExportPolicySpec) GetTemplateIntervalOk() (*string, bool)`

GetTemplateIntervalOk returns a tuple with the TemplateInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateInterval

`func (o *MonitoringFlowExportPolicySpec) SetTemplateInterval(v string)`

SetTemplateInterval sets TemplateInterval field to given value.

### HasTemplateInterval

`func (o *MonitoringFlowExportPolicySpec) HasTemplateInterval() bool`

HasTemplateInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


