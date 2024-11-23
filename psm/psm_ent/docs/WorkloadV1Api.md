# \WorkloadV1Api

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AbortMigration**](WorkloadV1Api.md#AbortMigration) | **Post** /configs/workload/v1/tenant/{O.Tenant}/workloads/{O.Name}/AbortMigration | Abort Workload Migration operation
[**AbortMigration1**](WorkloadV1Api.md#AbortMigration1) | **Post** /configs/workload/v1/workloads/{O.Name}/AbortMigration | Abort Workload Migration operation
[**AddWorkload**](WorkloadV1Api.md#AddWorkload) | **Post** /configs/workload/v1/tenant/{O.Tenant}/workloads | Create Workload object
[**AddWorkload1**](WorkloadV1Api.md#AddWorkload1) | **Post** /configs/workload/v1/workloads | Create Workload object
[**DeleteWorkload**](WorkloadV1Api.md#DeleteWorkload) | **Delete** /configs/workload/v1/tenant/{O.Tenant}/workloads/{O.Name} | Delete Workload object
[**DeleteWorkload1**](WorkloadV1Api.md#DeleteWorkload1) | **Delete** /configs/workload/v1/workloads/{O.Name} | Delete Workload object
[**FinalSyncMigration**](WorkloadV1Api.md#FinalSyncMigration) | **Post** /configs/workload/v1/tenant/{O.Tenant}/workloads/{O.Name}/FinalSyncMigration | Initiates the final sync for the Workload Migration operation
[**FinalSyncMigration1**](WorkloadV1Api.md#FinalSyncMigration1) | **Post** /configs/workload/v1/workloads/{O.Name}/FinalSyncMigration | Initiates the final sync for the Workload Migration operation
[**FinishMigration**](WorkloadV1Api.md#FinishMigration) | **Post** /configs/workload/v1/tenant/{O.Tenant}/workloads/{O.Name}/FinishMigration | Finish Workload Migration operation
[**FinishMigration1**](WorkloadV1Api.md#FinishMigration1) | **Post** /configs/workload/v1/workloads/{O.Name}/FinishMigration | Finish Workload Migration operation
[**GetEndpoint**](WorkloadV1Api.md#GetEndpoint) | **Get** /configs/workload/v1/tenant/{O.Tenant}/endpoints/{O.Name} | Get Endpoint object
[**GetEndpoint1**](WorkloadV1Api.md#GetEndpoint1) | **Get** /configs/workload/v1/endpoints/{O.Name} | Get Endpoint object
[**GetWorkload**](WorkloadV1Api.md#GetWorkload) | **Get** /configs/workload/v1/tenant/{O.Tenant}/workloads/{O.Name} | Get Workload object
[**GetWorkload1**](WorkloadV1Api.md#GetWorkload1) | **Get** /configs/workload/v1/workloads/{O.Name} | Get Workload object
[**LabelWorkload**](WorkloadV1Api.md#LabelWorkload) | **Post** /configs/workload/v1/tenant/{O.Tenant}/workloads/{O.Name}/label | Label Workload object
[**LabelWorkload1**](WorkloadV1Api.md#LabelWorkload1) | **Post** /configs/workload/v1/workloads/{O.Name}/label | Label Workload object
[**ListEndpoint**](WorkloadV1Api.md#ListEndpoint) | **Get** /configs/workload/v1/tenant/{O.Tenant}/endpoints | List Endpoint objects
[**ListEndpoint1**](WorkloadV1Api.md#ListEndpoint1) | **Get** /configs/workload/v1/endpoints | List Endpoint objects
[**ListWorkload**](WorkloadV1Api.md#ListWorkload) | **Get** /configs/workload/v1/tenant/{O.Tenant}/workloads | List Workload objects
[**ListWorkload1**](WorkloadV1Api.md#ListWorkload1) | **Get** /configs/workload/v1/workloads | List Workload objects
[**StartMigration**](WorkloadV1Api.md#StartMigration) | **Post** /configs/workload/v1/tenant/{O.Tenant}/workloads/{O.Name}/StartMigration | Start Workload Migration operation
[**StartMigration1**](WorkloadV1Api.md#StartMigration1) | **Post** /configs/workload/v1/workloads/{O.Name}/StartMigration | Start Workload Migration operation
[**UpdateWorkload**](WorkloadV1Api.md#UpdateWorkload) | **Put** /configs/workload/v1/tenant/{O.Tenant}/workloads/{O.Name} | Update Workload object
[**UpdateWorkload1**](WorkloadV1Api.md#UpdateWorkload1) | **Put** /configs/workload/v1/workloads/{O.Name} | Update Workload object
[**WatchEndpoint**](WorkloadV1Api.md#WatchEndpoint) | **Get** /configs/workload/v1/watch/tenant/{O.Tenant}/endpoints | Watch Endpoint objects. Supports WebSockets or HTTP long poll
[**WatchEndpoint1**](WorkloadV1Api.md#WatchEndpoint1) | **Get** /configs/workload/v1/watch/endpoints | Watch Endpoint objects. Supports WebSockets or HTTP long poll
[**WatchWorkload**](WorkloadV1Api.md#WatchWorkload) | **Get** /configs/workload/v1/watch/tenant/{O.Tenant}/workloads | Watch Workload objects. Supports WebSockets or HTTP long poll
[**WatchWorkload1**](WorkloadV1Api.md#WatchWorkload1) | **Get** /configs/workload/v1/watch/workloads | Watch Workload objects. Supports WebSockets or HTTP long poll



## AbortMigration

> WorkloadWorkload AbortMigration(ctx, oTenant).Body(body).Execute()

Abort Workload Migration operation

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
    oTenant := "oTenant_example" // string | 
    body := *openapiclient.NewWorkloadWorkload() // WorkloadWorkload | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkloadV1Api.AbortMigration(context.Background(), oTenant).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkloadV1Api.AbortMigration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AbortMigration`: WorkloadWorkload
    fmt.Fprintf(os.Stdout, "Response from `WorkloadV1Api.AbortMigration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAbortMigrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**WorkloadWorkload**](WorkloadWorkload.md) |  | 

### Return type

[**WorkloadWorkload**](workloadWorkload.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AbortMigration1

> WorkloadWorkload AbortMigration1(ctx, oName).Body(body).Execute()

Abort Workload Migration operation

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
    oName := "oName_example" // string | 
    body := *openapiclient.NewWorkloadWorkload() // WorkloadWorkload | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkloadV1Api.AbortMigration1(context.Background(), oName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkloadV1Api.AbortMigration1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AbortMigration1`: WorkloadWorkload
    fmt.Fprintf(os.Stdout, "Response from `WorkloadV1Api.AbortMigration1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAbortMigration1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**WorkloadWorkload**](WorkloadWorkload.md) |  | 

### Return type

[**WorkloadWorkload**](workloadWorkload.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddWorkload

> WorkloadWorkload AddWorkload(ctx, oTenant).Body(body).Execute()

Create Workload object

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
    oTenant := "oTenant_example" // string | 
    body := *openapiclient.NewWorkloadWorkload() // WorkloadWorkload | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkloadV1Api.AddWorkload(context.Background(), oTenant).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkloadV1Api.AddWorkload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddWorkload`: WorkloadWorkload
    fmt.Fprintf(os.Stdout, "Response from `WorkloadV1Api.AddWorkload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddWorkloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**WorkloadWorkload**](WorkloadWorkload.md) |  | 

### Return type

[**WorkloadWorkload**](workloadWorkload.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddWorkload1

> WorkloadWorkload AddWorkload1(ctx).Body(body).Execute()

Create Workload object

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
    body := *openapiclient.NewWorkloadWorkload() // WorkloadWorkload | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkloadV1Api.AddWorkload1(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkloadV1Api.AddWorkload1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddWorkload1`: WorkloadWorkload
    fmt.Fprintf(os.Stdout, "Response from `WorkloadV1Api.AddWorkload1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddWorkload1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**WorkloadWorkload**](WorkloadWorkload.md) |  | 

### Return type

[**WorkloadWorkload**](workloadWorkload.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWorkload

> WorkloadWorkload DeleteWorkload(ctx, oTenant).Execute()

Delete Workload object

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
    oTenant := "oTenant_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkloadV1Api.DeleteWorkload(context.Background(), oTenant).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkloadV1Api.DeleteWorkload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteWorkload`: WorkloadWorkload
    fmt.Fprintf(os.Stdout, "Response from `WorkloadV1Api.DeleteWorkload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkloadWorkload**](workloadWorkload.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWorkload1

> WorkloadWorkload DeleteWorkload1(ctx, oName).Execute()

Delete Workload object

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
    oName := "oName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkloadV1Api.DeleteWorkload1(context.Background(), oName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkloadV1Api.DeleteWorkload1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteWorkload1`: WorkloadWorkload
    fmt.Fprintf(os.Stdout, "Response from `WorkloadV1Api.DeleteWorkload1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkload1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkloadWorkload**](workloadWorkload.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FinalSyncMigration

> WorkloadWorkload FinalSyncMigration(ctx, oTenant).Body(body).Execute()

Initiates the final sync for the Workload Migration operation

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
    oTenant := "oTenant_example" // string | 
    body := *openapiclient.NewWorkloadWorkload() // WorkloadWorkload | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkloadV1Api.FinalSyncMigration(context.Background(), oTenant).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkloadV1Api.FinalSyncMigration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FinalSyncMigration`: WorkloadWorkload
    fmt.Fprintf(os.Stdout, "Response from `WorkloadV1Api.FinalSyncMigration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFinalSyncMigrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**WorkloadWorkload**](WorkloadWorkload.md) |  | 

### Return type

[**WorkloadWorkload**](workloadWorkload.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FinalSyncMigration1

> WorkloadWorkload FinalSyncMigration1(ctx, oName).Body(body).Execute()

Initiates the final sync for the Workload Migration operation

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
    oName := "oName_example" // string | 
    body := *openapiclient.NewWorkloadWorkload() // WorkloadWorkload | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkloadV1Api.FinalSyncMigration1(context.Background(), oName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkloadV1Api.FinalSyncMigration1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FinalSyncMigration1`: WorkloadWorkload
    fmt.Fprintf(os.Stdout, "Response from `WorkloadV1Api.FinalSyncMigration1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFinalSyncMigration1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**WorkloadWorkload**](WorkloadWorkload.md) |  | 

### Return type

[**WorkloadWorkload**](workloadWorkload.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FinishMigration

> WorkloadWorkload FinishMigration(ctx, oTenant).Body(body).Execute()

Finish Workload Migration operation

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
    oTenant := "oTenant_example" // string | 
    body := *openapiclient.NewWorkloadWorkload() // WorkloadWorkload | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkloadV1Api.FinishMigration(context.Background(), oTenant).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkloadV1Api.FinishMigration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FinishMigration`: WorkloadWorkload
    fmt.Fprintf(os.Stdout, "Response from `WorkloadV1Api.FinishMigration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFinishMigrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**WorkloadWorkload**](WorkloadWorkload.md) |  | 

### Return type

[**WorkloadWorkload**](workloadWorkload.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FinishMigration1

> WorkloadWorkload FinishMigration1(ctx, oName).Body(body).Execute()

Finish Workload Migration operation

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
    oName := "oName_example" // string | 
    body := *openapiclient.NewWorkloadWorkload() // WorkloadWorkload | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkloadV1Api.FinishMigration1(context.Background(), oName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkloadV1Api.FinishMigration1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FinishMigration1`: WorkloadWorkload
    fmt.Fprintf(os.Stdout, "Response from `WorkloadV1Api.FinishMigration1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFinishMigration1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**WorkloadWorkload**](WorkloadWorkload.md) |  | 

### Return type

[**WorkloadWorkload**](workloadWorkload.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEndpoint

> WorkloadEndpoint GetEndpoint(ctx, oTenant).TKind(tKind).MetaCreationTime(metaCreationTime).SpecNodeUuidList(specNodeUuidList).Execute()

Get Endpoint object

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
    tKind := "tKind_example" // string | Kind represents the type of the API object. (optional)
    metaCreationTime := time.Now() // time.Time | CreationTime is the creation time of the object. System generated and updated, not updatable by user. (optional)
    specNodeUuidList := []string{"Inner_example"} // []string | NodeUUIDList has the list of DSCs where a EP is supposed to go to. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkloadV1Api.GetEndpoint(context.Background(), oTenant).TKind(tKind).MetaCreationTime(metaCreationTime).SpecNodeUuidList(specNodeUuidList).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkloadV1Api.GetEndpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEndpoint`: WorkloadEndpoint
    fmt.Fprintf(os.Stdout, "Response from `WorkloadV1Api.GetEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tKind** | **string** | Kind represents the type of the API object. | 
 **metaCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **specNodeUuidList** | **[]string** | NodeUUIDList has the list of DSCs where a EP is supposed to go to. | 

### Return type

[**WorkloadEndpoint**](workloadEndpoint.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEndpoint1

> WorkloadEndpoint GetEndpoint1(ctx, oName).TKind(tKind).MetaCreationTime(metaCreationTime).SpecNodeUuidList(specNodeUuidList).Execute()

Get Endpoint object

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
    tKind := "tKind_example" // string | Kind represents the type of the API object. (optional)
    metaCreationTime := time.Now() // time.Time | CreationTime is the creation time of the object. System generated and updated, not updatable by user. (optional)
    specNodeUuidList := []string{"Inner_example"} // []string | NodeUUIDList has the list of DSCs where a EP is supposed to go to. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkloadV1Api.GetEndpoint1(context.Background(), oName).TKind(tKind).MetaCreationTime(metaCreationTime).SpecNodeUuidList(specNodeUuidList).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkloadV1Api.GetEndpoint1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEndpoint1`: WorkloadEndpoint
    fmt.Fprintf(os.Stdout, "Response from `WorkloadV1Api.GetEndpoint1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEndpoint1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tKind** | **string** | Kind represents the type of the API object. | 
 **metaCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **specNodeUuidList** | **[]string** | NodeUUIDList has the list of DSCs where a EP is supposed to go to. | 

### Return type

[**WorkloadEndpoint**](workloadEndpoint.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkload

> WorkloadWorkload GetWorkload(ctx, oTenant).TKind(tKind).MetaCreationTime(metaCreationTime).PropagationStatusPendingDscs(propagationStatusPendingDscs).Execute()

Get Workload object

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
    tKind := "tKind_example" // string | Kind represents the type of the API object. (optional)
    metaCreationTime := time.Now() // time.Time | CreationTime is the creation time of the object. System generated and updated, not updatable by user. (optional)
    propagationStatusPendingDscs := []string{"Inner_example"} // []string | list of DSCs where propagation did not complete. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkloadV1Api.GetWorkload(context.Background(), oTenant).TKind(tKind).MetaCreationTime(metaCreationTime).PropagationStatusPendingDscs(propagationStatusPendingDscs).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkloadV1Api.GetWorkload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkload`: WorkloadWorkload
    fmt.Fprintf(os.Stdout, "Response from `WorkloadV1Api.GetWorkload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tKind** | **string** | Kind represents the type of the API object. | 
 **metaCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **propagationStatusPendingDscs** | **[]string** | list of DSCs where propagation did not complete. | 

### Return type

[**WorkloadWorkload**](workloadWorkload.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkload1

> WorkloadWorkload GetWorkload1(ctx, oName).TKind(tKind).MetaCreationTime(metaCreationTime).PropagationStatusPendingDscs(propagationStatusPendingDscs).Execute()

Get Workload object

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
    tKind := "tKind_example" // string | Kind represents the type of the API object. (optional)
    metaCreationTime := time.Now() // time.Time | CreationTime is the creation time of the object. System generated and updated, not updatable by user. (optional)
    propagationStatusPendingDscs := []string{"Inner_example"} // []string | list of DSCs where propagation did not complete. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkloadV1Api.GetWorkload1(context.Background(), oName).TKind(tKind).MetaCreationTime(metaCreationTime).PropagationStatusPendingDscs(propagationStatusPendingDscs).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkloadV1Api.GetWorkload1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkload1`: WorkloadWorkload
    fmt.Fprintf(os.Stdout, "Response from `WorkloadV1Api.GetWorkload1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkload1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tKind** | **string** | Kind represents the type of the API object. | 
 **metaCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **propagationStatusPendingDscs** | **[]string** | list of DSCs where propagation did not complete. | 

### Return type

[**WorkloadWorkload**](workloadWorkload.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LabelWorkload

> WorkloadWorkload LabelWorkload(ctx, oTenant).Body(body).Execute()

Label Workload object

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
    oTenant := "oTenant_example" // string | 
    body := *openapiclient.NewApiLabel() // ApiLabel | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkloadV1Api.LabelWorkload(context.Background(), oTenant).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkloadV1Api.LabelWorkload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LabelWorkload`: WorkloadWorkload
    fmt.Fprintf(os.Stdout, "Response from `WorkloadV1Api.LabelWorkload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLabelWorkloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ApiLabel**](ApiLabel.md) |  | 

### Return type

[**WorkloadWorkload**](workloadWorkload.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LabelWorkload1

> WorkloadWorkload LabelWorkload1(ctx, oName).Body(body).Execute()

Label Workload object

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
    oName := "oName_example" // string | 
    body := *openapiclient.NewApiLabel() // ApiLabel | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkloadV1Api.LabelWorkload1(context.Background(), oName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkloadV1Api.LabelWorkload1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LabelWorkload1`: WorkloadWorkload
    fmt.Fprintf(os.Stdout, "Response from `WorkloadV1Api.LabelWorkload1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLabelWorkload1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ApiLabel**](ApiLabel.md) |  | 

### Return type

[**WorkloadWorkload**](workloadWorkload.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEndpoint

> WorkloadEndpointList ListEndpoint(ctx, oTenant).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

List Endpoint objects

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
    oName := "oName_example" // string | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. (optional)
    oCreationTime := time.Now() // time.Time | CreationTime is the creation time of the object. System generated and updated, not updatable by user. (optional)
    fieldChangeSelector := []string{"Inner_example"} // []string | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkloadV1Api.ListEndpoint(context.Background(), oTenant).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkloadV1Api.ListEndpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEndpoint`: WorkloadEndpointList
    fmt.Fprintf(os.Stdout, "Response from `WorkloadV1Api.ListEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**WorkloadEndpointList**](workloadEndpointList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEndpoint1

> WorkloadEndpointList ListEndpoint1(ctx).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

List Endpoint objects

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
    oName := "oName_example" // string | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. (optional)
    oCreationTime := time.Now() // time.Time | CreationTime is the creation time of the object. System generated and updated, not updatable by user. (optional)
    fieldChangeSelector := []string{"Inner_example"} // []string | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkloadV1Api.ListEndpoint1(context.Background()).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkloadV1Api.ListEndpoint1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEndpoint1`: WorkloadEndpointList
    fmt.Fprintf(os.Stdout, "Response from `WorkloadV1Api.ListEndpoint1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEndpoint1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**WorkloadEndpointList**](workloadEndpointList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkload

> WorkloadWorkloadList ListWorkload(ctx, oTenant).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

List Workload objects

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
    oName := "oName_example" // string | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. (optional)
    oCreationTime := time.Now() // time.Time | CreationTime is the creation time of the object. System generated and updated, not updatable by user. (optional)
    fieldChangeSelector := []string{"Inner_example"} // []string | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkloadV1Api.ListWorkload(context.Background(), oTenant).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkloadV1Api.ListWorkload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorkload`: WorkloadWorkloadList
    fmt.Fprintf(os.Stdout, "Response from `WorkloadV1Api.ListWorkload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**WorkloadWorkloadList**](workloadWorkloadList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkload1

> WorkloadWorkloadList ListWorkload1(ctx).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

List Workload objects

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
    oName := "oName_example" // string | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. (optional)
    oCreationTime := time.Now() // time.Time | CreationTime is the creation time of the object. System generated and updated, not updatable by user. (optional)
    fieldChangeSelector := []string{"Inner_example"} // []string | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkloadV1Api.ListWorkload1(context.Background()).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkloadV1Api.ListWorkload1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorkload1`: WorkloadWorkloadList
    fmt.Fprintf(os.Stdout, "Response from `WorkloadV1Api.ListWorkload1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWorkload1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**WorkloadWorkloadList**](workloadWorkloadList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartMigration

> WorkloadWorkload StartMigration(ctx, oTenant).Body(body).Execute()

Start Workload Migration operation

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
    oTenant := "oTenant_example" // string | 
    body := *openapiclient.NewWorkloadWorkload() // WorkloadWorkload | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkloadV1Api.StartMigration(context.Background(), oTenant).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkloadV1Api.StartMigration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartMigration`: WorkloadWorkload
    fmt.Fprintf(os.Stdout, "Response from `WorkloadV1Api.StartMigration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartMigrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**WorkloadWorkload**](WorkloadWorkload.md) |  | 

### Return type

[**WorkloadWorkload**](workloadWorkload.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartMigration1

> WorkloadWorkload StartMigration1(ctx, oName).Body(body).Execute()

Start Workload Migration operation

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
    oName := "oName_example" // string | 
    body := *openapiclient.NewWorkloadWorkload() // WorkloadWorkload | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkloadV1Api.StartMigration1(context.Background(), oName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkloadV1Api.StartMigration1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartMigration1`: WorkloadWorkload
    fmt.Fprintf(os.Stdout, "Response from `WorkloadV1Api.StartMigration1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartMigration1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**WorkloadWorkload**](WorkloadWorkload.md) |  | 

### Return type

[**WorkloadWorkload**](workloadWorkload.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkload

> WorkloadWorkload UpdateWorkload(ctx, oTenant).Body(body).Execute()

Update Workload object

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
    oTenant := "oTenant_example" // string | 
    body := *openapiclient.NewWorkloadWorkload() // WorkloadWorkload | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkloadV1Api.UpdateWorkload(context.Background(), oTenant).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkloadV1Api.UpdateWorkload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorkload`: WorkloadWorkload
    fmt.Fprintf(os.Stdout, "Response from `WorkloadV1Api.UpdateWorkload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**WorkloadWorkload**](WorkloadWorkload.md) |  | 

### Return type

[**WorkloadWorkload**](workloadWorkload.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkload1

> WorkloadWorkload UpdateWorkload1(ctx, oName).Body(body).Execute()

Update Workload object

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
    oName := "oName_example" // string | 
    body := *openapiclient.NewWorkloadWorkload() // WorkloadWorkload | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkloadV1Api.UpdateWorkload1(context.Background(), oName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkloadV1Api.UpdateWorkload1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorkload1`: WorkloadWorkload
    fmt.Fprintf(os.Stdout, "Response from `WorkloadV1Api.UpdateWorkload1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkload1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**WorkloadWorkload**](WorkloadWorkload.md) |  | 

### Return type

[**WorkloadWorkload**](workloadWorkload.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WatchEndpoint

> WorkloadAutoMsgEndpointWatchHelper WatchEndpoint(ctx, oTenant).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

Watch Endpoint objects. Supports WebSockets or HTTP long poll

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
    oName := "oName_example" // string | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. (optional)
    oCreationTime := time.Now() // time.Time | CreationTime is the creation time of the object. System generated and updated, not updatable by user. (optional)
    fieldChangeSelector := []string{"Inner_example"} // []string | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkloadV1Api.WatchEndpoint(context.Background(), oTenant).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkloadV1Api.WatchEndpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WatchEndpoint`: WorkloadAutoMsgEndpointWatchHelper
    fmt.Fprintf(os.Stdout, "Response from `WorkloadV1Api.WatchEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWatchEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**WorkloadAutoMsgEndpointWatchHelper**](workloadAutoMsgEndpointWatchHelper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WatchEndpoint1

> WorkloadAutoMsgEndpointWatchHelper WatchEndpoint1(ctx).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

Watch Endpoint objects. Supports WebSockets or HTTP long poll

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
    oName := "oName_example" // string | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. (optional)
    oCreationTime := time.Now() // time.Time | CreationTime is the creation time of the object. System generated and updated, not updatable by user. (optional)
    fieldChangeSelector := []string{"Inner_example"} // []string | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkloadV1Api.WatchEndpoint1(context.Background()).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkloadV1Api.WatchEndpoint1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WatchEndpoint1`: WorkloadAutoMsgEndpointWatchHelper
    fmt.Fprintf(os.Stdout, "Response from `WorkloadV1Api.WatchEndpoint1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWatchEndpoint1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**WorkloadAutoMsgEndpointWatchHelper**](workloadAutoMsgEndpointWatchHelper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WatchWorkload

> WorkloadAutoMsgWorkloadWatchHelper WatchWorkload(ctx, oTenant).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

Watch Workload objects. Supports WebSockets or HTTP long poll

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
    oName := "oName_example" // string | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. (optional)
    oCreationTime := time.Now() // time.Time | CreationTime is the creation time of the object. System generated and updated, not updatable by user. (optional)
    fieldChangeSelector := []string{"Inner_example"} // []string | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkloadV1Api.WatchWorkload(context.Background(), oTenant).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkloadV1Api.WatchWorkload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WatchWorkload`: WorkloadAutoMsgWorkloadWatchHelper
    fmt.Fprintf(os.Stdout, "Response from `WorkloadV1Api.WatchWorkload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWatchWorkloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**WorkloadAutoMsgWorkloadWatchHelper**](workloadAutoMsgWorkloadWatchHelper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WatchWorkload1

> WorkloadAutoMsgWorkloadWatchHelper WatchWorkload1(ctx).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

Watch Workload objects. Supports WebSockets or HTTP long poll

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
    oName := "oName_example" // string | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. (optional)
    oCreationTime := time.Now() // time.Time | CreationTime is the creation time of the object. System generated and updated, not updatable by user. (optional)
    fieldChangeSelector := []string{"Inner_example"} // []string | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkloadV1Api.WatchWorkload1(context.Background()).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkloadV1Api.WatchWorkload1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WatchWorkload1`: WorkloadAutoMsgWorkloadWatchHelper
    fmt.Fprintf(os.Stdout, "Response from `WorkloadV1Api.WatchWorkload1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWatchWorkload1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**WorkloadAutoMsgWorkloadWatchHelper**](workloadAutoMsgWorkloadWatchHelper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

