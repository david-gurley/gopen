# \ObjstoreV1Api

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddObject**](ObjstoreV1Api.md#AddObject) | **Post** /objstore/v1/tenant/{O.Tenant}/{O.Namespace}/objects | Create Object object
[**AddObject1**](ObjstoreV1Api.md#AddObject1) | **Post** /objstore/v1/{O.Namespace}/objects | Create Object object
[**DeleteObject**](ObjstoreV1Api.md#DeleteObject) | **Delete** /objstore/v1/tenant/{O.Tenant}/{O.Namespace}/objects/{O.Name} | Delete Object object
[**DeleteObject1**](ObjstoreV1Api.md#DeleteObject1) | **Delete** /objstore/v1/{O.Namespace}/objects/{O.Name} | Delete Object object
[**GetDownloadFile**](ObjstoreV1Api.md#GetDownloadFile) | **Get** /objstore/v1/downloads/tenant/{O.Tenant}/{O.Namespace}/{O.Name} | Download file
[**GetDownloadFile1**](ObjstoreV1Api.md#GetDownloadFile1) | **Get** /objstore/v1/downloads/{O.Namespace}/{O.Name} | Download file
[**GetDownloadFileByPrefix**](ObjstoreV1Api.md#GetDownloadFileByPrefix) | **Get** /objstore/v1/downloads/all/tenant/{O.Tenant}/{O.Namespace}/{O.Name} | Download file by prefix
[**GetObject**](ObjstoreV1Api.md#GetObject) | **Get** /objstore/v1/tenant/{O.Tenant}/{O.Namespace}/objects/{O.Name} | Get Object object
[**GetObject1**](ObjstoreV1Api.md#GetObject1) | **Get** /objstore/v1/{O.Namespace}/objects/{O.Name} | Get Object object
[**ListObject**](ObjstoreV1Api.md#ListObject) | **Get** /objstore/v1/tenant/{O.Tenant}/{O.Namespace}/objects | List Object objects
[**ListObject1**](ObjstoreV1Api.md#ListObject1) | **Get** /objstore/v1/{O.Namespace}/objects | List Object objects
[**WatchObject**](ObjstoreV1Api.md#WatchObject) | **Get** /objstore/v1/watch/tenant/{O.Tenant}/{O.Namespace}/objects | Watch Object objects. Supports WebSockets or HTTP long poll
[**WatchObject1**](ObjstoreV1Api.md#WatchObject1) | **Get** /objstore/v1/watch/{O.Namespace}/objects | Watch Object objects. Supports WebSockets or HTTP long poll



## AddObject

> ObjstoreObject AddObject(ctx, oTenant).Body(body).Execute()

Create Object object

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
    body := *openapiclient.NewObjstoreObject() // ObjstoreObject | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjstoreV1Api.AddObject(context.Background(), oTenant).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjstoreV1Api.AddObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddObject`: ObjstoreObject
    fmt.Fprintf(os.Stdout, "Response from `ObjstoreV1Api.AddObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ObjstoreObject**](ObjstoreObject.md) |  | 

### Return type

[**ObjstoreObject**](objstoreObject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddObject1

> ObjstoreObject AddObject1(ctx, oNamespace).Body(body).Execute()

Create Object object

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
    oNamespace := "oNamespace_example" // string | 
    body := *openapiclient.NewObjstoreObject() // ObjstoreObject | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjstoreV1Api.AddObject1(context.Background(), oNamespace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjstoreV1Api.AddObject1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddObject1`: ObjstoreObject
    fmt.Fprintf(os.Stdout, "Response from `ObjstoreV1Api.AddObject1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oNamespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddObject1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ObjstoreObject**](ObjstoreObject.md) |  | 

### Return type

[**ObjstoreObject**](objstoreObject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteObject

> ObjstoreObject DeleteObject(ctx, oTenant).Execute()

Delete Object object

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
    resp, r, err := api_client.ObjstoreV1Api.DeleteObject(context.Background(), oTenant).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjstoreV1Api.DeleteObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteObject`: ObjstoreObject
    fmt.Fprintf(os.Stdout, "Response from `ObjstoreV1Api.DeleteObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ObjstoreObject**](objstoreObject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteObject1

> ObjstoreObject DeleteObject1(ctx, oNamespace).Execute()

Delete Object object

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
    oNamespace := "oNamespace_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjstoreV1Api.DeleteObject1(context.Background(), oNamespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjstoreV1Api.DeleteObject1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteObject1`: ObjstoreObject
    fmt.Fprintf(os.Stdout, "Response from `ObjstoreV1Api.DeleteObject1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oNamespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteObject1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ObjstoreObject**](objstoreObject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDownloadFile

> ObjstoreStreamChunk GetDownloadFile(ctx, oTenant).TKind(tKind).MetaCreationTime(metaCreationTime).Execute()

Download file

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjstoreV1Api.GetDownloadFile(context.Background(), oTenant).TKind(tKind).MetaCreationTime(metaCreationTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjstoreV1Api.GetDownloadFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDownloadFile`: ObjstoreStreamChunk
    fmt.Fprintf(os.Stdout, "Response from `ObjstoreV1Api.GetDownloadFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDownloadFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tKind** | **string** | Kind represents the type of the API object. | 
 **metaCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 

### Return type

[**ObjstoreStreamChunk**](objstoreStreamChunk.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDownloadFile1

> ObjstoreStreamChunk GetDownloadFile1(ctx, oNamespace).TKind(tKind).MetaCreationTime(metaCreationTime).Execute()

Download file

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
    oNamespace := "oNamespace_example" // string | 
    tKind := "tKind_example" // string | Kind represents the type of the API object. (optional)
    metaCreationTime := time.Now() // time.Time | CreationTime is the creation time of the object. System generated and updated, not updatable by user. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjstoreV1Api.GetDownloadFile1(context.Background(), oNamespace).TKind(tKind).MetaCreationTime(metaCreationTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjstoreV1Api.GetDownloadFile1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDownloadFile1`: ObjstoreStreamChunk
    fmt.Fprintf(os.Stdout, "Response from `ObjstoreV1Api.GetDownloadFile1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oNamespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDownloadFile1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tKind** | **string** | Kind represents the type of the API object. | 
 **metaCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 

### Return type

[**ObjstoreStreamChunk**](objstoreStreamChunk.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDownloadFileByPrefix

> ObjstoreStreamChunk GetDownloadFileByPrefix(ctx, oTenant).TKind(tKind).MetaCreationTime(metaCreationTime).Execute()

Download file by prefix

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjstoreV1Api.GetDownloadFileByPrefix(context.Background(), oTenant).TKind(tKind).MetaCreationTime(metaCreationTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjstoreV1Api.GetDownloadFileByPrefix``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDownloadFileByPrefix`: ObjstoreStreamChunk
    fmt.Fprintf(os.Stdout, "Response from `ObjstoreV1Api.GetDownloadFileByPrefix`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDownloadFileByPrefixRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tKind** | **string** | Kind represents the type of the API object. | 
 **metaCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 

### Return type

[**ObjstoreStreamChunk**](objstoreStreamChunk.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObject

> ObjstoreObject GetObject(ctx, oTenant).TKind(tKind).MetaCreationTime(metaCreationTime).Execute()

Get Object object

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjstoreV1Api.GetObject(context.Background(), oTenant).TKind(tKind).MetaCreationTime(metaCreationTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjstoreV1Api.GetObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetObject`: ObjstoreObject
    fmt.Fprintf(os.Stdout, "Response from `ObjstoreV1Api.GetObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tKind** | **string** | Kind represents the type of the API object. | 
 **metaCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 

### Return type

[**ObjstoreObject**](objstoreObject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObject1

> ObjstoreObject GetObject1(ctx, oNamespace).TKind(tKind).MetaCreationTime(metaCreationTime).Execute()

Get Object object

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
    oNamespace := "oNamespace_example" // string | 
    tKind := "tKind_example" // string | Kind represents the type of the API object. (optional)
    metaCreationTime := time.Now() // time.Time | CreationTime is the creation time of the object. System generated and updated, not updatable by user. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjstoreV1Api.GetObject1(context.Background(), oNamespace).TKind(tKind).MetaCreationTime(metaCreationTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjstoreV1Api.GetObject1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetObject1`: ObjstoreObject
    fmt.Fprintf(os.Stdout, "Response from `ObjstoreV1Api.GetObject1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oNamespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetObject1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tKind** | **string** | Kind represents the type of the API object. | 
 **metaCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 

### Return type

[**ObjstoreObject**](objstoreObject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListObject

> ObjstoreObjectList ListObject(ctx, oTenant).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

List Object objects

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
    resp, r, err := api_client.ObjstoreV1Api.ListObject(context.Background(), oTenant).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjstoreV1Api.ListObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListObject`: ObjstoreObjectList
    fmt.Fprintf(os.Stdout, "Response from `ObjstoreV1Api.ListObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**ObjstoreObjectList**](objstoreObjectList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListObject1

> ObjstoreObjectList ListObject1(ctx, oNamespace).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

List Object objects

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
    oNamespace := "oNamespace_example" // string | 
    oName := "oName_example" // string | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. (optional)
    oCreationTime := time.Now() // time.Time | CreationTime is the creation time of the object. System generated and updated, not updatable by user. (optional)
    fieldChangeSelector := []string{"Inner_example"} // []string | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjstoreV1Api.ListObject1(context.Background(), oNamespace).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjstoreV1Api.ListObject1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListObject1`: ObjstoreObjectList
    fmt.Fprintf(os.Stdout, "Response from `ObjstoreV1Api.ListObject1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oNamespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListObject1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**ObjstoreObjectList**](objstoreObjectList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WatchObject

> ObjstoreAutoMsgObjectWatchHelper WatchObject(ctx, oTenant).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

Watch Object objects. Supports WebSockets or HTTP long poll

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
    resp, r, err := api_client.ObjstoreV1Api.WatchObject(context.Background(), oTenant).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjstoreV1Api.WatchObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WatchObject`: ObjstoreAutoMsgObjectWatchHelper
    fmt.Fprintf(os.Stdout, "Response from `ObjstoreV1Api.WatchObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWatchObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**ObjstoreAutoMsgObjectWatchHelper**](objstoreAutoMsgObjectWatchHelper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WatchObject1

> ObjstoreAutoMsgObjectWatchHelper WatchObject1(ctx, oNamespace).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

Watch Object objects. Supports WebSockets or HTTP long poll

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
    oNamespace := "oNamespace_example" // string | 
    oName := "oName_example" // string | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. (optional)
    oCreationTime := time.Now() // time.Time | CreationTime is the creation time of the object. System generated and updated, not updatable by user. (optional)
    fieldChangeSelector := []string{"Inner_example"} // []string | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjstoreV1Api.WatchObject1(context.Background(), oNamespace).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjstoreV1Api.WatchObject1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WatchObject1`: ObjstoreAutoMsgObjectWatchHelper
    fmt.Fprintf(os.Stdout, "Response from `ObjstoreV1Api.WatchObject1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oNamespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWatchObject1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**ObjstoreAutoMsgObjectWatchHelper**](objstoreAutoMsgObjectWatchHelper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

