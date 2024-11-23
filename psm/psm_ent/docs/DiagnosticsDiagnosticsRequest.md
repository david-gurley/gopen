# DiagnosticsDiagnosticsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**ApiObjectMeta**](apiObjectMeta.md) |  | [optional] 
**Parameters** | Pointer to **map[string]string** | Parameters to be passed for a diagnostic query. | [optional] 
**Query** | Pointer to **string** | Query is type of diagnostic information requested like Profile, Log. This string is service specific and meaning is assigned by the service. | [optional] 
**ServicePort** | Pointer to [**DiagnosticsServicePort**](diagnosticsServicePort.md) |  | [optional] 

## Methods

### NewDiagnosticsDiagnosticsRequest

`func NewDiagnosticsDiagnosticsRequest() *DiagnosticsDiagnosticsRequest`

NewDiagnosticsDiagnosticsRequest instantiates a new DiagnosticsDiagnosticsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiagnosticsDiagnosticsRequestWithDefaults

`func NewDiagnosticsDiagnosticsRequestWithDefaults() *DiagnosticsDiagnosticsRequest`

NewDiagnosticsDiagnosticsRequestWithDefaults instantiates a new DiagnosticsDiagnosticsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *DiagnosticsDiagnosticsRequest) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *DiagnosticsDiagnosticsRequest) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *DiagnosticsDiagnosticsRequest) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *DiagnosticsDiagnosticsRequest) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *DiagnosticsDiagnosticsRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *DiagnosticsDiagnosticsRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *DiagnosticsDiagnosticsRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *DiagnosticsDiagnosticsRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMeta

`func (o *DiagnosticsDiagnosticsRequest) GetMeta() ApiObjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *DiagnosticsDiagnosticsRequest) GetMetaOk() (*ApiObjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *DiagnosticsDiagnosticsRequest) SetMeta(v ApiObjectMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *DiagnosticsDiagnosticsRequest) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetParameters

`func (o *DiagnosticsDiagnosticsRequest) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *DiagnosticsDiagnosticsRequest) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *DiagnosticsDiagnosticsRequest) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *DiagnosticsDiagnosticsRequest) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetQuery

`func (o *DiagnosticsDiagnosticsRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *DiagnosticsDiagnosticsRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *DiagnosticsDiagnosticsRequest) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *DiagnosticsDiagnosticsRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetServicePort

`func (o *DiagnosticsDiagnosticsRequest) GetServicePort() DiagnosticsServicePort`

GetServicePort returns the ServicePort field if non-nil, zero value otherwise.

### GetServicePortOk

`func (o *DiagnosticsDiagnosticsRequest) GetServicePortOk() (*DiagnosticsServicePort, bool)`

GetServicePortOk returns a tuple with the ServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePort

`func (o *DiagnosticsDiagnosticsRequest) SetServicePort(v DiagnosticsServicePort)`

SetServicePort sets ServicePort field to given value.

### HasServicePort

`func (o *DiagnosticsDiagnosticsRequest) HasServicePort() bool`

HasServicePort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


