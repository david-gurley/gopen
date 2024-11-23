# NetworkHealthCheckSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeclareHealthyCount** | Pointer to **int64** | # of successful probes before we declare the backend back up. | [optional] 
**Interval** | Pointer to **int64** | Health check interval. | [optional] 
**MaxTimeouts** | Pointer to **int64** | timeout for declaring backend down. | [optional] 
**ProbePortOrUrl** | Pointer to **string** | probe URL. | [optional] 
**ProbesPerInterval** | Pointer to **int64** | # of probes per interval. | [optional] 

## Methods

### NewNetworkHealthCheckSpec

`func NewNetworkHealthCheckSpec() *NetworkHealthCheckSpec`

NewNetworkHealthCheckSpec instantiates a new NetworkHealthCheckSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkHealthCheckSpecWithDefaults

`func NewNetworkHealthCheckSpecWithDefaults() *NetworkHealthCheckSpec`

NewNetworkHealthCheckSpecWithDefaults instantiates a new NetworkHealthCheckSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeclareHealthyCount

`func (o *NetworkHealthCheckSpec) GetDeclareHealthyCount() int64`

GetDeclareHealthyCount returns the DeclareHealthyCount field if non-nil, zero value otherwise.

### GetDeclareHealthyCountOk

`func (o *NetworkHealthCheckSpec) GetDeclareHealthyCountOk() (*int64, bool)`

GetDeclareHealthyCountOk returns a tuple with the DeclareHealthyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclareHealthyCount

`func (o *NetworkHealthCheckSpec) SetDeclareHealthyCount(v int64)`

SetDeclareHealthyCount sets DeclareHealthyCount field to given value.

### HasDeclareHealthyCount

`func (o *NetworkHealthCheckSpec) HasDeclareHealthyCount() bool`

HasDeclareHealthyCount returns a boolean if a field has been set.

### GetInterval

`func (o *NetworkHealthCheckSpec) GetInterval() int64`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *NetworkHealthCheckSpec) GetIntervalOk() (*int64, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *NetworkHealthCheckSpec) SetInterval(v int64)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *NetworkHealthCheckSpec) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetMaxTimeouts

`func (o *NetworkHealthCheckSpec) GetMaxTimeouts() int64`

GetMaxTimeouts returns the MaxTimeouts field if non-nil, zero value otherwise.

### GetMaxTimeoutsOk

`func (o *NetworkHealthCheckSpec) GetMaxTimeoutsOk() (*int64, bool)`

GetMaxTimeoutsOk returns a tuple with the MaxTimeouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTimeouts

`func (o *NetworkHealthCheckSpec) SetMaxTimeouts(v int64)`

SetMaxTimeouts sets MaxTimeouts field to given value.

### HasMaxTimeouts

`func (o *NetworkHealthCheckSpec) HasMaxTimeouts() bool`

HasMaxTimeouts returns a boolean if a field has been set.

### GetProbePortOrUrl

`func (o *NetworkHealthCheckSpec) GetProbePortOrUrl() string`

GetProbePortOrUrl returns the ProbePortOrUrl field if non-nil, zero value otherwise.

### GetProbePortOrUrlOk

`func (o *NetworkHealthCheckSpec) GetProbePortOrUrlOk() (*string, bool)`

GetProbePortOrUrlOk returns a tuple with the ProbePortOrUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbePortOrUrl

`func (o *NetworkHealthCheckSpec) SetProbePortOrUrl(v string)`

SetProbePortOrUrl sets ProbePortOrUrl field to given value.

### HasProbePortOrUrl

`func (o *NetworkHealthCheckSpec) HasProbePortOrUrl() bool`

HasProbePortOrUrl returns a boolean if a field has been set.

### GetProbesPerInterval

`func (o *NetworkHealthCheckSpec) GetProbesPerInterval() int64`

GetProbesPerInterval returns the ProbesPerInterval field if non-nil, zero value otherwise.

### GetProbesPerIntervalOk

`func (o *NetworkHealthCheckSpec) GetProbesPerIntervalOk() (*int64, bool)`

GetProbesPerIntervalOk returns a tuple with the ProbesPerInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbesPerInterval

`func (o *NetworkHealthCheckSpec) SetProbesPerInterval(v int64)`

SetProbesPerInterval sets ProbesPerInterval field to given value.

### HasProbesPerInterval

`func (o *NetworkHealthCheckSpec) HasProbesPerInterval() bool`

HasProbesPerInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


