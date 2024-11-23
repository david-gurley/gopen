# SysruntimeHWConnectionSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InitiatorFlow** | Pointer to [**SysruntimeFlowSpec**](sysruntimeFlowSpec.md) |  | [optional] 
**PeerInitiatorFlow** | Pointer to [**SysruntimeFlowSpec**](sysruntimeFlowSpec.md) |  | [optional] 
**PeerResponderFlow** | Pointer to [**SysruntimeFlowSpec**](sysruntimeFlowSpec.md) |  | [optional] 
**ResponderFlow** | Pointer to [**SysruntimeFlowSpec**](sysruntimeFlowSpec.md) |  | [optional] 
**SessionId** | Pointer to **string** | sessionId is unique session identifier. | [optional] 

## Methods

### NewSysruntimeHWConnectionSpec

`func NewSysruntimeHWConnectionSpec() *SysruntimeHWConnectionSpec`

NewSysruntimeHWConnectionSpec instantiates a new SysruntimeHWConnectionSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSysruntimeHWConnectionSpecWithDefaults

`func NewSysruntimeHWConnectionSpecWithDefaults() *SysruntimeHWConnectionSpec`

NewSysruntimeHWConnectionSpecWithDefaults instantiates a new SysruntimeHWConnectionSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInitiatorFlow

`func (o *SysruntimeHWConnectionSpec) GetInitiatorFlow() SysruntimeFlowSpec`

GetInitiatorFlow returns the InitiatorFlow field if non-nil, zero value otherwise.

### GetInitiatorFlowOk

`func (o *SysruntimeHWConnectionSpec) GetInitiatorFlowOk() (*SysruntimeFlowSpec, bool)`

GetInitiatorFlowOk returns a tuple with the InitiatorFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorFlow

`func (o *SysruntimeHWConnectionSpec) SetInitiatorFlow(v SysruntimeFlowSpec)`

SetInitiatorFlow sets InitiatorFlow field to given value.

### HasInitiatorFlow

`func (o *SysruntimeHWConnectionSpec) HasInitiatorFlow() bool`

HasInitiatorFlow returns a boolean if a field has been set.

### GetPeerInitiatorFlow

`func (o *SysruntimeHWConnectionSpec) GetPeerInitiatorFlow() SysruntimeFlowSpec`

GetPeerInitiatorFlow returns the PeerInitiatorFlow field if non-nil, zero value otherwise.

### GetPeerInitiatorFlowOk

`func (o *SysruntimeHWConnectionSpec) GetPeerInitiatorFlowOk() (*SysruntimeFlowSpec, bool)`

GetPeerInitiatorFlowOk returns a tuple with the PeerInitiatorFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerInitiatorFlow

`func (o *SysruntimeHWConnectionSpec) SetPeerInitiatorFlow(v SysruntimeFlowSpec)`

SetPeerInitiatorFlow sets PeerInitiatorFlow field to given value.

### HasPeerInitiatorFlow

`func (o *SysruntimeHWConnectionSpec) HasPeerInitiatorFlow() bool`

HasPeerInitiatorFlow returns a boolean if a field has been set.

### GetPeerResponderFlow

`func (o *SysruntimeHWConnectionSpec) GetPeerResponderFlow() SysruntimeFlowSpec`

GetPeerResponderFlow returns the PeerResponderFlow field if non-nil, zero value otherwise.

### GetPeerResponderFlowOk

`func (o *SysruntimeHWConnectionSpec) GetPeerResponderFlowOk() (*SysruntimeFlowSpec, bool)`

GetPeerResponderFlowOk returns a tuple with the PeerResponderFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerResponderFlow

`func (o *SysruntimeHWConnectionSpec) SetPeerResponderFlow(v SysruntimeFlowSpec)`

SetPeerResponderFlow sets PeerResponderFlow field to given value.

### HasPeerResponderFlow

`func (o *SysruntimeHWConnectionSpec) HasPeerResponderFlow() bool`

HasPeerResponderFlow returns a boolean if a field has been set.

### GetResponderFlow

`func (o *SysruntimeHWConnectionSpec) GetResponderFlow() SysruntimeFlowSpec`

GetResponderFlow returns the ResponderFlow field if non-nil, zero value otherwise.

### GetResponderFlowOk

`func (o *SysruntimeHWConnectionSpec) GetResponderFlowOk() (*SysruntimeFlowSpec, bool)`

GetResponderFlowOk returns a tuple with the ResponderFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponderFlow

`func (o *SysruntimeHWConnectionSpec) SetResponderFlow(v SysruntimeFlowSpec)`

SetResponderFlow sets ResponderFlow field to given value.

### HasResponderFlow

`func (o *SysruntimeHWConnectionSpec) HasResponderFlow() bool`

HasResponderFlow returns a boolean if a field has been set.

### GetSessionId

`func (o *SysruntimeHWConnectionSpec) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *SysruntimeHWConnectionSpec) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *SysruntimeHWConnectionSpec) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *SysruntimeHWConnectionSpec) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


