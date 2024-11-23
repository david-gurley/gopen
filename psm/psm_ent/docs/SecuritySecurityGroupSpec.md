# SecuritySecurityGroupSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MatchPrefixes** | Pointer to **[]string** | list of CIDRs that are part of this security group. | [optional] 
**ServiceLabels** | Pointer to **[]string** | Service object selector. | [optional] 
**WorkloadSelector** | Pointer to [**LabelsSelector**](labelsSelector.md) |  | [optional] 

## Methods

### NewSecuritySecurityGroupSpec

`func NewSecuritySecurityGroupSpec() *SecuritySecurityGroupSpec`

NewSecuritySecurityGroupSpec instantiates a new SecuritySecurityGroupSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecuritySecurityGroupSpecWithDefaults

`func NewSecuritySecurityGroupSpecWithDefaults() *SecuritySecurityGroupSpec`

NewSecuritySecurityGroupSpecWithDefaults instantiates a new SecuritySecurityGroupSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatchPrefixes

`func (o *SecuritySecurityGroupSpec) GetMatchPrefixes() []string`

GetMatchPrefixes returns the MatchPrefixes field if non-nil, zero value otherwise.

### GetMatchPrefixesOk

`func (o *SecuritySecurityGroupSpec) GetMatchPrefixesOk() (*[]string, bool)`

GetMatchPrefixesOk returns a tuple with the MatchPrefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchPrefixes

`func (o *SecuritySecurityGroupSpec) SetMatchPrefixes(v []string)`

SetMatchPrefixes sets MatchPrefixes field to given value.

### HasMatchPrefixes

`func (o *SecuritySecurityGroupSpec) HasMatchPrefixes() bool`

HasMatchPrefixes returns a boolean if a field has been set.

### GetServiceLabels

`func (o *SecuritySecurityGroupSpec) GetServiceLabels() []string`

GetServiceLabels returns the ServiceLabels field if non-nil, zero value otherwise.

### GetServiceLabelsOk

`func (o *SecuritySecurityGroupSpec) GetServiceLabelsOk() (*[]string, bool)`

GetServiceLabelsOk returns a tuple with the ServiceLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLabels

`func (o *SecuritySecurityGroupSpec) SetServiceLabels(v []string)`

SetServiceLabels sets ServiceLabels field to given value.

### HasServiceLabels

`func (o *SecuritySecurityGroupSpec) HasServiceLabels() bool`

HasServiceLabels returns a boolean if a field has been set.

### GetWorkloadSelector

`func (o *SecuritySecurityGroupSpec) GetWorkloadSelector() LabelsSelector`

GetWorkloadSelector returns the WorkloadSelector field if non-nil, zero value otherwise.

### GetWorkloadSelectorOk

`func (o *SecuritySecurityGroupSpec) GetWorkloadSelectorOk() (*LabelsSelector, bool)`

GetWorkloadSelectorOk returns a tuple with the WorkloadSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadSelector

`func (o *SecuritySecurityGroupSpec) SetWorkloadSelector(v LabelsSelector)`

SetWorkloadSelector sets WorkloadSelector field to given value.

### HasWorkloadSelector

`func (o *SecuritySecurityGroupSpec) HasWorkloadSelector() bool`

HasWorkloadSelector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


