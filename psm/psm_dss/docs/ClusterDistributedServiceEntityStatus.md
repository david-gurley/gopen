# ClusterDistributedServiceEntityStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DSCSku** | Pointer to **string** |  | [optional] 
**DSCVersion** | Pointer to **string** |  | [optional] 
**AdmPhaseReason** | Pointer to **string** |  | [optional] 
**AdmissionPhase** | Pointer to **string** |  | [optional] [default to "unknown"]
**AlomPresent** | Pointer to **bool** |  | [optional] 
**Conditions** | Pointer to [**[]ClusterDSCCondition**](ClusterDSCCondition.md) |  | [optional] 
**ControlPlaneStatus** | Pointer to [**ClusterDSCControlPlaneStatus**](clusterDSCControlPlaneStatus.md) |  | [optional] 
**DssInfo** | Pointer to [**ClusterDSSInfo**](clusterDSSInfo.md) |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**InbandIpConfig** | Pointer to [**ClusterIPConfig**](clusterIPConfig.md) |  | [optional] 
**Interfaces** | Pointer to **[]string** |  | [optional] 
**IpConfig** | Pointer to [**ClusterIPConfig**](clusterIPConfig.md) |  | [optional] 
**IsConnectedToPsm** | Pointer to **bool** |  | [optional] 
**NumMacAddress** | Pointer to **int64** | Value should be between 0 and 256. | [optional] [default to 24]
**PackageType** | Pointer to **string** |  | [optional] [default to "dsc"]
**PrimaryMac** | Pointer to **string** | Should be a valid MAC address. | [optional] 
**SecureBooted** | Pointer to **bool** |  | [optional] 
**SecurityPolicyRuleScaleProfile** | Pointer to **string** |  | [optional] 
**SerialNum** | Pointer to **string** |  | [optional] 
**SystemInfo** | Pointer to [**ClusterDSCInfo**](clusterDSCInfo.md) |  | [optional] 
**UnhealthyServices** | Pointer to **[]string** |  | [optional] 
**VersionMismatch** | Pointer to **bool** |  | [optional] 

## Methods

### NewClusterDistributedServiceEntityStatus

`func NewClusterDistributedServiceEntityStatus() *ClusterDistributedServiceEntityStatus`

NewClusterDistributedServiceEntityStatus instantiates a new ClusterDistributedServiceEntityStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterDistributedServiceEntityStatusWithDefaults

`func NewClusterDistributedServiceEntityStatusWithDefaults() *ClusterDistributedServiceEntityStatus`

NewClusterDistributedServiceEntityStatusWithDefaults instantiates a new ClusterDistributedServiceEntityStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDSCSku

`func (o *ClusterDistributedServiceEntityStatus) GetDSCSku() string`

GetDSCSku returns the DSCSku field if non-nil, zero value otherwise.

### GetDSCSkuOk

`func (o *ClusterDistributedServiceEntityStatus) GetDSCSkuOk() (*string, bool)`

GetDSCSkuOk returns a tuple with the DSCSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDSCSku

`func (o *ClusterDistributedServiceEntityStatus) SetDSCSku(v string)`

SetDSCSku sets DSCSku field to given value.

### HasDSCSku

`func (o *ClusterDistributedServiceEntityStatus) HasDSCSku() bool`

HasDSCSku returns a boolean if a field has been set.

### GetDSCVersion

`func (o *ClusterDistributedServiceEntityStatus) GetDSCVersion() string`

GetDSCVersion returns the DSCVersion field if non-nil, zero value otherwise.

### GetDSCVersionOk

`func (o *ClusterDistributedServiceEntityStatus) GetDSCVersionOk() (*string, bool)`

GetDSCVersionOk returns a tuple with the DSCVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDSCVersion

`func (o *ClusterDistributedServiceEntityStatus) SetDSCVersion(v string)`

SetDSCVersion sets DSCVersion field to given value.

### HasDSCVersion

`func (o *ClusterDistributedServiceEntityStatus) HasDSCVersion() bool`

HasDSCVersion returns a boolean if a field has been set.

### GetAdmPhaseReason

`func (o *ClusterDistributedServiceEntityStatus) GetAdmPhaseReason() string`

GetAdmPhaseReason returns the AdmPhaseReason field if non-nil, zero value otherwise.

### GetAdmPhaseReasonOk

`func (o *ClusterDistributedServiceEntityStatus) GetAdmPhaseReasonOk() (*string, bool)`

GetAdmPhaseReasonOk returns a tuple with the AdmPhaseReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmPhaseReason

`func (o *ClusterDistributedServiceEntityStatus) SetAdmPhaseReason(v string)`

