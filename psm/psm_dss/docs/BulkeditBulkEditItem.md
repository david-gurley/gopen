# BulkeditBulkEditItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | Pointer to **string** |  | [optional] [default to "create"]
**Object** | Pointer to [**ApiAny**](apiAny.md) |  | [optional] 
**Uri** | Pointer to **string** | URI field: For Create, update and delete operations, the backend uses the Kind, tenant and name fields in the Object to Infer the URI The URI field is expected to be filled in only for label(required) and delete(optional) operations For delete operation, if the URI is populated, the Object field can be left empty. If the URI is empty, the object field is expected to be populated with the Object to be deleted For label operations, the URI must be populated and the Object must be of api.Label type. | [optional] 

## Methods

### NewBulkeditBulkEditItem

`func NewBulkeditBulkEditItem() *BulkeditBulkEditItem`

NewBulkeditBulkEditItem instantiates a new BulkeditBulkEditItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkeditBulkEditItemWithDefaults

`func NewBulkeditBulkEditItemWithDefaults() *BulkeditBulkEditItem`

NewBulkeditBulkEditItemWithDefaults instantiates a new BulkeditBulkEditItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *BulkeditBulkEditItem) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *BulkeditBulkEditItem) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *BulkeditBulkEditItem) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *BulkeditBulkEditItem) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetObject

`func (o *BulkeditBulkEditItem) GetObject() ApiAny`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *BulkeditBulkEditItem) GetObjectOk() (*ApiAny, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *BulkeditBulkEditItem) SetObject(v ApiAny)`

SetObject sets Object field to given value.

### HasObject

`func (o *BulkeditBulkEditItem) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetUri

`func (o *BulkeditBulkEditItem) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *BulkeditBulkEditItem) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *BulkeditBulkEditItem) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *BulkeditBulkEditItem) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


