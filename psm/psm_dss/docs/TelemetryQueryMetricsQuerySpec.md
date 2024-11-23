# TelemetryQueryMetricsQuerySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**BottomParam** | Pointer to [**TelemetryQueryBottomSpec**](telemetry_queryBottomSpec.md) |  | [optional] 
**EndTime** | Pointer to **time.Time** | EndTime selects all metrics with timestamp less than the EndTime, example 2018-09-18T00:12:00Z. | [optional] 
**Fields** | Pointer to **[]string** | Fields select the metric fields to be included in the result Empty will include all fields, must contain at least one non-tag field. Must start and end with alpha numeric and can have alphanumeric, -, _, . | [optional] 
**Function** | Pointer to **string** | Functions specify an operation function to be applied, example mean()/max(). | [optional] [default to "none"]
**GroupByField** | Pointer to **string** | GroupbyField groups series based on the field specified. Must start and end with alpha numeric and can have alphanumeric, -, _, . and ,. | [optional] 
**GroupByTime** | Pointer to **string** | GroupbyTime groups series based on the interval specified. Should be a valid time duration. | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | Name is the name of the API object. Must start and end with alpha numeric and can have alphanumeric, -, _, . | [optional] 
**Pagination** | Pointer to [**TelemetryQueryPaginationSpec**](telemetry_queryPaginationSpec.md) |  | [optional] 
**Selector** | Pointer to [**FieldsSelector**](fieldsSelector.md) |  | [optional] 
**SortOrder** | Pointer to **string** | SortOrder specifies time ordering of results. | [optional] [default to "ascending"]
**StartTime** | Pointer to **time.Time** | StartTime selects all metrics with timestamp greater than the StartTime, example 2018-10-18T00:12:00Z. | [optional] 
**Subquery** | Pointer to [**TelemetryQueryMetricsQuerySpec**](telemetry_queryMetricsQuerySpec.md) |  | [optional] 
**TopParam** | Pointer to [**TelemetryQueryTopSpec**](telemetry_queryTopSpec.md) |  | [optional] 

## Methods

### NewTelemetryQueryMetricsQuerySpec

`func NewTelemetryQueryMetricsQuerySpec() *TelemetryQueryMetricsQuerySpec`

NewTelemetryQueryMetricsQuerySpec instantiates a new TelemetryQueryMetricsQuerySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryQueryMetricsQuerySpecWithDefaults

`func NewTelemetryQueryMetricsQuerySpecWithDefaults() *TelemetryQueryMetricsQuerySpec`

NewTelemetryQueryMetricsQuerySpecWithDefaults instantiates a new TelemetryQueryMetricsQuerySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *TelemetryQueryMetricsQuerySpec) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *TelemetryQueryMetricsQuerySpec) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *TelemetryQueryMetricsQuerySpec) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *TelemetryQueryMetricsQuerySpec) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetBottomParam

`func (o *TelemetryQueryMetricsQuerySpec) GetBottomParam() TelemetryQueryBottomSpec`

GetBottomParam returns the BottomParam field if non-nil, zero value otherwise.

### GetBottomParamOk

`func (o *TelemetryQueryMetricsQuerySpec) GetBottomParamOk() (*TelemetryQueryBottomSpec, bool)`

GetBottomParamOk returns a tuple with the BottomParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBottomParam

`func (o *TelemetryQueryMetricsQuerySpec) SetBottomParam(v TelemetryQueryBottomSpec)`

SetBottomParam sets BottomParam field to given value.

### HasBottomParam

`func (o *TelemetryQueryMetricsQuerySpec) HasBottomParam() bool`

HasBottomParam returns a boolean if a field has been set.

### GetEndTime

`func (o *TelemetryQueryMetricsQuerySpec) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *TelemetryQueryMetricsQuerySpec) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *TelemetryQueryMetricsQuerySpec) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *TelemetryQueryMetricsQuerySpec) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetFields

`func (o *TelemetryQueryMetricsQuerySpec) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *TelemetryQueryMetricsQuerySpec) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *TelemetryQueryMetricsQuerySpec) SetFields(v []string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *TelemetryQueryMetricsQuerySpec) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFunction

