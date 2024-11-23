# ClusterSnapshotRestoreStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupSnapshotPath** | Pointer to **string** |  | [optional] 
**EndTime** | Pointer to **time.Time** |  | [optional] 
**StartTime** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] [default to "unknown"]

## Methods

### NewClusterSnapshotRestoreStatus

`func NewClusterSnapshotRestoreStatus() *ClusterSnapshotRestoreStatus`

NewClusterSnapshotRestoreStatus instantiates a new ClusterSnapshotRestoreStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterSnapshotRestoreStatusWithDefaults

`func NewClusterSnapshotRestoreStatusWithDefaults() *ClusterSnapshotRestoreStatus`

NewClusterSnapshotRestoreStatusWithDefaults instantiates a new ClusterSnapshotRestoreStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupSnapshotPath

`func (o *ClusterSnapshotRestoreStatus) GetBackupSnapshotPath() string`

GetBackupSnapshotPath returns the BackupSnapshotPath field if non-nil, zero value otherwise.

### GetBackupSnapshotPathOk

`func (o *ClusterSnapshotRestoreStatus) GetBackupSnapshotPathOk() (*string, bool)`

GetBackupSnapshotPathOk returns a tuple with the BackupSnapshotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupSnapshotPath

`func (o *ClusterSnapshotRestoreStatus) SetBackupSnapshotPath(v string)`

SetBackupSnapshotPath sets BackupSnapshotPath field to given value.

### HasBackupSnapshotPath

`func (o *ClusterSnapshotRestoreStatus) HasBackupSnapshotPath() bool`

HasBackupSnapshotPath returns a boolean if a field has been set.

### GetEndTime

`func (o *ClusterSnapshotRestoreStatus) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ClusterSnapshotRestoreStatus) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ClusterSnapshotRestoreStatus) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *ClusterSnapshotRestoreStatus) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetStartTime

`func (o *ClusterSnapshotRestoreStatus) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ClusterSnapshotRestoreStatus) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ClusterSnapshotRestoreStatus) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ClusterSnapshotRestoreStatus) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *ClusterSnapshotRestoreStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterSnapshotRestoreStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterSnapshotRestoreStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClusterSnapshotRestoreStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


