# EventsEventAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | Category represents the category of an event. e.g. Cluster/Network/Datapath. | [optional] [default to "cluster"]
**Count** | Pointer to **int64** | Number of occurrence of this event in the active interval. | [optional] 
**Message** | Pointer to **string** | Message represents the human readable description of an event. | [optional] 
**ObjectRef** | Pointer to [**ApiObjectRef**](apiObjectRef.md) |  | [optional] 
**Severity** | Pointer to **string** | Severity represents the criticality level of an event. | [optional] [default to "info"]
**Source** | Pointer to [**EventsEventSource**](eventsEventSource.md) |  | [optional] 
**Type** | Pointer to **string** | Type represents the type of an event. e.g. NICAdmittedEvent, NodeJoined. | [optional] 

## Methods

### NewEventsEventAttributes

`func NewEventsEventAttributes() *EventsEventAttributes`

NewEventsEventAttributes instantiates a new EventsEventAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsEventAttributesWithDefaults

`func NewEventsEventAttributesWithDefaults() *EventsEventAttributes`

NewEventsEventAttributesWithDefaults instantiates a new EventsEventAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *EventsEventAttributes) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *EventsEventAttributes) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *EventsEventAttributes) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *EventsEventAttributes) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCount

`func (o *EventsEventAttributes) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *EventsEventAttributes) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *EventsEventAttributes) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *EventsEventAttributes) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetMessage

`func (o *EventsEventAttributes) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *EventsEventAttributes) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *EventsEventAttributes) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *EventsEventAttributes) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetObjectRef

`func (o *EventsEventAttributes) GetObjectRef() ApiObjectRef`

GetObjectRef returns the ObjectRef field if non-nil, zero value otherwise.

### GetObjectRefOk

`func (o *EventsEventAttributes) GetObjectRefOk() (*ApiObjectRef, bool)`

GetObjectRefOk returns a tuple with the ObjectRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectRef

`func (o *EventsEventAttributes) SetObjectRef(v ApiObjectRef)`

SetObjectRef sets ObjectRef field to given value.

### HasObjectRef

`func (o *EventsEventAttributes) HasObjectRef() bool`

HasObjectRef returns a boolean if a field has been set.

### GetSeverity

`func (o *EventsEventAttributes) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *EventsEventAttributes) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *EventsEventAttributes) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *EventsEventAttributes) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSource

`func (o *EventsEventAttributes) GetSource() EventsEventSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *EventsEventAttributes) GetSourceOk() (*EventsEventSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *EventsEventAttributes) SetSource(v EventsEventSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *EventsEventAttributes) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetType

`func (o *EventsEventAttributes) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventsEventAttributes) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventsEventAttributes) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EventsEventAttributes) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


