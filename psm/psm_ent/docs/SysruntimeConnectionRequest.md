# SysruntimeConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**DscName** | Pointer to **string** | DSC on which the active connections are to be queried. | [optional] 
**Filters** | Pointer to [**[]SysruntimeConnectionFilter**](SysruntimeConnectionFilter.md) |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**List** | Pointer to [**ApiListWatchOptions**](apiListWatchOptions.md) |  | [optional] 

## Methods

### NewSysruntimeConnectionRequest

`func NewSysruntimeConnectionRequest() *SysruntimeConnectionRequest`

NewSysruntimeConnectionRequest instantiates a new SysruntimeConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSysruntimeConnectionRequestWithDefaults

`func NewSysruntimeConnectionRequestWithDefaults() *SysruntimeConnectionRequest`

NewSysruntimeConnectionRequestWithDefaults instantiates a new SysruntimeConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *SysruntimeConnectionRequest) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *SysruntimeConnectionRequest) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *SysruntimeConnectionRequest) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *SysruntimeConnectionRequest) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetDscName

`func (o *SysruntimeConnectionRequest) GetDscName() string`

GetDscName returns the DscName field if non-nil, zero value otherwise.

### GetDscNameOk

`func (o *SysruntimeConnectionRequest) GetDscNameOk() (*string, bool)`

GetDscNameOk returns a tuple with the DscName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDscName

`func (o *SysruntimeConnectionRequest) SetDscName(v string)`

SetDscName sets DscName field to given value.

### HasDscName

`func (o *SysruntimeConnectionRequest) HasDscName() bool`

HasDscName returns a boolean if a field has been set.

### GetFilters

`func (o *SysruntimeConnectionRequest) GetFilters() []SysruntimeConnectionFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *SysruntimeConnectionRequest) GetFiltersOk() (*[]SysruntimeConnectionFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *SysruntimeConnectionRequest) SetFilters(v []SysruntimeConnectionFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *SysruntimeConnectionRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetKind

`func (o *SysruntimeConnectionRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SysruntimeConnectionRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SysruntimeConnectionRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SysruntimeConnectionRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetList

`func (o *SysruntimeConnectionRequest) GetList() ApiListWatchOptions`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *SysruntimeConnectionRequest) GetListOk() (*ApiListWatchOptions, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *SysruntimeConnectionRequest) SetList(v ApiListWatchOptions)`

SetList sets List field to given value.

### HasList

`func (o *SysruntimeConnectionRequest) HasList() bool`

HasList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


