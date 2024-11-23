# SecurityIcmp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | ICMP Code is sub-command for a given ICMP Type. | [optional] 
**Type** | Pointer to **string** | ICMP Type. | [optional] 

## Methods

### NewSecurityIcmp

`func NewSecurityIcmp() *SecurityIcmp`

NewSecurityIcmp instantiates a new SecurityIcmp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityIcmpWithDefaults

`func NewSecurityIcmpWithDefaults() *SecurityIcmp`

NewSecurityIcmpWithDefaults instantiates a new SecurityIcmp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *SecurityIcmp) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *SecurityIcmp) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *SecurityIcmp) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *SecurityIcmp) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetType

`func (o *SecurityIcmp) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SecurityIcmp) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SecurityIcmp) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SecurityIcmp) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


