# NetworkNetworkInterfaceStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterNode** | Pointer to **string** | Set only if interface is on Venice Node. | [optional] 
**Dsc** | Pointer to **string** |  | [optional] 
**DscId** | Pointer to **string** |  | [optional] 
**IfHostStatus** | Pointer to [**NetworkNetworkInterfaceHostStatus**](networkNetworkInterfaceHostStatus.md) |  | [optional] 
**IfUplinkStatus** | Pointer to [**NetworkNetworkInterfaceUplinkStatus**](networkNetworkInterfaceUplinkStatus.md) |  | [optional] 
**MirrorSessions** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OperStatus** | Pointer to **string** |  | [optional] [default to "up"]
**PrimaryMac** | Pointer to **string** | Should be a valid MAC address. | [optional] 
**PropagationStatus** | Pointer to [**SecurityPropagationStatus**](securityPropagationStatus.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] [default to "none"]

## Methods

### NewNetworkNetworkInterfaceStatus

`func NewNetworkNetworkInterfaceStatus() *NetworkNetworkInterfaceStatus`

NewNetworkNetworkInterfaceStatus instantiates a new NetworkNetworkInterfaceStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkNetworkInterfaceStatusWithDefaults

`func NewNetworkNetworkInterfaceStatusWithDefaults() *NetworkNetworkInterfaceStatus`

NewNetworkNetworkInterfaceStatusWithDefaults instantiates a new NetworkNetworkInterfaceStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterNode

`func (o *NetworkNetworkInterfaceStatus) GetClusterNode() string`

GetClusterNode returns the ClusterNode field if non-nil, zero value otherwise.

### GetClusterNodeOk

`func (o *NetworkNetworkInterfaceStatus) GetClusterNodeOk() (*string, bool)`

GetClusterNodeOk returns a tuple with the ClusterNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterNode

`func (o *NetworkNetworkInterfaceStatus) SetClusterNode(v string)`

SetClusterNode sets ClusterNode field to given value.

### HasClusterNode

`func (o *NetworkNetworkInterfaceStatus) HasClusterNode() bool`

HasClusterNode returns a boolean if a field has been set.

### GetDsc

`func (o *NetworkNetworkInterfaceStatus) GetDsc() string`

GetDsc returns the Dsc field if non-nil, zero value otherwise.

### GetDscOk

`func (o *NetworkNetworkInterfaceStatus) GetDscOk() (*string, bool)`

GetDscOk returns a tuple with the Dsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsc

`func (o *NetworkNetworkInterfaceStatus) SetDsc(v string)`

SetDsc sets Dsc field to given value.

### HasDsc

`func (o *NetworkNetworkInterfaceStatus) HasDsc() bool`

HasDsc returns a boolean if a field has been set.

### GetDscId

`func (o *NetworkNetworkInterfaceStatus) GetDscId() string`

GetDscId returns the DscId field if non-nil, zero value otherwise.

### GetDscIdOk

`func (o *NetworkNetworkInterfaceStatus) GetDscIdOk() (*string, bool)`

GetDscIdOk returns a tuple with the DscId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDscId

`func (o *NetworkNetworkInterfaceStatus) SetDscId(v string)`

SetDscId sets DscId field to given value.

### HasDscId

`func (o *NetworkNetworkInterfaceStatus) HasDscId() bool`

HasDscId returns a boolean if a field has been set.

### GetIfHostStatus

`func (o *NetworkNetworkInterfaceStatus) GetIfHostStatus() NetworkNetworkInterfaceHostStatus`

GetIfHostStatus returns the IfHostStatus field if non-nil, zero value otherwise.

### GetIfHostStatusOk

`func (o *NetworkNetworkInterfaceStatus) GetIfHostStatusOk() (*NetworkNetworkInterfaceHostStatus, bool)`

GetIfHostStatusOk returns a tuple with the IfHostStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfHostStatus

`func (o *NetworkNetworkInterfaceStatus) SetIfHostStatus(v NetworkNetworkInterfaceHostStatus)`

SetIfHostStatus sets IfHostStatus field to given value.

### HasIfHostStatus

`func (o *NetworkNetworkInterfaceStatus) HasIfHostStatus() bool`

HasIfHostStatus returns a boolean if a field has been set.

### GetIfUplinkStatus

