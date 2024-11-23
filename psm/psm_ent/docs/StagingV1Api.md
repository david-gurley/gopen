# \StagingV1Api

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddBuffer**](StagingV1Api.md#AddBuffer) | **Post** /configs/staging/v1/tenant/{O.Tenant}/buffers | Create Buffer object
[**AddBuffer1**](StagingV1Api.md#AddBuffer1) | **Post** /configs/staging/v1/buffers | Create Buffer object
[**Bulkedit**](StagingV1Api.md#Bulkedit) | **Post** /configs/staging/v1/tenant/{O.Tenant}/buffers/{O.Name}/bulkedit | Create/Update/Delete multiple objects as part of a single request
[**Bulkedit1**](StagingV1Api.md#Bulkedit1) | **Post** /configs/staging/v1/buffers/{O.Name}/bulkedit | Create/Update/Delete multiple objects as part of a single request
[**Clear**](StagingV1Api.md#Clear) | **Post** /configs/staging/v1/tenant/{O.Tenant}/buffers/{O.Name}/clear | Clear operations from a configuration buffer
[**Clear1**](StagingV1Api.md#Clear1) | **Post** /configs/staging/v1/buffers/{O.Name}/clear | Clear operations from a configuration buffer
[**Commit**](StagingV1Api.md#Commit) | **Post** /configs/staging/v1/tenant/{O.Tenant}/buffers/{O.Name}/commit | Commit a staged configuration buffer
[**Commit1**](StagingV1Api.md#Commit1) | **Post** /configs/staging/v1/buffers/{O.Name}/commit | Commit a staged configuration buffer
[**DeleteBuffer**](StagingV1Api.md#DeleteBuffer) | **Delete** /configs/staging/v1/tenant/{O.Tenant}/buffers/{O.Name} | Delete Buffer object
[**DeleteBuffer1**](StagingV1Api.md#DeleteBuffer1) | **Delete** /configs/staging/v1/buffers/{O.Name} | Delete Buffer object
[**GetBuffer**](StagingV1Api.md#GetBuffer) | **Get** /configs/staging/v1/tenant/{O.Tenant}/buffers/{O.Name} | Get Buffer object
[**GetBuffer1**](StagingV1Api.md#GetBuffer1) | **Get** /configs/staging/v1/buffers/{O.Name} | Get Buffer object
[**ListBuffer**](StagingV1Api.md#ListBuffer) | **Get** /configs/staging/v1/tenant/{O.Tenant}/buffers | List Buffer objects
[**ListBuffer1**](StagingV1Api.md#ListBuffer1) | **Get** /configs/staging/v1/buffers | List Buffer objects



## AddBuffer

> StagingBuffer AddBuffer(ctx, oTenant).Body(body).Execute()

Create Buffer object

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
    body := *openapiclient.NewStagingBuffer() // StagingBuffer | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StagingV1Api.AddBuffer(context.Background(), oTenant).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagingV1Api.AddBuffer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddBuffer`: StagingBuffer
    fmt.Fprintf(os.Stdout, "Response from `StagingV1Api.AddBuffer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddBufferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**StagingBuffer**](StagingBuffer.md) |  | 

### Return type

[**StagingBuffer**](stagingBuffer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddBuffer1

> StagingBuffer AddBuffer1(ctx).Body(body).Execute()

Create Buffer object

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
    body := *openapiclient.NewStagingBuffer() // StagingBuffer | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StagingV1Api.AddBuffer1(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagingV1Api.AddBuffer1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddBuffer1`: StagingBuffer
    fmt.Fprintf(os.Stdout, "Response from `StagingV1Api.AddBuffer1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddBuffer1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**StagingBuffer**](StagingBuffer.md) |  | 

### Return type

[**StagingBuffer**](stagingBuffer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Bulkedit

> StagingBulkEditAction Bulkedit(ctx, oTenant).Body(body).Execute()

Create/Update/Delete multiple objects as part of a single request

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
    body := *openapiclient.NewStagingBulkEditAction() // StagingBulkEditAction | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StagingV1Api.Bulkedit(context.Background(), oTenant).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagingV1Api.Bulkedit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Bulkedit`: StagingBulkEditAction
    fmt.Fprintf(os.Stdout, "Response from `StagingV1Api.Bulkedit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkeditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**StagingBulkEditAction**](StagingBulkEditAction.md) |  | 

### Return type

[**StagingBulkEditAction**](stagingBulkEditAction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Bulkedit1

> StagingBulkEditAction Bulkedit1(ctx, oName).Body(body).Execute()

Create/Update/Delete multiple objects as part of a single request

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
    body := *openapiclient.NewStagingBulkEditAction() // StagingBulkEditAction | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StagingV1Api.Bulkedit1(context.Background(), oName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagingV1Api.Bulkedit1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Bulkedit1`: StagingBulkEditAction
    fmt.Fprintf(os.Stdout, "Response from `StagingV1Api.Bulkedit1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkedit1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**StagingBulkEditAction**](StagingBulkEditAction.md) |  | 

### Return type

[**StagingBulkEditAction**](stagingBulkEditAction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Clear

> StagingClearAction Clear(ctx, oTenant).Body(body).Execute()

Clear operations from a configuration buffer

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
    body := *openapiclient.NewStagingClearAction() // StagingClearAction | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StagingV1Api.Clear(context.Background(), oTenant).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagingV1Api.Clear``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Clear`: StagingClearAction
    fmt.Fprintf(os.Stdout, "Response from `StagingV1Api.Clear`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiClearRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**StagingClearAction**](StagingClearAction.md) |  | 

### Return type

[**StagingClearAction**](stagingClearAction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Clear1

> StagingClearAction Clear1(ctx, oName).Body(body).Execute()

Clear operations from a configuration buffer

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
    body := *openapiclient.NewStagingClearAction() // StagingClearAction | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StagingV1Api.Clear1(context.Background(), oName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagingV1Api.Clear1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Clear1`: StagingClearAction
    fmt.Fprintf(os.Stdout, "Response from `StagingV1Api.Clear1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiClear1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**StagingClearAction**](StagingClearAction.md) |  | 

### Return type

[**StagingClearAction**](stagingClearAction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Commit

> StagingCommitAction Commit(ctx, oTenant).Body(body).Execute()

Commit a staged configuration buffer

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
    body := *openapiclient.NewStagingCommitAction() // StagingCommitAction | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StagingV1Api.Commit(context.Background(), oTenant).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagingV1Api.Commit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Commit`: StagingCommitAction
    fmt.Fprintf(os.Stdout, "Response from `StagingV1Api.Commit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**StagingCommitAction**](StagingCommitAction.md) |  | 

### Return type

[**StagingCommitAction**](stagingCommitAction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Commit1

> StagingCommitAction Commit1(ctx, oName).Body(body).Execute()

Commit a staged configuration buffer

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
    body := *openapiclient.NewStagingCommitAction() // StagingCommitAction | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StagingV1Api.Commit1(context.Background(), oName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagingV1Api.Commit1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Commit1`: StagingCommitAction
    fmt.Fprintf(os.Stdout, "Response from `StagingV1Api.Commit1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommit1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**StagingCommitAction**](StagingCommitAction.md) |  | 

### Return type

[**StagingCommitAction**](stagingCommitAction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBuffer

> StagingBuffer DeleteBuffer(ctx, oTenant).Execute()

Delete Buffer object

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
    resp, r, err := api_client.StagingV1Api.DeleteBuffer(context.Background(), oTenant).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagingV1Api.DeleteBuffer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteBuffer`: StagingBuffer
    fmt.Fprintf(os.Stdout, "Response from `StagingV1Api.DeleteBuffer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBufferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StagingBuffer**](stagingBuffer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBuffer1

> StagingBuffer DeleteBuffer1(ctx, oName).Execute()

Delete Buffer object

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
    resp, r, err := api_client.StagingV1Api.DeleteBuffer1(context.Background(), oName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagingV1Api.DeleteBuffer1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteBuffer1`: StagingBuffer
    fmt.Fprintf(os.Stdout, "Response from `StagingV1Api.DeleteBuffer1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBuffer1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StagingBuffer**](stagingBuffer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBuffer

> StagingBuffer GetBuffer(ctx, oTenant).TKind(tKind).MetaCreationTime(metaCreationTime).SpecContact(specContact).Execute()

Get Buffer object

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
    specContact := "specContact_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StagingV1Api.GetBuffer(context.Background(), oTenant).TKind(tKind).MetaCreationTime(metaCreationTime).SpecContact(specContact).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagingV1Api.GetBuffer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuffer`: StagingBuffer
    fmt.Fprintf(os.Stdout, "Response from `StagingV1Api.GetBuffer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBufferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tKind** | **string** | Kind represents the type of the API object. | 
 **metaCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **specContact** | **string** |  | 

### Return type

[**StagingBuffer**](stagingBuffer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBuffer1

> StagingBuffer GetBuffer1(ctx, oName).TKind(tKind).MetaCreationTime(metaCreationTime).SpecContact(specContact).Execute()

Get Buffer object

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
    specContact := "specContact_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StagingV1Api.GetBuffer1(context.Background(), oName).TKind(tKind).MetaCreationTime(metaCreationTime).SpecContact(specContact).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagingV1Api.GetBuffer1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuffer1`: StagingBuffer
    fmt.Fprintf(os.Stdout, "Response from `StagingV1Api.GetBuffer1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuffer1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tKind** | **string** | Kind represents the type of the API object. | 
 **metaCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **specContact** | **string** |  | 

### Return type

[**StagingBuffer**](stagingBuffer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBuffer

> StagingBufferList ListBuffer(ctx, oTenant).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

List Buffer objects

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
    resp, r, err := api_client.StagingV1Api.ListBuffer(context.Background(), oTenant).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagingV1Api.ListBuffer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuffer`: StagingBufferList
    fmt.Fprintf(os.Stdout, "Response from `StagingV1Api.ListBuffer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBufferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**StagingBufferList**](stagingBufferList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBuffer1

> StagingBufferList ListBuffer1(ctx).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

List Buffer objects

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
    resp, r, err := api_client.StagingV1Api.ListBuffer1(context.Background()).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagingV1Api.ListBuffer1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuffer1`: StagingBufferList
    fmt.Fprintf(os.Stdout, "Response from `StagingV1Api.ListBuffer1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBuffer1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**StagingBufferList**](stagingBufferList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

