# \DistributedservicecardsV1Api

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostQueryConnection**](DistributedservicecardsV1Api.md#PostQueryConnection) | **Post** /sysruntime/distributedservicecards/v1/{DSCName}/connections | Active Connection Query



## PostQueryConnection

> SysruntimeConnectionStatus PostQueryConnection(ctx, dSCName).Body(body).Execute()

Active Connection Query

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
    dSCName := "dSCName_example" // string | 
    body := *openapiclient.NewSysruntimeConnectionRequest() // SysruntimeConnectionRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DistributedservicecardsV1Api.PostQueryConnection(context.Background(), dSCName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DistributedservicecardsV1Api.PostQueryConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostQueryConnection`: SysruntimeConnectionStatus
    fmt.Fprintf(os.Stdout, "Response from `DistributedservicecardsV1Api.PostQueryConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dSCName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostQueryConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SysruntimeConnectionRequest**](SysruntimeConnectionRequest.md) |  | 

### Return type

[**SysruntimeConnectionStatus**](sysruntimeConnectionStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

