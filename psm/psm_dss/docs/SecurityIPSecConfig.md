# SecurityIPSecConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IkeLifetime** | Pointer to **string** | Time to schedule Internet Key Exchange (IKE) reauthentication. IKE reauthentication recreates the IKE SA from scratch and re-evaluates the credentials Default 24h, set it to an empty string to disable reauthentication. Should be a valid time duration between 1h0m0s and 24h0m0s. | [optional] [default to "24h"]
**SaLifetime** | Pointer to **string** | How long a particular instance of a connection should last, from successful negotiation to expiry The connection will be re-negotiated before it expires. Post succesful re-negotation, the new connection will have new(different) keys and a new SPI Default 8h, Max 24h, set it to an empty string to disable rekeying. Should be a valid time duration between 1h0m0s and 24h0m0s. | [optional] [default to "8h"]

## Methods

### NewSecurityIPSecConfig

`func NewSecurityIPSecConfig() *SecurityIPSecConfig`

NewSecurityIPSecConfig instantiates a new SecurityIPSecConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityIPSecConfigWithDefaults

`func NewSecurityIPSecConfigWithDefaults() *SecurityIPSecConfig`

NewSecurityIPSecConfigWithDefaults instantiates a new SecurityIPSecConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIkeLifetime

`func (o *SecurityIPSecConfig) GetIkeLifetime() string`

GetIkeLifetime returns the IkeLifetime field if non-nil, zero value otherwise.

### GetIkeLifetimeOk

`func (o *SecurityIPSecConfig) GetIkeLifetimeOk() (*string, bool)`

GetIkeLifetimeOk returns a tuple with the IkeLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIkeLifetime

`func (o *SecurityIPSecConfig) SetIkeLifetime(v string)`

SetIkeLifetime sets IkeLifetime field to given value.

### HasIkeLifetime

`func (o *SecurityIPSecConfig) HasIkeLifetime() bool`

HasIkeLifetime returns a boolean if a field has been set.

### GetSaLifetime

`func (o *SecurityIPSecConfig) GetSaLifetime() string`

GetSaLifetime returns the SaLifetime field if non-nil, zero value otherwise.

### GetSaLifetimeOk

`func (o *SecurityIPSecConfig) GetSaLifetimeOk() (*string, bool)`

GetSaLifetimeOk returns a tuple with the SaLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaLifetime

`func (o *SecurityIPSecConfig) SetSaLifetime(v string)`

SetSaLifetime sets SaLifetime field to given value.

### HasSaLifetime

`func (o *SecurityIPSecConfig) HasSaLifetime() bool`

HasSaLifetime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


