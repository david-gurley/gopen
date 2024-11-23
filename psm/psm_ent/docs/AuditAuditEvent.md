# AuditAuditEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**ApiVersion** | Pointer to **string** |  | [optional] 
**ClientIps** | Pointer to **[]string** |  | [optional] 
**Data** | Pointer to **map[string]string** |  | [optional] 
**ExternalId** | Pointer to **string** | Length of string should be between 0 and 64. Must be alpha numeric and can have -. | [optional] 
**GatewayIp** | Pointer to **string** |  | [optional] 
**GatewayNode** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Level** | Pointer to **string** |  | [optional] [default to "basic"]
**Meta** | Pointer to [**ApiObjectMeta**](apiObjectMeta.md) |  | [optional] 
**Outcome** | Pointer to **string** |  | [optional] [default to "success"]
**RequestObject** | Pointer to **string** |  | [optional] 
**RequestUri** | Pointer to **string** | Should be a valid URI. | [optional] 
**Resource** | Pointer to [**ApiObjectRef**](apiObjectRef.md) |  | [optional] 
**ResponseObject** | Pointer to **string** |  | [optional] 
**ServiceName** | Pointer to **string** |  | [optional] 
**Stage** | Pointer to **string** |  | [optional] [default to "requestauthorization"]
**User** | Pointer to [**ApiObjectRef**](apiObjectRef.md) |  | [optional] 

## Methods

### NewAuditAuditEvent

`func NewAuditAuditEvent() *AuditAuditEvent`

NewAuditAuditEvent instantiates a new AuditAuditEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditAuditEventWithDefaults

`func NewAuditAuditEventWithDefaults() *AuditAuditEvent`

NewAuditAuditEventWithDefaults instantiates a new AuditAuditEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *AuditAuditEvent) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AuditAuditEvent) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AuditAuditEvent) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *AuditAuditEvent) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetApiVersion

`func (o *AuditAuditEvent) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AuditAuditEvent) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AuditAuditEvent) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *AuditAuditEvent) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetClientIps

`func (o *AuditAuditEvent) GetClientIps() []string`

GetClientIps returns the ClientIps field if non-nil, zero value otherwise.

### GetClientIpsOk

`func (o *AuditAuditEvent) GetClientIpsOk() (*[]string, bool)`

GetClientIpsOk returns a tuple with the ClientIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIps

`func (o *AuditAuditEvent) SetClientIps(v []string)`

SetClientIps sets ClientIps field to given value.

### HasClientIps

`func (o *AuditAuditEvent) HasClientIps() bool`

HasClientIps returns a boolean if a field has been set.

### GetData

`func (o *AuditAuditEvent) GetData() map[string]string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AuditAuditEvent) GetDataOk() (*map[string]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AuditAuditEvent) SetData(v map[string]string)`

SetData sets Data field to given value.

### HasData

`func (o *AuditAuditEvent) HasData() bool`

HasData returns a boolean if a field has been set.

### GetExternalId

`func (o *AuditAuditEvent) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AuditAuditEvent) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AuditAuditEvent) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *AuditAuditEvent) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetGatewayIp

`func (o *AuditAuditEvent) GetGatewayIp() string`

GetGatewayIp returns the GatewayIp field if non-nil, zero value otherwise.

### GetGatewayIpOk

`func (o *AuditAuditEvent) GetGatewayIpOk() (*string, bool)`

GetGatewayIpOk returns a tuple with the GatewayIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIp

`func (o *AuditAuditEvent) SetGatewayIp(v string)`

SetGatewayIp sets GatewayIp field to given value.

### HasGatewayIp

`func (o *AuditAuditEvent) HasGatewayIp() bool`

HasGatewayIp returns a boolean if a field has been set.

### GetGatewayNode

`func (o *AuditAuditEvent) GetGatewayNode() string`

GetGatewayNode returns the GatewayNode field if non-nil, zero value otherwise.

### GetGatewayNodeOk

