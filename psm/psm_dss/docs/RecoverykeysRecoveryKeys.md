# RecoverykeysRecoveryKeys

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrivateKey** | Pointer to **string** | PrivateKey is the private key generated at cluster bootstrap. | [optional] 
**PsmVersion** | Pointer to **string** | PsmVersion is the version of the PSM SW corresponding to the recovery keys. | [optional] 
**TrustChain** | Pointer to **[]string** | TrustChain is the chain of intermediate certificates that are needed to establish the validity of the public key. | [optional] 
**TrustRoots** | Pointer to **[]string** | TrustRoot is the certificate of an entity used as a root CA. | [optional] 

## Methods

### NewRecoverykeysRecoveryKeys

`func NewRecoverykeysRecoveryKeys() *RecoverykeysRecoveryKeys`

NewRecoverykeysRecoveryKeys instantiates a new RecoverykeysRecoveryKeys object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverykeysRecoveryKeysWithDefaults

`func NewRecoverykeysRecoveryKeysWithDefaults() *RecoverykeysRecoveryKeys`

NewRecoverykeysRecoveryKeysWithDefaults instantiates a new RecoverykeysRecoveryKeys object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivateKey

`func (o *RecoverykeysRecoveryKeys) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *RecoverykeysRecoveryKeys) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *RecoverykeysRecoveryKeys) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *RecoverykeysRecoveryKeys) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetPsmVersion

`func (o *RecoverykeysRecoveryKeys) GetPsmVersion() string`

GetPsmVersion returns the PsmVersion field if non-nil, zero value otherwise.

### GetPsmVersionOk

`func (o *RecoverykeysRecoveryKeys) GetPsmVersionOk() (*string, bool)`

GetPsmVersionOk returns a tuple with the PsmVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsmVersion

`func (o *RecoverykeysRecoveryKeys) SetPsmVersion(v string)`

SetPsmVersion sets PsmVersion field to given value.

### HasPsmVersion

`func (o *RecoverykeysRecoveryKeys) HasPsmVersion() bool`

HasPsmVersion returns a boolean if a field has been set.

### GetTrustChain

`func (o *RecoverykeysRecoveryKeys) GetTrustChain() []string`

GetTrustChain returns the TrustChain field if non-nil, zero value otherwise.

### GetTrustChainOk

`func (o *RecoverykeysRecoveryKeys) GetTrustChainOk() (*[]string, bool)`

GetTrustChainOk returns a tuple with the TrustChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustChain

`func (o *RecoverykeysRecoveryKeys) SetTrustChain(v []string)`

SetTrustChain sets TrustChain field to given value.

### HasTrustChain

`func (o *RecoverykeysRecoveryKeys) HasTrustChain() bool`

HasTrustChain returns a boolean if a field has been set.

### GetTrustRoots

`func (o *RecoverykeysRecoveryKeys) GetTrustRoots() []string`

GetTrustRoots returns the TrustRoots field if non-nil, zero value otherwise.

### GetTrustRootsOk

`func (o *RecoverykeysRecoveryKeys) GetTrustRootsOk() (*[]string, bool)`

GetTrustRootsOk returns a tuple with the TrustRoots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustRoots

`func (o *RecoverykeysRecoveryKeys) SetTrustRoots(v []string)`

SetTrustRoots sets TrustRoots field to given value.

### HasTrustRoots

`func (o *RecoverykeysRecoveryKeys) HasTrustRoots() bool`

HasTrustRoots returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


