# PreferencesIdleTimeout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | Pointer to **string** | Time of inactivity after which user logout countdown warning pops up. Should be a valid time duration. | [optional] [default to "60m"]
**WarningTime** | Pointer to **string** | Warning duration before logout and after system idle time. Should be a valid time duration of at most 5m0s. | [optional] [default to "10s"]

## Methods

### NewPreferencesIdleTimeout

`func NewPreferencesIdleTimeout() *PreferencesIdleTimeout`

NewPreferencesIdleTimeout instantiates a new PreferencesIdleTimeout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreferencesIdleTimeoutWithDefaults

`func NewPreferencesIdleTimeoutWithDefaults() *PreferencesIdleTimeout`

NewPreferencesIdleTimeoutWithDefaults instantiates a new PreferencesIdleTimeout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *PreferencesIdleTimeout) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *PreferencesIdleTimeout) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *PreferencesIdleTimeout) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *PreferencesIdleTimeout) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetWarningTime

`func (o *PreferencesIdleTimeout) GetWarningTime() string`

GetWarningTime returns the WarningTime field if non-nil, zero value otherwise.

### GetWarningTimeOk

`func (o *PreferencesIdleTimeout) GetWarningTimeOk() (*string, bool)`

GetWarningTimeOk returns a tuple with the WarningTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningTime

`func (o *PreferencesIdleTimeout) SetWarningTime(v string)`

SetWarningTime sets WarningTime field to given value.

### HasWarningTime

`func (o *PreferencesIdleTimeout) HasWarningTime() bool`

HasWarningTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


