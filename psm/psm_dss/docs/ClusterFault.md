# ClusterFault

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the fault occured. | [optional] 
**LastOccuredTime** | Pointer to **string** | Time at which the fault occured. | [optional] 
**Mitigation** | Pointer to **string** | Mitigation action,if any possible. | [optional] [default to "system-reboot"]
**Severity** | Pointer to **string** | Severity of fault occured at DSS end. | [optional] [default to "info"]

## Methods

### NewClusterFault

`func NewClusterFault() *ClusterFault`

NewClusterFault instantiates a new ClusterFault object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterFaultWithDefaults

`func NewClusterFaultWithDefaults() *ClusterFault`

NewClusterFaultWithDefaults instantiates a new ClusterFault object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ClusterFault) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClusterFault) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClusterFault) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClusterFault) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLastOccuredTime

`func (o *ClusterFault) GetLastOccuredTime() string`

GetLastOccuredTime returns the LastOccuredTime field if non-nil, zero value otherwise.

### GetLastOccuredTimeOk

`func (o *ClusterFault) GetLastOccuredTimeOk() (*string, bool)`

GetLastOccuredTimeOk returns a tuple with the LastOccuredTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOccuredTime

`func (o *ClusterFault) SetLastOccuredTime(v string)`

SetLastOccuredTime sets LastOccuredTime field to given value.

### HasLastOccuredTime

`func (o *ClusterFault) HasLastOccuredTime() bool`

HasLastOccuredTime returns a boolean if a field has been set.

### GetMitigation

`func (o *ClusterFault) GetMitigation() string`

GetMitigation returns the Mitigation field if non-nil, zero value otherwise.

### GetMitigationOk

`func (o *ClusterFault) GetMitigationOk() (*string, bool)`

GetMitigationOk returns a tuple with the Mitigation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitigation

`func (o *ClusterFault) SetMitigation(v string)`

SetMitigation sets Mitigation field to given value.

### HasMitigation

`func (o *ClusterFault) HasMitigation() bool`

HasMitigation returns a boolean if a field has been set.

### GetSeverity

`func (o *ClusterFault) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ClusterFault) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ClusterFault) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ClusterFault) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


