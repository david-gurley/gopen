# FwlogFwLogQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to **[]string** | OR of actions to be matched. Only one action can be specified and can only be specified if either source IP or destination IP is present. | [optional] 
**BatchSize** | Pointer to **int32** | BatchSize is the number of results returned in one batch while scrolling. | [optional] [default to 25]
**CountOnly** | Pointer to **bool** | if set, returns only number of hits for the search query and not flow logs. Number of hits are greater than or equal to TotalCount value returned in list-meta. | [optional] 
**DestinationIps** | Pointer to **[]string** | OR of destination IPs to be matched. Only one destination IP is allowed. Should be a valid v4 or v6 IP address. | [optional] 
**DestinationPorts** | Pointer to **[]int64** | OR of destination ports to be matched. Only one port can be specified and if present, destination IP must also be specified. Value should be between 0 and 65535. | [optional] 
**EncryptionStatus** | Pointer to **string** | if set, search logs that match the specified encryption status. | [optional] [default to "all"]
**EndTime** | Pointer to **time.Time** | EndTime selects all logs with timestamp less than the EndTime, example 2018-09-18T00:12:00Z. | [optional] 
**MaxResults** | Pointer to **int32** | MaxResults is the max-count of search results Default value is 50 and valid range is 0..8192. Value should be between 0 and 8192. | [optional] [default to 50]
**Protocols** | Pointer to **[]string** | OR of protocols to be matched. Only one protocol can be specified and can only be specified if either source IP or destination IP is present. | [optional] 
**ReporterIds** | Pointer to **[]string** | OR of reporter names to be matched. Only one reporter ID can be specified. | [optional] 
**ScrollAction** | Pointer to **string** | ScrollAction specifies actions related to scroll if its duration needs to be extended or scroll needs to be deleted. | [optional] [default to "none"]
**ScrollExpiry** | Pointer to **string** | ScrollExpiry is time duration after which scroll id expires. Default is 5 min. A duration string is a sequence of decimal number and a unit suffix, such as \&quot;300ms\&quot; or \&quot;2h45m\&quot;. Valid time units are \&quot;ns\&quot;, \&quot;us\&quot; (or \&quot;µs\&quot;), \&quot;ms\&quot;, \&quot;s\&quot;, \&quot;m\&quot;, \&quot;h\&quot;. Subsequent requests with a scroll id resets the timer for expiry. Should be a valid time duration. | [optional] [default to "5m"]
**ScrollId** | Pointer to **string** | ScrollID to scroll through results fetched by an earlier query. | [optional] 
**SortOrder** | Pointer to **string** | SortOrder specifies time ordering of results. | [optional] [default to "descending"]
**SourceIps** | Pointer to **[]string** | OR of sources IPs to be matched. Only one source IP is allowed. Should be a valid v4 or v6 IP address. | [optional] 
**SourcePorts** | Pointer to **[]int64** | OR of source ports to be matched. Only one port can be specified and if present, source IP must also be specified. Value should be between 0 and 65535. | [optional] 
**StartTime** | Pointer to **time.Time** | StartTime selects all logs with timestamp greater than the StartTime, example 2018-10-18T00:12:00Z. | [optional] 
**Tenants** | Pointer to **[]string** | OR of tenants within the scope of which search needs to be performed. If not specified, it will be set to tenant of the logged in user. Also users in non default tenant can search fwlogs in their tenant scope only. | [optional] 

## Methods

### NewFwlogFwLogQuery

`func NewFwlogFwLogQuery() *FwlogFwLogQuery`

NewFwlogFwLogQuery instantiates a new FwlogFwLogQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFwlogFwLogQueryWithDefaults

`func NewFwlogFwLogQueryWithDefaults() *FwlogFwLogQuery`

NewFwlogFwLogQueryWithDefaults instantiates a new FwlogFwLogQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *FwlogFwLogQuery) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *FwlogFwLogQuery) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *FwlogFwLogQuery) SetActions(v []string)`

SetActions sets Actions field to given value.

### HasActions

`func (o *FwlogFwLogQuery) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetBatchSize

`func (o *FwlogFwLogQuery) GetBatchSize() int32`

GetBatchSize returns the BatchSize field if non-nil, zero value otherwise.

### GetBatchSizeOk

`func (o *FwlogFwLogQuery) GetBatchSizeOk() (*int32, bool)`

GetBatchSizeOk returns a tuple with the BatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchSize

`func (o *FwlogFwLogQuery) SetBatchSize(v int32)`

SetBatchSize sets BatchSize field to given value.

