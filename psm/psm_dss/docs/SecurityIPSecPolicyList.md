# SecurityIPSecPolicyList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Items** | Pointer to [**[]SecurityIPSecPolicy**](SecurityIPSecPolicy.md) | List of IPSecPolicy objects. | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**ListMeta** | Pointer to [**ApiListMeta**](apiListMeta.md) |  | [optional] 

## Methods

### NewSecurityIPSecPolicyList

`func NewSecurityIPSecPolicyList() *SecurityIPSecPolicyList`

NewSecurityIPSecPolicyList instantiates a new SecurityIPSecPolicyList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityIPSecPolicyListWithDefaults

`func NewSecurityIPSecPolicyListWithDefaults() *SecurityIPSecPolicyList`

NewSecurityIPSecPolicyListWithDefaults instantiates a new SecurityIPSecPolicyList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *SecurityIPSecPolicyList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *SecurityIPSecPolicyList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *SecurityIPSecPolicyList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *SecurityIPSecPolicyList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *SecurityIPSecPolicyList) GetItems() []SecurityIPSecPolicy`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SecurityIPSecPolicyList) GetItemsOk() (*[]SecurityIPSecPolicy, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SecurityIPSecPolicyList) SetItems(v []SecurityIPSecPolicy)`

SetItems sets Items field to given value.

### HasItems

`func (o *SecurityIPSecPolicyList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetKind

`func (o *SecurityIPSecPolicyList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SecurityIPSecPolicyList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SecurityIPSecPolicyList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SecurityIPSecPolicyList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetListMeta

`func (o *SecurityIPSecPolicyList) GetListMeta() ApiListMeta`

GetListMeta returns the ListMeta field if non-nil, zero value otherwise.

### GetListMetaOk

`func (o *SecurityIPSecPolicyList) GetListMetaOk() (*ApiListMeta, bool)`

GetListMetaOk returns a tuple with the ListMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListMeta

`func (o *SecurityIPSecPolicyList) SetListMeta(v ApiListMeta)`

SetListMeta sets ListMeta field to given value.

### HasListMeta

`func (o *SecurityIPSecPolicyList) HasListMeta() bool`

HasListMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


