# \FwlogV1Api

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDownloadFwLogFileContent**](FwlogV1Api.md#GetDownloadFwLogFileContent) | **Get** /fwlog/v1/tenants/{O.Tenant}/objects/{O.Name} | fwlog/v1/tenants/default/objects/&lt;objectName&gt;
[**GetDownloadFwLogFileContent1**](FwlogV1Api.md#GetDownloadFwLogFileContent1) | **Get** /fwlog/v1/objects/{O.Name} | fwlog/v1/tenants/default/objects/&lt;objectName&gt;
[**GetGetLogs1**](FwlogV1Api.md#GetGetLogs1) | **Get** /fwlog/v1/query | Queries firewall logs
[**GetWatchLogs**](FwlogV1Api.md#GetWatchLogs) | **Get** /fwlog/v1/watch/query | 
[**PostGetLogs**](FwlogV1Api.md#PostGetLogs) | **Post** /fwlog/v1/query | Queries firewall logs



## GetDownloadFwLogFileContent

> FwlogFwLogList GetDownloadFwLogFileContent(ctx, oTenant).ONamespace(oNamespace).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

fwlog/v1/tenants/default/objects/<objectName>

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    oTenant := "oTenant_example" // string | 
    oNamespace := "oNamespace_example" // string | Namespace of the object, for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. (optional)
    oCreationTime := time.Now() // time.Time | CreationTime is the creation time of the object. System generated and updated, not updatable by user. (optional)
    fieldChangeSelector := []string{"Inner_example"} // []string | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FwlogV1Api.GetDownloadFwLogFileContent(context.Background(), oTenant).ONamespace(oNamespace).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FwlogV1Api.GetDownloadFwLogFileContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDownloadFwLogFileContent`: FwlogFwLogList
    fmt.Fprintf(os.Stdout, "Response from `FwlogV1Api.GetDownloadFwLogFileContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDownloadFwLogFileContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oNamespace** | **string** | Namespace of the object, for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**FwlogFwLogList**](fwlogFwLogList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDownloadFwLogFileContent1

> FwlogFwLogList GetDownloadFwLogFileContent1(ctx, oName).OTenant(oTenant).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

fwlog/v1/tenants/default/objects/<objectName>

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    oName := "oName_example" // string | 
    oTenant := "oTenant_example" // string | Tenant to which the object belongs to. This can be automatically filled in many cases based on the tenant the user, who created the object, belongs to. Must be alpha-numerics. Length of string should be between 1 and 48. (optional)
    oCreationTime := time.Now() // time.Time | CreationTime is the creation time of the object. System generated and updated, not updatable by user. (optional)
    fieldChangeSelector := []string{"Inner_example"} // []string | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FwlogV1Api.GetDownloadFwLogFileContent1(context.Background(), oName).OTenant(oTenant).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FwlogV1Api.GetDownloadFwLogFileContent1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDownloadFwLogFileContent1`: FwlogFwLogList
    fmt.Fprintf(os.Stdout, "Response from `FwlogV1Api.GetDownloadFwLogFileContent1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDownloadFwLogFileContent1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oTenant** | **string** | Tenant to which the object belongs to. This can be automatically filled in many cases based on the tenant the user, who created the object, belongs to. Must be alpha-numerics. Length of string should be between 1 and 48. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**FwlogFwLogList**](fwlogFwLogList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGetLogs1

> FwlogFwLogList GetGetLogs1(ctx).SourceIps(sourceIps).StartTime(startTime).SortOrder(sortOrder).Execute()

Queries firewall logs

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    sourceIps := []string{"Inner_example"} // []string | OR of sources IPs to be matched. Only one source IP is allowed. Should be a valid v4 or v6 IP address. (optional)
    startTime := time.Now() // time.Time | StartTime selects all logs with timestamp greater than the StartTime, example 2018-10-18T00:12:00Z. (optional)
    sortOrder := "sortOrder_example" // string | SortOrder specifies time ordering of results. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FwlogV1Api.GetGetLogs1(context.Background()).SourceIps(sourceIps).StartTime(startTime).SortOrder(sortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FwlogV1Api.GetGetLogs1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGetLogs1`: FwlogFwLogList
    fmt.Fprintf(os.Stdout, "Response from `FwlogV1Api.GetGetLogs1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGetLogs1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceIps** | **[]string** | OR of sources IPs to be matched. Only one source IP is allowed. Should be a valid v4 or v6 IP address. | 
 **startTime** | **time.Time** | StartTime selects all logs with timestamp greater than the StartTime, example 2018-10-18T00:12:00Z. | 
 **sortOrder** | **string** | SortOrder specifies time ordering of results. | 

### Return type

[**FwlogFwLogList**](fwlogFwLogList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWatchLogs

> FwlogFwLogList GetWatchLogs(ctx).SourceIps(sourceIps).StartTime(startTime).SortOrder(sortOrder).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    sourceIps := []string{"Inner_example"} // []string | OR of sources IPs to be matched. Only one source IP is allowed. Should be a valid v4 or v6 IP address. (optional)
    startTime := time.Now() // time.Time | StartTime selects all logs with timestamp greater than the StartTime, example 2018-10-18T00:12:00Z. (optional)
    sortOrder := "sortOrder_example" // string | SortOrder specifies time ordering of results. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FwlogV1Api.GetWatchLogs(context.Background()).SourceIps(sourceIps).StartTime(startTime).SortOrder(sortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FwlogV1Api.GetWatchLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWatchLogs`: FwlogFwLogList
    fmt.Fprintf(os.Stdout, "Response from `FwlogV1Api.GetWatchLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWatchLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceIps** | **[]string** | OR of sources IPs to be matched. Only one source IP is allowed. Should be a valid v4 or v6 IP address. | 
 **startTime** | **time.Time** | StartTime selects all logs with timestamp greater than the StartTime, example 2018-10-18T00:12:00Z. | 
 **sortOrder** | **string** | SortOrder specifies time ordering of results. | 

### Return type

[**FwlogFwLogList**](fwlogFwLogList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGetLogs

> FwlogFwLogList PostGetLogs(ctx).Body(body).Execute()

Queries firewall logs

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewFwlogFwLogQuery() // FwlogFwLogQuery | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FwlogV1Api.PostGetLogs(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FwlogV1Api.PostGetLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostGetLogs`: FwlogFwLogList
    fmt.Fprintf(os.Stdout, "Response from `FwlogV1Api.PostGetLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostGetLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**FwlogFwLogQuery**](FwlogFwLogQuery.md) |  | 

### Return type

[**FwlogFwLogList**](fwlogFwLogList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

