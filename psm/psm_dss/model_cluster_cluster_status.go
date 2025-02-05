/*
 * Cluster API reference
 *
 * Service name  
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package psm_dss

import (
	"encoding/json"
	"time"
)

// ClusterClusterStatus ClusterStatus contains the current state of the Cluster.
type ClusterClusterStatus struct {
	// AuthBootstrapped indicates whether the Cluster has Completed BootStrap of Auth.
	AuthBootstrapped *bool `json:"auth-bootstrapped,omitempty"`
	// List of current cluster conditions.
	Conditions *[]ClusterClusterCondition `json:"conditions,omitempty"`
	// CurrentTime is current time of a cluster.
	CurrentTime *time.Time `json:"current-time,omitempty"`
	// LastLeaderTransitionTime is when the leadership changed last time.
	LastLeaderTransitionTime *time.Time `json:"last-leader-transition-time,omitempty"`
	// Leader contains the node name of the cluster leader.
	Leader *string `json:"leader,omitempty"`
	QuorumStatus *ClusterQuorumStatus `json:"quorum-status,omitempty"`
	// RecoveryKeysDownloaded indicates whether keys have been downloaded since the cluster has been bootstrapped.
	RecoveryKeysDownloaded *bool `json:"recovery-keys-downloaded,omitempty"`
}

// NewClusterClusterStatus instantiates a new ClusterClusterStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterClusterStatus() *ClusterClusterStatus {
	this := ClusterClusterStatus{}
	return &this
}

// NewClusterClusterStatusWithDefaults instantiates a new ClusterClusterStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterClusterStatusWithDefaults() *ClusterClusterStatus {
	this := ClusterClusterStatus{}
	return &this
}

// GetAuthBootstrapped returns the AuthBootstrapped field value if set, zero value otherwise.
func (o *ClusterClusterStatus) GetAuthBootstrapped() bool {
	if o == nil || o.AuthBootstrapped == nil {
		var ret bool
		return ret
	}
	return *o.AuthBootstrapped
}

// GetAuthBootstrappedOk returns a tuple with the AuthBootstrapped field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterClusterStatus) GetAuthBootstrappedOk() (*bool, bool) {
	if o == nil || o.AuthBootstrapped == nil {
		return nil, false
	}
	return o.AuthBootstrapped, true
}

// HasAuthBootstrapped returns a boolean if a field has been set.
func (o *ClusterClusterStatus) HasAuthBootstrapped() bool {
	if o != nil && o.AuthBootstrapped != nil {
		return true
	}

	return false
}

// SetAuthBootstrapped gets a reference to the given bool and assigns it to the AuthBootstrapped field.
func (o *ClusterClusterStatus) SetAuthBootstrapped(v bool) {
	o.AuthBootstrapped = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *ClusterClusterStatus) GetConditions() []ClusterClusterCondition {
	if o == nil || o.Conditions == nil {
		var ret []ClusterClusterCondition
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterClusterStatus) GetConditionsOk() (*[]ClusterClusterCondition, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *ClusterClusterStatus) HasConditions() bool {
	if o != nil && o.Conditions != nil {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []ClusterClusterCondition and assigns it to the Conditions field.
func (o *ClusterClusterStatus) SetConditions(v []ClusterClusterCondition) {
	o.Conditions = &v
}

// GetCurrentTime returns the CurrentTime field value if set, zero value otherwise.
func (o *ClusterClusterStatus) GetCurrentTime() time.Time {
	if o == nil || o.CurrentTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CurrentTime
}

// GetCurrentTimeOk returns a tuple with the CurrentTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterClusterStatus) GetCurrentTimeOk() (*time.Time, bool) {
	if o == nil || o.CurrentTime == nil {
		return nil, false
	}
	return o.CurrentTime, true
}

// HasCurrentTime returns a boolean if a field has been set.
func (o *ClusterClusterStatus) HasCurrentTime() bool {
	if o != nil && o.CurrentTime != nil {
		return true
	}

	return false
}

// SetCurrentTime gets a reference to the given time.Time and assigns it to the CurrentTime field.
func (o *ClusterClusterStatus) SetCurrentTime(v time.Time) {
	o.CurrentTime = &v
}

// GetLastLeaderTransitionTime returns the LastLeaderTransitionTime field value if set, zero value otherwise.
func (o *ClusterClusterStatus) GetLastLeaderTransitionTime() time.Time {
	if o == nil || o.LastLeaderTransitionTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastLeaderTransitionTime
}

// GetLastLeaderTransitionTimeOk returns a tuple with the LastLeaderTransitionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterClusterStatus) GetLastLeaderTransitionTimeOk() (*time.Time, bool) {
	if o == nil || o.LastLeaderTransitionTime == nil {
		return nil, false
	}
	return o.LastLeaderTransitionTime, true
}

// HasLastLeaderTransitionTime returns a boolean if a field has been set.
func (o *ClusterClusterStatus) HasLastLeaderTransitionTime() bool {
	if o != nil && o.LastLeaderTransitionTime != nil {
		return true
	}

	return false
}

// SetLastLeaderTransitionTime gets a reference to the given time.Time and assigns it to the LastLeaderTransitionTime field.
func (o *ClusterClusterStatus) SetLastLeaderTransitionTime(v time.Time) {
	o.LastLeaderTransitionTime = &v
}

// GetLeader returns the Leader field value if set, zero value otherwise.
func (o *ClusterClusterStatus) GetLeader() string {
	if o == nil || o.Leader == nil {
		var ret string
		return ret
	}
	return *o.Leader
}

// GetLeaderOk returns a tuple with the Leader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterClusterStatus) GetLeaderOk() (*string, bool) {
	if o == nil || o.Leader == nil {
		return nil, false
	}
	return o.Leader, true
}

// HasLeader returns a boolean if a field has been set.
func (o *ClusterClusterStatus) HasLeader() bool {
	if o != nil && o.Leader != nil {
		return true
	}

	return false
}

// SetLeader gets a reference to the given string and assigns it to the Leader field.
func (o *ClusterClusterStatus) SetLeader(v string) {
	o.Leader = &v
}

// GetQuorumStatus returns the QuorumStatus field value if set, zero value otherwise.
func (o *ClusterClusterStatus) GetQuorumStatus() ClusterQuorumStatus {
	if o == nil || o.QuorumStatus == nil {
		var ret ClusterQuorumStatus
		return ret
	}
	return *o.QuorumStatus
}

// GetQuorumStatusOk returns a tuple with the QuorumStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterClusterStatus) GetQuorumStatusOk() (*ClusterQuorumStatus, bool) {
	if o == nil || o.QuorumStatus == nil {
		return nil, false
	}
	return o.QuorumStatus, true
}

// HasQuorumStatus returns a boolean if a field has been set.
func (o *ClusterClusterStatus) HasQuorumStatus() bool {
	if o != nil && o.QuorumStatus != nil {
		return true
	}

	return false
}

// SetQuorumStatus gets a reference to the given ClusterQuorumStatus and assigns it to the QuorumStatus field.
func (o *ClusterClusterStatus) SetQuorumStatus(v ClusterQuorumStatus) {
	o.QuorumStatus = &v
}

// GetRecoveryKeysDownloaded returns the RecoveryKeysDownloaded field value if set, zero value otherwise.
func (o *ClusterClusterStatus) GetRecoveryKeysDownloaded() bool {
	if o == nil || o.RecoveryKeysDownloaded == nil {
		var ret bool
		return ret
	}
	return *o.RecoveryKeysDownloaded
}

// GetRecoveryKeysDownloadedOk returns a tuple with the RecoveryKeysDownloaded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterClusterStatus) GetRecoveryKeysDownloadedOk() (*bool, bool) {
	if o == nil || o.RecoveryKeysDownloaded == nil {
		return nil, false
	}
	return o.RecoveryKeysDownloaded, true
}

// HasRecoveryKeysDownloaded returns a boolean if a field has been set.
func (o *ClusterClusterStatus) HasRecoveryKeysDownloaded() bool {
	if o != nil && o.RecoveryKeysDownloaded != nil {
		return true
	}

	return false
}

// SetRecoveryKeysDownloaded gets a reference to the given bool and assigns it to the RecoveryKeysDownloaded field.
func (o *ClusterClusterStatus) SetRecoveryKeysDownloaded(v bool) {
	o.RecoveryKeysDownloaded = &v
}

func (o ClusterClusterStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthBootstrapped != nil {
		toSerialize["auth-bootstrapped"] = o.AuthBootstrapped
	}
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}
	if o.CurrentTime != nil {
		toSerialize["current-time"] = o.CurrentTime
	}
	if o.LastLeaderTransitionTime != nil {
		toSerialize["last-leader-transition-time"] = o.LastLeaderTransitionTime
	}
	if o.Leader != nil {
		toSerialize["leader"] = o.Leader
	}
	if o.QuorumStatus != nil {
		toSerialize["quorum-status"] = o.QuorumStatus
	}
	if o.RecoveryKeysDownloaded != nil {
		toSerialize["recovery-keys-downloaded"] = o.RecoveryKeysDownloaded
	}
	return json.Marshal(toSerialize)
}

type NullableClusterClusterStatus struct {
	value *ClusterClusterStatus
	isSet bool
}

func (v NullableClusterClusterStatus) Get() *ClusterClusterStatus {
	return v.value
}

func (v *NullableClusterClusterStatus) Set(val *ClusterClusterStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterClusterStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterClusterStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterClusterStatus(val *ClusterClusterStatus) *NullableClusterClusterStatus {
	return &NullableClusterClusterStatus{value: val, isSet: true}
}

func (v NullableClusterClusterStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterClusterStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


