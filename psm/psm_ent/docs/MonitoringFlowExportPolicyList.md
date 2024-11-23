# MonitoringFlowExportPolicyList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Items** | Pointer to [**[]MonitoringFlowExportPolicy**](MonitoringFlowExportPolicy.md) | List of FlowExportPolicy objects. | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**ListMeta** | Pointer to [**ApiListMeta**](apiListMeta.md) |  | [optional] 

## Methods

### NewMonitoringFlowExportPolicyList

`func NewMonitoringFlowExportPolicyList() *MonitoringFlowExportPolicyList`

NewMonitoringFlowExportPolicyList instantiates a new MonitoringFlowExportPolicyList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringFlowExportPolicyListWithDefaults

`func NewMonitoringFlowExportPolicyListWithDefaults() *MonitoringFlowExportPolicyList`

NewMonitoringFlowExportPolicyListWithDefaults instantiates a new MonitoringFlowExportPolicyList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *MonitoringFlowExportPolicyList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *MonitoringFlowExportPolicyList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *MonitoringFlowExportPolicyList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *MonitoringFlowExportPolicyList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *MonitoringFlowExportPolicyList) GetItems() []MonitoringFlowExportPolicy`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *MonitoringFlowExportPolicyList) GetItemsOk() (*[]MonitoringFlowExportPolicy, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *MonitoringFlowExportPolicyList) SetItems(v []MonitoringFlowExportPolicy)`

SetItems sets Items field to given value.

### HasItems

`func (o *MonitoringFlowExportPolicyList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetKind

`func (o *MonitoringFlowExportPolicyList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *MonitoringFlowExportPolicyList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *MonitoringFlowExportPolicyList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *MonitoringFlowExportPolicyList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetListMeta

`func (o *MonitoringFlowExportPolicyList) GetListMeta() ApiListMeta`

GetListMeta returns the ListMeta field if non-nil, zero value otherwise.

### GetListMetaOk

`func (o *MonitoringFlowExportPolicyList) GetListMetaOk() (*ApiListMeta, bool)`

GetListMetaOk returns a tuple with the ListMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListMeta

`func (o *MonitoringFlowExportPolicyList) SetListMeta(v ApiListMeta)`

SetListMeta sets ListMeta field to given value.

### HasListMeta

`func (o *MonitoringFlowExportPolicyList) HasListMeta() bool`

HasListMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


