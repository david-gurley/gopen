# RolloutRolloutSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DscMustMatchConstraint** | Pointer to **bool** | When DSCMustMatchConstraint is true, Any DSC which does not match OrderConstraints does not go through rollout. | [optional] 
**DscsOnly** | Pointer to **bool** | Dont upgrade Controller but only upgrade DistributedServiceCards. | [optional] 
**MaxNicFailuresBeforeAbort** | Pointer to **int64** | After these many failures are observed during DSC upgrade, the rollout process stops This setting applies to DSCs only. The controller nodes are rollout first and any failure there stops the rollout of DistributedServiceCards. | [optional] 
**MaxParallel** | Pointer to **int64** | MaxParallel is the maximum number of nodes getting updated at any time This setting is applicable only to DistributedServiceCards. Controller nodes are always upgraded one after another. | [optional] [default to 2]
**OrderConstraints** | Pointer to [**[]LabelsSelector**](LabelsSelector.md) | If specified, this is the sequence in which the DistributedServiceCards are upgraded based on their labels. if a DistributedServiceCard matches multiple constraints, the first one is used. Any DistributedServiceCard which does not match the constraints is upgraded at the end. This order is mainly for the DSCs on Hosts Controller nodes are always rollout one after other. | [optional] 
**Retry** | Pointer to **bool** | If enabled, will retry rollout of failed naples within the maintenance window upto a max of 5 times. | [optional] 
**ScheduledEndTime** | Pointer to **time.Time** | ScheduledEndTime, if specified, after which the rollout is supposed to stop, if not completed by that time Typically represents the end of Maintenance window. (example:\&quot;2002-10-02T15:00:00.05Z\&quot;). | [optional] 
**ScheduledStartTime** | Pointer to **time.Time** | Time, if specified, at which the rollout is supposed to start. (example:\&quot;2002-10-02T15:00:00.05Z\&quot;). | [optional] 
**Strategy** | Pointer to **string** |  | [optional] [default to "linear"]
**Suspend** | Pointer to **bool** | When Set to true, the current rollout is suspended. Existing Nodes/Services/DistributedServiceCards in the middle of rollout continue rollout execution but any Nodes/Services/DistributedServiceCards which has not started Rollout will not be scheduled one. | [optional] 
**UpgradeType** | Pointer to **string** |  | [optional] [default to "graceful"]
**Version** | Pointer to **string** | New Version of the image to rollout to. | [optional] 

## Methods

### NewRolloutRolloutSpec

`func NewRolloutRolloutSpec() *RolloutRolloutSpec`

NewRolloutRolloutSpec instantiates a new RolloutRolloutSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRolloutRolloutSpecWithDefaults

`func NewRolloutRolloutSpecWithDefaults() *RolloutRolloutSpec`

NewRolloutRolloutSpecWithDefaults instantiates a new RolloutRolloutSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDscMustMatchConstraint

`func (o *RolloutRolloutSpec) GetDscMustMatchConstraint() bool`

GetDscMustMatchConstraint returns the DscMustMatchConstraint field if non-nil, zero value otherwise.

### GetDscMustMatchConstraintOk

`func (o *RolloutRolloutSpec) GetDscMustMatchConstraintOk() (*bool, bool)`

GetDscMustMatchConstraintOk returns a tuple with the DscMustMatchConstraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDscMustMatchConstraint

`func (o *RolloutRolloutSpec) SetDscMustMatchConstraint(v bool)`

SetDscMustMatchConstraint sets DscMustMatchConstraint field to given value.

### HasDscMustMatchConstraint

`func (o *RolloutRolloutSpec) HasDscMustMatchConstraint() bool`

HasDscMustMatchConstraint returns a boolean if a field has been set.

### GetDscsOnly

`func (o *RolloutRolloutSpec) GetDscsOnly() bool`

GetDscsOnly returns the DscsOnly field if non-nil, zero value otherwise.

### GetDscsOnlyOk

`func (o *RolloutRolloutSpec) GetDscsOnlyOk() (*bool, bool)`

GetDscsOnlyOk returns a tuple with the DscsOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDscsOnly

`func (o *RolloutRolloutSpec) SetDscsOnly(v bool)`

SetDscsOnly sets DscsOnly field to given value.

### HasDscsOnly

`func (o *RolloutRolloutSpec) HasDscsOnly() bool`

HasDscsOnly returns a boolean if a field has been set.

### GetMaxNicFailuresBeforeAbort

`func (o *RolloutRolloutSpec) GetMaxNicFailuresBeforeAbort() int64`

GetMaxNicFailuresBeforeAbort returns the MaxNicFailuresBeforeAbort field if non-nil, zero value otherwise.

### GetMaxNicFailuresBeforeAbortOk

`func (o *RolloutRolloutSpec) GetMaxNicFailuresBeforeAbortOk() (*int64, bool)`

GetMaxNicFailuresBeforeAbortOk returns a tuple with the MaxNicFailuresBeforeAbort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNicFailuresBeforeAbort

`func (o *RolloutRolloutSpec) SetMaxNicFailuresBeforeAbort(v int64)`

SetMaxNicFailuresBeforeAbort sets MaxNicFailuresBeforeAbort field to given value.

### HasMaxNicFailuresBeforeAbort

`func (o *RolloutRolloutSpec) HasMaxNicFailuresBeforeAbort() bool`

