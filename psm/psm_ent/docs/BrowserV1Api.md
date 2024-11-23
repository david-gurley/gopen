# \BrowserV1Api

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetQuery1**](BrowserV1Api.md#GetQuery1) | **Get** /configs/browser/v1/query | 
[**GetReferences**](BrowserV1Api.md#GetReferences) | **Get** /configs/browser/v1/dependencies/** | 
[**GetReferrers**](BrowserV1Api.md#GetReferrers) | **Get** /configs/browser/v1/dependedby/** | 
[**PostQuery**](BrowserV1Api.md#PostQuery) | **Post** /configs/browser/v1/query | 



## GetQuery1

> BrowserBrowseResponseList GetQuery1(ctx).TKind(tKind).MetaCreationTime(metaCreationTime).Execute()



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
    tKind := "tKind_example" // string | Kind represents the type of the API object. (optional)
    metaCreationTime := time.Now() // time.Time | CreationTime is the creation time of the object. System generated and updated, not updatable by user. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BrowserV1Api.GetQuery1(context.Background()).TKind(tKind).MetaCreationTime(metaCreationTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrowserV1Api.GetQuery1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetQuery1`: BrowserBrowseResponseList
    fmt.Fprintf(os.Stdout, "Response from `BrowserV1Api.GetQuery1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetQuery1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tKind** | **string** | Kind represents the type of the API object. | 
 **metaCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 

### Return type

[**BrowserBrowseResponseList**](browserBrowseResponseList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReferences

> BrowserBrowseResponse GetReferences(ctx).TKind(tKind).MetaCreationTime(metaCreationTime).Execute()



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
    tKind := "tKind_example" // string | Kind represents the type of the API object. (optional)
    metaCreationTime := time.Now() // time.Time | CreationTime is the creation time of the object. System generated and updated, not updatable by user. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BrowserV1Api.GetReferences(context.Background()).TKind(tKind).MetaCreationTime(metaCreationTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrowserV1Api.GetReferences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReferences`: BrowserBrowseResponse
    fmt.Fprintf(os.Stdout, "Response from `BrowserV1Api.GetReferences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetReferencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tKind** | **string** | Kind represents the type of the API object. | 
 **metaCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 

### Return type

[**BrowserBrowseResponse**](browserBrowseResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReferrers

> BrowserBrowseResponse GetReferrers(ctx).TKind(tKind).MetaCreationTime(metaCreationTime).Execute()



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
    tKind := "tKind_example" // string | Kind represents the type of the API object. (optional)
    metaCreationTime := time.Now() // time.Time | CreationTime is the creation time of the object. System generated and updated, not updatable by user. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BrowserV1Api.GetReferrers(context.Background()).TKind(tKind).MetaCreationTime(metaCreationTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrowserV1Api.GetReferrers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReferrers`: BrowserBrowseResponse
    fmt.Fprintf(os.Stdout, "Response from `BrowserV1Api.GetReferrers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetReferrersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tKind** | **string** | Kind represents the type of the API object. | 
 **metaCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 

### Return type

[**BrowserBrowseResponse**](browserBrowseResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostQuery

> BrowserBrowseResponseList PostQuery(ctx).Body(body).Execute()



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
    body := *openapiclient.NewBrowserBrowseRequestList() // BrowserBrowseRequestList | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BrowserV1Api.PostQuery(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrowserV1Api.PostQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostQuery`: BrowserBrowseResponseList
    fmt.Fprintf(os.Stdout, "Response from `BrowserV1Api.PostQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**BrowserBrowseRequestList**](BrowserBrowseRequestList.md) |  | 

### Return type

[**BrowserBrowseResponseList**](browserBrowseResponseList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

