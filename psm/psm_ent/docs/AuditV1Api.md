# \AuditV1Api

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGetEvent**](AuditV1Api.md#GetGetEvent) | **Get** /audit/v1/events/{UUID} | Fetches an audit event given its uuid



## GetGetEvent

> AuditAuditEvent GetGetEvent(ctx, uUID).Execute()

Fetches an audit event given its uuid

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
    uUID := "uUID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuditV1Api.GetGetEvent(context.Background(), uUID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditV1Api.GetGetEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGetEvent`: AuditAuditEvent
    fmt.Fprintf(os.Stdout, "Response from `AuditV1Api.GetGetEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGetEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuditAuditEvent**](auditAuditEvent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

