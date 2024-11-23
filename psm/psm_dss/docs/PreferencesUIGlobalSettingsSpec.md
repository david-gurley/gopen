# PreferencesUIGlobalSettingsSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdleTimeout** | Pointer to [**PreferencesIdleTimeout**](preferencesIdleTimeout.md) |  | [optional] 
**StyleOptions** | Pointer to **string** | Can contain any UI style preferences. Provide typing through UI code. | [optional] 

## Methods

### NewPreferencesUIGlobalSettingsSpec

`func NewPreferencesUIGlobalSettingsSpec() *PreferencesUIGlobalSettingsSpec`

NewPreferencesUIGlobalSettingsSpec instantiates a new PreferencesUIGlobalSettingsSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreferencesUIGlobalSettingsSpecWithDefaults

`func NewPreferencesUIGlobalSettingsSpecWithDefaults() *PreferencesUIGlobalSettingsSpec`

NewPreferencesUIGlobalSettingsSpecWithDefaults instantiates a new PreferencesUIGlobalSettingsSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdleTimeout

`func (o *PreferencesUIGlobalSettingsSpec) GetIdleTimeout() PreferencesIdleTimeout`

GetIdleTimeout returns the IdleTimeout field if non-nil, zero value otherwise.

### GetIdleTimeoutOk

`func (o *PreferencesUIGlobalSettingsSpec) GetIdleTimeoutOk() (*PreferencesIdleTimeout, bool)`

GetIdleTimeoutOk returns a tuple with the IdleTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleTimeout

`func (o *PreferencesUIGlobalSettingsSpec) SetIdleTimeout(v PreferencesIdleTimeout)`

SetIdleTimeout sets IdleTimeout field to given value.

### HasIdleTimeout

`func (o *PreferencesUIGlobalSettingsSpec) HasIdleTimeout() bool`

HasIdleTimeout returns a boolean if a field has been set.

### GetStyleOptions

`func (o *PreferencesUIGlobalSettingsSpec) GetStyleOptions() string`

GetStyleOptions returns the StyleOptions field if non-nil, zero value otherwise.

### GetStyleOptionsOk

`func (o *PreferencesUIGlobalSettingsSpec) GetStyleOptionsOk() (*string, bool)`

GetStyleOptionsOk returns a tuple with the StyleOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyleOptions

`func (o *PreferencesUIGlobalSettingsSpec) SetStyleOptions(v string)`

SetStyleOptions sets StyleOptions field to given value.

### HasStyleOptions

`func (o *PreferencesUIGlobalSettingsSpec) HasStyleOptions() bool`

HasStyleOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


