# SecurityPropagationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DscStatus** | Pointer to [**[]SecurityDSCStatus**](SecurityDSCStatus.md) | list of DSCs status. | [optional] 
**GenerationId** | Pointer to **string** | The Generation ID this status is for. | [optional] 
**MinVersion** | Pointer to **string** | The Version running on the slowest Naples. | [optional] 
**Pending** | Pointer to **int32** | Number of Naples pending. If this is 0 it can be assumed that everything is up to date. | [optional] 
**PendingDscs** | Pointer to **[]string** | list of DSCs where propagation did not complete. | [optional] 
**Status** | Pointer to **string** | Textual description of propagation status. | [optional] 
**Updated** | Pointer to **int32** | The number of Naples that this version has already been pushed to. | [optional] 

## Methods

### NewSecurityPropagationStatus

`func NewSecurityPropagationStatus() *SecurityPropagationStatus`

NewSecurityPropagationStatus instantiates a new SecurityPropagationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityPropagationStatusWithDefaults

`func NewSecurityPropagationStatusWithDefaults() *SecurityPropagationStatus`

NewSecurityPropagationStatusWithDefaults instantiates a new SecurityPropagationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDscStatus

`func (o *SecurityPropagationStatus) GetDscStatus() []SecurityDSCStatus`

GetDscStatus returns the DscStatus field if non-nil, zero value otherwise.

### GetDscStatusOk

`func (o *SecurityPropagationStatus) GetDscStatusOk() (*[]SecurityDSCStatus, bool)`

GetDscStatusOk returns a tuple with the DscStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDscStatus

`func (o *SecurityPropagationStatus) SetDscStatus(v []SecurityDSCStatus)`

SetDscStatus sets DscStatus field to given value.

### HasDscStatus

`func (o *SecurityPropagationStatus) HasDscStatus() bool`

HasDscStatus returns a boolean if a field has been set.

### GetGenerationId

`func (o *SecurityPropagationStatus) GetGenerationId() string`

GetGenerationId returns the GenerationId field if non-nil, zero value otherwise.

### GetGenerationIdOk

`func (o *SecurityPropagationStatus) GetGenerationIdOk() (*string, bool)`

GetGenerationIdOk returns a tuple with the GenerationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerationId

`func (o *SecurityPropagationStatus) SetGenerationId(v string)`

SetGenerationId sets GenerationId field to given value.

### HasGenerationId

`func (o *SecurityPropagationStatus) HasGenerationId() bool`

HasGenerationId returns a boolean if a field has been set.

### GetMinVersion

`func (o *SecurityPropagationStatus) GetMinVersion() string`

GetMinVersion returns the MinVersion field if non-nil, zero value otherwise.

### GetMinVersionOk

`func (o *SecurityPropagationStatus) GetMinVersionOk() (*string, bool)`

GetMinVersionOk returns a tuple with the MinVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinVersion

`func (o *SecurityPropagationStatus) SetMinVersion(v string)`

SetMinVersion sets MinVersion field to given value.

### HasMinVersion

`func (o *SecurityPropagationStatus) HasMinVersion() bool`

HasMinVersion returns a boolean if a field has been set.

### GetPending

`func (o *SecurityPropagationStatus) GetPending() int32`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *SecurityPropagationStatus) GetPendingOk() (*int32, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *SecurityPropagationStatus) SetPending(v int32)`

SetPending sets Pending field to given value.

### HasPending

`func (o *SecurityPropagationStatus) HasPending() bool`

HasPending returns a boolean if a field has been set.

### GetPendingDscs

`func (o *SecurityPropagationStatus) GetPendingDscs() []string`

GetPendingDscs returns the PendingDscs field if non-nil, zero value otherwise.

### GetPendingDscsOk

`func (o *SecurityPropagationStatus) GetPendingDscsOk() (*[]string, bool)`

GetPendingDscsOk returns a tuple with the PendingDscs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingDscs

`func (o *SecurityPropagationStatus) SetPendingDscs(v []string)`

SetPendingDscs sets PendingDscs field to given value.

### HasPendingDscs

`func (o *SecurityPropagationStatus) HasPendingDscs() bool`

HasPendingDscs returns a boolean if a field has been set.

### GetStatus

`func (o *SecurityPropagationStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SecurityPropagationStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SecurityPropagationStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SecurityPropagationStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdated

`func (o *SecurityPropagationStatus) GetUpdated() int32`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *SecurityPropagationStatus) GetUpdatedOk() (*int32, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *SecurityPropagationStatus) SetUpdated(v int32)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *SecurityPropagationStatus) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


