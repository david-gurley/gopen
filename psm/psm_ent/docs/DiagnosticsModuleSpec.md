# DiagnosticsModuleSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Args** | Pointer to **[]string** | Args are command line arguments passed to the module. | [optional] 
**EnableTrace** | Pointer to **bool** | EnableTrace enables traces for a module. Default is false. | [optional] 
**LogLevel** | Pointer to **string** | LogLevel is the logging level of this module. Default is Info. | [optional] [default to "info"]

## Methods

### NewDiagnosticsModuleSpec

`func NewDiagnosticsModuleSpec() *DiagnosticsModuleSpec`

NewDiagnosticsModuleSpec instantiates a new DiagnosticsModuleSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiagnosticsModuleSpecWithDefaults

`func NewDiagnosticsModuleSpecWithDefaults() *DiagnosticsModuleSpec`

NewDiagnosticsModuleSpecWithDefaults instantiates a new DiagnosticsModuleSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArgs

`func (o *DiagnosticsModuleSpec) GetArgs() []string`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *DiagnosticsModuleSpec) GetArgsOk() (*[]string, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *DiagnosticsModuleSpec) SetArgs(v []string)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *DiagnosticsModuleSpec) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetEnableTrace

`func (o *DiagnosticsModuleSpec) GetEnableTrace() bool`

GetEnableTrace returns the EnableTrace field if non-nil, zero value otherwise.

### GetEnableTraceOk

`func (o *DiagnosticsModuleSpec) GetEnableTraceOk() (*bool, bool)`

GetEnableTraceOk returns a tuple with the EnableTrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTrace

`func (o *DiagnosticsModuleSpec) SetEnableTrace(v bool)`

SetEnableTrace sets EnableTrace field to given value.

### HasEnableTrace

`func (o *DiagnosticsModuleSpec) HasEnableTrace() bool`

HasEnableTrace returns a boolean if a field has been set.

### GetLogLevel

`func (o *DiagnosticsModuleSpec) GetLogLevel() string`

GetLogLevel returns the LogLevel field if non-nil, zero value otherwise.

### GetLogLevelOk

`func (o *DiagnosticsModuleSpec) GetLogLevelOk() (*string, bool)`

GetLogLevelOk returns a tuple with the LogLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLevel

`func (o *DiagnosticsModuleSpec) SetLogLevel(v string)`

SetLogLevel sets LogLevel field to given value.

### HasLogLevel

`func (o *DiagnosticsModuleSpec) HasLogLevel() bool`

HasLogLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