### HasBatchSize

`func (o *FwlogFwLogQuery) HasBatchSize() bool`

HasBatchSize returns a boolean if a field has been set.

### GetCountOnly

`func (o *FwlogFwLogQuery) GetCountOnly() bool`

GetCountOnly returns the CountOnly field if non-nil, zero value otherwise.

### GetCountOnlyOk

`func (o *FwlogFwLogQuery) GetCountOnlyOk() (*bool, bool)`

GetCountOnlyOk returns a tuple with the CountOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountOnly

`func (o *FwlogFwLogQuery) SetCountOnly(v bool)`

SetCountOnly sets CountOnly field to given value.

### HasCountOnly

`func (o *FwlogFwLogQuery) HasCountOnly() bool`

HasCountOnly returns a boolean if a field has been set.

### GetDestinationIps

`func (o *FwlogFwLogQuery) GetDestinationIps() []string`

GetDestinationIps returns the DestinationIps field if non-nil, zero value otherwise.

### GetDestinationIpsOk

`func (o *FwlogFwLogQuery) GetDestinationIpsOk() (*[]string, bool)`

GetDestinationIpsOk returns a tuple with the DestinationIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationIps

`func (o *FwlogFwLogQuery) SetDestinationIps(v []string)`

SetDestinationIps sets DestinationIps field to given value.

### HasDestinationIps

`func (o *FwlogFwLogQuery) HasDestinationIps() bool`

HasDestinationIps returns a boolean if a field has been set.

### GetDestinationPorts

`func (o *FwlogFwLogQuery) GetDestinationPorts() []int64`

GetDestinationPorts returns the DestinationPorts field if non-nil, zero value otherwise.

### GetDestinationPortsOk

`func (o *FwlogFwLogQuery) GetDestinationPortsOk() (*[]int64, bool)`

GetDestinationPortsOk returns a tuple with the DestinationPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPorts

`func (o *FwlogFwLogQuery) SetDestinationPorts(v []int64)`

SetDestinationPorts sets DestinationPorts field to given value.

### HasDestinationPorts

`func (o *FwlogFwLogQuery) HasDestinationPorts() bool`

HasDestinationPorts returns a boolean if a field has been set.

### GetEncryptionStatus

`func (o *FwlogFwLogQuery) GetEncryptionStatus() string`

GetEncryptionStatus returns the EncryptionStatus field if non-nil, zero value otherwise.

### GetEncryptionStatusOk

`func (o *FwlogFwLogQuery) GetEncryptionStatusOk() (*string, bool)`

GetEncryptionStatusOk returns a tuple with the EncryptionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionStatus

`func (o *FwlogFwLogQuery) SetEncryptionStatus(v string)`

SetEncryptionStatus sets EncryptionStatus field to given value.

### HasEncryptionStatus

`func (o *FwlogFwLogQuery) HasEncryptionStatus() bool`

HasEncryptionStatus returns a boolean if a field has been set.

### GetEndTime

