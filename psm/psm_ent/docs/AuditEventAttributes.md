# AuditEventAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Action that was requested/performed on the referred object. For non API server resources, it is the http method. | [optional] 
**ClientIps** | Pointer to **[]string** | IP addresses of client and intermediate proxies from where API request was made. | [optional] 
**Data** | Pointer to **map[string]string** | Data is unstructured key value map stored with audit log that may be set by hooks in API Gateway. We can store Signature in JWS compact serialization format in this map. Data in this map will not be signed. | [optional] 
**ExternalId** | Pointer to **string** | ID passed in by an external application to link audit event to the request. It should be AlphaNumeric and can contain -. Maximum length supported is 64. Length of string should be between 0 and 64. Must be alpha numeric and can have -. | [optional] 
**GatewayIp** | Pointer to **string** | IP address of API Gateway where action was observed. | [optional] 
**GatewayNode** | Pointer to **string** | Name of the venice node where action was observed. | [optional] 
**Level** | Pointer to **string** | Level to control amount of audit information logged. | [optional] [default to "basic"]
**Outcome** | Pointer to **string** | Outcome represents the outcome of action on resource. | [optional] [default to "success"]
**RequestObject** | Pointer to **string** | Object from the request in JSON format. | [optional] 
**RequestUri** | Pointer to **string** | RequestURI is the request URI as sent by the client. Should be a valid URI. | [optional] 
**Resource** | Pointer to [**ApiObjectRef**](apiObjectRef.md) |  | [optional] 
**ResponseObject** | Pointer to **string** | Object from the response in JSON format to be sent to the client. | [optional] 
**ServiceName** | Pointer to **string** | Name of service that handled the request and performed the requested operation for ex: search, events etc. | [optional] 
**Stage** | Pointer to **string** | Request handling stage at which audit log was generated. | [optional] [default to "requestauthorization"]
**User** | Pointer to [**ApiObjectRef**](apiObjectRef.md) |  | [optional] 

## Methods

### NewAuditEventAttributes

`func NewAuditEventAttributes() *AuditEventAttributes`

NewAuditEventAttributes instantiates a new AuditEventAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditEventAttributesWithDefaults

`func NewAuditEventAttributesWithDefaults() *AuditEventAttributes`

NewAuditEventAttributesWithDefaults instantiates a new AuditEventAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *AuditEventAttributes) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AuditEventAttributes) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AuditEventAttributes) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *AuditEventAttributes) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetClientIps

`func (o *AuditEventAttributes) GetClientIps() []string`

GetClientIps returns the ClientIps field if non-nil, zero value otherwise.

### GetClientIpsOk

`func (o *AuditEventAttributes) GetClientIpsOk() (*[]string, bool)`

GetClientIpsOk returns a tuple with the ClientIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIps

`func (o *AuditEventAttributes) SetClientIps(v []string)`

SetClientIps sets ClientIps field to given value.

### HasClientIps

`func (o *AuditEventAttributes) HasClientIps() bool`

HasClientIps returns a boolean if a field has been set.

### GetData

`func (o *AuditEventAttributes) GetData() map[string]string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AuditEventAttributes) GetDataOk() (*map[string]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AuditEventAttributes) SetData(v map[string]string)`

SetData sets Data field to given value.

### HasData

`func (o *AuditEventAttributes) HasData() bool`

HasData returns a boolean if a field has been set.

### GetExternalId

