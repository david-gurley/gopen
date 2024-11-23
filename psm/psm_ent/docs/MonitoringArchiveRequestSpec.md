# MonitoringArchiveRequestSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | Pointer to [**MonitoringArchiveQuery**](monitoringArchiveQuery.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] [default to "event"]

## Methods

### NewMonitoringArchiveRequestSpec

`func NewMonitoringArchiveRequestSpec() *MonitoringArchiveRequestSpec`

NewMonitoringArchiveRequestSpec instantiates a new MonitoringArchiveRequestSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringArchiveRequestSpecWithDefaults

`func NewMonitoringArchiveRequestSpecWithDefaults() *MonitoringArchiveRequestSpec`

NewMonitoringArchiveRequestSpecWithDefaults instantiates a new MonitoringArchiveRequestSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *MonitoringArchiveRequestSpec) GetQuery() MonitoringArchiveQuery`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *MonitoringArchiveRequestSpec) GetQueryOk() (*MonitoringArchiveQuery, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *MonitoringArchiveRequestSpec) SetQuery(v MonitoringArchiveQuery)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *MonitoringArchiveRequestSpec) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetType

`func (o *MonitoringArchiveRequestSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MonitoringArchiveRequestSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MonitoringArchiveRequestSpec) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MonitoringArchiveRequestSpec) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


