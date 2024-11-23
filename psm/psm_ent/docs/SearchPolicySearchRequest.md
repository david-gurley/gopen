# SearchPolicySearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Action, e.g. PERMIT, DENY or REJECT/CLEAR, PROTECT_PERMISSIVE or PROTECT_STRICT. | [optional] 
**App** | Pointer to **string** | App specification,  predefined apps and alg config. | [optional] 
**FromIpAddress** | Pointer to **string** | Inbound ip-address, use any to refer to all ipaddresses eg: 10.1.1.1, any. | [optional] 
**FromSecurityGroup** | Pointer to **string** | Inbound security group. | [optional] 
**Kinds** | Pointer to **[]string** | Kind of the policy that this request should act on. It should be either NetworkSecurityPolicy or IPSecPolicy. | [optional] 
**Name** | Pointer to **string** | Name is optional. If provided policy-search will be limited to the specified policy of the given name on the given kind. If empty, then all the policies of the given kind will be searched. | [optional] 
**Namespace** | Pointer to **string** | Namespace is optional. If provided policy-search will be limited to the specified namespace. | [optional] [default to "default"]
**Port** | Pointer to **string** | TCP or UDP Port number. | [optional] 
**Protocol** | Pointer to **string** | Protocol eg: tcp, udp, icmp. | [optional] 
**Tenant** | Pointer to **string** | Tenant Name, to perform query within a Tenant&#39;s scope. The default tenant is \&quot;default\&quot;. In the backend this field gets auto-filled &amp; validated by apigw-hook based on user login context. | [optional] [default to "default"]
**ToIpAddress** | Pointer to **string** | Outbound ip-address, use any to refer to all ipaddresses eg: 20.1.1.1, any. | [optional] 
**ToSecurityGroup** | Pointer to **string** | Outbound security group. | [optional] 

## Methods

### NewSearchPolicySearchRequest

`func NewSearchPolicySearchRequest() *SearchPolicySearchRequest`

NewSearchPolicySearchRequest instantiates a new SearchPolicySearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchPolicySearchRequestWithDefaults

`func NewSearchPolicySearchRequestWithDefaults() *SearchPolicySearchRequest`

NewSearchPolicySearchRequestWithDefaults instantiates a new SearchPolicySearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *SearchPolicySearchRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SearchPolicySearchRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SearchPolicySearchRequest) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *SearchPolicySearchRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetApp

`func (o *SearchPolicySearchRequest) GetApp() string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *SearchPolicySearchRequest) GetAppOk() (*string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *SearchPolicySearchRequest) SetApp(v string)`

SetApp sets App field to given value.

### HasApp

`func (o *SearchPolicySearchRequest) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetFromIpAddress

`func (o *SearchPolicySearchRequest) GetFromIpAddress() string`

GetFromIpAddress returns the FromIpAddress field if non-nil, zero value otherwise.

### GetFromIpAddressOk

`func (o *SearchPolicySearchRequest) GetFromIpAddressOk() (*string, bool)`

GetFromIpAddressOk returns a tuple with the FromIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromIpAddress

`func (o *SearchPolicySearchRequest) SetFromIpAddress(v string)`

SetFromIpAddress sets FromIpAddress field to given value.

### HasFromIpAddress

`func (o *SearchPolicySearchRequest) HasFromIpAddress() bool`

HasFromIpAddress returns a boolean if a field has been set.

### GetFromSecurityGroup

`func (o *SearchPolicySearchRequest) GetFromSecurityGroup() string`

GetFromSecurityGroup returns the FromSecurityGroup field if non-nil, zero value otherwise.

### GetFromSecurityGroupOk

`func (o *SearchPolicySearchRequest) GetFromSecurityGroupOk() (*string, bool)`

GetFromSecurityGroupOk returns a tuple with the FromSecurityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromSecurityGroup

`func (o *SearchPolicySearchRequest) SetFromSecurityGroup(v string)`

SetFromSecurityGroup sets FromSecurityGroup field to given value.

### HasFromSecurityGroup

`func (o *SearchPolicySearchRequest) HasFromSecurityGroup() bool`

HasFromSecurityGroup returns a boolean if a field has been set.

### GetKinds

`func (o *SearchPolicySearchRequest) GetKinds() []string`

GetKinds returns the Kinds field if non-nil, zero value otherwise.

### GetKindsOk

`func (o *SearchPolicySearchRequest) GetKindsOk() (*[]string, bool)`

GetKindsOk returns a tuple with the Kinds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKinds

`func (o *SearchPolicySearchRequest) SetKinds(v []string)`

SetKinds sets Kinds field to given value.

### HasKinds

`func (o *SearchPolicySearchRequest) HasKinds() bool`

HasKinds returns a boolean if a field has been set.

### GetName

`func (o *SearchPolicySearchRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SearchPolicySearchRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SearchPolicySearchRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SearchPolicySearchRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *SearchPolicySearchRequest) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *SearchPolicySearchRequest) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *SearchPolicySearchRequest) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *SearchPolicySearchRequest) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetPort

`func (o *SearchPolicySearchRequest) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SearchPolicySearchRequest) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SearchPolicySearchRequest) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *SearchPolicySearchRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProtocol

`func (o *SearchPolicySearchRequest) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *SearchPolicySearchRequest) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *SearchPolicySearchRequest) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *SearchPolicySearchRequest) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetTenant

`func (o *SearchPolicySearchRequest) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *SearchPolicySearchRequest) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *SearchPolicySearchRequest) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *SearchPolicySearchRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetToIpAddress

`func (o *SearchPolicySearchRequest) GetToIpAddress() string`

GetToIpAddress returns the ToIpAddress field if non-nil, zero value otherwise.

### GetToIpAddressOk

`func (o *SearchPolicySearchRequest) GetToIpAddressOk() (*string, bool)`

GetToIpAddressOk returns a tuple with the ToIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToIpAddress

`func (o *SearchPolicySearchRequest) SetToIpAddress(v string)`

SetToIpAddress sets ToIpAddress field to given value.

### HasToIpAddress

`func (o *SearchPolicySearchRequest) HasToIpAddress() bool`

HasToIpAddress returns a boolean if a field has been set.

### GetToSecurityGroup

`func (o *SearchPolicySearchRequest) GetToSecurityGroup() string`

GetToSecurityGroup returns the ToSecurityGroup field if non-nil, zero value otherwise.

### GetToSecurityGroupOk

`func (o *SearchPolicySearchRequest) GetToSecurityGroupOk() (*string, bool)`

GetToSecurityGroupOk returns a tuple with the ToSecurityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToSecurityGroup

`func (o *SearchPolicySearchRequest) SetToSecurityGroup(v string)`

SetToSecurityGroup sets ToSecurityGroup field to given value.

### HasToSecurityGroup

`func (o *SearchPolicySearchRequest) HasToSecurityGroup() bool`

HasToSecurityGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