`func (o *AuditEventAttributes) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AuditEventAttributes) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AuditEventAttributes) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *AuditEventAttributes) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetGatewayIp

`func (o *AuditEventAttributes) GetGatewayIp() string`

GetGatewayIp returns the GatewayIp field if non-nil, zero value otherwise.

### GetGatewayIpOk

`func (o *AuditEventAttributes) GetGatewayIpOk() (*string, bool)`

GetGatewayIpOk returns a tuple with the GatewayIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIp

`func (o *AuditEventAttributes) SetGatewayIp(v string)`

SetGatewayIp sets GatewayIp field to given value.

### HasGatewayIp

`func (o *AuditEventAttributes) HasGatewayIp() bool`

HasGatewayIp returns a boolean if a field has been set.

### GetGatewayNode

`func (o *AuditEventAttributes) GetGatewayNode() string`

GetGatewayNode returns the GatewayNode field if non-nil, zero value otherwise.

### GetGatewayNodeOk

`func (o *AuditEventAttributes) GetGatewayNodeOk() (*string, bool)`

GetGatewayNodeOk returns a tuple with the GatewayNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayNode

`func (o *AuditEventAttributes) SetGatewayNode(v string)`

SetGatewayNode sets GatewayNode field to given value.

### HasGatewayNode

`func (o *AuditEventAttributes) HasGatewayNode() bool`

HasGatewayNode returns a boolean if a field has been set.

### GetLevel

`func (o *AuditEventAttributes) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *AuditEventAttributes) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *AuditEventAttributes) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *AuditEventAttributes) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetOutcome

`func (o *AuditEventAttributes) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *AuditEventAttributes) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *AuditEventAttributes) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *AuditEventAttributes) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### GetRequestObject

`func (o *AuditEventAttributes) GetRequestObject() string`

GetRequestObject returns the RequestObject field if non-nil, zero value otherwise.

### GetRequestObjectOk

`func (o *AuditEventAttributes) GetRequestObjectOk() (*string, bool)`

GetRequestObjectOk returns a tuple with the RequestObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestObject

`func (o *AuditEventAttributes) SetRequestObject(v string)`

SetRequestObject sets RequestObject field to given value.

### HasRequestObject

`func (o *AuditEventAttributes) HasRequestObject() bool`

HasRequestObject returns a boolean if a field has been set.

### GetRequestUri

`func (o *AuditEventAttributes) GetRequestUri() string`

GetRequestUri returns the RequestUri field if non-nil, zero value otherwise.

### GetRequestUriOk

`func (o *AuditEventAttributes) GetRequestUriOk() (*string, bool)`

GetRequestUriOk returns a tuple with the RequestUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestUri

`func (o *AuditEventAttributes) SetRequestUri(v string)`

SetRequestUri sets RequestUri field to given value.

### HasRequestUri

`func (o *AuditEventAttributes) HasRequestUri() bool`

HasRequestUri returns a boolean if a field has been set.

### GetResource

`func (o *AuditEventAttributes) GetResource() ApiObjectRef`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *AuditEventAttributes) GetResourceOk() (*ApiObjectRef, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *AuditEventAttributes) SetResource(v ApiObjectRef)`

SetResource sets Resource field to given value.

### HasResource

`func (o *AuditEventAttributes) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetResponseObject

`func (o *AuditEventAttributes) GetResponseObject() string`

GetResponseObject returns the ResponseObject field if non-nil, zero value otherwise.

### GetResponseObjectOk

`func (o *AuditEventAttributes) GetResponseObjectOk() (*string, bool)`

GetResponseObjectOk returns a tuple with the ResponseObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseObject

`func (o *AuditEventAttributes) SetResponseObject(v string)`

SetResponseObject sets ResponseObject field to given value.

### HasResponseObject

`func (o *AuditEventAttributes) HasResponseObject() bool`

HasResponseObject returns a boolean if a field has been set.

### GetServiceName

`func (o *AuditEventAttributes) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *AuditEventAttributes) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *AuditEventAttributes) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *AuditEventAttributes) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetStage

`func (o *AuditEventAttributes) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *AuditEventAttributes) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *AuditEventAttributes) SetStage(v string)`

SetStage sets Stage field to given value.

### HasStage

`func (o *AuditEventAttributes) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetUser

`func (o *AuditEventAttributes) GetUser() ApiObjectRef`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AuditEventAttributes) GetUserOk() (*ApiObjectRef, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AuditEventAttributes) SetUser(v ApiObjectRef)`

SetUser sets User field to given value.

### HasUser

`func (o *AuditEventAttributes) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


