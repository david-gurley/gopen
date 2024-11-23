# ClusterDistributedServiceCardStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DSCSku** | Pointer to **string** | DSC SKU. | [optional] 
**DSCVersion** | Pointer to **string** | DSC Version. | [optional] 
**AdmPhaseReason** | Pointer to **string** | The reason why the DistributedServiceCard is not in ADMITTED state. | [optional] 
**AdmissionPhase** | Pointer to **string** | Current admission phase of the DistributedServiceCard. When auto-admission is enabled, AdmissionPhase will be set to NIC_ADMITTED by CMD for validated NICs. When auto-admission is not enabled, AdmissionPhase will be set to NIC_PENDING by CMD for validated NICs since it requires manual approval. To admit the NIC as a part of manual admission, user is expected to set Spec.Admit to true for the NICs that are in NIC_PENDING state. Note : Whitelist mode is not supported yet. | [optional] [default to "unknown"]
**AlomPresent** | Pointer to **bool** | ALOMPresent true value indicates ALOM cable is present. | [optional] 
**Conditions** | Pointer to [**[]ClusterDSCCondition**](ClusterDSCCondition.md) | List of current NIC conditions. | [optional] 
**ControlPlaneStatus** | Pointer to [**ClusterDSCControlPlaneStatus**](clusterDSCControlPlaneStatus.md) |  | [optional] 
**DssInfo** | Pointer to [**ClusterDSSInfo**](clusterDSSInfo.md) |  | [optional] 
**Host** | Pointer to **string** | The name of the host this DistributedServiceCard is plugged into. | [optional] 
**InbandIpConfig** | Pointer to [**ClusterIPConfig**](clusterIPConfig.md) |  | [optional] 
**Interfaces** | Pointer to **[]string** | Network Interfaces. | [optional] 
**IpConfig** | Pointer to [**ClusterIPConfig**](clusterIPConfig.md) |  | [optional] 
**IsConnectedToPsm** | Pointer to **bool** | IsConnectedToPSM is set to true if connected to PSM. | [optional] 
**NumMacAddress** | Pointer to **int64** | NumMacAddress has the number of mac addresses that is available on this DSC. Value should be between 0 and 256. | [optional] [default to 24]
**PackageType** | Pointer to **string** | Type of DSC. | [optional] [default to "dsc"]
**PrimaryMac** | Pointer to **string** | PrimaryMAC is the MAC address of the primary PF exposed by DistributedServiceCard. Should be a valid MAC address. | [optional] 
**SecureBooted** | Pointer to **bool** | SecureBooted a true value indicates, secure boot is enabled. | [optional] 
**SecurityPolicyRuleScaleProfile** | Pointer to **string** | SecurityPolicyRuleScaleProfile is the active security policy rule scale profile in the DSE. | [optional] 
**SerialNum** | Pointer to **string** | Serial number. | [optional] 
**SystemInfo** | Pointer to [**ClusterDSCInfo**](clusterDSCInfo.md) |  | [optional] 
**UnhealthyServices** | Pointer to **[]string** | Lists the unhealthy services of a distributed service card. | [optional] 
**VersionMismatch** | Pointer to **bool** | Set to true if venice and dsc versions are incompatible. | [optional] 

## Methods

### NewClusterDistributedServiceCardStatus

`func NewClusterDistributedServiceCardStatus() *ClusterDistributedServiceCardStatus`

NewClusterDistributedServiceCardStatus instantiates a new ClusterDistributedServiceCardStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterDistributedServiceCardStatusWithDefaults

`func NewClusterDistributedServiceCardStatusWithDefaults() *ClusterDistributedServiceCardStatus`

NewClusterDistributedServiceCardStatusWithDefaults instantiates a new ClusterDistributedServiceCardStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDSCSku

`func (o *ClusterDistributedServiceCardStatus) GetDSCSku() string`

GetDSCSku returns the DSCSku field if non-nil, zero value otherwise.