`func (o *AuditAuditEvent) GetGatewayNodeOk() (*string, bool)`

GetGatewayNodeOk returns a tuple with the GatewayNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayNode

`func (o *AuditAuditEvent) SetGatewayNode(v string)`

SetGatewayNode sets GatewayNode field to given value.

### HasGatewayNode

`func (o *AuditAuditEvent) HasGatewayNode() bool`

HasGatewayNode returns a boolean if a field has been set.

### GetKind

`func (o *AuditAuditEvent) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AuditAuditEvent) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AuditAuditEvent) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *AuditAuditEvent) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLevel

`func (o *AuditAuditEvent) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *AuditAuditEvent) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *AuditAuditEvent) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *AuditAuditEvent) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetMeta

`func (o *AuditAuditEvent) GetMeta() ApiObjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AuditAuditEvent) GetMetaOk() (*ApiObjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AuditAuditEvent) SetMeta(v ApiObjectMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AuditAuditEvent) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetOutcome

`func (o *AuditAuditEvent) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *AuditAuditEvent) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *AuditAuditEvent) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *AuditAuditEvent) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### GetRequestObject

`func (o *AuditAuditEvent) GetRequestObject() string`

GetRequestObject returns the RequestObject field if non-nil, zero value otherwise.

### GetRequestObjectOk

`func (o *AuditAuditEvent) GetRequestObjectOk() (*string, bool)`

GetRequestObjectOk returns a tuple with the RequestObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestObject

`func (o *AuditAuditEvent) SetRequestObject(v string)`

SetRequestObject sets RequestObject field to given value.

### HasRequestObject

`func (o *AuditAuditEvent) HasRequestObject() bool`

HasRequestObject returns a boolean if a field has been set.

### GetRequestUri

`func (o *AuditAuditEvent) GetRequestUri() string`

GetRequestUri returns the RequestUri field if non-nil, zero value otherwise.

### GetRequestUriOk

`func (o *AuditAuditEvent) GetRequestUriOk() (*string, bool)`

GetRequestUriOk returns a tuple with the RequestUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestUri

`func (o *AuditAuditEvent) SetRequestUri(v string)`

SetRequestUri sets RequestUri field to given value.

### HasRequestUri

`func (o *AuditAuditEvent) HasRequestUri() bool`

HasRequestUri returns a boolean if a field has been set.

### GetResource

`func (o *AuditAuditEvent) GetResource() ApiObjectRef`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *AuditAuditEvent) GetResourceOk() (*ApiObjectRef, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *AuditAuditEvent) SetResource(v ApiObjectRef)`

SetResource sets Resource field to given value.

### HasResource

`func (o *AuditAuditEvent) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetResponseObject

`func (o *AuditAuditEvent) GetResponseObject() string`

GetResponseObject returns the ResponseObject field if non-nil, zero value otherwise.

### GetResponseObjectOk

`func (o *AuditAuditEvent) GetResponseObjectOk() (*string, bool)`

GetResponseObjectOk returns a tuple with the ResponseObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseObject

`func (o *AuditAuditEvent) SetResponseObject(v string)`

SetResponseObject sets ResponseObject field to given value.

### HasResponseObject

`func (o *AuditAuditEvent) HasResponseObject() bool`

HasResponseObject returns a boolean if a field has been set.

### GetServiceName

`func (o *AuditAuditEvent) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *AuditAuditEvent) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *AuditAuditEvent) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *AuditAuditEvent) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetStage

`func (o *AuditAuditEvent) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *AuditAuditEvent) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *AuditAuditEvent) SetStage(v string)`

SetStage sets Stage field to given value.

### HasStage

`func (o *AuditAuditEvent) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetUser

`func (o *AuditAuditEvent) GetUser() ApiObjectRef`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AuditAuditEvent) GetUserOk() (*ApiObjectRef, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AuditAuditEvent) SetUser(v ApiObjectRef)`

SetUser sets User field to given value.

### HasUser

`func (o *AuditAuditEvent) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


