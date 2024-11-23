# NetworkPauseSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RxPauseEnabled** | Pointer to **bool** | RX Pause enabled. | [optional] 
**TxPauseEnabled** | Pointer to **bool** | TX Pause enabled. | [optional] 
**Type** | Pointer to **string** | Pause type. | [optional] [default to "disable"]

## Methods

### NewNetworkPauseSpec

`func NewNetworkPauseSpec() *NetworkPauseSpec`

NewNetworkPauseSpec instantiates a new NetworkPauseSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkPauseSpecWithDefaults

`func NewNetworkPauseSpecWithDefaults() *NetworkPauseSpec`

NewNetworkPauseSpecWithDefaults instantiates a new NetworkPauseSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRxPauseEnabled

`func (o *NetworkPauseSpec) GetRxPauseEnabled() bool`

GetRxPauseEnabled returns the RxPauseEnabled field if non-nil, zero value otherwise.

### GetRxPauseEnabledOk

`func (o *NetworkPauseSpec) GetRxPauseEnabledOk() (*bool, bool)`

GetRxPauseEnabledOk returns a tuple with the RxPauseEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxPauseEnabled

`func (o *NetworkPauseSpec) SetRxPauseEnabled(v bool)`

SetRxPauseEnabled sets RxPauseEnabled field to given value.

### HasRxPauseEnabled

`func (o *NetworkPauseSpec) HasRxPauseEnabled() bool`

HasRxPauseEnabled returns a boolean if a field has been set.

### GetTxPauseEnabled

`func (o *NetworkPauseSpec) GetTxPauseEnabled() bool`

GetTxPauseEnabled returns the TxPauseEnabled field if non-nil, zero value otherwise.

### GetTxPauseEnabledOk

`func (o *NetworkPauseSpec) GetTxPauseEnabledOk() (*bool, bool)`

GetTxPauseEnabledOk returns a tuple with the TxPauseEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxPauseEnabled

`func (o *NetworkPauseSpec) SetTxPauseEnabled(v bool)`

SetTxPauseEnabled sets TxPauseEnabled field to given value.

### HasTxPauseEnabled

`func (o *NetworkPauseSpec) HasTxPauseEnabled() bool`

HasTxPauseEnabled returns a boolean if a field has been set.

### GetType

`func (o *NetworkPauseSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkPauseSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkPauseSpec) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NetworkPauseSpec) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


