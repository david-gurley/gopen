# NetworkTransceiverStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CableType** | Pointer to **string** |  | [optional] [default to "none"]
**Pid** | Pointer to **string** |  | [optional] [default to "unknown"]
**State** | Pointer to **string** |  | [optional] [default to "state_na"]

## Methods

### NewNetworkTransceiverStatus

`func NewNetworkTransceiverStatus() *NetworkTransceiverStatus`

NewNetworkTransceiverStatus instantiates a new NetworkTransceiverStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkTransceiverStatusWithDefaults

`func NewNetworkTransceiverStatusWithDefaults() *NetworkTransceiverStatus`

NewNetworkTransceiverStatusWithDefaults instantiates a new NetworkTransceiverStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCableType

`func (o *NetworkTransceiverStatus) GetCableType() string`

GetCableType returns the CableType field if non-nil, zero value otherwise.

### GetCableTypeOk

`func (o *NetworkTransceiverStatus) GetCableTypeOk() (*string, bool)`

GetCableTypeOk returns a tuple with the CableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCableType

`func (o *NetworkTransceiverStatus) SetCableType(v string)`

SetCableType sets CableType field to given value.

### HasCableType

`func (o *NetworkTransceiverStatus) HasCableType() bool`

HasCableType returns a boolean if a field has been set.

### GetPid

`func (o *NetworkTransceiverStatus) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *NetworkTransceiverStatus) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *NetworkTransceiverStatus) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *NetworkTransceiverStatus) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetState

`func (o *NetworkTransceiverStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NetworkTransceiverStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NetworkTransceiverStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *NetworkTransceiverStatus) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