HasMaxNicFailuresBeforeAbort returns a boolean if a field has been set.

### GetMaxParallel

`func (o *RolloutRolloutSpec) GetMaxParallel() int64`

GetMaxParallel returns the MaxParallel field if non-nil, zero value otherwise.

### GetMaxParallelOk

`func (o *RolloutRolloutSpec) GetMaxParallelOk() (*int64, bool)`

GetMaxParallelOk returns a tuple with the MaxParallel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxParallel

`func (o *RolloutRolloutSpec) SetMaxParallel(v int64)`

SetMaxParallel sets MaxParallel field to given value.

### HasMaxParallel

`func (o *RolloutRolloutSpec) HasMaxParallel() bool`

HasMaxParallel returns a boolean if a field has been set.

### GetOrderConstraints

`func (o *RolloutRolloutSpec) GetOrderConstraints() []LabelsSelector`

GetOrderConstraints returns the OrderConstraints field if non-nil, zero value otherwise.

### GetOrderConstraintsOk

`func (o *RolloutRolloutSpec) GetOrderConstraintsOk() (*[]LabelsSelector, bool)`

GetOrderConstraintsOk returns a tuple with the OrderConstraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderConstraints

`func (o *RolloutRolloutSpec) SetOrderConstraints(v []LabelsSelector)`

SetOrderConstraints sets OrderConstraints field to given value.

### HasOrderConstraints

`func (o *RolloutRolloutSpec) HasOrderConstraints() bool`

HasOrderConstraints returns a boolean if a field has been set.

### GetRetry

`func (o *RolloutRolloutSpec) GetRetry() bool`

GetRetry returns the Retry field if non-nil, zero value otherwise.

### GetRetryOk

`func (o *RolloutRolloutSpec) GetRetryOk() (*bool, bool)`

GetRetryOk returns a tuple with the Retry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetry

`func (o *RolloutRolloutSpec) SetRetry(v bool)`

SetRetry sets Retry field to given value.

### HasRetry

`func (o *RolloutRolloutSpec) HasRetry() bool`

HasRetry returns a boolean if a field has been set.

### GetScheduledEndTime

`func (o *RolloutRolloutSpec) GetScheduledEndTime() time.Time`

GetScheduledEndTime returns the ScheduledEndTime field if non-nil, zero value otherwise.

### GetScheduledEndTimeOk

`func (o *RolloutRolloutSpec) GetScheduledEndTimeOk() (*time.Time, bool)`

GetScheduledEndTimeOk returns a tuple with the ScheduledEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledEndTime

`func (o *RolloutRolloutSpec) SetScheduledEndTime(v time.Time)`

SetScheduledEndTime sets ScheduledEndTime field to given value.

### HasScheduledEndTime

`func (o *RolloutRolloutSpec) HasScheduledEndTime() bool`

HasScheduledEndTime returns a boolean if a field has been set.

### GetScheduledStartTime

`func (o *RolloutRolloutSpec) GetScheduledStartTime() time.Time`

GetScheduledStartTime returns the ScheduledStartTime field if non-nil, zero value otherwise.

### GetScheduledStartTimeOk

`func (o *RolloutRolloutSpec) GetScheduledStartTimeOk() (*time.Time, bool)`

GetScheduledStartTimeOk returns a tuple with the ScheduledStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledStartTime

`func (o *RolloutRolloutSpec) SetScheduledStartTime(v time.Time)`

SetScheduledStartTime sets ScheduledStartTime field to given value.

### HasScheduledStartTime

`func (o *RolloutRolloutSpec) HasScheduledStartTime() bool`

HasScheduledStartTime returns a boolean if a field has been set.

### GetStrategy

`func (o *RolloutRolloutSpec) GetStrategy() string`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *RolloutRolloutSpec) GetStrategyOk() (*string, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *RolloutRolloutSpec) SetStrategy(v string)`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *RolloutRolloutSpec) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.

### GetSuspend

`func (o *RolloutRolloutSpec) GetSuspend() bool`

GetSuspend returns the Suspend field if non-nil, zero value otherwise.

### GetSuspendOk

`func (o *RolloutRolloutSpec) GetSuspendOk() (*bool, bool)`

GetSuspendOk returns a tuple with the Suspend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspend

`func (o *RolloutRolloutSpec) SetSuspend(v bool)`

SetSuspend sets Suspend field to given value.

### HasSuspend

`func (o *RolloutRolloutSpec) HasSuspend() bool`

HasSuspend returns a boolean if a field has been set.

### GetUpgradeType

`func (o *RolloutRolloutSpec) GetUpgradeType() string`

GetUpgradeType returns the UpgradeType field if non-nil, zero value otherwise.

### GetUpgradeTypeOk

`func (o *RolloutRolloutSpec) GetUpgradeTypeOk() (*string, bool)`

GetUpgradeTypeOk returns a tuple with the UpgradeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeType

`func (o *RolloutRolloutSpec) SetUpgradeType(v string)`

SetUpgradeType sets UpgradeType field to given value.

### HasUpgradeType

`func (o *RolloutRolloutSpec) HasUpgradeType() bool`

HasUpgradeType returns a boolean if a field has been set.

### GetVersion

`func (o *RolloutRolloutSpec) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RolloutRolloutSpec) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RolloutRolloutSpec) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *RolloutRolloutSpec) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