SetAdmPhaseReason sets AdmPhaseReason field to given value.

### HasAdmPhaseReason

`func (o *ClusterDistributedServiceEntityStatus) HasAdmPhaseReason() bool`

HasAdmPhaseReason returns a boolean if a field has been set.

### GetAdmissionPhase

`func (o *ClusterDistributedServiceEntityStatus) GetAdmissionPhase() string`

GetAdmissionPhase returns the AdmissionPhase field if non-nil, zero value otherwise.

### GetAdmissionPhaseOk

`func (o *ClusterDistributedServiceEntityStatus) GetAdmissionPhaseOk() (*string, bool)`

GetAdmissionPhaseOk returns a tuple with the AdmissionPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmissionPhase

`func (o *ClusterDistributedServiceEntityStatus) SetAdmissionPhase(v string)`

SetAdmissionPhase sets AdmissionPhase field to given value.

### HasAdmissionPhase

`func (o *ClusterDistributedServiceEntityStatus) HasAdmissionPhase() bool`

HasAdmissionPhase returns a boolean if a field has been set.

### GetAlomPresent

`func (o *ClusterDistributedServiceEntityStatus) GetAlomPresent() bool`

GetAlomPresent returns the AlomPresent field if non-nil, zero value otherwise.

### GetAlomPresentOk

`func (o *ClusterDistributedServiceEntityStatus) GetAlomPresentOk() (*bool, bool)`

GetAlomPresentOk returns a tuple with the AlomPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlomPresent

`func (o *ClusterDistributedServiceEntityStatus) SetAlomPresent(v bool)`

SetAlomPresent sets AlomPresent field to given value.

### HasAlomPresent

`func (o *ClusterDistributedServiceEntityStatus) HasAlomPresent() bool`

HasAlomPresent returns a boolean if a field has been set.

### GetConditions

`func (o *ClusterDistributedServiceEntityStatus) GetConditions() []ClusterDSCCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *ClusterDistributedServiceEntityStatus) GetConditionsOk() (*[]ClusterDSCCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *ClusterDistributedServiceEntityStatus) SetConditions(v []ClusterDSCCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *ClusterDistributedServiceEntityStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetControlPlaneStatus

`func (o *ClusterDistributedServiceEntityStatus) GetControlPlaneStatus() ClusterDSCControlPlaneStatus`

GetControlPlaneStatus returns the ControlPlaneStatus field if non-nil, zero value otherwise.

### GetControlPlaneStatusOk

`func (o *ClusterDistributedServiceEntityStatus) GetControlPlaneStatusOk() (*ClusterDSCControlPlaneStatus, bool)`

GetControlPlaneStatusOk returns a tuple with the ControlPlaneStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlPlaneStatus

`func (o *ClusterDistributedServiceEntityStatus) SetControlPlaneStatus(v ClusterDSCControlPlaneStatus)`

SetControlPlaneStatus sets ControlPlaneStatus field to given value.

### HasControlPlaneStatus

`func (o *ClusterDistributedServiceEntityStatus) HasControlPlaneStatus() bool`

HasControlPlaneStatus returns a boolean if a field has been set.

### GetDssInfo

`func (o *ClusterDistributedServiceEntityStatus) GetDssInfo() ClusterDSSInfo`

GetDssInfo returns the DssInfo field if non-nil, zero value otherwise.

### GetDssInfoOk

`func (o *ClusterDistributedServiceEntityStatus) GetDssInfoOk() (*ClusterDSSInfo, bool)`

GetDssInfoOk returns a tuple with the DssInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDssInfo

`func (o *ClusterDistributedServiceEntityStatus) SetDssInfo(v ClusterDSSInfo)`

SetDssInfo sets DssInfo field to given value.

### HasDssInfo

`func (o *ClusterDistributedServiceEntityStatus) HasDssInfo() bool`

HasDssInfo returns a boolean if a field has been set.

### GetHost

`func (o *ClusterDistributedServiceEntityStatus) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ClusterDistributedServiceEntityStatus) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ClusterDistributedServiceEntityStatus) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *ClusterDistributedServiceEntityStatus) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetInbandIpConfig

`func (o *ClusterDistributedServiceEntityStatus) GetInbandIpConfig() ClusterIPConfig`

GetInbandIpConfig returns the InbandIpConfig field if non-nil, zero value otherwise.

### GetInbandIpConfigOk

`func (o *ClusterDistributedServiceEntityStatus) GetInbandIpConfigOk() (*ClusterIPConfig, bool)`

GetInbandIpConfigOk returns a tuple with the InbandIpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbandIpConfig

`func (o *ClusterDistributedServiceEntityStatus) SetInbandIpConfig(v ClusterIPConfig)`

