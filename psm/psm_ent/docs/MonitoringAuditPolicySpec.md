# MonitoringAuditPolicySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SyslogAuditor** | Pointer to [**MonitoringSyslogAuditor**](monitoringSyslogAuditor.md) |  | [optional] 

## Methods

### NewMonitoringAuditPolicySpec

`func NewMonitoringAuditPolicySpec() *MonitoringAuditPolicySpec`

NewMonitoringAuditPolicySpec instantiates a new MonitoringAuditPolicySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringAuditPolicySpecWithDefaults

`func NewMonitoringAuditPolicySpecWithDefaults() *MonitoringAuditPolicySpec`

NewMonitoringAuditPolicySpecWithDefaults instantiates a new MonitoringAuditPolicySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSyslogAuditor

`func (o *MonitoringAuditPolicySpec) GetSyslogAuditor() MonitoringSyslogAuditor`

GetSyslogAuditor returns the SyslogAuditor field if non-nil, zero value otherwise.

### GetSyslogAuditorOk

`func (o *MonitoringAuditPolicySpec) GetSyslogAuditorOk() (*MonitoringSyslogAuditor, bool)`

GetSyslogAuditorOk returns a tuple with the SyslogAuditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogAuditor

`func (o *MonitoringAuditPolicySpec) SetSyslogAuditor(v MonitoringSyslogAuditor)`

SetSyslogAuditor sets SyslogAuditor field to given value.

### HasSyslogAuditor

`func (o *MonitoringAuditPolicySpec) HasSyslogAuditor() bool`

HasSyslogAuditor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


