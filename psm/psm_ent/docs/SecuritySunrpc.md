# SecuritySunrpc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProgramId** | Pointer to **string** | RPC Program identifier. | [optional] 
**Timeout** | Pointer to **string** | Timeout for this program id. Should be a valid time duration. | [optional] 

## Methods

### NewSecuritySunrpc

`func NewSecuritySunrpc() *SecuritySunrpc`

NewSecuritySunrpc instantiates a new SecuritySunrpc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecuritySunrpcWithDefaults

`func NewSecuritySunrpcWithDefaults() *SecuritySunrpc`

NewSecuritySunrpcWithDefaults instantiates a new SecuritySunrpc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProgramId

`func (o *SecuritySunrpc) GetProgramId() string`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *SecuritySunrpc) GetProgramIdOk() (*string, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramId

`func (o *SecuritySunrpc) SetProgramId(v string)`

SetProgramId sets ProgramId field to given value.

### HasProgramId

`func (o *SecuritySunrpc) HasProgramId() bool`

HasProgramId returns a boolean if a field has been set.

### GetTimeout

`func (o *SecuritySunrpc) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *SecuritySunrpc) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *SecuritySunrpc) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *SecuritySunrpc) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