SetInbandIpConfig sets InbandIpConfig field to given value.

### HasInbandIpConfig

`func (o *ClusterDistributedServiceEntityStatus) HasInbandIpConfig() bool`

HasInbandIpConfig returns a boolean if a field has been set.

### GetInterfaces

`func (o *ClusterDistributedServiceEntityStatus) GetInterfaces() []string`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *ClusterDistributedServiceEntityStatus) GetInterfacesOk() (*[]string, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *ClusterDistributedServiceEntityStatus) SetInterfaces(v []string)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *ClusterDistributedServiceEntityStatus) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetIpConfig

`func (o *ClusterDistributedServiceEntityStatus) GetIpConfig() ClusterIPConfig`

GetIpConfig returns the IpConfig field if non-nil, zero value otherwise.

### GetIpConfigOk

`func (o *ClusterDistributedServiceEntityStatus) GetIpConfigOk() (*ClusterIPConfig, bool)`

GetIpConfigOk returns a tuple with the IpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpConfig

`func (o *ClusterDistributedServiceEntityStatus) SetIpConfig(v ClusterIPConfig)`

SetIpConfig sets IpConfig field to given value.

### HasIpConfig

`func (o *ClusterDistributedServiceEntityStatus) HasIpConfig() bool`

HasIpConfig returns a boolean if a field has been set.

### GetIsConnectedToPsm

`func (o *ClusterDistributedServiceEntityStatus) GetIsConnectedToPsm() bool`

GetIsConnectedToPsm returns the IsConnectedToPsm field if non-nil, zero value otherwise.

### GetIsConnectedToPsmOk

`func (o *ClusterDistributedServiceEntityStatus) GetIsConnectedToPsmOk() (*bool, bool)`

GetIsConnectedToPsmOk returns a tuple with the IsConnectedToPsm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConnectedToPsm

`func (o *ClusterDistributedServiceEntityStatus) SetIsConnectedToPsm(v bool)`

SetIsConnectedToPsm sets IsConnectedToPsm field to given value.

### HasIsConnectedToPsm

`func (o *ClusterDistributedServiceEntityStatus) HasIsConnectedToPsm() bool`

HasIsConnectedToPsm returns a boolean if a field has been set.

### GetNumMacAddress

`func (o *ClusterDistributedServiceEntityStatus) GetNumMacAddress() int64`

GetNumMacAddress returns the NumMacAddress field if non-nil, zero value otherwise.

### GetNumMacAddressOk

`func (o *ClusterDistributedServiceEntityStatus) GetNumMacAddressOk() (*int64, bool)`

GetNumMacAddressOk returns a tuple with the NumMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumMacAddress

`func (o *ClusterDistributedServiceEntityStatus) SetNumMacAddress(v int64)`

SetNumMacAddress sets NumMacAddress field to given value.

### HasNumMacAddress

`func (o *ClusterDistributedServiceEntityStatus) HasNumMacAddress() bool`

HasNumMacAddress returns a boolean if a field has been set.

### GetPackageType

`func (o *ClusterDistributedServiceEntityStatus) GetPackageType() string`

GetPackageType returns the PackageType field if non-nil, zero value otherwise.

### GetPackageTypeOk

`func (o *ClusterDistributedServiceEntityStatus) GetPackageTypeOk() (*string, bool)`

GetPackageTypeOk returns a tuple with the PackageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageType

`func (o *ClusterDistributedServiceEntityStatus) SetPackageType(v string)`

SetPackageType sets PackageType field to given value.

### HasPackageType

`func (o *ClusterDistributedServiceEntityStatus) HasPackageType() bool`

HasPackageType returns a boolean if a field has been set.

### GetPrimaryMac

`func (o *ClusterDistributedServiceEntityStatus) GetPrimaryMac() string`

GetPrimaryMac returns the PrimaryMac field if non-nil, zero value otherwise.

### GetPrimaryMacOk

`func (o *ClusterDistributedServiceEntityStatus) GetPrimaryMacOk() (*string, bool)`

GetPrimaryMacOk returns a tuple with the PrimaryMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryMac

`func (o *ClusterDistributedServiceEntityStatus) SetPrimaryMac(v string)`

SetPrimaryMac sets PrimaryMac field to given value.

### HasPrimaryMac

`func (o *ClusterDistributedServiceEntityStatus) HasPrimaryMac() bool`

HasPrimaryMac returns a boolean if a field has been set.

### GetSecureBooted

`func (o *ClusterDistributedServiceEntityStatus) GetSecureBooted() bool`

GetSecureBooted returns the SecureBooted field if non-nil, zero value otherwise.

### GetSecureBootedOk

