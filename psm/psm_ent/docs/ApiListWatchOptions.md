# ApiListWatchOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationTime** | Pointer to **time.Time** |  | [optional] 
**FieldChangeSelector** | Pointer to **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | [optional] 
**FieldSelector** | Pointer to **string** | FieldSelector to select on field values in list or watch results. | [optional] 
**From** | Pointer to **int32** | From represents the start index number (1 based - first object starts from index 1), of the results list. The results returned would be in the range [from ... (from + (max-results - 1))]. If From &#x3D; 0, the server will attempt to return all the results in the list without pagination. | [optional] 
**GenerationId** | Pointer to **string** |  | [optional] 
**LabelSelector** | Pointer to **string** | LabelSelector to select on labels in list or watch results. | [optional] 
**Labels** | Pointer to **map[string]string** |  | [optional] 
**MaxResults** | Pointer to **int32** | MaxResults is the maximum number of results to be returned as part of the response, per page If MaxResults is more than the maximum number of results per page supported by the server, the server will return an err If MaxResults is 0, the server will return all the results without pagination. | [optional] 
**MetaOnly** | Pointer to **bool** | If MetaOnly is set to true, the watch event notification that matches the watch criteria will not contain the full object. It will only contain the information about the object that changed, i.e. which object and what changed. MetaOnly is not set by default. | [optional] [default to false]
**ModTime** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** | Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | [optional] 
**Namespace** | Pointer to **string** | Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | [optional] 
**ResourceVersion** | Pointer to **string** |  | [optional] 
**SelfLink** | Pointer to **string** |  | [optional] 
**SortOrder** | Pointer to **string** | order to sort List results in. | [optional] [default to "none"]
**Tenant** | Pointer to **string** | Must be alpha-numerics. Length of string should be between 1 and 48. | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 

## Methods

### NewApiListWatchOptions

`func NewApiListWatchOptions() *ApiListWatchOptions`

NewApiListWatchOptions instantiates a new ApiListWatchOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiListWatchOptionsWithDefaults

`func NewApiListWatchOptionsWithDefaults() *ApiListWatchOptions`

NewApiListWatchOptionsWithDefaults instantiates a new ApiListWatchOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationTime

`func (o *ApiListWatchOptions) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *ApiListWatchOptions) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *ApiListWatchOptions) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *ApiListWatchOptions) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetFieldChangeSelector

`func (o *ApiListWatchOptions) GetFieldChangeSelector() []string`

GetFieldChangeSelector returns the FieldChangeSelector field if non-nil, zero value otherwise.

### GetFieldChangeSelectorOk

`func (o *ApiListWatchOptions) GetFieldChangeSelectorOk() (*[]string, bool)`

GetFieldChangeSelectorOk returns a tuple with the FieldChangeSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldChangeSelector

`func (o *ApiListWatchOptions) SetFieldChangeSelector(v []string)`

SetFieldChangeSelector sets FieldChangeSelector field to given value.

### HasFieldChangeSelector

`func (o *ApiListWatchOptions) HasFieldChangeSelector() bool`

HasFieldChangeSelector returns a boolean if a field has been set.

### GetFieldSelector

`func (o *ApiListWatchOptions) GetFieldSelector() string`

GetFieldSelector returns the FieldSelector field if non-nil, zero value otherwise.

### GetFieldSelectorOk

`func (o *ApiListWatchOptions) GetFieldSelectorOk() (*string, bool)`

GetFieldSelectorOk returns a tuple with the FieldSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldSelector

`func (o *ApiListWatchOptions) SetFieldSelector(v string)`

SetFieldSelector sets FieldSelector field to given value.

### HasFieldSelector

`func (o *ApiListWatchOptions) HasFieldSelector() bool`

HasFieldSelector returns a boolean if a field has been set.

### GetFrom

`func (o *ApiListWatchOptions) GetFrom() int32`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ApiListWatchOptions) GetFromOk() (*int32, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ApiListWatchOptions) SetFrom(v int32)`

SetFrom sets From field to given value.

### HasFrom

