# NetworkVirtualRouterStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Handle allocated in the system. | [optional] 
**PropagationStatus** | Pointer to [**SecurityPropagationStatus**](securityPropagationStatus.md) |  | [optional] 
**Rd** | Pointer to [**NetworkRouteDistinguisher**](networkRouteDistinguisher.md) |  | [optional] 
**RouteTable** | Pointer to **string** |  | [optional] 

## Methods

### NewNetworkVirtualRouterStatus

`func NewNetworkVirtualRouterStatus() *NetworkVirtualRouterStatus`

NewNetworkVirtualRouterStatus instantiates a new NetworkVirtualRouterStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkVirtualRouterStatusWithDefaults

`func NewNetworkVirtualRouterStatusWithDefaults() *NetworkVirtualRouterStatus`

NewNetworkVirtualRouterStatusWithDefaults instantiates a new NetworkVirtualRouterStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkVirtualRouterStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkVirtualRouterStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkVirtualRouterStatus) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkVirtualRouterStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPropagationStatus

`func (o *NetworkVirtualRouterStatus) GetPropagationStatus() SecurityPropagationStatus`

GetPropagationStatus returns the PropagationStatus field if non-nil, zero value otherwise.

### GetPropagationStatusOk

`func (o *NetworkVirtualRouterStatus) GetPropagationStatusOk() (*SecurityPropagationStatus, bool)`

GetPropagationStatusOk returns a tuple with the PropagationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropagationStatus

`func (o *NetworkVirtualRouterStatus) SetPropagationStatus(v SecurityPropagationStatus)`

SetPropagationStatus sets PropagationStatus field to given value.

### HasPropagationStatus

`func (o *NetworkVirtualRouterStatus) HasPropagationStatus() bool`

HasPropagationStatus returns a boolean if a field has been set.

### GetRd

`func (o *NetworkVirtualRouterStatus) GetRd() NetworkRouteDistinguisher`

GetRd returns the Rd field if non-nil, zero value otherwise.

### GetRdOk

`func (o *NetworkVirtualRouterStatus) GetRdOk() (*NetworkRouteDistinguisher, bool)`

GetRdOk returns a tuple with the Rd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRd

`func (o *NetworkVirtualRouterStatus) SetRd(v NetworkRouteDistinguisher)`

SetRd sets Rd field to given value.

### HasRd

`func (o *NetworkVirtualRouterStatus) HasRd() bool`

HasRd returns a boolean if a field has been set.

### GetRouteTable

`func (o *NetworkVirtualRouterStatus) GetRouteTable() string`

GetRouteTable returns the RouteTable field if non-nil, zero value otherwise.

### GetRouteTableOk

`func (o *NetworkVirtualRouterStatus) GetRouteTableOk() (*string, bool)`

GetRouteTableOk returns a tuple with the RouteTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteTable

`func (o *NetworkVirtualRouterStatus) SetRouteTable(v string)`

SetRouteTable sets RouteTable field to given value.

### HasRouteTable

`func (o *NetworkVirtualRouterStatus) HasRouteTable() bool`

HasRouteTable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


