# MonitoringFwlogPolicyList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Items** | Pointer to [**[]MonitoringFwlogPolicy**](MonitoringFwlogPolicy.md) | List of FwlogPolicy objects. | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**ListMeta** | Pointer to [**ApiListMeta**](apiListMeta.md) |  | [optional] 

## Methods

### NewMonitoringFwlogPolicyList

`func NewMonitoringFwlogPolicyList() *MonitoringFwlogPolicyList`

NewMonitoringFwlogPolicyList instantiates a new MonitoringFwlogPolicyList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringFwlogPolicyListWithDefaults

`func NewMonitoringFwlogPolicyListWithDefaults() *MonitoringFwlogPolicyList`

NewMonitoringFwlogPolicyListWithDefaults instantiates a new MonitoringFwlogPolicyList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *MonitoringFwlogPolicyList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *MonitoringFwlogPolicyList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *MonitoringFwlogPolicyList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *MonitoringFwlogPolicyList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *MonitoringFwlogPolicyList) GetItems() []MonitoringFwlogPolicy`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *MonitoringFwlogPolicyList) GetItemsOk() (*[]MonitoringFwlogPolicy, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *MonitoringFwlogPolicyList) SetItems(v []MonitoringFwlogPolicy)`

SetItems sets Items field to given value.

### HasItems

`func (o *MonitoringFwlogPolicyList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetKind

`func (o *MonitoringFwlogPolicyList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *MonitoringFwlogPolicyList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *MonitoringFwlogPolicyList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *MonitoringFwlogPolicyList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetListMeta

`func (o *MonitoringFwlogPolicyList) GetListMeta() ApiListMeta`

GetListMeta returns the ListMeta field if non-nil, zero value otherwise.

### GetListMetaOk

`func (o *MonitoringFwlogPolicyList) GetListMetaOk() (*ApiListMeta, bool)`

GetListMetaOk returns a tuple with the ListMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListMeta

`func (o *MonitoringFwlogPolicyList) SetListMeta(v ApiListMeta)`

SetListMeta sets ListMeta field to given value.

### HasListMeta

`func (o *MonitoringFwlogPolicyList) HasListMeta() bool`

HasListMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

