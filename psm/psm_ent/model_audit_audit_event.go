/*
 * Audit API reference
 *
 * Service name  
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package psm_ent

import (
	"encoding/json"
)

// AuditAuditEvent struct for AuditAuditEvent
type AuditAuditEvent struct {
	Action *string `json:"action,omitempty"`
	ApiVersion *string `json:"api-version,omitempty"`
	ClientIps *[]string `json:"client-ips,omitempty"`
	Data *map[string]string `json:"data,omitempty"`
	// Length of string should be between 0 and 64. Must be alpha numeric and can have -.
	ExternalId *string `json:"external-id,omitempty"`
	GatewayIp *string `json:"gateway-ip,omitempty"`
	GatewayNode *string `json:"gateway-node,omitempty"`
	Kind *string `json:"kind,omitempty"`
	Level *string `json:"level,omitempty"`
	Meta *ApiObjectMeta `json:"meta,omitempty"`
	Outcome *string `json:"outcome,omitempty"`
	RequestObject *string `json:"request-object,omitempty"`
	// Should be a valid URI.
	RequestUri *string `json:"request-uri,omitempty"`
	Resource *ApiObjectRef `json:"resource,omitempty"`
	ResponseObject *string `json:"response-object,omitempty"`
	ServiceName *string `json:"service-name,omitempty"`
	Stage *string `json:"stage,omitempty"`
	User *ApiObjectRef `json:"user,omitempty"`
}

// NewAuditAuditEvent instantiates a new AuditAuditEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditAuditEvent() *AuditAuditEvent {
	this := AuditAuditEvent{}
	var level string = "basic"
	this.Level = &level
	var outcome string = "success"
	this.Outcome = &outcome
	var stage string = "requestauthorization"
	this.Stage = &stage
	return &this
}

// NewAuditAuditEventWithDefaults instantiates a new AuditAuditEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditAuditEventWithDefaults() *AuditAuditEvent {
	this := AuditAuditEvent{}
	var level string = "basic"
	this.Level = &level
	var outcome string = "success"
	this.Outcome = &outcome
	var stage string = "requestauthorization"
	this.Stage = &stage
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *AuditAuditEvent) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditAuditEvent) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *AuditAuditEvent) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *AuditAuditEvent) SetAction(v string) {
	o.Action = &v
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *AuditAuditEvent) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditAuditEvent) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *AuditAuditEvent) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *AuditAuditEvent) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetClientIps returns the ClientIps field value if set, zero value otherwise.
func (o *AuditAuditEvent) GetClientIps() []string {
	if o == nil || o.ClientIps == nil {
		var ret []string
		return ret
	}
	return *o.ClientIps
}

// GetClientIpsOk returns a tuple with the ClientIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditAuditEvent) GetClientIpsOk() (*[]string, bool) {
	if o == nil || o.ClientIps == nil {
		return nil, false
	}
	return o.ClientIps, true
}

// HasClientIps returns a boolean if a field has been set.
func (o *AuditAuditEvent) HasClientIps() bool {
	if o != nil && o.ClientIps != nil {
		return true
	}

	return false
}

// SetClientIps gets a reference to the given []string and assigns it to the ClientIps field.
func (o *AuditAuditEvent) SetClientIps(v []string) {
	o.ClientIps = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AuditAuditEvent) GetData() map[string]string {
	if o == nil || o.Data == nil {
		var ret map[string]string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditAuditEvent) GetDataOk() (*map[string]string, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AuditAuditEvent) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]string and assigns it to the Data field.
func (o *AuditAuditEvent) SetData(v map[string]string) {
	o.Data = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *AuditAuditEvent) GetExternalId() string {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditAuditEvent) GetExternalIdOk() (*string, bool) {
	if o == nil || o.ExternalId == nil {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *AuditAuditEvent) HasExternalId() bool {
	if o != nil && o.ExternalId != nil {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *AuditAuditEvent) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetGatewayIp returns the GatewayIp field value if set, zero value otherwise.
func (o *AuditAuditEvent) GetGatewayIp() string {
	if o == nil || o.GatewayIp == nil {
		var ret string
		return ret
	}
	return *o.GatewayIp
}

// GetGatewayIpOk returns a tuple with the GatewayIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditAuditEvent) GetGatewayIpOk() (*string, bool) {
	if o == nil || o.GatewayIp == nil {
		return nil, false
	}
	return o.GatewayIp, true
}

// HasGatewayIp returns a boolean if a field has been set.
func (o *AuditAuditEvent) HasGatewayIp() bool {
	if o != nil && o.GatewayIp != nil {
		return true
	}

	return false
}

// SetGatewayIp gets a reference to the given string and assigns it to the GatewayIp field.
func (o *AuditAuditEvent) SetGatewayIp(v string) {
	o.GatewayIp = &v
}

// GetGatewayNode returns the GatewayNode field value if set, zero value otherwise.
func (o *AuditAuditEvent) GetGatewayNode() string {
	if o == nil || o.GatewayNode == nil {
		var ret string
		return ret
	}
	return *o.GatewayNode
}

// GetGatewayNodeOk returns a tuple with the GatewayNode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditAuditEvent) GetGatewayNodeOk() (*string, bool) {
	if o == nil || o.GatewayNode == nil {
		return nil, false
	}
	return o.GatewayNode, true
}

// HasGatewayNode returns a boolean if a field has been set.
func (o *AuditAuditEvent) HasGatewayNode() bool {
	if o != nil && o.GatewayNode != nil {
		return true
	}

	return false
}

// SetGatewayNode gets a reference to the given string and assigns it to the GatewayNode field.
func (o *AuditAuditEvent) SetGatewayNode(v string) {
	o.GatewayNode = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *AuditAuditEvent) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditAuditEvent) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *AuditAuditEvent) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *AuditAuditEvent) SetKind(v string) {
	o.Kind = &v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *AuditAuditEvent) GetLevel() string {
	if o == nil || o.Level == nil {
		var ret string
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditAuditEvent) GetLevelOk() (*string, bool) {
	if o == nil || o.Level == nil {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *AuditAuditEvent) HasLevel() bool {
	if o != nil && o.Level != nil {
		return true
	}

	return false
}

// SetLevel gets a reference to the given string and assigns it to the Level field.
func (o *AuditAuditEvent) SetLevel(v string) {
	o.Level = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AuditAuditEvent) GetMeta() ApiObjectMeta {
	if o == nil || o.Meta == nil {
		var ret ApiObjectMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditAuditEvent) GetMetaOk() (*ApiObjectMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AuditAuditEvent) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ApiObjectMeta and assigns it to the Meta field.
func (o *AuditAuditEvent) SetMeta(v ApiObjectMeta) {
	o.Meta = &v
}

// GetOutcome returns the Outcome field value if set, zero value otherwise.
func (o *AuditAuditEvent) GetOutcome() string {
	if o == nil || o.Outcome == nil {
		var ret string
		return ret
	}
	return *o.Outcome
}

// GetOutcomeOk returns a tuple with the Outcome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditAuditEvent) GetOutcomeOk() (*string, bool) {
	if o == nil || o.Outcome == nil {
		return nil, false
	}
	return o.Outcome, true
}

// HasOutcome returns a boolean if a field has been set.
func (o *AuditAuditEvent) HasOutcome() bool {
	if o != nil && o.Outcome != nil {
		return true
	}

	return false
}

// SetOutcome gets a reference to the given string and assigns it to the Outcome field.
func (o *AuditAuditEvent) SetOutcome(v string) {
	o.Outcome = &v
}

// GetRequestObject returns the RequestObject field value if set, zero value otherwise.
func (o *AuditAuditEvent) GetRequestObject() string {
	if o == nil || o.RequestObject == nil {
		var ret string
		return ret
	}
	return *o.RequestObject
}

// GetRequestObjectOk returns a tuple with the RequestObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditAuditEvent) GetRequestObjectOk() (*string, bool) {
	if o == nil || o.RequestObject == nil {
		return nil, false
	}
	return o.RequestObject, true
}

// HasRequestObject returns a boolean if a field has been set.
func (o *AuditAuditEvent) HasRequestObject() bool {
	if o != nil && o.RequestObject != nil {
		return true
	}

	return false
}

// SetRequestObject gets a reference to the given string and assigns it to the RequestObject field.
func (o *AuditAuditEvent) SetRequestObject(v string) {
	o.RequestObject = &v
}

// GetRequestUri returns the RequestUri field value if set, zero value otherwise.
func (o *AuditAuditEvent) GetRequestUri() string {
	if o == nil || o.RequestUri == nil {
		var ret string
		return ret
	}
	return *o.RequestUri
}

// GetRequestUriOk returns a tuple with the RequestUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditAuditEvent) GetRequestUriOk() (*string, bool) {
	if o == nil || o.RequestUri == nil {
		return nil, false
	}
	return o.RequestUri, true
}

// HasRequestUri returns a boolean if a field has been set.
func (o *AuditAuditEvent) HasRequestUri() bool {
	if o != nil && o.RequestUri != nil {
		return true
	}

	return false
}

// SetRequestUri gets a reference to the given string and assigns it to the RequestUri field.
func (o *AuditAuditEvent) SetRequestUri(v string) {
	o.RequestUri = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *AuditAuditEvent) GetResource() ApiObjectRef {
	if o == nil || o.Resource == nil {
		var ret ApiObjectRef
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditAuditEvent) GetResourceOk() (*ApiObjectRef, bool) {
	if o == nil || o.Resource == nil {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *AuditAuditEvent) HasResource() bool {
	if o != nil && o.Resource != nil {
		return true
	}

	return false
}

// SetResource gets a reference to the given ApiObjectRef and assigns it to the Resource field.
func (o *AuditAuditEvent) SetResource(v ApiObjectRef) {
	o.Resource = &v
}

// GetResponseObject returns the ResponseObject field value if set, zero value otherwise.
func (o *AuditAuditEvent) GetResponseObject() string {
	if o == nil || o.ResponseObject == nil {
		var ret string
		return ret
	}
	return *o.ResponseObject
}

// GetResponseObjectOk returns a tuple with the ResponseObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditAuditEvent) GetResponseObjectOk() (*string, bool) {
	if o == nil || o.ResponseObject == nil {
		return nil, false
	}
	return o.ResponseObject, true
}

// HasResponseObject returns a boolean if a field has been set.
func (o *AuditAuditEvent) HasResponseObject() bool {
	if o != nil && o.ResponseObject != nil {
		return true
	}

	return false
}

// SetResponseObject gets a reference to the given string and assigns it to the ResponseObject field.
func (o *AuditAuditEvent) SetResponseObject(v string) {
	o.ResponseObject = &v
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise.
func (o *AuditAuditEvent) GetServiceName() string {
	if o == nil || o.ServiceName == nil {
		var ret string
		return ret
	}
	return *o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditAuditEvent) GetServiceNameOk() (*string, bool) {
	if o == nil || o.ServiceName == nil {
		return nil, false
	}
	return o.ServiceName, true
}

// HasServiceName returns a boolean if a field has been set.
func (o *AuditAuditEvent) HasServiceName() bool {
	if o != nil && o.ServiceName != nil {
		return true
	}

	return false
}

// SetServiceName gets a reference to the given string and assigns it to the ServiceName field.
func (o *AuditAuditEvent) SetServiceName(v string) {
	o.ServiceName = &v
}

// GetStage returns the Stage field value if set, zero value otherwise.
func (o *AuditAuditEvent) GetStage() string {
	if o == nil || o.Stage == nil {
		var ret string
		return ret
	}
	return *o.Stage
}

// GetStageOk returns a tuple with the Stage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditAuditEvent) GetStageOk() (*string, bool) {
	if o == nil || o.Stage == nil {
		return nil, false
	}
	return o.Stage, true
}

// HasStage returns a boolean if a field has been set.
func (o *AuditAuditEvent) HasStage() bool {
	if o != nil && o.Stage != nil {
		return true
	}

	return false
}

// SetStage gets a reference to the given string and assigns it to the Stage field.
func (o *AuditAuditEvent) SetStage(v string) {
	o.Stage = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *AuditAuditEvent) GetUser() ApiObjectRef {
	if o == nil || o.User == nil {
		var ret ApiObjectRef
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditAuditEvent) GetUserOk() (*ApiObjectRef, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *AuditAuditEvent) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given ApiObjectRef and assigns it to the User field.
func (o *AuditAuditEvent) SetUser(v ApiObjectRef) {
	o.User = &v
}

func (o AuditAuditEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.ApiVersion != nil {
		toSerialize["api-version"] = o.ApiVersion
	}
	if o.ClientIps != nil {
		toSerialize["client-ips"] = o.ClientIps
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.ExternalId != nil {
		toSerialize["external-id"] = o.ExternalId
	}
	if o.GatewayIp != nil {
		toSerialize["gateway-ip"] = o.GatewayIp
	}
	if o.GatewayNode != nil {
		toSerialize["gateway-node"] = o.GatewayNode
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Level != nil {
		toSerialize["level"] = o.Level
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.Outcome != nil {
		toSerialize["outcome"] = o.Outcome
	}
	if o.RequestObject != nil {
		toSerialize["request-object"] = o.RequestObject
	}
	if o.RequestUri != nil {
		toSerialize["request-uri"] = o.RequestUri
	}
	if o.Resource != nil {
		toSerialize["resource"] = o.Resource
	}
	if o.ResponseObject != nil {
		toSerialize["response-object"] = o.ResponseObject
	}
	if o.ServiceName != nil {
		toSerialize["service-name"] = o.ServiceName
	}
	if o.Stage != nil {
		toSerialize["stage"] = o.Stage
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	return json.Marshal(toSerialize)
}

type NullableAuditAuditEvent struct {
	value *AuditAuditEvent
	isSet bool
}

func (v NullableAuditAuditEvent) Get() *AuditAuditEvent {
	return v.value
}

func (v *NullableAuditAuditEvent) Set(val *AuditAuditEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditAuditEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditAuditEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditAuditEvent(val *AuditAuditEvent) *NullableAuditAuditEvent {
	return &NullableAuditAuditEvent{value: val, isSet: true}
}

func (v NullableAuditAuditEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditAuditEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


