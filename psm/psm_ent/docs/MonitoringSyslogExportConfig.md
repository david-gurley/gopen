# MonitoringSyslogExportConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FacilityOverride** | Pointer to **string** | override default facility with this in exported logs. | [optional] [default to "user"]
**Prefix** | Pointer to **string** | add prefix in exported logs. | [optional] 

## Methods

### NewMonitoringSyslogExportConfig

`func NewMonitoringSyslogExportConfig() *MonitoringSyslogExportConfig`

NewMonitoringSyslogExportConfig instantiates a new MonitoringSyslogExportConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringSyslogExportConfigWithDefaults

`func NewMonitoringSyslogExportConfigWithDefaults() *MonitoringSyslogExportConfig`

NewMonitoringSyslogExportConfigWithDefaults instantiates a new MonitoringSyslogExportConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFacilityOverride

`func (o *MonitoringSyslogExportConfig) GetFacilityOverride() string`

GetFacilityOverride returns the FacilityOverride field if non-nil, zero value otherwise.

### GetFacilityOverrideOk

`func (o *MonitoringSyslogExportConfig) GetFacilityOverrideOk() (*string, bool)`

GetFacilityOverrideOk returns a tuple with the FacilityOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityOverride

`func (o *MonitoringSyslogExportConfig) SetFacilityOverride(v string)`

SetFacilityOverride sets FacilityOverride field to given value.

### HasFacilityOverride

`func (o *MonitoringSyslogExportConfig) HasFacilityOverride() bool`

HasFacilityOverride returns a boolean if a field has been set.

### GetPrefix

`func (o *MonitoringSyslogExportConfig) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *MonitoringSyslogExportConfig) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *MonitoringSyslogExportConfig) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *MonitoringSyslogExportConfig) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


