# ClusterHostStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdmittedDscs** | Pointer to **[]string** | AdmittedDSCs contains a list of admitted DistributedServiceCards that are on this host. | [optional] 
**MirrorSessions** | Pointer to **[]string** | MirrorSessions list of mirror sessions enabled on this host. | [optional] 

## Methods

### NewClusterHostStatus

`func NewClusterHostStatus() *ClusterHostStatus`

NewClusterHostStatus instantiates a new ClusterHostStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterHostStatusWithDefaults

`func NewClusterHostStatusWithDefaults() *ClusterHostStatus`

NewClusterHostStatusWithDefaults instantiates a new ClusterHostStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdmittedDscs

`func (o *ClusterHostStatus) GetAdmittedDscs() []string`

GetAdmittedDscs returns the AdmittedDscs field if non-nil, zero value otherwise.

### GetAdmittedDscsOk

`func (o *ClusterHostStatus) GetAdmittedDscsOk() (*[]string, bool)`

GetAdmittedDscsOk returns a tuple with the AdmittedDscs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmittedDscs

`func (o *ClusterHostStatus) SetAdmittedDscs(v []string)`

SetAdmittedDscs sets AdmittedDscs field to given value.

### HasAdmittedDscs

`func (o *ClusterHostStatus) HasAdmittedDscs() bool`

HasAdmittedDscs returns a boolean if a field has been set.

### GetMirrorSessions

`func (o *ClusterHostStatus) GetMirrorSessions() []string`

GetMirrorSessions returns the MirrorSessions field if non-nil, zero value otherwise.

### GetMirrorSessionsOk

`func (o *ClusterHostStatus) GetMirrorSessionsOk() (*[]string, bool)`

GetMirrorSessionsOk returns a tuple with the MirrorSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirrorSessions

`func (o *ClusterHostStatus) SetMirrorSessions(v []string)`

SetMirrorSessions sets MirrorSessions field to given value.

### HasMirrorSessions

`func (o *ClusterHostStatus) HasMirrorSessions() bool`

HasMirrorSessions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


