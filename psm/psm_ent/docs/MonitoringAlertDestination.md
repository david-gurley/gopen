# MonitoringAlertDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**ApiObjectMeta**](apiObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**MonitoringAlertDestinationSpec**](monitoringAlertDestinationSpec.md) |  | [optional] 
**Status** | Pointer to [**MonitoringAlertDestinationStatus**](monitoringAlertDestinationStatus.md) |  | [optional] 

## Methods

### NewMonitoringAlertDestination

`func NewMonitoringAlertDestination() *MonitoringAlertDestination`

NewMonitoringAlertDestination instantiates a new MonitoringAlertDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringAlertDestinationWithDefaults

`func NewMonitoringAlertDestinationWithDefaults() *MonitoringAlertDestination`

NewMonitoringAlertDestinationWithDefaults instantiates a new MonitoringAlertDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *MonitoringAlertDestination) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *MonitoringAlertDestination) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *MonitoringAlertDestination) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *MonitoringAlertDestination) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *MonitoringAlertDestination) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *MonitoringAlertDestination) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *MonitoringAlertDestination) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *MonitoringAlertDestination) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMeta

`func (o *MonitoringAlertDestination) GetMeta() ApiObjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MonitoringAlertDestination) GetMetaOk() (*ApiObjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MonitoringAlertDestination) SetMeta(v ApiObjectMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *MonitoringAlertDestination) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSpec

`func (o *MonitoringAlertDestination) GetSpec() MonitoringAlertDestinationSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *MonitoringAlertDestination) GetSpecOk() (*MonitoringAlertDestinationSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *MonitoringAlertDestination) SetSpec(v MonitoringAlertDestinationSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *MonitoringAlertDestination) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *MonitoringAlertDestination) GetStatus() MonitoringAlertDestinationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MonitoringAlertDestination) GetStatusOk() (*MonitoringAlertDestinationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MonitoringAlertDestination) SetStatus(v MonitoringAlertDestinationStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MonitoringAlertDestination) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


