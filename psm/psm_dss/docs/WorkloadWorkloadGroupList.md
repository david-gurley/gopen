# WorkloadWorkloadGroupList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Items** | Pointer to [**[]WorkloadWorkloadGroup**](WorkloadWorkloadGroup.md) | List of WorkloadGroup objects. | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**ListMeta** | Pointer to [**ApiListMeta**](apiListMeta.md) |  | [optional] 

## Methods

### NewWorkloadWorkloadGroupList

`func NewWorkloadWorkloadGroupList() *WorkloadWorkloadGroupList`

NewWorkloadWorkloadGroupList instantiates a new WorkloadWorkloadGroupList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadWorkloadGroupListWithDefaults

`func NewWorkloadWorkloadGroupListWithDefaults() *WorkloadWorkloadGroupList`

NewWorkloadWorkloadGroupListWithDefaults instantiates a new WorkloadWorkloadGroupList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *WorkloadWorkloadGroupList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *WorkloadWorkloadGroupList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *WorkloadWorkloadGroupList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *WorkloadWorkloadGroupList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *WorkloadWorkloadGroupList) GetItems() []WorkloadWorkloadGroup`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *WorkloadWorkloadGroupList) GetItemsOk() (*[]WorkloadWorkloadGroup, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *WorkloadWorkloadGroupList) SetItems(v []WorkloadWorkloadGroup)`

SetItems sets Items field to given value.

### HasItems

`func (o *WorkloadWorkloadGroupList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetKind

`func (o *WorkloadWorkloadGroupList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *WorkloadWorkloadGroupList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *WorkloadWorkloadGroupList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *WorkloadWorkloadGroupList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetListMeta

`func (o *WorkloadWorkloadGroupList) GetListMeta() ApiListMeta`

GetListMeta returns the ListMeta field if non-nil, zero value otherwise.

### GetListMetaOk

`func (o *WorkloadWorkloadGroupList) GetListMetaOk() (*ApiListMeta, bool)`

GetListMetaOk returns a tuple with the ListMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListMeta

`func (o *WorkloadWorkloadGroupList) SetListMeta(v ApiListMeta)`

SetListMeta sets ListMeta field to given value.

### HasListMeta

`func (o *WorkloadWorkloadGroupList) HasListMeta() bool`

HasListMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


