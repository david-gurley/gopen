# EventsEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] [default to "cluster"]
**Count** | Pointer to **int64** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**ApiObjectMeta**](apiObjectMeta.md) |  | [optional] 
**ObjectRef** | Pointer to [**ApiObjectRef**](apiObjectRef.md) |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] [default to "info"]
**Source** | Pointer to [**EventsEventSource**](eventsEventSource.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewEventsEvent

`func NewEventsEvent() *EventsEvent`

NewEventsEvent instantiates a new EventsEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsEventWithDefaults

`func NewEventsEventWithDefaults() *EventsEvent`

NewEventsEventWithDefaults instantiates a new EventsEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *EventsEvent) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *EventsEvent) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *EventsEvent) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *EventsEvent) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetCategory

`func (o *EventsEvent) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *EventsEvent) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *EventsEvent) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *EventsEvent) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCount

`func (o *EventsEvent) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *EventsEvent) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *EventsEvent) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *EventsEvent) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetKind

`func (o *EventsEvent) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *EventsEvent) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *EventsEvent) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *EventsEvent) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMessage

`func (o *EventsEvent) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *EventsEvent) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *EventsEvent) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *EventsEvent) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMeta

`func (o *EventsEvent) GetMeta() ApiObjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *EventsEvent) GetMetaOk() (*ApiObjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *EventsEvent) SetMeta(v ApiObjectMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *EventsEvent) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetObjectRef

`func (o *EventsEvent) GetObjectRef() ApiObjectRef`

GetObjectRef returns the ObjectRef field if non-nil, zero value otherwise.

### GetObjectRefOk

`func (o *EventsEvent) GetObjectRefOk() (*ApiObjectRef, bool)`

GetObjectRefOk returns a tuple with the ObjectRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectRef

`func (o *EventsEvent) SetObjectRef(v ApiObjectRef)`

SetObjectRef sets ObjectRef field to given value.

### HasObjectRef

`func (o *EventsEvent) HasObjectRef() bool`

HasObjectRef returns a boolean if a field has been set.

### GetSeverity

`func (o *EventsEvent) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *EventsEvent) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *EventsEvent) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *EventsEvent) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSource

`func (o *EventsEvent) GetSource() EventsEventSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *EventsEvent) GetSourceOk() (*EventsEventSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *EventsEvent) SetSource(v EventsEventSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *EventsEvent) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetType

`func (o *EventsEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventsEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventsEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EventsEvent) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


