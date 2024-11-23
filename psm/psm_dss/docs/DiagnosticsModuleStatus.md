# DiagnosticsModuleStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | Category specifies whether process is part of Venice(controller) or Naples(io) subsystem. | [optional] [default to "venice"]
**LastRestartReason** | Pointer to **string** | Arbitrary string, json, backtrace, etc. offering clues for restart. | [optional] 
**LastStart** | Pointer to **time.Time** | Last start time. | [optional] 
**MacAddress** | Pointer to **string** | MACAddress of the smart nic on which this module runs. | [optional] 
**Module** | Pointer to **string** | Module is the name of the process/container. | [optional] 
**Node** | Pointer to **string** | Node on which this process is running. | [optional] 
**RestartCount** | Pointer to **int32** | Number of times process got restarted. zero if never restarted. | [optional] 
**Service** | Pointer to **string** | Service is the name of the service/pod this process is part of. | [optional] 
**ServicePorts** | Pointer to [**[]DiagnosticsServicePort**](DiagnosticsServicePort.md) | Ports on which this process is listening. | [optional] 

## Methods

### NewDiagnosticsModuleStatus

`func NewDiagnosticsModuleStatus() *DiagnosticsModuleStatus`

NewDiagnosticsModuleStatus instantiates a new DiagnosticsModuleStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiagnosticsModuleStatusWithDefaults

`func NewDiagnosticsModuleStatusWithDefaults() *DiagnosticsModuleStatus`

NewDiagnosticsModuleStatusWithDefaults instantiates a new DiagnosticsModuleStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *DiagnosticsModuleStatus) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *DiagnosticsModuleStatus) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *DiagnosticsModuleStatus) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *DiagnosticsModuleStatus) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetLastRestartReason

`func (o *DiagnosticsModuleStatus) GetLastRestartReason() string`

GetLastRestartReason returns the LastRestartReason field if non-nil, zero value otherwise.

### GetLastRestartReasonOk

`func (o *DiagnosticsModuleStatus) GetLastRestartReasonOk() (*string, bool)`

GetLastRestartReasonOk returns a tuple with the LastRestartReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRestartReason

`func (o *DiagnosticsModuleStatus) SetLastRestartReason(v string)`

SetLastRestartReason sets LastRestartReason field to given value.

### HasLastRestartReason

`func (o *DiagnosticsModuleStatus) HasLastRestartReason() bool`

HasLastRestartReason returns a boolean if a field has been set.

### GetLastStart

`func (o *DiagnosticsModuleStatus) GetLastStart() time.Time`

GetLastStart returns the LastStart field if non-nil, zero value otherwise.

### GetLastStartOk

`func (o *DiagnosticsModuleStatus) GetLastStartOk() (*time.Time, bool)`

GetLastStartOk returns a tuple with the LastStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStart

`func (o *DiagnosticsModuleStatus) SetLastStart(v time.Time)`

SetLastStart sets LastStart field to given value.

### HasLastStart

`func (o *DiagnosticsModuleStatus) HasLastStart() bool`

HasLastStart returns a boolean if a field has been set.

### GetMacAddress

`func (o *DiagnosticsModuleStatus) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *DiagnosticsModuleStatus) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *DiagnosticsModuleStatus) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *DiagnosticsModuleStatus) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetModule

`func (o *DiagnosticsModuleStatus) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *DiagnosticsModuleStatus) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *DiagnosticsModuleStatus) SetModule(v string)`

SetModule sets Module field to given value.

### HasModule

`func (o *DiagnosticsModuleStatus) HasModule() bool`

HasModule returns a boolean if a field has been set.

### GetNode

`func (o *DiagnosticsModuleStatus) GetNode() string`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *DiagnosticsModuleStatus) GetNodeOk() (*string, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *DiagnosticsModuleStatus) SetNode(v string)`

SetNode sets Node field to given value.

### HasNode

`func (o *DiagnosticsModuleStatus) HasNode() bool`

HasNode returns a boolean if a field has been set.

### GetRestartCount

`func (o *DiagnosticsModuleStatus) GetRestartCount() int32`

GetRestartCount returns the RestartCount field if non-nil, zero value otherwise.

### GetRestartCountOk

`func (o *DiagnosticsModuleStatus) GetRestartCountOk() (*int32, bool)`

GetRestartCountOk returns a tuple with the RestartCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartCount

`func (o *DiagnosticsModuleStatus) SetRestartCount(v int32)`

SetRestartCount sets RestartCount field to given value.

### HasRestartCount

`func (o *DiagnosticsModuleStatus) HasRestartCount() bool`

HasRestartCount returns a boolean if a field has been set.

### GetService

`func (o *DiagnosticsModuleStatus) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *DiagnosticsModuleStatus) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *DiagnosticsModuleStatus) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *DiagnosticsModuleStatus) HasService() bool`

HasService returns a boolean if a field has been set.

### GetServicePorts

`func (o *DiagnosticsModuleStatus) GetServicePorts() []DiagnosticsServicePort`

GetServicePorts returns the ServicePorts field if non-nil, zero value otherwise.

### GetServicePortsOk

`func (o *DiagnosticsModuleStatus) GetServicePortsOk() (*[]DiagnosticsServicePort, bool)`

GetServicePortsOk returns a tuple with the ServicePorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePorts

`func (o *DiagnosticsModuleStatus) SetServicePorts(v []DiagnosticsServicePort)`

SetServicePorts sets ServicePorts field to given value.

### HasServicePorts

`func (o *DiagnosticsModuleStatus) HasServicePorts() bool`

HasServicePorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


