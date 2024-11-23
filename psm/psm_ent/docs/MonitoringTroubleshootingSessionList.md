# MonitoringTroubleshootingSessionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Items** | Pointer to [**[]MonitoringTroubleshootingSession**](MonitoringTroubleshootingSession.md) | List of TroubleshootingSession objects. | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**ListMeta** | Pointer to [**ApiListMeta**](apiListMeta.md) |  | [optional] 

## Methods

### NewMonitoringTroubleshootingSessionList

`func NewMonitoringTroubleshootingSessionList() *MonitoringTroubleshootingSessionList`

NewMonitoringTroubleshootingSessionList instantiates a new MonitoringTroubleshootingSessionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringTroubleshootingSessionListWithDefaults

`func NewMonitoringTroubleshootingSessionListWithDefaults() *MonitoringTroubleshootingSessionList`

NewMonitoringTroubleshootingSessionListWithDefaults instantiates a new MonitoringTroubleshootingSessionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *MonitoringTroubleshootingSessionList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *MonitoringTroubleshootingSessionList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *MonitoringTroubleshootingSessionList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *MonitoringTroubleshootingSessionList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *MonitoringTroubleshootingSessionList) GetItems() []MonitoringTroubleshootingSession`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *MonitoringTroubleshootingSessionList) GetItemsOk() (*[]MonitoringTroubleshootingSession, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *MonitoringTroubleshootingSessionList) SetItems(v []MonitoringTroubleshootingSession)`

SetItems sets Items field to given value.

### HasItems

`func (o *MonitoringTroubleshootingSessionList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetKind

`func (o *MonitoringTroubleshootingSessionList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *MonitoringTroubleshootingSessionList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *MonitoringTroubleshootingSessionList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *MonitoringTroubleshootingSessionList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetListMeta

`func (o *MonitoringTroubleshootingSessionList) GetListMeta() ApiListMeta`

GetListMeta returns the ListMeta field if non-nil, zero value otherwise.

### GetListMetaOk

`func (o *MonitoringTroubleshootingSessionList) GetListMetaOk() (*ApiListMeta, bool)`

GetListMetaOk returns a tuple with the ListMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListMeta

`func (o *MonitoringTroubleshootingSessionList) SetListMeta(v ApiListMeta)`

SetListMeta sets ListMeta field to given value.

### HasListMeta

`func (o *MonitoringTroubleshootingSessionList) HasListMeta() bool`

HasListMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


