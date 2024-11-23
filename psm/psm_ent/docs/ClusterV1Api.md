# \ClusterV1Api

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddConfigurationSnapshot**](ClusterV1Api.md#AddConfigurationSnapshot) | **Post** /configs/cluster/v1/config-snapshot | Create ConfigurationSnapshot object
[**AddDSCProfile**](ClusterV1Api.md#AddDSCProfile) | **Post** /configs/cluster/v1/dscprofiles | Create DSCProfile object
[**AddHost**](ClusterV1Api.md#AddHost) | **Post** /configs/cluster/v1/hosts | Create Host object
[**AddLicense**](ClusterV1Api.md#AddLicense) | **Post** /configs/cluster/v1/licenses | Create License object
[**AddNode**](ClusterV1Api.md#AddNode) | **Post** /configs/cluster/v1/nodes | Create Node object
[**AddTenant**](ClusterV1Api.md#AddTenant) | **Post** /configs/cluster/v1/tenants | Create Tenant object
[**AuthBootstrapComplete**](ClusterV1Api.md#AuthBootstrapComplete) | **Post** /configs/cluster/v1/cluster/AuthBootstrapComplete | Mark bootstrapping as complete for the cluster
[**DeleteConfigurationSnapshot**](ClusterV1Api.md#DeleteConfigurationSnapshot) | **Delete** /configs/cluster/v1/config-snapshot | Delete ConfigurationSnapshot object
[**DeleteDSCProfile**](ClusterV1Api.md#DeleteDSCProfile) | **Delete** /configs/cluster/v1/dscprofiles/{O.Name} | Delete DSCProfile object
[**DeleteDistributedServiceCard**](ClusterV1Api.md#DeleteDistributedServiceCard) | **Delete** /configs/cluster/v1/distributedservicecards/{O.Name} | Delete DistributedServiceCard object
[**DeleteHost**](ClusterV1Api.md#DeleteHost) | **Delete** /configs/cluster/v1/hosts/{O.Name} | Delete Host object
[**DeleteNode**](ClusterV1Api.md#DeleteNode) | **Delete** /configs/cluster/v1/nodes/{O.Name} | Delete Node object
[**DeleteTenant**](ClusterV1Api.md#DeleteTenant) | **Delete** /configs/cluster/v1/tenants/{O.Name} | Delete Tenant object
[**GetCluster**](ClusterV1Api.md#GetCluster) | **Get** /configs/cluster/v1/cluster | Get Cluster object
[**GetConfigurationSnapshot**](ClusterV1Api.md#GetConfigurationSnapshot) | **Get** /configs/cluster/v1/config-snapshot | Get ConfigurationSnapshot object
[**GetDSCProfile**](ClusterV1Api.md#GetDSCProfile) | **Get** /configs/cluster/v1/dscprofiles/{O.Name} | Get DSCProfile object
[**GetDistributedServiceCard**](ClusterV1Api.md#GetDistributedServiceCard) | **Get** /configs/cluster/v1/distributedservicecards/{O.Name} | Get DistributedServiceCard object
[**GetHost**](ClusterV1Api.md#GetHost) | **Get** /configs/cluster/v1/hosts/{O.Name} | Get Host object
[**GetLicense**](ClusterV1Api.md#GetLicense) | **Get** /configs/cluster/v1/licenses | Get License object
[**GetNode**](ClusterV1Api.md#GetNode) | **Get** /configs/cluster/v1/nodes/{O.Name} | Get Node object
[**GetSnapshotRestore**](ClusterV1Api.md#GetSnapshotRestore) | **Get** /configs/cluster/v1/config-restore | Get SnapshotRestore object
[**GetTenant**](ClusterV1Api.md#GetTenant) | **Get** /configs/cluster/v1/tenants/{O.Name} | Get Tenant object
[**GetVersion**](ClusterV1Api.md#GetVersion) | **Get** /configs/cluster/v1/version | Get Version object
[**LabelCluster**](ClusterV1Api.md#LabelCluster) | **Post** /configs/cluster/v1/cluster/label | Label Cluster object
[**LabelConfigurationSnapshot**](ClusterV1Api.md#LabelConfigurationSnapshot) | **Post** /configs/cluster/v1/config-snapshot/label | Label ConfigurationSnapshot object
[**LabelDSCProfile**](ClusterV1Api.md#LabelDSCProfile) | **Post** /configs/cluster/v1/dscprofiles/{O.Name}/label | Label DSCProfile object
[**LabelDistributedServiceCard**](ClusterV1Api.md#LabelDistributedServiceCard) | **Post** /configs/cluster/v1/distributedservicecards/{O.Name}/label | Label DistributedServiceCard object
[**LabelHost**](ClusterV1Api.md#LabelHost) | **Post** /configs/cluster/v1/hosts/{O.Name}/label | Label Host object
[**LabelLicense**](ClusterV1Api.md#LabelLicense) | **Post** /configs/cluster/v1/licenses/label | Label License object
[**LabelNode**](ClusterV1Api.md#LabelNode) | **Post** /configs/cluster/v1/nodes/{O.Name}/label | Label Node object
[**LabelTenant**](ClusterV1Api.md#LabelTenant) | **Post** /configs/cluster/v1/tenants/{O.Name}/label | Label Tenant object
[**ListDSCProfile**](ClusterV1Api.md#ListDSCProfile) | **Get** /configs/cluster/v1/dscprofiles | List DSCProfile objects
[**ListDistributedServiceCard**](ClusterV1Api.md#ListDistributedServiceCard) | **Get** /configs/cluster/v1/distributedservicecards | List DistributedServiceCard objects
[**ListHost**](ClusterV1Api.md#ListHost) | **Get** /configs/cluster/v1/hosts | List Host objects
[**ListNode**](ClusterV1Api.md#ListNode) | **Get** /configs/cluster/v1/nodes | List Node objects
[**ListTenant**](ClusterV1Api.md#ListTenant) | **Get** /configs/cluster/v1/tenants | List Tenant objects
[**Restore**](ClusterV1Api.md#Restore) | **Post** /configs/cluster/v1/config-restore/restore | Restore Configuration
[**Save**](ClusterV1Api.md#Save) | **Post** /configs/cluster/v1/config-snapshot/save | Perform a Configuation Snapshot
[**UpdateCluster**](ClusterV1Api.md#UpdateCluster) | **Put** /configs/cluster/v1/cluster | Update Cluster object
[**UpdateConfigurationSnapshot**](ClusterV1Api.md#UpdateConfigurationSnapshot) | **Put** /configs/cluster/v1/config-snapshot | Update ConfigurationSnapshot object
[**UpdateDSCProfile**](ClusterV1Api.md#UpdateDSCProfile) | **Put** /configs/cluster/v1/dscprofiles/{O.Name} | Update DSCProfile object
[**UpdateDistributedServiceCard**](ClusterV1Api.md#UpdateDistributedServiceCard) | **Put** /configs/cluster/v1/distributedservicecards/{O.Name} | Update DistributedServiceCard object
[**UpdateHost**](ClusterV1Api.md#UpdateHost) | **Put** /configs/cluster/v1/hosts/{O.Name} | Update Host object
[**UpdateLicense**](ClusterV1Api.md#UpdateLicense) | **Put** /configs/cluster/v1/licenses | Update License object
[**UpdateNode**](ClusterV1Api.md#UpdateNode) | **Put** /configs/cluster/v1/nodes/{O.Name} | Update Node object
[**UpdateTLSConfig**](ClusterV1Api.md#UpdateTLSConfig) | **Post** /configs/cluster/v1/cluster/UpdateTLSConfig | Update TLS Configuration for cluster
[**UpdateTenant**](ClusterV1Api.md#UpdateTenant) | **Put** /configs/cluster/v1/tenants/{O.Name} | Update Tenant object
[**UpdateVersion**](ClusterV1Api.md#UpdateVersion) | **Put** /configs/cluster/v1/version | Update Version object
[**WatchCluster**](ClusterV1Api.md#WatchCluster) | **Get** /configs/cluster/v1/watch/cluster | Watch Cluster objects. Supports WebSockets or HTTP long poll
[**WatchConfigurationSnapshot**](ClusterV1Api.md#WatchConfigurationSnapshot) | **Get** /configs/cluster/v1/watch/config-snapshot | Watch ConfigurationSnapshot objects. Supports WebSockets or HTTP long poll
[**WatchDSCProfile**](ClusterV1Api.md#WatchDSCProfile) | **Get** /configs/cluster/v1/watch/dscprofiles | Watch DSCProfile objects. Supports WebSockets or HTTP long poll
[**WatchDistributedServiceCard**](ClusterV1Api.md#WatchDistributedServiceCard) | **Get** /configs/cluster/v1/watch/distributedservicecards | Watch DistributedServiceCard objects. Supports WebSockets or HTTP long poll
[**WatchHost**](ClusterV1Api.md#WatchHost) | **Get** /configs/cluster/v1/watch/hosts | Watch Host objects. Supports WebSockets or HTTP long poll
[**WatchNode**](ClusterV1Api.md#WatchNode) | **Get** /configs/cluster/v1/watch/nodes | Watch Node objects. Supports WebSockets or HTTP long poll
[**WatchTenant**](ClusterV1Api.md#WatchTenant) | **Get** /configs/cluster/v1/watch/tenants | Watch Tenant objects. Supports WebSockets or HTTP long poll
[**WatchVersion**](ClusterV1Api.md#WatchVersion) | **Get** /configs/cluster/v1/watch/version | Watch Version objects. Supports WebSockets or HTTP long poll



## AddConfigurationSnapshot

> ClusterConfigurationSnapshot AddConfigurationSnapshot(ctx).Body(body).Execute()

Create ConfigurationSnapshot object

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
    body := *openapiclient.NewClusterConfigurationSnapshot() // ClusterConfigurationSnapshot | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterV1Api.AddConfigurationSnapshot(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV1Api.AddConfigurationSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddConfigurationSnapshot`: ClusterConfigurationSnapshot
    fmt.Fprintf(os.Stdout, "Response from `ClusterV1Api.AddConfigurationSnapshot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddConfigurationSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ClusterConfigurationSnapshot**](ClusterConfigurationSnapshot.md) |  | 

### Return type

[**ClusterConfigurationSnapshot**](clusterConfigurationSnapshot.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddDSCProfile

> ClusterDSCProfile AddDSCProfile(ctx).Body(body).Execute()

Create DSCProfile object

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
    body := *openapiclient.NewClusterDSCProfile() // ClusterDSCProfile | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterV1Api.AddDSCProfile(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV1Api.AddDSCProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddDSCProfile`: ClusterDSCProfile
    fmt.Fprintf(os.Stdout, "Response from `ClusterV1Api.AddDSCProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddDSCProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ClusterDSCProfile**](ClusterDSCProfile.md) |  | 

### Return type

[**ClusterDSCProfile**](clusterDSCProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddHost

> ClusterHost AddHost(ctx).Body(body).Execute()

Create Host object

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
    body := *openapiclient.NewClusterHost() // ClusterHost | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterV1Api.AddHost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV1Api.AddHost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddHost`: ClusterHost
    fmt.Fprintf(os.Stdout, "Response from `ClusterV1Api.AddHost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ClusterHost**](ClusterHost.md) |  | 

### Return type

[**ClusterHost**](clusterHost.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddLicense

> ClusterLicense AddLicense(ctx).Body(body).Execute()

Create License object

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
    body := *openapiclient.NewClusterLicense() // ClusterLicense | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterV1Api.AddLicense(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV1Api.AddLicense``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddLicense`: ClusterLicense
    fmt.Fprintf(os.Stdout, "Response from `ClusterV1Api.AddLicense`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddLicenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ClusterLicense**](ClusterLicense.md) |  | 

### Return type

[**ClusterLicense**](clusterLicense.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddNode

> ClusterNode AddNode(ctx).Body(body).Execute()

Create Node object

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
    body := *openapiclient.NewClusterNode() // ClusterNode | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterV1Api.AddNode(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV1Api.AddNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddNode`: ClusterNode
    fmt.Fprintf(os.Stdout, "Response from `ClusterV1Api.AddNode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ClusterNode**](ClusterNode.md) |  | 

### Return type

[**ClusterNode**](clusterNode.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddTenant

> ClusterTenant AddTenant(ctx).Body(body).Execute()

Create Tenant object

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
    body := *openapiclient.NewClusterTenant() // ClusterTenant | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterV1Api.AddTenant(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV1Api.AddTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddTenant`: ClusterTenant
    fmt.Fprintf(os.Stdout, "Response from `ClusterV1Api.AddTenant`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ClusterTenant**](ClusterTenant.md) |  | 

### Return type

[**ClusterTenant**](clusterTenant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthBootstrapComplete

> ClusterCluster AuthBootstrapComplete(ctx).Body(body).Execute()

Mark bootstrapping as complete for the cluster

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
    body := *openapiclient.NewClusterClusterAuthBootstrapRequest() // ClusterClusterAuthBootstrapRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterV1Api.AuthBootstrapComplete(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV1Api.AuthBootstrapComplete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthBootstrapComplete`: ClusterCluster
    fmt.Fprintf(os.Stdout, "Response from `ClusterV1Api.AuthBootstrapComplete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthBootstrapCompleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ClusterClusterAuthBootstrapRequest**](ClusterClusterAuthBootstrapRequest.md) |  | 

### Return type

[**ClusterCluster**](clusterCluster.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConfigurationSnapshot

> ClusterConfigurationSnapshot DeleteConfigurationSnapshot(ctx).Execute()

Delete ConfigurationSnapshot object

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
    resp, r, err := api_client.ClusterV1Api.DeleteConfigurationSnapshot(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV1Api.DeleteConfigurationSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteConfigurationSnapshot`: ClusterConfigurationSnapshot
    fmt.Fprintf(os.Stdout, "Response from `ClusterV1Api.DeleteConfigurationSnapshot`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConfigurationSnapshotRequest struct via the builder pattern


### Return type

[**ClusterConfigurationSnapshot**](clusterConfigurationSnapshot.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDSCProfile

> ClusterDSCProfile DeleteDSCProfile(ctx, oName).Execute()

Delete DSCProfile object

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
    resp, r, err := api_client.ClusterV1Api.DeleteDSCProfile(context.Background(), oName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV1Api.DeleteDSCProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDSCProfile`: ClusterDSCProfile
    fmt.Fprintf(os.Stdout, "Response from `ClusterV1Api.DeleteDSCProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDSCProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClusterDSCProfile**](clusterDSCProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDistributedServiceCard

> ClusterDistributedServiceCard DeleteDistributedServiceCard(ctx, oName).Execute()

Delete DistributedServiceCard object

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
    resp, r, err := api_client.ClusterV1Api.DeleteDistributedServiceCard(context.Background(), oName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV1Api.DeleteDistributedServiceCard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDistributedServiceCard`: ClusterDistributedServiceCard
    fmt.Fprintf(os.Stdout, "Response from `ClusterV1Api.DeleteDistributedServiceCard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDistributedServiceCardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClusterDistributedServiceCard**](clusterDistributedServiceCard.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteHost

> ClusterHost DeleteHost(ctx, oName).Execute()

Delete Host object

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
    resp, r, err := api_client.ClusterV1Api.DeleteHost(context.Background(), oName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV1Api.DeleteHost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteHost`: ClusterHost
    fmt.Fprintf(os.Stdout, "Response from `ClusterV1Api.DeleteHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClusterHost**](clusterHost.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNode

> ClusterNode DeleteNode(ctx, oName).Execute()

Delete Node object

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
    resp, r, err := api_client.ClusterV1Api.DeleteNode(context.Background(), oName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV1Api.DeleteNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteNode`: ClusterNode
    fmt.Fprintf(os.Stdout, "Response from `ClusterV1Api.DeleteNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClusterNode**](clusterNode.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTenant

> ClusterTenant DeleteTenant(ctx, oName).Execute()

Delete Tenant object

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
    resp, r, err := api_client.ClusterV1Api.DeleteTenant(context.Background(), oName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV1Api.DeleteTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteTenant`: ClusterTenant
    fmt.Fprintf(os.Stdout, "Response from `ClusterV1Api.DeleteTenant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClusterTenant**](clusterTenant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCluster

> ClusterCluster GetCluster(ctx).TKind(tKind).MetaCreationTime(metaCreationTime).SpecQuorumNodes(specQuorumNodes).Execute()

Get Cluster object

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
    specQuorumNodes := []string{"Inner_example"} // []string | QuorumNodes contains the list of hostnames for nodes configured to be quorum nodes in the cluster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterV1Api.GetCluster(context.Background()).TKind(tKind).MetaCreationTime(metaCreationTime).SpecQuorumNodes(specQuorumNodes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV1Api.GetCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCluster`: ClusterCluster
    fmt.Fprintf(os.Stdout, "Response from `ClusterV1Api.GetCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tKind** | **string** | Kind represents the type of the API object. | 
 **metaCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **specQuorumNodes** | **[]string** | QuorumNodes contains the list of hostnames for nodes configured to be quorum nodes in the cluster. | 

### Return type

[**ClusterCluster**](clusterCluster.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfigurationSnapshot

> ClusterConfigurationSnapshot GetConfigurationSnapshot(ctx).TKind(tKind).MetaCreationTime(metaCreationTime).DestinationType(destinationType).Execute()

Get ConfigurationSnapshot object

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
    destinationType := "destinationType_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterV1Api.GetConfigurationSnapshot(context.Background()).TKind(tKind).MetaCreationTime(metaCreationTime).DestinationType(destinationType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV1Api.GetConfigurationSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfigurationSnapshot`: ClusterConfigurationSnapshot
    fmt.Fprintf(os.Stdout, "Response from `ClusterV1Api.GetConfigurationSnapshot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigurationSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tKind** | **string** | Kind represents the type of the API object. | 
 **metaCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **destinationType** | **string** |  | 

### Return type

[**ClusterConfigurationSnapshot**](clusterConfigurationSnapshot.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDSCProfile

> ClusterDSCProfile GetDSCProfile(ctx, oName).TKind(tKind).MetaCreationTime(metaCreationTime).SpecDeploymentTarget(specDeploymentTarget).InterfaceProfileNumPfs(interfaceProfileNumPfs).PropagationStatusPendingDscs(propagationStatusPendingDscs).Execute()

Get DSCProfile object

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
    specDeploymentTarget := "specDeploymentTarget_example" // string |  (optional)
    interfaceProfileNumPfs := int64(789) // int64 |  (optional)
    propagationStatusPendingDscs := []string{"Inner_example"} // []string | list of DSCs where propagation did not complete. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterV1Api.GetDSCProfile(context.Background(), oName).TKind(tKind).MetaCreationTime(metaCreationTime).SpecDeploymentTarget(specDeploymentTarget).InterfaceProfileNumPfs(interfaceProfileNumPfs).PropagationStatusPendingDscs(propagationStatusPendingDscs).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV1Api.GetDSCProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDSCProfile`: ClusterDSCProfile
    fmt.Fprintf(os.Stdout, "Response from `ClusterV1Api.GetDSCProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDSCProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tKind** | **string** | Kind represents the type of the API object. | 
 **metaCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **specDeploymentTarget** | **string** |  | 
 **interfaceProfileNumPfs** | **int64** |  | 
 **propagationStatusPendingDscs** | **[]string** | list of DSCs where propagation did not complete. | 

### Return type

[**ClusterDSCProfile**](clusterDSCProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDistributedServiceCard

> ClusterDistributedServiceCard GetDistributedServiceCard(ctx, oName).TKind(tKind).MetaCreationTime(metaCreationTime).IpConfigDnsServers(ipConfigDnsServers).SpecDscprofile(specDscprofile).Execute()

Get DistributedServiceCard object

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
    ipConfigDnsServers := []string{"Inner_example"} // []string | DNSServers contains a list of DNS Servers that can be used on DistributedServiceCard. (optional)
    specDscprofile := "specDscprofile_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterV1Api.GetDistributedServiceCard(context.Background(), oName).TKind(tKind).MetaCreationTime(metaCreationTime).IpConfigDnsServers(ipConfigDnsServers).SpecDscprofile(specDscprofile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV1Api.GetDistributedServiceCard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDistributedServiceCard`: ClusterDistributedServiceCard
    fmt.Fprintf(os.Stdout, "Response from `ClusterV1Api.GetDistributedServiceCard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDistributedServiceCardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tKind** | **string** | Kind represents the type of the API object. | 
 **metaCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **ipConfigDnsServers** | **[]string** | DNSServers contains a list of DNS Servers that can be used on DistributedServiceCard. | 
 **specDscprofile** | **string** |  | 

### Return type

[**ClusterDistributedServiceCard**](clusterDistributedServiceCard.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHost

> ClusterHost GetHost(ctx, oName).TKind(tKind).MetaCreationTime(metaCreationTime).StatusAdmittedDscs(statusAdmittedDscs).Execute()

Get Host object

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
    statusAdmittedDscs := []string{"Inner_example"} // []string | AdmittedDSCs contains a list of admitted DistributedServiceCards that are on this host. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterV1Api.GetHost(context.Background(), oName).TKind(tKind).MetaCreationTime(metaCreationTime).StatusAdmittedDscs(statusAdmittedDscs).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV1Api.GetHost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHost`: ClusterHost
    fmt.Fprintf(os.Stdout, "Response from `ClusterV1Api.GetHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tKind** | **string** | Kind represents the type of the API object. | 
 **metaCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **statusAdmittedDscs** | **[]string** | AdmittedDSCs contains a list of admitted DistributedServiceCards that are on this host. | 

### Return type

[**ClusterHost**](clusterHost.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLicense

> ClusterLicense GetLicense(ctx).TKind(tKind).MetaCreationTime(metaCreationTime).StatusUnknown(statusUnknown).Execute()

Get License object

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
    statusUnknown := []string{"Inner_example"} // []string | Licenses that are not understood by the current running version of software. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterV1Api.GetLicense(context.Background()).TKind(tKind).MetaCreationTime(metaCreationTime).StatusUnknown(statusUnknown).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV1Api.GetLicense``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLicense`: ClusterLicense
    fmt.Fprintf(os.Stdout, "Response from `ClusterV1Api.GetLicense`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLicenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tKind** | **string** | Kind represents the type of the API object. | 
 **metaCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **statusUnknown** | **[]string** | Licenses that are not understood by the current running version of software. | 

### Return type

[**ClusterLicense**](clusterLicense.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNode

> ClusterNode GetNode(ctx, oName).TKind(tKind).MetaCreationTime(metaCreationTime).Execute()

Get Node object

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterV1Api.GetNode(context.Background(), oName).TKind(tKind).MetaCreationTime(metaCreationTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV1Api.GetNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNode`: ClusterNode
    fmt.Fprintf(os.Stdout, "Response from `ClusterV1Api.GetNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tKind** | **string** | Kind represents the type of the API object. | 
 **metaCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 

### Return type

[**ClusterNode**](clusterNode.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSnapshotRestore

> ClusterSnapshotRestore GetSnapshotRestore(ctx).TKind(tKind).MetaCreationTime(metaCreationTime).SpecPath(specPath).StatusStartTime(statusStartTime).Execute()

Get SnapshotRestore object

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
    specPath := "specPath_example" // string |  (optional)
    statusStartTime := time.Now() // time.Time |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterV1Api.GetSnapshotRestore(context.Background()).TKind(tKind).MetaCreationTime(metaCreationTime).SpecPath(specPath).StatusStartTime(statusStartTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV1Api.GetSnapshotRestore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSnapshotRestore`: ClusterSnapshotRestore
    fmt.Fprintf(os.Stdout, "Response from `ClusterV1Api.GetSnapshotRestore`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSnapshotRestoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tKind** | **string** | Kind represents the type of the API object. | 
 **metaCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **specPath** | **string** |  | 
 **statusStartTime** | **time.Time** |  | 

### Return type

[**ClusterSnapshotRestore**](clusterSnapshotRestore.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTenant

> ClusterTenant GetTenant(ctx, oName).TKind(tKind).MetaCreationTime(metaCreationTime).Execute()

Get Tenant object

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterV1Api.GetTenant(context.Background(), oName).TKind(tKind).MetaCreationTime(metaCreationTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV1Api.GetTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTenant`: ClusterTenant
    fmt.Fprintf(os.Stdout, "Response from `ClusterV1Api.GetTenant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tKind** | **string** | Kind represents the type of the API object. | 
 **metaCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 

### Return type

[**ClusterTenant**](clusterTenant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVersion

> ClusterVersion GetVersion(ctx).TKind(tKind).MetaCreationTime(metaCreationTime).Execute()

Get Version object

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
    resp, r, err := api_client.ClusterV1Api.GetVersion(context.Background()).TKind(tKind).MetaCreationTime(metaCreationTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV1Api.GetVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVersion`: ClusterVersion
    fmt.Fprintf(os.Stdout, "Response from `ClusterV1Api.GetVersion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tKind** | **string** | Kind represents the type of the API object. | 
 **metaCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 

### Return type

[**ClusterVersion**](clusterVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LabelCluster

> ClusterCluster LabelCluster(ctx).Body(body).Execute()

Label Cluster object

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
    resp, r, err := api_client.ClusterV1Api.LabelCluster(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV1Api.LabelCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LabelCluster`: ClusterCluster
    fmt.Fprintf(os.Stdout, "Response from `ClusterV1Api.LabelCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLabelClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApiLabel**](ApiLabel.md) |  | 

### Return type

[**ClusterCluster**](clusterCluster.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LabelConfigurationSnapshot

> ClusterConfigurationSnapshot LabelConfigurationSnapshot(ctx).Body(body).Execute()

Label ConfigurationSnapshot object

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
    resp, r, err := api_client.ClusterV1Api.LabelConfigurationSnapshot(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV1Api.LabelConfigurationSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LabelConfigurationSnapshot`: ClusterConfigurationSnapshot
    fmt.Fprintf(os.Stdout, "Response from `ClusterV1Api.LabelConfigurationSnapshot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLabelConfigurationSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApiLabel**](ApiLabel.md) |  | 

### Return type

[**ClusterConfigurationSnapshot**](clusterConfigurationSnapshot.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LabelDSCProfile

> ClusterDSCProfile LabelDSCProfile(ctx, oName).Body(body).Execute()

Label DSCProfile object

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
    resp, r, err := api_client.ClusterV1Api.LabelDSCProfile(context.Background(), oName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV1Api.LabelDSCProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LabelDSCProfile`: ClusterDSCProfile
    fmt.Fprintf(os.Stdout, "Response from `ClusterV1Api.LabelDSCProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLabelDSCProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ApiLabel**](ApiLabel.md) |  | 

### Return type

[**ClusterDSCProfile**](clusterDSCProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LabelDistributedServiceCard

> ClusterDistributedServiceCard LabelDistributedServiceCard(ctx, oName).Body(body).Execute()

Label DistributedServiceCard object

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
    resp, r, err := api_client.ClusterV1Api.LabelDistributedServiceCard(context.Background(), oName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV1Api.LabelDistributedServiceCard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LabelDistributedServiceCard`: ClusterDistributedServiceCard
    fmt.Fprintf(os.Stdout, "Response from `ClusterV1Api.LabelDistributedServiceCard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLabelDistributedServiceCardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ApiLabel**](ApiLabel.md) |  | 

### Return type

[**ClusterDistributedServiceCard**](clusterDistributedServiceCard.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LabelHost

> ClusterHost LabelHost(ctx, oName).Body(body).Execute()

Label Host object

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
    resp, r, err := api_client.ClusterV1Api.LabelHost(context.Background(), oName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV1Api.LabelHost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LabelHost`: ClusterHost
    fmt.Fprintf(os.Stdout, "Response from `ClusterV1Api.LabelHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLabelHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ApiLabel**](ApiLabel.md) |  | 

### Return type

[**ClusterHost**](clusterHost.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LabelLicense

> ClusterLicense LabelLicense(ctx).Body(body).Execute()

Label License object

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
    resp, r, err := api_client.ClusterV1Api.LabelLicense(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV1Api.LabelLicense``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LabelLicense`: ClusterLicense
    fmt.Fprintf(os.Stdout, "Response from `ClusterV1Api.LabelLicense`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLabelLicenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApiLabel**](ApiLabel.md) |  | 

### Return type

[**ClusterLicense**](clusterLicense.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LabelNode

> ClusterNode LabelNode(ctx, oName).Body(body).Execute()

Label Node object

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
    resp, r, err := api_client.ClusterV1Api.LabelNode(context.Background(), oName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV1Api.LabelNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LabelNode`: ClusterNode
    fmt.Fprintf(os.Stdout, "Response from `ClusterV1Api.LabelNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLabelNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ApiLabel**](ApiLabel.md) |  | 

### Return type

[**ClusterNode**](clusterNode.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LabelTenant

> ClusterTenant LabelTenant(ctx, oName).Body(body).Execute()

Label Tenant object

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
    resp, r, err := api_client.ClusterV1Api.LabelTenant(context.Background(), oName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV1Api.LabelTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LabelTenant`: ClusterTenant
    fmt.Fprintf(os.Stdout, "Response from `ClusterV1Api.LabelTenant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLabelTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ApiLabel**](ApiLabel.md) |  | 

### Return type

[**ClusterTenant**](clusterTenant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDSCProfile

> ClusterDSCProfileList ListDSCProfile(ctx).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

List DSCProfile objects

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
    resp, r, err := api_client.ClusterV1Api.ListDSCProfile(context.Background()).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV1Api.ListDSCProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDSCProfile`: ClusterDSCProfileList
    fmt.Fprintf(os.Stdout, "Response from `ClusterV1Api.ListDSCProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDSCProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**ClusterDSCProfileList**](clusterDSCProfileList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDistributedServiceCard

> ClusterDistributedServiceCardList ListDistributedServiceCard(ctx).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

List DistributedServiceCard objects

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
    resp, r, err := api_client.ClusterV1Api.ListDistributedServiceCard(context.Background()).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV1Api.ListDistributedServiceCard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDistributedServiceCard`: ClusterDistributedServiceCardList
    fmt.Fprintf(os.Stdout, "Response from `ClusterV1Api.ListDistributedServiceCard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDistributedServiceCardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**ClusterDistributedServiceCardList**](clusterDistributedServiceCardList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHost

> ClusterHostList ListHost(ctx).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

List Host objects

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
    resp, r, err := api_client.ClusterV1Api.ListHost(context.Background()).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV1Api.ListHost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListHost`: ClusterHostList
    fmt.Fprintf(os.Stdout, "Response from `ClusterV1Api.ListHost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**ClusterHostList**](clusterHostList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNode

> ClusterNodeList ListNode(ctx).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

List Node objects

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
    resp, r, err := api_client.ClusterV1Api.ListNode(context.Background()).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV1Api.ListNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNode`: ClusterNodeList
    fmt.Fprintf(os.Stdout, "Response from `ClusterV1Api.ListNode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**ClusterNodeList**](clusterNodeList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTenant

> ClusterTenantList ListTenant(ctx).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

List Tenant objects

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
    resp, r, err := api_client.ClusterV1Api.ListTenant(context.Background()).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV1Api.ListTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTenant`: ClusterTenantList
    fmt.Fprintf(os.Stdout, "Response from `ClusterV1Api.ListTenant`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**ClusterTenantList**](clusterTenantList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Restore

> ClusterSnapshotRestore Restore(ctx).Body(body).Execute()

Restore Configuration

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
    body := *openapiclient.NewClusterSnapshotRestore() // ClusterSnapshotRestore | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterV1Api.Restore(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV1Api.Restore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Restore`: ClusterSnapshotRestore
    fmt.Fprintf(os.Stdout, "Response from `ClusterV1Api.Restore`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRestoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ClusterSnapshotRestore**](ClusterSnapshotRestore.md) |  | 

### Return type

[**ClusterSnapshotRestore**](clusterSnapshotRestore.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Save

> ClusterConfigurationSnapshot Save(ctx).Body(body).Execute()

Perform a Configuation Snapshot

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
    body := *openapiclient.NewClusterConfigurationSnapshotRequest() // ClusterConfigurationSnapshotRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterV1Api.Save(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV1Api.Save``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Save`: ClusterConfigurationSnapshot
    fmt.Fprintf(os.Stdout, "Response from `ClusterV1Api.Save`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSaveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ClusterConfigurationSnapshotRequest**](ClusterConfigurationSnapshotRequest.md) |  | 

### Return type

[**ClusterConfigurationSnapshot**](clusterConfigurationSnapshot.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCluster

> ClusterCluster UpdateCluster(ctx).Body(body).Execute()

Update Cluster object

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
    body := *openapiclient.NewClusterCluster() // ClusterCluster | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterV1Api.UpdateCluster(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV1Api.UpdateCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCluster`: ClusterCluster
    fmt.Fprintf(os.Stdout, "Response from `ClusterV1Api.UpdateCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ClusterCluster**](ClusterCluster.md) |  | 

### Return type

[**ClusterCluster**](clusterCluster.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfigurationSnapshot

> ClusterConfigurationSnapshot UpdateConfigurationSnapshot(ctx).Body(body).Execute()

Update ConfigurationSnapshot object

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
    body := *openapiclient.NewClusterConfigurationSnapshot() // ClusterConfigurationSnapshot | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterV1Api.UpdateConfigurationSnapshot(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV1Api.UpdateConfigurationSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConfigurationSnapshot`: ClusterConfigurationSnapshot
    fmt.Fprintf(os.Stdout, "Response from `ClusterV1Api.UpdateConfigurationSnapshot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConfigurationSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ClusterConfigurationSnapshot**](ClusterConfigurationSnapshot.md) |  | 

### Return type

[**ClusterConfigurationSnapshot**](clusterConfigurationSnapshot.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDSCProfile

> ClusterDSCProfile UpdateDSCProfile(ctx, oName).Body(body).Execute()

Update DSCProfile object

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
    body := *openapiclient.NewClusterDSCProfile() // ClusterDSCProfile | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterV1Api.UpdateDSCProfile(context.Background(), oName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV1Api.UpdateDSCProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDSCProfile`: ClusterDSCProfile
    fmt.Fprintf(os.Stdout, "Response from `ClusterV1Api.UpdateDSCProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDSCProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ClusterDSCProfile**](ClusterDSCProfile.md) |  | 

### Return type

[**ClusterDSCProfile**](clusterDSCProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDistributedServiceCard

> ClusterDistributedServiceCard UpdateDistributedServiceCard(ctx, oName).Body(body).Execute()

Update DistributedServiceCard object

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
    body := *openapiclient.NewClusterDistributedServiceCard() // ClusterDistributedServiceCard | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterV1Api.UpdateDistributedServiceCard(context.Background(), oName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV1Api.UpdateDistributedServiceCard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDistributedServiceCard`: ClusterDistributedServiceCard
    fmt.Fprintf(os.Stdout, "Response from `ClusterV1Api.UpdateDistributedServiceCard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDistributedServiceCardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ClusterDistributedServiceCard**](ClusterDistributedServiceCard.md) |  | 

### Return type

[**ClusterDistributedServiceCard**](clusterDistributedServiceCard.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHost

> ClusterHost UpdateHost(ctx, oName).Body(body).Execute()

Update Host object

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
    body := *openapiclient.NewClusterHost() // ClusterHost | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterV1Api.UpdateHost(context.Background(), oName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV1Api.UpdateHost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateHost`: ClusterHost
    fmt.Fprintf(os.Stdout, "Response from `ClusterV1Api.UpdateHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ClusterHost**](ClusterHost.md) |  | 

### Return type

[**ClusterHost**](clusterHost.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLicense

> ClusterLicense UpdateLicense(ctx).Body(body).Execute()

Update License object

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
    body := *openapiclient.NewClusterLicense() // ClusterLicense | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterV1Api.UpdateLicense(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV1Api.UpdateLicense``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLicense`: ClusterLicense
    fmt.Fprintf(os.Stdout, "Response from `ClusterV1Api.UpdateLicense`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLicenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ClusterLicense**](ClusterLicense.md) |  | 

### Return type

[**ClusterLicense**](clusterLicense.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNode

> ClusterNode UpdateNode(ctx, oName).Body(body).Execute()

Update Node object

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
    body := *openapiclient.NewClusterNode() // ClusterNode | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterV1Api.UpdateNode(context.Background(), oName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV1Api.UpdateNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNode`: ClusterNode
    fmt.Fprintf(os.Stdout, "Response from `ClusterV1Api.UpdateNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ClusterNode**](ClusterNode.md) |  | 

### Return type

[**ClusterNode**](clusterNode.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTLSConfig

> ClusterCluster UpdateTLSConfig(ctx).Body(body).Execute()

Update TLS Configuration for cluster

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
    body := *openapiclient.NewClusterUpdateTLSConfigRequest() // ClusterUpdateTLSConfigRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterV1Api.UpdateTLSConfig(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV1Api.UpdateTLSConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTLSConfig`: ClusterCluster
    fmt.Fprintf(os.Stdout, "Response from `ClusterV1Api.UpdateTLSConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTLSConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ClusterUpdateTLSConfigRequest**](ClusterUpdateTLSConfigRequest.md) |  | 

### Return type

[**ClusterCluster**](clusterCluster.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTenant

> ClusterTenant UpdateTenant(ctx, oName).Body(body).Execute()

Update Tenant object

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
    body := *openapiclient.NewClusterTenant() // ClusterTenant | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterV1Api.UpdateTenant(context.Background(), oName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV1Api.UpdateTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTenant`: ClusterTenant
    fmt.Fprintf(os.Stdout, "Response from `ClusterV1Api.UpdateTenant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ClusterTenant**](ClusterTenant.md) |  | 

### Return type

[**ClusterTenant**](clusterTenant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVersion

> ClusterVersion UpdateVersion(ctx).Body(body).Execute()

Update Version object

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
    body := *openapiclient.NewClusterVersion() // ClusterVersion | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ClusterV1Api.UpdateVersion(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV1Api.UpdateVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVersion`: ClusterVersion
    fmt.Fprintf(os.Stdout, "Response from `ClusterV1Api.UpdateVersion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ClusterVersion**](ClusterVersion.md) |  | 

### Return type

[**ClusterVersion**](clusterVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WatchCluster

> ClusterAutoMsgClusterWatchHelper WatchCluster(ctx).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

Watch Cluster objects. Supports WebSockets or HTTP long poll

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
    resp, r, err := api_client.ClusterV1Api.WatchCluster(context.Background()).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV1Api.WatchCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WatchCluster`: ClusterAutoMsgClusterWatchHelper
    fmt.Fprintf(os.Stdout, "Response from `ClusterV1Api.WatchCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWatchClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**ClusterAutoMsgClusterWatchHelper**](clusterAutoMsgClusterWatchHelper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WatchConfigurationSnapshot

> ClusterAutoMsgConfigurationSnapshotWatchHelper WatchConfigurationSnapshot(ctx).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

Watch ConfigurationSnapshot objects. Supports WebSockets or HTTP long poll

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
    resp, r, err := api_client.ClusterV1Api.WatchConfigurationSnapshot(context.Background()).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV1Api.WatchConfigurationSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WatchConfigurationSnapshot`: ClusterAutoMsgConfigurationSnapshotWatchHelper
    fmt.Fprintf(os.Stdout, "Response from `ClusterV1Api.WatchConfigurationSnapshot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWatchConfigurationSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**ClusterAutoMsgConfigurationSnapshotWatchHelper**](clusterAutoMsgConfigurationSnapshotWatchHelper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WatchDSCProfile

> ClusterAutoMsgDSCProfileWatchHelper WatchDSCProfile(ctx).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

Watch DSCProfile objects. Supports WebSockets or HTTP long poll

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
    resp, r, err := api_client.ClusterV1Api.WatchDSCProfile(context.Background()).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV1Api.WatchDSCProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WatchDSCProfile`: ClusterAutoMsgDSCProfileWatchHelper
    fmt.Fprintf(os.Stdout, "Response from `ClusterV1Api.WatchDSCProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWatchDSCProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**ClusterAutoMsgDSCProfileWatchHelper**](clusterAutoMsgDSCProfileWatchHelper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WatchDistributedServiceCard

> ClusterAutoMsgDistributedServiceCardWatchHelper WatchDistributedServiceCard(ctx).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

Watch DistributedServiceCard objects. Supports WebSockets or HTTP long poll

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
    resp, r, err := api_client.ClusterV1Api.WatchDistributedServiceCard(context.Background()).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV1Api.WatchDistributedServiceCard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WatchDistributedServiceCard`: ClusterAutoMsgDistributedServiceCardWatchHelper
    fmt.Fprintf(os.Stdout, "Response from `ClusterV1Api.WatchDistributedServiceCard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWatchDistributedServiceCardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**ClusterAutoMsgDistributedServiceCardWatchHelper**](clusterAutoMsgDistributedServiceCardWatchHelper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WatchHost

> ClusterAutoMsgHostWatchHelper WatchHost(ctx).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

Watch Host objects. Supports WebSockets or HTTP long poll

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
    resp, r, err := api_client.ClusterV1Api.WatchHost(context.Background()).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV1Api.WatchHost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WatchHost`: ClusterAutoMsgHostWatchHelper
    fmt.Fprintf(os.Stdout, "Response from `ClusterV1Api.WatchHost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWatchHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**ClusterAutoMsgHostWatchHelper**](clusterAutoMsgHostWatchHelper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WatchNode

> ClusterAutoMsgNodeWatchHelper WatchNode(ctx).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

Watch Node objects. Supports WebSockets or HTTP long poll

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
    resp, r, err := api_client.ClusterV1Api.WatchNode(context.Background()).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV1Api.WatchNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WatchNode`: ClusterAutoMsgNodeWatchHelper
    fmt.Fprintf(os.Stdout, "Response from `ClusterV1Api.WatchNode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWatchNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**ClusterAutoMsgNodeWatchHelper**](clusterAutoMsgNodeWatchHelper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WatchTenant

> ClusterAutoMsgTenantWatchHelper WatchTenant(ctx).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

Watch Tenant objects. Supports WebSockets or HTTP long poll

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
    resp, r, err := api_client.ClusterV1Api.WatchTenant(context.Background()).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV1Api.WatchTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WatchTenant`: ClusterAutoMsgTenantWatchHelper
    fmt.Fprintf(os.Stdout, "Response from `ClusterV1Api.WatchTenant`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWatchTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**ClusterAutoMsgTenantWatchHelper**](clusterAutoMsgTenantWatchHelper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WatchVersion

> ClusterAutoMsgVersionWatchHelper WatchVersion(ctx).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()

Watch Version objects. Supports WebSockets or HTTP long poll

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
    resp, r, err := api_client.ClusterV1Api.WatchVersion(context.Background()).OName(oName).OCreationTime(oCreationTime).FieldChangeSelector(fieldChangeSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterV1Api.WatchVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WatchVersion`: ClusterAutoMsgVersionWatchHelper
    fmt.Fprintf(os.Stdout, "Response from `ClusterV1Api.WatchVersion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWatchVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oName** | **string** | Name of the object, unique within a Namespace for scoped objects. Must start and end with alpha numeric and can have alphanumeric, -, _, . Length of string should be between 2 and 64. | 
 **oCreationTime** | **time.Time** | CreationTime is the creation time of the object. System generated and updated, not updatable by user. | 
 **fieldChangeSelector** | **[]string** | FieldChangeSelector specifies to generate a watch notification on change in field(s) specified. | 

### Return type

[**ClusterAutoMsgVersionWatchHelper**](clusterAutoMsgVersionWatchHelper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

