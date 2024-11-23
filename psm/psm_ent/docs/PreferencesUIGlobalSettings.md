# PreferencesUIGlobalSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**ApiObjectMeta**](apiObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**PreferencesUIGlobalSettingsSpec**](preferencesUIGlobalSettingsSpec.md) |  | [optional] 
**Status** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPreferencesUIGlobalSettings

`func NewPreferencesUIGlobalSettings() *PreferencesUIGlobalSettings`

NewPreferencesUIGlobalSettings instantiates a new PreferencesUIGlobalSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreferencesUIGlobalSettingsWithDefaults

`func NewPreferencesUIGlobalSettingsWithDefaults() *PreferencesUIGlobalSettings`

NewPreferencesUIGlobalSettingsWithDefaults instantiates a new PreferencesUIGlobalSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *PreferencesUIGlobalSettings) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *PreferencesUIGlobalSettings) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *PreferencesUIGlobalSettings) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *PreferencesUIGlobalSettings) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *PreferencesUIGlobalSettings) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PreferencesUIGlobalSettings) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PreferencesUIGlobalSettings) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *PreferencesUIGlobalSettings) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMeta

`func (o *PreferencesUIGlobalSettings) GetMeta() ApiObjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *PreferencesUIGlobalSettings) GetMetaOk() (*ApiObjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *PreferencesUIGlobalSettings) SetMeta(v ApiObjectMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *PreferencesUIGlobalSettings) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSpec

`func (o *PreferencesUIGlobalSettings) GetSpec() PreferencesUIGlobalSettingsSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *PreferencesUIGlobalSettings) GetSpecOk() (*PreferencesUIGlobalSettingsSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *PreferencesUIGlobalSettings) SetSpec(v PreferencesUIGlobalSettingsSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *PreferencesUIGlobalSettings) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *PreferencesUIGlobalSettings) GetStatus() map[string]interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PreferencesUIGlobalSettings) GetStatusOk() (*map[string]interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PreferencesUIGlobalSettings) SetStatus(v map[string]interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PreferencesUIGlobalSettings) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


