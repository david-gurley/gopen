# \TelemetryQueryV1Api

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMetrics1**](TelemetryQueryV1Api.md#GetMetrics1) | **Get** /telemetry/v1/metrics | telemetry metrics query
[**PostMetrics**](TelemetryQueryV1Api.md#PostMetrics) | **Post** /telemetry/v1/metrics | telemetry metrics query



## GetMetrics1

> TelemetryQueryMetricsQueryResponse GetMetrics1(ctx).Tenant(tenant).Execute()

telemetry metrics query

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
    tenant := "tenant_example" // string | Tenant for the request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TelemetryQueryV1Api.GetMetrics1(context.Background()).Tenant(tenant).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TelemetryQueryV1Api.GetMetrics1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMetrics1`: TelemetryQueryMetricsQueryResponse
    fmt.Fprintf(os.Stdout, "Response from `TelemetryQueryV1Api.GetMetrics1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMetrics1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenant** | **string** | Tenant for the request. | 

### Return type

[**TelemetryQueryMetricsQueryResponse**](telemetry_queryMetricsQueryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMetrics

> TelemetryQueryMetricsQueryResponse PostMetrics(ctx).Body(body).Execute()

telemetry metrics query

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
    body := *openapiclient.NewTelemetryQueryMetricsQueryList() // TelemetryQueryMetricsQueryList | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TelemetryQueryV1Api.PostMetrics(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TelemetryQueryV1Api.PostMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostMetrics`: TelemetryQueryMetricsQueryResponse
    fmt.Fprintf(os.Stdout, "Response from `TelemetryQueryV1Api.PostMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**TelemetryQueryMetricsQueryList**](TelemetryQueryMetricsQueryList.md) |  | 

### Return type

[**TelemetryQueryMetricsQueryResponse**](telemetry_queryMetricsQueryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

