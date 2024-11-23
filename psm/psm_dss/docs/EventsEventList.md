# EventsEventList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Items** | Pointer to [**[]EventsEvent**](EventsEvent.md) |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**ResourceVersion** | Pointer to **string** |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewEventsEventList

`func NewEventsEventList() *EventsEventList`

NewEventsEventList instantiates a new EventsEventList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsEventListWithDefaults

`func NewEventsEventListWithDefaults() *EventsEventList`

NewEventsEventListWithDefaults instantiates a new EventsEventList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *EventsEventList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *EventsEventList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *EventsEventList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *EventsEventList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *EventsEventList) GetItems() []EventsEvent`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *EventsEventList) GetItemsOk() (*[]EventsEvent, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *EventsEventList) SetItems(v []EventsEvent)`

SetItems sets Items field to given value.

### HasItems

`func (o *EventsEventList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetKind

`func (o *EventsEventList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *EventsEventList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *EventsEventList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *EventsEventList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetResourceVersion

`func (o *EventsEventList) GetResourceVersion() string`

GetResourceVersion returns the ResourceVersion field if non-nil, zero value otherwise.

### GetResourceVersionOk

`func (o *EventsEventList) GetResourceVersionOk() (*string, bool)`

GetResourceVersionOk returns a tuple with the ResourceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceVersion

`func (o *EventsEventList) SetResourceVersion(v string)`

SetResourceVersion sets ResourceVersion field to given value.

### HasResourceVersion

`func (o *EventsEventList) HasResourceVersion() bool`

HasResourceVersion returns a boolean if a field has been set.

### GetTotalCount

`func (o *EventsEventList) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *EventsEventList) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *EventsEventList) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *EventsEventList) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


