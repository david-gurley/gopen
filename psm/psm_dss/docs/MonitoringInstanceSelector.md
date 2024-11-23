# MonitoringInstanceSelector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** | Kind of the instances to be selected using names/label. | [optional] 
**Labels** | Pointer to [**LabelsSelector**](labelsSelector.md) |  | [optional] 
**Names** | Pointer to **[]string** | List of names/reporter IDs. | [optional] 

## Methods

### NewMonitoringInstanceSelector

`func NewMonitoringInstanceSelector() *MonitoringInstanceSelector`

NewMonitoringInstanceSelector instantiates a new MonitoringInstanceSelector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringInstanceSelectorWithDefaults

`func NewMonitoringInstanceSelectorWithDefaults() *MonitoringInstanceSelector`

NewMonitoringInstanceSelectorWithDefaults instantiates a new MonitoringInstanceSelector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *MonitoringInstanceSelector) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *MonitoringInstanceSelector) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *MonitoringInstanceSelector) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *MonitoringInstanceSelector) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLabels

`func (o *MonitoringInstanceSelector) GetLabels() LabelsSelector`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *MonitoringInstanceSelector) GetLabelsOk() (*LabelsSelector, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *MonitoringInstanceSelector) SetLabels(v LabelsSelector)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *MonitoringInstanceSelector) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetNames

`func (o *MonitoringInstanceSelector) GetNames() []string`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *MonitoringInstanceSelector) GetNamesOk() (*[]string, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *MonitoringInstanceSelector) SetNames(v []string)`

SetNames sets Names field to given value.

### HasNames

`func (o *MonitoringInstanceSelector) HasNames() bool`

HasNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


