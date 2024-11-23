# SysruntimeFlowInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlowAction** | Pointer to **string** |  | [optional] 
**FlowAge** | Pointer to **int64** |  | [optional] 
**NatDestination** | Pointer to **string** |  | [optional] 
**NatDestinationPort** | Pointer to **int64** |  | [optional] 
**NatSource** | Pointer to **string** |  | [optional] 
**NatSourcePort** | Pointer to **int64** |  | [optional] 
**NatType** | Pointer to **string** |  | [optional] 
**TcpState** | Pointer to **string** |  | [optional] 
**TimeToAge** | Pointer to **int64** |  | [optional] 

## Methods

### NewSysruntimeFlowInfo

`func NewSysruntimeFlowInfo() *SysruntimeFlowInfo`

NewSysruntimeFlowInfo instantiates a new SysruntimeFlowInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSysruntimeFlowInfoWithDefaults

`func NewSysruntimeFlowInfoWithDefaults() *SysruntimeFlowInfo`

NewSysruntimeFlowInfoWithDefaults instantiates a new SysruntimeFlowInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlowAction

`func (o *SysruntimeFlowInfo) GetFlowAction() string`

GetFlowAction returns the FlowAction field if non-nil, zero value otherwise.

### GetFlowActionOk

`func (o *SysruntimeFlowInfo) GetFlowActionOk() (*string, bool)`

GetFlowActionOk returns a tuple with the FlowAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowAction

`func (o *SysruntimeFlowInfo) SetFlowAction(v string)`

SetFlowAction sets FlowAction field to given value.

### HasFlowAction

`func (o *SysruntimeFlowInfo) HasFlowAction() bool`

HasFlowAction returns a boolean if a field has been set.

### GetFlowAge

`func (o *SysruntimeFlowInfo) GetFlowAge() int64`

GetFlowAge returns the FlowAge field if non-nil, zero value otherwise.

### GetFlowAgeOk

`func (o *SysruntimeFlowInfo) GetFlowAgeOk() (*int64, bool)`

GetFlowAgeOk returns a tuple with the FlowAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowAge

`func (o *SysruntimeFlowInfo) SetFlowAge(v int64)`

SetFlowAge sets FlowAge field to given value.

### HasFlowAge

`func (o *SysruntimeFlowInfo) HasFlowAge() bool`

HasFlowAge returns a boolean if a field has been set.

### GetNatDestination

`func (o *SysruntimeFlowInfo) GetNatDestination() string`

GetNatDestination returns the NatDestination field if non-nil, zero value otherwise.

### GetNatDestinationOk

`func (o *SysruntimeFlowInfo) GetNatDestinationOk() (*string, bool)`

GetNatDestinationOk returns a tuple with the NatDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatDestination

`func (o *SysruntimeFlowInfo) SetNatDestination(v string)`

SetNatDestination sets NatDestination field to given value.

### HasNatDestination

`func (o *SysruntimeFlowInfo) HasNatDestination() bool`

HasNatDestination returns a boolean if a field has been set.

### GetNatDestinationPort

`func (o *SysruntimeFlowInfo) GetNatDestinationPort() int64`

GetNatDestinationPort returns the NatDestinationPort field if non-nil, zero value otherwise.

### GetNatDestinationPortOk

`func (o *SysruntimeFlowInfo) GetNatDestinationPortOk() (*int64, bool)`

GetNatDestinationPortOk returns a tuple with the NatDestinationPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatDestinationPort

`func (o *SysruntimeFlowInfo) SetNatDestinationPort(v int64)`

SetNatDestinationPort sets NatDestinationPort field to given value.

### HasNatDestinationPort

`func (o *SysruntimeFlowInfo) HasNatDestinationPort() bool`

HasNatDestinationPort returns a boolean if a field has been set.

### GetNatSource

`func (o *SysruntimeFlowInfo) GetNatSource() string`

GetNatSource returns the NatSource field if non-nil, zero value otherwise.

### GetNatSourceOk

`func (o *SysruntimeFlowInfo) GetNatSourceOk() (*string, bool)`

GetNatSourceOk returns a tuple with the NatSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatSource

`func (o *SysruntimeFlowInfo) SetNatSource(v string)`

SetNatSource sets NatSource field to given value.

### HasNatSource

`func (o *SysruntimeFlowInfo) HasNatSource() bool`

HasNatSource returns a boolean if a field has been set.

### GetNatSourcePort

`func (o *SysruntimeFlowInfo) GetNatSourcePort() int64`

GetNatSourcePort returns the NatSourcePort field if non-nil, zero value otherwise.

### GetNatSourcePortOk

`func (o *SysruntimeFlowInfo) GetNatSourcePortOk() (*int64, bool)`

GetNatSourcePortOk returns a tuple with the NatSourcePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatSourcePort

`func (o *SysruntimeFlowInfo) SetNatSourcePort(v int64)`

SetNatSourcePort sets NatSourcePort field to given value.

### HasNatSourcePort

`func (o *SysruntimeFlowInfo) HasNatSourcePort() bool`

HasNatSourcePort returns a boolean if a field has been set.

### GetNatType

`func (o *SysruntimeFlowInfo) GetNatType() string`

GetNatType returns the NatType field if non-nil, zero value otherwise.

### GetNatTypeOk

`func (o *SysruntimeFlowInfo) GetNatTypeOk() (*string, bool)`

GetNatTypeOk returns a tuple with the NatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatType

`func (o *SysruntimeFlowInfo) SetNatType(v string)`

SetNatType sets NatType field to given value.

### HasNatType

`func (o *SysruntimeFlowInfo) HasNatType() bool`

HasNatType returns a boolean if a field has been set.

### GetTcpState

`func (o *SysruntimeFlowInfo) GetTcpState() string`

GetTcpState returns the TcpState field if non-nil, zero value otherwise.

### GetTcpStateOk

`func (o *SysruntimeFlowInfo) GetTcpStateOk() (*string, bool)`

GetTcpStateOk returns a tuple with the TcpState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpState

`func (o *SysruntimeFlowInfo) SetTcpState(v string)`

SetTcpState sets TcpState field to given value.

### HasTcpState

`func (o *SysruntimeFlowInfo) HasTcpState() bool`

HasTcpState returns a boolean if a field has been set.

### GetTimeToAge

`func (o *SysruntimeFlowInfo) GetTimeToAge() int64`

GetTimeToAge returns the TimeToAge field if non-nil, zero value otherwise.

### GetTimeToAgeOk

`func (o *SysruntimeFlowInfo) GetTimeToAgeOk() (*int64, bool)`

GetTimeToAgeOk returns a tuple with the TimeToAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToAge

`func (o *SysruntimeFlowInfo) SetTimeToAge(v int64)`

SetTimeToAge sets TimeToAge field to given value.

### HasTimeToAge

`func (o *SysruntimeFlowInfo) HasTimeToAge() bool`

HasTimeToAge returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


