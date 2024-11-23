# ObjstoreBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**ApiObjectMeta**](apiObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**ObjstoreBucketSpec**](objstoreBucketSpec.md) |  | [optional] 
**Status** | Pointer to [**ObjstoreBucketStatus**](objstoreBucketStatus.md) |  | [optional] 

## Methods

### NewObjstoreBucket

`func NewObjstoreBucket() *ObjstoreBucket`

NewObjstoreBucket instantiates a new ObjstoreBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjstoreBucketWithDefaults

`func NewObjstoreBucketWithDefaults() *ObjstoreBucket`

NewObjstoreBucketWithDefaults instantiates a new ObjstoreBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ObjstoreBucket) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ObjstoreBucket) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ObjstoreBucket) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ObjstoreBucket) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *ObjstoreBucket) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ObjstoreBucket) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ObjstoreBucket) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ObjstoreBucket) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMeta

`func (o *ObjstoreBucket) GetMeta() ApiObjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ObjstoreBucket) GetMetaOk() (*ApiObjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ObjstoreBucket) SetMeta(v ApiObjectMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ObjstoreBucket) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSpec

`func (o *ObjstoreBucket) GetSpec() ObjstoreBucketSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ObjstoreBucket) GetSpecOk() (*ObjstoreBucketSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ObjstoreBucket) SetSpec(v ObjstoreBucketSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *ObjstoreBucket) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *ObjstoreBucket) GetStatus() ObjstoreBucketStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ObjstoreBucket) GetStatusOk() (*ObjstoreBucketStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ObjstoreBucket) SetStatus(v ObjstoreBucketStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ObjstoreBucket) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


