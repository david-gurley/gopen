# SysruntimeConnectionFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | Pointer to [**SysruntimeWorkloadSelector**](sysruntimeWorkloadSelector.md) |  | [optional] 
**DestinationPort** | Pointer to **int64** | destination port to be matched. Value should be between 1 and 65535. | [optional] 
**Protocol** | Pointer to **string** | protocol to be matched. | [optional] [default to "none"]
**Source** | Pointer to [**SysruntimeWorkloadSelector**](sysruntimeWorkloadSelector.md) |  | [optional] 
**SourcePort** | Pointer to **int64** | source port to be matched. Value should be between 1 and 65535. | [optional] 

## Methods

### NewSysruntimeConnectionFilter

`func NewSysruntimeConnectionFilter() *SysruntimeConnectionFilter`

NewSysruntimeConnectionFilter instantiates a new SysruntimeConnectionFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSysruntimeConnectionFilterWithDefaults

`func NewSysruntimeConnectionFilterWithDefaults() *SysruntimeConnectionFilter`

NewSysruntimeConnectionFilterWithDefaults instantiates a new SysruntimeConnectionFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *SysruntimeConnectionFilter) GetDestination() SysruntimeWorkloadSelector`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *SysruntimeConnectionFilter) GetDestinationOk() (*SysruntimeWorkloadSelector, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *SysruntimeConnectionFilter) SetDestination(v SysruntimeWorkloadSelector)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *SysruntimeConnectionFilter) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetDestinationPort

`func (o *SysruntimeConnectionFilter) GetDestinationPort() int64`

GetDestinationPort returns the DestinationPort field if non-nil, zero value otherwise.

### GetDestinationPortOk

`func (o *SysruntimeConnectionFilter) GetDestinationPortOk() (*int64, bool)`

GetDestinationPortOk returns a tuple with the DestinationPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPort

`func (o *SysruntimeConnectionFilter) SetDestinationPort(v int64)`

SetDestinationPort sets DestinationPort field to given value.

### HasDestinationPort

`func (o *SysruntimeConnectionFilter) HasDestinationPort() bool`

HasDestinationPort returns a boolean if a field has been set.

### GetProtocol

`func (o *SysruntimeConnectionFilter) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *SysruntimeConnectionFilter) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *SysruntimeConnectionFilter) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *SysruntimeConnectionFilter) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSource

`func (o *SysruntimeConnectionFilter) GetSource() SysruntimeWorkloadSelector`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SysruntimeConnectionFilter) GetSourceOk() (*SysruntimeWorkloadSelector, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SysruntimeConnectionFilter) SetSource(v SysruntimeWorkloadSelector)`

SetSource sets Source field to given value.

### HasSource

`func (o *SysruntimeConnectionFilter) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSourcePort

`func (o *SysruntimeConnectionFilter) GetSourcePort() int64`

GetSourcePort returns the SourcePort field if non-nil, zero value otherwise.

### GetSourcePortOk

`func (o *SysruntimeConnectionFilter) GetSourcePortOk() (*int64, bool)`

GetSourcePortOk returns a tuple with the SourcePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePort

`func (o *SysruntimeConnectionFilter) SetSourcePort(v int64)`

SetSourcePort sets SourcePort field to given value.

### HasSourcePort

`func (o *SysruntimeConnectionFilter) HasSourcePort() bool`

HasSourcePort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


