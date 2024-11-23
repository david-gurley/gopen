# MonitoringInterfaceMirror

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | Pointer to **string** |  | [optional] [default to "both"]
**Selectors** | Pointer to [**[]LabelsSelector**](LabelsSelector.md) |  | [optional] 

## Methods

### NewMonitoringInterfaceMirror

`func NewMonitoringInterfaceMirror() *MonitoringInterfaceMirror`

NewMonitoringInterfaceMirror instantiates a new MonitoringInterfaceMirror object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringInterfaceMirrorWithDefaults

`func NewMonitoringInterfaceMirrorWithDefaults() *MonitoringInterfaceMirror`

NewMonitoringInterfaceMirrorWithDefaults instantiates a new MonitoringInterfaceMirror object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *MonitoringInterfaceMirror) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *MonitoringInterfaceMirror) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *MonitoringInterfaceMirror) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *MonitoringInterfaceMirror) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetSelectors

`func (o *MonitoringInterfaceMirror) GetSelectors() []LabelsSelector`

GetSelectors returns the Selectors field if non-nil, zero value otherwise.

### GetSelectorsOk

`func (o *MonitoringInterfaceMirror) GetSelectorsOk() (*[]LabelsSelector, bool)`

GetSelectorsOk returns a tuple with the Selectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectors

`func (o *MonitoringInterfaceMirror) SetSelectors(v []LabelsSelector)`

SetSelectors sets Selectors field to given value.

### HasSelectors

`func (o *MonitoringInterfaceMirror) HasSelectors() bool`

HasSelectors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


