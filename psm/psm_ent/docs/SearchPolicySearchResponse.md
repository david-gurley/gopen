# SearchPolicySearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**SearchError**](searchError.md) |  | [optional] 
**Results** | Pointer to [**map[string]SearchRulesByPolicyName**](searchRulesByPolicyName.md) | Result is Map of &lt;Kind, Object name, PolicyMatch Entry&gt;. | [optional] 
**Status** | Pointer to **string** | Status of firewall policy search. | [optional] [default to "match"]

## Methods

### NewSearchPolicySearchResponse

`func NewSearchPolicySearchResponse() *SearchPolicySearchResponse`

NewSearchPolicySearchResponse instantiates a new SearchPolicySearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchPolicySearchResponseWithDefaults

`func NewSearchPolicySearchResponseWithDefaults() *SearchPolicySearchResponse`

NewSearchPolicySearchResponseWithDefaults instantiates a new SearchPolicySearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *SearchPolicySearchResponse) GetError() SearchError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SearchPolicySearchResponse) GetErrorOk() (*SearchError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SearchPolicySearchResponse) SetError(v SearchError)`

SetError sets Error field to given value.

### HasError

`func (o *SearchPolicySearchResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetResults

`func (o *SearchPolicySearchResponse) GetResults() map[string]SearchRulesByPolicyName`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SearchPolicySearchResponse) GetResultsOk() (*map[string]SearchRulesByPolicyName, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SearchPolicySearchResponse) SetResults(v map[string]SearchRulesByPolicyName)`

SetResults sets Results field to given value.

### HasResults

`func (o *SearchPolicySearchResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetStatus

`func (o *SearchPolicySearchResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SearchPolicySearchResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SearchPolicySearchResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SearchPolicySearchResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


