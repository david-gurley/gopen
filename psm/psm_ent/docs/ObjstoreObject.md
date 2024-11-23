# ObjstoreObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**ApiObjectMeta**](apiObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**ObjstoreObjectSpec**](objstoreObjectSpec.md) |  | [optional] 
**Status** | Pointer to [**ObjstoreObjectStatus**](objstoreObjectStatus.md) |  | [optional] 

## Methods

### NewObjstoreObject

`func NewObjstoreObject() *ObjstoreObject`

NewObjstoreObject instantiates a new ObjstoreObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjstoreObjectWithDefaults

`func NewObjstoreObjectWithDefaults() *ObjstoreObject`

NewObjstoreObjectWithDefaults instantiates a new ObjstoreObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ObjstoreObject) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ObjstoreObject) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ObjstoreObject) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ObjstoreObject) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *ObjstoreObject) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ObjstoreObject) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ObjstoreObject) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ObjstoreObject) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMeta

`func (o *ObjstoreObject) GetMeta() ApiObjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ObjstoreObject) GetMetaOk() (*ApiObjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ObjstoreObject) SetMeta(v ApiObjectMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ObjstoreObject) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSpec

`func (o *ObjstoreObject) GetSpec() ObjstoreObjectSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ObjstoreObject) GetSpecOk() (*ObjstoreObjectSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ObjstoreObject) SetSpec(v ObjstoreObjectSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *ObjstoreObject) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *ObjstoreObject) GetStatus() ObjstoreObjectStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ObjstoreObject) GetStatusOk() (*ObjstoreObjectStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ObjstoreObject) SetStatus(v ObjstoreObjectStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ObjstoreObject) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


