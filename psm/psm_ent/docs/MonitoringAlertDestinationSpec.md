# MonitoringAlertDestinationSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailExport** | Pointer to [**MonitoringEmailExport**](monitoringEmailExport.md) |  | [optional] 
**Selector** | Pointer to [**FieldsSelector**](fieldsSelector.md) |  | [optional] 
**SnmpExport** | Pointer to [**MonitoringSNMPExport**](monitoringSNMPExport.md) |  | [optional] 
**SyslogExport** | Pointer to [**MonitoringSyslogExport**](monitoringSyslogExport.md) |  | [optional] 

## Methods

### NewMonitoringAlertDestinationSpec

`func NewMonitoringAlertDestinationSpec() *MonitoringAlertDestinationSpec`

NewMonitoringAlertDestinationSpec instantiates a new MonitoringAlertDestinationSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringAlertDestinationSpecWithDefaults

`func NewMonitoringAlertDestinationSpecWithDefaults() *MonitoringAlertDestinationSpec`

NewMonitoringAlertDestinationSpecWithDefaults instantiates a new MonitoringAlertDestinationSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailExport

`func (o *MonitoringAlertDestinationSpec) GetEmailExport() MonitoringEmailExport`

GetEmailExport returns the EmailExport field if non-nil, zero value otherwise.

### GetEmailExportOk

`func (o *MonitoringAlertDestinationSpec) GetEmailExportOk() (*MonitoringEmailExport, bool)`

GetEmailExportOk returns a tuple with the EmailExport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailExport

`func (o *MonitoringAlertDestinationSpec) SetEmailExport(v MonitoringEmailExport)`

SetEmailExport sets EmailExport field to given value.

### HasEmailExport

`func (o *MonitoringAlertDestinationSpec) HasEmailExport() bool`

HasEmailExport returns a boolean if a field has been set.

### GetSelector

`func (o *MonitoringAlertDestinationSpec) GetSelector() FieldsSelector`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *MonitoringAlertDestinationSpec) GetSelectorOk() (*FieldsSelector, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *MonitoringAlertDestinationSpec) SetSelector(v FieldsSelector)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *MonitoringAlertDestinationSpec) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### GetSnmpExport

`func (o *MonitoringAlertDestinationSpec) GetSnmpExport() MonitoringSNMPExport`

GetSnmpExport returns the SnmpExport field if non-nil, zero value otherwise.

### GetSnmpExportOk

`func (o *MonitoringAlertDestinationSpec) GetSnmpExportOk() (*MonitoringSNMPExport, bool)`

GetSnmpExportOk returns a tuple with the SnmpExport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpExport

`func (o *MonitoringAlertDestinationSpec) SetSnmpExport(v MonitoringSNMPExport)`

SetSnmpExport sets SnmpExport field to given value.

### HasSnmpExport

`func (o *MonitoringAlertDestinationSpec) HasSnmpExport() bool`

HasSnmpExport returns a boolean if a field has been set.

### GetSyslogExport

`func (o *MonitoringAlertDestinationSpec) GetSyslogExport() MonitoringSyslogExport`

GetSyslogExport returns the SyslogExport field if non-nil, zero value otherwise.

### GetSyslogExportOk

`func (o *MonitoringAlertDestinationSpec) GetSyslogExportOk() (*MonitoringSyslogExport, bool)`

GetSyslogExportOk returns a tuple with the SyslogExport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogExport

`func (o *MonitoringAlertDestinationSpec) SetSyslogExport(v MonitoringSyslogExport)`

SetSyslogExport sets SyslogExport field to given value.

### HasSyslogExport

`func (o *MonitoringAlertDestinationSpec) HasSyslogExport() bool`

HasSyslogExport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


