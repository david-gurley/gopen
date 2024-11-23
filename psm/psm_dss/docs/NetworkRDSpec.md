# NetworkRDSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressFamily** | Pointer to **string** | Address family where this config applies. | [optional] [default to "ipv4-unicast"]
**Rd** | Pointer to [**NetworkRouteDistinguisher**](networkRouteDistinguisher.md) |  | [optional] 
**RdAuto** | Pointer to **bool** | True indicates the system will generate the RD automatically. | [optional] 
**RtExport** | Pointer to [**[]NetworkRouteDistinguisher**](NetworkRouteDistinguisher.md) | Route Targets to Export. | [optional] 
**RtImport** | Pointer to [**[]NetworkRouteDistinguisher**](NetworkRouteDistinguisher.md) | Route Targets to Import. | [optional] 

## Methods

### NewNetworkRDSpec

`func NewNetworkRDSpec() *NetworkRDSpec`

NewNetworkRDSpec instantiates a new NetworkRDSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkRDSpecWithDefaults

`func NewNetworkRDSpecWithDefaults() *NetworkRDSpec`

NewNetworkRDSpecWithDefaults instantiates a new NetworkRDSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressFamily

`func (o *NetworkRDSpec) GetAddressFamily() string`

GetAddressFamily returns the AddressFamily field if non-nil, zero value otherwise.

### GetAddressFamilyOk

`func (o *NetworkRDSpec) GetAddressFamilyOk() (*string, bool)`

GetAddressFamilyOk returns a tuple with the AddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFamily

`func (o *NetworkRDSpec) SetAddressFamily(v string)`

SetAddressFamily sets AddressFamily field to given value.

### HasAddressFamily

`func (o *NetworkRDSpec) HasAddressFamily() bool`

HasAddressFamily returns a boolean if a field has been set.

### GetRd

`func (o *NetworkRDSpec) GetRd() NetworkRouteDistinguisher`

GetRd returns the Rd field if non-nil, zero value otherwise.

### GetRdOk

`func (o *NetworkRDSpec) GetRdOk() (*NetworkRouteDistinguisher, bool)`

GetRdOk returns a tuple with the Rd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRd

`func (o *NetworkRDSpec) SetRd(v NetworkRouteDistinguisher)`

SetRd sets Rd field to given value.

### HasRd

`func (o *NetworkRDSpec) HasRd() bool`

HasRd returns a boolean if a field has been set.

### GetRdAuto

`func (o *NetworkRDSpec) GetRdAuto() bool`

GetRdAuto returns the RdAuto field if non-nil, zero value otherwise.

### GetRdAutoOk

`func (o *NetworkRDSpec) GetRdAutoOk() (*bool, bool)`

GetRdAutoOk returns a tuple with the RdAuto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdAuto

`func (o *NetworkRDSpec) SetRdAuto(v bool)`

SetRdAuto sets RdAuto field to given value.

### HasRdAuto

`func (o *NetworkRDSpec) HasRdAuto() bool`

HasRdAuto returns a boolean if a field has been set.

### GetRtExport

`func (o *NetworkRDSpec) GetRtExport() []NetworkRouteDistinguisher`

GetRtExport returns the RtExport field if non-nil, zero value otherwise.

### GetRtExportOk

`func (o *NetworkRDSpec) GetRtExportOk() (*[]NetworkRouteDistinguisher, bool)`

GetRtExportOk returns a tuple with the RtExport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtExport

`func (o *NetworkRDSpec) SetRtExport(v []NetworkRouteDistinguisher)`

SetRtExport sets RtExport field to given value.

### HasRtExport

`func (o *NetworkRDSpec) HasRtExport() bool`

HasRtExport returns a boolean if a field has been set.

### GetRtImport

`func (o *NetworkRDSpec) GetRtImport() []NetworkRouteDistinguisher`

GetRtImport returns the RtImport field if non-nil, zero value otherwise.

### GetRtImportOk

`func (o *NetworkRDSpec) GetRtImportOk() (*[]NetworkRouteDistinguisher, bool)`

GetRtImportOk returns a tuple with the RtImport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtImport

`func (o *NetworkRDSpec) SetRtImport(v []NetworkRouteDistinguisher)`

SetRtImport sets RtImport field to given value.

### HasRtImport

`func (o *NetworkRDSpec) HasRtImport() bool`

HasRtImport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


