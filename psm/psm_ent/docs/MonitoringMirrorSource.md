# MonitoringMirrorSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | Pointer to **string** | rx is towards the Source and tx is from Source. | [optional] [default to "both"]
**TargetSelectors** | Pointer to [**[]LabelsSelector**](LabelsSelector.md) |  | [optional] 
**TargetType** | Pointer to **string** |  | [optional] [default to "none"]

## Methods

### NewMonitoringMirrorSource

`func NewMonitoringMirrorSource() *MonitoringMirrorSource`

NewMonitoringMirrorSource instantiates a new MonitoringMirrorSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringMirrorSourceWithDefaults

`func NewMonitoringMirrorSourceWithDefaults() *MonitoringMirrorSource`

NewMonitoringMirrorSourceWithDefaults instantiates a new MonitoringMirrorSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *MonitoringMirrorSource) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *MonitoringMirrorSource) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *MonitoringMirrorSource) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *MonitoringMirrorSource) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetTargetSelectors

`func (o *MonitoringMirrorSource) GetTargetSelectors() []LabelsSelector`

GetTargetSelectors returns the TargetSelectors field if non-nil, zero value otherwise.

### GetTargetSelectorsOk

`func (o *MonitoringMirrorSource) GetTargetSelectorsOk() (*[]LabelsSelector, bool)`

GetTargetSelectorsOk returns a tuple with the TargetSelectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSelectors

`func (o *MonitoringMirrorSource) SetTargetSelectors(v []LabelsSelector)`

SetTargetSelectors sets TargetSelectors field to given value.

### HasTargetSelectors

`func (o *MonitoringMirrorSource) HasTargetSelectors() bool`

HasTargetSelectors returns a boolean if a field has been set.

### GetTargetType

`func (o *MonitoringMirrorSource) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *MonitoringMirrorSource) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *MonitoringMirrorSource) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *MonitoringMirrorSource) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