`func (o *FwlogFwLogQuery) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *FwlogFwLogQuery) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *FwlogFwLogQuery) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *FwlogFwLogQuery) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetMaxResults

`func (o *FwlogFwLogQuery) GetMaxResults() int32`

GetMaxResults returns the MaxResults field if non-nil, zero value otherwise.

### GetMaxResultsOk

`func (o *FwlogFwLogQuery) GetMaxResultsOk() (*int32, bool)`

GetMaxResultsOk returns a tuple with the MaxResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResults

`func (o *FwlogFwLogQuery) SetMaxResults(v int32)`

SetMaxResults sets MaxResults field to given value.

### HasMaxResults

`func (o *FwlogFwLogQuery) HasMaxResults() bool`

HasMaxResults returns a boolean if a field has been set.

### GetProtocols

`func (o *FwlogFwLogQuery) GetProtocols() []string`

GetProtocols returns the Protocols field if non-nil, zero value otherwise.

### GetProtocolsOk

`func (o *FwlogFwLogQuery) GetProtocolsOk() (*[]string, bool)`

GetProtocolsOk returns a tuple with the Protocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocols

`func (o *FwlogFwLogQuery) SetProtocols(v []string)`

SetProtocols sets Protocols field to given value.

### HasProtocols

`func (o *FwlogFwLogQuery) HasProtocols() bool`

HasProtocols returns a boolean if a field has been set.

### GetReporterIds

`func (o *FwlogFwLogQuery) GetReporterIds() []string`

GetReporterIds returns the ReporterIds field if non-nil, zero value otherwise.

### GetReporterIdsOk

`func (o *FwlogFwLogQuery) GetReporterIdsOk() (*[]string, bool)`

GetReporterIdsOk returns a tuple with the ReporterIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReporterIds

`func (o *FwlogFwLogQuery) SetReporterIds(v []string)`

SetReporterIds sets ReporterIds field to given value.

### HasReporterIds

`func (o *FwlogFwLogQuery) HasReporterIds() bool`

HasReporterIds returns a boolean if a field has been set.

### GetScrollAction

`func (o *FwlogFwLogQuery) GetScrollAction() string`

GetScrollAction returns the ScrollAction field if non-nil, zero value otherwise.

### GetScrollActionOk

`func (o *FwlogFwLogQuery) GetScrollActionOk() (*string, bool)`

GetScrollActionOk returns a tuple with the ScrollAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrollAction

`func (o *FwlogFwLogQuery) SetScrollAction(v string)`

SetScrollAction sets ScrollAction field to given value.

### HasScrollAction

`func (o *FwlogFwLogQuery) HasScrollAction() bool`

HasScrollAction returns a boolean if a field has been set.

### GetScrollExpiry

`func (o *FwlogFwLogQuery) GetScrollExpiry() string`

GetScrollExpiry returns the ScrollExpiry field if non-nil, zero value otherwise.

### GetScrollExpiryOk

`func (o *FwlogFwLogQuery) GetScrollExpiryOk() (*string, bool)`

GetScrollExpiryOk returns a tuple with the ScrollExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrollExpiry

`func (o *FwlogFwLogQuery) SetScrollExpiry(v string)`

SetScrollExpiry sets ScrollExpiry field to given value.

### HasScrollExpiry

`func (o *FwlogFwLogQuery) HasScrollExpiry() bool`

HasScrollExpiry returns a boolean if a field has been set.

### GetScrollId

`func (o *FwlogFwLogQuery) GetScrollId() string`

GetScrollId returns the ScrollId field if non-nil, zero value otherwise.

### GetScrollIdOk

`func (o *FwlogFwLogQuery) GetScrollIdOk() (*string, bool)`

GetScrollIdOk returns a tuple with the ScrollId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrollId

`func (o *FwlogFwLogQuery) SetScrollId(v string)`

SetScrollId sets ScrollId field to given value.

### HasScrollId

`func (o *FwlogFwLogQuery) HasScrollId() bool`

HasScrollId returns a boolean if a field has been set.

### GetSortOrder

`func (o *FwlogFwLogQuery) GetSortOrder() string`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *FwlogFwLogQuery) GetSortOrderOk() (*string, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *FwlogFwLogQuery) SetSortOrder(v string)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *FwlogFwLogQuery) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetSourceIps

`func (o *FwlogFwLogQuery) GetSourceIps() []string`

GetSourceIps returns the SourceIps field if non-nil, zero value otherwise.

### GetSourceIpsOk

`func (o *FwlogFwLogQuery) GetSourceIpsOk() (*[]string, bool)`

GetSourceIpsOk returns a tuple with the SourceIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIps

`func (o *FwlogFwLogQuery) SetSourceIps(v []string)`

SetSourceIps sets SourceIps field to given value.

### HasSourceIps

`func (o *FwlogFwLogQuery) HasSourceIps() bool`

HasSourceIps returns a boolean if a field has been set.

### GetSourcePorts

`func (o *FwlogFwLogQuery) GetSourcePorts() []int64`

GetSourcePorts returns the SourcePorts field if non-nil, zero value otherwise.

### GetSourcePortsOk

`func (o *FwlogFwLogQuery) GetSourcePortsOk() (*[]int64, bool)`

GetSourcePortsOk returns a tuple with the SourcePorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePorts

`func (o *FwlogFwLogQuery) SetSourcePorts(v []int64)`

SetSourcePorts sets SourcePorts field to given value.

### HasSourcePorts

`func (o *FwlogFwLogQuery) HasSourcePorts() bool`

HasSourcePorts returns a boolean if a field has been set.

### GetStartTime

`func (o *FwlogFwLogQuery) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *FwlogFwLogQuery) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *FwlogFwLogQuery) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *FwlogFwLogQuery) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetTenants

`func (o *FwlogFwLogQuery) GetTenants() []string`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *FwlogFwLogQuery) GetTenantsOk() (*[]string, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *FwlogFwLogQuery) SetTenants(v []string)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *FwlogFwLogQuery) HasTenants() bool`

HasTenants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


