# SysruntimeHWConnectionGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | Pointer to [**SysruntimeHWConnectionSpec**](sysruntimeHWConnectionSpec.md) |  | [optional] 
**Stats** | Pointer to [**SysruntimeHWConnectionStats**](sysruntimeHWConnectionStats.md) |  | [optional] 
**Status** | Pointer to [**SysruntimeHWConnectionStatus**](sysruntimeHWConnectionStatus.md) |  | [optional] 

## Methods

### NewSysruntimeHWConnectionGetResponse

`func NewSysruntimeHWConnectionGetResponse() *SysruntimeHWConnectionGetResponse`

NewSysruntimeHWConnectionGetResponse instantiates a new SysruntimeHWConnectionGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSysruntimeHWConnectionGetResponseWithDefaults

`func NewSysruntimeHWConnectionGetResponseWithDefaults() *SysruntimeHWConnectionGetResponse`

NewSysruntimeHWConnectionGetResponseWithDefaults instantiates a new SysruntimeHWConnectionGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *SysruntimeHWConnectionGetResponse) GetSpec() SysruntimeHWConnectionSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *SysruntimeHWConnectionGetResponse) GetSpecOk() (*SysruntimeHWConnectionSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *SysruntimeHWConnectionGetResponse) SetSpec(v SysruntimeHWConnectionSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *SysruntimeHWConnectionGetResponse) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStats

`func (o *SysruntimeHWConnectionGetResponse) GetStats() SysruntimeHWConnectionStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *SysruntimeHWConnectionGetResponse) GetStatsOk() (*SysruntimeHWConnectionStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *SysruntimeHWConnectionGetResponse) SetStats(v SysruntimeHWConnectionStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *SysruntimeHWConnectionGetResponse) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetStatus

`func (o *SysruntimeHWConnectionGetResponse) GetStatus() SysruntimeHWConnectionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SysruntimeHWConnectionGetResponse) GetStatusOk() (*SysruntimeHWConnectionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SysruntimeHWConnectionGetResponse) SetStatus(v SysruntimeHWConnectionStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SysruntimeHWConnectionGetResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


