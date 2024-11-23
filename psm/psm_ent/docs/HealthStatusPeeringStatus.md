# HealthStatusPeeringStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configured** | Pointer to **int32** |  | [optional] 
**DownPeers** | Pointer to **[]string** |  | [optional] 
**Established** | Pointer to **int32** |  | [optional] 

## Methods

### NewHealthStatusPeeringStatus

`func NewHealthStatusPeeringStatus() *HealthStatusPeeringStatus`

NewHealthStatusPeeringStatus instantiates a new HealthStatusPeeringStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthStatusPeeringStatusWithDefaults

`func NewHealthStatusPeeringStatusWithDefaults() *HealthStatusPeeringStatus`

NewHealthStatusPeeringStatusWithDefaults instantiates a new HealthStatusPeeringStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigured

`func (o *HealthStatusPeeringStatus) GetConfigured() int32`

GetConfigured returns the Configured field if non-nil, zero value otherwise.

### GetConfiguredOk

`func (o *HealthStatusPeeringStatus) GetConfiguredOk() (*int32, bool)`

GetConfiguredOk returns a tuple with the Configured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigured

`func (o *HealthStatusPeeringStatus) SetConfigured(v int32)`

SetConfigured sets Configured field to given value.

### HasConfigured

`func (o *HealthStatusPeeringStatus) HasConfigured() bool`

HasConfigured returns a boolean if a field has been set.

### GetDownPeers

`func (o *HealthStatusPeeringStatus) GetDownPeers() []string`

GetDownPeers returns the DownPeers field if non-nil, zero value otherwise.

### GetDownPeersOk

`func (o *HealthStatusPeeringStatus) GetDownPeersOk() (*[]string, bool)`

GetDownPeersOk returns a tuple with the DownPeers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownPeers

`func (o *HealthStatusPeeringStatus) SetDownPeers(v []string)`

SetDownPeers sets DownPeers field to given value.

### HasDownPeers

`func (o *HealthStatusPeeringStatus) HasDownPeers() bool`

HasDownPeers returns a boolean if a field has been set.

### GetEstablished

`func (o *HealthStatusPeeringStatus) GetEstablished() int32`

GetEstablished returns the Established field if non-nil, zero value otherwise.

### GetEstablishedOk

`func (o *HealthStatusPeeringStatus) GetEstablishedOk() (*int32, bool)`

GetEstablishedOk returns a tuple with the Established field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstablished

`func (o *HealthStatusPeeringStatus) SetEstablished(v int32)`

SetEstablished sets Established field to given value.

### HasEstablished

`func (o *HealthStatusPeeringStatus) HasEstablished() bool`

HasEstablished returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


