# NetworkRouteDistinguisher

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminValue** | Pointer to **map[string]interface{}** | Administrator subfield of Value. Length depends on Type. swagger-tags: type&#x3D;any. | [optional] 
**AssignedValue** | Pointer to **int64** | Assigned subfield of Value. Length depends on Type. | [optional] 
**Type** | Pointer to **string** | RD Type as in rfc4364. | [optional] [default to "type0"]

## Methods

### NewNetworkRouteDistinguisher

`func NewNetworkRouteDistinguisher() *NetworkRouteDistinguisher`

NewNetworkRouteDistinguisher instantiates a new NetworkRouteDistinguisher object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkRouteDistinguisherWithDefaults

`func NewNetworkRouteDistinguisherWithDefaults() *NetworkRouteDistinguisher`

NewNetworkRouteDistinguisherWithDefaults instantiates a new NetworkRouteDistinguisher object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminValue

`func (o *NetworkRouteDistinguisher) GetAdminValue() map[string]interface{}`

GetAdminValue returns the AdminValue field if non-nil, zero value otherwise.

### GetAdminValueOk

`func (o *NetworkRouteDistinguisher) GetAdminValueOk() (*map[string]interface{}, bool)`

GetAdminValueOk returns a tuple with the AdminValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminValue

`func (o *NetworkRouteDistinguisher) SetAdminValue(v map[string]interface{})`

SetAdminValue sets AdminValue field to given value.

### HasAdminValue

`func (o *NetworkRouteDistinguisher) HasAdminValue() bool`

HasAdminValue returns a boolean if a field has been set.

### GetAssignedValue

`func (o *NetworkRouteDistinguisher) GetAssignedValue() int64`

GetAssignedValue returns the AssignedValue field if non-nil, zero value otherwise.

### GetAssignedValueOk

`func (o *NetworkRouteDistinguisher) GetAssignedValueOk() (*int64, bool)`

GetAssignedValueOk returns a tuple with the AssignedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedValue

`func (o *NetworkRouteDistinguisher) SetAssignedValue(v int64)`

SetAssignedValue sets AssignedValue field to given value.

### HasAssignedValue

`func (o *NetworkRouteDistinguisher) HasAssignedValue() bool`

HasAssignedValue returns a boolean if a field has been set.

### GetType

`func (o *NetworkRouteDistinguisher) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkRouteDistinguisher) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkRouteDistinguisher) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NetworkRouteDistinguisher) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