`func (o *ApiListWatchOptions) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetGenerationId

`func (o *ApiListWatchOptions) GetGenerationId() string`

GetGenerationId returns the GenerationId field if non-nil, zero value otherwise.

### GetGenerationIdOk

`func (o *ApiListWatchOptions) GetGenerationIdOk() (*string, bool)`

GetGenerationIdOk returns a tuple with the GenerationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerationId

`func (o *ApiListWatchOptions) SetGenerationId(v string)`

SetGenerationId sets GenerationId field to given value.

### HasGenerationId

`func (o *ApiListWatchOptions) HasGenerationId() bool`

HasGenerationId returns a boolean if a field has been set.

### GetLabelSelector

`func (o *ApiListWatchOptions) GetLabelSelector() string`

GetLabelSelector returns the LabelSelector field if non-nil, zero value otherwise.

### GetLabelSelectorOk

`func (o *ApiListWatchOptions) GetLabelSelectorOk() (*string, bool)`

GetLabelSelectorOk returns a tuple with the LabelSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelSelector

`func (o *ApiListWatchOptions) SetLabelSelector(v string)`

SetLabelSelector sets LabelSelector field to given value.

### HasLabelSelector

`func (o *ApiListWatchOptions) HasLabelSelector() bool`

HasLabelSelector returns a boolean if a field has been set.

### GetLabels

`func (o *ApiListWatchOptions) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ApiListWatchOptions) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ApiListWatchOptions) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ApiListWatchOptions) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetMaxResults

`func (o *ApiListWatchOptions) GetMaxResults() int32`

GetMaxResults returns the MaxResults field if non-nil, zero value otherwise.

### GetMaxResultsOk

`func (o *ApiListWatchOptions) GetMaxResultsOk() (*int32, bool)`

GetMaxResultsOk returns a tuple with the MaxResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResults

`func (o *ApiListWatchOptions) SetMaxResults(v int32)`

SetMaxResults sets MaxResults field to given value.

### HasMaxResults

`func (o *ApiListWatchOptions) HasMaxResults() bool`

HasMaxResults returns a boolean if a field has been set.

### GetMetaOnly

`func (o *ApiListWatchOptions) GetMetaOnly() bool`

GetMetaOnly returns the MetaOnly field if non-nil, zero value otherwise.

### GetMetaOnlyOk

`func (o *ApiListWatchOptions) GetMetaOnlyOk() (*bool, bool)`

GetMetaOnlyOk returns a tuple with the MetaOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaOnly

`func (o *ApiListWatchOptions) SetMetaOnly(v bool)`

SetMetaOnly sets MetaOnly field to given value.

### HasMetaOnly

`func (o *ApiListWatchOptions) HasMetaOnly() bool`

HasMetaOnly returns a boolean if a field has been set.

### GetModTime

`func (o *ApiListWatchOptions) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *ApiListWatchOptions) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *ApiListWatchOptions) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *ApiListWatchOptions) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetName

`func (o *ApiListWatchOptions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiListWatchOptions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiListWatchOptions) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiListWatchOptions) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *ApiListWatchOptions) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ApiListWatchOptions) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ApiListWatchOptions) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ApiListWatchOptions) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetResourceVersion

`func (o *ApiListWatchOptions) GetResourceVersion() string`

GetResourceVersion returns the ResourceVersion field if non-nil, zero value otherwise.

### GetResourceVersionOk

`func (o *ApiListWatchOptions) GetResourceVersionOk() (*string, bool)`

GetResourceVersionOk returns a tuple with the ResourceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceVersion

`func (o *ApiListWatchOptions) SetResourceVersion(v string)`

SetResourceVersion sets ResourceVersion field to given value.

### HasResourceVersion

`func (o *ApiListWatchOptions) HasResourceVersion() bool`

HasResourceVersion returns a boolean if a field has been set.

### GetSelfLink

`func (o *ApiListWatchOptions) GetSelfLink() string`

GetSelfLink returns the SelfLink field if non-nil, zero value otherwise.

### GetSelfLinkOk

`func (o *ApiListWatchOptions) GetSelfLinkOk() (*string, bool)`

GetSelfLinkOk returns a tuple with the SelfLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfLink

`func (o *ApiListWatchOptions) SetSelfLink(v string)`

SetSelfLink sets SelfLink field to given value.

### HasSelfLink

`func (o *ApiListWatchOptions) HasSelfLink() bool`

HasSelfLink returns a boolean if a field has been set.

### GetSortOrder

`func (o *ApiListWatchOptions) GetSortOrder() string`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *ApiListWatchOptions) GetSortOrderOk() (*string, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *ApiListWatchOptions) SetSortOrder(v string)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *ApiListWatchOptions) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetTenant

`func (o *ApiListWatchOptions) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *ApiListWatchOptions) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *ApiListWatchOptions) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *ApiListWatchOptions) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetUuid

`func (o *ApiListWatchOptions) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ApiListWatchOptions) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ApiListWatchOptions) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ApiListWatchOptions) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