`func (o *TelemetryQueryMetricsQuerySpec) GetFunction() string`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *TelemetryQueryMetricsQuerySpec) GetFunctionOk() (*string, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *TelemetryQueryMetricsQuerySpec) SetFunction(v string)`

SetFunction sets Function field to given value.

### HasFunction

`func (o *TelemetryQueryMetricsQuerySpec) HasFunction() bool`

HasFunction returns a boolean if a field has been set.

### GetGroupByField

`func (o *TelemetryQueryMetricsQuerySpec) GetGroupByField() string`

GetGroupByField returns the GroupByField field if non-nil, zero value otherwise.

### GetGroupByFieldOk

`func (o *TelemetryQueryMetricsQuerySpec) GetGroupByFieldOk() (*string, bool)`

GetGroupByFieldOk returns a tuple with the GroupByField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByField

`func (o *TelemetryQueryMetricsQuerySpec) SetGroupByField(v string)`

SetGroupByField sets GroupByField field to given value.

### HasGroupByField

`func (o *TelemetryQueryMetricsQuerySpec) HasGroupByField() bool`

HasGroupByField returns a boolean if a field has been set.

### GetGroupByTime

`func (o *TelemetryQueryMetricsQuerySpec) GetGroupByTime() string`

GetGroupByTime returns the GroupByTime field if non-nil, zero value otherwise.

### GetGroupByTimeOk

`func (o *TelemetryQueryMetricsQuerySpec) GetGroupByTimeOk() (*string, bool)`

GetGroupByTimeOk returns a tuple with the GroupByTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupByTime

`func (o *TelemetryQueryMetricsQuerySpec) SetGroupByTime(v string)`

SetGroupByTime sets GroupByTime field to given value.

### HasGroupByTime

`func (o *TelemetryQueryMetricsQuerySpec) HasGroupByTime() bool`

HasGroupByTime returns a boolean if a field has been set.

### GetKind

`func (o *TelemetryQueryMetricsQuerySpec) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TelemetryQueryMetricsQuerySpec) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TelemetryQueryMetricsQuerySpec) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *TelemetryQueryMetricsQuerySpec) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *TelemetryQueryMetricsQuerySpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TelemetryQueryMetricsQuerySpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TelemetryQueryMetricsQuerySpec) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TelemetryQueryMetricsQuerySpec) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPagination

`func (o *TelemetryQueryMetricsQuerySpec) GetPagination() TelemetryQueryPaginationSpec`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *TelemetryQueryMetricsQuerySpec) GetPaginationOk() (*TelemetryQueryPaginationSpec, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *TelemetryQueryMetricsQuerySpec) SetPagination(v TelemetryQueryPaginationSpec)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *TelemetryQueryMetricsQuerySpec) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetSelector

`func (o *TelemetryQueryMetricsQuerySpec) GetSelector() FieldsSelector`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *TelemetryQueryMetricsQuerySpec) GetSelectorOk() (*FieldsSelector, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *TelemetryQueryMetricsQuerySpec) SetSelector(v FieldsSelector)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *TelemetryQueryMetricsQuerySpec) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### GetSortOrder

`func (o *TelemetryQueryMetricsQuerySpec) GetSortOrder() string`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *TelemetryQueryMetricsQuerySpec) GetSortOrderOk() (*string, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *TelemetryQueryMetricsQuerySpec) SetSortOrder(v string)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *TelemetryQueryMetricsQuerySpec) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetStartTime

`func (o *TelemetryQueryMetricsQuerySpec) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *TelemetryQueryMetricsQuerySpec) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *TelemetryQueryMetricsQuerySpec) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *TelemetryQueryMetricsQuerySpec) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetSubquery

`func (o *TelemetryQueryMetricsQuerySpec) GetSubquery() TelemetryQueryMetricsQuerySpec`

GetSubquery returns the Subquery field if non-nil, zero value otherwise.

### GetSubqueryOk

`func (o *TelemetryQueryMetricsQuerySpec) GetSubqueryOk() (*TelemetryQueryMetricsQuerySpec, bool)`

GetSubqueryOk returns a tuple with the Subquery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubquery

`func (o *TelemetryQueryMetricsQuerySpec) SetSubquery(v TelemetryQueryMetricsQuerySpec)`

SetSubquery sets Subquery field to given value.

### HasSubquery

`func (o *TelemetryQueryMetricsQuerySpec) HasSubquery() bool`

HasSubquery returns a boolean if a field has been set.

### GetTopParam

`func (o *TelemetryQueryMetricsQuerySpec) GetTopParam() TelemetryQueryTopSpec`

GetTopParam returns the TopParam field if non-nil, zero value otherwise.

### GetTopParamOk

`func (o *TelemetryQueryMetricsQuerySpec) GetTopParamOk() (*TelemetryQueryTopSpec, bool)`

GetTopParamOk returns a tuple with the TopParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopParam

`func (o *TelemetryQueryMetricsQuerySpec) SetTopParam(v TelemetryQueryTopSpec)`

SetTopParam sets TopParam field to given value.

### HasTopParam

`func (o *TelemetryQueryMetricsQuerySpec) HasTopParam() bool`

HasTopParam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


