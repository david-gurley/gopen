# NetworkPolicerCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BurstSize** | Pointer to **string** | Burst size in number of packets/bytes as policer criteria. | [optional] 
**BytesPerSecond** | Pointer to **string** | Maximum permissible bytes per second before policer will start dropping traffic. Either BytesPerSecond/PacketsPerSecond can be specified. | [optional] 
**PacketsPerSecond** | Pointer to **string** | Maximum permissible packets per second before policer will start dropping traffic. Either BytesPerSecond/PacketsPerSecond can be specified. | [optional] 

## Methods

### NewNetworkPolicerCriteria

`func NewNetworkPolicerCriteria() *NetworkPolicerCriteria`

NewNetworkPolicerCriteria instantiates a new NetworkPolicerCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkPolicerCriteriaWithDefaults

`func NewNetworkPolicerCriteriaWithDefaults() *NetworkPolicerCriteria`

NewNetworkPolicerCriteriaWithDefaults instantiates a new NetworkPolicerCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBurstSize

`func (o *NetworkPolicerCriteria) GetBurstSize() string`

GetBurstSize returns the BurstSize field if non-nil, zero value otherwise.

### GetBurstSizeOk

`func (o *NetworkPolicerCriteria) GetBurstSizeOk() (*string, bool)`

GetBurstSizeOk returns a tuple with the BurstSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurstSize

`func (o *NetworkPolicerCriteria) SetBurstSize(v string)`

SetBurstSize sets BurstSize field to given value.

### HasBurstSize

`func (o *NetworkPolicerCriteria) HasBurstSize() bool`

HasBurstSize returns a boolean if a field has been set.

### GetBytesPerSecond

`func (o *NetworkPolicerCriteria) GetBytesPerSecond() string`

GetBytesPerSecond returns the BytesPerSecond field if non-nil, zero value otherwise.

### GetBytesPerSecondOk

`func (o *NetworkPolicerCriteria) GetBytesPerSecondOk() (*string, bool)`

GetBytesPerSecondOk returns a tuple with the BytesPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesPerSecond

`func (o *NetworkPolicerCriteria) SetBytesPerSecond(v string)`

SetBytesPerSecond sets BytesPerSecond field to given value.

### HasBytesPerSecond

`func (o *NetworkPolicerCriteria) HasBytesPerSecond() bool`

HasBytesPerSecond returns a boolean if a field has been set.

### GetPacketsPerSecond

`func (o *NetworkPolicerCriteria) GetPacketsPerSecond() string`

GetPacketsPerSecond returns the PacketsPerSecond field if non-nil, zero value otherwise.

### GetPacketsPerSecondOk

`func (o *NetworkPolicerCriteria) GetPacketsPerSecondOk() (*string, bool)`

GetPacketsPerSecondOk returns a tuple with the PacketsPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketsPerSecond

`func (o *NetworkPolicerCriteria) SetPacketsPerSecond(v string)`

SetPacketsPerSecond sets PacketsPerSecond field to given value.

### HasPacketsPerSecond

`func (o *NetworkPolicerCriteria) HasPacketsPerSecond() bool`

HasPacketsPerSecond returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


