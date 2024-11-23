# ObjectURIs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RefType** | Pointer to **string** |  | [optional] [default to "named-reference"]
**Uri** | Pointer to [**[]ApiObjectRef**](ApiObjectRef.md) |  | [optional] 

## Methods

### NewObjectURIs

`func NewObjectURIs() *ObjectURIs`

NewObjectURIs instantiates a new ObjectURIs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectURIsWithDefaults

`func NewObjectURIsWithDefaults() *ObjectURIs`

NewObjectURIsWithDefaults instantiates a new ObjectURIs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefType

`func (o *ObjectURIs) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *ObjectURIs) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *ObjectURIs) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *ObjectURIs) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetUri

`func (o *ObjectURIs) GetUri() []ApiObjectRef`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ObjectURIs) GetUriOk() (*[]ApiObjectRef, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ObjectURIs) SetUri(v []ApiObjectRef)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ObjectURIs) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


