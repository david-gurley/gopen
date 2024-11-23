# MonitoringPrivacyConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algo** | Pointer to **string** |  | [optional] [default to "des56"]
**Password** | Pointer to **string** | Password contains the privacy password. | [optional] 

## Methods

### NewMonitoringPrivacyConfig

`func NewMonitoringPrivacyConfig() *MonitoringPrivacyConfig`

NewMonitoringPrivacyConfig instantiates a new MonitoringPrivacyConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringPrivacyConfigWithDefaults

`func NewMonitoringPrivacyConfigWithDefaults() *MonitoringPrivacyConfig`

NewMonitoringPrivacyConfigWithDefaults instantiates a new MonitoringPrivacyConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgo

`func (o *MonitoringPrivacyConfig) GetAlgo() string`

GetAlgo returns the Algo field if non-nil, zero value otherwise.

### GetAlgoOk

`func (o *MonitoringPrivacyConfig) GetAlgoOk() (*string, bool)`

GetAlgoOk returns a tuple with the Algo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgo

`func (o *MonitoringPrivacyConfig) SetAlgo(v string)`

SetAlgo sets Algo field to given value.

### HasAlgo

`func (o *MonitoringPrivacyConfig) HasAlgo() bool`

HasAlgo returns a boolean if a field has been set.

### GetPassword

`func (o *MonitoringPrivacyConfig) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *MonitoringPrivacyConfig) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *MonitoringPrivacyConfig) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *MonitoringPrivacyConfig) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


