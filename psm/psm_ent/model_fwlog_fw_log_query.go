/*
 * Fwlog API reference
 *
 * Service name  
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package psm_ent

import (
	"encoding/json"
	"time"
)

// FwlogFwLogQuery FwLogQuery allows selecting logs by all attributes All fields are ANDed together.
type FwlogFwLogQuery struct {
	// OR of actions to be matched. Only one action can be specified and can only be specified if either source IP or destination IP is present.
	Actions *[]string `json:"actions,omitempty"`
	// BatchSize is the number of results returned in one batch while scrolling.
	BatchSize *int32 `json:"batch-size,omitempty"`
	// if set, returns only number of hits for the search query and not flow logs. Number of hits are greater than or equal to TotalCount value returned in list-meta.
	CountOnly *bool `json:"count-only,omitempty"`
	// OR of destination IPs to be matched. Only one destination IP is allowed. Should be a valid v4 or v6 IP address.
	DestinationIps *[]string `json:"destination-ips,omitempty"`
	// OR of destination ports to be matched. Only one port can be specified and if present, destination IP must also be specified. Value should be between 0 and 65535.
	DestinationPorts *[]int64 `json:"destination-ports,omitempty"`
	// if set, search logs that match the specified encryption status.
	EncryptionStatus *string `json:"encryption-status,omitempty"`
	// EndTime selects all logs with timestamp less than the EndTime, example 2018-09-18T00:12:00Z.
	EndTime *time.Time `json:"end-time,omitempty"`
	// MaxResults is the max-count of search results Default value is 50 and valid range is 0..8192. Value should be between 0 and 8192.
	MaxResults *int32 `json:"max-results,omitempty"`
	// OR of protocols to be matched. Only one protocol can be specified and can only be specified if either source IP or destination IP is present.
	Protocols *[]string `json:"protocols,omitempty"`
	// OR of reporter names to be matched. Only one reporter ID can be specified.
	ReporterIds *[]string `json:"reporter-ids,omitempty"`
	// ScrollAction specifies actions related to scroll if its duration needs to be extended or scroll needs to be deleted.
	ScrollAction *string `json:"scroll-action,omitempty"`
	// ScrollExpiry is time duration after which scroll id expires. Default is 5 min. A duration string is a sequence of decimal number and a unit suffix, such as \"300ms\" or \"2h45m\". Valid time units are \"ns\", \"us\" (or \"µs\"), \"ms\", \"s\", \"m\", \"h\". Subsequent requests with a scroll id resets the timer for expiry. Should be a valid time duration.
	ScrollExpiry *string `json:"scroll-expiry,omitempty"`
	// ScrollID to scroll through results fetched by an earlier query.
	ScrollId *string `json:"scroll-id,omitempty"`
	// SortOrder specifies time ordering of results.
	SortOrder *string `json:"sort-order,omitempty"`
	// OR of sources IPs to be matched. Only one source IP is allowed. Should be a valid v4 or v6 IP address.
	SourceIps *[]string `json:"source-ips,omitempty"`
	// OR of source ports to be matched. Only one port can be specified and if present, source IP must also be specified. Value should be between 0 and 65535.
	SourcePorts *[]int64 `json:"source-ports,omitempty"`
	// StartTime selects all logs with timestamp greater than the StartTime, example 2018-10-18T00:12:00Z.
	StartTime *time.Time `json:"start-time,omitempty"`
	// OR of tenants within the scope of which search needs to be performed. If not specified, it will be set to tenant of the logged in user. Also users in non default tenant can search fwlogs in their tenant scope only.
	Tenants *[]string `json:"tenants,omitempty"`
}

// NewFwlogFwLogQuery instantiates a new FwlogFwLogQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFwlogFwLogQuery() *FwlogFwLogQuery {
	this := FwlogFwLogQuery{}
	var batchSize int32 = 25
	this.BatchSize = &batchSize
	var encryptionStatus string = "all"
	this.EncryptionStatus = &encryptionStatus
	var maxResults int32 = 50
	this.MaxResults = &maxResults
	var scrollAction string = "none"
	this.ScrollAction = &scrollAction
	var scrollExpiry string = "5m"
	this.ScrollExpiry = &scrollExpiry
	var sortOrder string = "descending"
	this.SortOrder = &sortOrder
	return &this
}

// NewFwlogFwLogQueryWithDefaults instantiates a new FwlogFwLogQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFwlogFwLogQueryWithDefaults() *FwlogFwLogQuery {
	this := FwlogFwLogQuery{}
	var batchSize int32 = 25
	this.BatchSize = &batchSize
	var encryptionStatus string = "all"
	this.EncryptionStatus = &encryptionStatus
	var maxResults int32 = 50
	this.MaxResults = &maxResults
	var scrollAction string = "none"
	this.ScrollAction = &scrollAction
	var scrollExpiry string = "5m"
	this.ScrollExpiry = &scrollExpiry
	var sortOrder string = "descending"
	this.SortOrder = &sortOrder
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *FwlogFwLogQuery) GetActions() []string {
	if o == nil || o.Actions == nil {
		var ret []string
		return ret
	}
	return *o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FwlogFwLogQuery) GetActionsOk() (*[]string, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *FwlogFwLogQuery) HasActions() bool {
	if o != nil && o.Actions != nil {
		return true
	}

	return false
}

// SetActions gets a reference to the given []string and assigns it to the Actions field.
func (o *FwlogFwLogQuery) SetActions(v []string) {
	o.Actions = &v
}

// GetBatchSize returns the BatchSize field value if set, zero value otherwise.
func (o *FwlogFwLogQuery) GetBatchSize() int32 {
	if o == nil || o.BatchSize == nil {
		var ret int32
		return ret
	}
	return *o.BatchSize
}

// GetBatchSizeOk returns a tuple with the BatchSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FwlogFwLogQuery) GetBatchSizeOk() (*int32, bool) {
	if o == nil || o.BatchSize == nil {
		return nil, false
	}
	return o.BatchSize, true
}

// HasBatchSize returns a boolean if a field has been set.
func (o *FwlogFwLogQuery) HasBatchSize() bool {
	if o != nil && o.BatchSize != nil {
		return true
	}

	return false
}

// SetBatchSize gets a reference to the given int32 and assigns it to the BatchSize field.
func (o *FwlogFwLogQuery) SetBatchSize(v int32) {
	o.BatchSize = &v
}

// GetCountOnly returns the CountOnly field value if set, zero value otherwise.
func (o *FwlogFwLogQuery) GetCountOnly() bool {
	if o == nil || o.CountOnly == nil {
		var ret bool
		return ret
	}
	return *o.CountOnly
}

// GetCountOnlyOk returns a tuple with the CountOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FwlogFwLogQuery) GetCountOnlyOk() (*bool, bool) {
	if o == nil || o.CountOnly == nil {
		return nil, false
	}
	return o.CountOnly, true
}

// HasCountOnly returns a boolean if a field has been set.
func (o *FwlogFwLogQuery) HasCountOnly() bool {
	if o != nil && o.CountOnly != nil {
		return true
	}

	return false
}

// SetCountOnly gets a reference to the given bool and assigns it to the CountOnly field.
func (o *FwlogFwLogQuery) SetCountOnly(v bool) {
	o.CountOnly = &v
}

// GetDestinationIps returns the DestinationIps field value if set, zero value otherwise.
func (o *FwlogFwLogQuery) GetDestinationIps() []string {
	if o == nil || o.DestinationIps == nil {
		var ret []string
		return ret
	}
	return *o.DestinationIps
}

// GetDestinationIpsOk returns a tuple with the DestinationIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FwlogFwLogQuery) GetDestinationIpsOk() (*[]string, bool) {
	if o == nil || o.DestinationIps == nil {
		return nil, false
	}
	return o.DestinationIps, true
}

// HasDestinationIps returns a boolean if a field has been set.
func (o *FwlogFwLogQuery) HasDestinationIps() bool {
	if o != nil && o.DestinationIps != nil {
		return true
	}

	return false
}

// SetDestinationIps gets a reference to the given []string and assigns it to the DestinationIps field.
func (o *FwlogFwLogQuery) SetDestinationIps(v []string) {
	o.DestinationIps = &v
}

// GetDestinationPorts returns the DestinationPorts field value if set, zero value otherwise.
func (o *FwlogFwLogQuery) GetDestinationPorts() []int64 {
	if o == nil || o.DestinationPorts == nil {
		var ret []int64
		return ret
	}
	return *o.DestinationPorts
}

// GetDestinationPortsOk returns a tuple with the DestinationPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FwlogFwLogQuery) GetDestinationPortsOk() (*[]int64, bool) {
	if o == nil || o.DestinationPorts == nil {
		return nil, false
	}
	return o.DestinationPorts, true
}

// HasDestinationPorts returns a boolean if a field has been set.
func (o *FwlogFwLogQuery) HasDestinationPorts() bool {
	if o != nil && o.DestinationPorts != nil {
		return true
	}

	return false
}

// SetDestinationPorts gets a reference to the given []int64 and assigns it to the DestinationPorts field.
func (o *FwlogFwLogQuery) SetDestinationPorts(v []int64) {
	o.DestinationPorts = &v
}

// GetEncryptionStatus returns the EncryptionStatus field value if set, zero value otherwise.
func (o *FwlogFwLogQuery) GetEncryptionStatus() string {
	if o == nil || o.EncryptionStatus == nil {
		var ret string
		return ret
	}
	return *o.EncryptionStatus
}

// GetEncryptionStatusOk returns a tuple with the EncryptionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FwlogFwLogQuery) GetEncryptionStatusOk() (*string, bool) {
	if o == nil || o.EncryptionStatus == nil {
		return nil, false
	}
	return o.EncryptionStatus, true
}

// HasEncryptionStatus returns a boolean if a field has been set.
func (o *FwlogFwLogQuery) HasEncryptionStatus() bool {
	if o != nil && o.EncryptionStatus != nil {
		return true
	}

	return false
}

// SetEncryptionStatus gets a reference to the given string and assigns it to the EncryptionStatus field.
func (o *FwlogFwLogQuery) SetEncryptionStatus(v string) {
	o.EncryptionStatus = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *FwlogFwLogQuery) GetEndTime() time.Time {
	if o == nil || o.EndTime == nil {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FwlogFwLogQuery) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || o.EndTime == nil {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *FwlogFwLogQuery) HasEndTime() bool {
	if o != nil && o.EndTime != nil {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *FwlogFwLogQuery) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetMaxResults returns the MaxResults field value if set, zero value otherwise.
func (o *FwlogFwLogQuery) GetMaxResults() int32 {
	if o == nil || o.MaxResults == nil {
		var ret int32
		return ret
	}
	return *o.MaxResults
}

// GetMaxResultsOk returns a tuple with the MaxResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FwlogFwLogQuery) GetMaxResultsOk() (*int32, bool) {
	if o == nil || o.MaxResults == nil {
		return nil, false
	}
	return o.MaxResults, true
}

// HasMaxResults returns a boolean if a field has been set.
func (o *FwlogFwLogQuery) HasMaxResults() bool {
	if o != nil && o.MaxResults != nil {
		return true
	}

	return false
}

// SetMaxResults gets a reference to the given int32 and assigns it to the MaxResults field.
func (o *FwlogFwLogQuery) SetMaxResults(v int32) {
	o.MaxResults = &v
}

// GetProtocols returns the Protocols field value if set, zero value otherwise.
func (o *FwlogFwLogQuery) GetProtocols() []string {
	if o == nil || o.Protocols == nil {
		var ret []string
		return ret
	}
	return *o.Protocols
}

// GetProtocolsOk returns a tuple with the Protocols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FwlogFwLogQuery) GetProtocolsOk() (*[]string, bool) {
	if o == nil || o.Protocols == nil {
		return nil, false
	}
	return o.Protocols, true
}

// HasProtocols returns a boolean if a field has been set.
func (o *FwlogFwLogQuery) HasProtocols() bool {
	if o != nil && o.Protocols != nil {
		return true
	}

	return false
}

// SetProtocols gets a reference to the given []string and assigns it to the Protocols field.
func (o *FwlogFwLogQuery) SetProtocols(v []string) {
	o.Protocols = &v
}

// GetReporterIds returns the ReporterIds field value if set, zero value otherwise.
func (o *FwlogFwLogQuery) GetReporterIds() []string {
	if o == nil || o.ReporterIds == nil {
		var ret []string
		return ret
	}
	return *o.ReporterIds
}

// GetReporterIdsOk returns a tuple with the ReporterIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FwlogFwLogQuery) GetReporterIdsOk() (*[]string, bool) {
	if o == nil || o.ReporterIds == nil {
		return nil, false
	}
	return o.ReporterIds, true
}

// HasReporterIds returns a boolean if a field has been set.
func (o *FwlogFwLogQuery) HasReporterIds() bool {
	if o != nil && o.ReporterIds != nil {
		return true
	}

	return false
}

// SetReporterIds gets a reference to the given []string and assigns it to the ReporterIds field.
func (o *FwlogFwLogQuery) SetReporterIds(v []string) {
	o.ReporterIds = &v
}

// GetScrollAction returns the ScrollAction field value if set, zero value otherwise.
func (o *FwlogFwLogQuery) GetScrollAction() string {
	if o == nil || o.ScrollAction == nil {
		var ret string
		return ret
	}
	return *o.ScrollAction
}

// GetScrollActionOk returns a tuple with the ScrollAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FwlogFwLogQuery) GetScrollActionOk() (*string, bool) {
	if o == nil || o.ScrollAction == nil {
		return nil, false
	}
	return o.ScrollAction, true
}

// HasScrollAction returns a boolean if a field has been set.
func (o *FwlogFwLogQuery) HasScrollAction() bool {
	if o != nil && o.ScrollAction != nil {
		return true
	}

	return false
}

// SetScrollAction gets a reference to the given string and assigns it to the ScrollAction field.
func (o *FwlogFwLogQuery) SetScrollAction(v string) {
	o.ScrollAction = &v
}

// GetScrollExpiry returns the ScrollExpiry field value if set, zero value otherwise.
func (o *FwlogFwLogQuery) GetScrollExpiry() string {
	if o == nil || o.ScrollExpiry == nil {
		var ret string
		return ret
	}
	return *o.ScrollExpiry
}

// GetScrollExpiryOk returns a tuple with the ScrollExpiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FwlogFwLogQuery) GetScrollExpiryOk() (*string, bool) {
	if o == nil || o.ScrollExpiry == nil {
		return nil, false
	}
	return o.ScrollExpiry, true
}

// HasScrollExpiry returns a boolean if a field has been set.
func (o *FwlogFwLogQuery) HasScrollExpiry() bool {
	if o != nil && o.ScrollExpiry != nil {
		return true
	}

	return false
}

// SetScrollExpiry gets a reference to the given string and assigns it to the ScrollExpiry field.
func (o *FwlogFwLogQuery) SetScrollExpiry(v string) {
	o.ScrollExpiry = &v
}

// GetScrollId returns the ScrollId field value if set, zero value otherwise.
func (o *FwlogFwLogQuery) GetScrollId() string {
	if o == nil || o.ScrollId == nil {
		var ret string
		return ret
	}
	return *o.ScrollId
}

// GetScrollIdOk returns a tuple with the ScrollId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FwlogFwLogQuery) GetScrollIdOk() (*string, bool) {
	if o == nil || o.ScrollId == nil {
		return nil, false
	}
	return o.ScrollId, true
}

// HasScrollId returns a boolean if a field has been set.
func (o *FwlogFwLogQuery) HasScrollId() bool {
	if o != nil && o.ScrollId != nil {
		return true
	}

	return false
}

// SetScrollId gets a reference to the given string and assigns it to the ScrollId field.
func (o *FwlogFwLogQuery) SetScrollId(v string) {
	o.ScrollId = &v
}

// GetSortOrder returns the SortOrder field value if set, zero value otherwise.
func (o *FwlogFwLogQuery) GetSortOrder() string {
	if o == nil || o.SortOrder == nil {
		var ret string
		return ret
	}
	return *o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FwlogFwLogQuery) GetSortOrderOk() (*string, bool) {
	if o == nil || o.SortOrder == nil {
		return nil, false
	}
	return o.SortOrder, true
}

// HasSortOrder returns a boolean if a field has been set.
func (o *FwlogFwLogQuery) HasSortOrder() bool {
	if o != nil && o.SortOrder != nil {
		return true
	}

	return false
}

// SetSortOrder gets a reference to the given string and assigns it to the SortOrder field.
func (o *FwlogFwLogQuery) SetSortOrder(v string) {
	o.SortOrder = &v
}

// GetSourceIps returns the SourceIps field value if set, zero value otherwise.
func (o *FwlogFwLogQuery) GetSourceIps() []string {
	if o == nil || o.SourceIps == nil {
		var ret []string
		return ret
	}
	return *o.SourceIps
}

// GetSourceIpsOk returns a tuple with the SourceIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FwlogFwLogQuery) GetSourceIpsOk() (*[]string, bool) {
	if o == nil || o.SourceIps == nil {
		return nil, false
	}
	return o.SourceIps, true
}

// HasSourceIps returns a boolean if a field has been set.
func (o *FwlogFwLogQuery) HasSourceIps() bool {
	if o != nil && o.SourceIps != nil {
		return true
	}

	return false
}

// SetSourceIps gets a reference to the given []string and assigns it to the SourceIps field.
func (o *FwlogFwLogQuery) SetSourceIps(v []string) {
	o.SourceIps = &v
}

// GetSourcePorts returns the SourcePorts field value if set, zero value otherwise.
func (o *FwlogFwLogQuery) GetSourcePorts() []int64 {
	if o == nil || o.SourcePorts == nil {
		var ret []int64
		return ret
	}
	return *o.SourcePorts
}

// GetSourcePortsOk returns a tuple with the SourcePorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FwlogFwLogQuery) GetSourcePortsOk() (*[]int64, bool) {
	if o == nil || o.SourcePorts == nil {
		return nil, false
	}
	return o.SourcePorts, true
}

// HasSourcePorts returns a boolean if a field has been set.
func (o *FwlogFwLogQuery) HasSourcePorts() bool {
	if o != nil && o.SourcePorts != nil {
		return true
	}

	return false
}

// SetSourcePorts gets a reference to the given []int64 and assigns it to the SourcePorts field.
func (o *FwlogFwLogQuery) SetSourcePorts(v []int64) {
	o.SourcePorts = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *FwlogFwLogQuery) GetStartTime() time.Time {
	if o == nil || o.StartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FwlogFwLogQuery) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *FwlogFwLogQuery) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *FwlogFwLogQuery) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetTenants returns the Tenants field value if set, zero value otherwise.
func (o *FwlogFwLogQuery) GetTenants() []string {
	if o == nil || o.Tenants == nil {
		var ret []string
		return ret
	}
	return *o.Tenants
}

// GetTenantsOk returns a tuple with the Tenants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FwlogFwLogQuery) GetTenantsOk() (*[]string, bool) {
	if o == nil || o.Tenants == nil {
		return nil, false
	}
	return o.Tenants, true
}

// HasTenants returns a boolean if a field has been set.
func (o *FwlogFwLogQuery) HasTenants() bool {
	if o != nil && o.Tenants != nil {
		return true
	}

	return false
}

// SetTenants gets a reference to the given []string and assigns it to the Tenants field.
func (o *FwlogFwLogQuery) SetTenants(v []string) {
	o.Tenants = &v
}

func (o FwlogFwLogQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Actions != nil {
		toSerialize["actions"] = o.Actions
	}
	if o.BatchSize != nil {
		toSerialize["batch-size"] = o.BatchSize
	}
	if o.CountOnly != nil {
		toSerialize["count-only"] = o.CountOnly
	}
	if o.DestinationIps != nil {
		toSerialize["destination-ips"] = o.DestinationIps
	}
	if o.DestinationPorts != nil {
		toSerialize["destination-ports"] = o.DestinationPorts
	}
	if o.EncryptionStatus != nil {
		toSerialize["encryption-status"] = o.EncryptionStatus
	}
	if o.EndTime != nil {
		toSerialize["end-time"] = o.EndTime
	}
	if o.MaxResults != nil {
		toSerialize["max-results"] = o.MaxResults
	}
	if o.Protocols != nil {
		toSerialize["protocols"] = o.Protocols
	}
	if o.ReporterIds != nil {
		toSerialize["reporter-ids"] = o.ReporterIds
	}
	if o.ScrollAction != nil {
		toSerialize["scroll-action"] = o.ScrollAction
	}
	if o.ScrollExpiry != nil {
		toSerialize["scroll-expiry"] = o.ScrollExpiry
	}
	if o.ScrollId != nil {
		toSerialize["scroll-id"] = o.ScrollId
	}
	if o.SortOrder != nil {
		toSerialize["sort-order"] = o.SortOrder
	}
	if o.SourceIps != nil {
		toSerialize["source-ips"] = o.SourceIps
	}
	if o.SourcePorts != nil {
		toSerialize["source-ports"] = o.SourcePorts
	}
	if o.StartTime != nil {
		toSerialize["start-time"] = o.StartTime
	}
	if o.Tenants != nil {
		toSerialize["tenants"] = o.Tenants
	}
	return json.Marshal(toSerialize)
}

type NullableFwlogFwLogQuery struct {
	value *FwlogFwLogQuery
	isSet bool
}

func (v NullableFwlogFwLogQuery) Get() *FwlogFwLogQuery {
	return v.value
}

func (v *NullableFwlogFwLogQuery) Set(val *FwlogFwLogQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableFwlogFwLogQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableFwlogFwLogQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFwlogFwLogQuery(val *FwlogFwLogQuery) *NullableFwlogFwLogQuery {
	return &NullableFwlogFwLogQuery{value: val, isSet: true}
}

func (v NullableFwlogFwLogQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFwlogFwLogQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