`func (o *NetworkNetworkInterfaceStatus) GetIfUplinkStatus() NetworkNetworkInterfaceUplinkStatus`

GetIfUplinkStatus returns the IfUplinkStatus field if non-nil, zero value otherwise.

### GetIfUplinkStatusOk

`func (o *NetworkNetworkInterfaceStatus) GetIfUplinkStatusOk() (*NetworkNetworkInterfaceUplinkStatus, bool)`

GetIfUplinkStatusOk returns a tuple with the IfUplinkStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfUplinkStatus

`func (o *NetworkNetworkInterfaceStatus) SetIfUplinkStatus(v NetworkNetworkInterfaceUplinkStatus)`

SetIfUplinkStatus sets IfUplinkStatus field to given value.

### HasIfUplinkStatus

`func (o *NetworkNetworkInterfaceStatus) HasIfUplinkStatus() bool`

HasIfUplinkStatus returns a boolean if a field has been set.

### GetMirrorSessions

`func (o *NetworkNetworkInterfaceStatus) GetMirrorSessions() []string`

GetMirrorSessions returns the MirrorSessions field if non-nil, zero value otherwise.

### GetMirrorSessionsOk

`func (o *NetworkNetworkInterfaceStatus) GetMirrorSessionsOk() (*[]string, bool)`

GetMirrorSessionsOk returns a tuple with the MirrorSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirrorSessions

`func (o *NetworkNetworkInterfaceStatus) SetMirrorSessions(v []string)`

SetMirrorSessions sets MirrorSessions field to given value.

### HasMirrorSessions

`func (o *NetworkNetworkInterfaceStatus) HasMirrorSessions() bool`

HasMirrorSessions returns a boolean if a field has been set.

### GetName

`func (o *NetworkNetworkInterfaceStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkNetworkInterfaceStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkNetworkInterfaceStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkNetworkInterfaceStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperStatus

`func (o *NetworkNetworkInterfaceStatus) GetOperStatus() string`

GetOperStatus returns the OperStatus field if non-nil, zero value otherwise.

### GetOperStatusOk

`func (o *NetworkNetworkInterfaceStatus) GetOperStatusOk() (*string, bool)`

GetOperStatusOk returns a tuple with the OperStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperStatus

`func (o *NetworkNetworkInterfaceStatus) SetOperStatus(v string)`

SetOperStatus sets OperStatus field to given value.

### HasOperStatus

`func (o *NetworkNetworkInterfaceStatus) HasOperStatus() bool`

HasOperStatus returns a boolean if a field has been set.

### GetPrimaryMac

`func (o *NetworkNetworkInterfaceStatus) GetPrimaryMac() string`

GetPrimaryMac returns the PrimaryMac field if non-nil, zero value otherwise.

### GetPrimaryMacOk

`func (o *NetworkNetworkInterfaceStatus) GetPrimaryMacOk() (*string, bool)`

GetPrimaryMacOk returns a tuple with the PrimaryMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryMac

`func (o *NetworkNetworkInterfaceStatus) SetPrimaryMac(v string)`

SetPrimaryMac sets PrimaryMac field to given value.

### HasPrimaryMac

`func (o *NetworkNetworkInterfaceStatus) HasPrimaryMac() bool`

HasPrimaryMac returns a boolean if a field has been set.

### GetPropagationStatus

`func (o *NetworkNetworkInterfaceStatus) GetPropagationStatus() SecurityPropagationStatus`

GetPropagationStatus returns the PropagationStatus field if non-nil, zero value otherwise.

### GetPropagationStatusOk

`func (o *NetworkNetworkInterfaceStatus) GetPropagationStatusOk() (*SecurityPropagationStatus, bool)`

GetPropagationStatusOk returns a tuple with the PropagationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropagationStatus

`func (o *NetworkNetworkInterfaceStatus) SetPropagationStatus(v SecurityPropagationStatus)`

SetPropagationStatus sets PropagationStatus field to given value.

### HasPropagationStatus

`func (o *NetworkNetworkInterfaceStatus) HasPropagationStatus() bool`

HasPropagationStatus returns a boolean if a field has been set.

### GetType

`func (o *NetworkNetworkInterfaceStatus) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkNetworkInterfaceStatus) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkNetworkInterfaceStatus) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NetworkNetworkInterfaceStatus) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


