# MonitoringAlertList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Items** | Pointer to [**[]MonitoringAlert**](MonitoringAlert.md) | List of Alert objects. | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**ListMeta** | Pointer to [**ApiListMeta**](apiListMeta.md) |  | [optional] 

## Methods

### NewMonitoringAlertList

`func NewMonitoringAlertList() *MonitoringAlertList`

NewMonitoringAlertList instantiates a new MonitoringAlertList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringAlertListWithDefaults

`func NewMonitoringAlertListWithDefaults() *MonitoringAlertList`

NewMonitoringAlertListWithDefaults instantiates a new MonitoringAlertList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *MonitoringAlertList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *MonitoringAlertList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *MonitoringAlertList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *MonitoringAlertList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *MonitoringAlertList) GetItems() []MonitoringAlert`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *MonitoringAlertList) GetItemsOk() (*[]MonitoringAlert, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *MonitoringAlertList) SetItems(v []MonitoringAlert)`

SetItems sets Items field to given value.

### HasItems

`func (o *MonitoringAlertList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetKind

`func (o *MonitoringAlertList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *MonitoringAlertList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *MonitoringAlertList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *MonitoringAlertList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetListMeta

`func (o *MonitoringAlertList) GetListMeta() ApiListMeta`

GetListMeta returns the ListMeta field if non-nil, zero value otherwise.

### GetListMetaOk

`func (o *MonitoringAlertList) GetListMetaOk() (*ApiListMeta, bool)`

GetListMetaOk returns a tuple with the ListMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListMeta

`func (o *MonitoringAlertList) SetListMeta(v ApiListMeta)`

SetListMeta sets ListMeta field to given value.

### HasListMeta

`func (o *MonitoringAlertList) HasListMeta() bool`

HasListMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