### GetDSCSkuOk

`func (o *ClusterDistributedServiceCardStatus) GetDSCSkuOk() (*string, bool)`

GetDSCSkuOk returns a tuple with the DSCSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDSCSku

`func (o *ClusterDistributedServiceCardStatus) SetDSCSku(v string)`

SetDSCSku sets DSCSku field to given value.

### HasDSCSku

`func (o *ClusterDistributedServiceCardStatus) HasDSCSku() bool`

HasDSCSku returns a boolean if a field has been set.

### GetDSCVersion

`func (o *ClusterDistributedServiceCardStatus) GetDSCVersion() string`

GetDSCVersion returns the DSCVersion field if non-nil, zero value otherwise.

### GetDSCVersionOk

`func (o *ClusterDistributedServiceCardStatus) GetDSCVersionOk() (*string, bool)`

GetDSCVersionOk returns a tuple with the DSCVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDSCVersion

`func (o *ClusterDistributedServiceCardStatus) SetDSCVersion(v string)`

SetDSCVersion sets DSCVersion field to given value.

### HasDSCVersion

`func (o *ClusterDistributedServiceCardStatus) HasDSCVersion() bool`

HasDSCVersion returns a boolean if a field has been set.

### GetAdmPhaseReason

`func (o *ClusterDistributedServiceCardStatus) GetAdmPhaseReason() string`

GetAdmPhaseReason returns the AdmPhaseReason field if non-nil, zero value otherwise.

### GetAdmPhaseReasonOk

`func (o *ClusterDistributedServiceCardStatus) GetAdmPhaseReasonOk() (*string, bool)`

GetAdmPhaseReasonOk returns a tuple with the AdmPhaseReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmPhaseReason

`func (o *ClusterDistributedServiceCardStatus) SetAdmPhaseReason(v string)`

SetAdmPhaseReason sets AdmPhaseReason field to given value.

### HasAdmPhaseReason

`func (o *ClusterDistributedServiceCardStatus) HasAdmPhaseReason() bool`

HasAdmPhaseReason returns a boolean if a field has been set.

### GetAdmissionPhase

`func (o *ClusterDistributedServiceCardStatus) GetAdmissionPhase() string`

GetAdmissionPhase returns the AdmissionPhase field if non-nil, zero value otherwise.

### GetAdmissionPhaseOk

`func (o *ClusterDistributedServiceCardStatus) GetAdmissionPhaseOk() (*string, bool)`

GetAdmissionPhaseOk returns a tuple with the AdmissionPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmissionPhase

`func (o *ClusterDistributedServiceCardStatus) SetAdmissionPhase(v string)`

SetAdmissionPhase sets AdmissionPhase field to given value.

### HasAdmissionPhase

`func (o *ClusterDistributedServiceCardStatus) HasAdmissionPhase() bool`

HasAdmissionPhase returns a boolean if a field has been set.

### GetAlomPresent

`func (o *ClusterDistributedServiceCardStatus) GetAlomPresent() bool`

GetAlomPresent returns the AlomPresent field if non-nil, zero value otherwise.

### GetAlomPresentOk

`func (o *ClusterDistributedServiceCardStatus) GetAlomPresentOk() (*bool, bool)`

GetAlomPresentOk returns a tuple with the AlomPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlomPresent

`func (o *ClusterDistributedServiceCardStatus) SetAlomPresent(v bool)`

SetAlomPresent sets AlomPresent field to given value.

### HasAlomPresent

`func (o *ClusterDistributedServiceCardStatus) HasAlomPresent() bool`

HasAlomPresent returns a boolean if a field has been set.

### GetConditions

