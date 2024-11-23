# \PreferencesV1Api

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUIGlobalSettings**](PreferencesV1Api.md#GetUIGlobalSettings) | **Get** /configs/preferences/v1/tenant/{O.Tenant}/uiglobalsettings | Get UIGlobalSettings object
[**GetUIGlobalSettings1**](PreferencesV1Api.md#GetUIGlobalSettings1) | **Get** /configs/preferences/v1/uiglobalsettings | Get UIGlobalSettings object
[**LabelUIGlobalSettings**](PreferencesV1Api.md#LabelUIGlobalSettings) | **Post** /configs/preferences/v1/tenant/{O.Tenant}/uiglobalsettings/label | Label UIGlobalSettings object
[**LabelUIGlobalSettings1**](PreferencesV1Api.md#LabelUIGlobalSettings1) | **Post** /configs/preferences/v1/uiglobalsettings/label | Label UIGlobalSettings object
[**UpdateUIGlobalSettings**](PreferencesV1Api.md#UpdateUIGlobalSettings) | **Put** /configs/preferences/v1/tenant/{O.Tenant}/uiglobalsettings | Update UIGlobalSettings object
[**UpdateUIGlobalSettings1**](PreferencesV1Api.md#UpdateUIGlobalSettings1) | **Put** /configs/preferences/v1/uiglobalsettings | Update UIGlobalSettings object
[**WatchUIGlobalSettings**](PreferencesV1Api.md#WatchUIGlobalSettings) | **Get** /configs/preferences/v1/watch/tenant/{O.Tenant}/uiglobalsettings | Watch UIGlobalSettings objects. Supports WebSockets or HTTP long poll
[**WatchUIGlobalSettings1**](PreferencesV1Api.md#WatchUIGlobalSettings1) | **Get** /configs/preferences/v1/watch/uiglobalsettings | Watch UIGlobalSettings objects. Supports WebSockets or HTTP long poll



## GetUIGlobalSettings

> PreferencesUIGlobalSettings GetUIGlobalSettings(ctx, oTenant).TKind(tKind).MetaCreationTime(metaCreationTime).Execute()

Get UIGlobalSettings object

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
    resp, r, err := api_client.PreferencesV1Api.GetUIGlobalSettings(context.Background(), oTenant).TKind(tKind).MetaCreationTime(metaCreationTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PreferencesV1Api.GetUIGlobalSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUIGlobalSettings`: PreferencesUIGlobalSettings
    fmt.Fprintf(os.Stdout, "Response from `PreferencesV1Api.GetUIGlobalSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUIGlobalSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tKind** | **string** | Kind represents the type of the API object. | 
 **metaCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 

### Return type

[**PreferencesUIGlobalSettings**](preferencesUIGlobalSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUIGlobalSettings1

> PreferencesUIGlobalSettings GetUIGlobalSettings1(ctx).TKind(tKind).MetaCreationTime(metaCreationTime).Execute()

Get UIGlobalSettings object

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
    resp, r, err := api_client.PreferencesV1Api.GetUIGlobalSettings1(context.Background()).TKind(tKind).MetaCreationTime(metaCreationTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PreferencesV1Api.GetUIGlobalSettings1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUIGlobalSettings1`: PreferencesUIGlobalSettings
    fmt.Fprintf(os.Stdout, "Response from `PreferencesV1Api.GetUIGlobalSettings1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUIGlobalSettings1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tKind** | **string** | Kind represents the type of the API object. | 
 **metaCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 

### Return type

[**PreferencesUIGlobalSettings**](preferencesUIGlobalSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LabelUIGlobalSettings

> PreferencesUIGlobalSettings LabelUIGlobalSettings(ctx, oTenant).Body(body).Execute()

Label UIGlobalSettings object

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
    resp, r, err := api_client.PreferencesV1Api.LabelUIGlobalSettings(context.Background(), oTenant).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PreferencesV1Api.LabelUIGlobalSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LabelUIGlobalSettings`: PreferencesUIGlobalSettings
    fmt.Fprintf(os.Stdout, "Response from `PreferencesV1Api.LabelUIGlobalSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLabelUIGlobalSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ApiLabel**](ApiLabel.md) |  | 

### Return type

[**PreferencesUIGlobalSettings**](preferencesUIGlobalSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LabelUIGlobalSettings1

> PreferencesUIGlobalSettings LabelUIGlobalSettings1(ctx).Body(body).Execute()

Label UIGlobalSettings object

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
    body := *openapiclient.NewApiLabel() // ApiLabel | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PreferencesV1Api.LabelUIGlobalSettings1(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PreferencesV1Api.LabelUIGlobalSettings1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LabelUIGlobalSettings1`: PreferencesUIGlobalSettings
    fmt.Fprintf(os.Stdout, "Response from `PreferencesV1Api.LabelUIGlobalSettings1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLabelUIGlobalSettings1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApiLabel**](ApiLabel.md) |  | 

### Return type

[**PreferencesUIGlobalSettings**](preferencesUIGlobalSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUIGlobalSettings

> PreferencesUIGlobalSettings UpdateUIGlobalSettings(ctx, oTenant).Body(body).Execute()

Update UIGlobalSettings object

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
    body := *openapiclient.NewPreferencesUIGlobalSettings() // PreferencesUIGlobalSettings | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PreferencesV1Api.UpdateUIGlobalSettings(context.Background(), oTenant).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PreferencesV1Api.UpdateUIGlobalSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUIGlobalSettings`: PreferencesUIGlobalSettings
    fmt.Fprintf(os.Stdout, "Response from `PreferencesV1Api.UpdateUIGlobalSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUIGlobalSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**PreferencesUIGlobalSettings**](PreferencesUIGlobalSettings.md) |  | 

### Return type

[**PreferencesUIGlobalSettings**](preferencesUIGlobalSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUIGlobalSettings1

> PreferencesUIGlobalSettings UpdateUIGlobalSettings1(ctx).Body(body).Execute()

Update UIGlobalSettings object

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
    body := *openapiclient.NewPreferencesUIGlobalSettings() // PreferencesUIGlobalSettings | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PreferencesV1Api.UpdateUIGlobalSettings1(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PreferencesV1Api.UpdateUIGlobalSettings1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUIGlobalSettings1`: PreferencesUIGlobalSettings
    fmt.Fprintf(os.Stdout, "Response from `PreferencesV1Api.UpdateUIGlobalSettings1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUIGlobalSettings1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PreferencesUIGlobalSettings**](PreferencesUIGlobalSettings.md) |  | 

### Return type

[**PreferencesUIGlobalSettings**](preferencesUIGlobalSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WatchUIGlobalSettings

> PreferencesAutoMsgUIGlobalSettingsWatchHelper WatchUIGlobalSettings(ctx, oTenant).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

Watch UIGlobalSettings objects. Supports WebSockets or HTTP long poll

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
    resp, r, err := api_client.PreferencesV1Api.WatchUIGlobalSettings(context.Background(), oTenant).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PreferencesV1Api.WatchUIGlobalSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WatchUIGlobalSettings`: PreferencesAutoMsgUIGlobalSettingsWatchHelper
    fmt.Fprintf(os.Stdout, "Response from `PreferencesV1Api.WatchUIGlobalSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oTenant** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWatchUIGlobalSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**PreferencesAutoMsgUIGlobalSettingsWatchHelper**](preferencesAutoMsgUIGlobalSettingsWatchHelper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WatchUIGlobalSettings1

> PreferencesAutoMsgUIGlobalSettingsWatchHelper WatchUIGlobalSettings1(ctx).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

Watch UIGlobalSettings objects. Supports WebSockets or HTTP long poll

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
    resp, r, err := api_client.PreferencesV1Api.WatchUIGlobalSettings1(context.Background()).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PreferencesV1Api.WatchUIGlobalSettings1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WatchUIGlobalSettings1`: PreferencesAutoMsgUIGlobalSettingsWatchHelper
    fmt.Fprintf(os.Stdout, "Response from `PreferencesV1Api.WatchUIGlobalSettings1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWatchUIGlobalSettings1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**PreferencesAutoMsgUIGlobalSettingsWatchHelper**](preferencesAutoMsgUIGlobalSettingsWatchHelper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

