# \TokenauthV1Api

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGenerateNodeToken**](TokenauthV1Api.md#GetGenerateNodeToken) | **Get** /tokenauth/v1/node | 



## GetGenerateNodeToken

> TokenauthNodeTokenResponse GetGenerateNodeToken(ctx).Audience(audience).ValidityStart(validityStart).Execute()



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
    audience := []string{"Inner_example"} // []string | Audience represents a list of nodes the token is valid for. \"*\" indicates all nodes. (optional)
    validityStart := time.Now() // time.Time | ValidityStart indicates the time at which the token becomes valid. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TokenauthV1Api.GetGenerateNodeToken(context.Background()).Audience(audience).ValidityStart(validityStart).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenauthV1Api.GetGenerateNodeToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGenerateNodeToken`: TokenauthNodeTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `TokenauthV1Api.GetGenerateNodeToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGenerateNodeTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **audience** | **[]string** | Audience represents a list of nodes the token is valid for. \&quot;*\&quot; indicates all nodes. | 
 **validityStart** | **time.Time** | ValidityStart indicates the time at which the token becomes valid. | 

### Return type

[**TokenauthNodeTokenResponse**](tokenauthNodeTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