`func (o *ClusterDistributedServiceCardStatus) GetConditions() []ClusterDSCCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *ClusterDistributedServiceCardStatus) GetConditionsOk() (*[]ClusterDSCCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *ClusterDistributedServiceCardStatus) SetConditions(v []ClusterDSCCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *ClusterDistributedServiceCardStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetControlPlaneStatus

`func (o *ClusterDistributedServiceCardStatus) GetControlPlaneStatus() ClusterDSCControlPlaneStatus`

GetControlPlaneStatus returns the ControlPlaneStatus field if non-nil, zero value otherwise.

### GetControlPlaneStatusOk

`func (o *ClusterDistributedServiceCardStatus) GetControlPlaneStatusOk() (*ClusterDSCControlPlaneStatus, bool)`

GetControlPlaneStatusOk returns a tuple with the ControlPlaneStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlPlaneStatus

`func (o *ClusterDistributedServiceCardStatus) SetControlPlaneStatus(v ClusterDSCControlPlaneStatus)`

SetControlPlaneStatus sets ControlPlaneStatus field to given value.

### HasControlPlaneStatus

`func (o *ClusterDistributedServiceCardStatus) HasControlPlaneStatus() bool`

HasControlPlaneStatus returns a boolean if a field has been set.

### GetDssInfo

`func (o *ClusterDistributedServiceCardStatus) GetDssInfo() ClusterDSSInfo`

GetDssInfo returns the DssInfo field if non-nil, zero value otherwise.

### GetDssInfoOk

`func (o *ClusterDistributedServiceCardStatus) GetDssInfoOk() (*ClusterDSSInfo, bool)`

GetDssInfoOk returns a tuple with the DssInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDssInfo

`func (o *ClusterDistributedServiceCardStatus) SetDssInfo(v ClusterDSSInfo)`

SetDssInfo sets DssInfo field to given value.

### HasDssInfo

`func (o *ClusterDistributedServiceCardStatus) HasDssInfo() bool`

HasDssInfo returns a boolean if a field has been set.

### GetHost

`func (o *ClusterDistributedServiceCardStatus) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ClusterDistributedServiceCardStatus) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ClusterDistributedServiceCardStatus) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *ClusterDistributedServiceCardStatus) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetInbandIpConfig

`func (o *ClusterDistributedServiceCardStatus) GetInbandIpConfig() ClusterIPConfig`

GetInbandIpConfig returns the InbandIpConfig field if non-nil, zero value otherwise.

### GetInbandIpConfigOk

`func (o *ClusterDistributedServiceCardStatus) GetInbandIpConfigOk() (*ClusterIPConfig, bool)`

GetInbandIpConfigOk returns a tuple with the InbandIpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbandIpConfig

`func (o *ClusterDistributedServiceCardStatus) SetInbandIpConfig(v ClusterIPConfig)`

SetInbandIpConfig sets InbandIpConfig field to given value.

### HasInbandIpConfig

`func (o *ClusterDistributedServiceCardStatus) HasInbandIpConfig() bool`

HasInbandIpConfig returns a boolean if a field has been set.

### GetInterfaces

`func (o *ClusterDistributedServiceCardStatus) GetInterfaces() []string`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *ClusterDistributedServiceCardStatus) GetInterfacesOk() (*[]string, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *ClusterDistributedServiceCardStatus) SetInterfaces(v []string)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *ClusterDistributedServiceCardStatus) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetIpConfig

`func (o *ClusterDistributedServiceCardStatus) GetIpConfig() ClusterIPConfig`

GetIpConfig returns the IpConfig field if non-nil, zero value otherwise.

### GetIpConfigOk

`func (o *ClusterDistributedServiceCardStatus) GetIpConfigOk() (*ClusterIPConfig, bool)`

GetIpConfigOk returns a tuple with the IpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpConfig

`func (o *ClusterDistributedServiceCardStatus) SetIpConfig(v ClusterIPConfig)`

SetIpConfig sets IpConfig field to given value.

### HasIpConfig

`func (o *ClusterDistributedServiceCardStatus) HasIpConfig() bool`

HasIpConfig returns a boolean if a field has been set.

### GetIsConnectedToPsm

