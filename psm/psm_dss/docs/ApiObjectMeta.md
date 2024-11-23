# ApiObjectMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationTime** | Pointer to **time.Time** | System generated and updated, not updatable by user. | [optional] 
**GenerationId** | Pointer to **string** | This is incremented anytime there is an update to the user intent, including Spec update and any update to ObjectMeta. System generated and updated, not updatable by user. | [optional] 
**Labels** | Pointer to **map[string]string** |  | [optional] 
**ModTime** | Pointer to **time.Time** | System generated and updated, not updatable by user. | [optional] 
**Name** | Pointer to **string** | Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | [optional] 
**Namespace** | Pointer to **string** | Namespace of the object, for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | [optional] 
**ResourceVersion** | Pointer to **string** | This is updated anytime there is any change to the object. System generated and updated, not updatable by user. | [optional] 
**SelfLink** | Pointer to **string** | When the object is served from the API-GW it is the URI path. Example: - \&quot;/v1/tenants/tenants/tenant2\&quot; System generated and updated, not updatable by user. | [optional] 
**Tenant** | Pointer to **string** | This can be automatically filled in many cases based on the tenant the user, who created the object, belongs to. Must be alpha-numerics. Length of string should be between 1 and 48. | [optional] 
**Uuid** | Pointer to **string** | This is generated on creation of the object. System generated, not updatable by user. | [optional] 

## Methods

### NewApiObjectMeta

`func NewApiObjectMeta() *ApiObjectMeta`

NewApiObjectMeta instantiates a new ApiObjectMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiObjectMetaWithDefaults

`func NewApiObjectMetaWithDefaults() *ApiObjectMeta`

NewApiObjectMetaWithDefaults instantiates a new ApiObjectMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationTime

`func (o *ApiObjectMeta) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *ApiObjectMeta) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *ApiObjectMeta) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *ApiObjectMeta) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetGenerationId

`func (o *ApiObjectMeta) GetGenerationId() string`

GetGenerationId returns the GenerationId field if non-nil, zero value otherwise.

### GetGenerationIdOk

`func (o *ApiObjectMeta) GetGenerationIdOk() (*string, bool)`

GetGenerationIdOk returns a tuple with the GenerationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerationId

`func (o *ApiObjectMeta) SetGenerationId(v string)`

SetGenerationId sets GenerationId field to given value.

### HasGenerationId

`func (o *ApiObjectMeta) HasGenerationId() bool`

HasGenerationId returns a boolean if a field has been set.

### GetLabels

`func (o *ApiObjectMeta) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ApiObjectMeta) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ApiObjectMeta) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ApiObjectMeta) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetModTime

`func (o *ApiObjectMeta) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *ApiObjectMeta) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *ApiObjectMeta) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *ApiObjectMeta) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetName

`func (o *ApiObjectMeta) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiObjectMeta) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiObjectMeta) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiObjectMeta) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *ApiObjectMeta) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ApiObjectMeta) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ApiObjectMeta) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ApiObjectMeta) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetResourceVersion

`func (o *ApiObjectMeta) GetResourceVersion() string`

GetResourceVersion returns the ResourceVersion field if non-nil, zero value otherwise.

### GetResourceVersionOk

`func (o *ApiObjectMeta) GetResourceVersionOk() (*string, bool)`

GetResourceVersionOk returns a tuple with the ResourceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceVersion

`func (o *ApiObjectMeta) SetResourceVersion(v string)`

SetResourceVersion sets ResourceVersion field to given value.

### HasResourceVersion

`func (o *ApiObjectMeta) HasResourceVersion() bool`

HasResourceVersion returns a boolean if a field has been set.

### GetSelfLink

`func (o *ApiObjectMeta) GetSelfLink() string`

GetSelfLink returns the SelfLink field if non-nil, zero value otherwise.

### GetSelfLinkOk

`func (o *ApiObjectMeta) GetSelfLinkOk() (*string, bool)`

GetSelfLinkOk returns a tuple with the SelfLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfLink

`func (o *ApiObjectMeta) SetSelfLink(v string)`

SetSelfLink sets SelfLink field to given value.

### HasSelfLink

`func (o *ApiObjectMeta) HasSelfLink() bool`

HasSelfLink returns a boolean if a field has been set.

### GetTenant

`func (o *ApiObjectMeta) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *ApiObjectMeta) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *ApiObjectMeta) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *ApiObjectMeta) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetUuid

`func (o *ApiObjectMeta) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ApiObjectMeta) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ApiObjectMeta) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ApiObjectMeta) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


