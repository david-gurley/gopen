# ApiObjectRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** | Kind represents the type of the API object. | [optional] 
**Name** | Pointer to **string** | Name of the object, unique within a Namespace for scoped objects. | [optional] 
**Namespace** | Pointer to **string** | Namespace of the object, for scoped objects. | [optional] 
**Tenant** | Pointer to **string** | Tenant of the object. | [optional] 
**Uri** | Pointer to **string** | URI is a link to accessing the referenced object. | [optional] 

## Methods

### NewApiObjectRef

`func NewApiObjectRef() *ApiObjectRef`

NewApiObjectRef instantiates a new ApiObjectRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiObjectRefWithDefaults

`func NewApiObjectRefWithDefaults() *ApiObjectRef`

NewApiObjectRefWithDefaults instantiates a new ApiObjectRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ApiObjectRef) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ApiObjectRef) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ApiObjectRef) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ApiObjectRef) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *ApiObjectRef) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiObjectRef) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiObjectRef) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiObjectRef) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *ApiObjectRef) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ApiObjectRef) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ApiObjectRef) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ApiObjectRef) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetTenant

`func (o *ApiObjectRef) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *ApiObjectRef) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *ApiObjectRef) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *ApiObjectRef) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetUri

`func (o *ApiObjectRef) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ApiObjectRef) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ApiObjectRef) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ApiObjectRef) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