`func (o *ClusterDistributedServiceEntityStatus) GetSecureBootedOk() (*bool, bool)`

GetSecureBootedOk returns a tuple with the SecureBooted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureBooted

`func (o *ClusterDistributedServiceEntityStatus) SetSecureBooted(v bool)`

SetSecureBooted sets SecureBooted field to given value.

### HasSecureBooted

`func (o *ClusterDistributedServiceEntityStatus) HasSecureBooted() bool`

HasSecureBooted returns a boolean if a field has been set.

### GetSecurityPolicyRuleScaleProfile

`func (o *ClusterDistributedServiceEntityStatus) GetSecurityPolicyRuleScaleProfile() string`

GetSecurityPolicyRuleScaleProfile returns the SecurityPolicyRuleScaleProfile field if non-nil, zero value otherwise.

### GetSecurityPolicyRuleScaleProfileOk

`func (o *ClusterDistributedServiceEntityStatus) GetSecurityPolicyRuleScaleProfileOk() (*string, bool)`

GetSecurityPolicyRuleScaleProfileOk returns a tuple with the SecurityPolicyRuleScaleProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityPolicyRuleScaleProfile

`func (o *ClusterDistributedServiceEntityStatus) SetSecurityPolicyRuleScaleProfile(v string)`

SetSecurityPolicyRuleScaleProfile sets SecurityPolicyRuleScaleProfile field to given value.

### HasSecurityPolicyRuleScaleProfile

`func (o *ClusterDistributedServiceEntityStatus) HasSecurityPolicyRuleScaleProfile() bool`

HasSecurityPolicyRuleScaleProfile returns a boolean if a field has been set.

### GetSerialNum

`func (o *ClusterDistributedServiceEntityStatus) GetSerialNum() string`

GetSerialNum returns the SerialNum field if non-nil, zero value otherwise.

### GetSerialNumOk

`func (o *ClusterDistributedServiceEntityStatus) GetSerialNumOk() (*string, bool)`

GetSerialNumOk returns a tuple with the SerialNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNum

`func (o *ClusterDistributedServiceEntityStatus) SetSerialNum(v string)`

SetSerialNum sets SerialNum field to given value.

### HasSerialNum

`func (o *ClusterDistributedServiceEntityStatus) HasSerialNum() bool`

HasSerialNum returns a boolean if a field has been set.

### GetSystemInfo

`func (o *ClusterDistributedServiceEntityStatus) GetSystemInfo() ClusterDSCInfo`

GetSystemInfo returns the SystemInfo field if non-nil, zero value otherwise.

### GetSystemInfoOk

`func (o *ClusterDistributedServiceEntityStatus) GetSystemInfoOk() (*ClusterDSCInfo, bool)`

GetSystemInfoOk returns a tuple with the SystemInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemInfo

`func (o *ClusterDistributedServiceEntityStatus) SetSystemInfo(v ClusterDSCInfo)`

SetSystemInfo sets SystemInfo field to given value.

### HasSystemInfo

`func (o *ClusterDistributedServiceEntityStatus) HasSystemInfo() bool`

HasSystemInfo returns a boolean if a field has been set.

### GetUnhealthyServices

`func (o *ClusterDistributedServiceEntityStatus) GetUnhealthyServices() []string`

GetUnhealthyServices returns the UnhealthyServices field if non-nil, zero value otherwise.

### GetUnhealthyServicesOk

`func (o *ClusterDistributedServiceEntityStatus) GetUnhealthyServicesOk() (*[]string, bool)`

GetUnhealthyServicesOk returns a tuple with the UnhealthyServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnhealthyServices

`func (o *ClusterDistributedServiceEntityStatus) SetUnhealthyServices(v []string)`

SetUnhealthyServices sets UnhealthyServices field to given value.

### HasUnhealthyServices

`func (o *ClusterDistributedServiceEntityStatus) HasUnhealthyServices() bool`

HasUnhealthyServices returns a boolean if a field has been set.

### GetVersionMismatch

`func (o *ClusterDistributedServiceEntityStatus) GetVersionMismatch() bool`

GetVersionMismatch returns the VersionMismatch field if non-nil, zero value otherwise.

### GetVersionMismatchOk

`func (o *ClusterDistributedServiceEntityStatus) GetVersionMismatchOk() (*bool, bool)`

GetVersionMismatchOk returns a tuple with the VersionMismatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionMismatch

`func (o *ClusterDistributedServiceEntityStatus) SetVersionMismatch(v bool)`

SetVersionMismatch sets VersionMismatch field to given value.

### HasVersionMismatch

`func (o *ClusterDistributedServiceEntityStatus) HasVersionMismatch() bool`

HasVersionMismatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