`func (o *ClusterDistributedServiceCardStatus) GetIsConnectedToPsm() bool`

GetIsConnectedToPsm returns the IsConnectedToPsm field if non-nil, zero value otherwise.

### GetIsConnectedToPsmOk

`func (o *ClusterDistributedServiceCardStatus) GetIsConnectedToPsmOk() (*bool, bool)`

GetIsConnectedToPsmOk returns a tuple with the IsConnectedToPsm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConnectedToPsm

`func (o *ClusterDistributedServiceCardStatus) SetIsConnectedToPsm(v bool)`

SetIsConnectedToPsm sets IsConnectedToPsm field to given value.

### HasIsConnectedToPsm

`func (o *ClusterDistributedServiceCardStatus) HasIsConnectedToPsm() bool`

HasIsConnectedToPsm returns a boolean if a field has been set.

### GetNumMacAddress

`func (o *ClusterDistributedServiceCardStatus) GetNumMacAddress() int64`

GetNumMacAddress returns the NumMacAddress field if non-nil, zero value otherwise.

### GetNumMacAddressOk

`func (o *ClusterDistributedServiceCardStatus) GetNumMacAddressOk() (*int64, bool)`

GetNumMacAddressOk returns a tuple with the NumMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumMacAddress

`func (o *ClusterDistributedServiceCardStatus) SetNumMacAddress(v int64)`

SetNumMacAddress sets NumMacAddress field to given value.

### HasNumMacAddress

`func (o *ClusterDistributedServiceCardStatus) HasNumMacAddress() bool`

HasNumMacAddress returns a boolean if a field has been set.

### GetPackageType

`func (o *ClusterDistributedServiceCardStatus) GetPackageType() string`

GetPackageType returns the PackageType field if non-nil, zero value otherwise.

### GetPackageTypeOk

`func (o *ClusterDistributedServiceCardStatus) GetPackageTypeOk() (*string, bool)`

GetPackageTypeOk returns a tuple with the PackageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageType

`func (o *ClusterDistributedServiceCardStatus) SetPackageType(v string)`

SetPackageType sets PackageType field to given value.

### HasPackageType

`func (o *ClusterDistributedServiceCardStatus) HasPackageType() bool`

HasPackageType returns a boolean if a field has been set.

### GetPrimaryMac

`func (o *ClusterDistributedServiceCardStatus) GetPrimaryMac() string`

GetPrimaryMac returns the PrimaryMac field if non-nil, zero value otherwise.

### GetPrimaryMacOk

`func (o *ClusterDistributedServiceCardStatus) GetPrimaryMacOk() (*string, bool)`

GetPrimaryMacOk returns a tuple with the PrimaryMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryMac

`func (o *ClusterDistributedServiceCardStatus) SetPrimaryMac(v string)`

SetPrimaryMac sets PrimaryMac field to given value.

### HasPrimaryMac

`func (o *ClusterDistributedServiceCardStatus) HasPrimaryMac() bool`

HasPrimaryMac returns a boolean if a field has been set.

### GetSecureBooted

`func (o *ClusterDistributedServiceCardStatus) GetSecureBooted() bool`

GetSecureBooted returns the SecureBooted field if non-nil, zero value otherwise.

### GetSecureBootedOk

`func (o *ClusterDistributedServiceCardStatus) GetSecureBootedOk() (*bool, bool)`

GetSecureBootedOk returns a tuple with the SecureBooted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureBooted

`func (o *ClusterDistributedServiceCardStatus) SetSecureBooted(v bool)`

SetSecureBooted sets SecureBooted field to given value.

### HasSecureBooted

`func (o *ClusterDistributedServiceCardStatus) HasSecureBooted() bool`

HasSecureBooted returns a boolean if a field has been set.

### GetSecurityPolicyRuleScaleProfile

`func (o *ClusterDistributedServiceCardStatus) GetSecurityPolicyRuleScaleProfile() string`

