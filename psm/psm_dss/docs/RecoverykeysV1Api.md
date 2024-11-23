# \RecoverykeysV1Api

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDownloadRecoveryKeys**](RecoverykeysV1Api.md#GetDownloadRecoveryKeys) | **Get** /sysruntime/v1/cluster/recoverykeys | 



## GetDownloadRecoveryKeys

> RecoverykeysRecoveryKeys GetDownloadRecoveryKeys(ctx).Execute()



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RecoverykeysV1Api.GetDownloadRecoveryKeys(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecoverykeysV1Api.GetDownloadRecoveryKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDownloadRecoveryKeys`: RecoverykeysRecoveryKeys
    fmt.Fprintf(os.Stdout, "Response from `RecoverykeysV1Api.GetDownloadRecoveryKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDownloadRecoveryKeysRequest struct via the builder pattern


### Return type

[**RecoverykeysRecoveryKeys**](recoverykeysRecoveryKeys.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

