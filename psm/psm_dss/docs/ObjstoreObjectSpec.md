# ObjstoreObjectSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentType** | Pointer to **string** | Content-Type for the stored object. Can either be specified when uploading. or the backend guesses one if possible. | [optional] 

## Methods

### NewObjstoreObjectSpec

`func NewObjstoreObjectSpec() *ObjstoreObjectSpec`

NewObjstoreObjectSpec instantiates a new ObjstoreObjectSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjstoreObjectSpecWithDefaults

`func NewObjstoreObjectSpecWithDefaults() *ObjstoreObjectSpec`

NewObjstoreObjectSpecWithDefaults instantiates a new ObjstoreObjectSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentType

`func (o *ObjstoreObjectSpec) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ObjstoreObjectSpec) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ObjstoreObjectSpec) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *ObjstoreObjectSpec) HasContentType() bool`

HasContentType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


