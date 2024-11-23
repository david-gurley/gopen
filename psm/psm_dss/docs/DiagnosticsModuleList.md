# DiagnosticsModuleList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Items** | Pointer to [**[]DiagnosticsModule**](DiagnosticsModule.md) | List of Module objects. | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**ListMeta** | Pointer to [**ApiListMeta**](apiListMeta.md) |  | [optional] 

## Methods

### NewDiagnosticsModuleList

`func NewDiagnosticsModuleList() *DiagnosticsModuleList`

NewDiagnosticsModuleList instantiates a new DiagnosticsModuleList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiagnosticsModuleListWithDefaults

`func NewDiagnosticsModuleListWithDefaults() *DiagnosticsModuleList`

NewDiagnosticsModuleListWithDefaults instantiates a new DiagnosticsModuleList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *DiagnosticsModuleList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *DiagnosticsModuleList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *DiagnosticsModuleList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *DiagnosticsModuleList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *DiagnosticsModuleList) GetItems() []DiagnosticsModule`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *DiagnosticsModuleList) GetItemsOk() (*[]DiagnosticsModule, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *DiagnosticsModuleList) SetItems(v []DiagnosticsModule)`

SetItems sets Items field to given value.

### HasItems

`func (o *DiagnosticsModuleList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetKind

`func (o *DiagnosticsModuleList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *DiagnosticsModuleList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *DiagnosticsModuleList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *DiagnosticsModuleList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetListMeta

`func (o *DiagnosticsModuleList) GetListMeta() ApiListMeta`

GetListMeta returns the ListMeta field if non-nil, zero value otherwise.

### GetListMetaOk

`func (o *DiagnosticsModuleList) GetListMetaOk() (*ApiListMeta, bool)`

GetListMetaOk returns a tuple with the ListMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListMeta

`func (o *DiagnosticsModuleList) SetListMeta(v ApiListMeta)`

SetListMeta sets ListMeta field to given value.

### HasListMeta

`func (o *DiagnosticsModuleList) HasListMeta() bool`

HasListMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


