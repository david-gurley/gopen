# SecurityAppSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alg** | Pointer to [**SecurityALG**](securityALG.md) |  | [optional] 
**ProtoPorts** | Pointer to [**[]SecurityProtoPort**](SecurityProtoPort.md) | list of (protocol, ports) pairs. | [optional] 
**Timeout** | Pointer to **string** | Timeout specifies for how long the connection be kept before removing the flow entry, specified in string as &#39;200s&#39;, &#39;5m&#39;, etc. Should be a valid time duration. | [optional] 

## Methods

### NewSecurityAppSpec

`func NewSecurityAppSpec() *SecurityAppSpec`

NewSecurityAppSpec instantiates a new SecurityAppSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityAppSpecWithDefaults

`func NewSecurityAppSpecWithDefaults() *SecurityAppSpec`

NewSecurityAppSpecWithDefaults instantiates a new SecurityAppSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlg

`func (o *SecurityAppSpec) GetAlg() SecurityALG`

GetAlg returns the Alg field if non-nil, zero value otherwise.

### GetAlgOk

`func (o *SecurityAppSpec) GetAlgOk() (*SecurityALG, bool)`

GetAlgOk returns a tuple with the Alg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlg

`func (o *SecurityAppSpec) SetAlg(v SecurityALG)`

SetAlg sets Alg field to given value.

### HasAlg

`func (o *SecurityAppSpec) HasAlg() bool`

HasAlg returns a boolean if a field has been set.

### GetProtoPorts

`func (o *SecurityAppSpec) GetProtoPorts() []SecurityProtoPort`

GetProtoPorts returns the ProtoPorts field if non-nil, zero value otherwise.

### GetProtoPortsOk

`func (o *SecurityAppSpec) GetProtoPortsOk() (*[]SecurityProtoPort, bool)`

GetProtoPortsOk returns a tuple with the ProtoPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtoPorts

`func (o *SecurityAppSpec) SetProtoPorts(v []SecurityProtoPort)`

SetProtoPorts sets ProtoPorts field to given value.

### HasProtoPorts

`func (o *SecurityAppSpec) HasProtoPorts() bool`

HasProtoPorts returns a boolean if a field has been set.

### GetTimeout

`func (o *SecurityAppSpec) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *SecurityAppSpec) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *SecurityAppSpec) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *SecurityAppSpec) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


