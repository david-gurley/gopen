# SysruntimeFlowKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationVrfId** | Pointer to **string** |  | [optional] 
**Ipv4** | Pointer to [**SysruntimeFlowKeyV4**](sysruntimeFlowKeyV4.md) |  | [optional] 
**L2** | Pointer to [**SysruntimeFlowKeyL2**](sysruntimeFlowKeyL2.md) |  | [optional] 
**SourceVrfId** | Pointer to **string** |  | [optional] 

## Methods

### NewSysruntimeFlowKey

`func NewSysruntimeFlowKey() *SysruntimeFlowKey`

NewSysruntimeFlowKey instantiates a new SysruntimeFlowKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSysruntimeFlowKeyWithDefaults

`func NewSysruntimeFlowKeyWithDefaults() *SysruntimeFlowKey`

NewSysruntimeFlowKeyWithDefaults instantiates a new SysruntimeFlowKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationVrfId

`func (o *SysruntimeFlowKey) GetDestinationVrfId() string`

GetDestinationVrfId returns the DestinationVrfId field if non-nil, zero value otherwise.

### GetDestinationVrfIdOk

`func (o *SysruntimeFlowKey) GetDestinationVrfIdOk() (*string, bool)`

GetDestinationVrfIdOk returns a tuple with the DestinationVrfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationVrfId

`func (o *SysruntimeFlowKey) SetDestinationVrfId(v string)`

SetDestinationVrfId sets DestinationVrfId field to given value.

### HasDestinationVrfId

`func (o *SysruntimeFlowKey) HasDestinationVrfId() bool`

HasDestinationVrfId returns a boolean if a field has been set.

### GetIpv4

`func (o *SysruntimeFlowKey) GetIpv4() SysruntimeFlowKeyV4`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *SysruntimeFlowKey) GetIpv4Ok() (*SysruntimeFlowKeyV4, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *SysruntimeFlowKey) SetIpv4(v SysruntimeFlowKeyV4)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *SysruntimeFlowKey) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetL2

`func (o *SysruntimeFlowKey) GetL2() SysruntimeFlowKeyL2`

GetL2 returns the L2 field if non-nil, zero value otherwise.

### GetL2Ok

`func (o *SysruntimeFlowKey) GetL2Ok() (*SysruntimeFlowKeyL2, bool)`

GetL2Ok returns a tuple with the L2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL2

`func (o *SysruntimeFlowKey) SetL2(v SysruntimeFlowKeyL2)`

SetL2 sets L2 field to given value.

### HasL2

`func (o *SysruntimeFlowKey) HasL2() bool`

HasL2 returns a boolean if a field has been set.

### GetSourceVrfId

`func (o *SysruntimeFlowKey) GetSourceVrfId() string`

GetSourceVrfId returns the SourceVrfId field if non-nil, zero value otherwise.

### GetSourceVrfIdOk

`func (o *SysruntimeFlowKey) GetSourceVrfIdOk() (*string, bool)`

GetSourceVrfIdOk returns a tuple with the SourceVrfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVrfId

`func (o *SysruntimeFlowKey) SetSourceVrfId(v string)`

SetSourceVrfId sets SourceVrfId field to given value.

### HasSourceVrfId

`func (o *SysruntimeFlowKey) HasSourceVrfId() bool`

HasSourceVrfId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


