# ApiWatchEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Control** | Pointer to [**ApiWatchControl**](apiWatchControl.md) |  | [optional] 
**Object** | Pointer to [**GoogleprotobufAny**](googleprotobufAny.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewApiWatchEvent

`func NewApiWatchEvent() *ApiWatchEvent`

NewApiWatchEvent instantiates a new ApiWatchEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiWatchEventWithDefaults

`func NewApiWatchEventWithDefaults() *ApiWatchEvent`

NewApiWatchEventWithDefaults instantiates a new ApiWatchEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetControl

`func (o *ApiWatchEvent) GetControl() ApiWatchControl`

GetControl returns the Control field if non-nil, zero value otherwise.

### GetControlOk

`func (o *ApiWatchEvent) GetControlOk() (*ApiWatchControl, bool)`

GetControlOk returns a tuple with the Control field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControl

`func (o *ApiWatchEvent) SetControl(v ApiWatchControl)`

SetControl sets Control field to given value.

### HasControl

`func (o *ApiWatchEvent) HasControl() bool`

HasControl returns a boolean if a field has been set.

### GetObject

`func (o *ApiWatchEvent) GetObject() GoogleprotobufAny`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ApiWatchEvent) GetObjectOk() (*GoogleprotobufAny, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ApiWatchEvent) SetObject(v GoogleprotobufAny)`

SetObject sets Object field to given value.

### HasObject

`func (o *ApiWatchEvent) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetType

`func (o *ApiWatchEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiWatchEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiWatchEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApiWatchEvent) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


