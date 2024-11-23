# DiagnosticsServicePort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is the name of the port. | [optional] 
**Port** | Pointer to **int32** | Port is port number. | [optional] 

## Methods

### NewDiagnosticsServicePort

`func NewDiagnosticsServicePort() *DiagnosticsServicePort`

NewDiagnosticsServicePort instantiates a new DiagnosticsServicePort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiagnosticsServicePortWithDefaults

`func NewDiagnosticsServicePortWithDefaults() *DiagnosticsServicePort`

NewDiagnosticsServicePortWithDefaults instantiates a new DiagnosticsServicePort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DiagnosticsServicePort) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DiagnosticsServicePort) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DiagnosticsServicePort) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DiagnosticsServicePort) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *DiagnosticsServicePort) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DiagnosticsServicePort) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DiagnosticsServicePort) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *DiagnosticsServicePort) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


