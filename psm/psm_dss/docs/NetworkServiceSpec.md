# NetworkServiceSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LbPolicy** | Pointer to **string** | load balancing policy (reference to LbPolicy object). | [optional] 
**Ports** | Pointer to **string** | load balancer port. | [optional] 
**TlsClientPolicy** | Pointer to [**NetworkTLSClientPolicySpec**](networkTLSClientPolicySpec.md) |  | [optional] 
**TlsServerPolicy** | Pointer to [**NetworkTLSServerPolicySpec**](networkTLSServerPolicySpec.md) |  | [optional] 
**VirtualIp** | Pointer to **string** | Virtual IP of the load balancer. | [optional] 
**WorkloadLabels** | Pointer to **[]string** | FIXME: maps are not working. change this after its fixed map&lt;string, string&gt; WorkloadSelector  &#x3D; 1 [(gogoproto.nullable) &#x3D; true, (gogoproto.jsontag) &#x3D; \&quot;workload-labels,omitempty\&quot;]; workload selector for the service (list of labels to match). | [optional] 

## Methods

### NewNetworkServiceSpec

`func NewNetworkServiceSpec() *NetworkServiceSpec`

NewNetworkServiceSpec instantiates a new NetworkServiceSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkServiceSpecWithDefaults

`func NewNetworkServiceSpecWithDefaults() *NetworkServiceSpec`

NewNetworkServiceSpecWithDefaults instantiates a new NetworkServiceSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLbPolicy

`func (o *NetworkServiceSpec) GetLbPolicy() string`

GetLbPolicy returns the LbPolicy field if non-nil, zero value otherwise.

### GetLbPolicyOk

`func (o *NetworkServiceSpec) GetLbPolicyOk() (*string, bool)`

GetLbPolicyOk returns a tuple with the LbPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbPolicy

`func (o *NetworkServiceSpec) SetLbPolicy(v string)`

SetLbPolicy sets LbPolicy field to given value.

### HasLbPolicy

`func (o *NetworkServiceSpec) HasLbPolicy() bool`

HasLbPolicy returns a boolean if a field has been set.

### GetPorts

`func (o *NetworkServiceSpec) GetPorts() string`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *NetworkServiceSpec) GetPortsOk() (*string, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *NetworkServiceSpec) SetPorts(v string)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *NetworkServiceSpec) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetTlsClientPolicy

`func (o *NetworkServiceSpec) GetTlsClientPolicy() NetworkTLSClientPolicySpec`

GetTlsClientPolicy returns the TlsClientPolicy field if non-nil, zero value otherwise.

### GetTlsClientPolicyOk

`func (o *NetworkServiceSpec) GetTlsClientPolicyOk() (*NetworkTLSClientPolicySpec, bool)`

GetTlsClientPolicyOk returns a tuple with the TlsClientPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsClientPolicy

`func (o *NetworkServiceSpec) SetTlsClientPolicy(v NetworkTLSClientPolicySpec)`

SetTlsClientPolicy sets TlsClientPolicy field to given value.

### HasTlsClientPolicy

`func (o *NetworkServiceSpec) HasTlsClientPolicy() bool`

HasTlsClientPolicy returns a boolean if a field has been set.

### GetTlsServerPolicy

`func (o *NetworkServiceSpec) GetTlsServerPolicy() NetworkTLSServerPolicySpec`

GetTlsServerPolicy returns the TlsServerPolicy field if non-nil, zero value otherwise.

### GetTlsServerPolicyOk

`func (o *NetworkServiceSpec) GetTlsServerPolicyOk() (*NetworkTLSServerPolicySpec, bool)`

GetTlsServerPolicyOk returns a tuple with the TlsServerPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsServerPolicy

`func (o *NetworkServiceSpec) SetTlsServerPolicy(v NetworkTLSServerPolicySpec)`

SetTlsServerPolicy sets TlsServerPolicy field to given value.

### HasTlsServerPolicy

`func (o *NetworkServiceSpec) HasTlsServerPolicy() bool`

HasTlsServerPolicy returns a boolean if a field has been set.

### GetVirtualIp

`func (o *NetworkServiceSpec) GetVirtualIp() string`

GetVirtualIp returns the VirtualIp field if non-nil, zero value otherwise.

### GetVirtualIpOk

`func (o *NetworkServiceSpec) GetVirtualIpOk() (*string, bool)`

GetVirtualIpOk returns a tuple with the VirtualIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualIp

`func (o *NetworkServiceSpec) SetVirtualIp(v string)`

SetVirtualIp sets VirtualIp field to given value.

### HasVirtualIp

`func (o *NetworkServiceSpec) HasVirtualIp() bool`

HasVirtualIp returns a boolean if a field has been set.

### GetWorkloadLabels

`func (o *NetworkServiceSpec) GetWorkloadLabels() []string`

GetWorkloadLabels returns the WorkloadLabels field if non-nil, zero value otherwise.

### GetWorkloadLabelsOk

`func (o *NetworkServiceSpec) GetWorkloadLabelsOk() (*[]string, bool)`

GetWorkloadLabelsOk returns a tuple with the WorkloadLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadLabels

`func (o *NetworkServiceSpec) SetWorkloadLabels(v []string)`

SetWorkloadLabels sets WorkloadLabels field to given value.

### HasWorkloadLabels

`func (o *NetworkServiceSpec) HasWorkloadLabels() bool`

HasWorkloadLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


