# \RoutingV1Api

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHealthZ**](RoutingV1Api.md#GetHealthZ) | **Get** /routing/v1/{Instance}/health | 
[**GetListNeighbors**](RoutingV1Api.md#GetListNeighbors) | **Get** /routing/v1/{Instance}/neighbors | 
[**GetListRoutes1**](RoutingV1Api.md#GetListRoutes1) | **Get** /routing/v1/{Instance}/routes | 
[**PostListRoutes**](RoutingV1Api.md#PostListRoutes) | **Post** /routing/v1/{Instance}/routes | 



## GetHealthZ

> RoutingHealth GetHealthZ(ctx, instance).Execute()



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
    instance := "instance_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RoutingV1Api.GetHealthZ(context.Background(), instance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingV1Api.GetHealthZ``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHealthZ`: RoutingHealth
    fmt.Fprintf(os.Stdout, "Response from `RoutingV1Api.GetHealthZ`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instance** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHealthZRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RoutingHealth**](routingHealth.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListNeighbors

> RoutingNeighborList GetListNeighbors(ctx, instance).TKind(tKind).MetaCreationTime(metaCreationTime).Neighbor(neighbor).Execute()



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
    instance := "instance_example" // string | 
    tKind := "tKind_example" // string | Kind represents the type of the API object. (optional)
    metaCreationTime := time.Now() // time.Time | CreationTime is the creation time of the object. System generated and updated, not updatable by user. (optional)
    neighbor := "neighbor_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RoutingV1Api.GetListNeighbors(context.Background(), instance).TKind(tKind).MetaCreationTime(metaCreationTime).Neighbor(neighbor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingV1Api.GetListNeighbors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetListNeighbors`: RoutingNeighborList
    fmt.Fprintf(os.Stdout, "Response from `RoutingV1Api.GetListNeighbors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instance** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListNeighborsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tKind** | **string** | Kind represents the type of the API object. | 
 **metaCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **neighbor** | **string** |  | 

### Return type

[**RoutingNeighborList**](routingNeighborList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListRoutes1

> RoutingRouteList GetListRoutes1(ctx, instance).TKind(tKind).MetaCreationTime(metaCreationTime).Ipaddress(ipaddress).PageNumber(pageNumber).Execute()



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
    instance := "instance_example" // string | 
    tKind := "tKind_example" // string | Kind represents the type of the API object. (optional)
    metaCreationTime := time.Now() // time.Time | CreationTime is the creation time of the object. System generated and updated, not updatable by user. (optional)
    ipaddress := "ipaddress_example" // string |  (optional)
    pageNumber := int64(789) // int64 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RoutingV1Api.GetListRoutes1(context.Background(), instance).TKind(tKind).MetaCreationTime(metaCreationTime).Ipaddress(ipaddress).PageNumber(pageNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingV1Api.GetListRoutes1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetListRoutes1`: RoutingRouteList
    fmt.Fprintf(os.Stdout, "Response from `RoutingV1Api.GetListRoutes1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instance** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListRoutes1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tKind** | **string** | Kind represents the type of the API object. | 
 **metaCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **ipaddress** | **string** |  | 
 **pageNumber** | **int64** |  | 

### Return type

[**RoutingRouteList**](routingRouteList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostListRoutes

> RoutingRouteList PostListRoutes(ctx, instance).Body(body).Execute()



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
    instance := "instance_example" // string | 
    body := *openapiclient.NewRoutingRouteFilter() // RoutingRouteFilter | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RoutingV1Api.PostListRoutes(context.Background(), instance).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingV1Api.PostListRoutes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostListRoutes`: RoutingRouteList
    fmt.Fprintf(os.Stdout, "Response from `RoutingV1Api.PostListRoutes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instance** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostListRoutesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RoutingRouteFilter**](RoutingRouteFilter.md) |  | 

### Return type

[**RoutingRouteList**](routingRouteList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

