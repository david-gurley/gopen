# ClusterClusterProfileSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OverlayForwarding** | Pointer to [**ClusterOverlayForwarding**](clusterOverlayForwarding.md) |  | [optional] 
**SearchOptions** | Pointer to [**ClusterSearchOptions**](clusterSearchOptions.md) |  | [optional] 
**SecurityPolicyRuleScaleTemplate** | Pointer to **string** | SecurityPolicyRuleScaleTemplate provides options to configure security policy and rule scale for the cluster. | [optional] [default to "NONE"]

## Methods

### NewClusterClusterProfileSpec

`func NewClusterClusterProfileSpec() *ClusterClusterProfileSpec`

NewClusterClusterProfileSpec instantiates a new ClusterClusterProfileSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterClusterProfileSpecWithDefaults

`func NewClusterClusterProfileSpecWithDefaults() *ClusterClusterProfileSpec`

NewClusterClusterProfileSpecWithDefaults instantiates a new ClusterClusterProfileSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverlayForwarding

`func (o *ClusterClusterProfileSpec) GetOverlayForwarding() ClusterOverlayForwarding`

GetOverlayForwarding returns the OverlayForwarding field if non-nil, zero value otherwise.

### GetOverlayForwardingOk

`func (o *ClusterClusterProfileSpec) GetOverlayForwardingOk() (*ClusterOverlayForwarding, bool)`

GetOverlayForwardingOk returns a tuple with the OverlayForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverlayForwarding

`func (o *ClusterClusterProfileSpec) SetOverlayForwarding(v ClusterOverlayForwarding)`

SetOverlayForwarding sets OverlayForwarding field to given value.

### HasOverlayForwarding

`func (o *ClusterClusterProfileSpec) HasOverlayForwarding() bool`

HasOverlayForwarding returns a boolean if a field has been set.

### GetSearchOptions

`func (o *ClusterClusterProfileSpec) GetSearchOptions() ClusterSearchOptions`

GetSearchOptions returns the SearchOptions field if non-nil, zero value otherwise.

### GetSearchOptionsOk

`func (o *ClusterClusterProfileSpec) GetSearchOptionsOk() (*ClusterSearchOptions, bool)`

GetSearchOptionsOk returns a tuple with the SearchOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchOptions

`func (o *ClusterClusterProfileSpec) SetSearchOptions(v ClusterSearchOptions)`

SetSearchOptions sets SearchOptions field to given value.

### HasSearchOptions

`func (o *ClusterClusterProfileSpec) HasSearchOptions() bool`

HasSearchOptions returns a boolean if a field has been set.

### GetSecurityPolicyRuleScaleTemplate

`func (o *ClusterClusterProfileSpec) GetSecurityPolicyRuleScaleTemplate() string`

GetSecurityPolicyRuleScaleTemplate returns the SecurityPolicyRuleScaleTemplate field if non-nil, zero value otherwise.

### GetSecurityPolicyRuleScaleTemplateOk

`func (o *ClusterClusterProfileSpec) GetSecurityPolicyRuleScaleTemplateOk() (*string, bool)`

GetSecurityPolicyRuleScaleTemplateOk returns a tuple with the SecurityPolicyRuleScaleTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityPolicyRuleScaleTemplate

`func (o *ClusterClusterProfileSpec) SetSecurityPolicyRuleScaleTemplate(v string)`

SetSecurityPolicyRuleScaleTemplate sets SecurityPolicyRuleScaleTemplate field to given value.

### HasSecurityPolicyRuleScaleTemplate

`func (o *ClusterClusterProfileSpec) HasSecurityPolicyRuleScaleTemplate() bool`

HasSecurityPolicyRuleScaleTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


