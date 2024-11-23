# \SearchV1Api

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPolicyQuery1**](SearchV1Api.md#GetPolicyQuery1) | **Get** /search/v1/policy-query | Security Policy Query
[**GetQuery1**](SearchV1Api.md#GetQuery1) | **Get** /search/v1/query | Structured or free-form search
[**PostPolicyQuery**](SearchV1Api.md#PostPolicyQuery) | **Post** /search/v1/policy-query | Security Policy Query
[**PostQuery**](SearchV1Api.md#PostQuery) | **Post** /search/v1/query | Structured or free-form search



## GetPolicyQuery1

> SearchPolicySearchResponse GetPolicyQuery1(ctx).Tenant(tenant).Kinds(kinds).Execute()

Security Policy Query

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
    tenant := "tenant_example" // string | Tenant Name, to perform query within a Tenant's scope. The default tenant is \"default\". In the backend this field gets auto-filled & validated by apigw-hook based on user login context. (optional)
    kinds := []string{"Inner_example"} // []string | Kind of the policy that this request should act on. It should be either NetworkSecurityPolicy or IPSecPolicy. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SearchV1Api.GetPolicyQuery1(context.Background()).Tenant(tenant).Kinds(kinds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchV1Api.GetPolicyQuery1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPolicyQuery1`: SearchPolicySearchResponse
    fmt.Fprintf(os.Stdout, "Response from `SearchV1Api.GetPolicyQuery1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyQuery1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenant** | **string** | Tenant Name, to perform query within a Tenant&#39;s scope. The default tenant is \&quot;default\&quot;. In the backend this field gets auto-filled &amp; validated by apigw-hook based on user login context. | 
 **kinds** | **[]string** | Kind of the policy that this request should act on. It should be either NetworkSecurityPolicy or IPSecPolicy. | 

### Return type

[**SearchPolicySearchResponse**](searchPolicySearchResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuery1

> SearchSearchResponse GetQuery1(ctx).QueryString(queryString).From(from).QueryCategories(queryCategories).Execute()

Structured or free-form search

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
    queryString := "queryString_example" // string | Simple query string This can be specified as URI parameter. For advanced query cases, Users should use specify SearchQuery and pass the SearchRequest in a GET/POST Body The max query-string length is 256 bytes. Length of string should be between 0 and 256. (optional)
    from := int32(56) // int32 | From represents the start offset (zero based), used in paginated search requests The results returned would be in the range [From ... From+MaxResults-1] This can be specified as URI parameter. Default value is 0 and valid range is 0..8192. Value should be between 0 and 8192. (optional)
    queryCategories := []string{"Inner_example"} // []string | OR of Categories to be matched, AND and Exclude are not supported for this type The max category string length is 64 bytes. Length of string should be between 0 and 64. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SearchV1Api.GetQuery1(context.Background()).QueryString(queryString).From(from).QueryCategories(queryCategories).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchV1Api.GetQuery1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetQuery1`: SearchSearchResponse
    fmt.Fprintf(os.Stdout, "Response from `SearchV1Api.GetQuery1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetQuery1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **queryString** | **string** | Simple query string This can be specified as URI parameter. For advanced query cases, Users should use specify SearchQuery and pass the SearchRequest in a GET/POST Body The max query-string length is 256 bytes. Length of string should be between 0 and 256. | 
 **from** | **int32** | From represents the start offset (zero based), used in paginated search requests The results returned would be in the range [From ... From+MaxResults-1] This can be specified as URI parameter. Default value is 0 and valid range is 0..8192. Value should be between 0 and 8192. | 
 **queryCategories** | **[]string** | OR of Categories to be matched, AND and Exclude are not supported for this type The max category string length is 64 bytes. Length of string should be between 0 and 64. | 

### Return type

[**SearchSearchResponse**](searchSearchResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPolicyQuery

> SearchPolicySearchResponse PostPolicyQuery(ctx).Body(body).Execute()

Security Policy Query

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
    body := *openapiclient.NewSearchPolicySearchRequest() // SearchPolicySearchRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SearchV1Api.PostPolicyQuery(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchV1Api.PostPolicyQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostPolicyQuery`: SearchPolicySearchResponse
    fmt.Fprintf(os.Stdout, "Response from `SearchV1Api.PostPolicyQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPolicyQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SearchPolicySearchRequest**](SearchPolicySearchRequest.md) |  | 

### Return type

[**SearchPolicySearchResponse**](searchPolicySearchResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostQuery

> SearchSearchResponse PostQuery(ctx).Body(body).Execute()

Structured or free-form search

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
    body := *openapiclient.NewSearchSearchRequest() // SearchSearchRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SearchV1Api.PostQuery(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchV1Api.PostQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostQuery`: SearchSearchResponse
    fmt.Fprintf(os.Stdout, "Response from `SearchV1Api.PostQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SearchSearchRequest**](SearchSearchRequest.md) |  | 

### Return type

[**SearchSearchResponse**](searchSearchResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

