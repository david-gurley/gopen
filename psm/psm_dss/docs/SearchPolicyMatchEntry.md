# SearchPolicyMatchEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **int64** | Index of the matching SG rule. | [optional] 
**IpSecPolicyRule** | Pointer to [**SecurityIPSecPolicyRule**](securityIPSecPolicyRule.md) |  | [optional] 
**SgRule** | Pointer to [**SecuritySGRule**](securitySGRule.md) |  | [optional] 

## Methods

### NewSearchPolicyMatchEntry

`func NewSearchPolicyMatchEntry() *SearchPolicyMatchEntry`

NewSearchPolicyMatchEntry instantiates a new SearchPolicyMatchEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchPolicyMatchEntryWithDefaults

`func NewSearchPolicyMatchEntryWithDefaults() *SearchPolicyMatchEntry`

NewSearchPolicyMatchEntryWithDefaults instantiates a new SearchPolicyMatchEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *SearchPolicyMatchEntry) GetIndex() int64`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *SearchPolicyMatchEntry) GetIndexOk() (*int64, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *SearchPolicyMatchEntry) SetIndex(v int64)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *SearchPolicyMatchEntry) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetIpSecPolicyRule

`func (o *SearchPolicyMatchEntry) GetIpSecPolicyRule() SecurityIPSecPolicyRule`

GetIpSecPolicyRule returns the IpSecPolicyRule field if non-nil, zero value otherwise.

### GetIpSecPolicyRuleOk

`func (o *SearchPolicyMatchEntry) GetIpSecPolicyRuleOk() (*SecurityIPSecPolicyRule, bool)`

GetIpSecPolicyRuleOk returns a tuple with the IpSecPolicyRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSecPolicyRule

`func (o *SearchPolicyMatchEntry) SetIpSecPolicyRule(v SecurityIPSecPolicyRule)`

SetIpSecPolicyRule sets IpSecPolicyRule field to given value.

### HasIpSecPolicyRule

`func (o *SearchPolicyMatchEntry) HasIpSecPolicyRule() bool`

HasIpSecPolicyRule returns a boolean if a field has been set.

### GetSgRule

`func (o *SearchPolicyMatchEntry) GetSgRule() SecuritySGRule`

GetSgRule returns the SgRule field if non-nil, zero value otherwise.

### GetSgRuleOk

`func (o *SearchPolicyMatchEntry) GetSgRuleOk() (*SecuritySGRule, bool)`

GetSgRuleOk returns a tuple with the SgRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSgRule

`func (o *SearchPolicyMatchEntry) SetSgRule(v SecuritySGRule)`

SetSgRule sets SgRule field to given value.

### HasSgRule

`func (o *SearchPolicyMatchEntry) HasSgRule() bool`

HasSgRule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


