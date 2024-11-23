# \OrchestrationV1Api

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOrchestrator**](OrchestrationV1Api.md#AddOrchestrator) | **Post** /configs/orchestration/v1/orchestrator | Create Orchestrator object
[**DeleteOrchestrator**](OrchestrationV1Api.md#DeleteOrchestrator) | **Delete** /configs/orchestration/v1/orchestrator/{O.Name} | Delete Orchestrator object
[**GetOrchestrator**](OrchestrationV1Api.md#GetOrchestrator) | **Get** /configs/orchestration/v1/orchestrator/{O.Name} | Get Orchestrator object
[**LabelOrchestrator**](OrchestrationV1Api.md#LabelOrchestrator) | **Post** /configs/orchestration/v1/orchestrator/{O.Name}/label | Label Orchestrator object
[**ListOrchestrator**](OrchestrationV1Api.md#ListOrchestrator) | **Get** /configs/orchestration/v1/orchestrator | List Orchestrator objects
[**UpdateOrchestrator**](OrchestrationV1Api.md#UpdateOrchestrator) | **Put** /configs/orchestration/v1/orchestrator/{O.Name} | Update Orchestrator object
[**WatchOrchestrator**](OrchestrationV1Api.md#WatchOrchestrator) | **Get** /configs/orchestration/v1/watch/orchestrator | Watch Orchestrator objects. Supports WebSockets or HTTP long poll



## AddOrchestrator

> OrchestrationOrchestrator AddOrchestrator(ctx).Body(body).Execute()

Create Orchestrator object

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
    body := *openapiclient.NewOrchestrationOrchestrator() // OrchestrationOrchestrator | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrchestrationV1Api.AddOrchestrator(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrchestrationV1Api.AddOrchestrator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddOrchestrator`: OrchestrationOrchestrator
    fmt.Fprintf(os.Stdout, "Response from `OrchestrationV1Api.AddOrchestrator`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddOrchestratorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**OrchestrationOrchestrator**](OrchestrationOrchestrator.md) |  | 

### Return type

[**OrchestrationOrchestrator**](orchestrationOrchestrator.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrchestrator

> OrchestrationOrchestrator DeleteOrchestrator(ctx, oName).Execute()

Delete Orchestrator object

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
    resp, r, err := api_client.OrchestrationV1Api.DeleteOrchestrator(context.Background(), oName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrchestrationV1Api.DeleteOrchestrator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteOrchestrator`: OrchestrationOrchestrator
    fmt.Fprintf(os.Stdout, "Response from `OrchestrationV1Api.DeleteOrchestrator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrchestratorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrchestrationOrchestrator**](orchestrationOrchestrator.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrchestrator

> OrchestrationOrchestrator GetOrchestrator(ctx, oName).TKind(tKind).MetaCreationTime(metaCreationTime).StatusConnectionStatus(statusConnectionStatus).StatusLastTransitionTime(statusLastTransitionTime).StatusDiscoveredNamespaces(statusDiscoveredNamespaces).Execute()

Get Orchestrator object

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
    statusConnectionStatus := "statusConnectionStatus_example" // string |  (optional)
    statusLastTransitionTime := time.Now() // time.Time |  (optional)
    statusDiscoveredNamespaces := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrchestrationV1Api.GetOrchestrator(context.Background(), oName).TKind(tKind).MetaCreationTime(metaCreationTime).StatusConnectionStatus(statusConnectionStatus).StatusLastTransitionTime(statusLastTransitionTime).StatusDiscoveredNamespaces(statusDiscoveredNamespaces).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrchestrationV1Api.GetOrchestrator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrchestrator`: OrchestrationOrchestrator
    fmt.Fprintf(os.Stdout, "Response from `OrchestrationV1Api.GetOrchestrator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrchestratorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tKind** | **string** | Kind represents the type of the API object. | 
 **metaCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **statusConnectionStatus** | **string** |  | 
 **statusLastTransitionTime** | **time.Time** |  | 
 **statusDiscoveredNamespaces** | **[]string** |  | 

### Return type

[**OrchestrationOrchestrator**](orchestrationOrchestrator.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LabelOrchestrator

> OrchestrationOrchestrator LabelOrchestrator(ctx, oName).Body(body).Execute()

Label Orchestrator object

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
    resp, r, err := api_client.OrchestrationV1Api.LabelOrchestrator(context.Background(), oName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrchestrationV1Api.LabelOrchestrator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LabelOrchestrator`: OrchestrationOrchestrator
    fmt.Fprintf(os.Stdout, "Response from `OrchestrationV1Api.LabelOrchestrator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLabelOrchestratorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ApiLabel**](ApiLabel.md) |  | 

### Return type

[**OrchestrationOrchestrator**](orchestrationOrchestrator.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrchestrator

> OrchestrationOrchestratorList ListOrchestrator(ctx).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

List Orchestrator objects

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
    resp, r, err := api_client.OrchestrationV1Api.ListOrchestrator(context.Background()).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrchestrationV1Api.ListOrchestrator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrchestrator`: OrchestrationOrchestratorList
    fmt.Fprintf(os.Stdout, "Response from `OrchestrationV1Api.ListOrchestrator`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOrchestratorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**OrchestrationOrchestratorList**](orchestrationOrchestratorList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrchestrator

> OrchestrationOrchestrator UpdateOrchestrator(ctx, oName).Body(body).Execute()

Update Orchestrator object

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
    body := *openapiclient.NewOrchestrationOrchestrator() // OrchestrationOrchestrator | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrchestrationV1Api.UpdateOrchestrator(context.Background(), oName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrchestrationV1Api.UpdateOrchestrator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrchestrator`: OrchestrationOrchestrator
    fmt.Fprintf(os.Stdout, "Response from `OrchestrationV1Api.UpdateOrchestrator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrchestratorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OrchestrationOrchestrator**](OrchestrationOrchestrator.md) |  | 

### Return type

[**OrchestrationOrchestrator**](orchestrationOrchestrator.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WatchOrchestrator

> OrchestrationAutoMsgOrchestratorWatchHelper WatchOrchestrator(ctx).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

Watch Orchestrator objects. Supports WebSockets or HTTP long poll

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
    resp, r, err := api_client.OrchestrationV1Api.WatchOrchestrator(context.Background()).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrchestrationV1Api.WatchOrchestrator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WatchOrchestrator`: OrchestrationAutoMsgOrchestratorWatchHelper
    fmt.Fprintf(os.Stdout, "Response from `OrchestrationV1Api.WatchOrchestrator`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWatchOrchestratorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**OrchestrationAutoMsgOrchestratorWatchHelper**](orchestrationAutoMsgOrchestratorWatchHelper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

