# EventsEventSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Component** | Pointer to **string** | Component from which the event is generated. | [optional] 
**NodeName** | Pointer to **string** | Name of the venice or workload node which is generating the event. | [optional] 
**UnitId** | Pointer to **string** | Unit ID on the host which is generating the event. | [optional] 

## Methods

### NewEventsEventSource

`func NewEventsEventSource() *EventsEventSource`

NewEventsEventSource instantiates a new EventsEventSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsEventSourceWithDefaults

`func NewEventsEventSourceWithDefaults() *EventsEventSource`

NewEventsEventSourceWithDefaults instantiates a new EventsEventSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponent

`func (o *EventsEventSource) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *EventsEventSource) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *EventsEventSource) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *EventsEventSource) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetNodeName

`func (o *EventsEventSource) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *EventsEventSource) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *EventsEventSource) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.

### HasNodeName

`func (o *EventsEventSource) HasNodeName() bool`

HasNodeName returns a boolean if a field has been set.

### GetUnitId

`func (o *EventsEventSource) GetUnitId() string`

GetUnitId returns the UnitId field if non-nil, zero value otherwise.

### GetUnitIdOk

`func (o *EventsEventSource) GetUnitIdOk() (*string, bool)`

GetUnitIdOk returns a tuple with the UnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitId

`func (o *EventsEventSource) SetUnitId(v string)`

SetUnitId sets UnitId field to given value.

### HasUnitId

`func (o *EventsEventSource) HasUnitId() bool`

HasUnitId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


