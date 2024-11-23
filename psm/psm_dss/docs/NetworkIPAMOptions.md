# NetworkIPAMOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClasslessStaticRoutes** | Pointer to [**[]NetworkClasslessStaticRoute**](NetworkClasslessStaticRoute.md) |  | [optional] 
**Lease** | Pointer to **int64** |  | [optional] 
**Routers** | Pointer to **[]string** |  | [optional] 

## Methods

### NewNetworkIPAMOptions

`func NewNetworkIPAMOptions() *NetworkIPAMOptions`

NewNetworkIPAMOptions instantiates a new NetworkIPAMOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkIPAMOptionsWithDefaults

`func NewNetworkIPAMOptionsWithDefaults() *NetworkIPAMOptions`

NewNetworkIPAMOptionsWithDefaults instantiates a new NetworkIPAMOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClasslessStaticRoutes

`func (o *NetworkIPAMOptions) GetClasslessStaticRoutes() []NetworkClasslessStaticRoute`

GetClasslessStaticRoutes returns the ClasslessStaticRoutes field if non-nil, zero value otherwise.

### GetClasslessStaticRoutesOk

`func (o *NetworkIPAMOptions) GetClasslessStaticRoutesOk() (*[]NetworkClasslessStaticRoute, bool)`

GetClasslessStaticRoutesOk returns a tuple with the ClasslessStaticRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClasslessStaticRoutes

`func (o *NetworkIPAMOptions) SetClasslessStaticRoutes(v []NetworkClasslessStaticRoute)`

SetClasslessStaticRoutes sets ClasslessStaticRoutes field to given value.

### HasClasslessStaticRoutes

`func (o *NetworkIPAMOptions) HasClasslessStaticRoutes() bool`

HasClasslessStaticRoutes returns a boolean if a field has been set.

### GetLease

`func (o *NetworkIPAMOptions) GetLease() int64`

GetLease returns the Lease field if non-nil, zero value otherwise.

### GetLeaseOk

`func (o *NetworkIPAMOptions) GetLeaseOk() (*int64, bool)`

GetLeaseOk returns a tuple with the Lease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLease

`func (o *NetworkIPAMOptions) SetLease(v int64)`

SetLease sets Lease field to given value.

### HasLease

`func (o *NetworkIPAMOptions) HasLease() bool`

HasLease returns a boolean if a field has been set.

### GetRouters

`func (o *NetworkIPAMOptions) GetRouters() []string`

GetRouters returns the Routers field if non-nil, zero value otherwise.

### GetRoutersOk

`func (o *NetworkIPAMOptions) GetRoutersOk() (*[]string, bool)`

GetRoutersOk returns a tuple with the Routers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouters

`func (o *NetworkIPAMOptions) SetRouters(v []string)`

SetRouters sets Routers field to given value.

### HasRouters

`func (o *NetworkIPAMOptions) HasRouters() bool`

HasRouters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


