# MonitoringMetricIdentifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldName** | Pointer to **string** | Field belonging to the kind e.g. ConnectionsPerSecond. This is the attribute that will be monitored and alerts will be created/resolved based on the thresholds. | [optional] 
**Group** | Pointer to **string** | Metric group - e.g. ftestats, flowstats, etc. | [optional] 
**Kind** | Pointer to **string** | Sub-category within the group e.g. MaxSessionThresholdDrops, FlowMissPackets. | [optional] 

## Methods

### NewMonitoringMetricIdentifier

`func NewMonitoringMetricIdentifier() *MonitoringMetricIdentifier`

NewMonitoringMetricIdentifier instantiates a new MonitoringMetricIdentifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringMetricIdentifierWithDefaults

`func NewMonitoringMetricIdentifierWithDefaults() *MonitoringMetricIdentifier`

NewMonitoringMetricIdentifierWithDefaults instantiates a new MonitoringMetricIdentifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldName

`func (o *MonitoringMetricIdentifier) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *MonitoringMetricIdentifier) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *MonitoringMetricIdentifier) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.

### HasFieldName

`func (o *MonitoringMetricIdentifier) HasFieldName() bool`

HasFieldName returns a boolean if a field has been set.

### GetGroup

`func (o *MonitoringMetricIdentifier) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *MonitoringMetricIdentifier) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *MonitoringMetricIdentifier) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *MonitoringMetricIdentifier) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetKind

`func (o *MonitoringMetricIdentifier) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *MonitoringMetricIdentifier) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *MonitoringMetricIdentifier) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *MonitoringMetricIdentifier) HasKind() bool`

HasKind returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


