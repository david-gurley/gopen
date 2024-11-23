# ObjstoreObjectStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Digest** | Pointer to **string** | Digest is a hash digest of the object content. | [optional] 
**Size** | Pointer to **string** | Size is the total size of the object. | [optional] 

## Methods

### NewObjstoreObjectStatus

`func NewObjstoreObjectStatus() *ObjstoreObjectStatus`

NewObjstoreObjectStatus instantiates a new ObjstoreObjectStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjstoreObjectStatusWithDefaults

`func NewObjstoreObjectStatusWithDefaults() *ObjstoreObjectStatus`

NewObjstoreObjectStatusWithDefaults instantiates a new ObjstoreObjectStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDigest

`func (o *ObjstoreObjectStatus) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *ObjstoreObjectStatus) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *ObjstoreObjectStatus) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *ObjstoreObjectStatus) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetSize

`func (o *ObjstoreObjectStatus) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ObjstoreObjectStatus) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ObjstoreObjectStatus) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *ObjstoreObjectStatus) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


