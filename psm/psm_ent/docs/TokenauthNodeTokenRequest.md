# TokenauthNodeTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audience** | Pointer to **[]string** | Audience represents a list of nodes the token is valid for. \&quot;*\&quot; indicates all nodes. | [optional] 
**ValidityEnd** | Pointer to **time.Time** | ValidityEnd indicates the time at which the token becomes invalid. | [optional] 
**ValidityStart** | Pointer to **time.Time** | ValidityStart indicates the time at which the token becomes valid. | [optional] 

## Methods

### NewTokenauthNodeTokenRequest

`func NewTokenauthNodeTokenRequest() *TokenauthNodeTokenRequest`

NewTokenauthNodeTokenRequest instantiates a new TokenauthNodeTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenauthNodeTokenRequestWithDefaults

`func NewTokenauthNodeTokenRequestWithDefaults() *TokenauthNodeTokenRequest`

NewTokenauthNodeTokenRequestWithDefaults instantiates a new TokenauthNodeTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudience

`func (o *TokenauthNodeTokenRequest) GetAudience() []string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *TokenauthNodeTokenRequest) GetAudienceOk() (*[]string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *TokenauthNodeTokenRequest) SetAudience(v []string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *TokenauthNodeTokenRequest) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetValidityEnd

`func (o *TokenauthNodeTokenRequest) GetValidityEnd() time.Time`

GetValidityEnd returns the ValidityEnd field if non-nil, zero value otherwise.

### GetValidityEndOk

`func (o *TokenauthNodeTokenRequest) GetValidityEndOk() (*time.Time, bool)`

GetValidityEndOk returns a tuple with the ValidityEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityEnd

`func (o *TokenauthNodeTokenRequest) SetValidityEnd(v time.Time)`

SetValidityEnd sets ValidityEnd field to given value.

### HasValidityEnd

`func (o *TokenauthNodeTokenRequest) HasValidityEnd() bool`

HasValidityEnd returns a boolean if a field has been set.

### GetValidityStart

`func (o *TokenauthNodeTokenRequest) GetValidityStart() time.Time`

GetValidityStart returns the ValidityStart field if non-nil, zero value otherwise.

### GetValidityStartOk

`func (o *TokenauthNodeTokenRequest) GetValidityStartOk() (*time.Time, bool)`

GetValidityStartOk returns a tuple with the ValidityStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityStart

`func (o *TokenauthNodeTokenRequest) SetValidityStart(v time.Time)`

SetValidityStart sets ValidityStart field to given value.

### HasValidityStart

`func (o *TokenauthNodeTokenRequest) HasValidityStart() bool`

HasValidityStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


