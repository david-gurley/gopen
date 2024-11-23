# \RolloutV1Api

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRollout**](RolloutV1Api.md#CreateRollout) | **Post** /configs/rollout/v1/rollout/CreateRollout | Start Rollout operation
[**GetRollout**](RolloutV1Api.md#GetRollout) | **Get** /configs/rollout/v1/rollout/{O.Name} | Get Rollout object
[**ListRollout**](RolloutV1Api.md#ListRollout) | **Get** /configs/rollout/v1/rollout | List Rollout objects
[**RemoveRollout**](RolloutV1Api.md#RemoveRollout) | **Post** /configs/rollout/v1/rollout/RemoveRollout | Remove a Rollout
[**StopRollout**](RolloutV1Api.md#StopRollout) | **Post** /configs/rollout/v1/rollout/StopRollout | Stop a Rollout operation
[**UpdateRollout**](RolloutV1Api.md#UpdateRollout) | **Post** /configs/rollout/v1/rollout/UpdateRollout | Update Rollout configuration
[**WatchRollout**](RolloutV1Api.md#WatchRollout) | **Get** /configs/rollout/v1/watch/rollout | Watch Rollout objects. Supports WebSockets or HTTP long poll



## CreateRollout

> RolloutRollout CreateRollout(ctx).Body(body).Execute()

Start Rollout operation

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
    body := *openapiclient.NewRolloutRollout() // RolloutRollout | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RolloutV1Api.CreateRollout(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolloutV1Api.CreateRollout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRollout`: RolloutRollout
    fmt.Fprintf(os.Stdout, "Response from `RolloutV1Api.CreateRollout`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRolloutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RolloutRollout**](RolloutRollout.md) |  | 

### Return type

[**RolloutRollout**](rolloutRollout.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRollout

> RolloutRollout GetRollout(ctx, oName).TKind(tKind).MetaCreationTime(metaCreationTime).SpecStrategy(specStrategy).Execute()

Get Rollout object

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
    specStrategy := "specStrategy_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RolloutV1Api.GetRollout(context.Background(), oName).TKind(tKind).MetaCreationTime(metaCreationTime).SpecStrategy(specStrategy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolloutV1Api.GetRollout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRollout`: RolloutRollout
    fmt.Fprintf(os.Stdout, "Response from `RolloutV1Api.GetRollout`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRolloutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tKind** | **string** | Kind represents the type of the API object. | 
 **metaCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **specStrategy** | **string** |  | 

### Return type

[**RolloutRollout**](rolloutRollout.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRollout

> RolloutRolloutList ListRollout(ctx).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

List Rollout objects

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
    resp, r, err := api_client.RolloutV1Api.ListRollout(context.Background()).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolloutV1Api.ListRollout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRollout`: RolloutRolloutList
    fmt.Fprintf(os.Stdout, "Response from `RolloutV1Api.ListRollout`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRolloutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**RolloutRolloutList**](rolloutRolloutList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveRollout

> RolloutRollout RemoveRollout(ctx).Body(body).Execute()

Remove a Rollout

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
    body := *openapiclient.NewRolloutRollout() // RolloutRollout | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RolloutV1Api.RemoveRollout(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolloutV1Api.RemoveRollout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveRollout`: RolloutRollout
    fmt.Fprintf(os.Stdout, "Response from `RolloutV1Api.RemoveRollout`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveRolloutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RolloutRollout**](RolloutRollout.md) |  | 

### Return type

[**RolloutRollout**](rolloutRollout.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopRollout

> RolloutRollout StopRollout(ctx).Body(body).Execute()

Stop a Rollout operation

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
    body := *openapiclient.NewRolloutRollout() // RolloutRollout | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RolloutV1Api.StopRollout(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolloutV1Api.StopRollout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StopRollout`: RolloutRollout
    fmt.Fprintf(os.Stdout, "Response from `RolloutV1Api.StopRollout`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStopRolloutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RolloutRollout**](RolloutRollout.md) |  | 

### Return type

[**RolloutRollout**](rolloutRollout.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRollout

> RolloutRollout UpdateRollout(ctx).Body(body).Execute()

Update Rollout configuration

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
    body := *openapiclient.NewRolloutRollout() // RolloutRollout | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RolloutV1Api.UpdateRollout(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolloutV1Api.UpdateRollout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRollout`: RolloutRollout
    fmt.Fprintf(os.Stdout, "Response from `RolloutV1Api.UpdateRollout`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRolloutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RolloutRollout**](RolloutRollout.md) |  | 

### Return type

[**RolloutRollout**](rolloutRollout.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WatchRollout

> RolloutAutoMsgRolloutWatchHelper WatchRollout(ctx).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

Watch Rollout objects. Supports WebSockets or HTTP long poll

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
    resp, r, err := api_client.RolloutV1Api.WatchRollout(context.Background()).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolloutV1Api.WatchRollout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WatchRollout`: RolloutAutoMsgRolloutWatchHelper
    fmt.Fprintf(os.Stdout, "Response from `RolloutV1Api.WatchRollout`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWatchRolloutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**RolloutAutoMsgRolloutWatchHelper**](rolloutAutoMsgRolloutWatchHelper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

