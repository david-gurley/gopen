# SecurityMsrpc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProgramUuid** | Pointer to **string** | MSRPC Program identifier. | [optional] 
**Timeout** | Pointer to **string** | Timeout for this program id. Should be a valid time duration. | [optional] 

## Methods

### NewSecurityMsrpc

`func NewSecurityMsrpc() *SecurityMsrpc`

NewSecurityMsrpc instantiates a new SecurityMsrpc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityMsrpcWithDefaults

`func NewSecurityMsrpcWithDefaults() *SecurityMsrpc`

NewSecurityMsrpcWithDefaults instantiates a new SecurityMsrpc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProgramUuid

`func (o *SecurityMsrpc) GetProgramUuid() string`

GetProgramUuid returns the ProgramUuid field if non-nil, zero value otherwise.

### GetProgramUuidOk

`func (o *SecurityMsrpc) GetProgramUuidOk() (*string, bool)`

GetProgramUuidOk returns a tuple with the ProgramUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramUuid

`func (o *SecurityMsrpc) SetProgramUuid(v string)`

SetProgramUuid sets ProgramUuid field to given value.

### HasProgramUuid

`func (o *SecurityMsrpc) HasProgramUuid() bool`

HasProgramUuid returns a boolean if a field has been set.

### GetTimeout

`func (o *SecurityMsrpc) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *SecurityMsrpc) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *SecurityMsrpc) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *SecurityMsrpc) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