GetSecurityPolicyRuleScaleProfile returns the SecurityPolicyRuleScaleProfile field if non-nil, zero value otherwise.

### GetSecurityPolicyRuleScaleProfileOk

`func (o *ClusterDistributedServiceCardStatus) GetSecurityPolicyRuleScaleProfileOk() (*string, bool)`

GetSecurityPolicyRuleScaleProfileOk returns a tuple with the SecurityPolicyRuleScaleProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityPolicyRuleScaleProfile

`func (o *ClusterDistributedServiceCardStatus) SetSecurityPolicyRuleScaleProfile(v string)`

SetSecurityPolicyRuleScaleProfile sets SecurityPolicyRuleScaleProfile field to given value.

### HasSecurityPolicyRuleScaleProfile

`func (o *ClusterDistributedServiceCardStatus) HasSecurityPolicyRuleScaleProfile() bool`

HasSecurityPolicyRuleScaleProfile returns a boolean if a field has been set.

### GetSerialNum

`func (o *ClusterDistributedServiceCardStatus) GetSerialNum() string`

GetSerialNum returns the SerialNum field if non-nil, zero value otherwise.

### GetSerialNumOk

`func (o *ClusterDistributedServiceCardStatus) GetSerialNumOk() (*string, bool)`

GetSerialNumOk returns a tuple with the SerialNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNum

`func (o *ClusterDistributedServiceCardStatus) SetSerialNum(v string)`

SetSerialNum sets SerialNum field to given value.

### HasSerialNum

`func (o *ClusterDistributedServiceCardStatus) HasSerialNum() bool`

HasSerialNum returns a boolean if a field has been set.

### GetSystemInfo

`func (o *ClusterDistributedServiceCardStatus) GetSystemInfo() ClusterDSCInfo`

GetSystemInfo returns the SystemInfo field if non-nil, zero value otherwise.

### GetSystemInfoOk

`func (o *ClusterDistributedServiceCardStatus) GetSystemInfoOk() (*ClusterDSCInfo, bool)`

GetSystemInfoOk returns a tuple with the SystemInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemInfo

`func (o *ClusterDistributedServiceCardStatus) SetSystemInfo(v ClusterDSCInfo)`

SetSystemInfo sets SystemInfo field to given value.

### HasSystemInfo

`func (o *ClusterDistributedServiceCardStatus) HasSystemInfo() bool`

HasSystemInfo returns a boolean if a field has been set.

### GetUnhealthyServices

`func (o *ClusterDistributedServiceCardStatus) GetUnhealthyServices() []string`

GetUnhealthyServices returns the UnhealthyServices field if non-nil, zero value otherwise.

### GetUnhealthyServicesOk

`func (o *ClusterDistributedServiceCardStatus) GetUnhealthyServicesOk() (*[]string, bool)`

GetUnhealthyServicesOk returns a tuple with the UnhealthyServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnhealthyServices

`func (o *ClusterDistributedServiceCardStatus) SetUnhealthyServices(v []string)`

SetUnhealthyServices sets UnhealthyServices field to given value.

### HasUnhealthyServices

`func (o *ClusterDistributedServiceCardStatus) HasUnhealthyServices() bool`

HasUnhealthyServices returns a boolean if a field has been set.

### GetVersionMismatch

`func (o *ClusterDistributedServiceCardStatus) GetVersionMismatch() bool`

GetVersionMismatch returns the VersionMismatch field if non-nil, zero value otherwise.

### GetVersionMismatchOk

`func (o *ClusterDistributedServiceCardStatus) GetVersionMismatchOk() (*bool, bool)`

GetVersionMismatchOk returns a tuple with the VersionMismatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionMismatch

`func (o *ClusterDistributedServiceCardStatus) SetVersionMismatch(v bool)`

SetVersionMismatch sets VersionMismatch field to given value.

### HasVersionMismatch

`func (o *ClusterDistributedServiceCardStatus) HasVersionMismatch() bool`

HasVersionMismatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


